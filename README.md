
# terraform-provider-gitlabci

Given a registration token, register a runner with GitLab.

# Synopsis

```terraform
terraform {
    required_providers {
        gitlabci = {
            source = "registry.terraform.io/rsrchboy/gitlabci"
        }
    }
}

provider "gitlabci" {
    base_url = "https://gitlab.com/api/v4"
}

resource "gitlabci_runner_token" "this" {
    registration_token = "..."
    run_untagged       = true
    active             = true
    locked             = true
    tags = [
        "jinx",
        "powder",
        "cupcake",
    ]
}
```

# Description

The [GitLab provider for terraform](https://github.com/terraform-providers/terraform-provider-gitlab) is rather nice.  However, it (currently) has
a couple limitations:

* Runners cannot be registered; and
* API tokens are required.

This is a limited functionality provider, aimed only at making it trivial to
create / destroy registered runner tokens while requiring nothing more than
the relevant [registration token](https://docs.gitlab.com/ce/api/runners.html#registration-and-authentication-tokens).

This provider also provides data sources related to GitLab CI, including
[`gitlabci_runner_config`](data-sources/runner_config), a data source allowing
runner configuration to be generated (much as, say, the AWS provider's
`aws_iam_policy_document` does for IAM policies).

# Documentation

This module is published to the [public terraform registry](https://registry.terraform.io).
Please see the documentation there:

* [`terraform-provider-gitlabci` documentation](https://registry.terraform.io/providers/rsrchboy/gitlabci/latest/docs)

# Releases and Source

The primary home for this software is on GitLab.  However, AFAICT the
[terraform registry only supports GitHub releases](https://www.terraform.io/registry/providers/publishing#creating-a-github-release)
we also mirror-push to a repository on GitHub.  GitHub Actions are used to
build releases over there, and those releases are then imported by the
[terraform registry](https://registry.terraform.io/providers/rsrchboy/gitlabci/latest).
GitLab CI is used for everything else -- including building releases here, as
well.

* Home: https://gitlab.com/rsrchboy/terraform-provider-gitlabci
* Mirror: https://github.com/rsrchboy/terraform-provider-gitlabci
* Registry: https://registry.terraform.io/providers/rsrchboy/gitlabci/latest

# Author and Copyright

This software is Copyright 2019-2022 Chris Weyl <cweyl@alumni.drew.edu>.

This is free software, licensed under the GNU GPL v3+.  See the `LICENSE` file
for more information.
