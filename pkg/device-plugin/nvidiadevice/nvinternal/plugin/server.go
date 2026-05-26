/*
 * SPDX-License-Identifier: Apache-2.0
 *
 * The HAMi Contributors require contributions made to
 * this file be licensed under the Apache-2.0 license or a
 * compatible open source license.
 */

/*
 * Licensed to NVIDIA CORPORATION under one or more contributor
 * license agreements. See the NOTICE file distributed with
 * this work for additional information regarding copyright
 * ownership. NVIDIA CORPORATION licenses this file to you under
 * the Apache License, Version 2.0 (the "License"); you may
 * not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

/*
 * Modifications Copyright The HAMi Authors. See
 * GitHub history for details.
 */

package plugin

import (
	"context"
	"os"
	"sync"
	"time"

	spec "github.com/NVIDIA/k8s-device-plugin/api/config/v1"
	"google.golang.org/grpc"
	kubeletdevicepluginv1beta1 "k8s.io/kubelet/pkg/apis/deviceplugin/v1beta1"

	"github.com/Project-HAMi/HAMi/pkg/device"
	"github.com/Project-HAMi/HAMi/pkg/device-plugin/nvidiadevice/nvinternal/cdi"
	"github.com/Project-HAMi/HAMi/pkg/device-plugin/nvidiadevice/nvinternal/imex"
	"github.com/Project-HAMi/HAMi/pkg/device-plugin/nvidiadevice/nvinternal/rm"
	"github.com/Project-HAMi/HAMi/pkg/device/nvidia"
	"github.com/Project-HAMi/HAMi/pkg/scheduler/config"
	"github.com/Project-HAMi/HAMi/pkg/util"
)

// Constants for use by the 'volume-mounts' device list strategy
const (
	deviceListAsVolumeMountsHostPath          = "/dev/null"
	deviceListAsVolumeMountsContainerPathRoot = "/var/run/nvidia-container-devices"
	NodeLockNvidia                            = "hami.io/mutex.lock"
	ConfigFilePath                            = "/config/config.json"
	deviceListEnvVar                          = "NVIDIA_VISIBLE_DEVICES"
)

var (
	hostHookPath                 string
	ConfigFile                   *string
	getPendingPod                = util.GetPendingPod
	enableGetPreferredAllocation bool
)

func init() {
	hostHookPath, _ = os.LookupEnv("HOOK_PATH")
}

// NvidiaDevicePlugin implements the Kubernetes device plugin API
type NvidiaDevicePlugin struct {
	kubeletdevicepluginv1beta1.UnimplementedDevicePluginServer

	ctx                  context.Context
	rm                   rm.ResourceManager
	config               *nvidia.DeviceConfig
	deviceListEnvvar     string
	deviceListStrategies spec.DeviceListStrategies
	socket               string
	schedulerConfig      nvidia.NvidiaConfig

	applyMutex                 sync.Mutex
	disableHealthChecks        chan bool
	ackDisableHealthChecks     chan bool
	disableWatchAndRegister    chan bool
	ackDisableWatchAndRegister chan bool

	cdiHandler          cdi.Interface
	cdiAnnotationPrefix string

	operatingMode string
	migCurrent    nvidia.MigPartedSpec
	deviceCache   string

	imexChannels imex.Channels

	server *grpc.Server
	health chan *rm.Device
	stop   chan any
}

func readFromConfigFile(sConfig *nvidia.NvidiaConfig, path string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func LoadNvidiaDevicePluginConfig() (*config.Config, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

// getPluginSocketPath returns the socket to use for the specified resource.
func getPluginSocketPath(resource spec.ResourceName) string { _ = "STUB: not implemented"; return "" }

// NewNvidiaDevicePlugin returns an initialized NvidiaDevicePlugin
func (o *options) devicePluginForResource(ctx context.Context, nvconfig *nvidia.DeviceConfig, resourceManager rm.ResourceManager, sConfig *config.Config, mode string) (Interface, error) {
	_ = "STUB: not implemented"
	return *new(Interface), nil
}

// Initialize devices with configuration

// These will be reinitialized every
// time the plugin server is restarted.

func (plugin *NvidiaDevicePlugin) initialize() { _ = "STUB: not implemented"; return }

func (plugin *NvidiaDevicePlugin) cleanup() { _ = "STUB: not implemented"; return }

// Devices returns the full set of devices associated with the plugin.
func (plugin *NvidiaDevicePlugin) Devices() rm.Devices {
	_ = "STUB: not implemented"
	return *new(rm.Devices)
}

// Start starts the gRPC server, registers the device plugin with the Kubelet,
// and starts the device healthchecks.
func (plugin *NvidiaDevicePlugin) Start(kubeletSocket string) error {
	_ = "STUB: not implemented"
	return nil
}

// Prepare the lock file sub directory.Due to the sequence of startup processes, both the device plugin
// and the vGPU monitor should attempt to create this directory by default to ensure its creation.

// If the temporary lock file still exists, it may be a leftover from the last incomplete mig  application process.
// Delete the temporary lock file to make sure vgpu monitor can start.

// Stop stops the gRPC server.
func (plugin *NvidiaDevicePlugin) Stop() error { _ = "STUB: not implemented"; return nil }

// Serve starts the gRPC server of the device plugin.
func (plugin *NvidiaDevicePlugin) Serve() error { _ = "STUB: not implemented"; return nil }

// restart if it has not been too often
// i.e. if server has crashed more than 5 times and it didn't last more than one hour each time

// quit

// it has been one hour since the last crash.. reset the count
// to reflect on the frequency

// Wait for server to start by launching a blocking connexion

// Register registers the device plugin for the given resourceName with Kubelet.
func (plugin *NvidiaDevicePlugin) Register(kubeletSocket string) error {
	_ = "STUB: not implemented"
	return nil
}

// GetDevicePluginOptions returns the values of the optional settings for this plugin
func (plugin *NvidiaDevicePlugin) GetDevicePluginOptions(context.Context, *kubeletdevicepluginv1beta1.Empty) (*kubeletdevicepluginv1beta1.DevicePluginOptions, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ListAndWatch lists devices and update that list according to the health status
func (plugin *NvidiaDevicePlugin) ListAndWatch(e *kubeletdevicepluginv1beta1.Empty, s kubeletdevicepluginv1beta1.DevicePlugin_ListAndWatchServer) error {
	_ = "STUB: not implemented"
	return nil
}

// FIXME: there is no way to recover from the Unhealthy state.

// GetPreferredAllocation returns the preferred allocation from the set of devices specified in the request
func (plugin *NvidiaDevicePlugin) GetPreferredAllocation(ctx context.Context, r *kubeletdevicepluginv1beta1.PreferredAllocationRequest) (*kubeletdevicepluginv1beta1.PreferredAllocationResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Filter out empty annotations to match kubelet's ContainerRequests order.
// Kubelet only sends requests for containers that need GPUs, but annotations
// include all containers (init + regular), some of which may be empty.

func (plugin *NvidiaDevicePlugin) selectPreferredDeviceIDsFromAnnotatedDevices(available, required []string, desired device.ContainerDevices, allocationSize int) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func physicalDeviceID(id string) string { _ = "STUB: not implemented"; return "" }

// Handle MIG format: GPU-UUID[tidx-idx] -> GPU-UUID

// Handle virtual device format: GPU-UUID-N -> GPU-UUID
// NVIDIA GPU UUID has exactly 5 dashes (GPU-xxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx)
// Virtual devices append "-N" suffix, resulting in exactly 6 dashes

func (plugin *NvidiaDevicePlugin) alignContainerDevicesWithAllocatedIDs(devreq device.ContainerDevices, deviceIDs []string) (device.ContainerDevices, error) {
	_ = "STUB: not implemented"
	return *new(device.ContainerDevices), nil
}

// Allocate which return list of devices.
func (plugin *NvidiaDevicePlugin) Allocate(ctx context.Context, reqs *kubeletdevicepluginv1beta1.AllocateRequest) (*kubeletdevicepluginv1beta1.AllocateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nodelock.ReleaseNodeLock(nodename, NodeLockNvidia, current)

// If the devices being allocated are replicas, then (conditionally)
// error out if more than one resource is being allocated.

// if env existed but is set to false or can not be parsed, ignore

// only env existed and set to true, we mark it "found"

func (plugin *NvidiaDevicePlugin) getAllocateResponse(requestIds []string) (*kubeletdevicepluginv1beta1.ContainerAllocateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create an empty response that will be updated as required below.

// The following modifications are only made if at least one non-CDI device
// list strategy is selected.

// updateResponseForCDI updates the specified response for the given device IDs.
// This response contains the annotations required to trigger CDI injection in the container engine or nvidia-container-runtime.
func (plugin *NvidiaDevicePlugin) updateResponseForCDI(response *kubeletdevicepluginv1beta1.ContainerAllocateResponse, responseID string, deviceIDs ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func (plugin *NvidiaDevicePlugin) getCDIDeviceAnnotations(id string, devices ...string) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// update annotations if a custom CDI prefix is configured

// PreStartContainer is unimplemented for this plugin
func (plugin *NvidiaDevicePlugin) PreStartContainer(context.Context, *kubeletdevicepluginv1beta1.PreStartContainerRequest) (*kubeletdevicepluginv1beta1.PreStartContainerResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// dial establishes the gRPC communication with the registered device plugin.
func (plugin *NvidiaDevicePlugin) dial(unixSocketPath string, timeout time.Duration) (*grpc.ClientConn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:staticcheck  // TODO: Switch to grpc.NewClient

//nolint:staticcheck  // TODO: WithBlock is deprecated.

func (plugin *NvidiaDevicePlugin) uniqueDeviceIDsFromAnnotatedDeviceIDs(ids []string) []string {
	_ = "STUB: not implemented"
	return nil
}

// updateResponseForDeviceListEnvVar sets the environment variable for the requested devices.
func (plugin *NvidiaDevicePlugin) updateResponseForDeviceListEnvVar(response *kubeletdevicepluginv1beta1.ContainerAllocateResponse, deviceIDs ...string) {
	_ = "STUB: not implemented"
	return
}

// updateResponseForImexChannelsEnvVar sets the environment variable for the requested IMEX channels.
func (plugin *NvidiaDevicePlugin) updateResponseForImexChannelsEnvVar(response *kubeletdevicepluginv1beta1.ContainerAllocateResponse) {
	_ = "STUB: not implemented"
	return
}

// updateResponseForDeviceMounts sets the mounts required to request devices if volume mounts are used.
func (plugin *NvidiaDevicePlugin) updateResponseForDeviceMounts(response *kubeletdevicepluginv1beta1.ContainerAllocateResponse, deviceIDs ...string) {
	_ = "STUB: not implemented"
	return
}

func (plugin *NvidiaDevicePlugin) apiDeviceSpecs(devRoot string, ids []string) []*kubeletdevicepluginv1beta1.DeviceSpec {
	_ = "STUB: not implemented"
	return nil
}

// TODO: The HostPath property for a channel is not the correct value to use here.
// The `devRoot` there represents the devRoot in the current container when discovering devices
// and is set to "{{ .*config.Flags.Plugin.ContainerDriverRoot }}/dev".
// The devRoot in this context is the {{ .config.Flags.NvidiaDevRoot }} and defines the
// root for device nodes on the host. This is usually / or /run/nvidia/driver when the
// driver container is used.

func (plugin *NvidiaDevicePlugin) apiDevices() []*kubeletdevicepluginv1beta1.Device {
	_ = "STUB: not implemented"
	return nil
}

func (plugin *NvidiaDevicePlugin) processMigConfigs(migConfigs map[string]nvidia.MigConfigSpecSlice, deviceCount int) (nvidia.MigConfigSpecSlice, error) {
	_ = "STUB: not implemented"
	return *new(nvidia.MigConfigSpecSlice), nil
}
