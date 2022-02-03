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

	resource.UnitTest(t, resource.TestCase{
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceRunnerConfigPullPolicyValidation,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.gitlabci_runner_config.foo", "config", testAccDataSourceRunnerConfigPullPolicyValidationOutput,
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

// A note, here, if you're wondering why we're seeing thing like
// runners.docker.tls_verify = false even when it was never set: the toml
// library being used does not attempt to distinguish between "false" and
// "zero, was never explicitly set", see
// https://github.com/BurntSushi/toml/blob/4272474656f1b35414cdee32185a45e36b39246e/encode.go#L611
//
// This isn't something to "blame" anything for; AFAICT it's not possible to
// make that determination at all (hence the library not even trying).  This
// behaviour is also in line with how gitlab-runner handles it, so I'm not
// going to worry about it until such time as gitlab-runner does.

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

const testAccDataSourceRunnerConfigPullPolicyValidation = `
data "gitlabci_runner_config" "foo" {
  runners {
    docker {
      pull_policy = ["never"]
    }
  }
}
`

const testAccDataSourceRunnerConfigPullPolicyValidationOutput = `concurrent = 0
check_interval = 0

[session_server]
  session_timeout = 0

[[runners]]
  name = ""
  url = ""
  token = ""
  executor = ""
  [runners.referees]
  [runners.docker]
    tls_verify = false
    image = ""
    privileged = false
    disable_entrypoint_overwrite = false
    oom_kill_disable = false
    disable_cache = false
    pull_policy = ["never"]
    shm_size = 0
`
