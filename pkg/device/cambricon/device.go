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

package cambricon

import (
	"flag"

	"github.com/Project-HAMi/HAMi/pkg/device"

	corev1 "k8s.io/api/core/v1"
)

const (
	CambriconMLUDevice     = "MLU"
	CambriconMLUCommonWord = "MLU"
	MluMemSplitLimit       = "CAMBRICON_SPLIT_MEMS"
	MluMemSplitIndex       = "CAMBRICON_SPLIT_VISIBLE_DEVICES"
	MluMemSplitEnable      = "CAMBRICON_SPLIT_ENABLE"
	MLUInUse               = "cambricon.com/use-mlutype"
	MLUNoUse               = "cambricon.com/nouse-mlutype"
	// MLUUseUUID annotation specifies a comma-separated list of MLU UUIDs to use.
	MLUUseUUID = "cambricon.com/use-gpuuuid"
	// MLUNoUseUUID annotation specifies a comma-separated list of MLU UUIDs to exclude.
	MLUNoUseUUID          = "cambricon.com/nouse-gpuuuid"
	DsmluLockTime         = "cambricon.com/dsmlu.lock"
	DsmluProfile          = "CAMBRICON_DSMLU_PROFILE"
	DsmluResourceAssigned = "CAMBRICON_DSMLU_ASSIGNED"
	retry                 = 5
)

var (
	MLUResourceCount  string
	MLUResourceMemory string
	MLUResourceCores  string
)

type CambriconConfig struct {
	ResourceCountName  string `yaml:"resourceCountName"`
	ResourceMemoryName string `yaml:"resourceMemoryName"`
	ResourceCoreName   string `yaml:"resourceCoreName"`
}

type CambriconDevices struct {
}

func ParseConfig(fs *flag.FlagSet) { _ = "STUB: not implemented"; return }

func InitMLUDevice(config CambriconConfig) *CambriconDevices { _ = "STUB: not implemented"; return nil }

func (dev *CambriconDevices) CommonWord() string { _ = "STUB: not implemented"; return "" }

func (dev *CambriconDevices) setNodeLock(node *corev1.Node) error {
	_ = "STUB: not implemented"
	return nil
}

func (dev *CambriconDevices) LockNode(n *corev1.Node, p *corev1.Pod) error {
	_ = "STUB: not implemented"
	return nil
}

func (dev *CambriconDevices) ReleaseNodeLock(n *corev1.Node, p *corev1.Pod) error {
	_ = "STUB: not implemented"
	return nil
}

func (dev *CambriconDevices) NodeCleanUp(nn string) error { _ = "STUB: not implemented"; return nil }

func (dev *CambriconDevices) CheckHealth(devType string, n *corev1.Node) (bool, bool) {
	_ = "STUB: not implemented"
	return false, false
}

func (dev *CambriconDevices) GetNodeDevices(n corev1.Node) ([]*device.DeviceInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (dev *CambriconDevices) AssertNuma(annos map[string]string) bool {
	_ = "STUB: not implemented"
	return false
}

func (dev *CambriconDevices) MutateAdmission(ctr *corev1.Container, p *corev1.Pod) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (dev *CambriconDevices) checkType(annos map[string]string, d device.DeviceUsage, n device.ContainerDeviceRequest) (bool, bool, bool) {
	_ = "STUB: not implemented"
	return false, false, false
}

func (dev *CambriconDevices) GenerateResourceRequests(ctr *corev1.Container) device.ContainerDeviceRequest {
	_ = "STUB: not implemented"
	return *new(device.ContainerDeviceRequest)
}

func (dev *CambriconDevices) PatchAnnotations(pod *corev1.Pod, annoinput *map[string]string, pd device.PodDevices) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func (dev *CambriconDevices) ScoreNode(node *corev1.Node, podDevices device.PodSingleDevice, previous []*device.DeviceUsage, policy string) float32 {
	_ = "STUB: not implemented"
	return 0
}

func (dev *CambriconDevices) AddResourceUsage(pod *corev1.Pod, n *device.DeviceUsage, ctr *device.ContainerDevice) error {
	_ = "STUB: not implemented"
	return nil
}

func (cam *CambriconDevices) Fit(devices []*device.DeviceUsage, request device.ContainerDeviceRequest, pod *corev1.Pod, nodeInfo *device.NodeInfo, allocated *device.PodDevices) (bool, map[string]device.ContainerDevices, string) {
	_ = "STUB: not implemented"
	return false, nil, ""
}

//return false, tmpDevs

//This incurs an issue

// Coresreq=100 indicates it want this card exclusively

// You can't allocate core=0 job to an already full GPU

func (dev *CambriconDevices) GetResourceNames() device.ResourceNames {
	_ = "STUB: not implemented"
	return *new(device.ResourceNames)
}
