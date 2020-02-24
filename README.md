
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
    - [Data Source `gitlabci_runner_config`](#data-source-gitlabci_runner_config)
        - [Block `runners`](#block-runners)
            - [Block `cache`](#block-cache)
                - [Block `gcs`](#block-gcs)
                - [Block `s3`](#block-s3)
            - [Block `custom`](#block-custom)
            - [Block `custom_build_dir`](#block-custom_build_dir)
            - [Block `docker`](#block-docker)
                - [Block `services`](#block-services)
            - [Block `kubernetes`](#block-kubernetes)
                - [Block `pod_security_context`](#block-pod_security_context)
                - [Block `services`](#block-services-1)
                - [Block `volumes`](#block-volumes)
                    - [Block `config_map`](#block-config_map)
                    - [Block `empty_dir`](#block-empty_dir)
                    - [Block `host_path`](#block-host_path)
                    - [Block `pvc`](#block-pvc)
                    - [Block `secret`](#block-secret)
            - [Block `machine`](#block-machine)
            - [Block `parallels`](#block-parallels)
            - [Block `referees`](#block-referees)
                - [Block `metrics`](#block-metrics)
            - [Block `ssh`](#block-ssh)
            - [Block `virtualbox`](#block-virtualbox)
        - [Block `session_server`](#block-session_server)

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

* `gitlabci_runner_config`




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




## Data Source `gitlabci_runner_config`



**Attributes:**

| **Attribute** | **Description** | **Type** | **opt/req?** |
|-----------|-------------|------|----------|
| `check_interval` |Defines the interval length, in seconds, between new jobs check. The default value is 3; if set to 0 or lower, the default value will be used| `number` | |
| `concurrent` |Limits how many jobs globally can be run concurrently| `number` | |
| `config` |The computed runner configuration (toml)| `string` | |
| `id` |The computed configuation id| `string` | |
| `listen_address` |Address (<host>:<port>) on which the Prometheus metrics HTTP server should be listening| `string` | |
| `log_format` |Log format (options: runner, text, json). Note that this setting has lower priority than format set by command line argument --log-format| `string` | |
| `log_level` |Log level (options: debug, info, warn, error, fatal, panic). Note that this setting has lower priority than level set by command line argument --debug, -l or --log-level| `string` | |




**Blocks:**

This data source also takes the following blocks (defined below):

* `runners`
* `session_server`


### Block `runners`



**Attributes:**

| **Attribute** | **Description** | **Type** | **opt/req?** |
|-----------|-------------|------|----------|
| `builds_dir` |Directory where builds are stored| `string` | |
| `cache_dir` |Directory where build cache is stored| `string` | |
| `clone_url` |Overwrite the default URL used to clone or fetch the git ref| `string` | |
| `config_template` |Configuration template (toml)| `string` | |
| `debug_trace_disabled` |When set to true Runner will disable the possibility of using the CI_DEBUG_TRACE feature| `bool` | |
| `environment` |Custom environment variables injected to build environment| `list(string)` | |
| `executor` |Select executor, eg. shell, docker, etc.| `string` | |
| `limit` |Maximum number of builds processed by this runner| `number` | |
| `name` |Runner name| `string` | |
| `output_limit` |Maximum build trace size in kilobytes| `number` | |
| `post_build_script` |Runner-specific command script executed after code is pulled and just after build executes| `string` | |
| `pre_build_script` |Runner-specific command script executed after code is pulled, just before build executes| `string` | |
| `pre_clone_script` |Runner-specific command script executed before code is pulled| `string` | |
| `request_concurrency` |Maximum concurrency for job requests| `number` | |
| `shell` |Select bash, cmd or powershell| `string` | |
| `tls_ca_file` |File containing the certificates to verify the peer when using HTTPS| `string` | |
| `tls_cert_file` |File containing certificate for TLS client auth when using HTTPS| `string` | |
| `tls_key_file` |File containing private key for TLS client auth when using HTTPS| `string` | |
| `token` |Runner token| `string` | |
| `url` |Runner URL| `string` | |




**Blocks:**

This block also takes the following blocks (defined below):

* `cache`
* `custom`
* `custom_build_dir`
* `docker`
* `kubernetes`
* `machine`
* `parallels`
* `referees`
* `ssh`
* `virtualbox`


#### Block `cache`



**Attributes:**

| **Attribute** | **Description** | **Type** | **opt/req?** |
|-----------|-------------|------|----------|
| `path` |Name of the path to prepend to the cache URL| `string` | |
| `shared` |Enable cache sharing between runners.| `bool` | |
| `type` |Select caching method| `string` | |




**Blocks:**

This block also takes the following blocks (defined below):

* `gcs`
* `s3`


##### Block `gcs`



**Attributes:**

| **Attribute** | **Description** | **Type** | **opt/req?** |
|-----------|-------------|------|----------|
| `access_id` |ID of GCP Service Account used to access the storage| `string` | |
| `bucket_name` |Name of the bucket where cache will be stored| `string` | |
| `credentials_file` |File with GCP credentials, containing AccessID and PrivateKey| `string` | |
| `private_key` |Private key used to sign GCS requests| `string` | |




##### Block `s3`



**Attributes:**

| **Attribute** | **Description** | **Type** | **opt/req?** |
|-----------|-------------|------|----------|
| `access_key` |S3 Access Key| `string` | |
| `bucket_location` |Name of S3 region| `string` | |
| `bucket_name` |Name of the bucket where cache will be stored| `string` | |
| `insecure` |Use insecure mode (without https)| `bool` | |
| `secret_key` |S3 Secret Key| `string` | |
| `server_address` |A host:port to the used S3-compatible server| `string` | |




#### Block `custom`



**Attributes:**

| **Attribute** | **Description** | **Type** | **opt/req?** |
|-----------|-------------|------|----------|
| `cleanup_args` |Arguments for the cleanup executable| `list(string)` | |
| `cleanup_exec` |Executable that cleanups after executor run| `string` | |
| `cleanup_exec_timeout` |Timeout for the cleanup executable (in seconds)| `number` | |
| `config_args` |Arguments for the config executable| `list(string)` | |
| `config_exec` |Executable that allows to inject configuration values to the executor| `string` | |
| `config_exec_timeout` |Timeout for the config executable (in seconds)| `number` | |
| `force_kill_timeout` |Force timeout for scripts execution (in seconds). Counted from the force kill call; if process will be not terminated, Runner will abandon process termination and log an error| `number` | |
| `graceful_kill_timeout` |Graceful timeout for scripts execution after SIGTERM is sent to the process (in seconds). This limits the time given for scripts to perform the cleanup before exiting| `number` | |
| `prepare_args` |Arguments for the prepare executable| `list(string)` | |
| `prepare_exec` |Executable that prepares executor| `string` | |
| `prepare_exec_timeout` |Timeout for the prepare executable (in seconds)| `number` | |
| `run_args` |Arguments for the run executable| `list(string)` | |
| `run_exec` |Executable that runs the job script in executor| `string` | |




#### Block `custom_build_dir`



**Attributes:**

| **Attribute** | **Description** | **Type** | **opt/req?** |
|-----------|-------------|------|----------|
| `enabled` |Enable job specific build directories| `bool` | |




#### Block `docker`



**Attributes:**

| **Attribute** | **Description** | **Type** | **opt/req?** |
|-----------|-------------|------|----------|
| `allowed_images` |Whitelist allowed images| `list(string)` | |
| `allowed_services` |Whitelist allowed services| `list(string)` | |
| `cache_dir` |Directory where to store caches| `string` | |
| `cap_add` |Add Linux capabilities| `list(string)` | |
| `cap_drop` |Drop Linux capabilities| `list(string)` | |
| `cpu_shares` |Number of CPU shares| `number` | |
| `cpus` |Number of CPUs| `string` | |
| `cpuset_cpus` |String value containing the cgroups CpusetCpus to use| `string` | |
| `devices` |Add a host device to the container| `list(string)` | |
| `disable_cache` |Disable all container caching| `bool` | |
| `disable_entrypoint_overwrite` |Disable the possibility for a container to overwrite the default image entrypoint| `bool` | |
| `dns` |A list of DNS servers for the container to use| `list(string)` | |
| `dns_search` |A list of DNS search domains| `list(string)` | |
| `extra_hosts` |Add a custom host-to-IP mapping| `list(string)` | |
| `helper_image` |[ADVANCED] Override the default helper image used to clone repos and upload artifacts| `string` | |
| `host` |Docker daemon address| `string` | |
| `hostname` |Custom container hostname| `string` | |
| `image` |Docker image to be used| `string` | |
| `links` |Add link to another container| `list(string)` | |
| `memory` |Memory limit (format: <number>[<unit>]). Unit can be one of b, k, m, or g. Minimum is 4M.| `string` | |
| `memory_reservation` |Memory soft limit (format: <number>[<unit>]). Unit can be one of b, k, m, or g.| `string` | |
| `memory_swap` |Total memory limit (memory + swap, format: <number>[<unit>]). Unit can be one of b, k, m, or g.| `string` | |
| `network_mode` |Add container to a custom network| `string` | |
| `oom_kill_disable` |Do not kill processes in a container if an out-of-memory (OOM) error occurs| `bool` | |
| `oom_score_adjust` |Adjust OOM score| `number` | |
| `privileged` |Give extended privileges to container| `bool` | |
| `pull_policy` |Image pull policy: never, if-not-present, always| `string` | |
| `runtime` |Docker runtime to be used| `string` | |
| `security_opt` |Security Options| `list(string)` | |
| `services_tmpfs` |A toml table/json object with the format key=values. When set this will mount the specified path in the key as a tmpfs volume in all the service containers, using the options specified as key. For the supported options, see the documentation for the unix 'mount' command| `map(string)` | |
| `shm_size` |Shared memory size for docker images (in bytes)| `number` | |
| `sysctls` |Sysctl options, a toml table/json object of key=value. Value is expected to be a string.| `map(string)` | |
| `tls_cert_path` |Certificate path| `string` | |
| `tls_verify` |Use TLS and verify the remote| `bool` | |
| `tmpfs` |A toml table/json object with the format key=values. When set this will mount the specified path in the key as a tmpfs volume in the main container, using the options specified as key. For the supported options, see the documentation for the unix 'mount' command| `map(string)` | |
| `userns_mode` |User namespace to use| `string` | |
| `volume_driver` |Volume driver to be used| `string` | |
| `volumes` |Bind-mount a volume and create it if it doesn't exist prior to mounting. Can be specified multiple times once per mountpoint, e.g. --docker-volumes 'test0:/test0' --docker-volumes 'test1:/test1'| `list(string)` | |
| `volumes_from` |A list of volumes to inherit from another container| `list(string)` | |
| `wait_for_services_timeout` |How long to wait for service startup| `number` | |




**Blocks:**

This block also takes the following blocks (defined below):

* `services`


##### Block `services`



**Attributes:**

| **Attribute** | **Description** | **Type** | **opt/req?** |
|-----------|-------------|------|----------|
| `alias` |The alias of the service| `string` | |
| `name` |The image path for the service| `string` | |




#### Block `kubernetes`



**Attributes:**

| **Attribute** | **Description** | **Type** | **opt/req?** |
|-----------|-------------|------|----------|
| `bearer_token` |Optional Kubernetes service account token used to start build pods.| `string` | |
| `bearer_token_overwrite_allowed` |Bool to authorize builds to specify their own bearer token for creation.| `bool` | |
| `ca_file` |Optional Kubernetes master auth ca certificate| `string` | |
| `cert_file` |Optional Kubernetes master auth certificate| `string` | |
| `cpu_limit` |The CPU allocation given to build containers| `string` | |
| `cpu_request` |The CPU allocation requested for build containers| `string` | |
| `helper_cpu_limit` |The CPU allocation given to build helper containers| `string` | |
| `helper_cpu_request` |The CPU allocation requested for build helper containers| `string` | |
| `helper_image` |[ADVANCED] Override the default helper image used to clone repos and upload artifacts| `string` | |
| `helper_memory_limit` |The amount of memory allocated to build helper containers| `string` | |
| `helper_memory_request` |The amount of memory requested for build helper containers| `string` | |
| `host` |Optional Kubernetes master host URL (auto-discovery attempted if not specified)| `string` | |
| `image` |Default docker image to use for builds when none is specified| `string` | |
| `image_pull_secrets` |A list of image pull secrets that are used for pulling docker image| `list(string)` | |
| `key_file` |Optional Kubernetes master auth private key| `string` | |
| `memory_limit` |The amount of memory allocated to build containers| `string` | |
| `memory_request` |The amount of memory requested from build containers| `string` | |
| `namespace` |Namespace to run Kubernetes jobs in| `string` | |
| `namespace_overwrite_allowed` |Regex to validate 'KUBERNETES_NAMESPACE_OVERWRITE' value| `string` | |
| `node_selector` |A toml table/json object of key=value. Value is expected to be a string. When set this will create pods on k8s nodes that match all the key=value pairs.| `map(string)` | |
| `node_tolerations` |A toml table/json object of key=value:effect. Value and effect are expected to be strings. When set, pods will tolerate the given taints. Only one toleration is supported through environment variable configuration.| `map(string)` | |
| `pod_annotations` |A toml table/json object of key-value. Value is expected to be a string. When set, this will create pods with the given annotations. Can be overwritten in build with KUBERNETES_POD_ANNOTATION_* variables| `map(string)` | |
| `pod_annotations_overwrite_allowed` |Regex to validate 'KUBERNETES_POD_ANNOTATIONS_*' values| `string` | |
| `pod_labels` |A toml table/json object of key-value. Value is expected to be a string. When set, this will create pods with the given pod labels. Environment variables will be substituted for values here.| `map(string)` | |
| `poll_interval` |How frequently, in seconds, the runner will poll the Kubernetes pod it has just created to check its status| `number` | |
| `poll_timeout` |The total amount of time, in seconds, that needs to pass before the runner will timeout attempting to connect to the pod it has just created (useful for queueing more builds that the cluster can handle at a time)| `number` | |
| `privileged` |Run all containers with the privileged flag enabled| `bool` | |
| `pull_policy` |Policy for if/when to pull a container image (never, if-not-present, always). The cluster default will be used if not set| `string` | |
| `service_account` |Executor pods will use this Service Account to talk to kubernetes API| `string` | |
| `service_account_overwrite_allowed` |Regex to validate 'KUBERNETES_SERVICE_ACCOUNT' value| `string` | |
| `service_cpu_limit` |The CPU allocation given to build service containers| `string` | |
| `service_cpu_request` |The CPU allocation requested for build service containers| `string` | |
| `service_memory_limit` |The amount of memory allocated to build service containers| `string` | |
| `service_memory_request` |The amount of memory requested for build service containers| `string` | |
| `termination_grace_period_seconds` |Duration after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal.| `number` | |




**Blocks:**

This block also takes the following blocks (defined below):

* `pod_security_context`
* `services`
* `volumes`


##### Block `pod_security_context`



**Attributes:**

| **Attribute** | **Description** | **Type** | **opt/req?** |
|-----------|-------------|------|----------|
| `fs_group` |A special supplemental group that applies to all containers in a pod| `number` | |
| `run_as_group` |The GID to run the entrypoint of the container process| `number` | |
| `run_as_non_root` |Indicates that the container must run as a non-root user| `bool` | |
| `run_as_user` |The UID to run the entrypoint of the container process| `number` | |
| `supplemental_groups` |A list of groups applied to the first process run in each container, in addition to the container's primary GID| `list(number)` | |




##### Block `services`



**Attributes:**

| **Attribute** | **Description** | **Type** | **opt/req?** |
|-----------|-------------|------|----------|
| `name` |The image path for the service| `string` | |




##### Block `volumes`



**Attributes:**

This block has no attributes.

**Blocks:**

This block also takes the following blocks (defined below):

* `config_map`
* `empty_dir`
* `host_path`
* `pvc`
* `secret`


###### Block `config_map`



**Attributes:**

| **Attribute** | **Description** | **Type** | **opt/req?** |
|-----------|-------------|------|----------|
| `items` |Key-to-path mapping for keys from the config map that should be used.| `map(string)` | |
| `mount_path` |Path where volume should be mounted inside of container| `string` | |
| `name` |The name of the volume and ConfigMap to use| `string` | |
| `read_only` |If this volume should be mounted read only| `bool` | |




###### Block `empty_dir`



**Attributes:**

| **Attribute** | **Description** | **Type** | **opt/req?** |
|-----------|-------------|------|----------|
| `medium` |Set to 'Memory' to have a tmpfs| `string` | |
| `mount_path` |Path where volume should be mounted inside of container| `string` | |
| `name` |The name of the volume and EmptyDir to use| `string` | |




###### Block `host_path`



**Attributes:**

| **Attribute** | **Description** | **Type** | **opt/req?** |
|-----------|-------------|------|----------|
| `host_path` |Path from the host that should be mounted as a volume| `string` | |
| `mount_path` |Path where volume should be mounted inside of container| `string` | |
| `name` |The name of the volume| `string` | |
| `read_only` |If this volume should be mounted read only| `bool` | |




###### Block `pvc`



**Attributes:**

| **Attribute** | **Description** | **Type** | **opt/req?** |
|-----------|-------------|------|----------|
| `mount_path` |Path where volume should be mounted inside of container| `string` | |
| `name` |The name of the volume and PVC to use| `string` | |
| `read_only` |If this volume should be mounted read only| `bool` | |




###### Block `secret`



**Attributes:**

| **Attribute** | **Description** | **Type** | **opt/req?** |
|-----------|-------------|------|----------|
| `items` |Key-to-path mapping for keys from the secret that should be used.| `map(string)` | |
| `mount_path` |Path where volume should be mounted inside of container| `string` | |
| `name` |The name of the volume and Secret to use| `string` | |
| `read_only` |If this volume should be mounted read only| `bool` | |




#### Block `machine`



**Attributes:**

| **Attribute** | **Description** | **Type** | **opt/req?** |
|-----------|-------------|------|----------|
| `idle_count` |Maximum idle machines| `number` | |
| `idle_time` |Minimum time after node can be destroyed| `number` | |
| `machine_driver` |The driver to use when creating machine| `string` | |
| `machine_name` |The template for machine name (needs to include %s)| `string` | |
| `machine_options` |Additional machine creation options| `list(string)` | |
| `max_builds` |Maximum number of builds processed by machine| `number` | |
| `off_peak_idle_count` |Maximum idle machines when the scheduler is in the OffPeak mode| `number` | |
| `off_peak_idle_time` |Minimum time after machine can be destroyed when the scheduler is in the OffPeak mode| `number` | |
| `off_peak_periods` |Time periods when the scheduler is in the OffPeak mode| `list(string)` | |
| `off_peak_timezone` |Timezone for the OffPeak periods (defaults to Local)| `string` | |




#### Block `parallels`



**Attributes:**

| **Attribute** | **Description** | **Type** | **opt/req?** |
|-----------|-------------|------|----------|
| `base_name` |VM name to be used| `string` | |
| `disable_snapshots` |Disable snapshoting to speedup VM creation| `bool` | |
| `template_name` |VM template to be created| `string` | |
| `time_server` |Timeserver to sync the guests time from. Defaults to time.apple.com| `string` | |




#### Block `referees`



**Attributes:**

This block has no attributes.

**Blocks:**

This block also takes the following blocks (defined below):

* `metrics`


##### Block `metrics`



**Attributes:**

| **Attribute** | **Description** | **Type** | **opt/req?** |
|-----------|-------------|------|----------|
| `prometheus_address` |A host:port to a prometheus metrics server| `string` | |
| `queries` |A list of metrics to query (in PromQL)| `list(string)` | |
| `query_interval` |Query interval (in seconds)| `number` | |




#### Block `ssh`



**Attributes:**

| **Attribute** | **Description** | **Type** | **opt/req?** |
|-----------|-------------|------|----------|
| `host` |Remote host| `string` | |
| `identity_file` |Identity file to be used| `string` | |
| `password` |User password| `string` | |
| `port` |Remote host port| `string` | |
| `user` |User name| `string` | |




#### Block `virtualbox`



**Attributes:**

| **Attribute** | **Description** | **Type** | **opt/req?** |
|-----------|-------------|------|----------|
| `base_name` |VM name to be used| `string` | |
| `base_snapshot` |Name or UUID of a specific VM snapshot to clone| `string` | |
| `disable_snapshots` |Disable snapshoting to speedup VM creation| `bool` | |




### Block `session_server`



**Attributes:**

| **Attribute** | **Description** | **Type** | **opt/req?** |
|-----------|-------------|------|----------|
| `advertise_address` |FIXME| `string` | |
| `listen_address` |FIXME| `string` | |
| `session_timeout` |FIXME| `number` | |


