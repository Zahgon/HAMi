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
	"flag"

	"k8s.io/client-go/kubernetes"
)

var kubeConfig string

func init() {
	flag.StringVar(&kubeConfig, "kubeconfig", "", "Path to the kubeConfig file")
}

// resolveKubeConfigPath picks kubeconfig in order: --kubeconfig flag, KUBE_CONF, ~/.kube/config.
func resolveKubeConfigPath() string { _ = "STUB: not implemented"; return "" }

func validateKubeConfigPath(configPath string) string { _ = "STUB: not implemented"; return "" }

func DefaultKubeConfigPath() string { _ = "STUB: not implemented"; return "" }

func GetClientSet() *kubernetes.Clientset { _ = "STUB: not implemented"; return nil }

func GetRandom() string { _ = "STUB: not implemented"; return "" }

// KubectlExecInPod executes a shell command in a specified Pod using kubectl exec.
func KubectlExecInPod(namespace, podName, command string) ([]byte, error) {
	_ = "STUB: not implemented"
	// Wait for the container to stabilize
	return nil, nil
}

// Build the kubectl exec command

// Capture the command output (both stdout and stderr)
