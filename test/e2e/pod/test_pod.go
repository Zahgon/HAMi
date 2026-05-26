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

package e2e

import (
	"fmt"
	"time"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"

	"github.com/Project-HAMi/HAMi/test/utils"
)

var _ = ginkgo.Describe("Pod E2E Tests", ginkgo.Ordered, func() {
	const (
		Namespace      = utils.GPUNameSpace
		NodeLabelKey   = utils.GPUNodeLabelKey
		NodeLabelValue = utils.GPUNodeLabelValue
		DeleteTimeout  = 300 * time.Second
		DeleteInterval = 10 * time.Second
	)

	var (
		clientSet *kubernetes.Clientset
		newPod    *corev1.Pod
		nodeName  string
	)

	ginkgo.BeforeAll(func() {
		clientSet = utils.GetClientSet()

		var err error
		nodeName, err = utils.GetGPUNode(clientSet)
		gomega.Expect(err).NotTo(gomega.HaveOccurred())
		fmt.Printf("Using GPU node: %s\n", nodeName)

		ginkgo.By("Adding node labeling")
		_, err = utils.AddNodeLabel(clientSet, nodeName, NodeLabelKey, NodeLabelValue)
		gomega.Expect(err).NotTo(gomega.HaveOccurred())
	})

	ginkgo.AfterEach(func() {
		ginkgo.By("Cleanup pod after each test")
		cleanupPod(newPod, clientSet)
	})

	ginkgo.AfterAll(func() {
		ginkgo.By("Deleting node labeling")
		_, err := utils.RemoveNodeLabel(clientSet, nodeName, NodeLabelKey)
		gomega.Expect(err).NotTo(gomega.HaveOccurred())
	})

	ginkgo.It("creates a single pod with CUDA configuration", func() {
		newPod = utils.Pod.DeepCopy()
		newPod.Name += utils.GetRandom()

		// Ensure cleanup even if the test fails
		ginkgo.DeferCleanup(func() {
			ginkgo.By("DeferCleanup: Deleting pod after test")
			cleanupPod(newPod, clientSet)
		})

		createAndVerifyPod(newPod, clientSet)

		ginkgo.By("Verifying GPU memory in pod by executing: " + utils.GPUExecuteNvidiaSMI)
		output, err := utils.KubectlExecInPod(newPod.Namespace, newPod.Name, utils.GPUExecuteNvidiaSMI)
		if err != nil {
			fmt.Printf("nvidia-smi execution error: %v\n", err)
			fmt.Printf("nvidia-smi output: %s\n", output)
		}
		gomega.Expect(err).NotTo(gomega.HaveOccurred(), "nvidia-smi command failed")

		fmt.Println("nvidia-smi output:")
		fmt.Println(string(output))

		ginkgo.By("Verifying CUDA execution in pod by executing: " + utils.GPUExecuteCudaSample)
		output, err = utils.KubectlExecInPod(newPod.Namespace, newPod.Name, utils.GPUExecuteCudaSample)
		if err != nil {
			fmt.Printf("CUDA sample execution error: %v\n", err)
			fmt.Printf("CUDA sample output: %s\n", output)
		}
		gomega.Expect(err).NotTo(gomega.HaveOccurred(), "CUDA sample command failed")

		fmt.Println("CUDA sample output:")
		fmt.Println(string(output))
	})

	ginkgo.It("create overcommit pods", func() {
		newPod = prepareOvercommitPod(utils.Pod.DeepCopy(), Namespace) // Pass the namespace to the helper

		ginkgo.By("Creating overcommit pod in namespace " + Namespace)
		createdPod, err := utils.CreatePod(clientSet, newPod, Namespace)
		gomega.Expect(err).NotTo(gomega.HaveOccurred())
		gomega.Expect(createdPod.Name).To(gomega.Equal(newPod.Name), "Pod was not created successfully")

		ginkgo.By("Verifying pod is pending due to filtering")
		gomega.Eventually(func() bool {
			return checkPodPendingDueToFiltering(clientSet, newPod)
		}, DeleteTimeout, DeleteInterval).Should(gomega.BeTrue())
	})
})

func cleanupPod(pod *corev1.Pod, clientSet *kubernetes.Clientset) {
	_ = "STUB: not implemented"
	return
}

func podExists(namespace, podName string, clientSet *kubernetes.Clientset) bool {
	_ = "STUB: not implemented"
	return false
}

func createAndVerifyPod(pod *corev1.Pod, clientSet *kubernetes.Clientset) {
	_ = "STUB: not implemented"
	return
}

func prepareOvercommitPod(pod *corev1.Pod, namespace string) *corev1.Pod {
	_ = "STUB: not implemented"
	return nil
}

// Ensure the pod's namespace is correctly set

// Modify pod spec for overcommit scenario

func checkPodPendingDueToFiltering(clientSet *kubernetes.Clientset, pod *corev1.Pod) bool {
	_ = "STUB: not implemented"
	return false
}
