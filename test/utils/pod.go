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

package utils

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

var Pod = &corev1.Pod{
	ObjectMeta: metav1.ObjectMeta{
		Name:      "gpu-pod",
		Namespace: "default",
	},
	Spec: corev1.PodSpec{
		Containers: []corev1.Container{
			{
				Name:    "cuda-container",
				Image:   "nvcr.io/nvidia/k8s/cuda-sample:vectoradd-cuda12.5.0",
				Command: []string{"/bin/sh"},
				Args:    []string{"-c", "sleep 86400"},
				Resources: corev1.ResourceRequirements{
					Limits: corev1.ResourceList{
						"nvidia.com/gpu":      resource.MustParse("1"),
						"nvidia.com/gpumem":   resource.MustParse(GPUPodMemory),
						"nvidia.com/gpucores": resource.MustParse(GPUPodCore),
					},
				},
			},
		},
	},
}

func GetPods(clientSet *kubernetes.Clientset, namespace string) (*corev1.PodList, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func CreatePod(clientSet *kubernetes.Clientset, pod *corev1.Pod, namespace string) (*corev1.Pod, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func DeletePod(clientSet *kubernetes.Clientset, namespace, podName string) error {
	_ = "STUB: not implemented"
	return nil
}

func WaitForPodRunning(clientSet kubernetes.Interface, namespace, podName string) error {
	_ = "STUB: not implemented"
	return nil
}

// Interval for checking Pod status
// Increased timeout for GPU Pods

// Fetch the Pod object from the Kubernetes API

// Print Pod status for debugging

// Check if the Pod is in the Running state

// Check if the Pod is in a Failed or Unknown state

// Print Pod events for debugging

// If the Pod is not in Running, Failed, or Unknown state, continue waiting
