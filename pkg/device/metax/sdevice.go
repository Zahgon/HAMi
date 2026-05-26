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

	corev1 "k8s.io/api/core/v1"
)

const (
	MetaxSGPUCommonWord = "Metax-SGPU"
	MetaxSGPUDevice     = "Metax-SGPU"

	MetaxNodeLock = "hami.io/mutex.lock"
)

const (
	CardUnhealthy         = "CardUnhealthy"
	CardQosPolicyMismatch = "CardQosPolicyMismatch"
	CardAppClassMismatch  = "CardAppClassMismatch"
)

var (
	MetaxResourceNameVCount  string
	MetaxResourceNameVCore   string
	MetaxResourceNameVMemory string
	MetaxTopologyAware       bool
)

type MetaxSDevices struct {
	jqCache *JitteryQosCache
}

func InitMetaxSDevice(config MetaxConfig) *MetaxSDevices { _ = "STUB: not implemented"; return nil }

func (sdev *MetaxSDevices) CommonWord() string { _ = "STUB: not implemented"; return "" }

func (sdev *MetaxSDevices) MutateAdmission(ctr *corev1.Container, p *corev1.Pod) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (sdev *MetaxSDevices) GetNodeDevices(n corev1.Node) ([]*device.DeviceInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (sdev *MetaxSDevices) PatchAnnotations(pod *corev1.Pod, annoinput *map[string]string, pd device.PodDevices) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

// hami

// metax

func (sdev *MetaxSDevices) LockNode(n *corev1.Node, p *corev1.Pod) error {
	_ = "STUB: not implemented"
	return nil
}

func (sdev *MetaxSDevices) ReleaseNodeLock(n *corev1.Node, p *corev1.Pod) error {
	_ = "STUB: not implemented"
	return nil
}

func (sdev *MetaxSDevices) NodeCleanUp(nn string) error { _ = "STUB: not implemented"; return nil }

func (sdev *MetaxSDevices) checkType(annos map[string]string, d device.DeviceUsage, n device.ContainerDeviceRequest) bool {
	_ = "STUB: not implemented"
	return false
}

func (sdev *MetaxSDevices) CheckHealth(devType string, n *corev1.Node) (bool, bool) {
	_ = "STUB: not implemented"
	return false, false
}

func (sdev *MetaxSDevices) GenerateResourceRequests(ctr *corev1.Container) device.ContainerDeviceRequest {
	_ = "STUB: not implemented"
	return *new(device.ContainerDeviceRequest)
}

// if user not set unit, default unit is Gi

func (sdev *MetaxSDevices) ScoreNode(node *corev1.Node, podDevices device.PodSingleDevice, previous []*device.DeviceUsage, policy string) float32 {
	_ = "STUB: not implemented"
	// TODO: score should not depend on policy
	// we have to give it a smaller value because of Spread policy
	return 0
}

func (sdev *MetaxSDevices) AddResourceUsage(pod *corev1.Pod, n *device.DeviceUsage, ctr *device.ContainerDevice) error {
	_ = "STUB: not implemented"
	return nil
}

func (mats *MetaxSDevices) Fit(devices []*device.DeviceUsage, request device.ContainerDeviceRequest, pod *corev1.Pod, nodeInfo *device.NodeInfo, allocated *device.PodDevices) (bool, map[string]device.ContainerDevices, string) {
	_ = "STUB: not implemented"
	return false, nil, ""
}

// filter device

// Coresreq=100 indicates it want this card exclusively

// You can't allocate core=0 job to an already full GPU

// prioritize device

// online Pod need additional logic to prioritize devices,
// offline Pod follow hami gpu scheduling policy

// generate containerDevice

// WorkAround: add Pod annotations into ContainerDevice to pass annotations in `ScoreNode`

func (dev *MetaxSDevices) GetResourceNames() device.ResourceNames {
	_ = "STUB: not implemented"
	return *new(device.ResourceNames)
}

func (sdev *MetaxSDevices) getMetaxSDevices(n corev1.Node) ([]*MetaxSDeviceInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (sdev *MetaxSDevices) checkDeviceQos(reqQos string, usage device.DeviceUsage, request device.ContainerDeviceRequest) bool {
	_ = "STUB: not implemented"
	return false
}

func (sdev *MetaxSDevices) addJitteryQos(reqQos string, devs device.PodSingleDevice) {
	_ = "STUB: not implemented"
	return
}

func prioritizeExclusiveDevices(candidateDevices []*device.DeviceUsage, require int) []*device.DeviceUsage {
	_ = "STUB: not implemented"
	return nil
}

// 1. pickup devices within MetaLink

// 2. pickup devices cross MetaLink

// 3. if not satisfied, pick up devices no MetaLink

func topologyAwareEnable(podDevices device.PodSingleDevice) bool {
	_ = "STUB: not implemented"
	return false
}

func scoreExclusiveDevices(podDevices device.PodSingleDevice, previous []*device.DeviceUsage) int {
	_ = "STUB: not implemented"
	return 0
}

func checkAppClass(requestClass string, dev device.DeviceUsage) bool {
	_ = "STUB: not implemented"
	return false
}

// Online/Offline Pod can't scheduled to normal device
// normal Pod can't scheduled to device that runs Online/Offline Pod
// once device start to running Online Pod, Offline Pod in device will be suspended
// there is no need to schedule Offline Pod to device that runs Online Pod

func appClassOnlineEnable(podDevices device.PodSingleDevice) bool {
	_ = "STUB: not implemented"
	return false
}

func prioritizeOnlineDevices(candidateDevices []*device.DeviceUsage, require int) []*device.DeviceUsage {
	_ = "STUB: not implemented"
	return nil
}

// try choose device with fewer online tasks,
// and try not to choose device that only have offline tasks

func scoreOnlineDevices(podDevices device.PodSingleDevice, previous []*device.DeviceUsage) int {
	_ = "STUB: not implemented"
	return 0
}
