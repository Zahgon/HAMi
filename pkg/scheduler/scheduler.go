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

package scheduler

import (
	"context"
	"sync"
	"time"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes"
	coordinationv1 "k8s.io/client-go/listers/coordination/v1"
	listerscorev1 "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/record"
	extenderv1 "k8s.io/kube-scheduler/extender/v1"

	"github.com/Project-HAMi/HAMi/pkg/device"
	"github.com/Project-HAMi/HAMi/pkg/scheduler/policy"
	"github.com/Project-HAMi/HAMi/pkg/util/leaderelection"
)

const (
	defaultResync    = 1 * time.Hour
	syncedPollPeriod = 100 * time.Millisecond
)

type Scheduler struct {
	*nodeManager
	podManager    *device.PodManager
	quotaManager  *device.QuotaManager
	leaderManager leaderelection.LeaderManager

	stopCh       chan struct{}
	nodeNotify   chan struct{}
	leaderNotify chan struct{}

	kubeClient  kubernetes.Interface
	podLister   listerscorev1.PodLister
	nodeLister  listerscorev1.NodeLister
	quotaLister listerscorev1.ResourceQuotaLister
	leaseLister coordinationv1.LeaseLister
	//Node status returned by filter
	cachedstatus map[string]*NodeUsage
	//Node Overview
	overviewstatus map[string]*NodeUsage
	eventRecorder  record.EventRecorder
	started        uint32 // 0 = false, 1 = true

	lock   sync.RWMutex
	synced bool
}

func NewScheduler() *Scheduler { _ = "STUB: not implemented"; return nil }

// Use dummy leader manager when leaderElect is disabled
// This ensures IsLeader() always returns true and synced will not be set to false

func (s *Scheduler) GetQuotaManager() *device.QuotaManager { _ = "STUB: not implemented"; return nil }

func (s *Scheduler) GetPodManager() *device.PodManager { _ = "STUB: not implemented"; return nil }

func (s *Scheduler) GetLeaderManager() leaderelection.LeaderManager {
	_ = "STUB: not implemented"
	return *new(leaderelection.LeaderManager)
}

func (s *Scheduler) doNodeNotify() { _ = "STUB: not implemented"; return }

func (s *Scheduler) onAddPod(obj any) { _ = "STUB: not implemented"; return }

func (s *Scheduler) onUpdatePod(_, newObj any) { _ = "STUB: not implemented"; return }

func (s *Scheduler) onDelPod(obj any) { _ = "STUB: not implemented"; return }

// onDelNode handles node delete events. It removes any in-memory per-node
// lock bookkeeping to avoid unbounded growth when nodes are removed by
// autoscalers or administratively.
func (s *Scheduler) onDelNode(obj any) {
	_ = "STUB: not implemented"
	// Ensure downstream consumers are notified regardless of decoding success
	return
}

// cleanupNodeUsage removes the node from overviewstatus and cachedstatus maps
// to ensure metrics no longer report data for deleted nodes.
func (s *Scheduler) cleanupNodeUsage(nodeID string) { _ = "STUB: not implemented"; return }

func (s *Scheduler) onAddQuota(obj any) { _ = "STUB: not implemented"; return }

func (s *Scheduler) onUpdateQuota(oldObj, newObj any) { _ = "STUB: not implemented"; return }

func (s *Scheduler) onDelQuota(obj any) { _ = "STUB: not implemented"; return }

func (s *Scheduler) Start() error { _ = "STUB: not implemented"; return nil }

func (s *Scheduler) Stop() { _ = "STUB: not implemented"; return }

func (s *Scheduler) RegisterFromNodeAnnotations() { _ = "STUB: not implemented"; return }

func (s *Scheduler) register(labelSelector labels.Selector, printedLog map[string]bool) {
	_ = "STUB: not implemented"
	// Lock here to avoid setting s.synced to false, when we lost leadership, while doing register.
	// 1. lost leadership before register: synced will set to false in callbacks, and register will be skipped because IsLeader() returns false
	// 2. lost leadership during or after register: synced will set to true after finishing register, and callback will set it to false again after lock is acquired by callback
	return
}

// Only do registration when we are leader.

// Set synced to true only after getNodeUsage() succeeds

func (s *Scheduler) updateSchedulerLabel() { _ = "STUB: not implemented"; return }

// The pod is leader, apply the leader role label to it.

// The pod is not the leader, apply the follower role label to it.

func (s *Scheduler) WaitForCacheSync(ctx context.Context) bool {
	_ = "STUB: not implemented"
	return false
}

// InspectAllNodesUsage is used by metrics monitor.
func (s *Scheduler) InspectAllNodesUsage() *map[string]*NodeUsage {
	_ = "STUB: not implemented"
	return nil

	// returns all nodes and its device memory usage, and we filter it with nodeSelector, taints, nodeAffinity
	// unschedulerable and nodeName.
}

func (s *Scheduler) getNodesUsage(nodes *[]string, task *corev1.Pod) (*map[string]*NodeUsage, map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// The identified node does not have a gpu device, so the log here has no practical meaning,increase log priority.

func (s *Scheduler) getPodUsage() (map[string]device.PodUseDeviceStat, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Scheduler) Bind(args extenderv1.ExtenderBindingArgs) (*extenderv1.ExtenderBindingResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Scheduler) Filter(args extenderv1.ExtenderArgs) (*extenderv1.ExtenderFilterResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func genSuccessMsg(totalNodes int, target string, nodes []*policy.NodeScore) string {
	_ = "STUB: not implemented"
	return ""
}
