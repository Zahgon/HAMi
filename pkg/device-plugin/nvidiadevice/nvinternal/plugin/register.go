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
	"github.com/NVIDIA/go-nvml/pkg/nvml"

	"github.com/Project-HAMi/HAMi/pkg/device"
)

// uint8Slice wraps an []uint8 with more functions.
type uint8Slice []uint8

// String turns a nil terminated uint8Slice into a string
func (s uint8Slice) String() string { _ = "STUB: not implemented"; return "" }

// GetNumaNode returns the NUMA node associated with the GPU device
func GetNumaNode(d nvml.Device) (bool, int, error) { _ = "STUB: not implemented"; return false, 0, nil }

// Discard leading zeros.

func (plugin *NvidiaDevicePlugin) getAPIDevices() *[]*device.DeviceInfo {
	_ = "STUB: not implemented"
	return nil
}

// Log mode-related warnings once per scan instead of per device

// Unified memory architecture GPUs (e.g., NVIDIA GB10/DGX Spark) don't support
// traditional memory queries. Use PreConfiguredDeviceMemory from config as fallback.

// when NVIDIA-Tesla P4, the device info is : ID:GPU-e290caca-2f0c-9582-acab-67a142b61ffa,Health:Healthy,Topology:nil,
// it is more reasonable to think of healthy as case-insensitive

// If the model name does not start with "NVIDIA ", we assume it is a virtual GPU or a non-NVIDIA device.
// This is to handle cases where the model name might not be in the expected format.

// RegisterInAnnotation scans devices and patches node annotations.
// Returns (changed, error) where changed indicates whether the annotation was actually updated.
func (plugin *NvidiaDevicePlugin) RegisterInAnnotation() (bool, error) {
	_ = "STUB: not implemented"
	return false, nil

	// Log compact summary at V(3); full details at V(5)
}

func (plugin *NvidiaDevicePlugin) WatchAndRegister(disableNVML <-chan bool, ackDisableWatchAndRegister chan<- bool) {
	_ = "STUB: not implemented"
	return
}

// when received disableNVML signal, stop the watch and register all the time

// when received enableNVML signal, start the watch and register again
