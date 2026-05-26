/*
 * Copyright (c) 2024, HAMi.  All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package plugin

import (
	"context"
	"errors"
	"strings"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog/v2"
	"tags.cncf.io/container-device-interface/specs-go"

	"github.com/Project-HAMi/HAMi/pkg/device"
	"github.com/Project-HAMi/HAMi/pkg/device/nvidia"
	"github.com/Project-HAMi/HAMi/pkg/util"
	"github.com/Project-HAMi/HAMi/pkg/util/client"
)

// GetLibPath returns the path to the vGPU library.
func GetLibPath() string { _ = "STUB: not implemented"; return "" }

func GetNextDeviceRequest(dtype string, p corev1.Pod) (corev1.Container, device.ContainerDevices, error) {
	_ = "STUB: not implemented"
	return *new(corev1.Container), *new(device.ContainerDevices), nil
}

// The annotation format follows the order: init containers first, then regular containers
// Index mapping:
//   0 to len(InitContainers)-1: init containers
//   len(InitContainers) to len(InitContainers)+len(Containers)-1: regular containers

// This is an init container

// This is a regular container

var eraseNextDeviceTypeFromAnnotation = func(dtype string, p corev1.Pod) error {
	pdevices, err := device.DecodePodDevices(device.InRequestDevices, p.Annotations)
	if err != nil {
		return err
	}
	res := device.PodSingleDevice{}
	pd, ok := pdevices[dtype]
	if !ok {
		return errors.New("erase device annotation not found")
	}
	found := false
	for _, val := range pd {
		if found {
			res = append(res, val)
		} else {
			if len(val) > 0 {
				found = true
				res = append(res, device.ContainerDevices{})
			} else {
				res = append(res, val)
			}
		}
	}
	klog.Infoln("After erase res=", res)
	newannos := make(map[string]string)
	newannos[device.InRequestDevices[dtype]] = device.EncodePodSingleDevice(res)
	return util.PatchPodAnnotations(&p, newannos)
}

func GetIndexAndTypeFromUUID(uuid string) (string, int) { _ = "STUB: not implemented"; return "", 0 }

func GetMigUUIDFromSmiOutput(output string, uuid string, idx int) string {
	_ = "STUB: not implemented"
	return ""
}

func GetMigUUIDFromIndex(uuid string, idx int) string { _ = "STUB: not implemented"; return "" }

func GetMigGpuInstanceIdFromIndex(uuid string, idx int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func GetDeviceNums() (int, error) { _ = "STUB: not implemented"; return 0, nil }

func GetDeviceNames() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func (nv *NvidiaDevicePlugin) DisableOtherNVMLOperation() {
	_ = "STUB: not implemented"
	// Create MIG apply lock file
	return
}

// If the lock file creation fails, it is highly likely that the mig apply will be failed, so the plugin should terminate.

//wait for disableHealthChecks to be closed,signal must be true or wait forever

func (nv *NvidiaDevicePlugin) EnableOtherNVMLOperation() {
	_ = "STUB: not implemented"
	// Remove MIG apply lock file
	return
}

func (nv *NvidiaDevicePlugin) ApplyMigTemplate() { _ = "STUB: not implemented"; return }

func (nv *NvidiaDevicePlugin) GenerateMigTemplate(devtype string, devindex int, val device.ContainerDevice) (int, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

// Initialize to an invalid position

// Helper function to check if a model is in the list of models.
func containsModel(target string, models []string) bool { _ = "STUB: not implemented"; return false }

// Helper function to check if a device index is in the list of devices.
func containsDevice(target int, devices []int32) bool { _ = "STUB: not implemented"; return false }

// Helper function to deepcopy new mig spec
func deepCopyMigConfig(src nvidia.MigConfigSpec) nvidia.MigConfigSpec {
	_ = "STUB: not implemented"
	return *new(nvidia.MigConfigSpec)
}

func (nv *NvidiaDevicePlugin) GetContainerDeviceStrArray(c device.ContainerDevices) []string {
	_ = "STUB: not implemented"
	return nil
}

var podAllocationTrySuccess = func(nodeName string, devName string, lockName string, pod *corev1.Pod) {
	refreshed, err := client.GetClient().CoreV1().Pods(pod.Namespace).Get(context.Background(), pod.Name, metav1.GetOptions{})
	if err != nil {
		klog.Errorf("Error getting pod %s/%s: %v", pod.Namespace, pod.Name, err)
		return
	}
	annos := refreshed.Annotations[device.InRequestDevices[devName]]
	klog.Infof("Trying allocation success: %s", annos)
	for _, val := range device.DevicesToHandle {
		if strings.Contains(annos, val) {
			return
		}
	}
	klog.Infof("All devices allocate success, releasing lock")
	PodAllocationSuccess(nodeName, pod, lockName)
}

func PodAllocationSuccess(nodeName string, pod *corev1.Pod, lockName string) {
	_ = "STUB: not implemented"
	return
}

func updatePodAnnotationsAndReleaseLock(nodeName string, pod *corev1.Pod, lockName string, deviceBindPhase string) {
	_ = "STUB: not implemented"
	return
}

var podAllocationFailed = func(nodeName string, pod *corev1.Pod, lockName string) {
	klog.Infof("Pod allocation failed for pod %s/%s on node %s", pod.Namespace, pod.Name, nodeName)
	updatePodAnnotationsAndReleaseLock(nodeName, pod, lockName, util.DeviceBindFailed)
}

func checkCDISpecFile(filePath, kind string) error { _ = "STUB: not implemented"; return nil }

func checkCDISpec(spec specs.Spec, kind string) error { _ = "STUB: not implemented"; return nil }

func createSpecFile(outputPath string) error { _ = "STUB: not implemented"; return nil }
