package provider

import (
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceRunnerToken(t *testing.T) {

	token := os.Getenv("RUNNER_REGISTRATION_TOKEN")

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t); testRunnerTokenAccPreCheck(t, token) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testAccResourceRunnerToken, token),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("gitlabci_runner_token.foo", "token"),
					resource.TestCheckResourceAttrSet("gitlabci_runner_token.foo", "runner_id"),
					resource.TestMatchResourceAttr("gitlabci_runner_token.foo", "runner_id", regexp.MustCompile(`^\d+$`)),
				),
			},
		},
	})
}

const testAccResourceRunnerToken = `
resource "gitlabci_runner_token" "foo" {
  registration_token = "%s"
}
`

func testRunnerTokenAccPreCheck(t *testing.T, token string) {
	if token == "" {
		t.Fatal("$RUNNER_REGISTRATION_TOKEN not set; skipping registration tests")
	}
}
