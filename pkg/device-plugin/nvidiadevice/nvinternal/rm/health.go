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

const (
	// envDisableHealthChecks defines the environment variable that is checked to determine whether healthchecks
	// should be disabled. If this envvar is set to "all" or contains the string "xids", healthchecks are
	// disabled entirely. If set, the envvar is treated as a comma-separated list of Xids to ignore. Note that
	// this is in addition to the Application errors that are already ignored.
	envDisableHealthChecks = "DP_DISABLE_HEALTHCHECKS"
	// envEnableHealthChecks defines the environment variable that is checked to
	// determine which XIDs should be explicitly enabled. XIDs specified here
	// override the ones specified in the `DP_DISABLE_HEALTHCHECKS`.
	// Note that this also allows individual XIDs to be selected when ALL XIDs
	// are disabled.
	envEnableHealthChecks = "DP_ENABLE_HEALTHCHECKS"
)

// CheckHealth performs health checks on a set of devices, writing to the 'unhealthy' channel with any unhealthy devices
func (r *nvmlResourceManager) checkHealth(stop <-chan interface{}, devices Devices, unhealthy chan<- *Device, disableNVML <-chan bool) error {
	_ = "STUB: not implemented"
	return nil
}

// If we cannot reliably determine the device UUID, we mark all devices as unhealthy.

const allXIDs = 0

// disabledXIDs stores a map of explicitly disabled XIDs.
// The special XID `allXIDs` indicates that all XIDs are disabled, but does
// allow for specific XIDs to be enabled even if this is the case.
type disabledXIDs map[uint64]bool

// Disabled returns whether XID-based health checks are disabled.
// These are considered if all XIDs have been disabled AND no other XIDs have
// been explicitly enabled.
func (h disabledXIDs) IsAllDisabled() bool { _ = "STUB: not implemented"; return false }

// At this point we wither have explicitly disabled XIDs or explicitly
// enabled XIDs. Since ANY XID that's not specified is assumed enabled, we
// return here.

// IsDisabled checks whether the specified XID has been explicitly disalbled.
// An XID is considered disabled if it has been explicitly disabled, or all XIDs
// have been disabled.
func (h disabledXIDs) IsDisabled(xid uint64) bool {
	_ = "STUB: not implemented"
	// Handle the case where enabled=all.
	return false
}

// Handle the case where the XID has been specifically enabled (or disabled)

// getHealthCheckXids returns the XIDs that are considered fatal.
// Here we combine the following (in order of precedence):
// * A list of explicitly disabled XIDs (including all XIDs)
// * A list of hardcoded disabled XIDs
// * A list of explicitly enabled XIDs (including all XIDs)
//
// Note that if an XID is explicitly enabled, this takes precedence over it
// having been disabled either explicitly or implicitly.
func getHealthCheckXids() disabledXIDs { _ = "STUB: not implemented"; return *new(disabledXIDs) }

// TODO: We should not read the envvar here directly, but instead
// "upgrade" this to a top-level config option.

// TODO: We should not read the envvar here directly, but instead
// "upgrade" this to a top-level config option.

// Add the list of hardcoded disabled (ignored) XIDs:
// FIXME: formalize the full list and document it.
// http://docs.nvidia.com/deploy/xid-errors/index.html#topic_4
// Application errors: the GPU should still be healthy

// Graphics Engine Exception
// GPU memory page fault
// GPU stopped processing
// Preemptive cleanup, due to previous errors
// Video processor exception
// Context Switch Timeout Error

// Explicitly ENABLE specific XIDs,

// newHealthCheckXIDs converts a list of Xids to a healthCheckXIDs map.
// Special xid values 'all' and 'xids' return a special map that matches all
// xids.
// For other xids, these are converted to a uint64 values with invalid values
// being ignored.
func newHealthCheckXIDs(xids ...string) disabledXIDs {
	_ = "STUB: not implemented"
	return *new(disabledXIDs)
}

// TODO: We should have a different type for "all" and "all-except"

// getDevicePlacement returns the placement of the specified device.
// For a MIG device the placement is defined by the 3-tuple <parent UUID, GI, CI>
// For a full device the returned 3-tuple is the device's uuid and 0xFFFFFFFF for the other two elements.
func (r *nvmlResourceManager) getDevicePlacement(d *Device) (string, uint32, uint32, error) {
	_ = "STUB: not implemented"
	return "", 0, 0, nil
}

// getMigDeviceParts returns the parent GI and CI ids of the MIG device.
func (r *nvmlResourceManager) getMigDeviceParts(d *Device) (string, uint32, uint32, error) {
	_ = "STUB: not implemented"
	return "", 0, 0, nil
}

// For older driver versions, the call to DeviceGetHandleByUUID will fail for MIG devices.

//nolint:gosec  // We know that the values returned from Get*InstanceId are within the valid uint32 range.

// parseMigDeviceUUID splits the MIG device UUID into the parent device UUID and ci and gi
func parseMigDeviceUUID(mig string) (string, uint32, uint32, error) {
	_ = "STUB: not implemented"
	return "", 0, 0, nil
}

func toUint32(s string) (uint32, error) { _ = "STUB: not implemented"; return 0, nil }

//nolint:gosec  // Since we parse s with a 32-bit size this will not overflow.
