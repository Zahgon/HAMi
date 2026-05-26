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

package awsneuron

import (
	"flag"

	"github.com/Project-HAMi/HAMi/pkg/device"

	corev1 "k8s.io/api/core/v1"
)

type AWSNeuronDevices struct {
	resourceCountName string
	resourceCoreName  string
	coresPerAWSNeuron uint
	coremask          uint
}

const (
	AWSNeuronDevice          = "AWSNeuron"
	AWSNeuronCommonWord      = "AWSNeuron"
	AWSNeuronDeviceSelection = "aws.amazon.com/neuron-index"
	AWSNeuronUseUUID         = "aws.amazon.com/use-neuron-uuid"
	AWSNeuronNoUseUUID       = "aws.amazon.com/nouse-neuron-uuid"
	AWSNeuronAssignedIndex   = "AWS_NEURON_IDS"
	AWSNeuronAssignedNode    = "aws.amazon.com/predicate-node"
	AWSNeuronPredicateTime   = "NEURON_ALLOC_TIME"
	AWSNeuronResourceType    = "NEURON_RESOURCE_TYPE"
	AWSNeuronAllocated       = "NEURON_ALLOCATED"
	AWSUsageInfo             = "awsusageinfo"
	AWSNodeType              = "AWSNodeType"
)

type AWSNeuronConfig struct {
	ResourceCountName string `yaml:"resourceCountName"`
	ResourceCoreName  string `yaml:"resourceCoreName"`
}

func InitAWSNeuronDevice(config AWSNeuronConfig) *AWSNeuronDevices {
	_ = "STUB: not implemented"
	return nil
}

func (dev *AWSNeuronDevices) CommonWord() string { _ = "STUB: not implemented"; return "" }

func ParseConfig(fs *flag.FlagSet) { _ = "STUB: not implemented"; return }

func (dev *AWSNeuronDevices) MutateAdmission(ctr *corev1.Container, p *corev1.Pod) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (dev *AWSNeuronDevices) GetNodeDevices(n corev1.Node) ([]*device.DeviceInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (dev *AWSNeuronDevices) PatchAnnotations(pod *corev1.Pod, annoinput *map[string]string, pd device.PodDevices) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

// This needs to modify, it has to be core indexes?

func (dev *AWSNeuronDevices) LockNode(n *corev1.Node, p *corev1.Pod) error {
	_ = "STUB: not implemented"
	return nil
}

func (dev *AWSNeuronDevices) ReleaseNodeLock(n *corev1.Node, p *corev1.Pod) error {
	_ = "STUB: not implemented"
	return nil
}

func (dev *AWSNeuronDevices) NodeCleanUp(nn string) error { _ = "STUB: not implemented"; return nil }

func (dev *AWSNeuronDevices) checkType(n device.ContainerDeviceRequest) (bool, bool, bool) {
	_ = "STUB: not implemented"
	return false, false, false
}

func (dev *AWSNeuronDevices) CheckHealth(devType string, n *corev1.Node) (bool, bool) {
	_ = "STUB: not implemented"
	return false, false
}

func (dev *AWSNeuronDevices) GetResourceNames() device.ResourceNames {
	_ = "STUB: not implemented"
	return *new(device.ResourceNames)
}

func (dev *AWSNeuronDevices) GenerateResourceRequests(ctr *corev1.Container) device.ContainerDeviceRequest {
	_ = "STUB: not implemented"
	return *new(device.ContainerDeviceRequest)
}

func (dev *AWSNeuronDevices) ScoreNode(node *corev1.Node, podDevices device.PodSingleDevice, previous []*device.DeviceUsage, policy string) float32 {
	_ = "STUB: not implemented"
	return 0
}

func (dev *AWSNeuronDevices) AddResourceUsage(pod *corev1.Pod, n *device.DeviceUsage, ctr *device.ContainerDevice) error {
	_ = "STUB: not implemented"
	return nil
}

func countMaskAvailable(mask int32) int32 { _ = "STUB: not implemented"; return 0 }

func addCoreUsage(prev map[string]any, require int) map[string]any {
	_ = "STUB: not implemented"
	return nil
}

func continuousDeviceAvailable(devices []*device.DeviceUsage, start int, count int) []int {
	_ = "STUB: not implemented"
	return nil
}

func graphSelect(devices []*device.DeviceUsage, count int) []int {
	_ = "STUB: not implemented"
	return nil
}

//Deal with ring

func (neuron *AWSNeuronDevices) Fit(devices []*device.DeviceUsage, request device.ContainerDeviceRequest, pod *corev1.Pod, nodeinfo *device.NodeInfo, allocated *device.PodDevices) (bool, map[string]device.ContainerDevices, string) {
	_ = "STUB: not implemented"
	return false, nil, ""
}
