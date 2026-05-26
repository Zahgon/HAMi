/*
Copyright 2024 The HAMi Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package policy

import (
	"github.com/Project-HAMi/HAMi/pkg/device"
)

type DeviceListsScore struct {
	Device *device.DeviceUsage
	// Score recode every device user/allocate score
	Score float32
}

type DeviceUsageList struct {
	DeviceLists []*DeviceListsScore
	Policy      string
}

func (l DeviceUsageList) Len() int { _ = "STUB: not implemented"; return 0 }

func (l DeviceUsageList) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (l DeviceUsageList) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

// default policy is spread

func (l DeviceUsageList) DeepCopy() DeviceUsageList {
	_ = "STUB: not implemented"
	return *new(DeviceUsageList)
}

func (ds *DeviceListsScore) DeepCopy() *DeviceListsScore { _ = "STUB: not implemented"; return nil }

func (ds *DeviceListsScore) ComputeScore(requests device.ContainerDeviceRequests) {
	_ = "STUB: not implemented"
	return
}

// Here we are required to use the same type device
