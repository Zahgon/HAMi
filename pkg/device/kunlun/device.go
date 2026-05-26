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

	corev1 "k8s.io/api/core/v1"
)

const (
	KunlunGPUDevice       = "kunlun"
	KunlunGPUCommonWord   = "kunlun"
	KunlunDeviceSelection = "BAIDU_COM_DEVICE_IDX"
	KunlunUseUUID         = "baidu.com/use-gpuuuid"
	KunlunNoUseUUID       = "baidu.com/nouse-gpuuuid"
	InterGroupConnection  = "0-4,1-5,2-6,3-7"
	InterGroupConnection2 = "0-1-4-5,2-3-6-7,0-2-4-6,1-3-5-7,0-3-4-7,1-2-5-6"
	GroupConnection       = "0-1,0-2,0-3,1-2,1-3,2-3,4-5,4-6,4-7,5-6,5-7,6-7"
)

type KunlunDevices struct {
}

func InitKunlunDevice(config KunlunConfig) *KunlunDevices { _ = "STUB: not implemented"; return nil }

func (dev *KunlunDevices) CommonWord() string { _ = "STUB: not implemented"; return "" }

func (dev *KunlunDevices) MutateAdmission(ctr *corev1.Container, p *corev1.Pod) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (dev *KunlunDevices) GetNodeDevices(n corev1.Node) ([]*device.DeviceInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (dev *KunlunDevices) PatchAnnotations(pod *corev1.Pod, annoinput *map[string]string, pd device.PodDevices) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func (dev *KunlunDevices) LockNode(n *corev1.Node, p *corev1.Pod) error {
	_ = "STUB: not implemented"
	return nil
}

func (dev *KunlunDevices) ReleaseNodeLock(n *corev1.Node, p *corev1.Pod) error {
	_ = "STUB: not implemented"
	return nil
}

func (dev *KunlunDevices) NodeCleanUp(nn string) error { _ = "STUB: not implemented"; return nil }

func (dev *KunlunDevices) CheckType(annos map[string]string, d device.DeviceUsage, n device.ContainerDeviceRequest) (bool, bool) {
	_ = "STUB: not implemented"
	return false, false
}

func (dev *KunlunDevices) CheckHealth(devType string, n *corev1.Node) (bool, bool) {
	_ = "STUB: not implemented"
	return false, false
}

func (dev *KunlunDevices) GenerateResourceRequests(ctr *corev1.Container) device.ContainerDeviceRequest {
	_ = "STUB: not implemented"
	return *new(device.ContainerDeviceRequest)
}

func (dev *KunlunDevices) ScoreNode(node *corev1.Node, podDevices device.PodSingleDevice, previous []*device.DeviceUsage, policy string) float32 {
	_ = "STUB: not implemented"
	return 0
}

func (dev *KunlunDevices) AddResourceUsage(pod *corev1.Pod, n *device.DeviceUsage, ctr *device.ContainerDevice) error {
	_ = "STUB: not implemented"
	return nil
}

func (kl *KunlunDevices) Fit(devices []*device.DeviceUsage, request device.ContainerDeviceRequest, pod *corev1.Pod, nodeInfo *device.NodeInfo, allocated *device.PodDevices) (bool, map[string]device.ContainerDevices, string) {
	_ = "STUB: not implemented"
	return false, nil, ""
}

func (dev *KunlunDevices) GetResourceNames() device.ResourceNames {
	_ = "STUB: not implemented"
	return *new(device.ResourceNames)
}

func FitXPU(device *device.DeviceUsage, request device.ContainerDeviceRequest) bool {
	_ = "STUB: not implemented"
	return false
}
