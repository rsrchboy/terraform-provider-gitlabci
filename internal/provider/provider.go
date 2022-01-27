package provider

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/parnurzeal/gorequest"
)

// NewProvider creates a new Provider
func NewProvider(version, commit string) func() *schema.Provider {
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

		p.ConfigureContextFunc = providerConfigure(version, commit, p)
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

// Basic information best set at the provider level
type apiClient struct {
	baseURL   string
	userAgent string
	version   string
	commit    string
}

func (api apiClient) newAgent() *gorequest.SuperAgent {
	return gorequest.New().Set("User-Agent", api.userAgent)
}

func providerConfigure(version, commit string, p *schema.Provider) func(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		// Setup a User-Agent for the API client
		uaVersion := version
		if commit != "" {
			// version + "-" + commit
			uaVersion += "-" + commit
		}

		i := apiClient{
			baseURL:   strings.TrimRight(d.Get("base_url").(string), "/"),
			userAgent: p.UserAgent("terraform-provider-gitlabci", uaVersion),
			version:   version,
			commit:    commit,
		}

		return i, nil
	}
}
