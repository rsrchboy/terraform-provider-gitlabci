// The MIT License (MIT)
//
// Copyright (c) 2016-2020 Containous SAS; 2020-2022 Traefik Labs
//
// Permission is hereby granted, free of charge, to any person obtaining a
// copy of this software and associated documentation files (the "Software"),
// to deal in the Software without restriction, including without limitation
// the rights to use, copy, modify, merge, publish, distribute, sublicense,
// and/or sell copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER
// DEALINGS IN THE SOFTWARE.

package main

import (
	"fmt"
	"go/build"
	"go/types"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

const rootPkg = "github.com/traefik/traefik/v2/pkg/config/dynamic"

const (
	destModuleName = "github.com/traefik/genconf"
	destPkg        = "dynamic"
)

const marsh = `package %s

import "encoding/json"

type JSONPayload struct {
	*Configuration
}

func (c JSONPayload) MarshalJSON() ([]byte, error) {
	if c.Configuration == nil {
		return nil, nil
	}

	return json.Marshal(c.Configuration)
}
`

// main generate Go Structures from Go structures.
// Allows to create an external module (destModuleName) used by the plugin's providers
// that contains Go structs of the dynamic configuration and nothing else.
// These Go structs do not have any non-exported fields and do not rely on any external dependencies.
func main() {
	dest := filepath.Join(path.Join(build.Default.GOPATH, "src"), destModuleName, destPkg)

	log.Println("Output:", dest)

	err := run(dest)
	if err != nil {
		log.Fatal(err)
	}
}

func run(dest string) error {
	centrifuge, err := NewCentrifuge(rootPkg)
	if err != nil {
		return err
	}

	centrifuge.IncludedImports = []string{
		"github.com/traefik/traefik/v2/pkg/tls",
		"github.com/traefik/traefik/v2/pkg/types",
	}

	centrifuge.ExcludedTypes = []string{
		// tls
		"CertificateStore", "Manager",
		// dynamic
		"Message", "Configurations",
		// types
		"HTTPCodeRanges", "HostResolverConfig",
	}

	centrifuge.ExcludedFiles = []string{
		"github.com/traefik/traefik/v2/pkg/types/logs.go",
		"github.com/traefik/traefik/v2/pkg/types/metrics.go",
	}

	centrifuge.TypeCleaner = cleanType
	centrifuge.PackageCleaner = cleanPackage

	err = centrifuge.Run(dest, destPkg)
	if err != nil {
		return err
	}

	return os.WriteFile(filepath.Join(dest, "marshaler.go"), []byte(fmt.Sprintf(marsh, destPkg)), 0o666)
}

func cleanType(typ types.Type, base string) string {
	if typ.String() == "github.com/traefik/traefik/v2/pkg/tls.FileOrContent" {
		return "string"
	}

	if typ.String() == "[]github.com/traefik/traefik/v2/pkg/tls.FileOrContent" {
		return "[]string"
	}

	if typ.String() == "github.com/traefik/paerser/types.Duration" {
		return "string"
	}

	if strings.Contains(typ.String(), base) {
		return strings.ReplaceAll(typ.String(), base+".", "")
	}

	if strings.Contains(typ.String(), "github.com/traefik/traefik/v2/pkg/") {
		return strings.ReplaceAll(typ.String(), "github.com/traefik/traefik/v2/pkg/", "")
	}

	return typ.String()
}

func cleanPackage(src string) string {
	switch src {
	case "github.com/traefik/paerser/types":
		return ""
	case "github.com/traefik/traefik/v2/pkg/tls":
		return path.Join(destModuleName, destPkg, "tls")
	case "github.com/traefik/traefik/v2/pkg/types":
		return path.Join(destModuleName, destPkg, "types")
	default:
		return src
	}
}
