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

package version

import (
	"github.com/spf13/cobra"
)

var (
	version  = "v0.0.0-master"
	revision = "unknown" // sha1 from git, output of $(git rev-parse HEAD)

	buildDate = "unknown" // build date in ISO8601 format, output of $(date -u +'%Y-%m-%dT%H:%M:%SZ')
)

// Version information.
type Info struct {
	Version   string `json:"version"`
	Revision  string `json:"revision"`
	BuildDate string `json:"buildDate"`
	GoVersion string `json:"goVersion"`
	Compiler  string `json:"compiler"`
	Platform  string `json:"platform"`
}

// String returns a Go-syntax representation of the Info.
func (info Info) String() string { _ = "STUB: not implemented"; return "" }

// versionInfoTmpl contains the template used by Info.
var versionInfoTmpl = `
version:          {{.version}}
revision:         {{.revision}}
build date:       {{.buildDate}}
go version:       {{.goVersion}}
compiler:         {{.compiler}}
platform:         {{.platform}}
`

func Print() string { _ = "STUB: not implemented"; return "" }

func Version() Info { _ = "STUB: not implemented"; return *new(Info) }

var (
	VersionCmd = &cobra.Command{
		Use:   "version",
		Short: "print version",
		Run: func(cmd *cobra.Command, _ []string) {
			cmd.Println(Print())
		},
	}
)
