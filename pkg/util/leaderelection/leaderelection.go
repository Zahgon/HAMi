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

package leaderelection

import (
	"sync"
	"time"

	coordinationv1 "k8s.io/api/coordination/v1"
	"k8s.io/client-go/tools/cache"
)

type LeaderCallbacks struct {
	// OnStartedLeading is called when starts leading
	OnStartedLeading func()
	// OnStoppedLeading is called when stops leading
	OnStoppedLeading func()
}

type LeaderManager interface {
	IsLeader() bool

	cache.ResourceEventHandler
}

var _ LeaderManager = &leaderManager{}

type leaderManager struct {
	hostname          string
	resourceName      string
	resourceNamespace string

	leaseLock     sync.RWMutex
	observedLease *coordinationv1.Lease
	observedTime  time.Time

	callbacks LeaderCallbacks

	cache.FilteringResourceEventHandler
}

func NewLeaderManager(hostname, namespace, name string, callbacks LeaderCallbacks) *leaderManager {
	_ = "STUB: not implemented"
	return nil
}

func objectToLease(obj any) *coordinationv1.Lease { _ = "STUB: not implemented"; return nil }

func (m *leaderManager) setObservedRecord(lease *coordinationv1.Lease) {
	_ = "STUB: not implemented"
	return
}

// onAdd notifies if we are the leader when lease is created.
func (m *leaderManager) onAdd(obj any) { _ = "STUB: not implemented"; return }

// Notify if we are the leader from the very begging

// onUpdate notifies when we have been elected as leader.
func (m *leaderManager) onUpdate(oldObj, newObj any) { _ = "STUB: not implemented"; return }

// Notify if we have been elected to become the leader

func (m *leaderManager) onDelete(obj any) {
	_ = "STUB: not implemented"
	// Do nothing on delete
	return
}

func (m *leaderManager) isHolderOf(lease *coordinationv1.Lease) bool {
	_ = "STUB: not implemented"
	// kube-scheduler lease id take format of `hostname + "_" + string(uuid.NewUUID())`
	return false
}

func (m *leaderManager) isLeaseValid(now time.Time) bool { _ = "STUB: not implemented"; return false }

func (m *leaderManager) IsLeader() bool { _ = "STUB: not implemented"; return false }

type dummyLeaderManager struct {
	elected bool
	cache.ResourceEventHandlerFuncs
}

var _ LeaderManager = &dummyLeaderManager{}

// NewDummyLeaderManager creates a dummy leader manager which will not change its elected state during its lifetime.
// It will always return the elected state passed in the constructor when calling IsLeader() and you will never get notified by it's channel.
//
// This is useful when disabling leader-election.
func NewDummyLeaderManager(elected bool) *dummyLeaderManager { _ = "STUB: not implemented"; return nil }

func (d *dummyLeaderManager) IsLeader() bool { _ = "STUB: not implemented"; return false }
