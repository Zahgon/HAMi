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

package v0

const maxDevices = 16

type deviceMemory struct {
	contextSize uint64
	moduleSize  uint64
	bufferSize  uint64
	offset      uint64
	total       uint64
}

type deviceUtilization struct {
	decUtil uint64
	encUtil uint64
	smUtil  uint64
}

type shrregProcSlotT struct {
	pid         int32
	hostpid     int32
	used        [16]deviceMemory
	monitorused [16]uint64
	deviceUtil  [16]deviceUtilization
	status      int32
}

type uuid struct {
	uuid [96]byte
}

type semT struct {
	sem [32]byte
}

type sharedRegionT struct {
	initializedFlag int32
	smInitFlag      int32
	ownerPid        uint32
	sem             semT
	num             uint64
	uuids           [16]uuid

	limit   [16]uint64
	smLimit [16]uint64
	procs   [1024]shrregProcSlotT

	procnum           int32
	utilizationSwitch int32
	recentKernel      int32
	priority          int32
}

type Spec struct {
	sr *sharedRegionT
}

func (s Spec) DeviceMax() int { _ = "STUB: not implemented"; return 0 }

func (s Spec) DeviceNum() int { _ = "STUB: not implemented"; return 0 }

func (s Spec) DeviceMemoryContextSize(idx int) uint64 { _ = "STUB: not implemented"; return 0 }

func (s Spec) DeviceMemoryModuleSize(idx int) uint64 { _ = "STUB: not implemented"; return 0 }

func (s Spec) DeviceMemoryBufferSize(idx int) uint64 { _ = "STUB: not implemented"; return 0 }

func (s Spec) DeviceMemoryOffset(idx int) uint64 { _ = "STUB: not implemented"; return 0 }

func (s Spec) DeviceMemoryTotal(idx int) uint64 { _ = "STUB: not implemented"; return 0 }

func (s Spec) DeviceSmUtil(idx int) uint64 { _ = "STUB: not implemented"; return 0 }

func (s Spec) SetDeviceSmLimit(l uint64) { _ = "STUB: not implemented"; return }

func (s Spec) IsValidUUID(idx int) bool { _ = "STUB: not implemented"; return false }

func (s Spec) DeviceUUID(idx int) string { _ = "STUB: not implemented"; return "" }

func (s Spec) DeviceMemoryLimit(idx int) uint64 { _ = "STUB: not implemented"; return 0 }

func (s Spec) SetDeviceMemoryLimit(l uint64) { _ = "STUB: not implemented"; return }

func (s Spec) LastKernelTime() int64 { _ = "STUB: not implemented"; return 0 }

func CastSpec(data []byte) Spec { _ = "STUB: not implemented"; return *new(Spec) }

//	func (s *SharedRegionT) UsedMemory(idx int) (uint64, error) {
//		return 0, nil
//	}

func (s Spec) GetPriority() int { _ = "STUB: not implemented"; return 0 }

func (s Spec) GetRecentKernel() int32 { _ = "STUB: not implemented"; return 0 }

func (s Spec) SetRecentKernel(v int32) { _ = "STUB: not implemented"; return }

func (s Spec) GetUtilizationSwitch() int32 { _ = "STUB: not implemented"; return 0 }

func (s Spec) SetUtilizationSwitch(v int32) { _ = "STUB: not implemented"; return }
