package provider

import (
	"bytes"
	"context"
	"fmt"
	"strconv"

	"github.com/BurntSushi/toml"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/imdario/mergo"
	"gitlab.com/rsrchboy/terraform-provider-gitlabci/third_party/gitlab/runner/config"
)

type schemaMap map[string]*schema.Schema
type iMap map[string]interface{}

type stringMap map[string]string

const dsRunnerConfigDescription = `
This data source can be used to generate a valid TOML configuration for a
gitlab-runner.  The blocks and attributes largely follow the configuration and
naming of the configuration file, with the exception of the computed
attributes (e.g. "output") and inputs for templating.
`

func dataSourceGitlabCIRunnerConfig() *schema.Resource {

	schema := &schema.Resource{
		Description: dsRunnerConfigDescription,
		ReadContext: dataSourceGitlabCIRunnerConfigRead,
		Schema:      configDataSourceRawSchema,
	}

	return schema
}

func dataSourceGitlabCIRunnerConfigRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	tflog.Trace(ctx, "dataSourceGitlabCIRunnerConfigRead() (mark IV)")

	c, err := dsRunnerConfigReadStructConfigConfig(ctx, "", d)
	if err != nil {
		return diag.FromErr(err)
	}

	// process runner config templates (if any)
	for i, r := range c.Runners {
		key := fmt.Sprintf("runners.%d.config_template", i)
		if tstr, ok := d.GetOk(key); ok {
			tflog.Debug(ctx, fmt.Sprintf("template exists for key: %s", key))
			tc := config.Config{}
			if _, err = toml.Decode(tstr.(string), &tc); err != nil {
				return diag.FromErr(err)
			}
			if count := len(tc.Runners); count != 1 {
				return diag.Errorf("template %s has %d != 1 runners sections!", key, count)
			}
			// err = mergo.Merge(c.Runners[i], tc.Runners[0])
			err = mergo.Merge(r, *tc.Runners[0])
			if err != nil {
				return diag.FromErr(err)
			}
		} else {
			tflog.Debug(ctx, fmt.Sprintf("template does not exist for key: %s", key))
		}
	}

	// generate toml config
	buf := new(bytes.Buffer)
	if err := toml.NewEncoder(buf).Encode(c); err != nil {
		return diag.FromErr(err)
	}

	d.Set("config", fmt.Sprintf("%s", buf.String()))
	d.Set("config_not_sensitive", fmt.Sprintf("%s", buf.String()))

	// TODO how concerned should we be about this logging? (from a secrets
	// perspective)
	tflog.Debug(ctx, fmt.Sprintf("runner config toml:\n\n%s", buf.String()))

	configString := buf.String()

	d.SetId(strconv.Itoa(hashcode.String(configString)))

	tflog.Trace(ctx, "dataSourceGitlabCIRunnerConfigRead() finished")
	return nil
}
