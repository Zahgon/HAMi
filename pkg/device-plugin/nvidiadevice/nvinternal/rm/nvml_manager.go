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

type nvmlResourceManager struct {
	resourceManager
	nvml nvml.Interface
}

var _ ResourceManager = (*nvmlResourceManager)(nil)

// NewNVMLResourceManagers returns a set of ResourceManagers, one for each NVML resource in 'config'.
func NewNVMLResourceManagers(infolib info.Interface, nvmllib nvml.Interface, devicelib device.Interface, config *spec.Config) ([]ResourceManager, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetPreferredAllocation runs an allocation algorithm over the inputs.
// The algorithm chosen is based both on the incoming set of available devices and various config settings.
func (r *nvmlResourceManager) GetPreferredAllocation(available, required []string, size int) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetDevicePaths returns the required and optional device nodes for the requested resources
func (r *nvmlResourceManager) GetDevicePaths(ids []string) []string {
	_ = "STUB: not implemented"
	return nil
}

// CheckHealth performs health checks on a set of devices, writing to the 'unhealthy' channel with any unhealthy devices
func (r *nvmlResourceManager) CheckHealth(stop <-chan interface{}, unhealthy chan<- *Device, disableNVML <-chan bool, ackDisableHealthChecks chan<- bool) error {
	_ = "STUB: not implemented"

	// first check if disableNVML channel signal is pass close into checkHealth function
	// if signal is pass close, return error "close signal received"
	return nil
}

// when disableNVML channel signal is pass restart, continue to restart checkHealth function
// when disableNVML channel signal is not pass restart, wait for restart signal

// getPreferredAllocation runs an allocation algorithm over the inputs.
// The algorithm chosen is based both on the incoming set of available devices and various config settings.
func (r *nvmlResourceManager) getPreferredAllocation(available, required []string, size int) ([]string, error) {
	_ = "STUB: not implemented"
	// If all of the available devices are full GPUs without replicas, then
	// calculate an aligned allocation across those devices.
	return nil, nil
}

// Otherwise, distribute them evenly across all replicated GPUs

// alignedAlloc shells out to the alignedAllocationPolicy that is set in
// order to calculate the preferred allocation.
func (r *nvmlResourceManager) alignedAlloc(available, required []string, size int) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
