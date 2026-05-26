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

package device

import (
	corev1 "k8s.io/api/core/v1"
)

type Devices interface {
	CommonWord() string
	MutateAdmission(ctr *corev1.Container, pod *corev1.Pod) (bool, error)
	CheckHealth(devType string, n *corev1.Node) (bool, bool)
	NodeCleanUp(nn string) error
	GetResourceNames() ResourceNames
	GetNodeDevices(n corev1.Node) ([]*DeviceInfo, error)
	LockNode(n *corev1.Node, p *corev1.Pod) error
	ReleaseNodeLock(n *corev1.Node, p *corev1.Pod) error
	GenerateResourceRequests(ctr *corev1.Container) ContainerDeviceRequest
	PatchAnnotations(pod *corev1.Pod, annoinput *map[string]string, pd PodDevices) map[string]string
	ScoreNode(node *corev1.Node, podDevices PodSingleDevice, previous []*DeviceUsage, policy string) float32
	AddResourceUsage(pod *corev1.Pod, n *DeviceUsage, ctr *ContainerDevice) error
	Fit(devices []*DeviceUsage, request ContainerDeviceRequest, pod *corev1.Pod, nodeInfo *NodeInfo, allocated *PodDevices) (bool, map[string]ContainerDevices, string)
}

type MigTemplate struct {
	Name   string `yaml:"name"`
	Core   int32  `yaml:"core"`
	Memory int32  `yaml:"memory"`
	Count  int32  `yaml:"count"`
}

type MigTemplateUsage struct {
	Name   string `json:"name,omitempty"`
	Core   int32  `json:"core,omitempty"`
	Memory int32  `json:"memory,omitempty"`
	InUse  bool   `json:"inuse,omitempty"`
}

type Geometry []MigTemplate

type MIGS []MigTemplateUsage

type MigInUse struct {
	Index     int32
	UsageList MIGS
}

type AllowedMigGeometries struct {
	Models     []string   `yaml:"models"`
	Geometries []Geometry `yaml:"allowedGeometries"`
}

type DeviceUsage struct {
	ID          string
	Index       uint
	Used        int32
	Count       int32
	Usedmem     int32
	Totalmem    int32
	Totalcore   int32
	Usedcores   int32
	Mode        string
	MigTemplate []Geometry
	MigUsage    MigInUse
	Numa        int
	Type        string
	Health      bool
	PodInfos    []*PodInfo
	CustomInfo  map[string]any
}

type DeviceInfo struct {
	ID              string          `json:"id,omitempty"`
	Index           uint            `json:"index,omitempty"`
	Count           int32           `json:"count,omitempty"`
	Devmem          int32           `json:"devmem,omitempty"`
	Devcore         int32           `json:"devcore,omitempty"`
	Type            string          `json:"type,omitempty"`
	Numa            int             `json:"numa,omitempty"`
	Mode            string          `json:"mode,omitempty"`
	MIGTemplate     []Geometry      `json:"migtemplate,omitempty"`
	Health          bool            `json:"health,omitempty"`
	DeviceVendor    string          `json:"devicevendor,omitempty"`
	CustomInfo      map[string]any  `json:"custominfo,omitempty"`
	DevicePairScore DevicePairScore `json:"devicepairscore,omitempty"`
}

type DevicePairScores []DevicePairScore
type DevicePairScore struct {
	ID     string         `json:"uuid,omitempty"`
	Scores map[string]int `json:"score,omitempty"`
}

type NodeInfo struct {
	ID      string
	Node    *corev1.Node
	Devices map[string][]DeviceInfo
}

type ResourceNames struct {
	ResourceCountName  string
	ResourceMemoryName string
	ResourceCoreName   string
}

type ContainerDevice struct {
	// TODO current Idx cannot use, because EncodeContainerDevices method not encode this filed.
	Idx        int
	UUID       string
	Type       string
	Usedmem    int32
	Usedcores  int32
	CustomInfo map[string]any
}

type ContainerDeviceRequest struct {
	Nums             int32
	Type             string
	Memreq           int32
	MemPercentagereq int32
	Coresreq         int32
}

type ContainerDevices []ContainerDevice
type ContainerDeviceRequests map[string]ContainerDeviceRequest

// type ContainerAllDevices map[string]ContainerDevices.
type PodSingleDevice []ContainerDevices
type PodDeviceRequests []ContainerDeviceRequests
type PodDevices map[string]PodSingleDevice

const (
	// OneContainerMultiDeviceSplitSymbol this is when one container use multi device, use : symbol to join device info.
	OneContainerMultiDeviceSplitSymbol = ":"

	// OnePodMultiContainerSplitSymbol this is when one pod having multi container and more than one container use device, use ; symbol to join device info.
	OnePodMultiContainerSplitSymbol = ";"
)

var (
	GPUSchedulerPolicy string
	InRequestDevices   map[string]string
	SupportDevices     map[string]string
	DevicesMap         map[string]Devices
	DevicesToHandle    []string
)

func init() {
	InRequestDevices = make(map[string]string)
	SupportDevices = make(map[string]string)
}

func (d *DeviceUsage) DeepCopy() *DeviceUsage { _ = "STUB: not implemented"; return nil }

func (m MigInUse) DeepCopy() MigInUse { _ = "STUB: not implemented"; return *new(MigInUse) }

func GetDevices() map[string]Devices { _ = "STUB: not implemented"; return nil }

func DecodeNodeDevices(str string) ([]*DeviceInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func DecodePairScores(pairScores string) (*DevicePairScores, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func EncodeNodeDevices(dlist []*DeviceInfo) string { _ = "STUB: not implemented"; return "" }

//tmp += val.ID + "," + strconv.FormatInt(int64(val.Count), 10) + "," + strconv.Itoa(int(val.Devmem)) + "," + strconv.Itoa(int(val.Devcore)) + "," + val.Type + "," + strconv.Itoa(val.Numa) + "," + strconv.FormatBool(val.Health) + "," + strconv.Itoa(val.Index) + OneContainerMultiDeviceSplitSymbol

// MarshalNodeDevices will only marshal general information, customInfo is neglected.
func MarshalNodeDevices(dlist []*DeviceInfo) string { _ = "STUB: not implemented"; return "" }

func UnMarshalNodeDevices(str string) ([]*DeviceInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func EncodeContainerDevices(cd ContainerDevices) string { _ = "STUB: not implemented"; return "" }

func EncodeContainerDeviceType(cd ContainerDevices, t string) string {
	_ = "STUB: not implemented"
	return ""
}

func EncodePodSingleDevice(pd PodSingleDevice) string { _ = "STUB: not implemented"; return "" }

func EncodePodDevices(checklist map[string]string, pd PodDevices) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func DecodeContainerDevices(str string) (ContainerDevices, error) {
	_ = "STUB: not implemented"
	return *new(ContainerDevices), nil
}

func DecodePodDevices(checklist map[string]string, annos map[string]string) (PodDevices, error) {
	_ = "STUB: not implemented"
	return *new(PodDevices), nil
}

// IMPORTANT: Do NOT skip empty ContainerDevices!
// We must preserve the index mapping between annotation entries and pod containers.
// The annotation format is: "dev1:;dev2:;;dev3:;" where ; separates containers
// If we skip empty entries, the index mapping will be broken for multi-container pods
// (especially pods with init containers where some containers don't use devices)

func PlatternMIG(n *MigInUse, templates []Geometry, templateIdx int) {
	_ = "STUB: not implemented"
	return
}

func GetDevicesUUIDList(infos []*DeviceInfo) []string { _ = "STUB: not implemented"; return nil }

func CheckHealth(devType string, resourceCountName string, node *corev1.Node) (bool, bool) {
	_ = "STUB: not implemented"
	return false, false
}

// Mirror the timestamp logic used by the Requesting branch: a
// Deleted_<ts> older than 60s on a node whose devices are otherwise
// reporting healthy means the previous cleanup is stale and the
// scheduler should bring the node back into its cache. Stamp
// Requesting_<now> and return (true, true) so the caller re-adds
// node devices on the next reconcile.
//
// Bare "Deleted" without a timestamp (used in some unit tests) and
// any unparsable timestamp must keep the conservative (true, false)
// path so we never recover from a malformed value.

// Mirror the empty-annotation branch: GetNode also acts as a guard
// for an empty node.Name and an uninitialised client (PatchNodeAnnotations
// panics in unit tests without it). Worth a follow-up to dedupe with
// the else branch into a helper.

// Enhanced ExtractMigTemplatesFromUUID with error handling.
func ExtractMigTemplatesFromUUID(uuid string) (int, int, error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

func Resourcereqs(pod *corev1.Pod) (counts PodDeviceRequests) {
	_ = "STUB: not implemented"
	// Total containers = init containers + regular containers
	return *new(PodDeviceRequests)
}

//Count Nvidia GPU

// Process init containers first (indices 0 to len(InitContainers)-1)

// Process regular containers (indices len(InitContainers) to totalContainers-1)

func CheckUUID(annos map[string]string, id, useKey, noUseKey, deviceType string) bool {
	_ = "STUB: not implemented"
	return false
}

// use , symbol to connect multiple uuid

// use , symbol to connect multiple uuid
