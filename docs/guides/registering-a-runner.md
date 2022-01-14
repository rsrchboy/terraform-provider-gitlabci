---
layout: ""
page_title: "Registering a runner"
subcategory: ""
description: |-
  Registering a runner with a registration token.
---

# Registering a runner

Registering a GitLab Runner relies on a "registration token".  This token is
not tied to any specific user or group, but rather to the entity the runner is
being registered to: a project, group, or instance.  A user does not "own" a
runner, therefore user credentials are not required to register it.

If you haven't yet read it, the [Registering Runners](https://docs.gitlab.com/runner/register/)
page is useful.

## Our tokens are immutable

As we're registering tokens using a runner registration token, we only have
access to create / destroy runners.  Changing the `runner_token` resource
definition will force the destruction and recreation of the resource.

...except for the `runner_token.registration_token` argument

## Instance-level runners

...are only possible if you're registering runners against your own instance
of GitLab (or you work for GitLab :)).  Support for obtaining the
instance-level registration token is, AFAIK, not included in the
general-purpose GitLab provider, and as such will need to be provided
manually.

# Obtaining a registration token

You can obtain tokens either manually, or through the general-purpose
[GitLab provider](https://www.terraform.io/docs/providers/gitlab/).  Note that
`terraform-provider-gitlab` does require user authentication, etc.

* data sources
    * [`data.gitlab.project.XXX.runners_token`](https://registry.terraform.io/providers/gitlabhq/gitlab/latest/docs/data-sources/project#runners_token)
    * [`data.gitlab.user.XXX.runners_token`](https://registry.terraform.io/providers/gitlabhq/gitlab/latest/docs/data-sources/group#runners_token)
* resources
    * [`gitlab.project.XXX.runners_token`](https://registry.terraform.io/providers/gitlabhq/gitlab/latest/docs/resources/project#runners_token)
    * [`gitlab.user.XXX.runners_token`](https://registry.terraform.io/providers/gitlabhq/gitlab/latest/docs/resources/group#runners_token)

# Manual runner registration

If, for whatever reason, you want to do this manually, the procedure is pretty
obvious.  Note that this is (currently) the only way to do this when
registering runners at the instance level.

The [Registering Runners](https://docs.gitlab.com/runner/register/) contains
more information about finding these registration tokens.

```terraform
resource "gitlabci_runner_token" "this" {
  registration_token = var.registration_token
  locked             = true
  tags = [
    "jinx",
    "powder",
    "cupcake",
  ]
}

variable "registration_token" {
  type        = string
  description = "Runner registration token"
  sensitive   = true
}

output "token" {
  sensitive = true
  value     = gitlabci_runner_token.this.token
}
```

# Registration token lookup

The use of project and group resource / data-source lookups via the
general-purpose GitLab provider are extremely similar, so we'll just go over
the use of the project lookups.

## Via data-sources (e.g. `data.gitlab_project`)

```terraform
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
```

## Via resources (e.g. `gitlab_project`)

```terraform
resource "gitlab_project" "this" {
  name            = "arcane"
  description     = "No magic, plz!"
  visibility_leve = "public"
}

resource "gitlabci_runner_token" "this" {
  registration_token = gitlab_project.this.runners_token
  locked             = true
  tags = [
    "jinx",
    "powder",
    "cupcake",
  ]
}

output "token" {
  sensitive = true
  value     = gitlabci_runner_token.this.token
}
```
