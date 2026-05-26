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
	"github.com/NVIDIA/go-nvml/pkg/nvml"
)

const (
	nvidiaProcDriverPath   = "/proc/driver/nvidia"
	nvidiaCapabilitiesPath = nvidiaProcDriverPath + "/capabilities"
)

// nvmlDevice wraps an nvml.Device with more functions.
type nvmlDevice struct {
	nvml.Device
}

// nvmlMigDevice allows for specific functions of nvmlDevice to be overridden.
type nvmlMigDevice nvmlDevice

var _ deviceInfo = (*nvmlDevice)(nil)
var _ deviceInfo = (*nvmlMigDevice)(nil)

func newNvmlGPUDevice(i int, gpu nvml.Device) (string, deviceInfo) {
	_ = "STUB: not implemented"
	return "", *new(deviceInfo)
}

func newWslGPUDevice(i int, gpu nvml.Device) (string, deviceInfo) {
	_ = "STUB: not implemented"
	return "", *new(deviceInfo)
}

func newMigDevice(i int, j int, mig nvml.Device) (string, nvmlMigDevice) {
	_ = "STUB: not implemented"
	return "", *new(nvmlMigDevice)
}

// GetUUID returns the UUID of the device
func (d nvmlDevice) GetUUID() (string, error) { _ = "STUB: not implemented"; return "", nil }

// GetUUID returns the UUID of the device
func (d nvmlMigDevice) GetUUID() (string, error) { _ = "STUB: not implemented"; return "", nil }

// GetPaths returns the paths for a GPU device
func (d nvmlDevice) GetPaths() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

// GetComputeCapability returns the CUDA Compute Capability for the device.
func (d nvmlDevice) GetComputeCapability() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// GetComputeCapability returns the CUDA Compute Capability for the device.
func (d nvmlMigDevice) GetComputeCapability() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// GetPaths returns the paths for a MIG device
func (d nvmlMigDevice) GetPaths() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

// GetNumaNode returns the NUMA node associated with the GPU device
func (d nvmlDevice) GetNumaNode() (bool, int, error) {
	_ = "STUB: not implemented"
	return false, 0, nil
}

// Discard leading zeros.

// GetNumaNode for a MIG device is the NUMA node of the parent device.
func (d nvmlMigDevice) GetNumaNode() (bool, int, error) {
	_ = "STUB: not implemented"
	return false, 0, nil
}

// GetTotalMemory returns the total memory available on the device.
func (d nvmlDevice) GetTotalMemory() (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

// GetTotalMemory returns the total memory available on the device.
func (d nvmlMigDevice) GetTotalMemory() (uint64, error) { _ = "STUB: not implemented"; return 0, nil }
