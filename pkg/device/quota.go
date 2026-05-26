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

package device

import (
	"sync"

	corev1 "k8s.io/api/core/v1"
)

type Quota struct {
	Used  int64
	Limit int64
}

type DeviceQuota map[string]*Quota

type QuotaManager struct {
	Quotas map[string]*DeviceQuota
	mutex  sync.RWMutex
}

var localCache QuotaManager

func GetLocalCache() *QuotaManager { _ = "STUB: not implemented"; return nil }

var once sync.Once

func NewQuotaManager() *QuotaManager { _ = "STUB: not implemented"; return nil }

func (q *QuotaManager) FitQuota(ns string, memreq int64, memoryFactor int32, coresreq int64, deviceName string) bool {
	_ = "STUB: not implemented"
	return false
}

func countPodDevices(podDev PodDevices) map[string]int64 { _ = "STUB: not implemented"; return nil }

func (q *QuotaManager) AddUsage(pod *corev1.Pod, podDev PodDevices) {
	_ = "STUB: not implemented"
	return
}

func (q *QuotaManager) RmUsage(pod *corev1.Pod, podDev PodDevices) {
	_ = "STUB: not implemented"
	return
}

func IsManagedQuota(quotaName string) bool { _ = "STUB: not implemented"; return false }

func (q *QuotaManager) AddQuota(quota *corev1.ResourceQuota) { _ = "STUB: not implemented"; return }

func (q *QuotaManager) DelQuota(quota *corev1.ResourceQuota) { _ = "STUB: not implemented"; return }

func (q *QuotaManager) GetResourceQuota() map[string]*DeviceQuota {
	_ = "STUB: not implemented"
	return nil
}
