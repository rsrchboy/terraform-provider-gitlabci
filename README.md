
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
| base_url | The GitLab base API URL | string | %!s(bool=true) |


# Resources

We have the following resources:

* gitlabci_runner_token



## Resource `gitlabci_runner_token`

This resource will take a registration token and use it to register a new
runner.  Tags, etc, may be specified here at create time.

**N.B.** Changing any parameter will force the creation of a new resource.
_Registration info cannot be changed by this resource._

Generally, all options are as listed at [API Doc -- "Register a new Runner"](https://docs.gitlab.com/ce/api/runners.html#register-a-new-runner).


**Attributes:**

| **Attribute** | **Description** | **Type** | **opt/req?** |
|-----------|-------------|------|----------|
| access_level | string | Run against all refs, or protected only |  |
| active | bool | Create the runner active, or paused? |  |
| description | string | Runner description |  |
| id | string |  |  |
| locked | bool | Lock runner to project |  |
| maximum_timeout | number | Maximum timeout for jobs |  |
| registration_token | string | Runner registration token (shared, group, or project) |  |
| run_untagged | bool | Take and run untagged jobs? |  |
| runner_id | number | Runner ID |  |
| tags | [set string] | List of tags for the runner |  |
| token | string | Generated (registered) runner token |  |


# Data Sources

Provider `gitlabci` has the following data sources:


* gitlabci_environment

* gitlabci_runner_config




## Data Source `gitlabci_environment`



**Attributes:**

| **Attribute** | **Description** | **Type** | **opt/req?** |
|-----------|-------------|------|----------|
| artifact_download_attempts | string | Number of attempts to download artifacts running a job |  |
| chat_channel | string | Source chat channel which triggered the [ChatOps](../chatops/README.md) command |  |
| chat_input | string | Additional arguments passed in the [ChatOps](../chatops/README.md) command |  |
| ci | string | Mark that job is executed in CI environment |  |
| ci_api_v4_url | string | The GitLab API v4 root URL |  |
| ci_builds_dir | string | Top-level directory where builds are executed. |  |
| ci_commit_before_sha | string | The previous latest commit present on a branch before a merge request. Only populated when there is a merge request associated with the pipeline. |  |
| ci_commit_branch | string | The commit branch name. Present only when building branches. |  |
| ci_commit_description | string | The description of the commit: the message without first line, if the title is shorter than 100 characters; full message in other case. |  |
| ci_commit_message | string | The full commit message. |  |
| ci_commit_ref_name | string | The branch or tag name for which project is built |  |
| ci_commit_ref_protected | string | If the job is running on a protected branch |  |
| ci_commit_ref_slug | string | `$CI_COMMIT_REF_NAME` lowercased, shortened to 63 bytes, and with everything except `0-9` and `a-z` replaced with `-`. No leading / trailing `-`. Use in URLs, host names and domain names. |  |
| ci_commit_sha | string | The commit revision for which project is built |  |
| ci_commit_short_sha | string | The first eight characters of `CI_COMMIT_SHA` |  |
| ci_commit_tag | string | The commit tag name. Present only when building tags. |  |
| ci_commit_title | string | The title of the commit - the full first line of the message |  |
| ci_concurrent_id | string | Unique ID of build execution within a single executor. |  |
| ci_concurrent_project_id | string | Unique ID of build execution within a single executor and project. |  |
| ci_config_path | string | The path to CI config file. Defaults to `.gitlab-ci.yml` |  |
| ci_debug_trace | string | Whether [debug logging (tracing)](README.md#debug-logging) is enabled |  |
| ci_default_branch | string | The name of the default branch for the project. |  |
| ci_deploy_password | string | Authentication password of the [GitLab Deploy Token][gitlab-deploy-token], only present if the Project has one related. |  |
| ci_deploy_user | string | Authentication username of the [GitLab Deploy Token][gitlab-deploy-token], only present if the Project has one related. |  |
| ci_disposable_environment | string | Marks that the job is executed in a disposable environment (something that is created only for this job and disposed of/destroyed after the execution - all executors except `shell` and `ssh`). If the environment is disposable, it is set to true, otherwise it is not defined at all. |  |
| ci_environment_name | string | The name of the environment for this job. Only present if [`environment:name`](../yaml/README.md#environmentname) is set. |  |
| ci_environment_slug | string | A simplified version of the environment name, suitable for inclusion in DNS, URLs, Kubernetes labels, etc. Only present if [`environment:name`](../yaml/README.md#environmentname) is set. |  |
| ci_environment_url | string | The URL of the environment for this job. Only present if [`environment:url`](../yaml/README.md#environmenturl) is set. |  |
| ci_external_pull_request_iid | string | Pull Request ID from GitHub if the [pipelines are for external pull requests](../ci_cd_for_external_repos/index.md#pipelines-for-external-pull-requests). Available only if `only: [external_pull_requests]` is used and the pull request is open. |  |
| ci_external_pull_request_source_branch_name | string | The source branch name of the pull request if [the pipelines are for external pull requests](../ci_cd_for_external_repos/index.md#pipelines-for-external-pull-requests). Available only if `only: [external_pull_requests]` is used and the pull request is open. |  |
| ci_external_pull_request_source_branch_sha | string | The HEAD SHA of the source branch of the pull request if [the pipelines are for external pull requests](../ci_cd_for_external_repos/index.md#pipelines-for-external-pull-requests). Available only if `only: [external_pull_requests]` is used and the pull request is open. |  |
| ci_external_pull_request_target_branch_name | string | The target branch name of the pull request if [the pipelines are for external pull requests](../ci_cd_for_external_repos/index.md#pipelines-for-external-pull-requests). Available only if `only: [external_pull_requests]` is used and the pull request is open. |  |
| ci_external_pull_request_target_branch_sha | string | The HEAD SHA of the target branch of the pull request if [the pipelines are for external pull requests](../ci_cd_for_external_repos/index.md#pipelines-for-external-pull-requests). Available only if `only: [external_pull_requests]` is used and the pull request is open. |  |
| ci_job_id | string | The unique id of the current job that GitLab CI uses internally |  |
| ci_job_manual | string | The flag to indicate that job was manually started |  |
| ci_job_name | string | The name of the job as defined in `.gitlab-ci.yml` |  |
| ci_job_stage | string | The name of the stage as defined in `.gitlab-ci.yml` |  |
| ci_job_token | string | Token used for authenticating with the [GitLab Container Registry][registry] and downloading [dependent repositories][dependent-repositories] |  |
| ci_job_url | string | Job details URL |  |
| ci_merge_request_assignees | string | Comma-separated list of username(s) of assignee(s) for the merge request if [the pipelines are for merge requests](../merge_request_pipelines/index.md). Available only if `only: [merge_requests]` is used and the merge request is created. |  |
| ci_merge_request_event_type | string | The event type of the merge request, if [the pipelines are for merge requests](../merge_request_pipelines/index.md). Can be `detached`, `merged_result` or `merge_train`. |  |
| ci_merge_request_id | string | The ID of the merge request if [the pipelines are for merge requests](../merge_request_pipelines/index.md). Available only if `only: [merge_requests]` is used and the merge request is created. |  |
| ci_merge_request_iid | string | The IID of the merge request if [the pipelines are for merge requests](../merge_request_pipelines/index.md). Available only if `only: [merge_requests]` is used and the merge request is created. |  |
| ci_merge_request_labels | string | Comma-separated label names of the merge request if [the pipelines are for merge requests](../merge_request_pipelines/index.md). Available only if `only: [merge_requests]` is used and the merge request is created. |  |
| ci_merge_request_milestone | string | The milestone title of the merge request if [the pipelines are for merge requests](../merge_request_pipelines/index.md). Available only if `only: [merge_requests]` is used and the merge request is created. |  |
| ci_merge_request_project_id | string | The ID of the project of the merge request if [the pipelines are for merge requests](../merge_request_pipelines/index.md). Available only if `only: [merge_requests]` is used and the merge request is created. |  |
| ci_merge_request_project_path | string | The path of the project of the merge request if [the pipelines are for merge requests](../merge_request_pipelines/index.md) (e.g. `namespace/awesome-project`). Available only if `only: [merge_requests]` is used and the merge request is created. |  |
| ci_merge_request_project_url | string | The URL of the project of the merge request if [the pipelines are for merge requests](../merge_request_pipelines/index.md) (e.g. `http://192.168.10.15:3000/namespace/awesome-project`). Available only if `only: [merge_requests]` is used and the merge request is created. |  |
| ci_merge_request_ref_path | string | The ref path of the merge request if [the pipelines are for merge requests](../merge_request_pipelines/index.md). (e.g. `refs/merge-requests/1/head`). Available only if `only: [merge_requests]` is used and the merge request is created. |  |
| ci_merge_request_source_branch_name | string | The source branch name of the merge request if [the pipelines are for merge requests](../merge_request_pipelines/index.md). Available only if `only: [merge_requests]` is used and the merge request is created. |  |
| ci_merge_request_source_branch_sha | string | The HEAD SHA of the source branch of the merge request if [the pipelines are for merge requests](../merge_request_pipelines/index.md). Available only if `only: [merge_requests]` is used, the merge request is created, and the pipeline is a [merged result pipeline](../merge_request_pipelines/pipelines_for_merged_results/index.md). **(PREMIUM)** |  |
| ci_merge_request_source_project_id | string | The ID of the source project of the merge request if [the pipelines are for merge requests](../merge_request_pipelines/index.md). Available only if `only: [merge_requests]` is used and the merge request is created. |  |
| ci_merge_request_source_project_path | string | The path of the source project of the merge request if [the pipelines are for merge requests](../merge_request_pipelines/index.md). Available only if `only: [merge_requests]` is used and the merge request is created. |  |
| ci_merge_request_source_project_url | string | The URL of the source project of the merge request if [the pipelines are for merge requests](../merge_request_pipelines/index.md). Available only if `only: [merge_requests]` is used and the merge request is created. |  |
| ci_merge_request_target_branch_name | string | The target branch name of the merge request if [the pipelines are for merge requests](../merge_request_pipelines/index.md). Available only if `only: [merge_requests]` is used and the merge request is created. |  |
| ci_merge_request_target_branch_sha | string | The HEAD SHA of the target branch of the merge request if [the pipelines are for merge requests](../merge_request_pipelines/index.md). Available only if `only: [merge_requests]` is used, the merge request is created, and the pipeline is a [merged result pipeline](../merge_request_pipelines/pipelines_for_merged_results/index.md). **(PREMIUM)** |  |
| ci_merge_request_title | string | The title of the merge request if [the pipelines are for merge requests](../merge_request_pipelines/index.md). Available only if `only: [merge_requests]` is used and the merge request is created. |  |
| ci_node_index | string | Index of the job in the job set. If the job is not parallelized, this variable is not set. |  |
| ci_node_total | string | Total number of instances of this job running in parallel. If the job is not parallelized, this variable is set to `1`. |  |
| ci_pages_domain | string | The configured domain that hosts GitLab Pages. |  |
| ci_pages_url | string | URL to GitLab Pages-built pages. Always belongs to a subdomain of `CI_PAGES_DOMAIN`. |  |
| ci_pipeline_id | string | The unique id of the current pipeline that GitLab CI uses internally |  |
| ci_pipeline_iid | string | The unique id of the current pipeline scoped to project |  |
| ci_pipeline_source | string | Indicates how the pipeline was triggered. Possible options are: `push`, `web`, `trigger`, `schedule`, `api`, `pipeline`, `external`, `chat`, `merge_request_event`, and `external_pull_request_event`. For pipelines created before GitLab 9.5, this will show as `unknown` |  |
| ci_pipeline_triggered | string | The flag to indicate that job was [triggered](../triggers/README.md) |  |
| ci_pipeline_url | string | Pipeline details URL |  |
| ci_project_dir | string | The full path where the repository is cloned and where the job is run. If the GitLab Runner `builds_dir` parameter is set, this variable is set relative to the value of `builds_dir`. For more information, see [Advanced configuration](https://docs.gitlab.com/runner/configuration/advanced-configuration.html#the-runners-section) for GitLab Runner. |  |
| ci_project_id | string | The unique id of the current project that GitLab CI uses internally |  |
| ci_project_name | string | The name of the directory for the project that is currently being built. For example, if the project URL is `gitlab.example.com/group-name/project-1`, the `CI_PROJECT_NAME` would be `project-1`. |  |
| ci_project_namespace | string | The project namespace (username or groupname) that is currently being built |  |
| ci_project_path | string | The namespace with project name |  |
| ci_project_path_slug | string | `$CI_PROJECT_PATH` lowercased and with everything except `0-9` and `a-z` replaced with `-`. Use in URLs and domain names. |  |
| ci_project_repository_languages | string | Comma-separated, lowercased list of the languages used in the repository (e.g. `ruby,javascript,html,css`) |  |
| ci_project_title | string | The human-readable project name as displayed in the GitLab web interface. |  |
| ci_project_url | string | The HTTP(S) address to access project |  |
| ci_project_visibility | string | The project visibility (internal, private, public) |  |
| ci_registry | string | If the Container Registry is enabled it returns the address of GitLab's Container Registry.  This variable will include a `:port` value if one has been specified in the registry configuration. |  |
| ci_registry_image | string | If the Container Registry is enabled for the project it returns the address of the registry tied to the specific project |  |
| ci_registry_password | string | The password to use to push containers to the GitLab Container Registry |  |
| ci_registry_user | string | The username to use to push containers to the GitLab Container Registry |  |
| ci_repository_url | string | The URL to clone the Git repository |  |
| ci_runner_description | string | The description of the runner as saved in GitLab |  |
| ci_runner_executable_arch | string | The OS/architecture of the GitLab Runner executable (note that this is not necessarily the same as the environment of the executor) |  |
| ci_runner_id | string | The unique id of runner being used |  |
| ci_runner_revision | string | GitLab Runner revision that is executing the current job |  |
| ci_runner_short_token | string | First eight characters of GitLab Runner's token used to authenticate new job requests. Used as Runner's unique ID |  |
| ci_runner_tags | string | The defined runner tags |  |
| ci_runner_version | string | GitLab Runner version that is executing the current job |  |
| ci_server | string | Mark that job is executed in CI environment |  |
| ci_server_host | string | Host component of the GitLab instance URL, without protocol and port (like `gitlab.example.com`) |  |
| ci_server_name | string | The name of CI server that is used to coordinate jobs |  |
| ci_server_revision | string | GitLab revision that is used to schedule jobs |  |
| ci_server_url | string | The base URL of the GitLab instance, including protocol and port (like `https://gitlab.example.com:8080`) |  |
| ci_server_version | string | GitLab version that is used to schedule jobs |  |
| ci_server_version_major | string | GitLab version major component |  |
| ci_server_version_minor | string | GitLab version minor component |  |
| ci_server_version_patch | string | GitLab version patch component |  |
| ci_shared_environment | string | Marks that the job is executed in a shared environment (something that is persisted across CI invocations like `shell` or `ssh` executor). If the environment is shared, it is set to true, otherwise it is not defined at all. |  |
| get_sources_attempts | string | Number of attempts to fetch sources running a job |  |
| gitlab_ci | string | Mark that job is executed in GitLab CI environment |  |
| gitlab_features | string | The comma separated list of licensed features available for your instance and plan |  |
| gitlab_user_email | string | The email of the user who started the job |  |
| gitlab_user_id | string | The id of the user who started the job |  |
| gitlab_user_login | string | The login username of the user who started the job |  |
| gitlab_user_name | string | The real name of the user who started the job |  |
| id | string | The computed configuation id |  |
| restore_cache_attempts | string | Number of attempts to restore the cache running a job |  |
| running_under_ci | bool | True if we appear to be running as a CI job |  |




## Data Source `gitlabci_runner_config`



**Attributes:**

| **Attribute** | **Description** | **Type** | **opt/req?** |
|-----------|-------------|------|----------|
| check_interval | number | FIXME |  |
| concurrent | number | FIXME |  |
| config | string | The computed runner configuration (toml) |  |
| id | string | The computed configuation id |  |
| log_format | string | FIXME |  |
| log_level | string | FIXME |  |




**Blocks:**

This data source also takes the following blocks (defined below):

* runners
* session_server


### Block `runners`



**Attributes:**

| **Attribute** | **Description** | **Type** | **opt/req?** |
|-----------|-------------|------|----------|
| builds_dir | string | Directory where builds are stored |  |
| cache_dir | string | Directory where build cache is stored |  |
| clone_url | string | Overwrite the default URL used to clone or fetch the git ref |  |
| config_template | string | Configuration template (toml) |  |
| debug_trace_disabled | bool | When set to true Runner will disable the possibility of using the CI_DEBUG_TRACE feature |  |
| environment | [list string] | Custom environment variables injected to build environment |  |
| executor | string | Select executor, eg. shell, docker, etc. |  |
| limit | number | Maximum number of builds processed by this runner |  |
| name | string | Runner name |  |
| output_limit | number | Maximum build trace size in kilobytes |  |
| post_build_script | string | Runner-specific command script executed after code is pulled and just after build executes |  |
| pre_build_script | string | Runner-specific command script executed after code is pulled, just before build executes |  |
| pre_clone_script | string | Runner-specific command script executed before code is pulled |  |
| request_concurrency | number | Maximum concurrency for job requests |  |
| shell | string | Select bash, cmd or powershell |  |
| tls_ca_file | string | File containing the certificates to verify the peer when using HTTPS |  |
| tls_cert_file | string | File containing certificate for TLS client auth when using HTTPS |  |
| tls_key_file | string | File containing private key for TLS client auth when using HTTPS |  |
| token | string | Runner token |  |
| url | string | Runner URL |  |




**Blocks:**

This block also takes the following blocks (defined below):

* cache
* custom
* custom_build_dir
* docker
* kubernetes
* machine
* parallels
* referees
* ssh
* virtualbox


#### Block `cache`



**Attributes:**

| **Attribute** | **Description** | **Type** | **opt/req?** |
|-----------|-------------|------|----------|
| path | string | Name of the path to prepend to the cache URL |  |
| shared | bool | Enable cache sharing between runners. |  |
| type | string | Select caching method |  |




**Blocks:**

This block also takes the following blocks (defined below):

* gcs
* s3


##### Block `gcs`



**Attributes:**

| **Attribute** | **Description** | **Type** | **opt/req?** |
|-----------|-------------|------|----------|
| access_id | string | ID of GCP Service Account used to access the storage |  |
| bucket_name | string | Name of the bucket where cache will be stored |  |
| credentials_file | string | File with GCP credentials, containing AccessID and PrivateKey |  |
| private_key | string | Private key used to sign GCS requests |  |




##### Block `s3`



**Attributes:**

| **Attribute** | **Description** | **Type** | **opt/req?** |
|-----------|-------------|------|----------|
| access_key | string | S3 Access Key |  |
| bucket_location | string | Name of S3 region |  |
| bucket_name | string | Name of the bucket where cache will be stored |  |
| insecure | bool | Use insecure mode (without https) |  |
| secret_key | string | S3 Secret Key |  |
| server_address | string | A host:port to the used S3-compatible server |  |




#### Block `custom`



**Attributes:**

| **Attribute** | **Description** | **Type** | **opt/req?** |
|-----------|-------------|------|----------|
| cleanup_args | [list string] | Arguments for the cleanup executable |  |
| cleanup_exec | string | Executable that cleanups after executor run |  |
| cleanup_exec_timeout | number | Timeout for the cleanup executable (in seconds) |  |
| config_args | [list string] | Arguments for the config executable |  |
| config_exec | string | Executable that allows to inject configuration values to the executor |  |
| config_exec_timeout | number | Timeout for the config executable (in seconds) |  |
| force_kill_timeout | number | Force timeout for scripts execution (in seconds). Counted from the force kill call; if process will be not terminated, Runner will abandon process termination and log an error |  |
| graceful_kill_timeout | number | Graceful timeout for scripts execution after SIGTERM is sent to the process (in seconds). This limits the time given for scripts to perform the cleanup before exiting |  |
| prepare_args | [list string] | Arguments for the prepare executable |  |
| prepare_exec | string | Executable that prepares executor |  |
| prepare_exec_timeout | number | Timeout for the prepare executable (in seconds) |  |
| run_args | [list string] | Arguments for the run executable |  |
| run_exec | string | Executable that runs the job script in executor |  |




#### Block `custom_build_dir`



**Attributes:**

| **Attribute** | **Description** | **Type** | **opt/req?** |
|-----------|-------------|------|----------|
| enabled | bool | Enable job specific build directories |  |




#### Block `docker`



**Attributes:**

| **Attribute** | **Description** | **Type** | **opt/req?** |
|-----------|-------------|------|----------|
| allowed_images | [list string] | Whitelist allowed images |  |
| allowed_services | [list string] | Whitelist allowed services |  |
| cache_dir | string | Directory where to store caches |  |
| cap_add | [list string] | Add Linux capabilities |  |
| cap_drop | [list string] | Drop Linux capabilities |  |
| cpu_shares | number | Number of CPU shares |  |
| cpus | string | Number of CPUs |  |
| cpuset_cpus | string | String value containing the cgroups CpusetCpus to use |  |
| devices | [list string] | Add a host device to the container |  |
| disable_cache | bool | Disable all container caching |  |
| disable_entrypoint_overwrite | bool | Disable the possibility for a container to overwrite the default image entrypoint |  |
| dns | [list string] | A list of DNS servers for the container to use |  |
| dns_search | [list string] | A list of DNS search domains |  |
| extra_hosts | [list string] | Add a custom host-to-IP mapping |  |
| helper_image | string | [ADVANCED] Override the default helper image used to clone repos and upload artifacts |  |
| host | string | Docker daemon address |  |
| hostname | string | Custom container hostname |  |
| image | string | Docker image to be used |  |
| links | [list string] | Add link to another container |  |
| memory | string | Memory limit (format: <number>[<unit>]). Unit can be one of b, k, m, or g. Minimum is 4M. |  |
| memory_reservation | string | Memory soft limit (format: <number>[<unit>]). Unit can be one of b, k, m, or g. |  |
| memory_swap | string | Total memory limit (memory + swap, format: <number>[<unit>]). Unit can be one of b, k, m, or g. |  |
| network_mode | string | Add container to a custom network |  |
| oom_kill_disable | bool | Do not kill processes in a container if an out-of-memory (OOM) error occurs |  |
| oom_score_adjust | number | Adjust OOM score |  |
| privileged | bool | Give extended privileges to container |  |
| pull_policy | string | Image pull policy: never, if-not-present, always |  |
| runtime | string | Docker runtime to be used |  |
| security_opt | [list string] | Security Options |  |
| services_tmpfs | [map string] | A toml table/json object with the format key=values. When set this will mount the specified path in the key as a tmpfs volume in all the service containers, using the options specified as key. For the supported options, see the documentation for the unix 'mount' command |  |
| shm_size | number | Shared memory size for docker images (in bytes) |  |
| sysctls | [map string] | Sysctl options, a toml table/json object of key=value. Value is expected to be a string. |  |
| tls_cert_path | string | Certificate path |  |
| tls_verify | bool | Use TLS and verify the remote |  |
| tmpfs | [map string] | A toml table/json object with the format key=values. When set this will mount the specified path in the key as a tmpfs volume in the main container, using the options specified as key. For the supported options, see the documentation for the unix 'mount' command |  |
| userns_mode | string | User namespace to use |  |
| volume_driver | string | Volume driver to be used |  |
| volumes | [list string] | Bind-mount a volume and create it if it doesn't exist prior to mounting. Can be specified multiple times once per mountpoint, e.g. --docker-volumes 'test0:/test0' --docker-volumes 'test1:/test1' |  |
| volumes_from | [list string] | A list of volumes to inherit from another container |  |
| wait_for_services_timeout | number | How long to wait for service startup |  |




**Blocks:**

This block also takes the following blocks (defined below):

* services


##### Block `services`



**Attributes:**

| **Attribute** | **Description** | **Type** | **opt/req?** |
|-----------|-------------|------|----------|
| alias | string | The alias of the service |  |
| name | string | The image path for the service |  |




#### Block `kubernetes`



**Attributes:**

| **Attribute** | **Description** | **Type** | **opt/req?** |
|-----------|-------------|------|----------|
| bearer_token | string | Optional Kubernetes service account token used to start build pods. |  |
| bearer_token_overwrite_allowed | bool | Bool to authorize builds to specify their own bearer token for creation. |  |
| ca_file | string | Optional Kubernetes master auth ca certificate |  |
| cert_file | string | Optional Kubernetes master auth certificate |  |
| cpu_limit | string | The CPU allocation given to build containers |  |
| cpu_request | string | The CPU allocation requested for build containers |  |
| helper_cpu_limit | string | The CPU allocation given to build helper containers |  |
| helper_cpu_request | string | The CPU allocation requested for build helper containers |  |
| helper_image | string | [ADVANCED] Override the default helper image used to clone repos and upload artifacts |  |
| helper_memory_limit | string | The amount of memory allocated to build helper containers |  |
| helper_memory_request | string | The amount of memory requested for build helper containers |  |
| host | string | Optional Kubernetes master host URL (auto-discovery attempted if not specified) |  |
| image | string | Default docker image to use for builds when none is specified |  |
| image_pull_secrets | [list string] | A list of image pull secrets that are used for pulling docker image |  |
| key_file | string | Optional Kubernetes master auth private key |  |
| memory_limit | string | The amount of memory allocated to build containers |  |
| memory_request | string | The amount of memory requested from build containers |  |
| namespace | string | Namespace to run Kubernetes jobs in |  |
| namespace_overwrite_allowed | string | Regex to validate 'KUBERNETES_NAMESPACE_OVERWRITE' value |  |
| node_selector | [map string] | A toml table/json object of key=value. Value is expected to be a string. When set this will create pods on k8s nodes that match all the key=value pairs. |  |
| node_tolerations | [map string] | A toml table/json object of key=value:effect. Value and effect are expected to be strings. When set, pods will tolerate the given taints. Only one toleration is supported through environment variable configuration. |  |
| pod_annotations | [map string] | A toml table/json object of key-value. Value is expected to be a string. When set, this will create pods with the given annotations. Can be overwritten in build with KUBERNETES_POD_ANNOTATION_* variables |  |
| pod_annotations_overwrite_allowed | string | Regex to validate 'KUBERNETES_POD_ANNOTATIONS_*' values |  |
| pod_labels | [map string] | A toml table/json object of key-value. Value is expected to be a string. When set, this will create pods with the given pod labels. Environment variables will be substituted for values here. |  |
| poll_interval | number | How frequently, in seconds, the runner will poll the Kubernetes pod it has just created to check its status |  |
| poll_timeout | number | The total amount of time, in seconds, that needs to pass before the runner will timeout attempting to connect to the pod it has just created (useful for queueing more builds that the cluster can handle at a time) |  |
| privileged | bool | Run all containers with the privileged flag enabled |  |
| pull_policy | string | Policy for if/when to pull a container image (never, if-not-present, always). The cluster default will be used if not set |  |
| service_account | string | Executor pods will use this Service Account to talk to kubernetes API |  |
| service_account_overwrite_allowed | string | Regex to validate 'KUBERNETES_SERVICE_ACCOUNT' value |  |
| service_cpu_limit | string | The CPU allocation given to build service containers |  |
| service_cpu_request | string | The CPU allocation requested for build service containers |  |
| service_memory_limit | string | The amount of memory allocated to build service containers |  |
| service_memory_request | string | The amount of memory requested for build service containers |  |
| termination_grace_period_seconds | number | Duration after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. |  |




**Blocks:**

This block also takes the following blocks (defined below):

* pod_security_context
* services
* volumes


##### Block `pod_security_context`



**Attributes:**

| **Attribute** | **Description** | **Type** | **opt/req?** |
|-----------|-------------|------|----------|
| fs_group | number | A special supplemental group that applies to all containers in a pod |  |
| run_as_group | number | The GID to run the entrypoint of the container process |  |
| run_as_non_root | bool | Indicates that the container must run as a non-root user |  |
| run_as_user | number | The UID to run the entrypoint of the container process |  |
| supplemental_groups | [list number] | A list of groups applied to the first process run in each container, in addition to the container's primary GID |  |




##### Block `services`



**Attributes:**

| **Attribute** | **Description** | **Type** | **opt/req?** |
|-----------|-------------|------|----------|
| name | string | The image path for the service |  |




##### Block `volumes`



**Attributes:**

This block has no attributes.

**Blocks:**

This block also takes the following blocks (defined below):

* config_map
* empty_dir
* host_path
* pvc
* secret


###### Block `config_map`



**Attributes:**

| **Attribute** | **Description** | **Type** | **opt/req?** |
|-----------|-------------|------|----------|
| items | [map string] | Key-to-path mapping for keys from the config map that should be used. |  |
| mount_path | string | Path where volume should be mounted inside of container |  |
| name | string | The name of the volume and ConfigMap to use |  |
| read_only | bool | If this volume should be mounted read only |  |




###### Block `empty_dir`



**Attributes:**

| **Attribute** | **Description** | **Type** | **opt/req?** |
|-----------|-------------|------|----------|
| medium | string | Set to 'Memory' to have a tmpfs |  |
| mount_path | string | Path where volume should be mounted inside of container |  |
| name | string | The name of the volume and EmptyDir to use |  |




###### Block `host_path`



**Attributes:**

| **Attribute** | **Description** | **Type** | **opt/req?** |
|-----------|-------------|------|----------|
| host_path | string | Path from the host that should be mounted as a volume |  |
| mount_path | string | Path where volume should be mounted inside of container |  |
| name | string | The name of the volume |  |
| read_only | bool | If this volume should be mounted read only |  |




###### Block `pvc`



**Attributes:**

| **Attribute** | **Description** | **Type** | **opt/req?** |
|-----------|-------------|------|----------|
| mount_path | string | Path where volume should be mounted inside of container |  |
| name | string | The name of the volume and PVC to use |  |
| read_only | bool | If this volume should be mounted read only |  |




###### Block `secret`



**Attributes:**

| **Attribute** | **Description** | **Type** | **opt/req?** |
|-----------|-------------|------|----------|
| items | [map string] | Key-to-path mapping for keys from the secret that should be used. |  |
| mount_path | string | Path where volume should be mounted inside of container |  |
| name | string | The name of the volume and Secret to use |  |
| read_only | bool | If this volume should be mounted read only |  |




#### Block `machine`



**Attributes:**

| **Attribute** | **Description** | **Type** | **opt/req?** |
|-----------|-------------|------|----------|
| idle_count | number | Maximum idle machines |  |
| idle_time | number | Minimum time after node can be destroyed |  |
| machine_driver | string | The driver to use when creating machine |  |
| machine_name | string | The template for machine name (needs to include %s) |  |
| machine_options | [list string] | Additional machine creation options |  |
| max_builds | number | Maximum number of builds processed by machine |  |
| off_peak_idle_count | number | Maximum idle machines when the scheduler is in the OffPeak mode |  |
| off_peak_idle_time | number | Minimum time after machine can be destroyed when the scheduler is in the OffPeak mode |  |
| off_peak_periods | [list string] | Time periods when the scheduler is in the OffPeak mode |  |
| off_peak_timezone | string | Timezone for the OffPeak periods (defaults to Local) |  |




#### Block `parallels`



**Attributes:**

| **Attribute** | **Description** | **Type** | **opt/req?** |
|-----------|-------------|------|----------|
| base_name | string | VM name to be used |  |
| disable_snapshots | bool | Disable snapshoting to speedup VM creation |  |
| template_name | string | VM template to be created |  |
| time_server | string | Timeserver to sync the guests time from. Defaults to time.apple.com |  |




#### Block `referees`



**Attributes:**

This block has no attributes.

**Blocks:**

This block also takes the following blocks (defined below):

* metrics


##### Block `metrics`



**Attributes:**

| **Attribute** | **Description** | **Type** | **opt/req?** |
|-----------|-------------|------|----------|
| prometheus_address | string | A host:port to a prometheus metrics server |  |
| queries | [list string] | A list of metrics to query (in PromQL) |  |
| query_interval | number | Query interval (in seconds) |  |




#### Block `ssh`



**Attributes:**

| **Attribute** | **Description** | **Type** | **opt/req?** |
|-----------|-------------|------|----------|
| host | string | Remote host |  |
| identity_file | string | Identity file to be used |  |
| password | string | User password |  |
| port | string | Remote host port |  |
| user | string | User name |  |




#### Block `virtualbox`



**Attributes:**

| **Attribute** | **Description** | **Type** | **opt/req?** |
|-----------|-------------|------|----------|
| base_name | string | VM name to be used |  |
| base_snapshot | string | Name or UUID of a specific VM snapshot to clone |  |
| disable_snapshots | bool | Disable snapshoting to speedup VM creation |  |




### Block `session_server`



**Attributes:**

| **Attribute** | **Description** | **Type** | **opt/req?** |
|-----------|-------------|------|----------|
| advertise_address | string | FIXME |  |
| listen_address | string | FIXME |  |
| session_timeout | number | FIXME |  |


