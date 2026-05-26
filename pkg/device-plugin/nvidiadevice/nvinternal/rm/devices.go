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
	kubeletdevicepluginv1beta1 "k8s.io/kubelet/pkg/apis/deviceplugin/v1beta1"
)

// Device wraps kubeletdevicepluginv1beta1.Device with extra metadata and functions.
type Device struct {
	kubeletdevicepluginv1beta1.Device
	Paths             []string
	Index             string
	TotalMemory       uint64
	ComputeCapability string
	// Replicas stores the total number of times this device is replicated.
	// If this is 0 or 1 then the device is not shared.
	Replicas int
}

// deviceInfo defines the information the required to construct a Device
type deviceInfo interface {
	GetUUID() (string, error)
	GetPaths() ([]string, error)
	GetNumaNode() (bool, int, error)
	GetTotalMemory() (uint64, error)
	GetComputeCapability() (string, error)
}

// Devices wraps a map[string]*Device with some functions.
type Devices map[string]*Device

// AnnotatedID represents an ID with a replica number embedded in it.
type AnnotatedID string

// AnnotatedIDs can be used to treat a []string as a []AnnotatedID.
type AnnotatedIDs []string

// BuildDevice builds an rm.Device with the specified index and deviceInfo
func BuildDevice(index string, d deviceInfo) (*Device, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Contains checks if Devices contains devices matching all ids.
func (ds Devices) Contains(ids ...string) bool { _ = "STUB: not implemented"; return false }

// GetByID returns a reference to the device matching the specified ID (nil otherwise).
func (ds Devices) GetByID(id string) *Device {
	_ = "STUB: not implemented"

	// GetByIndex returns a reference to the device matching the specified Index (nil otherwise).
	return nil
}

func (ds Devices) GetByIndex(index string) *Device { _ = "STUB: not implemented"; return nil }

// Subset returns the subset of devices in Devices matching the provided ids.
// If any id in ids is not in Devices, then the subset that did match will be returned.
func (ds Devices) Subset(ids []string) Devices { _ = "STUB: not implemented"; return *new(Devices) }

// Difference returns the set of devices contained in ds but not in ods.
func (ds Devices) Difference(ods Devices) Devices { _ = "STUB: not implemented"; return *new(Devices) }

// GetIDs returns the ids from all devices in the Devices
func (ds Devices) GetIDs() []string { _ = "STUB: not implemented"; return nil }

// GetUUIDs returns the uuids associated with the Device in the set.
func (ds Devices) GetUUIDs() []string { _ = "STUB: not implemented"; return nil }

// GetPluginDevices returns the plugin Devices from all devices in the Devices
func (ds Devices) GetPluginDevices(count uint) []*kubeletdevicepluginv1beta1.Device {
	_ = "STUB: not implemented"
	return nil
}

// GetIndices returns the Indices from all devices in the Devices
func (ds Devices) GetIndices() []string { _ = "STUB: not implemented"; return nil }

// GetPaths returns the Paths from all devices in the Devices
func (ds Devices) GetPaths() []string { _ = "STUB: not implemented"; return nil }

// AlignedAllocationSupported checks whether all devices support an aligned allocation
func (ds Devices) AlignedAllocationSupported() bool { _ = "STUB: not implemented"; return false }

// AlignedAllocationSupported checks whether the device supports an aligned allocation
func (d Device) AlignedAllocationSupported() bool { _ = "STUB: not implemented"; return false }

// IsMigDevice returns checks whether d is a MIG device or not.
func (d Device) IsMigDevice() bool { _ = "STUB: not implemented"; return false }

// GetUUID returns the UUID for the device from the annotated ID.
func (d Device) GetUUID() string { _ = "STUB: not implemented"; return "" }

// NewAnnotatedID creates a new AnnotatedID from an ID and a replica number.
func NewAnnotatedID(id string, replica int) AnnotatedID {
	_ = "STUB: not implemented"
	return *new(AnnotatedID)
}

// HasAnnotations checks if an AnnotatedID has any annotations or not.
func (r AnnotatedID) HasAnnotations() bool { _ = "STUB: not implemented"; return false }

// Split splits a AnnotatedID into its ID and replica number parts.
func (r AnnotatedID) Split() (string, int) { _ = "STUB: not implemented"; return "", 0 }

// GetID returns just the ID part of the replicated ID
func (r AnnotatedID) GetID() string { _ = "STUB: not implemented"; return "" }

// AnyHasAnnotations checks if any ID has annotations or not.
func (rs AnnotatedIDs) AnyHasAnnotations() bool { _ = "STUB: not implemented"; return false }

// GetIDs returns just the ID parts of the annotated IDs as a []string
func (rs AnnotatedIDs) GetIDs() []string { _ = "STUB: not implemented"; return nil }
