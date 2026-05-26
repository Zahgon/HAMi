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
	"context"
	"errors"

	"github.com/NVIDIA/go-nvml/pkg/nvml"

	"github.com/Project-HAMi/HAMi/pkg/monitor/nvidia"
)

var cgroupDriver int
var errTemporaryClosed = errors.New("temporary closed")

//type hostGPUPid struct {
//	hostGPUPid int
//	mtime      uint64
//}

type UtilizationPerDevice []int

func setcGgroupDriver() int {
	_ = "STUB: not implemented"
	// 1 for cgroupfs 2 for systemd
	return 0
}

func getUsedGPUPid() ([]uint, nvml.Return) {
	_ = "STUB: not implemented"
	return nil, *new(nvml.Return)
}

func CheckBlocking(utSwitchOn map[string]UtilizationPerDevice, p int, c *nvidia.ContainerUsage) bool {
	_ = "STUB: not implemented"
	return false
}

// Check whether task with higher priority use GPU or there are other tasks with the same priority.
func CheckPriority(utSwitchOn map[string]UtilizationPerDevice, p int, c *nvidia.ContainerUsage) bool {
	_ = "STUB: not implemented"
	return false
}

func Observe(lister *nvidia.ContainerLister) { _ = "STUB: not implemented"; return }

//for _, devuuid := range val.sr.uuids {
// Null device condition

func watchAndFeedback(ctx context.Context, lister *nvidia.ContainerLister, migLockSignal <-chan bool) error {
	_ = "STUB: not implemented"
	return nil
}
