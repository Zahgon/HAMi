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

const DirectLinkScore = 10

type LinkDevice struct {
	uuid     string
	linkZone int32
}

func (from *LinkDevice) score(to *LinkDevice) int { _ = "STUB: not implemented"; return 0 }

type LinkDevices []*LinkDevice

func (devs LinkDevices) Score() int { _ = "STUB: not implemented"; return 0 }

func (devs LinkDevices) String() string { _ = "STUB: not implemented"; return "" }
