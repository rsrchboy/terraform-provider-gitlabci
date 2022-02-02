package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceRunnerConfig(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceRunnerConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.gitlabci_runner_config.foo", "log_format", "json",
					),
					resource.TestCheckResourceAttr(
						"data.gitlabci_runner_config.foo", "config", testAccDataSourceRunnerConfigOutput,
					),
				),
			},
		},
	})
}

const testAccDataSourceRunnerConfig = `
data "gitlabci_runner_config" "foo" {
  log_format = "json"
  runners {
    limit = 1234
    url = "https://a.gitlab.somewhere"
    clone_url = "https://a.sneaky.gitlab.somewhere"
    builds_dir = "yahhhhhhhhhhhhhhhhhhhhhhh!!!!!!!!!!!!!!!!!!!!!!"
    executor = "docker"
    docker {
       image = "an/image:here"
    }
  }
}
`

const testAccDataSourceRunnerConfigOutput = `concurrent = 0
check_interval = 0
log_format = "json"

[session_server]
  session_timeout = 0

[[runners]]
  name = ""
  limit = 1234
  url = "https://a.gitlab.somewhere"
  token = ""
  executor = "docker"
  builds_dir = "yahhhhhhhhhhhhhhhhhhhhhhh!!!!!!!!!!!!!!!!!!!!!!"
  clone_url = "https://a.sneaky.gitlab.somewhere"
  [runners.referees]
  [runners.docker]
    tls_verify = false
    image = "an/image:here"
    privileged = false
    disable_entrypoint_overwrite = false
    oom_kill_disable = false
    disable_cache = false
    shm_size = 0
`
