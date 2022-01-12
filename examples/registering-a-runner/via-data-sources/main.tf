data "gitlab_project" "this" {
  id = var.project_id
}

resource "gitlabci_runner_token" "this" {
  registration_token = data.gitlab_project.this.runners_token
  locked             = true
  tags = [
    "jinx",
    "powder",
    "cupcake",
  ]
}

variable "project_id" {
  # note this could also legitimately be a string; see
  # https://registry.terraform.io/providers/gitlabhq/gitlab/latest/docs/data-sources/project#id
  type        = number
  description = "Project ID"
}

output "token" {
  sensitive = true
  value     = gitlabci_runner_token.this.token
}
