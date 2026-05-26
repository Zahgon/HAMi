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

package main

import (
	"github.com/Project-HAMi/HAMi/pkg/device/nvidia"

	spec "github.com/NVIDIA/k8s-device-plugin/api/config/v1"
	cli "github.com/urfave/cli/v2"
)

func addFlags() []cli.Flag { _ = "STUB: not implemented"; return nil }

// updateFromCLIFlag conditionally updates the config flag at 'pflag' to the value of the CLI flag with name 'flagName'.
func updateFromCLIFlag[T any](pflag **T, c *cli.Context, flagName string) {
	_ = "STUB: not implemented"
	return
}

func generateDeviceConfigFromNvidia(cfg *spec.Config, c *cli.Context, flags []cli.Flag) (nvidia.DeviceConfig, error) {
	_ = "STUB: not implemented"
	return *new(nvidia.DeviceConfig), nil
}

// Common flags
