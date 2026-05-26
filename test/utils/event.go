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
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func GetEvents(clientSet *kubernetes.Clientset, namespace string, listOptions metav1.ListOptions) ([]v1.Event, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetPodEvents(clientSet *kubernetes.Clientset, namespace, podName string) ([]v1.Event, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
