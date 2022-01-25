package main

import (
	"bytes"
	"errors"
	"fmt"
	"go/format"
	"log"
	"os"
	"reflect"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stoewer/go-strcase"
	"gitlab.com/rsrchboy/terraform-provider-gitlabci/third_party/gitlab/runner/config"
	"golang.org/x/tools/imports"
)

// Yes, this is probably the messiest bit of code I've written in... ages.
// For now, however, it works, and I'd rather like to get the config
// templating working again!

// I'm not going to worry about spacing, alignment, etc, as we effectively run
// gofmt on the rendered output before spewing it out into a file.
const tmplString = `

{{define "schemaSchema" }}
	Computed:    false,
	Optional:    true,
	Description: "{{. | nodeDesc | replace "\"" "\\\"" }}",
{{ $tpl := nodeElemTemplate . }}
{{- if eq $tpl "elemNull" -}}
// elemNull
// {{ $tpl }}
{{- else if eq $tpl "elemString" -}}
	Type: schema.TypeString,
{{- else if eq $tpl "elemBool" -}}
	Type: schema.TypeBool,
{{- else if eq $tpl "elemInt" -}}
	Type: schema.TypeInt,
{{- else if eq $tpl "elemFloat" -}}
	Type: schema.TypeFloat,
{{- else if eq $tpl "elemMapString" -}}
	Type: schema.TypeMap,
	Elem: &schema.Schema{Type: schema.TypeString},
{{- else if eq $tpl "elemMapBool" -}}
	Type: schema.TypeMap,
	Elem: &schema.Schema{Type: schema.TypeBool},
{{- else if eq $tpl "elemSliceString" -}}
	Type: schema.TypeList,
	Elem: &schema.Schema{Type: schema.TypeString},
{{- else if eq $tpl "elemSliceInt" -}}
	Type: schema.TypeList,
	Elem: &schema.Schema{Type: schema.TypeInt},
{{- else if eq $tpl "elemSliceFloat" -}}
	Type: schema.TypeList,
	Elem: &schema.Schema{Type: schema.TypeFloat},
{{- else if eq $tpl "elemSliceStruct" -}}
	Type: schema.TypeList,
	Elem: &schema.Resource{
		Schema: {{ template "schemaSchemaSlice" .Type.Elem }} },
	},
{{- else if eq $tpl "elemPtrStruct" -}}
	Type: schema.TypeList,
	MinItems:    0,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: {{ template "schemaSchemaSlice" .Type.Elem }} },
	},
{{- else if eq $tpl "elemSlicePtrStruct" -}}
	Type: schema.TypeList,
	Elem: &schema.Resource{
		Schema: {{ template "schemaSchemaSlice" .Type.Elem.Elem }} },
	},
{{- else if eq $tpl "elemStruct" -}}
	Type: schema.TypeList,
	MinItems:    0,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: {{ template "schemaSchemaSlice" .Type }} },
	},
{{- else -}}

	// FIXME unknown: {{ $tpl }}

{{end -}}

{{end}}


{{- define "schemaSchemaSlice" -}}
map[string]*schema.Schema{
{{ range . | nodeFields -}}
"{{. | nodeName }}": {
{{- template "schemaSchema" .}}
},
{{ end }}
{{ end }}

{{ define "handleSlice" -}}
{{ $type := .Type.String }}
{{ $nname := nodeName . }}
	{{ $plainType := $type | trimPrefix "[]" }}
	// HERE
	if _, ok := d.GetOk(prefix + "{{$nname}}"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "{{$nname}}"))
		i := 0
		val.{{.Name}} = {{ $type }}{}
		for {
			pfx := fmt.Sprintf("%s.%s.%d", prefix, "{{$nname}}", i)
			if v, ok := d.GetOk(pfx); ok {
				tflog.Debug(ctx, fmt.Sprintf("key is set: %s", pfx))
				val.{{.Name}} = append(val.{{.Name}}, v.({{$plainType}}))
				i++
			} else {
				tflog.Debug(ctx, "not set: %s", pfx)
				break
			}
		}
	}
{{ end }}

package provider

// generated file. DO NOT EDIT!

import (
	"context"
	"fmt"

	"github.com/giantswarm/to"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"gitlab.com/rsrchboy/terraform-provider-gitlabci/third_party/gitlab/runner/config"
	"gitlab.com/rsrchboy/terraform-provider-gitlabci/third_party/gitlab/runner/config/docker"
	"gitlab.com/rsrchboy/terraform-provider-gitlabci/third_party/gitlab/runner/config/referees"
	"gitlab.com/rsrchboy/terraform-provider-gitlabci/third_party/gitlab/runner/config/ssh"
)

var configDataSourceRawSchema = {{ template "schemaSchemaSlice" . }}
}

func dataSourceGitlabCIRunnerConfigReadNEW(d *schema.ResourceData, meta interface{}) error {

	// c := config.Config{}

{{ $pfx := "" }}
{{ range . | nodeFields -}}
// {{ .Name }}: {{. | nodeName }} -- {{ .Type.Name }}, {{ .Type.String }}
// if v, ok := d.GetOk("{{$pfx}}{{nodeName .}}"); ok {
//c.{{.Name}} = v.(FIXME type)
// }
{{ end }}

	return nil
}

{{ define "readStructFunc" }}

func dsRunnerConfigReadStruct{{ .Name | title | replace "." "" }}(ctx context.Context, prefix string, d *schema.ResourceData) ({{ .Name }}, error) {

	if prefix != "" {
		prefix = prefix + "."
	}

	tflog.Debug(ctx, fmt.Sprintf("beginning dsRunnerConfigReadStruct{{ .Name | title | replace "." "" }} run; prefix is '%s'", prefix))

	val := {{ .Name }}{}

{{ range .Type | nodeFields -}}

{{ $type := .Type.String }}
{{ $nname := nodeName . }}
// {{ .Name }}: {{$nname}} -- {{ .Type.Name }}, {{ .Type.String }}
{{ if eq $type "config.StringOrArray" -}}
	if v, ok := d.GetOk(prefix + "{{$nname}}"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "{{$nname}}"))
		val.{{.Name}} = v.([]string)
	}
{{ else if eq "elemSliceString" (. | nodeElemTemplate)}}
{{ template "handleSlice" . }}
{{ else if eq "elemSliceBool" (. | nodeElemTemplate)}}
{{ template "handleSlice" . }}
{{ else if eq "elemSliceInt" (. | nodeElemTemplate)}}
{{ template "handleSlice" . }}
{{ else if .Type | isSimpleType -}}
	if v, ok := d.GetOk(prefix + "{{$nname}}"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "{{$nname}}"))
		val.{{.Name}} = v.({{ .Type.String }})
	}
{{ else if eq "elemStruct" (. | nodeElemTemplate)}}
	if _, ok := d.GetOk(prefix + "{{$nname}}.0"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "{{$nname}}"))
		thing, err := dsRunnerConfigReadStruct{{ $type | title | replace "." "" | replace "*" "" }}(ctx, prefix+"{{$nname}}.0", d)
		if err != nil {
			return val, err
		}
		val.{{.Name}} = thing
	}
{{ else if eq "elemSliceStruct" (. | nodeElemTemplate)}}
	{{ $plainType := $type | trimPrefix "[]" }}
	// HERE
	if _, ok := d.GetOk(prefix + "{{$nname}}"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "{{$nname}}"))
		i := 0
		val.{{.Name}} = {{ $type }}{}
		for {
			pfx := fmt.Sprintf("%s.%s.%d", prefix, "{{$nname}}", i)
			if _, ok := d.GetOk(pfx); ok {
				tflog.Debug(ctx, "set: %s", pfx)
				thing, err := dsRunnerConfigReadStruct{{ $plainType | title | replace "." "" | replace "*" "" }}(ctx, pfx, d)
				if err != nil {
					return val, err
				}
				val.{{.Name}} = append(val.{{.Name}}, thing)
				i++
			} else {
				tflog.Debug(ctx, "not set: %s", pfx)
				break
			}
		}
	}
{{ else if eq "elemSlicePtrStruct" (. | nodeElemTemplate)}}
	{{ $plainType := $type | trimPrefix "[]" }}
	// HERE
	if _, ok := d.GetOk(prefix + "{{$nname}}"); ok {
		tflog.Debug(ctx, fmt.Sprintf("key is set: %s%s", prefix, "{{$nname}}"))
		i := 0
		val.{{.Name}} = {{ $type }}{}
		for {
			pfx := fmt.Sprintf("%s%s.%d", prefix, "{{$nname}}", i)
			if _, ok := d.GetOk(pfx); ok {
				tflog.Debug(ctx, fmt.Sprintf("key is set: %s", pfx))
				thing, err := dsRunnerConfigReadStruct{{ $plainType | title | replace "." "" | replace "*" "" }}(ctx, pfx, d)
				if err != nil {
					return val, err
				}
				val.{{.Name}} = append(val.{{.Name}}, &thing)
				i++
			} else {
				tflog.Debug(ctx, fmt.Sprintf("not set: %s", pfx))
				break
			}
		}
	}
{{ else if eq "elemPtrStruct" (. | nodeElemTemplate)}}
	if _, ok := d.GetOk(prefix + "{{$nname}}.0"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "{{$nname}}"))
		thing, err := dsRunnerConfigReadStruct{{ .Type.String | title | replace "." "" | replace "*" "" }}(ctx, prefix+"{{$nname}}.0", d)
		if err != nil {
			return val, err
		}
		val.{{.Name}} = &thing
	} else {
		tflog.Trace(ctx, fmt.Sprintf("not set: %s", prefix + "{{$nname}}.0"))
	}
{{ else if eq $type "*string" -}}
	if v, ok := d.GetOk(prefix + "{{$nname}}"); ok {
		tflog.Debug(ctx, "set: %s.%s", prefix, "{{$nname}}")
		val.{{.Name}} = to.StringP(v.(string))
	}
{{ else if eq $type "*int" -}}
	if v, ok := d.GetOk(prefix + "{{$nname}}"); ok {
		tflog.Debug(ctx, "set: %s.%s", prefix, "{{$nname}}")
		val.{{.Name}} = to.IntP(v.(int))
	}
{{ else if eq $type "*int64" -}}
	if v, ok := d.GetOk(prefix + "{{$nname}}"); ok {
		tflog.Debug(ctx, "set: %s.%s", prefix, "{{$nname}}")
		val.{{.Name}} = to.Int64P(v.(int64))
	}
{{ else if eq $type "*bool" -}}
	if v, ok := d.GetOk(prefix + "{{$nname}}"); ok {
		tflog.Debug(ctx, "set: %s.%s", prefix, "{{$nname}}")
		val.{{.Name}} = to.BoolP(v.(bool))
	}
{{ else -}}
	// FIXME unhandled type {{ $type }}
{{ end -}}

{{ end }}

	return val, nil
}

{{ end }}


{{ $pfx = "" }}
{{ $all := . | allStructs | sortAlpha | uniq -}}
{{ range $all -}}
{{ $t := . | typeFor }}
// {{ . }}
{{ template "readStructFunc" dict "Name" . "Type" (typeFor .) }}
{{ end }}
`

const outFile = "internal/provider/generated.go"

var funcMap template.FuncMap
var tmpl *template.Template

func init() {

	thing := reflect.TypeOf(config.Config{})

	structsCache := map[string]reflect.Type{}
	for _, t := range childStructTypes(thing) {
		structsCache[t.String()] = t
	}
	funcMap = template.FuncMap{
		"nodeName": attrName,
		"nodeDesc": func(f reflect.StructField) string {
			return f.Tag.Get("description")
		},
		"nodeElemTemplate": nodeElemTemplate,
		// We prefer to deal with embedded fields as transparently as
		// possible.  This should be the least surprising to the consumer of
		// this provider!  (Also, filter anything not exported -- though we're
		// not allowing any of those in our generated structs at this
		// point...)
		"nodeFields": func(t reflect.Type) []reflect.StructField {
			fields := []reflect.StructField{}
			for _, f := range reflect.VisibleFields(t) {
				if f.Anonymous || !f.IsExported() {
					continue
				}
				fields = append(fields, f)
			}
			return fields
		},

		"allStructs": childStructs,
		"typeFor": func(name string) (reflect.Type, error) {
			if t, ok := structsCache[name]; ok {
				return t, nil
			}
			return nil, errors.New("unknown type name: " + name)
		},
		"isSimpleType": func(t reflect.Type) bool {
			switch t.String() {
			case "string", "int", "int32", "int64", "float", "float64", "bool":
				return true
			case "[]string", "[]int64":
				return false
			case "map[string]string", "map[string]bool":
				return true
			default:
				return false
			}
		},
	}
	tmpl = template.Must(
		template.
			New("fields").
			Funcs(sprig.TxtFuncMap()).
			Funcs(funcMap).
			Parse(tmplString),
	)
}

func main() {

	thing := reflect.TypeOf(config.Config{})

	var out bytes.Buffer
	err := tmpl.Execute(&out, thing)
	if err != nil {
		fmt.Printf("template error!: %s\n", err)
	}

	// gofmt
	source, err := format.Source(out.Bytes())
	if err != nil {
		// log.Println(out.String())
		fmt.Println(out.String())
		fmt.Printf("failed to format sources: %s\n", err)
		return
	}

	// goimports
	process, err := imports.Process(outFile, source, nil)
	if err != nil {
		log.Println(string(source))
		fmt.Printf("failed to format imports: %s\b", err)
		return
	}

	fmt.Println(string(process))

	// TODO this is begging for a flag
	err = os.WriteFile(outFile, process, 0o644)
	if err != nil {
		fmt.Printf("failed to write file: %s\n", err)
		return
	}

}

func attrName(f reflect.StructField) string {
	name := f.Name
	if tag := f.Tag.Get("toml"); tag != "" {
		name = strings.Split(tag, ",")[0]
	}
	return strcase.SnakeCase(name)
}

// return the name of the template to process this field with for a
// schema.Schema's Elem field
func nodeElemTemplate(f reflect.StructField) string {

	name := nodeElemTemplatePart(f.Type)

	switch name {
	case "PtrBool":
		return "elemBool"
	case "PtrString":
		return "elemString"
	case "PtrFloat":
		return "elemFloat"
	case "PtrInt":
		return "elemInt"
	default:
		return "elem" + name
	}

}

// return the name of the template to process this field with for a
// schema.Schema's Elem field
func nodeElemTemplatePart(t reflect.Type) string {
	tname := t.Kind().String()
	switch tname {
	case "ptr":
		return "Ptr" + nodeElemTemplatePart(t.Elem())
	case "slice":
		return "Slice" + nodeElemTemplatePart(t.Elem())
	case "map":
		return "Map" + nodeElemTemplatePart(t.Elem())
	case "struct":
		return "Struct"
	case "string", "*string": //, "common.DockerPullPolicy", "common.KubernetesPullPolicy":
		return "String"
	case "*int", "*int64", "int", "int32", "int64":
		return "Int"
	case "*float", "*float64", "float", "float64":
		return "Float"
	case "*bool", "bool":
		return "Bool"
	default:
		return tname
	}
}

func fieldToSchema(f reflect.StructField) *schema.Schema {

	me := &schema.Schema{}

	if tag := f.Tag.Get("description"); tag != "" {
		me.Description = tag
	}

	tname := f.Type.Kind().String()
	switch tname {
	case "[]string", "*[]string", "[]*string":
		me.Type = schema.TypeList
		me.Elem = &schema.Schema{Type: schema.TypeString}
		// info.IsList = true
	case "[]*int", "[]*int64", "[]int", "[]int64":
		me.Type = schema.TypeList
		me.Elem = &schema.Schema{Type: schema.TypeInt}
		// info.IsList = true
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
		fmt.Printf("unhandled type: %s\n", tname)
	}

	return me
}

// return a list of all the structs the given struct has as fields, as well as
// their children
func childStructTypes(t reflect.Type) []reflect.Type {
	switch t.Kind() {
	case reflect.Ptr, reflect.Slice, reflect.Map:
		return childStructTypes(t.Elem())
	case reflect.Struct:
		list := []reflect.Type{t}
		for _, f := range reflect.VisibleFields(t) {
			list = append(list, childStructTypes(f.Type)...)
		}
		return list
	default:
		return []reflect.Type{}
	}
}

func childStructs(t reflect.Type) []string {
	structs := childStructTypes(t)
	names := make([]string, len(structs))
	for i, s := range structs {
		names[i] = s.String()
	}
	return names
}
