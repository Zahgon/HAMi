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

package util

import (
	"context"
	"flag"

	corev1 "k8s.io/api/core/v1"
)

var (
	HandshakeAnnos map[string]string
)

func init() {
	HandshakeAnnos = make(map[string]string)
}

func GetNode(nodename string) (*corev1.Node, error) { _ = "STUB: not implemented"; return nil, nil }

func GetPendingPod(ctx context.Context, node string) (*corev1.Pod, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// filter pods for this node.

// Allow both "allocating" and "success" phases for multi-container pods
// where some containers have already been allocated but others are still pending

func GetAllocatePodByNode(ctx context.Context, nodeName string) (*corev1.Pod, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func PatchNodeAnnotations(node *corev1.Node, annotations map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

func PatchPodAnnotations(pod *corev1.Pod, annotations map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

func PatchPodLabels(namespace, name string, labels map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

func InitKlogFlags() *flag.FlagSet {
	_ = "STUB: not implemented"
	// Init log flags
	return nil
}

func MarkAnnotationsToDelete(devType string, nn string) error {
	_ = "STUB: not implemented"
	return nil
}

func GetGPUSchedulerPolicyByPod(defaultPolicy string, task *corev1.Pod) string {
	_ = "STUB: not implemented"
	return ""
}

func IsPodInTerminatedState(pod *corev1.Pod) bool { _ = "STUB: not implemented"; return false }

func IsPodTerminating(pod *corev1.Pod) bool { _ = "STUB: not implemented"; return false }

func AllContainersCreated(pod *corev1.Pod) bool { _ = "STUB: not implemented"; return false }
