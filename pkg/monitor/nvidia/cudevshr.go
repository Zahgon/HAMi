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

package nvidia

import (
	"os"
	"sync"
	"time"

	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	corelisters "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/klog/v2"
)

const SharedRegionMagicFlag = 19920718

type headerT struct {
	initializedFlag int32
	majorVersion    int32
	minorVersion    int32
}

type UsageInfo interface {
	DeviceMax() int
	DeviceNum() int
	DeviceMemoryContextSize(idx int) uint64
	DeviceMemoryModuleSize(idx int) uint64
	DeviceMemoryBufferSize(idx int) uint64
	DeviceMemoryOffset(idx int) uint64
	DeviceMemoryTotal(idx int) uint64
	DeviceSmUtil(idx int) uint64
	SetDeviceSmLimit(l uint64)
	IsValidUUID(idx int) bool
	DeviceUUID(idx int) string
	DeviceMemoryLimit(idx int) uint64
	SetDeviceMemoryLimit(l uint64)
	LastKernelTime() int64
	//UsedMemory(idx int) (uint64, error)
	GetPriority() int
	GetRecentKernel() int32
	SetRecentKernel(v int32)
	GetUtilizationSwitch() int32
	SetUtilizationSwitch(v int32)
}

type ContainerUsage struct {
	PodUID        string
	ContainerName string
	data          []byte
	Info          UsageInfo
}

type ContainerLister struct {
	containerPath string
	containers    map[string]*ContainerUsage
	mutex         sync.Mutex
	clientset     *kubernetes.Clientset
	nodeName      string

	// Fields for the informer-based pod cache mechanism
	informerFactory informers.SharedInformerFactory
	podInformer     cache.SharedIndexInformer
	podLister       corelisters.PodLister
	podListerSynced cache.InformerSynced
	stopCh          chan struct{}
}

var resyncInterval time.Duration = 5 * time.Minute

func init() {
	if os.Getenv("HAMI_RESYNC_INTERVAL") != "" {
		// If HAMI_RESYNC_INTERVAL is set, parse it
		if interval, err := time.ParseDuration(os.Getenv("HAMI_RESYNC_INTERVAL")); err == nil {
			resyncInterval = interval
		} else {
			klog.Warningf("Invalid HAMI_RESYNC_INTERVAL value: %s, using default %v", os.Getenv("HAMI_RESYNC_INTERVAL"), resyncInterval)
		}
	}
}

func NewContainerLister() (*ContainerLister, error) { _ = "STUB: not implemented"; return nil, nil }

// Initialize the informer

func (l *ContainerLister) Lock() { _ = "STUB: not implemented"; return }

func (l *ContainerLister) UnLock() { _ = "STUB: not implemented"; return }

func (l *ContainerLister) ListContainers() map[string]*ContainerUsage {
	_ = "STUB: not implemented"
	return nil
}

func (l *ContainerLister) Clientset() *kubernetes.Clientset { _ = "STUB: not implemented"; return nil }

func (l *ContainerLister) Update() error { _ = "STUB: not implemented"; return nil }

// no cuInit in container

func loadCache(fpath string) (*ContainerUsage, error) { _ = "STUB: not implemented"; return nil, nil }

func (l *ContainerLister) initInformerWithConfig(resyncInterval time.Duration) error {
	_ = "STUB: not implemented"
	// Create informer factory with a longer resync period to reduce API calls
	return nil
}

// Start the informer

// Wait for cache sync

func (l *ContainerLister) onPodDelete(obj any) { _ = "STUB: not implemented"; return }
