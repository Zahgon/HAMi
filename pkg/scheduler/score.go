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
package scheduler

import (
	corev1 "k8s.io/api/core/v1"

	"github.com/Project-HAMi/HAMi/pkg/device"
	"github.com/Project-HAMi/HAMi/pkg/scheduler/policy"
)

func viewStatus(usage NodeUsage) { _ = "STUB: not implemented"; return }

func getNodeResources(list NodeUsage, t string) []*device.DeviceUsage {
	_ = "STUB: not implemented"
	return nil
}

func fitInDevices(node *NodeUsage, requests device.ContainerDeviceRequests, pod *corev1.Pod, nodeInfo *device.NodeInfo, devinput *device.PodDevices) (bool, string) {
	_ = "STUB: not implemented"
	//devmap := make(map[string]device.ContainerDevices)
	return false, ""
}

// computer all device score for one node

//This loop is for requests for different devices

//bc node.Devices has been sorted, so we should find out the correct device

func (s *Scheduler) calcScore(nodes *map[string]*NodeUsage, resourceReqs device.PodDeviceRequests, task *corev1.Pod, failedNodes map[string]string) (*policy.NodeScoreList, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Assume the node is a fit by default. This handles pods with no device
// requests, which should be schedulable on any node.

//This loop is for different container request

// container need no device and we have got certain deviceType

// found certain deviceType, fill missing empty allocation for containers before this

// only pod scheduler failure will record failure event
