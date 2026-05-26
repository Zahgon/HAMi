/*
Copyright 2025 The HAMi Authors.

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

package kunlun

import (
	"github.com/Project-HAMi/HAMi/pkg/device"
)

type FitFn func(device *device.DeviceUsage, request device.ContainerDeviceRequest) bool

func parseUsage(devices []*device.DeviceUsage, request device.ContainerDeviceRequest, fitFn FitFn) []int {
	_ = "STUB: not implemented"
	return nil
}

func addidx(temp []int, value int) []int { _ = "STUB: not implemented"; return nil }

func getvalue(t int) int { _ = "STUB: not implemented"; return 0 }

func countbubble(t []int) int { _ = "STUB: not implemented"; return 0 }

func calcscore(p []int, c []int) float32 { _ = "STUB: not implemented"; return 0 }

func parseInterconnection() [][]int { _ = "STUB: not implemented"; return nil }

func parseInterconnection2() [][]int { _ = "STUB: not implemented"; return nil }

func interconnect(devices []*device.DeviceUsage, request device.ContainerDeviceRequest, fitFn FitFn) []int {
	_ = "STUB: not implemented"
	return nil
}

func canMeet(have, want []int) bool { _ = "STUB: not implemented"; return false }

func delta(have, want []int) []int { _ = "STUB: not implemented"; return nil }

func devicepick(devices []*device.DeviceUsage, start int, request device.ContainerDeviceRequest, fitFn FitFn) []int {
	_ = "STUB: not implemented"
	return nil
}

func graghSelect(devices []*device.DeviceUsage, request device.ContainerDeviceRequest, fitFn FitFn) []int {
	_ = "STUB: not implemented"
	return nil
}
