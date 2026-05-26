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

package ascend

import (
	"flag"

	"github.com/Project-HAMi/HAMi/pkg/device"

	corev1 "k8s.io/api/core/v1"
)

const (
	NodeLockAscend             = "hami.io/mutex.lock"
	Ascend910Prefix            = "Ascend910"
	Ascend910CType             = "Ascend910C"
	Ascend910NetworkWeight     = 10
	VNPUModeAnnotation         = "huawei.com/vnpu-mode"
	VNPUModeHamiCore           = "hami-core"
	VNPUNodeSelectorAnnotation = "hami-vnpu-core"
)

type Devices struct {
	config           VNPUConfig
	nodeRegisterAnno string
	useUUIDAnno      string
	noUseUUIDAnno    string
	handshakeAnno    string
	hamiVnpuCore     bool
}

type RuntimeInfo struct {
	UUID   string `json:"UUID,omitempty"`
	Temp   string `json:"temp,omitempty"`
	Memory int64  `json:"memory,omitempty"`
	Core   int32  `json:"core,omitempty"`
}

var (
	enableAscend bool
	configFile   string
)

func (dev *Devices) trimMemory(m int64) (int64, string) { _ = "STUB: not implemented"; return 0, "" }

func InitDevices(vnpus VNPUs) []*Devices { _ = "STUB: not implemented"; return nil }

func ParseConfig(fs *flag.FlagSet) { _ = "STUB: not implemented"; return }

func (dev *Devices) CommonWord() string { _ = "STUB: not implemented"; return "" }

func (dev *Devices) MutateAdmission(ctr *corev1.Container, p *corev1.Pod) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Since the minimum allocation unit is one physical module (2 NPUs), round up the limits and requests to 2.

// Reject any other odd-numbered request (e.g., 3, 5, 7...)

// Check if hami-core is declared

// Inject PostStart hook to start the limiter process

// Set runtime class name if it is not set by user and the runtime class name is configured

func (dev *Devices) GetNodeDevices(n corev1.Node) ([]*device.DeviceInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (dev *Devices) PatchAnnotations(pod *corev1.Pod, annoInput *map[string]string, pd device.PodDevices) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

// Check if hami-core is declared

// If is hami core, populate Memory and Core directly without using Temp

func (dev *Devices) LockNode(n *corev1.Node, p *corev1.Pod) error {
	_ = "STUB: not implemented"
	return nil
}

func (dev *Devices) ReleaseNodeLock(n *corev1.Node, p *corev1.Pod) error {
	_ = "STUB: not implemented"
	return nil
}

func (dev *Devices) NodeCleanUp(nn string) error { _ = "STUB: not implemented"; return nil }

func (dev *Devices) checkType(annos map[string]string, d device.DeviceUsage, n device.ContainerDeviceRequest) (bool, bool, bool) {
	_ = "STUB: not implemented"
	return false, false, false
}

func (dev *Devices) CheckHealth(devType string, n *corev1.Node) (bool, bool) {
	_ = "STUB: not implemented"
	return false, false
}

func (dev *Devices) GenerateResourceRequests(ctr *corev1.Container) device.ContainerDeviceRequest {
	_ = "STUB: not implemented"
	return *new(device.ContainerDeviceRequest)
}

// If "core" is requested, it explicitly indicates the use of soft-partitioning.

// Soft-partitioning: Use the raw value directly.

// Process Core Resources

func (dev *Devices) ScoreNode(node *corev1.Node, podDevices device.PodSingleDevice, previous []*device.DeviceUsage, policy string) float32 {
	_ = "STUB: not implemented"
	return 0
}

func (dev *Devices) AddResourceUsage(pod *corev1.Pod, n *device.DeviceUsage, ctr *device.ContainerDevice) error {
	_ = "STUB: not implemented"
	return nil
}

func (dev *Devices) GetResourceNames() device.ResourceNames {
	_ = "STUB: not implemented"
	return *new(device.ResourceNames)
}

func (npu *Devices) Fit(devices []*device.DeviceUsage, request device.ContainerDeviceRequest, pod *corev1.Pod, nodeInfo *device.NodeInfo, allocated *device.PodDevices) (bool, map[string]device.ContainerDevices, string) {
	_ = "STUB: not implemented"
	return false, nil, ""
}

// Verify whether the Node supports hami vnpu core.
// Global hamiVnpuCore config acts as the default; node-level annotation takes higher priority.

//return false, tmpDevs

//This incurs an issue

// Set dev.Totalcore to 100 if vnpuMode is hami-core

// Coresreq=100 indicates it want this card exclusively

// You can't allocate core=0 job to an already full GPU

// If requesting multiple devices, select the best combination of cards.

// Use topology-aware allocation for Ascend910C: only select full modules (2 NPUs per card).

func hasNetworkID(devices []*device.DeviceUsage) bool { _ = "STUB: not implemented"; return false }

func (npudev *Devices) computeBestCombination(nodeInfo *device.NodeInfo, reqNum int, containerDevices device.ContainerDevices) device.ContainerDevices {
	_ = "STUB: not implemented"
	return *new(device.ContainerDevices)
}

func (npudev *Devices) computeBestCombination910C(nodeInfo *device.NodeInfo, reqNum int, containerDevices device.ContainerDevices) device.ContainerDevices {
	_ = "STUB: not implemented"
	// Build a mapping from NPU index to device object for quick lookup.
	return *new(device.ContainerDevices)
}

// Each physical card hosts exactly 2 NPUs (Ascend 910C module design).

// Group NPU indices by the module and Sort

// Convert the card topology map into a slice for sorting.

// Sort cards by the number of available NPUs in ascending order.

// Select NPUs card by card, preferring full cards.

// Only consider cards that have both NPUs available (full card).
