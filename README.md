# terraform-provider-gitlabci

Given a registration token, register a runner with GitLab.

# Synopsis

    provider "gitlabci" {
        base_url = "https://gitlab.com/api/v4"
    }

    resource "gitlabci_runner_token" "this" {
        registration_token = "c0ffee..."
        run_untagged       = true
        active             = true
        locked             = true
        tags               = [
            "one",
            "two",
            "yipeeeeee",
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

# Provider Configuration

The provider takes only one argument, and it is required.

| Name       | Description                     | Type | Default | Required |
|------------|---------------------------------|:----:|:-------:|:--------:|
| `base_url` | API url of your GitLab instance | URL  | _none_  | Yes      |

# Resource Configuration

## `gitlabci_runner_token`

This resource will take a registration token and use it to register a new
runner.  Tags, etc, may be specified here at create time.

**N.B.** Changing any parameter will force the creation of a new resource.
_Registration info cannot be changed by this resource._

Generally, all options are as listed at [API Doc -- "Register a new Runner"](https://docs.gitlab.com/ce/api/runners.html#register-a-new-runner).

### Configuration Options

| Name                 | Description                      | Type          | Default         | Required |
|----------------------|----------------------------------|:-------------:|:---------------:|:--------:|
| `registration_token` | Registration token               | `string`      | _none_          | Yes      |
| `description`        | Description of runner            | `string`      | _none_          | No       |
| `tags`               | List of tags for the runner      | `set(string)` | `[]`            | No       |
| `run_untagged`       | Run untagged jobs                | `bool`        | `true`          | No       |
| `active`             | Create this runner active        | `bool`        | `true`          | No       |
| `locked`             | Lock this runner to this project | `bool`        | `true`          | No       |
| `access_level`       | Run only against protected refs  | `string`      | "not_protected" | No       |
| `maximum_timeout`    | Maximum timeout for jobs         | `int`         | _none_          | No       |

`access_level` can be:

* `ref_protected` (only take jobs on protected refs)
* `not_protected` (take jobs from any ref)

### Calculated Attributes

| Name        | Description  | Type     | Sensitive |
|-------------|--------------|:--------:|:---------:|
| `runner_id` | Runner id    | `int`    | No        |
| `token`     | Runner token | `string` | Yes       |
