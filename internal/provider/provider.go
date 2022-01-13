package provider

import (
	"fmt"
	"net/url"
	"strings"

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
			"gitlabci_environment": dataSourceGitlabCIEnvironment(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"gitlabci_runner_token": resourceGitlabRunner(),
		},
		ConfigureFunc: providerConfigure,
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

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	baseURL := strings.TrimRight(d.Get("base_url").(string), "/")
	return baseURL, nil
}
