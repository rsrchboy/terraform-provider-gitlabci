---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "gitlabci_environment Data Source - terraform-provider-gitlabci"
subcategory: ""
description: |-
  
---

# gitlabci_environment (Data Source)





<!-- schema generated by tfplugindocs -->
## Schema

### Read-Only

- `chat_channel` (String) The Source chat channel that triggered the [ChatOps](../chatops/index.md) command.
- `chat_input` (String) The additional arguments passed with the [ChatOps](../chatops/index.md) command.
- `chat_user_id` (String) The chat service's user ID of the user who triggered the [ChatOps](../chatops/index.md) command.
- `ci` (String) Available for all jobs executed in CI/CD. `true` when available.
- `ci_api_v4_url` (String) The GitLab API v4 root URL.
- `ci_builds_dir` (String) The top-level directory where builds are executed.
- `ci_commit_author` (String) The author of the commit in `Name <email>` format.
- `ci_commit_before_sha` (String) The previous latest commit present on a branch. Is always `0000000000000000000000000000000000000000` in pipelines for merge requests.
- `ci_commit_branch` (String) The commit branch name. Available in branch pipelines, including pipelines for the default branch. Not available in merge request pipelines or tag pipelines.
- `ci_commit_description` (String) The description of the commit. If the title is shorter than 100 characters, the message without the first line.
- `ci_commit_message` (String) The full commit message.
- `ci_commit_ref_name` (String) The branch or tag name for which project is built.
- `ci_commit_ref_protected` (String) `true` if the job is running for a protected reference.
- `ci_commit_ref_slug` (String) `CI_COMMIT_REF_NAME` in lowercase, shortened to 63 bytes, and with everything except `0-9` and `a-z` replaced with `-`. No leading / trailing `-`. Use in URLs, host names and domain names.
- `ci_commit_sha` (String) The commit revision the project is built for.
- `ci_commit_short_sha` (String) The first eight characters of `CI_COMMIT_SHA`.
- `ci_commit_tag` (String) The commit tag name. Available only in pipelines for tags.
- `ci_commit_timestamp` (String) The timestamp of the commit in the ISO 8601 format.
- `ci_commit_title` (String) The title of the commit. The full first line of the message.
- `ci_concurrent_id` (String) The unique ID of build execution in a single executor.
- `ci_concurrent_project_id` (String) The unique ID of build execution in a single executor and project.
- `ci_config_path` (String) The path to the CI/CD configuration file. Defaults to `.gitlab-ci.yml`. Read-only inside a running pipeline.
- `ci_debug_trace` (String) `true` if [debug logging (tracing)](index.md#debug-logging) is enabled.
- `ci_default_branch` (String) The name of the project's default branch.
- `ci_dependency_proxy_direct_group_image_prefix` (String) The direct group image prefix for pulling images through the Dependency Proxy.
- `ci_dependency_proxy_group_image_prefix` (String) The top-level group image prefix for pulling images through the Dependency Proxy.
- `ci_dependency_proxy_password` (String) The password to pull images through the Dependency Proxy.
- `ci_dependency_proxy_server` (String) The server for logging in to the Dependency Proxy. This is equivalent to `$CI_SERVER_HOST:$CI_SERVER_PORT`.
- `ci_dependency_proxy_user` (String) The username to pull images through the Dependency Proxy.
- `ci_deploy_freeze` (String) Only available if the pipeline runs during a [deploy freeze window](../../user/project/releases/index.md#prevent-unintentional-releases-by-setting-a-deploy-freeze). `true` when available.
- `ci_deploy_password` (String) The authentication password of the [GitLab Deploy Token](../../user/project/deploy_tokens/index.md#gitlab-deploy-token), if the project has one.
- `ci_deploy_user` (String) The authentication username of the [GitLab Deploy Token](../../user/project/deploy_tokens/index.md#gitlab-deploy-token), if the project has one.
- `ci_disposable_environment` (String) Only available if the job is executed in a disposable environment (something that is created only for this job and disposed of/destroyed after the execution - all executors except `shell` and `ssh`). `true` when available.
- `ci_environment_action` (String) The action annotation specified for this job's environment. Available if [`environment:action`](../yaml/index.md#environmentaction) is set. Can be `start`, `prepare`, or `stop`.
- `ci_environment_name` (String) The name of the environment for this job. Available if [`environment:name`](../yaml/index.md#environmentname) is set.
- `ci_environment_slug` (String) The simplified version of the environment name, suitable for inclusion in DNS, URLs, Kubernetes labels, and so on. Available if [`environment:name`](../yaml/index.md#environmentname) is set. The slug is [truncated to 24 characters](https://gitlab.com/gitlab-org/gitlab/-/issues/20941).
- `ci_environment_tier` (String) The [deployment tier of the environment](../environments/index.md#deployment-tier-of-environments) for this job.
- `ci_environment_url` (String) The URL of the environment for this job. Available if [`environment:url`](../yaml/index.md#environmenturl) is set.
- `ci_external_pull_request_iid` (String) Pull request ID from GitHub.
- `ci_external_pull_request_source_branch_name` (String) The source branch name of the pull request.
- `ci_external_pull_request_source_branch_sha` (String) The HEAD SHA of the source branch of the pull request.
- `ci_external_pull_request_source_repository` (String) The source repository name of the pull request.
- `ci_external_pull_request_target_branch_name` (String) The target branch name of the pull request.
- `ci_external_pull_request_target_branch_sha` (String) The HEAD SHA of the target branch of the pull request.
- `ci_external_pull_request_target_repository` (String) The target repository name of the pull request.
- `ci_has_open_requirements` (String) Only available if the pipeline's project has an open [requirement](../../user/project/requirements/index.md). `true` when available.
- `ci_job_id` (String) The internal ID of the job, unique across all jobs in the GitLab instance.
- `ci_job_image` (String) The name of the Docker image running the job.
- `ci_job_jwt` (String) A RS256 JSON web token to authenticate with third party systems that support JWT authentication, for example [HashiCorp's Vault](../secrets/index.md).
- `ci_job_jwt_v1` (String) The same value as `CI_JOB_JWT`.
- `ci_job_jwt_v2` (String) [**alpha:**](https://about.gitlab.com/handbook/product/gitlab-the-product/#alpha-beta-ga) A newly formatted RS256 JSON web token to increase compatibility. Similar to `CI_JOB_JWT`, except the issuer (`iss`) claim is changed from `gitlab.com` to `https://gitlab.com`, `sub` has changed from `job_id` to a string that contains the project path, and an `aud` claim is added. Format is subject to change. Be aware, the `aud` field is a constant value. Trusting JWTs in multiple relying parties can lead to [one RP sending a JWT to another one and acting maliciously as a job](https://gitlab.com/gitlab-org/gitlab/-/merge_requests/72555#note_769112331).
- `ci_job_manual` (String) `true` if a job was started manually.
- `ci_job_name` (String) The name of the job.
- `ci_job_stage` (String) The name of the job's stage.
- `ci_job_started_at` (String) The UTC datetime when a job started, in [ISO 8601](https://tools.ietf.org/html/rfc3339#appendix-A) format.
- `ci_job_status` (String) The status of the job as each runner stage is executed. Use with [`after_script`](../yaml/index.md#after_script). Can be `success`, `failed`, or `canceled`.
- `ci_job_token` (String) A token to authenticate with [certain API endpoints](../jobs/ci_job_token.md). The token is valid as long as the job is running.
- `ci_job_url` (String) The job details URL.
- `ci_kubernetes_active` (String) Only available if the pipeline has a Kubernetes cluster available for deployments. `true` when available.
- `ci_merge_request_approved` (String) Approval status of the merge request. `true` when [merge request approvals](../../user/project/merge_requests/approvals/index.md) is available and the merge request has been approved.
- `ci_merge_request_assignees` (String) Comma-separated list of usernames of assignees for the merge request.
- `ci_merge_request_diff_base_sha` (String) The base SHA of the merge request diff.
- `ci_merge_request_diff_id` (String) The version of the merge request diff.
- `ci_merge_request_event_type` (String) The event type of the merge request. Can be `detached`, `merged_result` or `merge_train`.
- `ci_merge_request_id` (String) The instance-level ID of the merge request. This is a unique ID across all projects on GitLab.
- `ci_merge_request_iid` (String) The project-level IID (internal ID) of the merge request. This ID is unique for the current project.
- `ci_merge_request_labels` (String) Comma-separated label names of the merge request.
- `ci_merge_request_milestone` (String) The milestone title of the merge request.
- `ci_merge_request_project_id` (String) The ID of the project of the merge request.
- `ci_merge_request_project_path` (String) The path of the project of the merge request. For example `namespace/awesome-project`.
- `ci_merge_request_project_url` (String) The URL of the project of the merge request. For example, `http://192.168.10.15:3000/namespace/awesome-project`.
- `ci_merge_request_ref_path` (String) The ref path of the merge request. For example, `refs/merge-requests/1/head`.
- `ci_merge_request_source_branch_name` (String) The source branch name of the merge request.
- `ci_merge_request_source_branch_sha` (String) The HEAD SHA of the source branch of the merge request. The variable is empty in merge request pipelines. The SHA is present only in [merged results pipelines](../pipelines/pipelines_for_merged_results.md). **(PREMIUM)**
- `ci_merge_request_source_project_id` (String) The ID of the source project of the merge request.
- `ci_merge_request_source_project_path` (String) The path of the source project of the merge request.
- `ci_merge_request_source_project_url` (String) The URL of the source project of the merge request.
- `ci_merge_request_target_branch_name` (String) The target branch name of the merge request.
- `ci_merge_request_target_branch_sha` (String) The HEAD SHA of the target branch of the merge request. The variable is empty in merge request pipelines. The SHA is present only in [merged results pipelines](../pipelines/pipelines_for_merged_results.md). **(PREMIUM)**
- `ci_merge_request_title` (String) The title of the merge request.
- `ci_node_index` (String) The index of the job in the job set. Only available if the job uses [`parallel`](../yaml/index.md#parallel).
- `ci_node_total` (String) The total number of instances of this job running in parallel. Set to `1` if the job does not use [`parallel`](../yaml/index.md#parallel).
- `ci_open_merge_requests` (String) A comma-separated list of up to four merge requests that use the current branch and project as the merge request source. Only available in branch and merge request pipelines if the branch has an associated merge request. For example, `gitlab-org/gitlab!333,gitlab-org/gitlab-foss!11`.
- `ci_pages_domain` (String) The configured domain that hosts GitLab Pages.
- `ci_pages_url` (String) The URL for a GitLab Pages site. Always a subdomain of `CI_PAGES_DOMAIN`.
- `ci_pipeline_created_at` (String) The UTC datetime when the pipeline was created, in [ISO 8601](https://tools.ietf.org/html/rfc3339#appendix-A) format.
- `ci_pipeline_id` (String) The instance-level ID of the current pipeline. This ID is unique across all projects on the GitLab instance.
- `ci_pipeline_iid` (String) The project-level IID (internal ID) of the current pipeline. This ID is unique only within the current project.
- `ci_pipeline_source` (String) How the pipeline was triggered. Can be `push`, `web`, `schedule`, `api`, `external`, `chat`, `webide`, `merge_request_event`, `external_pull_request_event`, `parent_pipeline`, [`trigger`, or `pipeline`](../triggers/index.md#configure-cicd-jobs-to-run-in-triggered-pipelines).
- `ci_pipeline_triggered` (String) `true` if the job was [triggered](../triggers/index.md).
- `ci_pipeline_url` (String) The URL for the pipeline details.
- `ci_project_classification_label` (String) The project [external authorization classification label](../../user/admin_area/settings/external_authorization.md).
- `ci_project_config_path` (String) [Removed](https://gitlab.com/gitlab-org/gitlab/-/issues/322807) in GitLab 14.0. Use `CI_CONFIG_PATH`.
- `ci_project_dir` (String) The full path the repository is cloned to, and where the job runs from. If the GitLab Runner `builds_dir` parameter is set, this variable is set relative to the value of `builds_dir`. For more information, see the [Advanced GitLab Runner configuration](https://docs.gitlab.com/runner/configuration/advanced-configuration.html#the-runners-section).
- `ci_project_id` (String) The ID of the current project. This ID is unique across all projects on the GitLab instance.
- `ci_project_name` (String) The name of the directory for the project. For example if the project URL is `gitlab.example.com/group-name/project-1`, `CI_PROJECT_NAME` is `project-1`.
- `ci_project_namespace` (String) The project namespace (username or group name) of the job.
- `ci_project_path` (String) The project namespace with the project name included.
- `ci_project_path_slug` (String) `$CI_PROJECT_PATH` in lowercase with characters that are not `a-z` or `0-9` replaced with `-` and shortened to 63 bytes. Use in URLs and domain names.
- `ci_project_repository_languages` (String) A comma-separated, lowercase list of the languages used in the repository. For example `ruby,javascript,html,css`.
- `ci_project_root_namespace` (String) The root project namespace (username or group name) of the job. For example, if `CI_PROJECT_NAMESPACE` is `root-group/child-group/grandchild-group`, `CI_PROJECT_ROOT_NAMESPACE` is `root-group`.
- `ci_project_title` (String) The human-readable project name as displayed in the GitLab web interface.
- `ci_project_url` (String) The HTTP(S) address of the project.
- `ci_project_visibility` (String) The project visibility. Can be `internal`, `private`, or `public`.
- `ci_registry` (String) The address of the GitLab Container Registry. Only available if the Container Registry is enabled for the project. This variable includes a `:port` value if one is specified in the registry configuration.
- `ci_registry_image` (String) The address of the project's Container Registry. Only available if the Container Registry is enabled for the project.
- `ci_registry_password` (String) The password to push containers to the project's GitLab Container Registry. Only available if the Container Registry is enabled for the project. This password value is the same as the `CI_JOB_TOKEN` and is valid only as long as the job is running. Use the `CI_DEPLOY_PASSWORD` for long-lived access to the registry
- `ci_registry_user` (String) The username to push containers to the project's GitLab Container Registry. Only available if the Container Registry is enabled for the project.
- `ci_repository_url` (String) The URL to clone the Git repository.
- `ci_runner_description` (String) The description of the runner.
- `ci_runner_executable_arch` (String) The OS/architecture of the GitLab Runner executable. Might not be the same as the environment of the executor.
- `ci_runner_id` (String) The unique ID of the runner being used.
- `ci_runner_revision` (String) The revision of the runner running the job.
- `ci_runner_short_token` (String) First eight characters of the runner's token used to authenticate new job requests. Used as the runner's unique ID.
- `ci_runner_tags` (String) A comma-separated list of the runner tags.
- `ci_runner_version` (String) The version of the GitLab Runner running the job.
- `ci_server` (String) Available for all jobs executed in CI/CD. `yes` when available.
- `ci_server_host` (String) The host of the GitLab instance URL, without protocol or port. For example `gitlab.example.com`.
- `ci_server_name` (String) The name of CI/CD server that coordinates jobs.
- `ci_server_port` (String) The port of the GitLab instance URL, without host or protocol. For example `8080`.
- `ci_server_protocol` (String) The protocol of the GitLab instance URL, without host or port. For example `https`.
- `ci_server_revision` (String) GitLab revision that schedules jobs.
- `ci_server_url` (String) The base URL of the GitLab instance, including protocol and port. For example `https://gitlab.example.com:8080`.
- `ci_server_version` (String) The full version of the GitLab instance.
- `ci_server_version_major` (String) The major version of the GitLab instance. For example, if the GitLab version is `13.6.1`, the `CI_SERVER_VERSION_MAJOR` is `13`.
- `ci_server_version_minor` (String) The minor version of the GitLab instance. For example, if the GitLab version is `13.6.1`, the `CI_SERVER_VERSION_MINOR` is `6`.
- `ci_server_version_patch` (String) The patch version of the GitLab instance. For example, if the GitLab version is `13.6.1`, the `CI_SERVER_VERSION_PATCH` is `1`.
- `ci_shared_environment` (String) Only available if the job is executed in a shared environment (something that is persisted across CI/CD invocations, like the `shell` or `ssh` executor). `true` when available.
- `gitlab_ci` (String) Available for all jobs executed in CI/CD. `true` when available.
- `gitlab_features` (String) The comma-separated list of licensed features available for the GitLab instance and license.
- `gitlab_user_email` (String) The email of the user who started the job.
- `gitlab_user_id` (String) The ID of the user who started the job.
- `gitlab_user_login` (String) The username of the user who started the job.
- `gitlab_user_name` (String) The name of the user who started the job.
- `id` (String) The computed configuation id
- `running_under_ci` (Boolean) True if we appear to be running as a CI job
- `trigger_payload` (String) The webhook payload. Only available when a pipeline is [triggered with a webhook](../triggers/index.md#use-a-webhook-payload).


