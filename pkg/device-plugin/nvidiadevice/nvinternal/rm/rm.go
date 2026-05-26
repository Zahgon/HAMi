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
	"errors"

	"github.com/NVIDIA/go-nvlib/pkg/nvlib/device"
	"github.com/NVIDIA/go-nvlib/pkg/nvlib/info"
	"github.com/NVIDIA/go-nvml/pkg/nvml"
	spec "github.com/NVIDIA/k8s-device-plugin/api/config/v1"
	"github.com/Project-HAMi/HAMi/pkg/device/nvidia"
)

// resourceManager forms the base type for specific resource manager implementations
type resourceManager struct {
	config   *spec.Config
	resource spec.ResourceName
	devices  Devices
}

// ResourceManager provides an interface for listing a set of Devices and checking health on them
type ResourceManager interface {
	Resource() spec.ResourceName
	Devices() Devices
	GetDevicePaths([]string) []string
	GetPreferredAllocation(available, required []string, size int) ([]string, error)
	CheckHealth(stop <-chan interface{}, unhealthy chan<- *Device, disableNVML <-chan bool, ackDisableHealthChecks chan<- bool) error
	ValidateRequest(AnnotatedIDs) error
}

// Resource gets the resource name associated with the ResourceManager
func (r *resourceManager) Resource() spec.ResourceName {
	_ = "STUB: not implemented"

	// Resource gets the devices managed by the ResourceManager
	return *new(spec.ResourceName)
}

func (r *resourceManager) Devices() Devices { _ = "STUB: not implemented"; return *new(Devices) }

var errInvalidRequest = errors.New("invalid request")

// ValidateRequest checks the requested IDs against the resource manager configuration.
// It asserts that all requested IDs are known to the resource manager and that the request is
// valid for a specified sharing configuration.
func (r *resourceManager) ValidateRequest(ids AnnotatedIDs) error {
	_ = "STUB: not implemented"
	// Assert that all requested IDs are known to the resource manager
	return nil
}

// If the devices being allocated are replicas, then (conditionally)
// error out if more than one resource is being allocated.

// For MPS sharing, we explicitly ignore the FailRequestsGreaterThanOne
// value in the sharing settings.
// This setting was added to timeslicing after the initial release and
// is set to `false` to maintain backward compatibility with existing
// deployments. If we do extend MPS to allow multiple devices to be
// requested, the MPS API will be extended separately from the
// time-slicing API.

// AddDefaultResourcesToConfig adds default resource matching rules to config.Resources
func AddDefaultResourcesToConfig(infolib info.Interface, nvmllib nvml.Interface, devicelib device.Interface, config *nvidia.DeviceConfig) error {
	_ = "STUB: not implemented"
	return nil
}
