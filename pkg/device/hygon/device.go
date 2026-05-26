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

package hygon

import (
	"flag"

	"github.com/Project-HAMi/HAMi/pkg/device"

	corev1 "k8s.io/api/core/v1"
)

type DCUDevices struct {
}

const (
	HandshakeAnnos     = "hami.io/node-handshake-dcu"
	RegisterAnnos      = "hami.io/node-dcu-register"
	HygonDCUDevice     = "DCU"
	HygonDCUCommonWord = "DCU"
	DCUInUse           = "hygon.com/use-dcutype"
	DCUNoUse           = "hygon.com/nouse-dcutype"
	// DCUUseUUID annotation specifies a comma-separated list of DCU UUIDs to use.
	DCUUseUUID = "hygon.com/use-gpuuuid"
	// DCUNoUseUUID annotation specifies a comma-separated list of DCU UUIDs to exclude.
	DCUNoUseUUID = "hygon.com/nouse-gpuuuid"

	// NodeLockDCU should same with device plugin node lock name
	// there is a bug with nodelock package utils, the key is hard coded as "hami.io/mutex.lock"
	// so we can only use this value now.
	NodeLockDCU = "hami.io/mutex.lock"
)

var (
	HygonResourceCount  string
	HygonResourceMemory string
	HygonResourceCores  string
	MemoryFactor        int32
)

type HygonConfig struct {
	ResourceCountName  string `yaml:"resourceCountName"`
	ResourceMemoryName string `yaml:"resourceMemoryName"`
	ResourceCoreName   string `yaml:"resourceCoreName"`
	MemoryFactor       int32  `yaml:"memoryFactor"`
}

func InitDCUDevice(config HygonConfig) *DCUDevices { _ = "STUB: not implemented"; return nil }

func (dev *DCUDevices) CommonWord() string { _ = "STUB: not implemented"; return "" }

func ParseConfig(fs *flag.FlagSet) { _ = "STUB: not implemented"; return }

func (dev *DCUDevices) MutateAdmission(ctr *corev1.Container, p *corev1.Pod) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func checkDCUtype(annos map[string]string, cardtype string) bool {
	_ = "STUB: not implemented"
	return false
}

func (dev *DCUDevices) LockNode(n *corev1.Node, p *corev1.Pod) error {
	_ = "STUB: not implemented"
	return nil
}

func (dev *DCUDevices) ReleaseNodeLock(n *corev1.Node, p *corev1.Pod) error {
	_ = "STUB: not implemented"
	return nil
}

func (dev *DCUDevices) GetNodeDevices(n corev1.Node) ([]*device.DeviceInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (dev *DCUDevices) NodeCleanUp(nn string) error { _ = "STUB: not implemented"; return nil }

func (dev *DCUDevices) CheckHealth(devType string, n *corev1.Node) (bool, bool) {
	_ = "STUB: not implemented"
	return false, false
}

func (dev *DCUDevices) checkType(annos map[string]string, d device.DeviceUsage, n device.ContainerDeviceRequest) (bool, bool, bool) {
	_ = "STUB: not implemented"
	return false, false, false
}

func (dev *DCUDevices) GenerateResourceRequests(ctr *corev1.Container) device.ContainerDeviceRequest {
	_ = "STUB: not implemented"
	return *new(device.ContainerDeviceRequest)
}

func (dev *DCUDevices) PatchAnnotations(pod *corev1.Pod, annoinput *map[string]string, pd device.PodDevices) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func (dev *DCUDevices) ScoreNode(node *corev1.Node, podDevices device.PodSingleDevice, previous []*device.DeviceUsage, policy string) float32 {
	_ = "STUB: not implemented"
	return 0
}

func (dev *DCUDevices) AddResourceUsage(pod *corev1.Pod, n *device.DeviceUsage, ctr *device.ContainerDevice) error {
	_ = "STUB: not implemented"
	return nil
}

func (dcu *DCUDevices) Fit(devices []*device.DeviceUsage, request device.ContainerDeviceRequest, pod *corev1.Pod, nodeInfo *device.NodeInfo, allocated *device.PodDevices) (bool, map[string]device.ContainerDevices, string) {
	_ = "STUB: not implemented"
	return false, nil, ""
}

//return false, tmpDevs

//This incurs an issue

// Coresreq=100 indicates it want this card exclusively

// You can't allocate core=0 job to an already full GPU

func (dev *DCUDevices) GetResourceNames() device.ResourceNames {
	_ = "STUB: not implemented"
	return *new(device.ResourceNames)
}
