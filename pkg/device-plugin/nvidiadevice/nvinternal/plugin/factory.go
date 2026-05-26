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

	"github.com/NVIDIA/go-nvlib/pkg/nvlib/device"
	"github.com/NVIDIA/go-nvlib/pkg/nvlib/info"
	"github.com/NVIDIA/go-nvml/pkg/nvml"

	spec "github.com/NVIDIA/k8s-device-plugin/api/config/v1"
	"github.com/Project-HAMi/HAMi/pkg/device-plugin/nvidiadevice/nvinternal/cdi"
	"github.com/Project-HAMi/HAMi/pkg/device-plugin/nvidiadevice/nvinternal/imex"
	"github.com/Project-HAMi/HAMi/pkg/device-plugin/nvidiadevice/nvinternal/rm"
	"github.com/Project-HAMi/HAMi/pkg/device/nvidia"
)

type options struct {
	infolib   info.Interface
	nvmllib   nvml.Interface
	devicelib device.Interface

	failOnInitError bool

	cdiHandler cdi.Interface
	config     *nvidia.DeviceConfig

	deviceListStrategies spec.DeviceListStrategies

	imexChannels imex.Channels
}

// New a new set of plugins with the supplied options.
func New(ctx context.Context, infolib info.Interface, nvmllib nvml.Interface, devicelib device.Interface, opts ...Option) ([]Interface, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getResourceManager constructs a set of resource managers.
// Each resource manager maps to a specific named extended resource and may
// include full GPUs or MIG devices.
func (o *options) getResourceManagers() ([]rm.ResourceManager, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (o *options) resolveStrategy(strategy string) string { _ = "STUB: not implemented"; return "" }
