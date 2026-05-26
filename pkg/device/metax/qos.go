/*
Copyright 2025 The HAMi Authors.

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

package metax

import (
	"sync"
)

type JitteryQosCache struct {
	sync.Mutex
	cache map[string]string
}

func NewJitteryQosCache() *JitteryQosCache { _ = "STUB: not implemented"; return nil }

func (c *JitteryQosCache) Sync(devices []*MetaxSDeviceInfo) { _ = "STUB: not implemented"; return }

func (c *JitteryQosCache) Add(uuid string, expectedQos string) { _ = "STUB: not implemented"; return }

func (c *JitteryQosCache) Get(uuid string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}
