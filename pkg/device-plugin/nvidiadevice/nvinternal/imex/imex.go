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

package imex

import (
	spec "github.com/NVIDIA/k8s-device-plugin/api/config/v1"
)

// Channels represents a set of IMEX channels.
type Channels []*Channel

// Channel represents an IMEX channel.
type Channel struct {
	ID       string
	Path     string
	HostPath string
}

// GetChannels returns the set of channels for the given config.
// If the selection of the default IMEX channel is disabled no channels are returned.
func GetChannels(config *spec.Config, devRoot string) (Channels, error) {
	_ = "STUB: not implemented"
	return *new(Channels), nil
}

// exists checks whether the IMEX channel exists.
// We check both the Path and HostPath since the location of the device node
// associated with the channel in the container is dependent on how it is
// injected.
// For example, if the host driver root is mounted at /driver-root the channel
// device node would be available at /driver-root/dev even if it was not
// injected into the container through any other mechanism.
// For the case of management containers using CDI to inject device nodes, these
// device nodes would exist at /dev in the container instead.
func (c Channel) exists() (bool, error) { _ = "STUB: not implemented"; return false, nil }
