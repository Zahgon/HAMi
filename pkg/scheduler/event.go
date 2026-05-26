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
	corev1 "k8s.io/api/core/v1"
)

// Define events for ResourceBinding, ResourceFilter objects and their associated resources.
const (
	// EventReasonFilteringFailed indicates that filtering failed.
	EventReasonFilteringFailed = "FilteringFailed"
	// EventReasonFilteringSucceed indicates that filtering succeed.
	EventReasonFilteringSucceed = "FilteringSucceed"

	// EventReasonBindingFailed indicates that  binding failed.
	EventReasonBindingFailed = "BindingFailed"
	// EventReasonBindingSucceed indicates that  binding succeed.
	EventReasonBindingSucceed = "BindingSucceed"
)

func (s *Scheduler) addAllEventHandlers() { _ = "STUB: not implemented"; return }

func (s *Scheduler) recordScheduleBindingResultEvent(pod *corev1.Pod, eventReason string, nodeResult []string, schedulerErr error) {
	_ = "STUB: not implemented"
	return
}

func (s *Scheduler) recordScheduleFilterResultEvent(pod *corev1.Pod, eventReason string, successMsg string, schedulerErr error) {
	_ = "STUB: not implemented"
	// eventRecorder maybe  nil
	return
}
