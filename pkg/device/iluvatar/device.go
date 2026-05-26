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

package iluvatar

import (
	"flag"

	"github.com/Project-HAMi/HAMi/pkg/device"

	corev1 "k8s.io/api/core/v1"
)

var (
	enableIluvatar bool
)

type IluvatarDevices struct {
	config           IluvatarConfig
	nodeRegisterAnno string
	useUUIDAnno      string
	noUseUUIDAnno    string
	handshakeAnno    string
}

type IluvatarConfig struct {
	CommonWord         string `yaml:"commonWord"`
	ChipName           string `yaml:"chipName"`
	ResourceCountName  string `yaml:"resourceCountName"`
	ResourceMemoryName string `yaml:"resourceMemoryName"`
	ResourceCoreName   string `yaml:"resourceCoreName"`
}

func InitIluvatarDevice(config []IluvatarConfig) []*IluvatarDevices {
	_ = "STUB: not implemented"
	return nil
}

func (dev *IluvatarDevices) CommonWord() string { _ = "STUB: not implemented"; return "" }

func ParseConfig(fs *flag.FlagSet) { _ = "STUB: not implemented"; return }

func (dev *IluvatarDevices) MutateAdmission(ctr *corev1.Container, p *corev1.Pod) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Note: This should be "SOL_CONTAINER_NAME", but keeping the incorrect "SOL_CONTINER_NAME" for backward compatibility

func (dev *IluvatarDevices) GetNodeDevices(n corev1.Node) ([]*device.DeviceInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (dev *IluvatarDevices) PatchAnnotations(pod *corev1.Pod, annoinput *map[string]string, pd device.PodDevices) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func (dev *IluvatarDevices) LockNode(n *corev1.Node, p *corev1.Pod) error {
	_ = "STUB: not implemented"
	return nil
}

func (dev *IluvatarDevices) ReleaseNodeLock(n *corev1.Node, p *corev1.Pod) error {
	_ = "STUB: not implemented"
	return nil
}

func (dev *IluvatarDevices) NodeCleanUp(nn string) error { _ = "STUB: not implemented"; return nil }

func (dev *IluvatarDevices) checkType(annos map[string]string, d device.DeviceUsage, n device.ContainerDeviceRequest) (bool, bool, bool) {
	_ = "STUB: not implemented"
	return false, false, false
}

func (dev *IluvatarDevices) CheckHealth(devType string, n *corev1.Node) (bool, bool) {
	_ = "STUB: not implemented"
	return false, false
}

func (dev *IluvatarDevices) GenerateResourceRequests(ctr *corev1.Container) device.ContainerDeviceRequest {
	_ = "STUB: not implemented"
	return *new(device.ContainerDeviceRequest)
}

func (dev *IluvatarDevices) ScoreNode(node *corev1.Node, podDevices device.PodSingleDevice, previous []*device.DeviceUsage, policy string) float32 {
	_ = "STUB: not implemented"
	return 0
}

func (dev *IluvatarDevices) AddResourceUsage(pod *corev1.Pod, n *device.DeviceUsage, ctr *device.ContainerDevice) error {
	_ = "STUB: not implemented"
	return nil
}

func (ilu *IluvatarDevices) Fit(devices []*device.DeviceUsage, request device.ContainerDeviceRequest, pod *corev1.Pod, nodeInfo *device.NodeInfo, allocated *device.PodDevices) (bool, map[string]device.ContainerDevices, string) {
	_ = "STUB: not implemented"
	return false, nil, ""
}

//return false, tmpDevs

//This incurs an issue

// Coresreq=100 indicates it want this card exclusively

// You can't allocate core=0 job to an already full GPU

func (dev *IluvatarDevices) GetResourceNames() device.ResourceNames {
	_ = "STUB: not implemented"
	return *new(device.ResourceNames)
}
