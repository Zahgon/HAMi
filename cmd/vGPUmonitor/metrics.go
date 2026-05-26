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

package main

import (
	"github.com/Project-HAMi/HAMi/pkg/monitor/nvidia"

	"github.com/NVIDIA/go-nvml/pkg/nvml"
	"github.com/prometheus/client_golang/prometheus"
	corev1 "k8s.io/api/core/v1"
	listerscorev1 "k8s.io/client-go/listers/core/v1"
)

// ClusterManager is an example for a system that might have been built without
// Prometheus in mind. It models a central manager of jobs running in a
// cluster. Thus, we implement a custom Collector called
// ClusterManagerCollector, which collects information from a ClusterManager
// using its provided methods and turns them into Prometheus Metrics for
// collection.
//
// An additional challenge is that multiple instances of the ClusterManager are
// run within the same binary, each in charge of a different zone. We need to
// make use of wrapping Registerers to be able to register each
// ClusterManagerCollector instance with Prometheus.
type ClusterManager struct {
	Zone string
	// Contains many more fields not listed in this example.
	PodLister       listerscorev1.PodLister
	containerLister *nvidia.ContainerLister
	LegacyMetrics   bool
}

// ReallyExpensiveAssessmentOfTheSystemState is a mock for the data gathering a
// real cluster manager would have to do. Since it may actually be really
// expensive, it must only be called once per collection. This implementation,
// obviously, only returns some made-up data.
func (c *ClusterManager) ReallyExpensiveAssessmentOfTheSystemState() (
	oomCountByHost map[string]int, ramUsageByHost map[string]float64,
) {
	_ = "STUB: not implemented"
	// Just example fake data.
	return nil, nil
}

// ClusterManagerCollector implements the Collector interface.
type ClusterManagerCollector struct {
	ClusterManager *ClusterManager
}

// Descriptors used by the ClusterManagerCollector below.
// Metric and label names follow Prometheus naming best practices:
// https://prometheus.io/docs/practices/naming/
var (
	hostGPUdesc = prometheus.NewDesc(
		"hami_host_gpu_memory_used_bytes",
		"GPU device memory usage in bytes",
		[]string{"device_index", "device_uuid", "device_type"}, nil,
	)

	hostGPUUtilizationdesc = prometheus.NewDesc(
		"hami_host_gpu_utilization_ratio",
		"GPU core utilization ratio (0-100)",
		[]string{"device_index", "device_uuid", "device_type"}, nil,
	)

	ctrvGPUdesc = prometheus.NewDesc(
		"hami_vgpu_memory_used_bytes",
		"vGPU device memory usage in bytes",
		[]string{"namespace", "pod", "container", "vdevice_index", "device_uuid"}, nil,
	)

	ctrvGPUlimitdesc = prometheus.NewDesc(
		"hami_vgpu_memory_limit_bytes",
		"vGPU device memory limit in bytes",
		[]string{"namespace", "pod", "container", "vdevice_index", "device_uuid"}, nil,
	)
	ctrDeviceMemorydesc = prometheus.NewDesc(
		"hami_container_device_memory_bytes",
		`Container device memory usage breakdown in bytes (The label "context_size", "module_size", "buffer_size" and "offset" will be deprecated in v2.10.0, use hami_vgpu_memory_context_bytes, hami_vgpu_memory_module_bytes and hami_vgpu_memory_buffer_bytes instead)`,
		[]string{"namespace", "pod", "container", "vdevice_index", "device_uuid", "context_size", "module_size", "buffer_size", "offset"}, nil,
	)
	ctrDeviceUtilizationdesc = prometheus.NewDesc(
		"hami_container_device_utilization_ratio",
		"Container device SM utilization ratio",
		[]string{"namespace", "pod", "container", "vdevice_index", "device_uuid"}, nil,
	)
	ctrDeviceLastKernelDesc = prometheus.NewDesc(
		"hami_container_last_kernel_elapsed_seconds",
		"Seconds since last kernel execution in container",
		[]string{"namespace", "pod", "container", "vdevice_index", "device_uuid"}, nil,
	)
	ctrDeviceMigInfo = prometheus.NewDesc(
		"hami_mig_device_info",
		"MIG device information for container",
		[]string{"namespace", "pod", "container", "vdevice_index", "device_uuid", "instance_id"}, nil,
	)
	ctrDeviceMemoryContextDesc = prometheus.NewDesc(
		"hami_vgpu_memory_context_bytes",
		"Container device memory context size in bytes",
		[]string{"namespace", "pod", "container", "vdevice_index", "device_uuid"}, nil,
	)

	ctrDeviceMemoryModuleDesc = prometheus.NewDesc(
		"hami_vgpu_memory_module_bytes",
		"Container device memory module size in bytes",
		[]string{"namespace", "pod", "container", "vdevice_index", "device_uuid"}, nil,
	)

	ctrDeviceMemoryBufferDesc = prometheus.NewDesc(
		"hami_vgpu_memory_buffer_bytes",
		"Container device memory buffer size in bytes",
		[]string{"namespace", "pod", "container", "vdevice_index", "device_uuid"}, nil,
	)
)

// Legacy metric descriptors (populated only when --legacy-metrics is enabled).
var (
	legacyHostGPUdesc              *prometheus.Desc
	legacyHostGPUUtilizationdesc   *prometheus.Desc
	legacyCtrvGPUdesc              *prometheus.Desc
	legacyCtrvGPUlimitdesc         *prometheus.Desc
	legacyCtrDeviceMemorydesc      *prometheus.Desc
	legacyCtrDeviceUtilizationdesc *prometheus.Desc
	legacyCtrDeviceLastKernelDesc  *prometheus.Desc
	legacyCtrDeviceMigInfo         *prometheus.Desc
)

func initLegacyDescriptors() { _ = "STUB: not implemented"; return }

func sendLegacyMetric(ch chan<- prometheus.Metric, desc *prometheus.Desc, valueType prometheus.ValueType, value float64, labels ...string) {
	_ = "STUB: not implemented"
	return
}

// Describe is implemented with DescribeByCollect. That's possible because the
// Collect method will always return the same two metrics with the same two
// descriptors.
func (cc ClusterManagerCollector) Describe(ch chan<- *prometheus.Desc) {
	_ = "STUB: not implemented"
	return
}

//func parseidstr(podusage string) (string, string, error) {
//	tmp := strings.Split(podusage, "_")
//	if len(tmp) > 1 {
//		return tmp[0], tmp[1], nil
//	} else {
//		return "", "", errors.New("parse error")
//	}
//}
//
//func gettotalusage(usage podusage, vidx int) (deviceMemory, error) {
//	added := deviceMemory{
//		bufferSize:  0,
//		contextSize: 0,
//		moduleSize:  0,
//		offset:      0,
//		total:       0,
//	}
//	for _, val := range usage.sr.procs {
//		added.bufferSize += val.used[vidx].bufferSize
//		added.contextSize += val.used[vidx].contextSize
//		added.moduleSize += val.used[vidx].moduleSize
//		added.offset += val.used[vidx].offset
//		added.total += val.used[vidx].total
//	}
//	return added, nil
//}
//
//func getTotalUtilization(usage podusage, vidx int) deviceUtilization {
//	added := deviceUtilization{
//		decUtil: 0,
//		encUtil: 0,
//		smUtil:  0,
//	}
//	for _, val := range usage.sr.procs {
//		added.decUtil += val.deviceUtil[vidx].decUtil
//		added.encUtil += val.deviceUtil[vidx].encUtil
//		added.smUtil += val.deviceUtil[vidx].smUtil
//	}
//	return added
//}

// Collect first triggers the ReallyExpensiveAssessmentOfTheSystemState. Then it
// creates constant metrics for each host on the fly based on the returned data.
//
// Note that Collect could be called concurrently, so we depend on
// ReallyExpensiveAssessmentOfTheSystemState to be concurrency-safe.
func (cc ClusterManagerCollector) Collect(ch chan<- prometheus.Metric) {
	_ = "STUB: not implemented"
	return
}

// Collect GPU information

// Decide whether to continue or return based on business requirements

// Collect Pod and Container information

// Decide whether to continue or return based on business requirements

// Collect Pod and Container Mig information

// Decide whether to continue or return based on business requirements

func (cc ClusterManagerCollector) collectGPUInfo(ch chan<- prometheus.Metric) error {
	_ = "STUB: not implemented"
	return nil
}

func (cc ClusterManagerCollector) initNVML() error { _ = "STUB: not implemented"; return nil }

func (cc ClusterManagerCollector) getDeviceCount() (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (cc ClusterManagerCollector) collectGPUDeviceMetrics(ch chan<- prometheus.Metric, index int) error {
	_ = "STUB: not implemented"
	return nil
}

func (cc ClusterManagerCollector) collectGPUMemoryMetrics(ch chan<- prometheus.Metric, hdev nvml.Device, index int) error {
	_ = "STUB: not implemented"
	return nil
}

func (cc ClusterManagerCollector) collectGPUUtilizationMetrics(ch chan<- prometheus.Metric, hdev nvml.Device, index int) error {
	_ = "STUB: not implemented"
	return nil
}

func (cc ClusterManagerCollector) collectPodAndContainerInfo(ch chan<- prometheus.Metric) error {
	_ = "STUB: not implemented"
	return nil
}

// podUID -> containers

// Iterate through each Pod

// Iterate through each container in the Pod

// Find the matching container

// Exit the inner loop after finding the matching container

func (cc ClusterManagerCollector) isPodUIDMatched(pod *corev1.Pod, podUID string) bool {
	_ = "STUB: not implemented"
	return false
}

func (cc ClusterManagerCollector) collectContainerMetrics(ch chan<- prometheus.Metric, pod *corev1.Pod, ctr corev1.Container, c *nvidia.ContainerUsage, nowSec int64) error {
	_ = "STUB: not implemented"
	// Validate inputs
	return nil
}

// Iterate through each device

// Ensure UUID is truncated to 40 characters

// Collect device metrics

func (cc ClusterManagerCollector) collectPodAndContainerMigInfo(ch chan<- prometheus.Metric) error {
	_ = "STUB: not implemented"
	return nil
}

func sendMetric(ch chan<- prometheus.Metric, desc *prometheus.Desc, valueType prometheus.ValueType, value float64, labels ...string) error {
	_ = "STUB: not implemented"
	return nil
}

// NewClusterManager first creates a Prometheus-ignorant ClusterManager
// instance. Then, it creates a ClusterManagerCollector for the just created
// ClusterManager. Finally, it registers the ClusterManagerCollector with a
// wrapping Registerer that adds the zone as a label. In this way, the metrics
// collected by different ClusterManagerCollectors do not collide.
func NewClusterManager(zone string, reg prometheus.Registerer, containerLister *nvidia.ContainerLister, legacyMetrics bool) *ClusterManager {
	_ = "STUB: not implemented"
	return nil
}
