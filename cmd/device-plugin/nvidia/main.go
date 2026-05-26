/*
Copyright 2024 The HAMi Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"flag"
	"fmt"
	"os"

	nvinfo "github.com/NVIDIA/go-nvlib/pkg/nvlib/info"
	spec "github.com/NVIDIA/k8s-device-plugin/api/config/v1"
	cli "github.com/urfave/cli/v2"
	"k8s.io/klog/v2"
	kubeletdevicepluginv1beta1 "k8s.io/kubelet/pkg/apis/deviceplugin/v1beta1"

	"github.com/Project-HAMi/HAMi/pkg/device-plugin/nvidiadevice/nvinternal/info"
	"github.com/Project-HAMi/HAMi/pkg/device-plugin/nvidiadevice/nvinternal/plugin"
	flagutil "github.com/Project-HAMi/HAMi/pkg/util/flag"
)

type options struct {
	flags         []cli.Flag
	configFile    string
	kubeletSocket string
}

func main() {
	c := cli.NewApp()
	o := &options{}
	c.Name = "NVIDIA Device Plugin"
	c.Usage = "NVIDIA device plugin for Kubernetes"
	c.Action = func(ctx *cli.Context) error {
		flagutil.PrintCliFlags(ctx)
		return start(ctx, o)
	}
	c.Commands = []*cli.Command{
		{
			Name:  "version",
			Usage: "Show the version of NVIDIA Device Plugin",
			Action: func(c *cli.Context) error {
				fmt.Printf("%s version: %s\n", c.App.Name, info.GetVersionString())
				return nil
			},
		},
	}

	flagset := flag.NewFlagSet("klog", flag.ExitOnError)
	klog.InitFlags(flagset)

	c.Before = func(ctx *cli.Context) error {
		logLevel := ctx.Int("v")
		if err := flagset.Set("v", fmt.Sprintf("%d", logLevel)); err != nil {
			return err
		}
		return nil
	}

	c.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "mig-strategy",
			Value:   spec.MigStrategyNone,
			Usage:   "the desired strategy for exposing MIG devices on GPUs that support it:\n\t\t[none | single | mixed]",
			EnvVars: []string{"MIG_STRATEGY"},
		},
		&cli.BoolFlag{
			Name:    "fail-on-init-error",
			Value:   true,
			Usage:   "fail the plugin if an error is encountered during initialization, otherwise block indefinitely",
			EnvVars: []string{"FAIL_ON_INIT_ERROR"},
		},
		&cli.StringFlag{
			Name:    "nvidia-driver-root",
			Value:   "/",
			Usage:   "the root path for the NVIDIA driver installation (typical values are '/' or '/run/nvidia/driver')",
			EnvVars: []string{"NVIDIA_DRIVER_ROOT"},
		},
		&cli.StringFlag{
			Name:    "dev-root",
			Aliases: []string{"nvidia-dev-root"},
			Usage:   "the root path for the NVIDIA device nodes on the host (typical values are '/' or '/run/nvidia/driver')",
			EnvVars: []string{"NVIDIA_DEV_ROOT"},
		},
		&cli.BoolFlag{
			Name:    "pass-device-specs",
			Value:   false,
			Usage:   "pass the list of DeviceSpecs to the kubelet on Allocate()",
			EnvVars: []string{"PASS_DEVICE_SPECS"},
		},
		&cli.StringSliceFlag{
			Name:    "device-list-strategy",
			Value:   cli.NewStringSlice(string(spec.DeviceListStrategyEnvVar)),
			Usage:   "the desired strategy for passing the device list to the underlying runtime:\n\t\t[envvar | volume-mounts | cdi-annotations]",
			EnvVars: []string{"DEVICE_LIST_STRATEGY"},
		},
		&cli.StringFlag{
			Name:    "device-id-strategy",
			Value:   spec.DeviceIDStrategyUUID,
			Usage:   "the desired strategy for passing device IDs to the underlying runtime:\n\t\t[uuid | index]",
			EnvVars: []string{"DEVICE_ID_STRATEGY"},
		},
		&cli.BoolFlag{
			Name:    "gdrcopy-enabled",
			Usage:   "ensure that containers that request NVIDIA GPU resources are started with GDRCopy support",
			EnvVars: []string{"GDRCOPY_ENABLED"},
		},
		&cli.BoolFlag{
			Name:    "gds-enabled",
			Usage:   "ensure that containers are started with NVIDIA_GDS=enabled",
			EnvVars: []string{"GDS_ENABLED"},
		},
		&cli.BoolFlag{
			Name:    "mofed-enabled",
			Usage:   "ensure that containers are started with NVIDIA_MOFED=enabled",
			EnvVars: []string{"MOFED_ENABLED"},
		},
		&cli.StringFlag{
			Name:        "kubelet-socket",
			Value:       kubeletdevicepluginv1beta1.KubeletSocket,
			Usage:       "specify the socket for communicating with the kubelet; if this is empty, no connection with the kubelet is attempted",
			Destination: &o.kubeletSocket,
			EnvVars:     []string{"KUBELET_SOCKET"},
		},
		&cli.StringFlag{
			Name:        "config-file",
			Usage:       "the path to a config file as an alternative to command line options or environment variables",
			Destination: &o.configFile,
			EnvVars:     []string{"CONFIG_FILE"},
		},
		&cli.StringFlag{
			Name:    "cdi-annotation-prefix",
			Value:   spec.DefaultCDIAnnotationPrefix,
			Usage:   "the prefix to use for CDI container annotation keys",
			EnvVars: []string{"CDI_ANNOTATION_PREFIX"},
		},
		&cli.StringFlag{
			Name:    "nvidia-cdi-hook-path",
			Aliases: []string{"nvidia-ctk-path"},
			Value:   spec.DefaultNvidiaCTKPath,
			Usage:   "the path to use for NVIDIA CDI hooks in the generated CDI specification",
			EnvVars: []string{"NVIDIA_CDI_HOOK_PATH", "NVIDIA_CTK_PATH"},
		},
		&cli.StringFlag{
			Name:    "driver-root-ctr-path",
			Aliases: []string{"container-driver-root"},
			Value:   spec.DefaultContainerDriverRoot,
			Usage:   "the path where the NVIDIA driver root is mounted in the container; used for generating CDI specifications",
			EnvVars: []string{"DRIVER_ROOT_CTR_PATH", "CONTAINER_DRIVER_ROOT"},
		},
		&cli.StringFlag{
			Name:    "device-discovery-strategy",
			Value:   "auto",
			Usage:   "the strategy to use to discover devices: 'auto', 'nvml', or 'tegra'",
			EnvVars: []string{"DEVICE_DISCOVERY_STRATEGY"},
		},
		&cli.IntSliceFlag{
			Name:    "imex-channel-ids",
			Usage:   "A list of IMEX channels to inject.",
			EnvVars: []string{"IMEX_CHANNEL_IDS"},
		},
		&cli.BoolFlag{
			Name:    "imex-required",
			Usage:   "The specified IMEX channels are required",
			EnvVars: []string{"IMEX_REQUIRED"},
		},
		&cli.IntFlag{
			Name:  "v",
			Usage: "number for the log level verbosity",
			Value: 0,
		},
	}
	c.Flags = append(c.Flags, addFlags()...)
	o.flags = c.Flags
	err := c.Run(os.Args)
	if err != nil {
		klog.Error(err)
		os.Exit(1)
	}
}

func validateFlags(infolib nvinfo.Interface, config *spec.Config) error {
	_ = "STUB: not implemented"
	return nil
}

func loadConfig(c *cli.Context, flags []cli.Flag) (*spec.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func start(c *cli.Context, o *options) error { _ = "STUB: not implemented"; return nil }

/*Loading config files*/

// If we are restarting, stop plugins from previous run.

// Start an infinite loop, waiting for several indicators to either log
// some messages, trigger a restart of the plugins, or exit the program.

// If the restart timeout has expired, then restart the plugins

// Detect a kubelet restart by watching for a newly created
// 'kubeletdevicepluginv1beta1.KubeletSocket' file. When this occurs, restart this loop,
// restarting all of the plugins in the process.

// Watch for any other fs errors and log them.

// Watch for any signals from the OS. On SIGHUP, restart this loop,
// restarting all of the plugins in the process. On all other
// signals, exit the loop and exit the program.

func startPlugins(c *cli.Context, o *options) ([]plugin.Interface, bool, error) {
	_ = "STUB: not implemented"
	// Load the configuration file
	return nil, false, nil
}

/*Loading config files*/
//fmt.Println("NodeName=", config.NodeName)

// We construct an NVML library specifying the path to libnvidia-ml.so.1
// explicitly so that we don't have to rely on the library path.

// Update the configuration file with default resources.

// Print the config to the output.

// Get the set of plugins.

// Loop through all plugins, starting them if they have any devices
// to serve. If even one plugin fails to start properly, try
// starting them all again.

// Just continue if there are no devices to serve for plugin p.

// Start the gRPC server for plugin p and connect it with the kubelet.

func stopPlugins(plugins []plugin.Interface) error { _ = "STUB: not implemented"; return nil }

// disableResourceRenamingInConfig temporarily disable the resource renaming feature of the plugin.
// We plan to reeenable this feature in a future release.
func disableResourceRenamingInConfig(config *spec.Config) {
	_ = "STUB: not implemented"
	// Disable resource renaming through config.Resource
	return
}

// Disable renaming / device selection in Sharing.TimeSlicing.Resources
