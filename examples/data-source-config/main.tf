provider "gitlabci" {
  base_url = "https://gitlab.com/api/v4"
}

data "gitlabci_runner_config" "this" {
  log_format = "json"
  session_server {
    listen_address = "0.0.0.0:12345"
  }
}

resource "local_file" "this" {
  filename = "${path.module}/config.toml"
  content  = data.gitlabci_runner_config.this.config
}

output "config" {
  value = data.gitlabci_runner_config.this.config
}
