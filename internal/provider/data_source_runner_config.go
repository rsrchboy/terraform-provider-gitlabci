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

	// structs.DefaultTagName = "toml"

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

	// FIXME need to restore template functionality
	// if runners, hasRunners := d.GetOk("runners"); hasRunners {
	// 	c.Runners = make([]*rcommon.RunnerConfig, len(runners.([]interface{})))
	// 	for i, _ := range runners.([]interface{}) {
	// 		log.Printf("working on runner block #%d", i)

	// 		// handle templates here
	// 		if templateData, hasTemplateData := d.GetOkExists(pfx + "config_template"); hasTemplateData {
	// 			// https://godoc.org/github.com/hashicorp/terraform-plugin-sdk/helper/pathorcontents
	// 			templateStr, _, err := pathorcontents.Read(templateData.(string))
	// 			if err != nil {
	// 				return err
	// 			}
	// 			tmpfile, err := ioutil.TempFile("", "runnercfgtemplate*")
	// 			if err != nil {
	// 				return err
	// 			}
	// 			defer os.Remove(tmpfile.Name())
	// 			if _, err = tmpfile.WriteString(templateStr); err != nil {
	// 				return err
	// 			}
	// 			ct, err := configtemplate.NewConfigTemplateFromFile(tmpfile.Name())
	// 			if err != nil {
	// 				return err
	// 			}
	// 			ct.MergeTo(&r)
	// 		}

	// 		c.Runners[i] = &r
	// 	}

	// generate toml config
	buf := new(bytes.Buffer)
	if err := toml.NewEncoder(buf).Encode(c); err != nil {
		return diag.FromErr(err)
	}

	d.Set("config", fmt.Sprintf("%s", buf.String()))
	d.Set("config_not_sensitive", fmt.Sprintf("%s", buf.String()))
	tflog.Debug(ctx, "runner config toml:\n\n%s", buf.String())

	configString := buf.String()

	d.SetId(strconv.Itoa(hashcode.String(configString)))

	tflog.Trace(ctx, "dataSourceGitlabCIRunnerConfigRead() finished")
	return nil
}
