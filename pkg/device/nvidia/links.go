/*
*
# Copyright 2024 NVIDIA CORPORATION
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
*
*/

package nvidia

import (
	"github.com/NVIDIA/go-nvlib/pkg/nvlib/device"
	"github.com/NVIDIA/go-nvml/pkg/nvml"
)

// P2PLinkType defines the link information between two devices.
type P2PLinkType uint

// The following constants define the nature of a link between two devices.
// These include peer-2-peer and NVLink information.
const (
	P2PLinkUnknown P2PLinkType = iota
	P2PLinkCrossCPU
	P2PLinkSameCPU
	P2PLinkHostBridge
	P2PLinkMultiSwitch
	P2PLinkSingleSwitch
	P2PLinkSameBoard
	SingleNVLINKLink
	TwoNVLINKLinks
	ThreeNVLINKLinks
	FourNVLINKLinks
	FiveNVLINKLinks
	SixNVLINKLinks
	SevenNVLINKLinks
	EightNVLINKLinks
	NineNVLINKLinks
	TenNVLINKLinks
	ElevenNVLINKLinks
	TwelveNVLINKLinks
	ThirteenNVLINKLinks
	FourteenNVLINKLinks
	FifteenNVLINKLinks
	SixteenNVLINKLinks
	SeventeenNVLINKLinks
	EighteenNVLINKLinks
)

// String returns the string representation of the P2PLink type.
func (l P2PLinkType) String() string { _ = "STUB: not implemented"; return "" }

// GetP2PLink gets the peer-to-peer connectivity between two devices.
func GetP2PLink(dev1 device.Device, dev2 device.Device) (P2PLinkType, error) {
	_ = "STUB: not implemented"
	return *new(P2PLinkType), nil
}

// NVML_TOPOLOGY_CPU was renamed NVML_TOPOLOGY_NODE

// GetNVLink gets the number of NVLinks between the specified devices.
func GetNVLink(dev1 device.Device, dev2 device.Device) (P2PLinkType, error) {
	_ = "STUB: not implemented"
	return *new(P2PLinkType), nil
}

// TODO(klueska): Handle NVSwitch semantics

// getAllNvLinkRemotePciInfo returns the PCI info for all devices attached to the specified device by an NVLink.
func getAllNvLinkRemotePciInfo(dev device.Device) ([]PciInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PciInfo is a type alias to nvml.PciInfo to allow for functions to be defined on the type.
type PciInfo nvml.PciInfo

// BusID provides a utility function that returns the string representation of the bus ID.
// Note that the []int8 slice member is named BusId.
func (p PciInfo) BusID() string { _ = "STUB: not implemented"; return "" }

// CPUAffinity returns the CPU affinity associated with a specified PCI device.
// If NUMA information is not available, this returns nil.
func (p PciInfo) CPUAffinity() *uint { _ = "STUB: not implemented"; return nil }

// NumaNode returns the numa node associates with a PCI device.
// If numa is unsupported, -1 is returned.
func (p PciInfo) NumaNode() int64 {
	_ = "STUB: not implemented"
	// Read the numa_node file associated with the PCI Device Info
	return 0
}
