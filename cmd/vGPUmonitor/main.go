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
	"context"

	"github.com/Project-HAMi/HAMi/pkg/monitor/nvidia"
	"github.com/Project-HAMi/HAMi/pkg/util"
	"github.com/Project-HAMi/HAMi/pkg/util/flag"
	"github.com/Project-HAMi/HAMi/pkg/version"

	"github.com/spf13/cobra"
	"k8s.io/klog/v2"
)

var (
	rootCmd = &cobra.Command{
		Use:   "vGPUmonitor",
		Short: "Hami vgpu vGPUmonitor",
		RunE: func(cmd *cobra.Command, args []string) error {
			flag.PrintPFlags(cmd.Flags())
			return start()
		},
	}
	metricsBindAddress string
	legacyMetrics      bool
)

func init() {
	rootCmd.Flags().SortFlags = false
	rootCmd.PersistentFlags().SortFlags = false
	rootCmd.Flags().AddGoFlagSet(util.InitKlogFlags())
	rootCmd.Flags().StringVar(&metricsBindAddress, "metrics-bind-address", ":9394", "The TCP address that the vGPUmonitor should bind to for serving prometheus metrics(e.g. 127.0.0.1:9394, :9394)")
	rootCmd.Flags().BoolVar(&legacyMetrics, "legacy-metrics", false, "Emit legacy metric names alongside new ones for backward compatibility")
	rootCmd.AddCommand(version.VersionCmd)
}

func start() error { _ = "STUB: not implemented"; return nil }

// Explicitly initialize

// Prepare the lock file sub directory.Due to the sequence of startup processes, both the device plugin
// and the vGPU monitor should attempt to create this directory by default to ensure its creation.

// Start the metrics service

// Start the monitoring and feedback service

// if err is temporary closed, wait for lock file to be removed

// Capture system signals

// Wait for all goroutines to complete

func initMetrics(ctx context.Context, containerLister *nvidia.ContainerLister) error {
	_ = "STUB: not implemented"
	return nil
}

//reg := prometheus.NewPedanticRegistry()

// Construct cluster managers. In real code, we would assign them to
// variables to then do something with them.

//NewClusterManager("ca", reg)

// Uncomment to add the standard process and Go metrics to the custom registry.
//reg.MustRegister(
//	prometheus.NewProcessCollector(prometheus.ProcessCollectorOpts{}),
//	prometheus.NewGoCollector(),
//)

// Starting the HTTP server in a goroutine

// Graceful shutdown on context cancellation

func main() {
	if err := rootCmd.Execute(); err != nil {
		klog.Fatal(err)
	}
}
