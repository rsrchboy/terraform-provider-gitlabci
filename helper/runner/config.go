package runner

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	tree "github.com/DiSiqueira/GoTree"
	"github.com/davecgh/go-spew/spew"
	// "github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	// "github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/imdario/mergo"
	"github.com/mitchellh/mapstructure"
	"github.com/mohae/deepcopy"
	"gitlab.com/rsrchboy/terraform-provider-gitlabci/internal/structs"
	// strcase "github.com/stoewer/go-strcase"
	rcommon "gitlab.com/gitlab-org/gitlab-runner/common"
	// rdhelpers "gitlab.com/gitlab-org/gitlab-runner/helpers/docker"
	// "gitlab.com/gitlab-org/gitlab-runner/helpers/ssh"
	// "gitlab.com/gitlab-org/gitlab-runner/referees"
)

type schemaMap map[string]*schema.Schema
type iMap map[string]interface{}
type stringMap map[string]string

type processFunc func(
	block *map[string]interface{},
	info *FieldInfo,
	dCfg *mapstructure.DecoderConfig,
) error

// type FieldInfoMap map[string]*FieldInfo
type FieldInfoMap map[string]*FieldInfo

type FieldInfo struct {
	Type       string
	NotStruct  bool
	NoFlatten  bool
	IsEmbedded bool
	IsList     bool
	// Fields      map[string]FieldInfo
	Fields      FieldInfoMap
	ProcessFunc processFunc
	// schema generation bits
	OverrideSchema *schema.Schema // use this instead of generating our own
	// schemaMap      map[string]*schema.Schema // all the fields in this Type
	// schemaFields *schemaMap // all the fields in this Type
	schemaFields schemaMap // all the fields in this Type
	schema       *schema.Schema
	Name         string
	Description  string
	Tree         tree.Tree
	Processed    bool
}

var cfgStructs = &FieldInfo{
	Type:      "common.RunnerConfig",
	Processed: false,
	Fields: FieldInfoMap{
		// parent struct
		// "config":           FieldInfo{Type: "common.Config"},
		// sibling
		// "session_server":                  FieldInfo{Type: "common.SessionServer"},
		// this struct
		// "runner_config":                   FieldInfo{Type: "common.RunnerConfig"},
		// embedded
		// "runner_credentials":              FieldInfo{Type: "common.RunnerCredentials"},
		// "runner_settings":                 FieldInfo{Type: "common.RunnerSettings"},
		"environment": &FieldInfo{
			Type: "[]string",
			OverrideSchema: &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				ForceNew: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
		},
		"custom_build_dir": &FieldInfo{Type: "common.CustomBuildDir"},
		"referees":         &FieldInfo{Type: "referees.Config"},
		"cache": &FieldInfo{
			Type: "common.CacheConfig",
			// Fields: map[string]&FieldInfo{
			Fields: FieldInfoMap{
				"s3":  &FieldInfo{Type: "common.CacheS3Config"},
				"gcs": &FieldInfo{Type: "common.CacheGCSConfig"},
				// embedded in gcs
				// "gcs_credentials": &FieldInfo{Type: "common.CacheGCSCredentials"},
			},
		},
		// TODO check ssh.Config
		"ssh": &FieldInfo{Type: "ssh.Config"},
		"docker": &FieldInfo{
			Type: "common.DockerConfig",
			// Fields: map[string]&FieldInfo{
			Fields: FieldInfoMap{
				"pull_policy": &FieldInfo{Type: "common.DockerPullPolicy"},
				"sysctls":     &FieldInfo{Type: "common.DockerSysCtls"},
				"services":    &FieldInfo{Type: "common.DockerService"},
			},
		},
		"custom":      &FieldInfo{Type: "common.CustomConfig"},
		"machine":     &FieldInfo{Type: "common.DockerMachine"},
		"parallels":   &FieldInfo{Type: "common.ParallelsConfig"},
		"virtual_box": &FieldInfo{Type: "common.VirtualBoxConfig"},
		"kubernetes": &FieldInfo{
			Type: "common.KubernetesConfig",
			// Fields: map[string]&FieldInfo{
			Fields: FieldInfoMap{
				"pod_security_context": &FieldInfo{Type: "common.KubernetesPodSecurityContext"},
				"volumes": &FieldInfo{
					Type: "common.KubernetesVolumes",
					// Fields: map[string]&FieldInfo{
					Fields: FieldInfoMap{
						// HostPaths  []KubernetesHostPath  `toml:"host_path" description:"The host paths which will be mounted" json:"host_paths"`
						// PVCs       []KubernetesPVC       `toml:"pvc" description:"The persistent volume claims that will be mounted" json:"pv_cs"`
						// ConfigMaps []KubernetesConfigMap `toml:"config_map" description:"The config maps which will be mounted as volumes" json:"config_maps"`
						// Secrets    []KubernetesSecret    `toml:"secret" description:"The secret maps which will be mounted" json:"secrets"`
						// EmptyDirs  []KubernetesEmptyDir  `toml:"empty_dir" description:"The empty dirs which will be mounted" json:"empty_dirs"`
						// "kubernetes_empty_dir":            &FieldInfo{Type: "common.KubernetesEmptyDir"},
						// "kubernetes_secret":               &FieldInfo{Type: "common.KubernetesSecret"},
						// "kubernetes_p_v_c":                &FieldInfo{Type: "common.KubernetesPVC"},
						// "kubernetes_host_path":            &FieldInfo{Type: "common.KubernetesHostPath"},
						// "kubernetes_config_map":           &FieldInfo{Type: "common.KubernetesConfigMap"},
					},
				},
				"services":    &FieldInfo{Type: "common.Service"},
				"pull_policy": &FieldInfo{Type: "common.KubernetesPullPolicy"},
			},
		},
	},
}

func RunnerConfigToTerraformSchema() schemaMap {

	if cfgStructs.Processed {
		return cfgStructs.SchemaFields()
	}

	tree := tree.New("runner_config")
	cfgStructs.Tree = tree
	schema := cfgStructs.SchemaFields()
	log.Printf("Schema tree looks like:\n%s", tree.Print())
	// return cfgStructs.SchemaFields()
	return schema
}

// func (self *FieldInfo) infoToSchema() map[string]*schema.Schema {
func (info *FieldInfo) ToSchema() *schema.Schema {

	if info.schema != nil {
		return info.schema
	}

	// easy case! :)
	if info.OverrideSchema != nil {
		info.schema = info.OverrideSchema
		return info.schema
	}

	me := schema.Schema{
		Optional:    true,
		Description: info.Description,
	}
	info.schema = &me

	handled := true

	switch info.Type {
	case "[]string", "*[]string", "[]*string":
		me.Type = schema.TypeList
		me.Elem = &schema.Schema{Type: schema.TypeString}
		info.IsList = true
	case "[]*int", "[]*int64", "[]int", "[]int64":
		me.Type = schema.TypeList
		me.Elem = &schema.Schema{Type: schema.TypeInt}
		info.IsList = true
	case "map[string]string", "common.DockerSysCtls":
		me.Type = schema.TypeMap
		me.Elem = &schema.Schema{Type: schema.TypeString}
	case "string", "*string", "common.DockerPullPolicy", "common.KubernetesPullPolicy":
		me.Type = schema.TypeString
	case "*int", "*int64", "int", "int64":
		me.Type = schema.TypeInt
	case "*bool", "bool":
		me.Type = schema.TypeBool
	default:
		log.Printf("unhandled type: %s", info.Type)
		handled = false
	}

	if !handled {
		schemaFields := info.SchemaFields()

		me.Type = schema.TypeList
		me.Elem = &schema.Resource{Schema: schemaFields}
	}

	// if TypeList and aren't flagged as being a list, only allow one
	if !info.IsList && me.Type == schema.TypeList {
		me.MinItems = 0
		me.MaxItems = 1
	}

	return info.schema
}

func (info *FieldInfo) SchemaFields() schemaMap {
	// info.schemaFields = make(map[string]*schema.Schema)
	info.schemaFields = make(schemaMap)
	info.Fields = make(FieldInfoMap)

	cs := NewConfigStruct(info.Type)

	if cs == nil {
		return info.schemaFields
	}

	s := structs.New(cs)

	for _, f := range s.Fields() {

		// skip this field entirely if it's unexported
		if !f.IsExported() {
			log.Printf("[INFO] FieldToSchemaResource() found field: %s: not exported! skipping", f.Name())
			continue
		}

		name := NameForSchema(f)
		log.Printf("[INFO] %s, tag %s", f.Name(), name)

		var child *FieldInfo

		// if we don't have this field, create a bog-standard one
		if _, hasField := info.Fields[name]; !hasField {
			typeName := f.ReflectValue().Type().String()
			// info.Fields[name] = &FieldInfo{Type: typeName, NotStruct: true}
			child = newFieldInfo(name, f, typeName)
		} else {
			info.Fields[name].Name = name
			child = info.Fields[name]
		}

		if child.IsEmbedded || name == "" {
			// mergo.Merge(&info.schemaFields, info.Fields[name].SchemaFields())
			log.Printf("[INFO] %s, tag %s -- embedded", f.Name(), name)
			child.Tree = info.Tree
			mergo.Merge(&info.schemaFields, child.SchemaFields())
		} else {
			log.Printf("[INFO] %s, tag %s -- not embedded", f.Name(), name)
			child.Tree = info.Tree.Add(name)
			info.Fields[name] = child
			info.schemaFields[name] = child.ToSchema()
			// schemaFields[name] = child.ToSchema()
		}
	}

	return info.schemaFields
}

func newFieldInfo(name string, f *structs.Field, typeName string) *FieldInfo {
	info := FieldInfo{
		Type:        typeName, // strings.TrimPrefix(typeName, "[]"),
		Name:        NameForSchema(f),
		Description: f.Tag("description"),
		IsEmbedded:  false, // f.IsEmbedded(),
		NotStruct:   true,
		NoFlatten:   true,
		IsList:      strings.HasPrefix(typeName, "[]"),
	}

	return &info
}

func flattenBlock(block *map[string]interface{}, info *FieldInfo, dCfg *mapstructure.DecoderConfig) (interface{}, error) {
	log.Printf("[TRACE] flattenBlock %s", info.Type)

	for col, colInfo := range info.Fields {
		// check to see if we need to do any flattening
		log.Printf("[TRACE] Flattening col: %s", col)
		if colInfo.NoFlatten || (*block)[col] == nil {
			continue
		}

		switch (*block)[col].(type) {
		case []interface{}:
			// log.Printf("flattening: %T", value)
		default:
			continue
		}
		colVal := (*block)[col].([]interface{})
		if len(colVal) > 0 {
			log.Printf("%s is an array > 0", col)
			thing := colVal[0].(map[string]interface{})
			_, err := flattenBlock(&thing, colInfo, nil)
			if err != nil {
				log.Printf("errored!: %v", err)
				return nil, err
			}
			(*block)[col] = thing
		} else {
			log.Printf("%s is an array == 0; deleting", col)
			delete(*block, col)
		}
	}

	log.Printf("[TRACE] runnerBlock redux: %s", spew.Sdump(block))

	if dCfg == nil {
		return nil, nil
	} else {
		var r rcommon.RunnerConfig

		log.Printf("[TRACE] block ended up as: %s", spew.Sdump(block))
		// bytes, err := json.Marshal(dCfg.Result)
		bytes, err := json.Marshal(block)
		if err != nil {
			log.Printf("errored!: %v", err)
			return nil, fmt.Errorf("serialization of runner config failed: %v", err)
		}
		log.Printf("[TRACE] json serialized to %s", bytes)

		err = json.Unmarshal(bytes, &r)
		// err = json.Unmarshal(bytes, dCfg.Result)
		if err != nil {
			log.Printf("errored!: %v", err)
			return nil, fmt.Errorf("deserialization of runner config failed: %v", err)
		}

		log.Printf("[TRACE] rc: %s", spew.Sdump(r))
		// log.Printf("[TRACE] *********** rc: %#v", dCfg.Result)
		// c.Runners = append(c.Runners, &rc)
		return &r, nil
	}
}

func flattenKey(block *interface{}) map[string]interface{} {
	runnerBlock := deepcopy.Copy(*block)
	for key, value := range runnerBlock.(map[string]interface{}) {
		// check to see if we need to do any flattening
		log.Printf("[TRACE] Flattening key: %s", key)
		if value == nil {
			log.Printf("skipping: %T", value)
			continue
		}
		switch value.(type) {
		case []interface{}:
			log.Printf("flattening: %T", value)
		default:
			continue
		}
		if len(value.([]interface{})) > 0 {
			log.Printf("%s is an array > 0", key)
			thing := value.([]interface{})[0]
			thing = flattenKey(&thing)
			runnerBlock.(map[string]interface{})[key] = thing
		} else {
			log.Printf("%s is an array == 0", key)
		}
		log.Printf("[TRACE] runnerBlock redux: %s", spew.Sdump(runnerBlock))
	}
	return runnerBlock.(map[string]interface{})
}
