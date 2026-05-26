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

package client

import (
	"sync"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type Client struct {
	// Embedded kubernetes.Interface to avoid name conflicts.
	kubernetes.Interface
	config *rest.Config
}

var (
	KubeClient kubernetes.Interface
	once       sync.Once
)

func init() {
	KubeClient = nil
}

// GetClient returns the global Kubernetes client.
func GetClient() kubernetes.Interface {
	_ = "STUB: not implemented"

	// NewClient creates a new Kubernetes client with the given options.
	return *new(kubernetes.Interface)
}

func NewClient(opts ...Option) (*Client, error) { _ = "STUB: not implemented"; return nil, nil }

// Apply WithDefaults option first to set default values.

// Then apply user-provided options that will override defaults if specified.

// InitGlobalClient initializes the global Kubernetes client with the given options.
func InitGlobalClient(opts ...Option) { _ = "STUB: not implemented"; return }

// loadKubeConfig loads Kubernetes configuration from the environment or in-cluster.
func loadKubeConfig() (*rest.Config, error) { _ = "STUB: not implemented"; return nil, nil }
