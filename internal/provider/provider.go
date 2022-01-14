package provider

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func New(version string) func() *schema.Provider {
	return Provider
}

func Provider() *schema.Provider {
	return &schema.Provider{
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
		ConfigureContextFunc: providerConfigure,
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

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	baseURL := strings.TrimRight(d.Get("base_url").(string), "/")
	return baseURL, nil
}
