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
  # note this could also legitimately be an int; see
  # https://registry.terraform.io/providers/gitlabhq/gitlab/latest/docs/data-sources/project#id
  type        = string
  description = "Project ID (path, e.g. rsrchboy/scratch)"
}

output "token" {
  sensitive = true
  value     = gitlabci_runner_token.this.token
}
