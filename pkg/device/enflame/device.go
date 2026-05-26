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

package enflame

import (
	"github.com/Project-HAMi/HAMi/pkg/device"

	corev1 "k8s.io/api/core/v1"
)

type EnflameDevices struct{}

const (
	EnflameVGCUDevice     = "Enflame"
	EnflameVGCUCommonWord = "Enflame"
	// EnflameUseUUID annotation specifies a comma-separated list of Enflame UUIDs to use.
	EnflameUseUUID = "enflame.com/use-gpuuuid"
	// EnflameNoUseUUID annotation specifies a comma-separated list of Enflame UUIDs to exclude.
	EnflameNoUseUUID   = "enflame.com/nouse-gpuuuid"
	PodRequestGCUSize  = "enflame.com/gcu-request-size"
	PodAssignedGCUID   = "enflame.com/gcu-assigned-id"
	PodHasAssignedGCU  = "enflame.com/gcu-assigned"
	PodAssignedGCUIdx  = "enflame.com/gcu-assigned-index"
	PodAssignedGCUMin  = "enflame.com/gcu-assigned-minor"
	PodAssignedGCUTime = "enflame.com/gcu-assigned-time"
	AssignedContainers = "assigned-containers"
	GCUDrsCapacity     = "enflame.com/gcu-drs-capacity"

	SharedResourceName = "enflame.com/shared-gcu"
	CountNoSharedName  = "enflame.com/gcu-count"

	enflameRequestModeDirect  int32 = 0
	enflameRequestModeBySpec  int32 = -1
	enflameUnknownCoreRequest int32 = 0
)

type drsCapacitySpec struct {
	Devices  []drsDeviceSpec   `json:"devices"`
	Profiles map[string]string `json:"profiles"`
}

type drsDeviceSpec struct {
	Index    string `json:"index"`
	Minor    string `json:"minor"`
	Capacity any    `json:"capacity"`
}

type assignedContainerInfo struct {
	Allocated    bool   `json:"allocated"`
	Request      int32  `json:"request"`
	ProfileID    string `json:"profileID,omitempty"`
	ProfileName  string `json:"profileName,omitempty"`
	InstanceID   string `json:"instanceID,omitempty"`
	InstanceUUID string `json:"instanceUUID,omitempty"`
}

func InitEnflameDevice(config EnflameConfig) *EnflameDevices { _ = "STUB: not implemented"; return nil }

func (dev *EnflameDevices) CommonWord() string { _ = "STUB: not implemented"; return "" }

func (dev *EnflameDevices) MutateAdmission(ctr *corev1.Container, p *corev1.Pod) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Direct DRS API: enflame.com/drs-gcu

// Unified API: request by memory/core, then convert profile in Fit().

func (dev *EnflameDevices) GetNodeDevices(n corev1.Node) ([]*device.DeviceInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (dev *EnflameDevices) PatchAnnotations(pod *corev1.Pod, annoinput *map[string]string, pd device.PodDevices) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func (dev *EnflameDevices) LockNode(n *corev1.Node, p *corev1.Pod) error {
	_ = "STUB: not implemented"
	return nil
}

func (dev *EnflameDevices) ReleaseNodeLock(n *corev1.Node, p *corev1.Pod) error {
	_ = "STUB: not implemented"
	return nil
}

func (dev *EnflameDevices) NodeCleanUp(nn string) error { _ = "STUB: not implemented"; return nil }

func (dev *EnflameDevices) checkType(annos map[string]string, d device.DeviceUsage, n device.ContainerDeviceRequest) (bool, bool, bool) {
	_ = "STUB: not implemented"
	return false, false, false
}

func (dev *EnflameDevices) CheckHealth(devType string, n *corev1.Node) (bool, bool) {
	_ = "STUB: not implemented"
	return false, false
}

func (dev *EnflameDevices) GenerateResourceRequests(ctr *corev1.Container) device.ContainerDeviceRequest {
	_ = "STUB: not implemented"
	return *new(device.ContainerDeviceRequest)
}

func (dev *EnflameDevices) ScoreNode(node *corev1.Node, podDevices device.PodSingleDevice, previous []*device.DeviceUsage, policy string) float32 {
	_ = "STUB: not implemented"
	return 0
}

func (dev *EnflameDevices) AddResourceUsage(pod *corev1.Pod, n *device.DeviceUsage, ctr *device.ContainerDevice) error {
	_ = "STUB: not implemented"
	return nil
}

func (enf *EnflameDevices) Fit(devices []*device.DeviceUsage, request device.ContainerDeviceRequest, pod *corev1.Pod, nodeInfo *device.NodeInfo, allocated *device.PodDevices) (bool, map[string]device.ContainerDevices, string) {
	_ = "STUB: not implemented"
	return false, nil, ""
}

func (dev *EnflameDevices) GetResourceNames() device.ResourceNames {
	_ = "STUB: not implemented"
	return *new(device.ResourceNames)
}

type drsProfileCandidate struct {
	Name               string
	ID                 string
	Size               int
	MemoryGB           int
	CorePercent        int
	RequestMemoryGB    int
	RequestCorePercent int
}

func (dev *EnflameDevices) selectProfileByRequest(devices []*device.DeviceUsage, request device.ContainerDeviceRequest) (drsProfileCandidate, bool) {
	_ = "STUB: not implemented"
	return *new(drsProfileCandidate), false
}

func parseProfilesFromCustomInfo(customInfo map[string]any) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func collectDRSProfiles(devices []*device.DeviceUsage) []drsProfileCandidate {
	_ = "STUB: not implemented"
	return nil
}

func parseProfile(profileName string) (int, int) { _ = "STUB: not implemented"; return 0, 0 }

func normalizeMemoryRequestToGB(rawMemory int32, maxProfileMemoryGB int) int {
	_ = "STUB: not implemented"
	return 0
}

// If the request value is much larger than profile-GB units, treat it as MiB.

func parseDRSCapacity(raw any) (int32, error) { _ = "STUB: not implemented"; return 0, nil }

func containerNameByIndex(pod *corev1.Pod, index int) string { _ = "STUB: not implemented"; return "" }

func readCustomInfoString(customInfo map[string]any, key string) string {
	_ = "STUB: not implemented"
	return ""
}

func readCustomInfoInt(customInfo map[string]any, key string) int {
	_ = "STUB: not implemented"
	return 0
}

func getContainerResourceRequest(ctr *corev1.Container, resourceName corev1.ResourceName) (int64, bool) {
	_ = "STUB: not implemented"
	return 0, false
}
