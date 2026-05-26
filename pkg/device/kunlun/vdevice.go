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
	XPUDevice      = "XPU"
	XPUCommonWord  = "XPU"
	NodeLock       = "hami.io/mutex.lock"
	RegisterAnnos  = "hami.io/node-register-xpu"
	HandshakeAnnos = "hami.io/node-handshake-xpu"
	UseUUIDAnno    = "hami.io/use-xpu-uuid"
	NoUseUUIDAnno  = "hami.io/no-use-xpu-uuid"
)

const (
	KunlunMaxMemory = 98304
)

type KunlunVDevices struct {
}

func InitKunlunVDevice(config KunlunConfig) *KunlunVDevices { _ = "STUB: not implemented"; return nil }

func (dev *KunlunVDevices) trimMemory(m int64) int64 { _ = "STUB: not implemented"; return 0 }

func (dev *KunlunVDevices) CommonWord() string { _ = "STUB: not implemented"; return "" }

func (dev *KunlunVDevices) MutateAdmission(ctr *corev1.Container, p *corev1.Pod) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (dev *KunlunVDevices) CheckHealth(devType string, n *corev1.Node) (bool, bool) {
	_ = "STUB: not implemented"
	return false, false
}

func (dev *KunlunVDevices) NodeCleanUp(nn string) error { _ = "STUB: not implemented"; return nil }

func (dev *KunlunVDevices) GetNodeDevices(n corev1.Node) ([]*device.DeviceInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (dev *KunlunVDevices) PatchAnnotations(pod *corev1.Pod, annoinput *map[string]string, pd device.PodDevices) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func (dev *KunlunVDevices) LockNode(n *corev1.Node, p *corev1.Pod) error {
	_ = "STUB: not implemented"
	return nil
}

func (dev *KunlunVDevices) ReleaseNodeLock(n *corev1.Node, p *corev1.Pod) error {
	_ = "STUB: not implemented"
	return nil
}

func (dev *KunlunVDevices) CheckType(annos map[string]string, d device.DeviceUsage, n device.ContainerDeviceRequest) (bool, bool) {
	_ = "STUB: not implemented"
	return false, false
}

func (dev *KunlunVDevices) GenerateResourceRequests(ctr *corev1.Container) device.ContainerDeviceRequest {
	_ = "STUB: not implemented"
	return *new(device.ContainerDeviceRequest)
}

//int32(dev.config.MemoryMax),

func (dev *KunlunVDevices) ScoreNode(node *corev1.Node, podDevices device.PodSingleDevice, previous []*device.DeviceUsage, policy string) float32 {
	_ = "STUB: not implemented"
	return 0
}

func (dev *KunlunVDevices) AddResourceUsage(pod *corev1.Pod, n *device.DeviceUsage, ctr *device.ContainerDevice) error {
	_ = "STUB: not implemented"
	return nil
}

func (dev *KunlunVDevices) GetResourceNames() device.ResourceNames {
	_ = "STUB: not implemented"
	return *new(device.ResourceNames)
}

func (dev *KunlunVDevices) Fit(devices []*device.DeviceUsage, request device.ContainerDeviceRequest, pod *corev1.Pod, nodeInfo *device.NodeInfo, allocated *device.PodDevices) (bool, map[string]device.ContainerDevices, string) {
	_ = "STUB: not implemented"
	return false, nil, ""
}

func FitVXPU(device *device.DeviceUsage, request device.ContainerDeviceRequest) bool {
	_ = "STUB: not implemented"
	return false
}
