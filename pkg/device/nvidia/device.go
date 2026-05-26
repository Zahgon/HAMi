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

package nvidia

import (
	"flag"
	"sync"

	spec "github.com/NVIDIA/k8s-device-plugin/api/config/v1"
	corev1 "k8s.io/api/core/v1"

	"github.com/Project-HAMi/HAMi/pkg/device"
)

const (
	HandshakeAnnos       = "hami.io/node-handshake"
	RegisterAnnos        = "hami.io/node-nvidia-register"
	RegisterGPUPairScore = "hami.io/node-nvidia-score"
	NvidiaGPUDevice      = "NVIDIA"
	GPUInUse             = "nvidia.com/use-gputype"
	GPUNoUse             = "nvidia.com/nouse-gputype"
	NumaBind             = "nvidia.com/numa-bind"
	NodeLockNvidia       = "hami.io/mutex.lock"
	// GPUUseUUID annotation specifies a comma-separated list of GPU UUIDs to use.
	GPUUseUUID = "nvidia.com/use-gpuuuid"
	// GPUNoUseUUID annotation specifies a comma-separated list of GPU UUIDs to exclude.
	GPUNoUseUUID = "nvidia.com/nouse-gpuuuid"
	AllocateMode = "nvidia.com/vgpu-mode"

	MigMode      = "mig"
	HamiCoreMode = "hami-core"
	MpsMode      = "mps"
)

var (
	NodeName          string
	RuntimeSocketFlag string
	DisableCoreLimit  *bool

	// DevicePluginFilterDevice need device-plugin filter this device, don't register this device.
	DevicePluginFilterDevice *FilterDevice
	MemoryFactor             int32 = 1
)

type MigPartedSpec struct {
	Version    string                        `json:"version"               yaml:"version"`
	MigConfigs map[string]MigConfigSpecSlice `json:"mig-configs,omitempty" yaml:"mig-configs,omitempty"`
}

// MigConfigSpec defines the spec to declare the desired MIG configuration for a set of GPUs.
type MigConfigSpec struct {
	DeviceFilter any              `json:"device-filter,omitempty" yaml:"device-filter,flow,omitempty"`
	Devices      []int32          `json:"devices"                 yaml:"devices,flow"`
	MigEnabled   bool             `json:"mig-enabled"             yaml:"mig-enabled"`
	MigDevices   map[string]int32 `json:"mig-devices"             yaml:"mig-devices"`
}

// MigConfigSpecSlice represents a slice of 'MigConfigSpec'.
type MigConfigSpecSlice []MigConfigSpec

// GPUCoreUtilizationPolicy is set nvidia gpu core isolation policy.
type GPUCoreUtilizationPolicy string

const (
	DefaultCorePolicy GPUCoreUtilizationPolicy = "default"
	ForceCorePolicy   GPUCoreUtilizationPolicy = "force"
	DisableCorePolicy GPUCoreUtilizationPolicy = "disable"
)

type LibCudaLogLevel string

const (
	Error    LibCudaLogLevel = "0"
	Warnings LibCudaLogLevel = "1"
	Infos    LibCudaLogLevel = "3"
	Debugs   LibCudaLogLevel = "4"
)

type NvidiaConfig struct {
	// These configs are shared and can be overwritten by Nodeconfig.
	NodeDefaultConfig            `yaml:",inline"`
	ResourceCountName            string `yaml:"resourceCountName"`
	ResourceMemoryName           string `yaml:"resourceMemoryName"`
	ResourceCoreName             string `yaml:"resourceCoreName"`
	ResourceMemoryPercentageName string `yaml:"resourceMemoryPercentageName"`
	ResourcePriority             string `yaml:"resourcePriorityName"`
	OverwriteEnv                 bool   `yaml:"overwriteEnv"`
	DefaultMemory                int32  `yaml:"defaultMemory"`
	DefaultCores                 int32  `yaml:"defaultCores"`
	DefaultGPUNum                int32  `yaml:"defaultGPUNum"`
	MemoryFactor                 int32  `yaml:"memoryFactor"`
	// TODO Whether these should be removed
	DisableCoreLimit  bool                          `yaml:"disableCoreLimit"`
	MigGeometriesList []device.AllowedMigGeometries `yaml:"knownMigGeometries"`
	// GPUCorePolicy through webhook automatic injected to container env
	GPUCorePolicy GPUCoreUtilizationPolicy `yaml:"gpuCorePolicy"`
	// RuntimeClassName is the name of the runtime class to be added to pod.spec.runtimeClassName
	RuntimeClassName string `yaml:"runtimeClassName"`
}

// These configs can be specified for each node by using Nodeconfig.
type NodeDefaultConfig struct {
	DeviceSplitCount          *uint    `yaml:"deviceSplitCount" json:"devicesplitcount"`
	DeviceMemoryScaling       *float64 `yaml:"deviceMemoryScaling" json:"devicememoryscaling"`
	DeviceCoreScaling         *float64 `yaml:"deviceCoreScaling" json:"devicecorescaling"`
	PreConfiguredDeviceMemory *int64   `yaml:"preConfiguredDeviceMemory" json:"preconfigureddevicememory"`
	// LogLevel is LIBCUDA_LOG_LEVEL value
	LogLevel *LibCudaLogLevel `yaml:"libCudaLogLevel" json:"libcudaloglevel"`
}

type FilterDevice struct {
	// UUID is the device ID.
	UUID []string `json:"uuid"`
	// Index is the device index.
	Index []uint `json:"index"`
}

type DevicePluginConfigs struct {
	Nodeconfig []struct {
		// These configs is shared and will overwrite those in NvidiaConfig.
		NodeDefaultConfig            `json:",inline"`
		Name                         string        `json:"name"`
		OperatingMode                string        `json:"operatingmode"`
		Migstrategy                  string        `json:"migstrategy"`
		FilterDevice                 *FilterDevice `json:"filterdevices"`
		EnableGetPreferredAllocation bool          `json:"enablegetpreferredallocation"`
	} `json:"nodeconfig"`
}

type DeviceConfig struct {
	*spec.Config

	ResourceName *string
	DebugMode    *bool
}

type NvidiaGPUDevices struct {
	config         NvidiaConfig
	ReportedGPUNum map[string]int64 // key: nodeName, value: reported GPU count
	mu             sync.Mutex       // protects concurrent access to ReportedGPUNum
}

func InitNvidiaDevice(nvconfig NvidiaConfig) *NvidiaGPUDevices {
	_ = "STUB: not implemented"
	return nil
}

func (dev *NvidiaGPUDevices) CommonWord() string { _ = "STUB: not implemented"; return "" }

func ParseConfig(fs *flag.FlagSet) { _ = "STUB: not implemented"; return }

func FilterDeviceToRegister(uuid, indexStr string) bool { _ = "STUB: not implemented"; return false }

func (dev *NvidiaGPUDevices) NodeCleanUp(nn string) error { _ = "STUB: not implemented"; return nil }

func (dev *NvidiaGPUDevices) CheckHealth(devType string, n *corev1.Node) (bool, bool) {
	_ = "STUB: not implemented"
	return false, false
}

func (dev *NvidiaGPUDevices) LockNode(n *corev1.Node, p *corev1.Pod) error {
	_ = "STUB: not implemented"
	return nil
}

func (dev *NvidiaGPUDevices) ReleaseNodeLock(n *corev1.Node, p *corev1.Pod) error {
	_ = "STUB: not implemented"
	return nil
}

func (dev *NvidiaGPUDevices) GetNodeDevices(n corev1.Node) ([]*device.DeviceInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// fit pair score to device info

func (dev *NvidiaGPUDevices) MutateAdmission(ctr *corev1.Container, p *corev1.Pod) (bool, error) {
	_ = "STUB: not implemented"
	/*gpu related */ return false, nil
}

// Set runtime class name if it is not set by user and the runtime class name is configured

func (dev *NvidiaGPUDevices) mutateContainerResource(ctr *corev1.Container) bool {
	_ = "STUB: not implemented"
	return false
}

func (dev *NvidiaGPUDevices) defaultExclusiveCoreIfNeeded(ctr *corev1.Container) bool {
	_ = "STUB: not implemented"
	return false
}

func resourceValue(ctr *corev1.Container, name corev1.ResourceName) (int64, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

func resourcePresent(ctr *corev1.Container, name corev1.ResourceName) bool {
	_ = "STUB: not implemented"
	return false
}

func checkGPUtype(annos map[string]string, cardtype string) bool {
	_ = "STUB: not implemented"
	return false
}

func assertNuma(annos map[string]string) bool { _ = "STUB: not implemented"; return false }

func (dev *NvidiaGPUDevices) checkType(annos map[string]string, d device.DeviceUsage, n device.ContainerDeviceRequest) (bool, bool) {
	_ = "STUB: not implemented"
	return false, false
}

func (dev *NvidiaGPUDevices) PatchAnnotations(pod *corev1.Pod, annoinput *map[string]string, pd device.PodDevices) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func (dev *NvidiaGPUDevices) GenerateResourceRequests(ctr *corev1.Container) device.ContainerDeviceRequest {
	_ = "STUB: not implemented"
	return *new(device.ContainerDeviceRequest)
}

func (dev *NvidiaGPUDevices) CustomFilterRule(allocated *device.PodDevices, request device.ContainerDeviceRequest, toAllocate device.ContainerDevices, devusage *device.DeviceUsage) bool {
	_ = "STUB: not implemented"
	//memreq := request.Memreq
	return false
}

// The same logic as in AddResourceUsage

func (dev *NvidiaGPUDevices) ScoreNode(node *corev1.Node, podDevices device.PodSingleDevice, previous []*device.DeviceUsage, policy string) float32 {
	_ = "STUB: not implemented"
	return 0
}

func (dev *NvidiaGPUDevices) migNeedsReset(n *device.DeviceUsage) bool {
	_ = "STUB: not implemented"
	return false
}

func (dev *NvidiaGPUDevices) AddResourceUsage(pod *corev1.Pod, n *device.DeviceUsage, ctr *device.ContainerDevice) error {
	_ = "STUB: not implemented"
	return nil
}

// Calculate the correct UsageList index by summing Count of all templates before idx

func fitQuota(tmpDevs map[string]device.ContainerDevices, allocated *device.PodDevices, ns string, memreq int64, coresreq int64) bool {
	_ = "STUB: not implemented"
	return false
}

func (nv *NvidiaGPUDevices) Fit(devices []*device.DeviceUsage, request device.ContainerDeviceRequest, pod *corev1.Pod, nodeInfo *device.NodeInfo, allocated *device.PodDevices) (bool, map[string]device.ContainerDevices, string) {
	_ = "STUB: not implemented"
	return false, nil, ""
}

//return false, tmpDevs

//This incurs an issue

// Coresreq=100 indicates it want this card exclusively

// You can't allocate core=0 job to an already full GPU

// If requesting a device, select the card with the worst connection to other cards (lowest total score).

// If requesting multiple devices, select the best combination of cards.

func (dev *NvidiaGPUDevices) GetResourceNames() device.ResourceNames {
	_ = "STUB: not implemented"
	return *new(device.ResourceNames)
}

func generateCombinations(request device.ContainerDeviceRequest, tmpDevs map[string]device.ContainerDevices) []device.ContainerDevices {
	_ = "STUB: not implemented"
	return nil
}

// This code mainly performs permutations and combinations to generate all non-repetitive subsets.
// For example, if the request is 3 GPUs, and the available GPUs are ["GPU0", "GPU1", "GPU2", "GPU3"],
// it will generate all combinations of 3 GPUs from the available GPUs.
// Result：[["GPU0", "GPU1", "GPU2"],["GPU0", "GPU1", "GPU3"],["GPU0", "GPU2", "GPU3"],["GPU1", "GPU2", "GPU3"]]

func getDevicePairScoreMap(nodeInfo *device.NodeInfo) map[string]*device.DevicePairScore {
	_ = "STUB: not implemented"
	return nil
}

func computeWorstSingleCard(nodeInfo *device.NodeInfo, request device.ContainerDeviceRequest, tmpDevs map[string]device.ContainerDevices) device.ContainerDevices {
	_ = "STUB: not implemented"
	return *new(device.ContainerDevices)
}

// Iterate through all devices to find the one with the lowest score

func computeBestCombination(nodeInfo *device.NodeInfo, combinations []device.ContainerDevices) device.ContainerDevices {
	_ = "STUB: not implemented"
	return *new(device.ContainerDevices)
}

// Iterate through all combinations to find the one with the highest score
