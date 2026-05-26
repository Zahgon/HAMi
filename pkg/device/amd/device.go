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

package amd

import (
	"flag"

	"github.com/Project-HAMi/HAMi/pkg/device"

	corev1 "k8s.io/api/core/v1"
)

type AMDDevices struct {
	resourceCountName  string
	resourceMemoryName string
}

const (
	AMDDevice          = "AMDGPU"
	AMDCommonWord      = "AMDGPU"
	AMDDeviceSelection = "amd.com/gpu-index"
	AMDUseUUID         = "amd.com/use-gpu-uuid"
	AMDNoUseUUID       = "amd.com/nouse-gpu-uuid"
	AMDAssignedNode    = "amd.com/predicate-node"
	Mi300xMemory       = 192000
)

type AMDConfig struct {
	ResourceCountName  string `yaml:"resourceCountName"`
	ResourceMemoryName string `yaml:"resourceMemoryName"`
}

func InitAMDGPUDevice(config AMDConfig) *AMDDevices { _ = "STUB: not implemented"; return nil }

func (dev *AMDDevices) CommonWord() string { _ = "STUB: not implemented"; return "" }

func ParseConfig(fs *flag.FlagSet) { _ = "STUB: not implemented"; return }

func (dev *AMDDevices) MutateAdmission(ctr *corev1.Container, p *corev1.Pod) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (dev *AMDDevices) GetNodeDevices(n corev1.Node) ([]*device.DeviceInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (dev *AMDDevices) PatchAnnotations(pod *corev1.Pod, annoinput *map[string]string, pd device.PodDevices) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func (dev *AMDDevices) LockNode(n *corev1.Node, p *corev1.Pod) error {
	_ = "STUB: not implemented"
	return nil
}

func (dev *AMDDevices) ReleaseNodeLock(n *corev1.Node, p *corev1.Pod) error {
	_ = "STUB: not implemented"
	return nil
}

func (dev *AMDDevices) NodeCleanUp(nn string) error { _ = "STUB: not implemented"; return nil }

func (dev *AMDDevices) checkType(n device.ContainerDeviceRequest) (bool, bool, bool) {
	_ = "STUB: not implemented"
	return false, false, false
}

func (dev *AMDDevices) CheckHealth(devType string, n *corev1.Node) (bool, bool) {
	_ = "STUB: not implemented"
	return false, false
}

func (dev *AMDDevices) GetResourceNames() device.ResourceNames {
	_ = "STUB: not implemented"
	return *new(device.ResourceNames)
}

func (dev *AMDDevices) GenerateResourceRequests(ctr *corev1.Container) device.ContainerDeviceRequest {
	_ = "STUB: not implemented"
	return *new(device.ContainerDeviceRequest)
}

//amdResourceMemory := corev1.ResourceName(dev.resourceMemoryName)

func (dev *AMDDevices) ScoreNode(node *corev1.Node, podDevices device.PodSingleDevice, previous []*device.DeviceUsage, policy string) float32 {
	_ = "STUB: not implemented"
	return 0
}

func (dev *AMDDevices) AddResourceUsage(pod *corev1.Pod, n *device.DeviceUsage, ctr *device.ContainerDevice) error {
	_ = "STUB: not implemented"
	return nil
}

func (amddevice *AMDDevices) Fit(devices []*device.DeviceUsage, request device.ContainerDeviceRequest, pod *corev1.Pod, nodeinfo *device.NodeInfo, allocated *device.PodDevices) (bool, map[string]device.ContainerDevices, string) {
	_ = "STUB: not implemented"
	return false, nil, ""
}
