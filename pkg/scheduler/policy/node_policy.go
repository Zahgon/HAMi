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

package policy

import (
	"github.com/Project-HAMi/HAMi/pkg/device"

	corev1 "k8s.io/api/core/v1"
)

type NodeScore struct {
	NodeID  string
	Node    *corev1.Node
	Devices device.PodDevices
	// Score recode every node all device user/allocate score
	Score float32
}

type NodeScoreList struct {
	NodeList []*NodeScore
	Policy   string
}

func (l NodeScoreList) Len() int { _ = "STUB: not implemented"; return 0 }

func (l NodeScoreList) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (l NodeScoreList) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

// default policy is Binpack

func (ns *NodeScore) OverrideScore(previous []*device.DeviceUsage, policy string) {
	_ = "STUB: not implemented"
	// current user having request resource
	return
}

func (ns *NodeScore) SnapshotDevice(devices DeviceUsageList) []*device.DeviceUsage {
	_ = "STUB: not implemented"
	return nil
}

func (ns *NodeScore) ComputeDefaultScore(devices DeviceUsageList) {
	_ = "STUB: not implemented"
	return
}
