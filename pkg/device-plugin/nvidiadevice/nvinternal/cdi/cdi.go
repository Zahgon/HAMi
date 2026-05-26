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
	"github.com/NVIDIA/go-nvlib/pkg/nvlib/device"
	"github.com/NVIDIA/go-nvlib/pkg/nvlib/info"
	"github.com/NVIDIA/go-nvml/pkg/nvml"
	"github.com/NVIDIA/nvidia-container-toolkit/pkg/nvcdi"
	"github.com/NVIDIA/nvidia-container-toolkit/pkg/nvcdi/transform"
	"github.com/sirupsen/logrus"

	spec "github.com/NVIDIA/k8s-device-plugin/api/config/v1"
	"github.com/Project-HAMi/HAMi/pkg/device-plugin/nvidiadevice/nvinternal/imex"
)

const (
	cdiRoot = "/var/run/cdi"
)

// cdiHandler creates CDI specs for devices assocatied with the device plugin
type cdiHandler struct {
	infolib   info.Interface
	nvmllib   nvml.Interface
	devicelib device.Interface

	logger           *logrus.Logger
	driverRoot       string
	devRoot          string
	targetDriverRoot string
	targetDevRoot    string
	nvidiaCTKPath    string
	vendor           string
	deviceIDStrategy string

	deviceListStrategies spec.DeviceListStrategies

	gdsEnabled     bool
	mofedEnabled   bool
	gdrcopyEnabled bool

	imexChannels imex.Channels

	cdilibs         map[string]nvcdi.SpecGenerator
	additionalModes []string
}

var _ Interface = &cdiHandler{}

// New constructs a new instance of the 'cdi' interface
func New(infolib info.Interface, nvmllib nvml.Interface, devicelib device.Interface, opts ...Option) (Interface, error) {
	_ = "STUB: not implemented"
	return *new(Interface), nil
}

// CreateSpecFile creates a CDI spec file for the specified devices.
func (cdi *cdiHandler) CreateSpecFile() error { _ = "STUB: not implemented"; return nil }

// TODO: Once the NewDriverTransformer is merged in container-toolkit we can instantiate it directly.

// TODO: This is a brittle check since it relies on exact string matches.
// We should pull this functionality into the CDI tooling instead.

// Remove the classes with empty specs from the supported types.

func (cdi *cdiHandler) getRootTransformer() transform.Transformer {
	_ = "STUB: not implemented"
	return *new(transform.Transformer)
}

// QualifiedName constructs a CDI qualified device name for the specified resources.
// Note: This assumes that the specified id matches the device name returned by the naming strategy.
func (cdi *cdiHandler) QualifiedName(class string, id string) string {
	_ = "STUB: not implemented"
	return ""
}

// AdditionalDevices returns the optional CDI devices based on the device plugin
// configuration.
// Here we check for requested modes as well as whether the modes have a valid
// CDI spec associated with them.
func (cdi *cdiHandler) AdditionalDevices() []string { _ = "STUB: not implemented"; return nil }
