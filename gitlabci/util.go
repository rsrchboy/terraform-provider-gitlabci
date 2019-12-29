package gitlabci

import (
	"crypto/sha256"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	// strcase "github.com/stoewer/go-strcase"
	// "github.com/davecgh/go-spew/spew"
	rcommon "gitlab.com/gitlab-org/gitlab-runner/common"
	rdhelpers "gitlab.com/gitlab-org/gitlab-runner/helpers/docker"
	rssh "gitlab.com/gitlab-org/gitlab-runner/helpers/ssh"
	"gitlab.com/gitlab-org/gitlab-runner/referees"
)

func toStringMap(key string, d *schema.ResourceData) stringMap {
	imap := d.Get(key).(map[string]interface{})
	// log.Printf("iMapToStringMap: %T, %s", imap, spew.Sdump(imap))
	smap := make(stringMap, len(imap))

	for k, v := range imap {
		smap[k] = v.(string)
	}

	return smap
}

func stringList(key string, d *schema.ResourceData) []string {
	stringsI := d.Get(key).([]interface{})
	strings := make([]string, len(stringsI))

	// I'm hopeful there's a better way I'm simply unaware of as of yet
	for i, str := range stringsI {
		strings[i] = str.(string)
	}

	return strings
}

// copied from the gitlab provider
func stringSetToStringSlice(stringSet *schema.Set) *[]string {
	ret := []string{}
	if stringSet == nil {
		return &ret
	}
	for _, envVal := range stringSet.List() {
		ret = append(ret, envVal.(string))
	}
	return &ret
}

// copied from the gitlab provider
func hashSum(contents interface{}) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(contents.(string))))
}

// remember: grep struct .../gitlab/gitlab-runner/common/config.go | awk '{ print $2 }' | sort | perl -nE 'chomp; say qq{\tcase "*$_" "$_":\n\t\treturn rcommon.$_} . qq!{}!'

// remember, this is only being invoked so that we have an instance of the
// struct we can deconstruct with struct.Struct et al
func newConfigStruct(s string) interface{} {
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
	default:
		panic("unhandled type: " + s)
	}
}
