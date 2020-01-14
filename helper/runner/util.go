package runner

import (
	"log"
	"strings"

	strcase "github.com/stoewer/go-strcase"
	rcommon "gitlab.com/gitlab-org/gitlab-runner/common"
	rdhelpers "gitlab.com/gitlab-org/gitlab-runner/helpers/docker"
	rssh "gitlab.com/gitlab-org/gitlab-runner/helpers/ssh"
	"gitlab.com/gitlab-org/gitlab-runner/referees"
	"gitlab.com/rsrchboy/terraform-provider-gitlabci/internal/structs"
)

func NameForSchema(f *structs.Field) string {

	if f.IsEmbedded() {
		return ""
	}

	name := ""

	if tag := f.Tag("tf"); tag != "" {
		name = tag
	} else if tag := f.Tag("toml"); tag != "" {
		name = tag
	} else if tag := f.Tag("json"); tag != "" {
		name = tag
		// }
	} else if tag := f.Name(); tag != "" {
		// return strcase.SnakeCase(tag)
		name = tag
	}

	name = strcase.SnakeCase(name)
	name = strings.ReplaceAll(name, "-", "_")
	name = strings.TrimSuffix(name, ",omitempty")
	name = strings.TrimSuffix(name, ",omitzero")

	log.Printf("field name is: %s", name)
	return name
}

// remember: grep struct .../gitlab/gitlab-runner/common/config.go | awk '{ print $2 }' | sort | perl -nE 'chomp; say qq{\tcase "*$_" "$_":\n\t\treturn rcommon.$_} . qq!{}!'

// remember, this is only being invoked so that we have an instance of the
// struct we can deconstruct with struct.Struct et al
func NewConfigStruct(s string) interface{} {
	log.Printf("[TRACE] newConfigStruct(): %s", s)

	switch strings.TrimPrefix(s, "[]") {
	case "ssh.Config", "*ssh.Config":
		return rssh.Config{}
	case "docker_helpers.DockerCredentials":
		return rdhelpers.DockerCredentials{}
	case "*common.CacheConfig", "common.CacheConfig":
		return rcommon.CacheConfig{}
	case "*common.CacheGCSConfig", "common.CacheGCSConfig":
		return rcommon.CacheGCSConfig{}
	case "*common.CacheGCSCredentials", "common.CacheGCSCredentials":
		return rcommon.CacheGCSCredentials{}
	case "*common.CacheS3Config", "common.CacheS3Config":
		return rcommon.CacheS3Config{}
	case "*common.Config", "common.Config":
		return rcommon.Config{}
	case "*common.CustomBuildDir", "common.CustomBuildDir":
		return rcommon.CustomBuildDir{}
	case "*common.CustomConfig", "common.CustomConfig":
		return rcommon.CustomConfig{}
	case "*common.DockerConfig", "common.DockerConfig":
		return rcommon.DockerConfig{}
	case "*common.DockerService", "common.DockerService":
		return rcommon.DockerService{}
	case "*common.DockerMachine", "common.DockerMachine":
		return rcommon.DockerMachine{}
	case "*common.KubernetesConfig", "common.KubernetesConfig":
		return rcommon.KubernetesConfig{}
	case "*common.KubernetesConfigMap", "common.KubernetesConfigMap":
		return rcommon.KubernetesConfigMap{}
	case "*common.KubernetesEmptyDir", "common.KubernetesEmptyDir":
		return rcommon.KubernetesEmptyDir{}
	case "*common.KubernetesHostPath", "common.KubernetesHostPath":
		return rcommon.KubernetesHostPath{}
	case "*common.KubernetesPodSecurityContext", "common.KubernetesPodSecurityContext":
		return rcommon.KubernetesPodSecurityContext{}
	case "*common.KubernetesPVC", "common.KubernetesPVC":
		return rcommon.KubernetesPVC{}
	case "*common.KubernetesSecret", "common.KubernetesSecret":
		return rcommon.KubernetesSecret{}
	case "*common.KubernetesVolumes", "common.KubernetesVolumes":
		return rcommon.KubernetesVolumes{}
	case "*common.ParallelsConfig", "common.ParallelsConfig":
		return rcommon.ParallelsConfig{}
	case "*common.RunnerConfig", "common.RunnerConfig":
		return rcommon.RunnerConfig{}
	case "*common.RunnerCredentials", "common.RunnerCredentials":
		return rcommon.RunnerCredentials{}
	case "*common.RunnerSettings", "common.RunnerSettings":
		return rcommon.RunnerSettings{}
	case "*common.Service", "common.Service":
		return rcommon.Service{}
	case "*common.SessionServer", "common.SessionServer":
		return rcommon.SessionServer{}
	case "*common.VirtualBoxConfig", "common.VirtualBoxConfig":
		return rcommon.VirtualBoxConfig{}
	case "*referees.Config":
		return referees.Config{}
	case "*referees.MetricsReferee", "referees.MetricsReferee":
		return referees.MetricsReferee{}
	case "*referees.MetricsRefereeConfig", "referees.MetricsRefereeConfig":
		return referees.MetricsRefereeConfig{}
	case "[]string":
		return []string{}
	case "int", "int64", "bool", "string":
		return nil
	default:
		panic("unhandled type: " + s)
	}
}
