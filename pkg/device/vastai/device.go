/*
Copyright 2026 The HAMi Authors.

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

package vastai

import (
	"github.com/Project-HAMi/HAMi/pkg/device"

	corev1 "k8s.io/api/core/v1"
)

type VastaiDevices struct {
}

const (
	HandshakeAnnos   = "hami.io/node-handshake-va"
	RegisterAnnos    = "hami.io/node-va-register"
	VastaiDevice     = "Vastai"
	VastaiCommonWord = "Vastai"
	VastaiInUse      = "vastaitech.com/use-va"
	VastaiNoUse      = "vastaitech.com/nouse-va"
	VastaiUseUUID    = "vastaitech.com/use-gpuuuid"
	VastaiNoUseUUID  = "vastaitech.com/nouse-gpuuuid"
)

var (
	VastaiResourceCount string
)

type VastaiConfig struct {
	ResourceCountName string `yaml:"resourceCountName"`
}

func InitVastaiDevice(config VastaiConfig) *VastaiDevices { _ = "STUB: not implemented"; return nil }

func (dev *VastaiDevices) CommonWord() string { _ = "STUB: not implemented"; return "" }

func (dev *VastaiDevices) GetNodeDevices(n corev1.Node) ([]*device.DeviceInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// only for calscore use

func (dev *VastaiDevices) MutateAdmission(ctr *corev1.Container, p *corev1.Pod) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (dev *VastaiDevices) LockNode(n *corev1.Node, p *corev1.Pod) error {
	_ = "STUB: not implemented"
	return nil
}

func (dev *VastaiDevices) ReleaseNodeLock(n *corev1.Node, p *corev1.Pod) error {
	_ = "STUB: not implemented"
	return nil
}

func (dev *VastaiDevices) NodeCleanUp(nn string) error { _ = "STUB: not implemented"; return nil }

func (dev *VastaiDevices) checkType(annos map[string]string, d device.DeviceUsage, n device.ContainerDeviceRequest) (bool, bool, bool) {
	_ = "STUB: not implemented"
	return false, false, false
}

func (dev *VastaiDevices) CheckHealth(devType string, n *corev1.Node) (bool, bool) {
	_ = "STUB: not implemented"
	return false, false
}

func (dev *VastaiDevices) GenerateResourceRequests(ctr *corev1.Container) device.ContainerDeviceRequest {
	_ = "STUB: not implemented"
	return *new(device.ContainerDeviceRequest)
}

func (dev *VastaiDevices) PatchAnnotations(pod *corev1.Pod, annoinput *map[string]string, pd device.PodDevices) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func (dev *VastaiDevices) ScoreNode(node *corev1.Node, podDevices device.PodSingleDevice, previous []*device.DeviceUsage, policy string) float32 {
	_ = "STUB: not implemented"
	return 0
}

func (dev *VastaiDevices) AddResourceUsage(pod *corev1.Pod, n *device.DeviceUsage, ctr *device.ContainerDevice) error {
	_ = "STUB: not implemented"
	return nil
}

func (va *VastaiDevices) Fit(devices []*device.DeviceUsage, request device.ContainerDeviceRequest, pod *corev1.Pod, nodeInfo *device.NodeInfo, allocated *device.PodDevices) (bool, map[string]device.ContainerDevices, string) {
	_ = "STUB: not implemented"
	return false, nil, ""
}

// If requesting multiple devices, select the best combination of cards.

func (dev *VastaiDevices) GetResourceNames() device.ResourceNames {
	_ = "STUB: not implemented"
	return *new(device.ResourceNames)
}

func (dev *VastaiDevices) computeBestCombination(reqNum int, containerDevices device.ContainerDevices) device.ContainerDevices {
	_ = "STUB: not implemented"
	return *new(device.ContainerDevices)
}

func isDieMode(devices []*device.DeviceUsage) bool { _ = "STUB: not implemented"; return false }
