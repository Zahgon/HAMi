/**
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
**/

package nvidia

import (
	"github.com/NVIDIA/go-nvlib/pkg/nvlib/device"
	"github.com/NVIDIA/go-nvml/pkg/nvml"
)

// Device represents a GPU device as reported by NVML, including all of its
// Point-to-Point link information.
type Device struct {
	nvlibDevice
	Index int
	Links map[int][]P2PLink
}

// DeviceList stores an ordered list of devices.
type DeviceList []*Device

// Filter filters out the selected devices from the list.
// Note that the specified uuids must exist in the list of devices.
func (d DeviceList) Filter(uuids []string) (DeviceList, error) {
	_ = "STUB: not implemented"
	return *new(DeviceList), nil
}

// P2PLink represents a Point-to-Point link between two GPU devices. The link
// is between the Device struct this struct is embedded in and the GPU Device
// contained in the P2PLink struct itself.
type P2PLink struct {
	GPU  *Device
	Type P2PLinkType
}

type nvlibDevice struct {
	device.Device
	// The previous binding implementation used to cache specific device properties.
	// These should be considered deprecated and the functions associated with device.Device
	// should be used instead.
	UUID string
	PCI  struct {
		BusID string
	}
	CPUAffinity *uint
}

// deviceListBuilder stores the options required to build a list of linked devices.
type deviceListBuilder struct {
	nvmllib   nvml.Interface
	devicelib device.Interface
}

// NewDevices creates a list of Devices from all available nvml.Devices using the specified options.
func NewDevices() (DeviceList, error) { _ = "STUB: not implemented"; return *new(DeviceList), nil }

// build uses the configured options to build a DeviceList.
func (o *deviceListBuilder) build() (DeviceList, error) {
	_ = "STUB: not implemented"
	return *new(DeviceList), nil
}

// newDevice constructs a Device for the specified index and nvml Device.
func newDevice(i int, d device.Device) (*Device, error) { _ = "STUB: not implemented"; return nil, nil }

type ListDeviceScore []DeviceScore

type DeviceScore struct {
	UUID string `json:"uuid"`
	// Score is record and other gpu communications score,
	// the value bigger communications bandwidth the higher.
	Score map[string]int `json:"score"`
}

func CalculateGPUScore(available []string) (ListDeviceScore, error) {
	_ = "STUB: not implemented"
	return *new(ListDeviceScore), nil
}

func calculateGPUScore(devices []*Device) ListDeviceScore {
	_ = "STUB: not implemented"
	return *new(ListDeviceScore)
}

func calculateGPUPairScore(gpu0 *Device, gpu1 *Device) int { _ = "STUB: not implemented"; return 0 }
