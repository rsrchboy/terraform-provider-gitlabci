
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

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->


- [Provider `gitlabci`](#provider-gitlabci)
- [Resources](#resources)
    - [Resource `gitlabci_runner_token`](#resource-gitlabci_runner_token)
- [Data Sources](#data-sources)
    - [Data Source `gitlabci_environment`](#data-source-gitlabci_environment)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# Provider `gitlabci`

... yada yada include original README.md here

| **Attribute** | **Description** | **Type** | **Optional** |
|-----------|-------------|------|----------|
| `base_url` |The GitLab base API URL| `string` | Optional |


# Resources

We have the following resources:

* `gitlabci_runner_token`



## Resource `gitlabci_runner_token`

This resource will take a registration token and use it to register a new
runner.  Tags, etc, may be specified here at create time.

**N.B.** Changing any parameter will force the creation of a new resource.
_Registration info cannot be changed by this resource._

Generally, all options are as listed at [API Doc -- "Register a new Runner"](https://docs.gitlab.com/ce/api/runners.html#register-a-new-runner).


**Attributes:**

| **Attribute** | **Description** | **Type** | **opt/req?** |
|-----------|-------------|------|----------|
| `access_level` |Run against all refs, or protected only| `string` | |
| `active` |Create the runner active, or paused?| `bool` | |
| `description` |Runner description| `string` | |
| `id` || `string` | |
| `locked` |Lock runner to project| `bool` | |
| `maximum_timeout` |Maximum timeout for jobs| `number` | |
| `registration_token` |Runner registration token (shared, group, or project)| `string` | |
| `run_untagged` |Take and run untagged jobs?| `bool` | |
| `runner_id` |Runner ID| `number` | |
| `tags` |List of tags for the runner| `set(string)` | |
| `token` |Generated (registered) runner token| `string` | |


# Data Sources

Provider `gitlabci` has the following data sources:


* `gitlabci_environment`




## Data Source `gitlabci_environment`



**Attributes:**

| **Attribute** | **Description** | **Type** | **opt/req?** |
|-----------|-------------|------|----------|
| `artifact_download_attempts` |Number of attempts to download artifacts running a job| `string` | |
| `chat_channel` |Source chat channel which triggered the [ChatOps](../chatops/README.md) command| `string` | |
| `chat_input` |Additional arguments passed in the [ChatOps](../chatops/README.md) command| `string` | |
| `ci` |Mark that job is executed in CI environment| `string` | |
| `ci_api_v4_url` |The GitLab API v4 root URL| `string` | |
| `ci_builds_dir` |Top-level directory where builds are executed.| `string` | |
| `ci_commit_before_sha` |The previous latest commit present on a branch before a merge request. Only populated when there is a merge request associated with the pipeline.| `string` | |
| `ci_commit_branch` |The commit branch name. Present only when building branches.| `string` | |
| `ci_commit_description` |The description of the commit: the message without first line, if the title is shorter than 100 characters; full message in other case.| `string` | |
| `ci_commit_message` |The full commit message.| `string` | |
| `ci_commit_ref_name` |The branch or tag name for which project is built| `string` | |
| `ci_commit_ref_protected` |If the job is running on a protected branch| `string` | |
| `ci_commit_ref_slug` |`$CI_COMMIT_REF_NAME` lowercased, shortened to 63 bytes, and with everything except `0-9` and `a-z` replaced with `-`. No leading / trailing `-`. Use in URLs, host names and domain names.| `string` | |
| `ci_commit_sha` |The commit revision for which project is built| `string` | |
| `ci_commit_short_sha` |The first eight characters of `CI_COMMIT_SHA`| `string` | |
| `ci_commit_tag` |The commit tag name. Present only when building tags.| `string` | |
| `ci_commit_title` |The title of the commit - the full first line of the message| `string` | |
| `ci_concurrent_id` |Unique ID of build execution within a single executor.| `string` | |
| `ci_concurrent_project_id` |Unique ID of build execution within a single executor and project.| `string` | |
| `ci_config_path` |The path to CI config file. Defaults to `.gitlab-ci.yml`| `string` | |
| `ci_debug_trace` |Whether [debug logging (tracing)](README.md#debug-logging) is enabled| `string` | |
| `ci_default_branch` |The name of the default branch for the project.| `string` | |
| `ci_deploy_password` |Authentication password of the [GitLab Deploy Token][gitlab-deploy-token], only present if the Project has one related.| `string` | |
| `ci_deploy_user` |Authentication username of the [GitLab Deploy Token][gitlab-deploy-token], only present if the Project has one related.| `string` | |
| `ci_disposable_environment` |Marks that the job is executed in a disposable environment (something that is created only for this job and disposed of/destroyed after the execution - all executors except `shell` and `ssh`). If the environment is disposable, it is set to true, otherwise it is not defined at all.| `string` | |
| `ci_environment_name` |The name of the environment for this job. Only present if [`environment:name`](../yaml/README.md#environmentname) is set.| `string` | |
| `ci_environment_slug` |A simplified version of the environment name, suitable for inclusion in DNS, URLs, Kubernetes labels, etc. Only present if [`environment:name`](../yaml/README.md#environmentname) is set.| `string` | |
| `ci_environment_url` |The URL of the environment for this job. Only present if [`environment:url`](../yaml/README.md#environmenturl) is set.| `string` | |
| `ci_external_pull_request_iid` |Pull Request ID from GitHub if the [pipelines are for external pull requests](../ci_cd_for_external_repos/index.md#pipelines-for-external-pull-requests). Available only if `only: [external_pull_requests]` is used and the pull request is open.| `string` | |
| `ci_external_pull_request_source_branch_name` |The source branch name of the pull request if [the pipelines are for external pull requests](../ci_cd_for_external_repos/index.md#pipelines-for-external-pull-requests). Available only if `only: [external_pull_requests]` is used and the pull request is open.| `string` | |
| `ci_external_pull_request_source_branch_sha` |The HEAD SHA of the source branch of the pull request if [the pipelines are for external pull requests](../ci_cd_for_external_repos/index.md#pipelines-for-external-pull-requests). Available only if `only: [external_pull_requests]` is used and the pull request is open.| `string` | |
| `ci_external_pull_request_target_branch_name` |The target branch name of the pull request if [the pipelines are for external pull requests](../ci_cd_for_external_repos/index.md#pipelines-for-external-pull-requests). Available only if `only: [external_pull_requests]` is used and the pull request is open.| `string` | |
| `ci_external_pull_request_target_branch_sha` |The HEAD SHA of the target branch of the pull request if [the pipelines are for external pull requests](../ci_cd_for_external_repos/index.md#pipelines-for-external-pull-requests). Available only if `only: [external_pull_requests]` is used and the pull request is open.| `string` | |
| `ci_job_id` |The unique id of the current job that GitLab CI uses internally| `string` | |
| `ci_job_manual` |The flag to indicate that job was manually started| `string` | |
| `ci_job_name` |The name of the job as defined in `.gitlab-ci.yml`| `string` | |
| `ci_job_stage` |The name of the stage as defined in `.gitlab-ci.yml`| `string` | |
| `ci_job_token` |Token used for authenticating with the [GitLab Container Registry][registry] and downloading [dependent repositories][dependent-repositories]| `string` | |
| `ci_job_url` |Job details URL| `string` | |
| `ci_merge_request_assignees` |Comma-separated list of username(s) of assignee(s) for the merge request if [the pipelines are for merge requests](../merge_request_pipelines/index.md). Available only if `only: [merge_requests]` is used and the merge request is created.| `string` | |
| `ci_merge_request_event_type` |The event type of the merge request, if [the pipelines are for merge requests](../merge_request_pipelines/index.md). Can be `detached`, `merged_result` or `merge_train`.| `string` | |
| `ci_merge_request_id` |The ID of the merge request if [the pipelines are for merge requests](../merge_request_pipelines/index.md). Available only if `only: [merge_requests]` is used and the merge request is created.| `string` | |
| `ci_merge_request_iid` |The IID of the merge request if [the pipelines are for merge requests](../merge_request_pipelines/index.md). Available only if `only: [merge_requests]` is used and the merge request is created.| `string` | |
| `ci_merge_request_labels` |Comma-separated label names of the merge request if [the pipelines are for merge requests](../merge_request_pipelines/index.md). Available only if `only: [merge_requests]` is used and the merge request is created.| `string` | |
| `ci_merge_request_milestone` |The milestone title of the merge request if [the pipelines are for merge requests](../merge_request_pipelines/index.md). Available only if `only: [merge_requests]` is used and the merge request is created.| `string` | |
| `ci_merge_request_project_id` |The ID of the project of the merge request if [the pipelines are for merge requests](../merge_request_pipelines/index.md). Available only if `only: [merge_requests]` is used and the merge request is created.| `string` | |
| `ci_merge_request_project_path` |The path of the project of the merge request if [the pipelines are for merge requests](../merge_request_pipelines/index.md) (e.g. `namespace/awesome-project`). Available only if `only: [merge_requests]` is used and the merge request is created.| `string` | |
| `ci_merge_request_project_url` |The URL of the project of the merge request if [the pipelines are for merge requests](../merge_request_pipelines/index.md) (e.g. `http://192.168.10.15:3000/namespace/awesome-project`). Available only if `only: [merge_requests]` is used and the merge request is created.| `string` | |
| `ci_merge_request_ref_path` |The ref path of the merge request if [the pipelines are for merge requests](../merge_request_pipelines/index.md). (e.g. `refs/merge-requests/1/head`). Available only if `only: [merge_requests]` is used and the merge request is created.| `string` | |
| `ci_merge_request_source_branch_name` |The source branch name of the merge request if [the pipelines are for merge requests](../merge_request_pipelines/index.md). Available only if `only: [merge_requests]` is used and the merge request is created.| `string` | |
| `ci_merge_request_source_branch_sha` |The HEAD SHA of the source branch of the merge request if [the pipelines are for merge requests](../merge_request_pipelines/index.md). Available only if `only: [merge_requests]` is used, the merge request is created, and the pipeline is a [merged result pipeline](../merge_request_pipelines/pipelines_for_merged_results/index.md). **(PREMIUM)**| `string` | |
| `ci_merge_request_source_project_id` |The ID of the source project of the merge request if [the pipelines are for merge requests](../merge_request_pipelines/index.md). Available only if `only: [merge_requests]` is used and the merge request is created.| `string` | |
| `ci_merge_request_source_project_path` |The path of the source project of the merge request if [the pipelines are for merge requests](../merge_request_pipelines/index.md). Available only if `only: [merge_requests]` is used and the merge request is created.| `string` | |
| `ci_merge_request_source_project_url` |The URL of the source project of the merge request if [the pipelines are for merge requests](../merge_request_pipelines/index.md). Available only if `only: [merge_requests]` is used and the merge request is created.| `string` | |
| `ci_merge_request_target_branch_name` |The target branch name of the merge request if [the pipelines are for merge requests](../merge_request_pipelines/index.md). Available only if `only: [merge_requests]` is used and the merge request is created.| `string` | |
| `ci_merge_request_target_branch_sha` |The HEAD SHA of the target branch of the merge request if [the pipelines are for merge requests](../merge_request_pipelines/index.md). Available only if `only: [merge_requests]` is used, the merge request is created, and the pipeline is a [merged result pipeline](../merge_request_pipelines/pipelines_for_merged_results/index.md). **(PREMIUM)**| `string` | |
| `ci_merge_request_title` |The title of the merge request if [the pipelines are for merge requests](../merge_request_pipelines/index.md). Available only if `only: [merge_requests]` is used and the merge request is created.| `string` | |
| `ci_node_index` |Index of the job in the job set. If the job is not parallelized, this variable is not set.| `string` | |
| `ci_node_total` |Total number of instances of this job running in parallel. If the job is not parallelized, this variable is set to `1`.| `string` | |
| `ci_pages_domain` |The configured domain that hosts GitLab Pages.| `string` | |
| `ci_pages_url` |URL to GitLab Pages-built pages. Always belongs to a subdomain of `CI_PAGES_DOMAIN`.| `string` | |
| `ci_pipeline_id` |The unique id of the current pipeline that GitLab CI uses internally| `string` | |
| `ci_pipeline_iid` |The unique id of the current pipeline scoped to project| `string` | |
| `ci_pipeline_source` |Indicates how the pipeline was triggered. Possible options are: `push`, `web`, `trigger`, `schedule`, `api`, `pipeline`, `external`, `chat`, `merge_request_event`, and `external_pull_request_event`. For pipelines created before GitLab 9.5, this will show as `unknown`| `string` | |
| `ci_pipeline_triggered` |The flag to indicate that job was [triggered](../triggers/README.md)| `string` | |
| `ci_pipeline_url` |Pipeline details URL| `string` | |
| `ci_project_dir` |The full path where the repository is cloned and where the job is run. If the GitLab Runner `builds_dir` parameter is set, this variable is set relative to the value of `builds_dir`. For more information, see [Advanced configuration](https://docs.gitlab.com/runner/configuration/advanced-configuration.html#the-runners-section) for GitLab Runner.| `string` | |
| `ci_project_id` |The unique id of the current project that GitLab CI uses internally| `string` | |
| `ci_project_name` |The name of the directory for the project that is currently being built. For example, if the project URL is `gitlab.example.com/group-name/project-1`, the `CI_PROJECT_NAME` would be `project-1`.| `string` | |
| `ci_project_namespace` |The project namespace (username or groupname) that is currently being built| `string` | |
| `ci_project_path` |The namespace with project name| `string` | |
| `ci_project_path_slug` |`$CI_PROJECT_PATH` lowercased and with everything except `0-9` and `a-z` replaced with `-`. Use in URLs and domain names.| `string` | |
| `ci_project_repository_languages` |Comma-separated, lowercased list of the languages used in the repository (e.g. `ruby,javascript,html,css`)| `string` | |
| `ci_project_title` |The human-readable project name as displayed in the GitLab web interface.| `string` | |
| `ci_project_url` |The HTTP(S) address to access project| `string` | |
| `ci_project_visibility` |The project visibility (internal, private, public)| `string` | |
| `ci_registry` |If the Container Registry is enabled it returns the address of GitLab's Container Registry.  This variable will include a `:port` value if one has been specified in the registry configuration.| `string` | |
| `ci_registry_image` |If the Container Registry is enabled for the project it returns the address of the registry tied to the specific project| `string` | |
| `ci_registry_password` |The password to use to push containers to the GitLab Container Registry| `string` | |
| `ci_registry_user` |The username to use to push containers to the GitLab Container Registry| `string` | |
| `ci_repository_url` |The URL to clone the Git repository| `string` | |
| `ci_runner_description` |The description of the runner as saved in GitLab| `string` | |
| `ci_runner_executable_arch` |The OS/architecture of the GitLab Runner executable (note that this is not necessarily the same as the environment of the executor)| `string` | |
| `ci_runner_id` |The unique id of runner being used| `string` | |
| `ci_runner_revision` |GitLab Runner revision that is executing the current job| `string` | |
| `ci_runner_short_token` |First eight characters of GitLab Runner's token used to authenticate new job requests. Used as Runner's unique ID| `string` | |
| `ci_runner_tags` |The defined runner tags| `string` | |
| `ci_runner_version` |GitLab Runner version that is executing the current job| `string` | |
| `ci_server` |Mark that job is executed in CI environment| `string` | |
| `ci_server_host` |Host component of the GitLab instance URL, without protocol and port (like `gitlab.example.com`)| `string` | |
| `ci_server_name` |The name of CI server that is used to coordinate jobs| `string` | |
| `ci_server_revision` |GitLab revision that is used to schedule jobs| `string` | |
| `ci_server_url` |The base URL of the GitLab instance, including protocol and port (like `https://gitlab.example.com:8080`)| `string` | |
| `ci_server_version` |GitLab version that is used to schedule jobs| `string` | |
| `ci_server_version_major` |GitLab version major component| `string` | |
| `ci_server_version_minor` |GitLab version minor component| `string` | |
| `ci_server_version_patch` |GitLab version patch component| `string` | |
| `ci_shared_environment` |Marks that the job is executed in a shared environment (something that is persisted across CI invocations like `shell` or `ssh` executor). If the environment is shared, it is set to true, otherwise it is not defined at all.| `string` | |
| `get_sources_attempts` |Number of attempts to fetch sources running a job| `string` | |
| `gitlab_ci` |Mark that job is executed in GitLab CI environment| `string` | |
| `gitlab_features` |The comma separated list of licensed features available for your instance and plan| `string` | |
| `gitlab_user_email` |The email of the user who started the job| `string` | |
| `gitlab_user_id` |The id of the user who started the job| `string` | |
| `gitlab_user_login` |The login username of the user who started the job| `string` | |
| `gitlab_user_name` |The real name of the user who started the job| `string` | |
| `id` |The computed configuation id| `string` | |
| `restore_cache_attempts` |Number of attempts to restore the cache running a job| `string` | |
| `running_under_ci` |True if we appear to be running as a CI job| `bool` | |


