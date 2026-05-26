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

package routes

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/Project-HAMi/HAMi/pkg/scheduler"
)

const maxRequestSize = 1024 * 1024 // 1MB limit

func checkBody(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

func PredicateRoute(s *scheduler.Scheduler) httprouter.Handle {
	_ = "STUB: not implemented"
	return *new(httprouter.Handle)
}

// Limit the body size to prevent deep nesting/resource exhaustion attacks

// Poll may return false when context is cancelled

func Bind(s *scheduler.Scheduler) httprouter.Handle {
	_ = "STUB: not implemented"
	return *new(httprouter.Handle)
}

// Limit the body size to prevent deep nesting/resource exhaustion attacks

func WebHookRoute() httprouter.Handle { _ = "STUB: not implemented"; return *new(httprouter.Handle) }

func HealthzRoute() httprouter.Handle { _ = "STUB: not implemented"; return *new(httprouter.Handle) }

func ReadyzRoute(s *scheduler.Scheduler) httprouter.Handle {
	_ = "STUB: not implemented"
	return *new(httprouter.Handle)
}
