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
	"go/types"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const rootPkg = "gitlab.com/gitlab-org/gitlab-runner/common"

const (
	destModuleName = "gitlab.com/rsrchboy/terraform-provider-gitlabci/third_party/gitlab/runner"
	destPkg        = "config"
	// TODO FIXME this just screams out for a flag...
	workingDir = "../../gitlab/gitlab-runner/common"
)

var destPath = filepath.Join("third_party", "gitlab", "runner", destPkg)

const marsh = `package %s

import (
	"fmt"
)

// StringOrArray is pulled from gitlab-runner/common

// // StringOrArray implements UnmarshalTOML to unmarshal either a string or array of strings.
// type StringOrArray []string

func (p *StringOrArray) UnmarshalTOML(data interface{}) error {
	switch v := data.(type) {
	case string:
		*p = StringOrArray{v}
	case []interface{}:
		for _, vv := range v {
			switch item := vv.(type) {
			case string:
				*p = append(*p, item)
			default:
				return fmt.Errorf("unexpected data type: %%v", item)
			}
		}
	default:
		return fmt.Errorf("unexpected data type: %%v", v)
	}

	return nil
}
`

// main generate Go Structures from Go structures.
// Allows to create an external module (destModuleName) used by the plugin's providers
// that contains Go structs of the dynamic configuration and nothing else.
// These Go structs do not have any non-exported fields and do not rely on any external dependencies.
func main() {
	dest := destPath

	log.Println("Output:", dest)

	err := run(dest)
	if err != nil {
		log.Fatal(err)
	}
}

func run(dest string) error {
	centrifuge, err := NewCentrifuge(rootPkg, workingDir)
	if err != nil {
		return err
	}

	centrifuge.IncludedImports = []string{
		"gitlab.com/gitlab-org/gitlab-runner/helpers/docker",
		"gitlab.com/gitlab-org/gitlab-runner/helpers/ssh",
		"gitlab.com/gitlab-org/gitlab-runner/referees",
	}

	centrifuge.ExcludedTypes = []string{
		"S3AuthType",              // string
		"InvalidTimePeriodsError", // unused
		"DockerPullPolicy",        // string
		"DockerSysCtls",           // map[string]string
	}

	centrifuge.ExcludedFiles = []string{
		"gitlab.com/gitlab-org/gitlab-runner/common/allowed_images.go",
		"gitlab.com/gitlab-org/gitlab-runner/common/build.go",
		"gitlab.com/gitlab-org/gitlab-runner/common/build_logger.go",
		"gitlab.com/gitlab-org/gitlab-runner/common/executor.go",
		"gitlab.com/gitlab-org/gitlab-runner/common/secrets.go",
		"gitlab.com/gitlab-org/gitlab-runner/common/shell.go",
		"gitlab.com/gitlab-org/gitlab-runner/common/trace.go",
		"gitlab.com/gitlab-org/gitlab-runner/common/version.go",
		"gitlab.com/gitlab-org/gitlab-runner/helpers/docker/client.go",
		"gitlab.com/gitlab-org/gitlab-runner/helpers/docker/machine.go",
		"gitlab.com/gitlab-org/gitlab-runner/helpers/docker/machine_command.go",
		"gitlab.com/gitlab-org/gitlab-runner/helpers/docker/sockets.go",
		"gitlab.com/gitlab-org/gitlab-runner/helpers/ssh/ssh_command.go",
		"gitlab.com/gitlab-org/gitlab-runner/helpers/ssh/stub_ssh_server.go",
		"gitlab.com/gitlab-org/gitlab-runner/referees/prometheus_api.go",
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

	// base is the module we're importing _from_

	if strings.Contains(typ.String(), "gitlab.com/gitlab-org/gitlab-runner/helpers/") {
		return strings.ReplaceAll(typ.String(), "gitlab.com/gitlab-org/gitlab-runner/helpers/", "")
	}

	switch typ.String() {
	case base + ".DockerPullPolicy":
		return "string"
	case base + ".DockerSysCtls":
		return "map[string]string"
	case base + ".S3AuthType":
		return "string"
	}

	if typ.String() == "[]k8s.io/api/core/v1.Capability" {
		return "[]string"
	}

	if typ.String() == "*gitlab.com/gitlab-org/gitlab-runner/referees.Config" {
		return "referees.Config"
	}

	if typ.String() == "gitlab.com/gitlab-org/gitlab-runner/helpers/docker.Credentials" {
		return "docker.Credentials"
	}

	if strings.Contains(typ.String(), base) {
		return strings.ReplaceAll(typ.String(), base+".", "")
	}

	return typ.String()
}

func cleanPackage(src string) string {
	switch src {
	case "gitlab.com/gitlab-org/gitlab-runner/helpers/docker":
		return fmt.Sprintf("%s/%s/%s", destModuleName, destPkg, "docker")
	case "gitlab.com/gitlab-org/gitlab-runner/helpers/ssh":
		return fmt.Sprintf("%s/%s/%s", destModuleName, destPkg, "ssh")
	case "gitlab.com/gitlab-org/gitlab-runner/referees":
		return fmt.Sprintf("%s/%s/%s", destModuleName, destPkg, "referees")
	default:
		return src
	}
}
