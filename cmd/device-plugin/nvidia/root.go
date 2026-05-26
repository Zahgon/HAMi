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

package main

type root string

func (r root) join(parts ...string) string { _ = "STUB: not implemented"; return "" }

// getDevRoot returns the dev root associated with the root.
// If the root is not a dev root, this defaults to "/".
func (r root) getDevRoot() string { _ = "STUB: not implemented"; return "" }

// isDevRoot checks whether the specified root is a dev root.
// A dev root is defined as a root containing a /dev folder.
func (r root) isDevRoot() bool { _ = "STUB: not implemented"; return false }

func (r root) tryResolveLibrary(libraryName string) string { _ = "STUB: not implemented"; return "" }

// resolveLink finds the target of a symlink or the file itself in the
// case of a regular file.
// This is equivalent to running `readlink -f ${l}`.
func resolveLink(l string) (string, error) { _ = "STUB: not implemented"; return "", nil }
