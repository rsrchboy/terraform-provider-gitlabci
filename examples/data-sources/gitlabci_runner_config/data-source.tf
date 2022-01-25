data "gitlab_project" "this" {
  id = var.project_id
}

resource "gitlabci_runner_token" "this" {
  registration_token = data.gitlab_project.this.runners_token
  locked             = true
  tags               = ["jinx", "powder", "cupcake"]
}

data "gitlabci_runner_config" "this" {
  log_format = "json"
  runners {
    url   = "https://gitlab.com"
    token = gitlabci_runner_token.this.token
    # ...
  }
}

resource "local_file" "this" {
  filename = "${path.module}/config.toml"
  content  = data.gitlabci_runner_config.this.config
}

output "config" {
  description = "Our generated runner configuration"
  value       = data.gitlabci_runner_config.this.config
}

variable "project_id" {
  # note this could also legitimately be an int; see
  # https://registry.terraform.io/providers/gitlabhq/gitlab/latest/docs/data-sources/project#id
  type        = string
  description = "Project ID (path, e.g. rsrchboy/scratch)"
}

