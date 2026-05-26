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

package config

import (
	"flag"
	"time"

	"github.com/Project-HAMi/HAMi/pkg/device/amd"
	"github.com/Project-HAMi/HAMi/pkg/device/ascend"
	"github.com/Project-HAMi/HAMi/pkg/device/awsneuron"
	"github.com/Project-HAMi/HAMi/pkg/device/cambricon"
	"github.com/Project-HAMi/HAMi/pkg/device/enflame"
	"github.com/Project-HAMi/HAMi/pkg/device/hygon"
	"github.com/Project-HAMi/HAMi/pkg/device/iluvatar"
	"github.com/Project-HAMi/HAMi/pkg/device/kunlun"
	"github.com/Project-HAMi/HAMi/pkg/device/metax"
	"github.com/Project-HAMi/HAMi/pkg/device/mthreads"
	"github.com/Project-HAMi/HAMi/pkg/device/nvidia"
	"github.com/Project-HAMi/HAMi/pkg/device/vastai"
	"github.com/Project-HAMi/HAMi/pkg/util"
)

var (
	QPS                float32
	Burst              int
	Timeout            int
	HTTPBind           string
	SchedulerName      string
	MetricsBindAddress string

	DefaultMem         int32
	DefaultCores       int32
	DefaultResourceNum int32

	// NodeSchedulerPolicy is config this scheduler node to use `binpack` or `spread`. default value is binpack.
	NodeSchedulerPolicy = util.NodeSchedulerPolicyBinpack.String()

	// NodeLabelSelector is scheduler filter node by node label.
	NodeLabelSelector map[string]string

	// NodeLockTimeout is the timeout for node locks.
	NodeLockTimeout time.Duration

	// If set to false, When Pod.Spec.SchedulerName equals to the const DefaultSchedulerName in k8s.io/api/core/v1 package, webhook will not overwrite it, default value is true.
	ForceOverwriteDefaultScheduler bool

	HostName                     string
	LeaderElect                  bool
	LeaderElectResourceName      string
	LeaderElectResourceNamespace string
)

type Config struct {
	NvidiaConfig    nvidia.NvidiaConfig       `yaml:"nvidia"`
	MetaxConfig     metax.MetaxConfig         `yaml:"metax"`
	HygonConfig     hygon.HygonConfig         `yaml:"hygon"`
	CambriconConfig cambricon.CambriconConfig `yaml:"cambricon"`
	MthreadsConfig  mthreads.MthreadsConfig   `yaml:"mthreads"`
	IluvatarConfig  []iluvatar.IluvatarConfig `yaml:"iluvatars"`
	EnflameConfig   enflame.EnflameConfig     `yaml:"enflame"`
	KunlunConfig    kunlun.KunlunConfig       `yaml:"kunlun"`
	AWSNeuronConfig awsneuron.AWSNeuronConfig `yaml:"awsneuron"`
	AMDGPUConfig    amd.AMDConfig             `yaml:"amd"`
	VastaiConfig    vastai.VastaiConfig       `yaml:"vastai"`
	VNPUs           ascend.VNPUs              `yaml:"vnpus"`
}

var (
	HandshakeAnnos = map[string]string{}
	RegisterAnnos  = map[string]string{}
	configFile     string
	DebugMode      bool
)

func InitDevicesWithConfig(config *Config) error { _ = "STUB: not implemented"; return nil }

// Helper function to initialize devices and handle errors

// Wrapper for each device's initialization function to include type assertion check

// Initialize all devices using the wrapped functions

// Initialize Ascend devices

// Initialize Iluvatar devices

// validateConfig validates the configuration object to ensure it is complete.
func validateConfig(config *Config) error { _ = "STUB: not implemented"; return nil }

func InitDevices() { _ = "STUB: not implemented"; return }

func InitDefaultDevices() { _ = "STUB: not implemented"; return }

// Initialize devices with configuration

func GlobalFlagSet() *flag.FlagSet { _ = "STUB: not implemented"; return nil }

func LoadConfig(path string) (*Config, error) { _ = "STUB: not implemented"; return nil, nil }
