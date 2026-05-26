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

package metax

import (
	"github.com/Project-HAMi/HAMi/pkg/device"

	corev1 "k8s.io/api/core/v1"
)

type MetaxDevices struct {
}

const (
	MetaxGPUDevice       = "Metax-GPU"
	MetaxGPUCommonWord   = "Metax-GPU"
	MetaxAnnotationLoss  = "metax-tech.com/gpu.topology.losses"
	MetaxAnnotationScore = "metax-tech.com/gpu.topology.scores"
)

var (
	MetaxResourceCount string
)

func InitMetaxDevice(config MetaxConfig) *MetaxDevices { _ = "STUB: not implemented"; return nil }

func (dev *MetaxDevices) CommonWord() string { _ = "STUB: not implemented"; return "" }

func (dev *MetaxDevices) MutateAdmission(ctr *corev1.Container, p *corev1.Pod) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (dev *MetaxDevices) GetNodeDevices(n corev1.Node) ([]*device.DeviceInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (dev *MetaxDevices) PatchAnnotations(pod *corev1.Pod, annoinput *map[string]string, pd device.PodDevices) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func (dev *MetaxDevices) LockNode(n *corev1.Node, p *corev1.Pod) error {
	_ = "STUB: not implemented"
	return nil
}

func (dev *MetaxDevices) ReleaseNodeLock(n *corev1.Node, p *corev1.Pod) error {
	_ = "STUB: not implemented"
	return nil
}

func (dev *MetaxDevices) NodeCleanUp(nn string) error { _ = "STUB: not implemented"; return nil }

func (dev *MetaxDevices) checkType(annos map[string]string, d device.DeviceUsage, n device.ContainerDeviceRequest) (bool, bool, bool) {
	_ = "STUB: not implemented"
	return false, false, false
}

func (dev *MetaxDevices) checkUUID(annos map[string]string, d device.DeviceUsage) bool {
	_ = "STUB: not implemented"
	return false
}

func (dev *MetaxDevices) CheckHealth(devType string, n *corev1.Node) (bool, bool) {
	_ = "STUB: not implemented"
	return false, false
}

func (dev *MetaxDevices) GenerateResourceRequests(ctr *corev1.Container) device.ContainerDeviceRequest {
	_ = "STUB: not implemented"
	return *new(device.ContainerDeviceRequest)
}

func (dev *MetaxDevices) customFilterRule(allocated *device.PodDevices, request device.ContainerDeviceRequest, toAllocate device.ContainerDevices, device *device.DeviceUsage) bool {
	_ = "STUB: not implemented"
	return false
}

func parseMetaxAnnos(annos string, index int) float32 { _ = "STUB: not implemented"; return 0 }

func (dev *MetaxDevices) ScoreNode(node *corev1.Node, podDevices device.PodSingleDevice, previous []*device.DeviceUsage, policy string) float32 {
	_ = "STUB: not implemented"
	return 0
}

// it's preferred to select the node with lower loss

// it's preferred to select the node with higher score
// But we have to give it a smaller value because of Spread policy

func (dev *MetaxDevices) AddResourceUsage(pod *corev1.Pod, n *device.DeviceUsage, ctr *device.ContainerDevice) error {
	_ = "STUB: not implemented"
	return nil
}

func (mat *MetaxDevices) Fit(devices []*device.DeviceUsage, request device.ContainerDeviceRequest, pod *corev1.Pod, nodeInfo *device.NodeInfo, allocated *device.PodDevices) (bool, map[string]device.ContainerDevices, string) {
	_ = "STUB: not implemented"
	return false, nil, ""
}

//return false, tmpDevs

//This incurs an issue

// Coresreq=100 indicates it want this card exclusively

// You can't allocate core=0 job to an already full GPU

func (dev *MetaxDevices) GetResourceNames() device.ResourceNames {
	_ = "STUB: not implemented"
	return *new(device.ResourceNames)
}
