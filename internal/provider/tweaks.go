package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// Tweak our generated schema for the runner_config data source.

func init() {

	// validation, etc
	configDataSourceRawSchema["concurrent"].ValidateFunc = validation.IntAtLeast(1)
	configDataSourceRawSchema["check_interval"].ValidateFunc = validation.IntAtLeast(0)
	configDataSourceRawSchema["log_level"].ValidateFunc =
		validation.StringInSlice([]string{"panic", "fatal", "error", "warning", "info", "debug"}, false)
	configDataSourceRawSchema["log_format"].ValidateFunc =
		validation.StringInSlice([]string{"runner", "text", "json"}, false)

	// computed attributes
	configDataSourceRawSchema["config"] = &schema.Schema{
		Type:        schema.TypeString,
		Computed:    true,
		Description: "The computed runner configuration (toml)",
	}
	configDataSourceRawSchema["id"] = &schema.Schema{
		Type:        schema.TypeString,
		Computed:    true,
		Description: "The computed configuation id",
	}

	// allow for some other config to be used as a template
	configDataSourceRawSchema["runners"].Elem.(*schema.Resource).Schema["config_template"] = &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Configuration template (toml)",
	}

	// FIXME make sure this flows out into our generated toml output iff set
	configDataSourceRawSchema["runners"].Elem.(*schema.Resource).
		Schema["runner_credentials"].Elem.(*schema.Resource).
		Schema["token"].Sensitive = true
}
