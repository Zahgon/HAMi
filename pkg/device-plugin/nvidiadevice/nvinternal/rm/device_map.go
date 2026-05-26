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

package rm

import (
	"github.com/NVIDIA/go-nvlib/pkg/nvlib/device"
	"github.com/NVIDIA/go-nvlib/pkg/nvlib/info"
	"github.com/NVIDIA/go-nvml/pkg/nvml"

	spec "github.com/NVIDIA/k8s-device-plugin/api/config/v1"
)

type deviceMapBuilder struct {
	device.Interface
	migStrategy         *string
	resources           *spec.Resources
	replicatedResources *spec.ReplicatedResources

	newGPUDevice func(i int, gpu nvml.Device) (string, deviceInfo)
}

// DeviceMap stores a set of devices per resource name.
type DeviceMap map[spec.ResourceName]Devices

// NewDeviceMap creates a device map for the specified NVML library and config.
func NewDeviceMap(infolib info.Interface, devicelib device.Interface, config *spec.Config) (DeviceMap, error) {
	_ = "STUB: not implemented"
	return *new(DeviceMap), nil
}

// build builds a map of resource names to devices.
func (b *deviceMapBuilder) build() (DeviceMap, error) {
	_ = "STUB: not implemented"
	return *new(DeviceMap), nil
}

// buildDeviceMapFromConfigResources builds a map of resource names to devices from spec.Config.Resources
func (b *deviceMapBuilder) buildDeviceMapFromConfigResources() (DeviceMap, error) {
	_ = "STUB: not implemented"
	return *new(DeviceMap), nil
}

// buildGPUDeviceMap builds a map of resource names to GPU devices
func (b *deviceMapBuilder) buildGPUDeviceMap() (DeviceMap, error) {
	_ = "STUB: not implemented"
	return *new(DeviceMap), nil
}

// buildMigDeviceMap builds a map of resource names to MIG devices
func (b *deviceMapBuilder) buildMigDeviceMap() (DeviceMap, error) {
	_ = "STUB: not implemented"
	return *new(DeviceMap), nil
}

// assertAllMigDevicesAreValid ensures that each MIG-enabled device has at least one MIG device
// associated with it.
func (b *deviceMapBuilder) assertAllMigDevicesAreValid(uniform bool) error {
	_ = "STUB: not implemented"
	return nil
}

// setEntry sets the DeviceMap entry for the specified resource
func (d DeviceMap) setEntry(name spec.ResourceName, index string, device deviceInfo) error {
	_ = "STUB: not implemented"
	return nil
}

// insert adds the specified device to the device map
func (d DeviceMap) insert(name spec.ResourceName, dev *Device) { _ = "STUB: not implemented"; return }

// merge merges two devices maps
func (d DeviceMap) merge(o DeviceMap) { _ = "STUB: not implemented"; return }

// isEmpty checks whether a device map is empty
func (d DeviceMap) isEmpty() bool { _ = "STUB: not implemented"; return false }

// getIDsOfDevicesToReplicate returns a list of dervice IDs that we want to replicate.
func (d DeviceMap) getIDsOfDevicesToReplicate(r *spec.ReplicatedResource) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If all devices for this resource type are to be replicated.

// If a specific number of devices for this resource type are to be replicated.

// If a specific set of devices for this resource type are to be replicated.

// updateDeviceMapWithReplicas returns an updated map of resource names to devices with replica
// information from the active replicated resources config.
func updateDeviceMapWithReplicas(replicatedResources *spec.ReplicatedResources, oDevices DeviceMap) (DeviceMap, error) {
	_ = "STUB: not implemented"
	return *new(DeviceMap), nil
}

// Begin by walking replicatedResources.Resources and building a map of just the resource names.

// Copy over all devices from oDevices without a resource reference in TimeSlicing.Resources.

// Walk shared Resources and update devices in the device map as appropriate.

// Get the IDs of the devices we want to replicate from oDevices

// Skip any resources not matched in oDevices

// Add any devices we don't want replicated directly into the device map.

// Create replicated devices add them to the device map.
// Rename the resource for replicated devices as requested.
