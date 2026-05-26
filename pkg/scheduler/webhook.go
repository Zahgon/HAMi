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

	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

const template = "Processing admission hook for pod %v/%v, UID: %v"

type webhook struct {
	decoder admission.Decoder
}

func NewWebHook() (*admission.Webhook, error) { _ = "STUB: not implemented"; return nil, nil }

func (h *webhook) Handle(_ context.Context, req admission.Request) admission.Response {
	_ = "STUB: not implemented"
	return *new(admission.Response)
}

//return admission.Allowed("no resource found")

func fitResourceQuota(pod *corev1.Pod) bool { _ = "STUB: not implemented"; return false }

// Only supports NVIDIA
