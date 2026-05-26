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
	"github.com/prometheus/client_golang/prometheus"
)

type ClusterManager struct {
	Zone          string
	LegacyMetrics bool
}

// ClusterManagerCollector implements the Collector interface.
type ClusterManagerCollector struct {
	ClusterManager *ClusterManager
}

// Describe is implemented with DescribeByCollect. That's possible because the
// Collect method will always return the same metrics with the same descriptors.
func (cc ClusterManagerCollector) Describe(ch chan<- *prometheus.Desc) {
	_ = "STUB: not implemented"
	return
}

// Collect creates constant metrics for each host on the fly based on the returned data.
func (cc ClusterManagerCollector) Collect(ch chan<- prometheus.Metric) {
	_ = "STUB: not implemented"
	return
}

// New metric descriptors

// Legacy metric descriptors (only created when legacy mode is enabled)

// NewClusterManager creates a ClusterManager and registers its collector.
func NewClusterManager(zone string, reg prometheus.Registerer, legacyMetrics bool) *ClusterManager {
	_ = "STUB: not implemented"
	return nil
}

func initMetrics(bindAddress string, legacyMetrics bool) { _ = "STUB: not implemented"; return }
