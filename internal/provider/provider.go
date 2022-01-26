package provider

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// NewProvider creates a new Provider
func NewProvider(version string) func() *schema.Provider {
	return func() *schema.Provider {
		p := &schema.Provider{
			Schema: map[string]*schema.Schema{
				"base_url": {
					Type:         schema.TypeString,
					Optional:     true,
					DefaultFunc:  schema.EnvDefaultFunc("GITLAB_BASE_URL", "https://gitlab.com/api/v4"),
					Description:  "The GitLab base API URL",
					ValidateFunc: validateURLFunc,
				},
			},
			DataSourcesMap: map[string]*schema.Resource{
				"gitlabci_runner_config": dataSourceGitlabCIRunnerConfig(),
				"gitlabci_environment":   dataSourceGitlabCIEnvironment(),
			},
			ResourcesMap: map[string]*schema.Resource{
				"gitlabci_runner_token": resourceGitlabRunner(),
			},
		}

		p.ConfigureContextFunc = providerConfigure(version, p)
		return p
	}
}

var validateURLFunc = func(v interface{}, k string) (s []string, errors []error) {
	value := v.(string)
	url, err := url.Parse(value)

	if err != nil || url.Host == "" || url.Scheme == "" {
		errors = append(errors, fmt.Errorf("%s is not a valid URL", value))
		return
	}

	return
}

func providerConfigure(version string, p *schema.Provider) func(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		baseURL := strings.TrimRight(d.Get("base_url").(string), "/")
		return baseURL, nil
	}
}
