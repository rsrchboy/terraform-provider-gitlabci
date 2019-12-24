package gitlabci

import (
	"fmt"
	"net/url"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	// "github.com/parnurzeal/gorequest"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"base_url": {
				Type: schema.TypeString,
				// Optional:     true,
				Required: true,
				// DefaultFunc:  schema.EnvDefaultFunc("GITLAB_BASE_URL", ""),
				Description: descriptions["base_url"],
				// FIXME schema url validation func exists?
				ValidateFunc: validateURLFunc,
			},
			// "cacert_file": {
			// 	Type:        schema.TypeString,
			// 	Optional:    true,
			// 	Default:     "",
			// 	Description: descriptions["cacert_file"],
			// },
			// "insecure": {
			// 	Type:        schema.TypeBool,
			// 	Optional:    true,
			// 	Default:     false,
			// 	Description: descriptions["insecure"],
			// },
		},
		ResourcesMap: map[string]*schema.Resource{
			"gitlabci_runner_token": resourceGitlabRunner(),
		},
		ConfigureFunc: providerConfigure,
	}
}

var descriptions map[string]string

var validateURLFunc = func(v interface{}, k string) (s []string, errors []error) {
	value := v.(string)
	url, err := url.Parse(value)

	if err != nil || url.Host == "" || url.Scheme == "" {
		errors = append(errors, fmt.Errorf("%s is not a valid URL", value))
		return
	}

	return
}

func init() {
	descriptions = map[string]string{
		// "token": "The OAuth token used to connect to GitLab.",

		"base_url": "The GitLab Base API URL",

		"cacert_file": "A file containing the ca certificate to use in case ssl certificate is not from a standard chain",

		"insecure": "Disable SSL verification of API calls",
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	// 	// config := Config{
	// 	// 	// Token:      d.Get("token").(string),
	// 	// 	BaseURL:    d.Get("base_url").(string),
	// 	// 	CACertFile: d.Get("cacert_file").(string),
	// 	// 	Insecure:   d.Get("insecure").(bool),
	// 	// }
	// 	// return config.Client()
	baseURL := d.Get("base_url").(string) + ""
	// return d.Get("base_url").(interface{}), nil
	return baseURL, nil
}
