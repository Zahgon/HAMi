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

package cdi

import (
	spec "github.com/NVIDIA/k8s-device-plugin/api/config/v1"
	"github.com/Project-HAMi/HAMi/pkg/device-plugin/nvidiadevice/nvinternal/imex"
)

// Option defines a function for passing options to the New() call
type Option func(*cdiHandler)

// WithDeviceListStrategies provides an Option to set the enabled flag used by the 'cdi' interface
func WithDeviceListStrategies(deviceListStrategies spec.DeviceListStrategies) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithDriverRoot provides an Option to set the driver root used by the 'cdi' interface.
func WithDriverRoot(root string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithDevRoot sets the dev root for the `cdi` interface.
func WithDevRoot(root string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithTargetDriverRoot provides an Option to set the target (host) driver root used by the 'cdi' interface
func WithTargetDriverRoot(root string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithTargetDevRoot provides an Option to set the target (host) dev root used by the 'cdi' interface
func WithTargetDevRoot(root string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithNvidiaCTKPath provides an Option to set the nvidia-ctk path used by the 'cdi' interface
func WithNvidiaCTKPath(path string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithDeviceIDStrategy provides an Option to set the device ID strategy used by the 'cdi' interface
func WithDeviceIDStrategy(strategy string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithVendor provides an Option to set the vendor used by the 'cdi' interface
func WithVendor(vendor string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithGdrcopyEnabled provides an option to set whether a GDS CDI spec should be generated
func WithGdrcopyEnabled(enabled bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithGdsEnabled provides an option to set whether a GDS CDI spec should be generated
func WithGdsEnabled(enabled bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithMofedEnabled provides an option to set whether a MOFED CDI spec should be generated
func WithMofedEnabled(enabled bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithImexChannels sets the IMEX channels for which CDI specs should be generated.
func WithImexChannels(imexChannels imex.Channels) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}
