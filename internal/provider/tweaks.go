package provider

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
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

	// FIXME make sure this flows out into our generated toml output iff set
	configDataSourceRawSchema["runners"].Elem.(*schema.Resource).
		Schema["token"].Sensitive = true
	configDataSourceRawSchema["runners"].Elem.(*schema.Resource).
		Schema["cache"].Elem.(*schema.Resource).
		Schema["s3"].Elem.(*schema.Resource).
		Schema["secret_key"].Sensitive = true
	configDataSourceRawSchema["runners"].Elem.(*schema.Resource).
		Schema["cache"].Elem.(*schema.Resource).
		Schema["gcs"].Elem.(*schema.Resource).
		Schema["private_key"].Sensitive = true
	configDataSourceRawSchema["runners"].Elem.(*schema.Resource).
		Schema["cache"].Elem.(*schema.Resource).
		Schema["azure"].Elem.(*schema.Resource).
		Schema["account_key"].Sensitive = true
	configDataSourceRawSchema["runners"].Elem.(*schema.Resource).
		Schema["kubernetes"].Elem.(*schema.Resource).
		Schema["bearer_token"].Sensitive = true
	configDataSourceRawSchema["runners"].Elem.(*schema.Resource).
		Schema["ssh"].Elem.(*schema.Resource).
		Schema["password"].Sensitive = true

	// validations
	configDataSourceRawSchema["concurrent"].ValidateFunc = validation.IntAtLeast(1)
	configDataSourceRawSchema["check_interval"].ValidateFunc = validation.IntAtLeast(0)
	configDataSourceRawSchema["log_level"].ValidateFunc =
		validation.StringInSlice([]string{"panic", "fatal", "error", "warning", "info", "debug"}, false)
	configDataSourceRawSchema["log_format"].ValidateFunc =
		validation.StringInSlice([]string{"runner", "text", "json"}, false)
	configDataSourceRawSchema["runners"].Elem.(*schema.Resource).
		Schema["url"].ValidateFunc = validation.IsURLWithHTTPorHTTPS
	configDataSourceRawSchema["runners"].Elem.(*schema.Resource).
		Schema["clone_url"].ValidateFunc = validation.IsURLWithHTTPorHTTPS
	configDataSourceRawSchema["runners"].Elem.(*schema.Resource).
		Schema["executor"].ValidateFunc =
		validation.StringInSlice([]string{
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
	configDataSourceRawSchema["runners"].Elem.(*schema.Resource).
		Schema["shell"].ValidateFunc =
		validation.StringInSlice([]string{"bash", "sh", "powershell", "pwsh"}, false)
	configDataSourceRawSchema["runners"].Elem.(*schema.Resource).
		Schema["machine"].Elem.(*schema.Resource).
		Schema["machine_name"].ValidateFunc =
		validation.StringMatch(regexp.MustCompile(`.*%s.*`), "string must include %s")
}

// vim: set nowrap :
