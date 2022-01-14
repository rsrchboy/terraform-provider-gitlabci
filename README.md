
# terraform-provider-gitlabci

Given a registration token, register a runner with GitLab.

# Synopsis

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

# Description

The [GitLab provider for terraform](https://github.com/terraform-providers/terraform-provider-gitlab) is rather nice.  However, it (currently) has
a couple limitations:

* Runners cannot be registered; and
* API tokens are required.

This is a limited functionality provider, aimed only at making it trivial to
create / destroy registered runner tokens while requiring nothing more than
the relevant [registration token](https://docs.gitlab.com/ce/api/runners.html#registration-and-authentication-tokens).

# Documentation

This module is published to the [public terraform registry](https://registry.terraform.io).
Please see the documentation there:

* [`terraform-provider-gitlabci` documentation](https://registry.terraform.io/providers/rsrchboy/gitlabci/latest/docs)

# Author and Copyright

This software is Copyright 2019-2022 Chris Weyl <cweyl@alumni.drew.edu>.

This is free software, licensed under the GNU GPL v3+.  See the `LICENSE` file
for more information.
