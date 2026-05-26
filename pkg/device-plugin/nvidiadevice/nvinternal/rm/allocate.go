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

// distributedAlloc returns a list of devices such that any replicated
// devices are distributed across all replicated GPUs equally. It takes into
// account already allocated replicas to ensure a proper balance across them.
func (r *resourceManager) distributedAlloc(available, required []string, size int) ([]string, error) {
	_ = "STUB: not implemented"
	// Get the set of candidate devices as the difference between available and required.
	return nil, nil
}

// For each candidate device, build a mapping of (stripped) device ID to
// total / available replicas for that device.

// Grab the set of 'needed' devices one-by-one from the candidates list.
// Before selecting each candidate, first sort the candidate list using the
// replicas map above. After sorting, the first element in the list will
// contain the device with the least difference between total and available
// replications (based on what's already been allocated). Add this device
// to the list of devices to allocate, remove it from the candidate list,
// down its available count in the replicas map, and repeat.

// Add the set of required devices to this list and return it.
