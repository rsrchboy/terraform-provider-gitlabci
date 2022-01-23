package provider

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGitlabCIEnvironment() *schema.Resource {

	log.SetFlags(log.Lshortfile)

	// remember: cat ~/Downloads/predefined_variables.md| awk -F\| '{ print $2 $5 }' | perl -nE '/`(\w+)`\s+(.*\S)\s+$/; say q{"} . lc($1) . qq{": {\nType: schema.TypeString,\nComputed: true,\nDescription: "$2",\n},}'

	schema := &schema.Resource{
		ReadContext: dataSourceGitlabCIEnvironmentRead,

		SchemaVersion: 1,
		Schema: map[string]*schema.Schema{
			// internal
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The computed configuation id",
			},
			// generated
			"running_under_ci": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "True if we appear to be running as a CI job",
			},

			// from the environment

			"chat_channel": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The Source chat channel that triggered the [ChatOps](../chatops/index.md) command.",
			},
			"chat_input": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The additional arguments passed with the [ChatOps](../chatops/index.md) command.",
			},
			"chat_user_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The chat service's user ID of the user who triggered the [ChatOps](../chatops/index.md) command.",
			},
			"ci": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Available for all jobs executed in CI/CD. `true` when available.",
			},
			"ci_api_v4_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The GitLab API v4 root URL.",
			},
			"ci_builds_dir": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The top-level directory where builds are executed.",
			},
			"ci_commit_author": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The author of the commit in `Name <email>` format.",
			},
			"ci_commit_before_sha": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The previous latest commit present on a branch. Is always `0000000000000000000000000000000000000000` in pipelines for merge requests.",
			},
			"ci_commit_branch": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The commit branch name. Available in branch pipelines, including pipelines for the default branch. Not available in merge request pipelines or tag pipelines.",
			},
			"ci_commit_description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The description of the commit. If the title is shorter than 100 characters, the message without the first line.",
			},
			"ci_commit_message": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The full commit message.",
			},
			"ci_commit_ref_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The branch or tag name for which project is built.",
			},
			"ci_commit_ref_protected": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "`true` if the job is running for a protected reference.",
			},
			"ci_commit_ref_slug": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "`CI_COMMIT_REF_NAME` in lowercase, shortened to 63 bytes, and with everything except `0-9` and `a-z` replaced with `-`. No leading / trailing `-`. Use in URLs, host names and domain names.",
			},
			"ci_commit_sha": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The commit revision the project is built for.",
			},
			"ci_commit_short_sha": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The first eight characters of `CI_COMMIT_SHA`.",
			},
			"ci_commit_tag": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The commit tag name. Available only in pipelines for tags.",
			},
			"ci_commit_timestamp": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The timestamp of the commit in the ISO 8601 format.",
			},
			"ci_commit_title": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The title of the commit. The full first line of the message.",
			},
			"ci_concurrent_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The unique ID of build execution in a single executor.",
			},
			"ci_concurrent_project_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The unique ID of build execution in a single executor and project.",
			},
			"ci_config_path": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The path to the CI/CD configuration file. Defaults to `.gitlab-ci.yml`. Read-only inside a running pipeline.",
			},
			"ci_debug_trace": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "`true` if [debug logging (tracing)](index.md#debug-logging) is enabled.",
			},
			"ci_default_branch": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the project's default branch.",
			},
			"ci_dependency_proxy_group_image_prefix": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The top-level group image prefix for pulling images through the Dependency Proxy.",
			},
			"ci_dependency_proxy_direct_group_image_prefix": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The direct group image prefix for pulling images through the Dependency Proxy.",
			},
			"ci_dependency_proxy_password": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The password to pull images through the Dependency Proxy.",
			},
			"ci_dependency_proxy_server": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The server for logging in to the Dependency Proxy. This is equivalent to `$CI_SERVER_HOST:$CI_SERVER_PORT`.",
			},
			"ci_dependency_proxy_user": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The username to pull images through the Dependency Proxy.",
			},
			"ci_deploy_freeze": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Only available if the pipeline runs during a [deploy freeze window](../../user/project/releases/index.md#prevent-unintentional-releases-by-setting-a-deploy-freeze). `true` when available.",
			},
			"ci_deploy_password": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The authentication password of the [GitLab Deploy Token](../../user/project/deploy_tokens/index.md#gitlab-deploy-token), if the project has one.",
			},
			"ci_deploy_user": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The authentication username of the [GitLab Deploy Token](../../user/project/deploy_tokens/index.md#gitlab-deploy-token), if the project has one.",
			},
			"ci_disposable_environment": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Only available if the job is executed in a disposable environment (something that is created only for this job and disposed of/destroyed after the execution - all executors except `shell` and `ssh`). `true` when available.",
			},
			"ci_environment_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the environment for this job. Available if [`environment:name`](../yaml/index.md#environmentname) is set.",
			},
			"ci_environment_slug": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The simplified version of the environment name, suitable for inclusion in DNS, URLs, Kubernetes labels, and so on. Available if [`environment:name`](../yaml/index.md#environmentname) is set. The slug is [truncated to 24 characters](https://gitlab.com/gitlab-org/gitlab/-/issues/20941).",
			},
			"ci_environment_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL of the environment for this job. Available if [`environment:url`](../yaml/index.md#environmenturl) is set.",
			},
			"ci_environment_action": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The action annotation specified for this job's environment. Available if [`environment:action`](../yaml/index.md#environmentaction) is set. Can be `start`, `prepare`, or `stop`.",
			},
			"ci_environment_tier": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The [deployment tier of the environment](../environments/index.md#deployment-tier-of-environments) for this job.",
			},
			"ci_has_open_requirements": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Only available if the pipeline's project has an open [requirement](../../user/project/requirements/index.md). `true` when available.",
			},
			"ci_job_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The internal ID of the job, unique across all jobs in the GitLab instance.",
			},
			"ci_job_image": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the Docker image running the job.",
			},
			"ci_job_jwt": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A RS256 JSON web token to authenticate with third party systems that support JWT authentication, for example [HashiCorp's Vault](../secrets/index.md).",
			},
			"ci_job_jwt_v1": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The same value as `CI_JOB_JWT`.",
			},
			"ci_job_jwt_v2": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "[**alpha:**](https://about.gitlab.com/handbook/product/gitlab-the-product/#alpha-beta-ga) A newly formatted RS256 JSON web token to increase compatibility. Similar to `CI_JOB_JWT`, except the issuer (`iss`) claim is changed from `gitlab.com` to `https://gitlab.com`, `sub` has changed from `job_id` to a string that contains the project path, and an `aud` claim is added. Format is subject to change. Be aware, the `aud` field is a constant value. Trusting JWTs in multiple relying parties can lead to [one RP sending a JWT to another one and acting maliciously as a job](https://gitlab.com/gitlab-org/gitlab/-/merge_requests/72555#note_769112331).",
			},
			"ci_job_manual": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "`true` if a job was started manually.",
			},
			"ci_job_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the job.",
			},
			"ci_job_stage": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the job's stage.",
			},
			"ci_job_status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the job as each runner stage is executed. Use with [`after_script`](../yaml/index.md#after_script). Can be `success`, `failed`, or `canceled`.",
			},
			"ci_job_token": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A token to authenticate with [certain API endpoints](../jobs/ci_job_token.md). The token is valid as long as the job is running.",
			},
			"ci_job_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The job details URL.",
			},
			"ci_job_started_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The UTC datetime when a job started, in [ISO 8601](https://tools.ietf.org/html/rfc3339#appendix-A) format.",
			},
			"ci_kubernetes_active": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Only available if the pipeline has a Kubernetes cluster available for deployments. `true` when available.",
			},
			"ci_node_index": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The index of the job in the job set. Only available if the job uses [`parallel`](../yaml/index.md#parallel).",
			},
			"ci_node_total": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The total number of instances of this job running in parallel. Set to `1` if the job does not use [`parallel`](../yaml/index.md#parallel).",
			},
			"ci_open_merge_requests": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A comma-separated list of up to four merge requests that use the current branch and project as the merge request source. Only available in branch and merge request pipelines if the branch has an associated merge request. For example, `gitlab-org/gitlab!333,gitlab-org/gitlab-foss!11`.",
			},
			"ci_pages_domain": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The configured domain that hosts GitLab Pages.",
			},
			"ci_pages_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL for a GitLab Pages site. Always a subdomain of `CI_PAGES_DOMAIN`.",
			},
			"ci_pipeline_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The instance-level ID of the current pipeline. This ID is unique across all projects on the GitLab instance.",
			},
			"ci_pipeline_iid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The project-level IID (internal ID) of the current pipeline. This ID is unique only within the current project.",
			},
			"ci_pipeline_source": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "How the pipeline was triggered. Can be `push`, `web`, `schedule`, `api`, `external`, `chat`, `webide`, `merge_request_event`, `external_pull_request_event`, `parent_pipeline`, [`trigger`, or `pipeline`](../triggers/index.md#configure-cicd-jobs-to-run-in-triggered-pipelines).",
			},
			"ci_pipeline_triggered": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "`true` if the job was [triggered](../triggers/index.md).",
			},
			"ci_pipeline_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL for the pipeline details.",
			},
			"ci_pipeline_created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The UTC datetime when the pipeline was created, in [ISO 8601](https://tools.ietf.org/html/rfc3339#appendix-A) format.",
			},
			"ci_project_config_path": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "[Removed](https://gitlab.com/gitlab-org/gitlab/-/issues/322807) in GitLab 14.0. Use `CI_CONFIG_PATH`.",
			},
			"ci_project_dir": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The full path the repository is cloned to, and where the job runs from. If the GitLab Runner `builds_dir` parameter is set, this variable is set relative to the value of `builds_dir`. For more information, see the [Advanced GitLab Runner configuration](https://docs.gitlab.com/runner/configuration/advanced-configuration.html#the-runners-section).",
			},
			"ci_project_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ID of the current project. This ID is unique across all projects on the GitLab instance.",
			},
			"ci_project_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the directory for the project. For example if the project URL is `gitlab.example.com/group-name/project-1`, `CI_PROJECT_NAME` is `project-1`.",
			},
			"ci_project_namespace": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The project namespace (username or group name) of the job.",
			},
			"ci_project_path_slug": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "`$CI_PROJECT_PATH` in lowercase with characters that are not `a-z` or `0-9` replaced with `-` and shortened to 63 bytes. Use in URLs and domain names.",
			},
			"ci_project_path": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The project namespace with the project name included.",
			},
			"ci_project_repository_languages": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A comma-separated, lowercase list of the languages used in the repository. For example `ruby,javascript,html,css`.",
			},
			"ci_project_root_namespace": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The root project namespace (username or group name) of the job. For example, if `CI_PROJECT_NAMESPACE` is `root-group/child-group/grandchild-group`, `CI_PROJECT_ROOT_NAMESPACE` is `root-group`.",
			},
			"ci_project_title": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The human-readable project name as displayed in the GitLab web interface.",
			},
			"ci_project_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The HTTP(S) address of the project.",
			},
			"ci_project_visibility": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The project visibility. Can be `internal`, `private`, or `public`.",
			},
			"ci_project_classification_label": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The project [external authorization classification label](../../user/admin_area/settings/external_authorization.md).",
			},
			"ci_registry_image": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The address of the project's Container Registry. Only available if the Container Registry is enabled for the project.",
			},
			"ci_registry_password": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The password to push containers to the project's GitLab Container Registry. Only available if the Container Registry is enabled for the project. This password value is the same as the `CI_JOB_TOKEN` and is valid only as long as the job is running. Use the `CI_DEPLOY_PASSWORD` for long-lived access to the registry",
			},
			"ci_registry_user": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The username to push containers to the project's GitLab Container Registry. Only available if the Container Registry is enabled for the project.",
			},
			"ci_registry": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The address of the GitLab Container Registry. Only available if the Container Registry is enabled for the project. This variable includes a `:port` value if one is specified in the registry configuration.",
			},
			"ci_repository_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL to clone the Git repository.",
			},
			"ci_runner_description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The description of the runner.",
			},
			"ci_runner_executable_arch": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The OS/architecture of the GitLab Runner executable. Might not be the same as the environment of the executor.",
			},
			"ci_runner_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The unique ID of the runner being used.",
			},
			"ci_runner_revision": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The revision of the runner running the job.",
			},
			"ci_runner_short_token": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "First eight characters of the runner's token used to authenticate new job requests. Used as the runner's unique ID.",
			},
			"ci_runner_tags": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A comma-separated list of the runner tags.",
			},
			"ci_runner_version": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The version of the GitLab Runner running the job.",
			},
			"ci_server_host": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The host of the GitLab instance URL, without protocol or port. For example `gitlab.example.com`.",
			},
			"ci_server_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of CI/CD server that coordinates jobs.",
			},
			"ci_server_port": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The port of the GitLab instance URL, without host or protocol. For example `8080`.",
			},
			"ci_server_protocol": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The protocol of the GitLab instance URL, without host or port. For example `https`.",
			},
			"ci_server_revision": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "GitLab revision that schedules jobs.",
			},
			"ci_server_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The base URL of the GitLab instance, including protocol and port. For example `https://gitlab.example.com:8080`.",
			},
			"ci_server_version_major": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The major version of the GitLab instance. For example, if the GitLab version is `13.6.1`, the `CI_SERVER_VERSION_MAJOR` is `13`.",
			},
			"ci_server_version_minor": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The minor version of the GitLab instance. For example, if the GitLab version is `13.6.1`, the `CI_SERVER_VERSION_MINOR` is `6`.",
			},
			"ci_server_version_patch": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The patch version of the GitLab instance. For example, if the GitLab version is `13.6.1`, the `CI_SERVER_VERSION_PATCH` is `1`.",
			},
			"ci_server_version": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The full version of the GitLab instance.",
			},
			"ci_server": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Available for all jobs executed in CI/CD. `yes` when available.",
			},
			"ci_shared_environment": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Only available if the job is executed in a shared environment (something that is persisted across CI/CD invocations, like the `shell` or `ssh` executor). `true` when available.",
			},
			"gitlab_ci": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Available for all jobs executed in CI/CD. `true` when available.",
			},
			"gitlab_features": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The comma-separated list of licensed features available for the GitLab instance and license.",
			},
			"gitlab_user_email": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The email of the user who started the job.",
			},
			"gitlab_user_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ID of the user who started the job.",
			},
			"gitlab_user_login": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The username of the user who started the job.",
			},
			"gitlab_user_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the user who started the job.",
			},
			"trigger_payload": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The webhook payload. Only available when a pipeline is [triggered with a webhook](../triggers/index.md#use-a-webhook-payload).",
			},
			"ci_merge_request_approved": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Approval status of the merge request. `true` when [merge request approvals](../../user/project/merge_requests/approvals/index.md) is available and the merge request has been approved.",
			},
			"ci_merge_request_assignees": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Comma-separated list of usernames of assignees for the merge request.",
			},
			"ci_merge_request_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The instance-level ID of the merge request. This is a unique ID across all projects on GitLab.",
			},
			"ci_merge_request_iid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The project-level IID (internal ID) of the merge request. This ID is unique for the current project.",
			},
			"ci_merge_request_labels": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Comma-separated label names of the merge request.",
			},
			"ci_merge_request_milestone": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The milestone title of the merge request.",
			},
			"ci_merge_request_project_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ID of the project of the merge request.",
			},
			"ci_merge_request_project_path": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The path of the project of the merge request. For example `namespace/awesome-project`.",
			},
			"ci_merge_request_project_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL of the project of the merge request. For example, `http://192.168.10.15:3000/namespace/awesome-project`.",
			},
			"ci_merge_request_ref_path": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ref path of the merge request. For example, `refs/merge-requests/1/head`.",
			},
			"ci_merge_request_source_branch_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The source branch name of the merge request.",
			},
			"ci_merge_request_source_branch_sha": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The HEAD SHA of the source branch of the merge request. The variable is empty in merge request pipelines. The SHA is present only in [merged results pipelines](../pipelines/pipelines_for_merged_results.md). **(PREMIUM)**",
			},
			"ci_merge_request_source_project_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ID of the source project of the merge request.",
			},
			"ci_merge_request_source_project_path": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The path of the source project of the merge request.",
			},
			"ci_merge_request_source_project_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL of the source project of the merge request.",
			},
			"ci_merge_request_target_branch_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The target branch name of the merge request.",
			},
			"ci_merge_request_target_branch_sha": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The HEAD SHA of the target branch of the merge request. The variable is empty in merge request pipelines. The SHA is present only in [merged results pipelines](../pipelines/pipelines_for_merged_results.md). **(PREMIUM)**",
			},
			"ci_merge_request_title": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The title of the merge request.",
			},
			"ci_merge_request_event_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The event type of the merge request. Can be `detached`, `merged_result` or `merge_train`.",
			},
			"ci_merge_request_diff_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The version of the merge request diff.",
			},
			"ci_merge_request_diff_base_sha": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The base SHA of the merge request diff.",
			},
			"ci_external_pull_request_iid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Pull request ID from GitHub.",
			},
			"ci_external_pull_request_source_repository": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The source repository name of the pull request.",
			},
			"ci_external_pull_request_target_repository": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The target repository name of the pull request.",
			},
			"ci_external_pull_request_source_branch_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The source branch name of the pull request.",
			},
			"ci_external_pull_request_source_branch_sha": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The HEAD SHA of the source branch of the pull request.",
			},
			"ci_external_pull_request_target_branch_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The target branch name of the pull request.",
			},
			"ci_external_pull_request_target_branch_sha": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The HEAD SHA of the target branch of the pull request.",
			},
		},
	}

	// log.Printf("[TRACE] generated schema is: %s", spew.Sdump(schema))
	return schema
}

func dataSourceGitlabCIEnvironmentRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	tflog.Trace(ctx, "dataSourceGitlabCIEnvironmentRead() (mark III)")

	// This is effectively a no-op: we'd be computing these values from the
	// environment, but we've also asked that the default values be set from
	// where we'd look in the environment.

	if d.Get("ci_server").(string) == "yes" {
		d.Set("running_under_ci", true)
	} else {
		d.Set("running_under_ci", false)
	}

	// cat ~/Downloads/predefined_variables.md| perl -naF'\|' -E 's/(\s+$|^\s+)//g for @F; $F[1] =~ s/`//g; say qq{d.Set("} . lc($F[1]). qq{", os.Getenv("$F[1]"))}'

	d.Set("chat_channel", os.Getenv("CHAT_CHANNEL"))
	d.Set("chat_input", os.Getenv("CHAT_INPUT"))
	d.Set("chat_user_id", os.Getenv("CHAT_USER_ID"))
	d.Set("ci", os.Getenv("CI"))
	d.Set("ci_api_v4_url", os.Getenv("CI_API_V4_URL"))
	d.Set("ci_builds_dir", os.Getenv("CI_BUILDS_DIR"))
	d.Set("ci_commit_author", os.Getenv("CI_COMMIT_AUTHOR"))
	d.Set("ci_commit_before_sha", os.Getenv("CI_COMMIT_BEFORE_SHA"))
	d.Set("ci_commit_branch", os.Getenv("CI_COMMIT_BRANCH"))
	d.Set("ci_commit_description", os.Getenv("CI_COMMIT_DESCRIPTION"))
	d.Set("ci_commit_message", os.Getenv("CI_COMMIT_MESSAGE"))
	d.Set("ci_commit_ref_name", os.Getenv("CI_COMMIT_REF_NAME"))
	d.Set("ci_commit_ref_protected", os.Getenv("CI_COMMIT_REF_PROTECTED"))
	d.Set("ci_commit_ref_slug", os.Getenv("CI_COMMIT_REF_SLUG"))
	d.Set("ci_commit_sha", os.Getenv("CI_COMMIT_SHA"))
	d.Set("ci_commit_short_sha", os.Getenv("CI_COMMIT_SHORT_SHA"))
	d.Set("ci_commit_tag", os.Getenv("CI_COMMIT_TAG"))
	d.Set("ci_commit_timestamp", os.Getenv("CI_COMMIT_TIMESTAMP"))
	d.Set("ci_commit_title", os.Getenv("CI_COMMIT_TITLE"))
	d.Set("ci_concurrent_id", os.Getenv("CI_CONCURRENT_ID"))
	d.Set("ci_concurrent_project_id", os.Getenv("CI_CONCURRENT_PROJECT_ID"))
	d.Set("ci_config_path", os.Getenv("CI_CONFIG_PATH"))
	d.Set("ci_debug_trace", os.Getenv("CI_DEBUG_TRACE"))
	d.Set("ci_default_branch", os.Getenv("CI_DEFAULT_BRANCH"))
	d.Set("ci_dependency_proxy_group_image_prefix", os.Getenv("CI_DEPENDENCY_PROXY_GROUP_IMAGE_PREFIX"))
	d.Set("ci_dependency_proxy_direct_group_image_prefix", os.Getenv("CI_DEPENDENCY_PROXY_DIRECT_GROUP_IMAGE_PREFIX"))
	d.Set("ci_dependency_proxy_password", os.Getenv("CI_DEPENDENCY_PROXY_PASSWORD"))
	d.Set("ci_dependency_proxy_server", os.Getenv("CI_DEPENDENCY_PROXY_SERVER"))
	d.Set("ci_dependency_proxy_user", os.Getenv("CI_DEPENDENCY_PROXY_USER"))
	d.Set("ci_deploy_freeze", os.Getenv("CI_DEPLOY_FREEZE"))
	d.Set("ci_deploy_password", os.Getenv("CI_DEPLOY_PASSWORD"))
	d.Set("ci_deploy_user", os.Getenv("CI_DEPLOY_USER"))
	d.Set("ci_disposable_environment", os.Getenv("CI_DISPOSABLE_ENVIRONMENT"))
	d.Set("ci_environment_name", os.Getenv("CI_ENVIRONMENT_NAME"))
	d.Set("ci_environment_slug", os.Getenv("CI_ENVIRONMENT_SLUG"))
	d.Set("ci_environment_url", os.Getenv("CI_ENVIRONMENT_URL"))
	d.Set("ci_environment_action", os.Getenv("CI_ENVIRONMENT_ACTION"))
	d.Set("ci_environment_tier", os.Getenv("CI_ENVIRONMENT_TIER"))
	d.Set("ci_has_open_requirements", os.Getenv("CI_HAS_OPEN_REQUIREMENTS"))
	d.Set("ci_job_id", os.Getenv("CI_JOB_ID"))
	d.Set("ci_job_image", os.Getenv("CI_JOB_IMAGE"))
	d.Set("ci_job_jwt", os.Getenv("CI_JOB_JWT"))
	d.Set("ci_job_jwt_v1", os.Getenv("CI_JOB_JWT_V1"))
	d.Set("ci_job_jwt_v2", os.Getenv("CI_JOB_JWT_V2"))
	d.Set("ci_job_manual", os.Getenv("CI_JOB_MANUAL"))
	d.Set("ci_job_name", os.Getenv("CI_JOB_NAME"))
	d.Set("ci_job_stage", os.Getenv("CI_JOB_STAGE"))
	d.Set("ci_job_status", os.Getenv("CI_JOB_STATUS"))
	d.Set("ci_job_token", os.Getenv("CI_JOB_TOKEN"))
	d.Set("ci_job_url", os.Getenv("CI_JOB_URL"))
	d.Set("ci_job_started_at", os.Getenv("CI_JOB_STARTED_AT"))
	d.Set("ci_kubernetes_active", os.Getenv("CI_KUBERNETES_ACTIVE"))
	d.Set("ci_node_index", os.Getenv("CI_NODE_INDEX"))
	d.Set("ci_node_total", os.Getenv("CI_NODE_TOTAL"))
	d.Set("ci_open_merge_requests", os.Getenv("CI_OPEN_MERGE_REQUESTS"))
	d.Set("ci_pages_domain", os.Getenv("CI_PAGES_DOMAIN"))
	d.Set("ci_pages_url", os.Getenv("CI_PAGES_URL"))
	d.Set("ci_pipeline_id", os.Getenv("CI_PIPELINE_ID"))
	d.Set("ci_pipeline_iid", os.Getenv("CI_PIPELINE_IID"))
	d.Set("ci_pipeline_source", os.Getenv("CI_PIPELINE_SOURCE"))
	d.Set("ci_pipeline_triggered", os.Getenv("CI_PIPELINE_TRIGGERED"))
	d.Set("ci_pipeline_url", os.Getenv("CI_PIPELINE_URL"))
	d.Set("ci_pipeline_created_at", os.Getenv("CI_PIPELINE_CREATED_AT"))
	d.Set("ci_project_config_path", os.Getenv("CI_PROJECT_CONFIG_PATH"))
	d.Set("ci_project_dir", os.Getenv("CI_PROJECT_DIR"))
	d.Set("ci_project_id", os.Getenv("CI_PROJECT_ID"))
	d.Set("ci_project_name", os.Getenv("CI_PROJECT_NAME"))
	d.Set("ci_project_namespace", os.Getenv("CI_PROJECT_NAMESPACE"))
	d.Set("ci_project_path_slug", os.Getenv("CI_PROJECT_PATH_SLUG"))
	d.Set("ci_project_path", os.Getenv("CI_PROJECT_PATH"))
	d.Set("ci_project_repository_languages", os.Getenv("CI_PROJECT_REPOSITORY_LANGUAGES"))
	d.Set("ci_project_root_namespace", os.Getenv("CI_PROJECT_ROOT_NAMESPACE"))
	d.Set("ci_project_title", os.Getenv("CI_PROJECT_TITLE"))
	d.Set("ci_project_url", os.Getenv("CI_PROJECT_URL"))
	d.Set("ci_project_visibility", os.Getenv("CI_PROJECT_VISIBILITY"))
	d.Set("ci_project_classification_label", os.Getenv("CI_PROJECT_CLASSIFICATION_LABEL"))
	d.Set("ci_registry_image", os.Getenv("CI_REGISTRY_IMAGE"))
	d.Set("ci_registry_password", os.Getenv("CI_REGISTRY_PASSWORD"))
	d.Set("ci_registry_user", os.Getenv("CI_REGISTRY_USER"))
	d.Set("ci_registry", os.Getenv("CI_REGISTRY"))
	d.Set("ci_repository_url", os.Getenv("CI_REPOSITORY_URL"))
	d.Set("ci_runner_description", os.Getenv("CI_RUNNER_DESCRIPTION"))
	d.Set("ci_runner_executable_arch", os.Getenv("CI_RUNNER_EXECUTABLE_ARCH"))
	d.Set("ci_runner_id", os.Getenv("CI_RUNNER_ID"))
	d.Set("ci_runner_revision", os.Getenv("CI_RUNNER_REVISION"))
	d.Set("ci_runner_short_token", os.Getenv("CI_RUNNER_SHORT_TOKEN"))
	d.Set("ci_runner_tags", os.Getenv("CI_RUNNER_TAGS"))
	d.Set("ci_runner_version", os.Getenv("CI_RUNNER_VERSION"))
	d.Set("ci_server_host", os.Getenv("CI_SERVER_HOST"))
	d.Set("ci_server_name", os.Getenv("CI_SERVER_NAME"))
	d.Set("ci_server_port", os.Getenv("CI_SERVER_PORT"))
	d.Set("ci_server_protocol", os.Getenv("CI_SERVER_PROTOCOL"))
	d.Set("ci_server_revision", os.Getenv("CI_SERVER_REVISION"))
	d.Set("ci_server_url", os.Getenv("CI_SERVER_URL"))
	d.Set("ci_server_version_major", os.Getenv("CI_SERVER_VERSION_MAJOR"))
	d.Set("ci_server_version_minor", os.Getenv("CI_SERVER_VERSION_MINOR"))
	d.Set("ci_server_version_patch", os.Getenv("CI_SERVER_VERSION_PATCH"))
	d.Set("ci_server_version", os.Getenv("CI_SERVER_VERSION"))
	d.Set("ci_server", os.Getenv("CI_SERVER"))
	d.Set("ci_shared_environment", os.Getenv("CI_SHARED_ENVIRONMENT"))
	d.Set("gitlab_ci", os.Getenv("GITLAB_CI"))
	d.Set("gitlab_features", os.Getenv("GITLAB_FEATURES"))
	d.Set("gitlab_user_email", os.Getenv("GITLAB_USER_EMAIL"))
	d.Set("gitlab_user_id", os.Getenv("GITLAB_USER_ID"))
	d.Set("gitlab_user_login", os.Getenv("GITLAB_USER_LOGIN"))
	d.Set("gitlab_user_name", os.Getenv("GITLAB_USER_NAME"))
	d.Set("trigger_payload", os.Getenv("TRIGGER_PAYLOAD"))
	d.Set("ci_merge_request_approved", os.Getenv("CI_MERGE_REQUEST_APPROVED"))
	d.Set("ci_merge_request_assignees", os.Getenv("CI_MERGE_REQUEST_ASSIGNEES"))
	d.Set("ci_merge_request_id", os.Getenv("CI_MERGE_REQUEST_ID"))
	d.Set("ci_merge_request_iid", os.Getenv("CI_MERGE_REQUEST_IID"))
	d.Set("ci_merge_request_labels", os.Getenv("CI_MERGE_REQUEST_LABELS"))
	d.Set("ci_merge_request_milestone", os.Getenv("CI_MERGE_REQUEST_MILESTONE"))
	d.Set("ci_merge_request_project_id", os.Getenv("CI_MERGE_REQUEST_PROJECT_ID"))
	d.Set("ci_merge_request_project_path", os.Getenv("CI_MERGE_REQUEST_PROJECT_PATH"))
	d.Set("ci_merge_request_project_url", os.Getenv("CI_MERGE_REQUEST_PROJECT_URL"))
	d.Set("ci_merge_request_ref_path", os.Getenv("CI_MERGE_REQUEST_REF_PATH"))
	d.Set("ci_merge_request_source_branch_name", os.Getenv("CI_MERGE_REQUEST_SOURCE_BRANCH_NAME"))
	d.Set("ci_merge_request_source_branch_sha", os.Getenv("CI_MERGE_REQUEST_SOURCE_BRANCH_SHA"))
	d.Set("ci_merge_request_source_project_id", os.Getenv("CI_MERGE_REQUEST_SOURCE_PROJECT_ID"))
	d.Set("ci_merge_request_source_project_path", os.Getenv("CI_MERGE_REQUEST_SOURCE_PROJECT_PATH"))
	d.Set("ci_merge_request_source_project_url", os.Getenv("CI_MERGE_REQUEST_SOURCE_PROJECT_URL"))
	d.Set("ci_merge_request_target_branch_name", os.Getenv("CI_MERGE_REQUEST_TARGET_BRANCH_NAME"))
	d.Set("ci_merge_request_target_branch_sha", os.Getenv("CI_MERGE_REQUEST_TARGET_BRANCH_SHA"))
	d.Set("ci_merge_request_title", os.Getenv("CI_MERGE_REQUEST_TITLE"))
	d.Set("ci_merge_request_event_type", os.Getenv("CI_MERGE_REQUEST_EVENT_TYPE"))
	d.Set("ci_merge_request_diff_id", os.Getenv("CI_MERGE_REQUEST_DIFF_ID"))
	d.Set("ci_merge_request_diff_base_sha", os.Getenv("CI_MERGE_REQUEST_DIFF_BASE_SHA"))
	d.Set("ci_external_pull_request_iid", os.Getenv("CI_EXTERNAL_PULL_REQUEST_IID"))
	d.Set("ci_external_pull_request_source_repository", os.Getenv("CI_EXTERNAL_PULL_REQUEST_SOURCE_REPOSITORY"))
	d.Set("ci_external_pull_request_target_repository", os.Getenv("CI_EXTERNAL_PULL_REQUEST_TARGET_REPOSITORY"))
	d.Set("ci_external_pull_request_source_branch_name", os.Getenv("CI_EXTERNAL_PULL_REQUEST_SOURCE_BRANCH_NAME"))
	d.Set("ci_external_pull_request_source_branch_sha", os.Getenv("CI_EXTERNAL_PULL_REQUEST_SOURCE_BRANCH_SHA"))
	d.Set("ci_external_pull_request_target_branch_name", os.Getenv("CI_EXTERNAL_PULL_REQUEST_TARGET_BRANCH_NAME"))
	d.Set("ci_external_pull_request_target_branch_sha", os.Getenv("CI_EXTERNAL_PULL_REQUEST_TARGET_BRANCH_SHA"))

	tflog.Trace(ctx, "dataSourceGitlabCIEnvironmentRead() finished")
	return nil
}
