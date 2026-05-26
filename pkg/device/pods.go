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
	"sync"

	corev1 "k8s.io/api/core/v1"
	k8stypes "k8s.io/apimachinery/pkg/types"
)

type PodInfo struct {
	*corev1.Pod
	NodeID  string
	Devices PodDevices
	CtrIDs  []string
}

// PodUseDeviceStat counts pod use device info.
type PodUseDeviceStat struct {
	TotalPod     int // Count of all running pods on the current node
	UseDevicePod int // Count of running pods that use devices
}

type PodManager struct {
	pods  map[k8stypes.UID]*PodInfo
	mutex sync.RWMutex
}

func NewPodManager() *PodManager { _ = "STUB: not implemented"; return nil }

func (m *PodManager) AddPod(pod *corev1.Pod, nodeID string, devices PodDevices) bool {
	_ = "STUB: not implemented"
	return false
}

func (m *PodManager) UpdatePod(pod *corev1.Pod) { _ = "STUB: not implemented"; return }

func (m *PodManager) DelPod(pod *corev1.Pod) { _ = "STUB: not implemented"; return }

func (m *PodManager) GetPod(pod *corev1.Pod) (*PodInfo, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (m *PodManager) TakeAndDeletePod(pod *corev1.Pod) (*PodInfo, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (m *PodManager) ListPodsUID() ([]*corev1.Pod, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *PodManager) ListPodsInfo() []*PodInfo { _ = "STUB: not implemented"; return nil }

func (p *PodInfo) DeepCopy() *PodInfo { _ = "STUB: not implemented"; return nil }

func (pd PodDevices) DeepCopy() PodDevices { _ = "STUB: not implemented"; return *new(PodDevices) }

func (psd PodSingleDevice) DeepCopy() PodSingleDevice {
	_ = "STUB: not implemented"
	return *new(PodSingleDevice)
}

func (cd ContainerDevices) DeepCopy() ContainerDevices {
	_ = "STUB: not implemented"
	return *new(ContainerDevices)
}

func (c ContainerDevice) DeepCopy() ContainerDevice {
	_ = "STUB: not implemented"
	return *new(ContainerDevice)
}

func (m *PodManager) GetScheduledPods() (map[k8stypes.UID]*PodInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Return a shallow copy of the pods map to avoid race conditions.
// This prevents a "concurrent map iteration and map write" fatal error.
