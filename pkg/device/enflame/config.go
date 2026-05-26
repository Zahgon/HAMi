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

package enflame

import "flag"

var (
	EnflameResourceNameGCU            string
	EnflameResourceNameDRSGCU         string
	EnflameResourceNameGCUMemory      string
	EnflameResourceNameGCUCore        string
	EnflameResourceNameVGCU           string
	EnflameResourceNameVGCUPercentage string
)

type EnflameConfig struct {
	// GCU
	ResourceNameGCU string `yaml:"resourceNameGCU"`

	// DRS-GCU (new hard-partition mode)
	ResourceNameDRSGCU string `yaml:"resourceNameDRSGCU"`
	ResourceNameMemory string `yaml:"resourceNameGCUMemory"`
	ResourceNameCore   string `yaml:"resourceNameGCUCore"`

	// Legacy shared-GCU key kept for compatibility.
	ResourceNameVGCU           string `yaml:"resourceNameVGCU"`
	ResourceNameVGCUPercentage string `yaml:"resourceNameVGCUPercentage"`
}

func ParseConfig(fs *flag.FlagSet) {
	_ = "STUB: not implemented"
	// GCU
	return
}

// DRS-GCU.

// Legacy flag alias for backward compatibility.

// Legacy shared-GCU related flags.
