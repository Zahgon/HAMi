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

package mthreads

import (
	"flag"

	"github.com/Project-HAMi/HAMi/pkg/device"

	corev1 "k8s.io/api/core/v1"
)

type MthreadsDevices struct {
}

const (
	MthreadsGPUDevice       = "Mthreads"
	MthreadsGPUCommonWord   = "Mthreads"
	MthreadsDeviceSelection = "mthreads.com/gpu-index"
	// MthreadsUseUUID annotation specifies a comma-separated list of Mthreads UUIDs to use.
	MthreadsUseUUID = "mthreads.ai/use-gpuuuid"
	// MthreadsNoUseUUID annotation specifies a comma-separated list of Mthreads UUIDs to exclude.
	MthreadsNoUseUUID        = "mthreads.ai/nouse-gpuuuid"
	MthreadsAssignedGPUIndex = "mthreads.com/gpu-index"
	MthreadsAssignedNode     = "mthreads.com/predicate-node"
	MthreadsPredicateTime    = "mthreads.com/predicate-time"
	coresPerMthreadsGPU      = 16
	memoryPerMthreadsGPU     = 96
)

var (
	MthreadsResourceCount  string
	MthreadsResourceMemory string
	MthreadsResourceCores  string
	legalMemoryslices      = []int64{2, 4, 8, 16, 32, 64, 96}
)

type MthreadsConfig struct {
	ResourceCountName  string `yaml:"resourceCountName"`
	ResourceMemoryName string `yaml:"resourceMemoryName"`
	ResourceCoreName   string `yaml:"resourceCoreName"`
}

func InitMthreadsDevice(config MthreadsConfig) *MthreadsDevices {
	_ = "STUB: not implemented"
	return nil
}

func (dev *MthreadsDevices) CommonWord() string { _ = "STUB: not implemented"; return "" }

func ParseConfig(fs *flag.FlagSet) { _ = "STUB: not implemented"; return }

func (dev *MthreadsDevices) MutateAdmission(ctr *corev1.Container, p *corev1.Pod) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (dev *MthreadsDevices) GetNodeDevices(n corev1.Node) ([]*device.DeviceInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (dev *MthreadsDevices) PatchAnnotations(pod *corev1.Pod, annoinput *map[string]string, pd device.PodDevices) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

//(*annoinput)[MthreadsAssignedNode]=

func (dev *MthreadsDevices) LockNode(n *corev1.Node, p *corev1.Pod) error {
	_ = "STUB: not implemented"
	return nil
}

func (dev *MthreadsDevices) ReleaseNodeLock(n *corev1.Node, p *corev1.Pod) error {
	_ = "STUB: not implemented"
	return nil
}

func (dev *MthreadsDevices) NodeCleanUp(nn string) error { _ = "STUB: not implemented"; return nil }

func (dev *MthreadsDevices) checkType(annos map[string]string, d device.DeviceUsage, n device.ContainerDeviceRequest) (bool, bool, bool) {
	_ = "STUB: not implemented"
	return false, false, false
}

func (dev *MthreadsDevices) CheckHealth(devType string, n *corev1.Node) (bool, bool) {
	_ = "STUB: not implemented"
	return false, false
}

func (dev *MthreadsDevices) GenerateResourceRequests(ctr *corev1.Container) device.ContainerDeviceRequest {
	_ = "STUB: not implemented"
	return *new(device.ContainerDeviceRequest)
}

func (dev *MthreadsDevices) customFilterRule(allocated *device.PodDevices, request device.ContainerDeviceRequest, toAllocate device.ContainerDevices, device *device.DeviceUsage) bool {
	_ = "STUB: not implemented"
	return false
}

func (dev *MthreadsDevices) ScoreNode(node *corev1.Node, podDevices device.PodSingleDevice, previous []*device.DeviceUsage, policy string) float32 {
	_ = "STUB: not implemented"
	return 0
}

func (dev *MthreadsDevices) AddResourceUsage(pod *corev1.Pod, n *device.DeviceUsage, ctr *device.ContainerDevice) error {
	_ = "STUB: not implemented"
	return nil
}

func (mth *MthreadsDevices) Fit(devices []*device.DeviceUsage, request device.ContainerDeviceRequest, pod *corev1.Pod, nodeInfo *device.NodeInfo, allocated *device.PodDevices) (bool, map[string]device.ContainerDevices, string) {
	_ = "STUB: not implemented"
	return false, nil, ""
}

//return false, tmpDevs

//This incurs an issue

// Coresreq=100 indicates it want this card exclusively

// You can't allocate core=0 job to an already full GPU

func (dev *MthreadsDevices) GetResourceNames() device.ResourceNames {
	_ = "STUB: not implemented"
	return *new(device.ResourceNames)
}
