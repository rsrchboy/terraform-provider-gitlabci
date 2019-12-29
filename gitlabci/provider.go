package gitlabci

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"base_url": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "The GitLab Base API URL",
				ValidateFunc: validateURLFunc,
			},
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
