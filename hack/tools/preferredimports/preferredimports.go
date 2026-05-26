/*
Copyright 2019 The Kubernetes Authors.

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

// This code is directly lifted from the Kubernetes codebase in order to avoid relying on the k8s.io/kubernetes package.
// For reference: https://github.com/kubernetes/kubernetes/blob/release-1.22/cmd/preferredimports/preferredimports.go

// verify that all the imports have our preferred alias(es).
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"go/ast"
	"go/build"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"

	"golang.org/x/term"
)

var (
	importAliases    = flag.String("import-aliases", "hack/.import-aliases", "json file with import aliases")
	confirm          = flag.Bool("confirm", false, "update file with the preferred aliases for imports")
	includePathRegex = flag.String("include-path", "(test/e2e/|test/e2e_node)", "only files with paths matching this regex is touched")
	excludePathRegex = flag.String("exclude-path", "(testing)", "files with paths matching this regex is ignored")
	isTerminal       = term.IsTerminal(int(os.Stdout.Fd()))
	logPrefix        = ""
	aliases          map[string]string
)

type analyzer struct {
	fset      *token.FileSet // positions are relative to fset
	ctx       build.Context
	failed    bool
	donePaths map[string]any
}

func newAnalyzer() *analyzer { _ = "STUB: not implemented"; return nil }

// collect extracts test metadata from a file.
func (a *analyzer) collect(dir string) { _ = "STUB: not implemented"; return }

// Just collect all the parsed files in a slice, no need for ast.Package

// Iterate directly over the files

func renameImportUsages(f *ast.File, old, new string) {
	_ = "STUB: not implemented"
	// use this to avoid renaming the package declaration, eg:
	//
	//	given: package foo; import foo "bar"; foo.Baz, rename foo->qux
	//	yield: package foo; import qux "bar"; qux.Baz
	return
}

// Rename top-level old to new, both unresolved names
// (probably defined in another file) and names that resolve
// to a declaration we renamed.

type collector struct {
	dirs             []string
	includePathRegex *regexp.Regexp
	excludePathRegex *regexp.Regexp
}

// handlePath walks the filesystem recursively, collecting directories,
// ignoring some unneeded directories (hidden/vendored) that are handled
// specially later.
func (c *collector) handlePath(path string, info os.FileInfo, err error) error {
	_ = "STUB: not implemented"
	return nil
}

// Ignore hidden directories (.git, .cache, etc)

// Staging code is symlinked from vendor/k8s.io, and uses import
// paths as if it were inside of vendor/. It fails typechecking
// inside of staging/, but works when typechecked as part of vendor/.

// OS-specific vendor code tends to be imported by OS-specific
// packages. We recursively typecheck imported vendored packages for
// each OS, but don't typecheck everything for every OS.

// This is a weird one. /testdata/ is *mostly* ignored by Go,
// and this translates to kubernetes/vendor not working.
// edit/record.go doesn't compile without gopkg.in/yaml.v2
// in $GOSRC/$GOROOT (both typecheck and the shell script).

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		args = append(args, ".")
	}

	includePathRegex, err := regexp.Compile(*includePathRegex)
	if err != nil {
		log.Fatalf("Error compiling regex: %v", err)
	}
	excludePathRegex, err := regexp.Compile(*excludePathRegex)
	if err != nil {
		log.Fatalf("Error compiling regex: %v", err)
	}
	c := collector{includePathRegex: includePathRegex, excludePathRegex: excludePathRegex}
	for _, arg := range args {
		err := filepath.Walk(arg, c.handlePath)
		if err != nil {
			log.Fatalf("Error walking: %v", err)
		}
	}
	sort.Strings(c.dirs)

	if len(*importAliases) > 0 {
		bytes, err := os.ReadFile(*importAliases)
		if err != nil {
			log.Fatalf("Error reading import aliases: %v", err)
		}
		err = json.Unmarshal(bytes, &aliases)
		if err != nil {
			log.Fatalf("Error loading aliases: %v", err)
		}
	}
	if isTerminal {
		logPrefix = "\r" // clear status bar when printing
	}
	fmt.Println("checking-imports: ")

	a := newAnalyzer()
	for _, dir := range c.dirs {
		if isTerminal {
			fmt.Printf("\r\033[0m %-80s\n", dir)
		}
		a.collect(dir)
	}
	fmt.Println()
	if a.failed {
		os.Exit(1)
	}
}
