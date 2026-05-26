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

package metax

import (
	"github.com/Project-HAMi/HAMi/pkg/device"
)

const (
	MetaxSDeviceAnno       = "metax-tech.com/node-gpu-devices"
	MetaxAllocatedSDevices = "metax-tech.com/gpu-devices-allocated"
	MetaxPredicateTime     = "metax-tech.com/predicate-time"

	MetaxUseUUID   = "metax-tech.com/use-gpuuuid"
	MetaxNoUseUUID = "metax-tech.com/nouse-gpuuuid"

	MetaxSGPUQosPolicy     = "metax-tech.com/sgpu-qos-policy"
	MetaxSGPUTopologyAware = "metax-tech.com/sgpu-topology-aware"
	MetaxSGPUAppClass      = "metax-tech.com/sgpu-app-class"
)

const (
	BestEffort = "best-effort"
	FixedShare = "fixed-share"
	BurstShare = "burst-share"
)

const (
	Online  = "online"
	Offline = "offline"
)

type MetaxSDeviceInfo struct {
	UUID              string `json:"uuid"`
	BDF               string `json:"bdf,omitempty"`
	Model             string `json:"model,omitempty"`
	TotalDevCount     int32  `json:"totalDevCount,omitempty"`
	TotalCompute      int32  `json:"totalCompute,omitempty"`
	TotalVRam         int32  `json:"totalVRam,omitempty"`
	AvailableDevCount int32  `json:"availableDevCount,omitempty"`
	AvailableCompute  int32  `json:"availableCompute,omitempty"`
	AvailableVRam     int32  `json:"availableVRam,omitempty"`
	Numa              int32  `json:"numa,omitempty"`
	Healthy           bool   `json:"healthy,omitempty"`
	QosPolicy         string `json:"qosPolicy,omitempty"`
	LinkZone          int32  `json:"linkZone,omitempty"`
}
type NodeMetaxSDeviceInfo []*MetaxSDeviceInfo

type ContainerMetaxSDevice struct {
	UUID    string `json:"uuid"`
	Compute int32  `json:"compute,omitempty"`
	VRam    int32  `json:"vRam,omitempty"`
}
type ContainerMetaxSDevices []ContainerMetaxSDevice
type PodMetaxSDevice []ContainerMetaxSDevices

func (ni NodeMetaxSDeviceInfo) String() string { _ = "STUB: not implemented"; return "" }

func (sdev *PodMetaxSDevice) String() string { _ = "STUB: not implemented"; return "" }

func convertMetaxSDeviceToHAMIDevice(metaxSDevices []*MetaxSDeviceInfo) []*device.DeviceInfo {
	_ = "STUB: not implemented"
	return nil
}

func convertHAMIPodDeviceToMetaxPodDevice(hamiPodDevices device.PodSingleDevice) PodMetaxSDevice {
	_ = "STUB: not implemented"
	return *new(PodMetaxSDevice)
}
