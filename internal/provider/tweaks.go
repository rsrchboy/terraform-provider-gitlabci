package provider

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	v "github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// Tweak our generated schema for the runner_config data source.
//
// This file contains any changes that we need to make to the schema, etc,
// prior to its consumption/usage.  These sorts of changes would be ...
// awkward to embed in the code generation, but reasonably straight-forward to
// handle here.
//
// We're not going to try to include defaults.  (Most of the time, anyways.)
// There's too many ways someone could reasonably use this datasource that
// defaults could make a mess of things.  Merge requests demonstrating the
// error of my ways are welcome :)
//
// Generally speaking, one can expect to find:
//
// *) computed attributes
//
//    e.g. config, or id
//
// *) attribute, um, attribute tweaks
//
//    Secrets should be marked "Sensitive", some blocks have a maximum, etc.
//
// *) validation functions
//
//    Some attributes we can clearly validate the inputs; others... not so
//    much.  We'll try to cover the basics/obvious ones; merge requests
//    welcome :)

func init() {

	// computed attributes
	configDataSourceRawSchema["config"] = &schema.Schema{
		Type:        schema.TypeString,
		Computed:    true,
		Sensitive:   true,
		Description: "The computed runner configuration (TOML).  This attribute is marked sensitive as it may include authentication tokens, cache bucket keys, etc.",
	}
	configDataSourceRawSchema["config_not_sensitive"] = &schema.Schema{
		Type:        schema.TypeString,
		Computed:    true,
		Description: "The computed runner configuration (TOML).  This attribute is NOT marked sensitive EVEN THOUGH it may include authentication tokens, cache bucket keys, etc.  If you're uncertain, you should probably just be using the `config` attribute.",
	}
	configDataSourceRawSchema["id"] = &schema.Schema{
		Type:        schema.TypeString,
		Computed:    true,
		Description: "The computed configuration id",
	}

	// // allow for some other config to be used as a template
	// configDataSourceRawSchema["runners"].Elem.(*schema.Resource).Schema["config_template"] = &schema.Schema{
	// 	Type:        schema.TypeString,
	// 	Optional:    true,
	// 	Sensitive:   true,
	// 	Description: "Configuration template (toml).  If included, this toml will be used as the base of this runner's configuration.  This attribute is marked sensitive as it may include authentication tokens, cache bucket keys, etc.",
	// }

	// secrets!
	findRes("runners", "token").Sensitive = true
	findRes("runners", "cache", "s3", "secret_key").Sensitive = true
	findRes("runners", "cache", "gcs", "private_key").Sensitive = true
	findRes("runners", "cache", "azure", "account_key").Sensitive = true
	findRes("runners", "kubernetes", "bearer_token").Sensitive = true
	findRes("runners", "ssh", "password").Sensitive = true

	findRes("runners", "docker", "services", "name").Required = true
	findRes("runners", "docker", "services", "name").Optional = false

	// validations
	findRes("concurrent").ValidateFunc = v.IntAtLeast(1)
	findRes("check_interval").ValidateFunc = v.IntAtLeast(0)
	findRes("log_level").ValidateFunc =
		v.StringInSlice([]string{"panic", "fatal", "error", "warning", "info", "debug"}, false)
	findRes("log_format").ValidateFunc =
		v.StringInSlice([]string{"runner", "text", "json"}, false)
	findRes("runners", "url").ValidateFunc = v.IsURLWithHTTPorHTTPS
	findRes("runners", "clone_url").ValidateFunc = v.IsURLWithHTTPorHTTPS
	findRes("runners", "executor").ValidateFunc =
		v.StringInSlice([]string{
			"custom",
			"docker",
			"docker+machine",
			"docker-ssh",
			"docker-ssh+machine",
			"docker-windows",
			"kubernetes",
			"parallels",
			"shell",
			"ssh",
			"virtualbox",
		}, false)
	findRes("runners", "shell").ValidateFunc =
		v.StringInSlice([]string{"bash", "sh", "powershell", "pwsh"}, false)
	findRes("runners", "machine", "machine_name").ValidateFunc =
		v.StringMatch(regexp.MustCompile(`.*%s.*`), "string must include %s")
}

func findRes(path ...string) *schema.Schema {

	return findResRec(configDataSourceRawSchema[path[0]], path[1:]...)
}

func findResRec(s *schema.Schema, path ...string) *schema.Schema {
	switch len(path) {
	case 0:
		return s
	case 1:
		return s.Elem.(*schema.Resource).Schema[path[0]]
	default:
		found := s.Elem.(*schema.Resource).Schema[path[0]]
		return findResRec(found, path[1:]...)
	}
}

// vim: set nowrap :
