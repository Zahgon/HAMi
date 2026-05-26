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
	spec "github.com/NVIDIA/k8s-device-plugin/api/config/v1"
)

const (
	tegraDeviceName = "tegra"
)

// buildTegraDeviceMap creates a DeviceMap for the tegra devices in the sytesm.
// NOTE: At present only a single tegra device is expected.
func buildTegraDeviceMap(config *spec.Config) (DeviceMap, error) {
	_ = "STUB: not implemented"
	return *new(DeviceMap), nil
}

type tegraDevice struct{}

var _ deviceInfo = (*tegraDevice)(nil)

// GetUUID returns the UUID of the tegra device.
// TODO: This is currently hardcoded to `tegra`
func (d *tegraDevice) GetUUID() (string, error) {
	_ = "STUB: not implemented"
	return "",

		// GetPaths returns the paths for a tegra device.
		// A tegra device does not have paths associated with it.
		nil
}

func (d *tegraDevice) GetPaths() ([]string, error) {
	_ = "STUB: not implemented"

	// GetNumaNode always returns unsupported for a Tegra device
	return nil, nil
}

func (d *tegraDevice) GetNumaNode() (bool, int, error) {
	_ = "STUB: not implemented"
	return false,

		// GetTotalMemory is unsupported for a Tegra device.
		0, nil
}

func (d *tegraDevice) GetTotalMemory() (uint64, error) {
	_ = "STUB: not implemented"

	// GetComputeCapability is unimplemented for a Tegra device.
	return 0, nil
}

func (d *tegraDevice) GetComputeCapability() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
