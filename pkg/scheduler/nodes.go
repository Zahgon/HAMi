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
	"sync"

	corev1 "k8s.io/api/core/v1"

	"github.com/Project-HAMi/HAMi/pkg/device"
	"github.com/Project-HAMi/HAMi/pkg/scheduler/policy"
)

type NodeUsage struct {
	Node    *corev1.Node
	Devices policy.DeviceUsageList
}

func (n *NodeUsage) DeepCopy() *NodeUsage { _ = "STUB: not implemented"; return nil }

type nodeManager struct {
	nodes map[string]*device.NodeInfo
	mutex sync.RWMutex
}

func newNodeManager() *nodeManager { _ = "STUB: not implemented"; return nil }

func (m *nodeManager) addNode(nodeID string, nodeInfo *device.NodeInfo) {
	_ = "STUB: not implemented"
	return
}

func (m *nodeManager) rmNodeDevices(nodeID string, deviceVendor string) {
	_ = "STUB: not implemented"
	return
}

func (m *nodeManager) rmNode(nodeID string) { _ = "STUB: not implemented"; return }

func (m *nodeManager) GetNode(nodeID string) (*device.NodeInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *nodeManager) ListNodes() (map[string]*device.NodeInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
