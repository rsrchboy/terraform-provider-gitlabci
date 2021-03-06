package gitlabci

import (
	"log"
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceGitlabCIEnvironment() *schema.Resource {

	log.SetFlags(log.Lshortfile)

	// remember: cat ~/Downloads/predefined_variables.md| awk -F\| '{ print $2 $5 }' | perl -nE '/`(\w+)`\s+(.*\S)\s+$/; say q{"} . lc($1) . qq{": {\nType: schema.TypeString,\nComputed: true,\nDescription: "$2",\n},}'

	schema := &schema.Resource{
		Read: dataSourceGitlabCIEnvironmentRead,

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
			"artifact_download_attempts": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Number of attempts to download artifacts running a job",
			},
			"chat_channel": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Source chat channel which triggered the [ChatOps](../chatops/README.md) command",
			},
			"chat_input": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Additional arguments passed in the [ChatOps](../chatops/README.md) command",
			},
			"ci": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Mark that job is executed in CI environment",
			},
			"ci_api_v4_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The GitLab API v4 root URL",
			},
			"ci_builds_dir": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Top-level directory where builds are executed.",
			},
			"ci_commit_before_sha": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The previous latest commit present on a branch before a merge request. Only populated when there is a merge request associated with the pipeline.",
			},
			"ci_commit_description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The description of the commit: the message without first line, if the title is shorter than 100 characters; full message in other case.",
			},
			"ci_commit_message": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The full commit message.",
			},
			"ci_commit_ref_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The branch or tag name for which project is built",
			},
			"ci_commit_ref_protected": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "If the job is running on a protected branch",
			},
			"ci_commit_ref_slug": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "`$CI_COMMIT_REF_NAME` lowercased, shortened to 63 bytes, and with everything except `0-9` and `a-z` replaced with `-`. No leading / trailing `-`. Use in URLs, host names and domain names.",
			},
			"ci_commit_sha": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The commit revision for which project is built",
			},
			"ci_commit_short_sha": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The first eight characters of `CI_COMMIT_SHA`",
			},
			"ci_commit_branch": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The commit branch name. Present only when building branches.",
			},
			"ci_commit_tag": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The commit tag name. Present only when building tags.",
			},
			"ci_commit_title": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The title of the commit - the full first line of the message",
			},
			"ci_concurrent_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Unique ID of build execution within a single executor.",
			},
			"ci_concurrent_project_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Unique ID of build execution within a single executor and project.",
			},
			"ci_config_path": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The path to CI config file. Defaults to `.gitlab-ci.yml`",
			},
			"ci_debug_trace": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Whether [debug logging (tracing)](README.md#debug-logging) is enabled",
			},
			"ci_default_branch": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the default branch for the project.",
			},
			"ci_deploy_password": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Authentication password of the [GitLab Deploy Token][gitlab-deploy-token], only present if the Project has one related.",
			},
			"ci_deploy_user": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Authentication username of the [GitLab Deploy Token][gitlab-deploy-token], only present if the Project has one related.",
			},
			"ci_disposable_environment": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Marks that the job is executed in a disposable environment (something that is created only for this job and disposed of/destroyed after the execution - all executors except `shell` and `ssh`). If the environment is disposable, it is set to true, otherwise it is not defined at all.",
			},
			"ci_environment_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the environment for this job. Only present if [`environment:name`](../yaml/README.md#environmentname) is set.",
			},
			"ci_environment_slug": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A simplified version of the environment name, suitable for inclusion in DNS, URLs, Kubernetes labels, etc. Only present if [`environment:name`](../yaml/README.md#environmentname) is set.",
			},
			"ci_environment_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL of the environment for this job. Only present if [`environment:url`](../yaml/README.md#environmenturl) is set.",
			},
			"ci_external_pull_request_iid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Pull Request ID from GitHub if the [pipelines are for external pull requests](../ci_cd_for_external_repos/index.md#pipelines-for-external-pull-requests). Available only if `only: [external_pull_requests]` is used and the pull request is open.",
			},
			"ci_external_pull_request_source_branch_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The source branch name of the pull request if [the pipelines are for external pull requests](../ci_cd_for_external_repos/index.md#pipelines-for-external-pull-requests). Available only if `only: [external_pull_requests]` is used and the pull request is open.",
			},
			"ci_external_pull_request_source_branch_sha": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The HEAD SHA of the source branch of the pull request if [the pipelines are for external pull requests](../ci_cd_for_external_repos/index.md#pipelines-for-external-pull-requests). Available only if `only: [external_pull_requests]` is used and the pull request is open.",
			},
			"ci_external_pull_request_target_branch_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The target branch name of the pull request if [the pipelines are for external pull requests](../ci_cd_for_external_repos/index.md#pipelines-for-external-pull-requests). Available only if `only: [external_pull_requests]` is used and the pull request is open.",
			},
			"ci_external_pull_request_target_branch_sha": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The HEAD SHA of the target branch of the pull request if [the pipelines are for external pull requests](../ci_cd_for_external_repos/index.md#pipelines-for-external-pull-requests). Available only if `only: [external_pull_requests]` is used and the pull request is open.",
			},
			"ci_job_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The unique id of the current job that GitLab CI uses internally",
			},
			"ci_job_manual": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The flag to indicate that job was manually started",
			},
			"ci_job_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the job as defined in `.gitlab-ci.yml`",
			},
			"ci_job_stage": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the stage as defined in `.gitlab-ci.yml`",
			},
			"ci_job_token": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Token used for authenticating with the [GitLab Container Registry][registry] and downloading [dependent repositories][dependent-repositories]",
			},
			"ci_job_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Job details URL",
			},
			"ci_merge_request_assignees": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Comma-separated list of username(s) of assignee(s) for the merge request if [the pipelines are for merge requests](../merge_request_pipelines/index.md). Available only if `only: [merge_requests]` is used and the merge request is created.",
			},
			"ci_merge_request_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ID of the merge request if [the pipelines are for merge requests](../merge_request_pipelines/index.md). Available only if `only: [merge_requests]` is used and the merge request is created.",
			},
			"ci_merge_request_iid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The IID of the merge request if [the pipelines are for merge requests](../merge_request_pipelines/index.md). Available only if `only: [merge_requests]` is used and the merge request is created.",
			},
			"ci_merge_request_labels": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Comma-separated label names of the merge request if [the pipelines are for merge requests](../merge_request_pipelines/index.md). Available only if `only: [merge_requests]` is used and the merge request is created.",
			},
			"ci_merge_request_milestone": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The milestone title of the merge request if [the pipelines are for merge requests](../merge_request_pipelines/index.md). Available only if `only: [merge_requests]` is used and the merge request is created.",
			},
			"ci_merge_request_project_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ID of the project of the merge request if [the pipelines are for merge requests](../merge_request_pipelines/index.md). Available only if `only: [merge_requests]` is used and the merge request is created.",
			},
			"ci_merge_request_project_path": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The path of the project of the merge request if [the pipelines are for merge requests](../merge_request_pipelines/index.md) (e.g. `namespace/awesome-project`). Available only if `only: [merge_requests]` is used and the merge request is created.",
			},
			"ci_merge_request_project_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL of the project of the merge request if [the pipelines are for merge requests](../merge_request_pipelines/index.md) (e.g. `http://192.168.10.15:3000/namespace/awesome-project`). Available only if `only: [merge_requests]` is used and the merge request is created.",
			},
			"ci_merge_request_ref_path": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ref path of the merge request if [the pipelines are for merge requests](../merge_request_pipelines/index.md). (e.g. `refs/merge-requests/1/head`). Available only if `only: [merge_requests]` is used and the merge request is created.",
			},
			"ci_merge_request_source_branch_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The source branch name of the merge request if [the pipelines are for merge requests](../merge_request_pipelines/index.md). Available only if `only: [merge_requests]` is used and the merge request is created.",
			},
			"ci_merge_request_source_branch_sha": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The HEAD SHA of the source branch of the merge request if [the pipelines are for merge requests](../merge_request_pipelines/index.md). Available only if `only: [merge_requests]` is used, the merge request is created, and the pipeline is a [merged result pipeline](../merge_request_pipelines/pipelines_for_merged_results/index.md). **(PREMIUM)**",
			},
			"ci_merge_request_source_project_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ID of the source project of the merge request if [the pipelines are for merge requests](../merge_request_pipelines/index.md). Available only if `only: [merge_requests]` is used and the merge request is created.",
			},
			"ci_merge_request_source_project_path": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The path of the source project of the merge request if [the pipelines are for merge requests](../merge_request_pipelines/index.md). Available only if `only: [merge_requests]` is used and the merge request is created.",
			},
			"ci_merge_request_source_project_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL of the source project of the merge request if [the pipelines are for merge requests](../merge_request_pipelines/index.md). Available only if `only: [merge_requests]` is used and the merge request is created.",
			},
			"ci_merge_request_target_branch_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The target branch name of the merge request if [the pipelines are for merge requests](../merge_request_pipelines/index.md). Available only if `only: [merge_requests]` is used and the merge request is created.",
			},
			"ci_merge_request_target_branch_sha": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The HEAD SHA of the target branch of the merge request if [the pipelines are for merge requests](../merge_request_pipelines/index.md). Available only if `only: [merge_requests]` is used, the merge request is created, and the pipeline is a [merged result pipeline](../merge_request_pipelines/pipelines_for_merged_results/index.md). **(PREMIUM)**",
			},
			"ci_merge_request_title": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The title of the merge request if [the pipelines are for merge requests](../merge_request_pipelines/index.md). Available only if `only: [merge_requests]` is used and the merge request is created.",
			},
			"ci_merge_request_event_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The event type of the merge request, if [the pipelines are for merge requests](../merge_request_pipelines/index.md). Can be `detached`, `merged_result` or `merge_train`.",
			},
			"ci_node_index": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Index of the job in the job set. If the job is not parallelized, this variable is not set.",
			},
			"ci_node_total": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Total number of instances of this job running in parallel. If the job is not parallelized, this variable is set to `1`.",
			},
			"ci_pages_domain": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The configured domain that hosts GitLab Pages.",
			},
			"ci_pages_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "URL to GitLab Pages-built pages. Always belongs to a subdomain of `CI_PAGES_DOMAIN`.",
			},
			"ci_pipeline_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The unique id of the current pipeline that GitLab CI uses internally",
			},
			"ci_pipeline_iid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The unique id of the current pipeline scoped to project",
			},
			"ci_pipeline_source": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Indicates how the pipeline was triggered. Possible options are: `push`, `web`, `trigger`, `schedule`, `api`, `pipeline`, `external`, `chat`, `merge_request_event`, and `external_pull_request_event`. For pipelines created before GitLab 9.5, this will show as `unknown`",
			},
			"ci_pipeline_triggered": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The flag to indicate that job was [triggered](../triggers/README.md)",
			},
			"ci_pipeline_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Pipeline details URL",
			},
			"ci_project_dir": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The full path where the repository is cloned and where the job is run. If the GitLab Runner `builds_dir` parameter is set, this variable is set relative to the value of `builds_dir`. For more information, see [Advanced configuration](https://docs.gitlab.com/runner/configuration/advanced-configuration.html#the-runners-section) for GitLab Runner.",
			},
			"ci_project_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The unique id of the current project that GitLab CI uses internally",
			},
			"ci_project_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the directory for the project that is currently being built. For example, if the project URL is `gitlab.example.com/group-name/project-1`, the `CI_PROJECT_NAME` would be `project-1`.",
			},
			"ci_project_namespace": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The project namespace (username or groupname) that is currently being built",
			},
			"ci_project_path": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The namespace with project name",
			},
			"ci_project_path_slug": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "`$CI_PROJECT_PATH` lowercased and with everything except `0-9` and `a-z` replaced with `-`. Use in URLs and domain names.",
			},
			"ci_project_repository_languages": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Comma-separated, lowercased list of the languages used in the repository (e.g. `ruby,javascript,html,css`)",
			},
			"ci_project_title": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The human-readable project name as displayed in the GitLab web interface.",
			},
			"ci_project_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The HTTP(S) address to access project",
			},
			"ci_project_visibility": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The project visibility (internal, private, public)",
			},
			"ci_registry": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "If the Container Registry is enabled it returns the address of GitLab's Container Registry.  This variable will include a `:port` value if one has been specified in the registry configuration.",
			},
			"ci_registry_image": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "If the Container Registry is enabled for the project it returns the address of the registry tied to the specific project",
			},
			"ci_registry_password": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The password to use to push containers to the GitLab Container Registry",
			},
			"ci_registry_user": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The username to use to push containers to the GitLab Container Registry",
			},
			"ci_repository_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL to clone the Git repository",
			},
			"ci_runner_description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The description of the runner as saved in GitLab",
			},
			"ci_runner_executable_arch": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The OS/architecture of the GitLab Runner executable (note that this is not necessarily the same as the environment of the executor)",
			},
			"ci_runner_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The unique id of runner being used",
			},
			"ci_runner_revision": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "GitLab Runner revision that is executing the current job",
			},
			"ci_runner_short_token": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "First eight characters of GitLab Runner's token used to authenticate new job requests. Used as Runner's unique ID",
			},
			"ci_runner_tags": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The defined runner tags",
			},
			"ci_runner_version": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "GitLab Runner version that is executing the current job",
			},
			"ci_server": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Mark that job is executed in CI environment",
			},
			"ci_server_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The base URL of the GitLab instance, including protocol and port (like `https://gitlab.example.com:8080`)",
			},
			"ci_server_host": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Host component of the GitLab instance URL, without protocol and port (like `gitlab.example.com`)",
			},
			"ci_server_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of CI server that is used to coordinate jobs",
			},
			"ci_server_revision": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "GitLab revision that is used to schedule jobs",
			},
			"ci_server_version": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "GitLab version that is used to schedule jobs",
			},
			"ci_server_version_major": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "GitLab version major component",
			},
			"ci_server_version_minor": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "GitLab version minor component",
			},
			"ci_server_version_patch": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "GitLab version patch component",
			},
			"ci_shared_environment": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Marks that the job is executed in a shared environment (something that is persisted across CI invocations like `shell` or `ssh` executor). If the environment is shared, it is set to true, otherwise it is not defined at all.",
			},
			"get_sources_attempts": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Number of attempts to fetch sources running a job",
			},
			"gitlab_ci": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Mark that job is executed in GitLab CI environment",
			},
			"gitlab_features": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The comma separated list of licensed features available for your instance and plan",
			},
			"gitlab_user_email": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The email of the user who started the job",
			},
			"gitlab_user_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The id of the user who started the job",
			},
			"gitlab_user_login": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The login username of the user who started the job",
			},
			"gitlab_user_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The real name of the user who started the job",
			},
			"restore_cache_attempts": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Number of attempts to restore the cache running a job",
			},
		},
	}

	// log.Printf("[TRACE] generated schema is: %s", spew.Sdump(schema))
	return schema
}

func dataSourceGitlabCIEnvironmentRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[TRACE] dataSourceGitlabCIEnvironmentRead() (mark III)")

	// This is effectively a no-op: we'd be computing these values from the
	// environment, but we've also asked that the default values be set from
	// where we'd look in the environment.

	if d.Get("ci_server").(string) == "yes" {
		d.Set("running_under_ci", true)
	} else {
		d.Set("running_under_ci", false)
	}

	// cat ~/Downloads/predefined_variables.md| perl -naF'\|' -E 's/(\s+$|^\s+)//g for @F; $F[1] =~ s/`//g; say qq{d.Set("} . lc($F[1]). qq{", os.Getenv("$F[1]"))}'

	d.Set("artifact_download_attempts", os.Getenv("ARTIFACT_DOWNLOAD_ATTEMPTS"))
	d.Set("chat_channel", os.Getenv("CHAT_CHANNEL"))
	d.Set("chat_input", os.Getenv("CHAT_INPUT"))
	d.Set("ci", os.Getenv("CI"))
	d.Set("ci_api_v4_url", os.Getenv("CI_API_V4_URL"))
	d.Set("ci_builds_dir", os.Getenv("CI_BUILDS_DIR"))
	d.Set("ci_commit_before_sha", os.Getenv("CI_COMMIT_BEFORE_SHA"))
	d.Set("ci_commit_description", os.Getenv("CI_COMMIT_DESCRIPTION"))
	d.Set("ci_commit_message", os.Getenv("CI_COMMIT_MESSAGE"))
	d.Set("ci_commit_ref_name", os.Getenv("CI_COMMIT_REF_NAME"))
	d.Set("ci_commit_ref_protected", os.Getenv("CI_COMMIT_REF_PROTECTED"))
	d.Set("ci_commit_ref_slug", os.Getenv("CI_COMMIT_REF_SLUG"))
	d.Set("ci_commit_sha", os.Getenv("CI_COMMIT_SHA"))
	d.Set("ci_commit_short_sha", os.Getenv("CI_COMMIT_SHORT_SHA"))
	d.Set("ci_commit_branch", os.Getenv("CI_COMMIT_BRANCH"))
	d.Set("ci_commit_tag", os.Getenv("CI_COMMIT_TAG"))
	d.Set("ci_commit_title", os.Getenv("CI_COMMIT_TITLE"))
	d.Set("ci_concurrent_id", os.Getenv("CI_CONCURRENT_ID"))
	d.Set("ci_concurrent_project_id", os.Getenv("CI_CONCURRENT_PROJECT_ID"))
	d.Set("ci_config_path", os.Getenv("CI_CONFIG_PATH"))
	d.Set("ci_debug_trace", os.Getenv("CI_DEBUG_TRACE"))
	d.Set("ci_default_branch", os.Getenv("CI_DEFAULT_BRANCH"))
	d.Set("ci_deploy_password", os.Getenv("CI_DEPLOY_PASSWORD"))
	d.Set("ci_deploy_user", os.Getenv("CI_DEPLOY_USER"))
	d.Set("ci_disposable_environment", os.Getenv("CI_DISPOSABLE_ENVIRONMENT"))
	d.Set("ci_environment_name", os.Getenv("CI_ENVIRONMENT_NAME"))
	d.Set("ci_environment_slug", os.Getenv("CI_ENVIRONMENT_SLUG"))
	d.Set("ci_environment_url", os.Getenv("CI_ENVIRONMENT_URL"))
	d.Set("ci_external_pull_request_iid", os.Getenv("CI_EXTERNAL_PULL_REQUEST_IID"))
	d.Set("ci_external_pull_request_source_branch_name", os.Getenv("CI_EXTERNAL_PULL_REQUEST_SOURCE_BRANCH_NAME"))
	d.Set("ci_external_pull_request_source_branch_sha", os.Getenv("CI_EXTERNAL_PULL_REQUEST_SOURCE_BRANCH_SHA"))
	d.Set("ci_external_pull_request_target_branch_name", os.Getenv("CI_EXTERNAL_PULL_REQUEST_TARGET_BRANCH_NAME"))
	d.Set("ci_external_pull_request_target_branch_sha", os.Getenv("CI_EXTERNAL_PULL_REQUEST_TARGET_BRANCH_SHA"))
	d.Set("ci_job_id", os.Getenv("CI_JOB_ID"))
	d.Set("ci_job_manual", os.Getenv("CI_JOB_MANUAL"))
	d.Set("ci_job_name", os.Getenv("CI_JOB_NAME"))
	d.Set("ci_job_stage", os.Getenv("CI_JOB_STAGE"))
	d.Set("ci_job_token", os.Getenv("CI_JOB_TOKEN"))
	d.Set("ci_job_url", os.Getenv("CI_JOB_URL"))
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
	d.Set("ci_node_index", os.Getenv("CI_NODE_INDEX"))
	d.Set("ci_node_total", os.Getenv("CI_NODE_TOTAL"))
	d.Set("ci_pages_domain", os.Getenv("CI_PAGES_DOMAIN"))
	d.Set("ci_pages_url", os.Getenv("CI_PAGES_URL"))
	d.Set("ci_pipeline_id", os.Getenv("CI_PIPELINE_ID"))
	d.Set("ci_pipeline_iid", os.Getenv("CI_PIPELINE_IID"))
	d.Set("ci_pipeline_source", os.Getenv("CI_PIPELINE_SOURCE"))
	d.Set("ci_pipeline_triggered", os.Getenv("CI_PIPELINE_TRIGGERED"))
	d.Set("ci_pipeline_url", os.Getenv("CI_PIPELINE_URL"))
	d.Set("ci_project_dir", os.Getenv("CI_PROJECT_DIR"))
	d.Set("ci_project_id", os.Getenv("CI_PROJECT_ID"))
	d.Set("ci_project_name", os.Getenv("CI_PROJECT_NAME"))
	d.Set("ci_project_namespace", os.Getenv("CI_PROJECT_NAMESPACE"))
	d.Set("ci_project_path", os.Getenv("CI_PROJECT_PATH"))
	d.Set("ci_project_path_slug", os.Getenv("CI_PROJECT_PATH_SLUG"))
	d.Set("ci_project_repository_languages", os.Getenv("CI_PROJECT_REPOSITORY_LANGUAGES"))
	d.Set("ci_project_title", os.Getenv("CI_PROJECT_TITLE"))
	d.Set("ci_project_url", os.Getenv("CI_PROJECT_URL"))
	d.Set("ci_project_visibility", os.Getenv("CI_PROJECT_VISIBILITY"))
	d.Set("ci_registry", os.Getenv("CI_REGISTRY"))
	d.Set("ci_registry_image", os.Getenv("CI_REGISTRY_IMAGE"))
	d.Set("ci_registry_password", os.Getenv("CI_REGISTRY_PASSWORD"))
	d.Set("ci_registry_user", os.Getenv("CI_REGISTRY_USER"))
	d.Set("ci_repository_url", os.Getenv("CI_REPOSITORY_URL"))
	d.Set("ci_runner_description", os.Getenv("CI_RUNNER_DESCRIPTION"))
	d.Set("ci_runner_executable_arch", os.Getenv("CI_RUNNER_EXECUTABLE_ARCH"))
	d.Set("ci_runner_id", os.Getenv("CI_RUNNER_ID"))
	d.Set("ci_runner_revision", os.Getenv("CI_RUNNER_REVISION"))
	d.Set("ci_runner_short_token", os.Getenv("CI_RUNNER_SHORT_TOKEN"))
	d.Set("ci_runner_tags", os.Getenv("CI_RUNNER_TAGS"))
	d.Set("ci_runner_version", os.Getenv("CI_RUNNER_VERSION"))
	d.Set("ci_server", os.Getenv("CI_SERVER"))
	d.Set("ci_server_url", os.Getenv("CI_SERVER_URL"))
	d.Set("ci_server_host", os.Getenv("CI_SERVER_HOST"))
	d.Set("ci_server_name", os.Getenv("CI_SERVER_NAME"))
	d.Set("ci_server_revision", os.Getenv("CI_SERVER_REVISION"))
	d.Set("ci_server_version", os.Getenv("CI_SERVER_VERSION"))
	d.Set("ci_server_version_major", os.Getenv("CI_SERVER_VERSION_MAJOR"))
	d.Set("ci_server_version_minor", os.Getenv("CI_SERVER_VERSION_MINOR"))
	d.Set("ci_server_version_patch", os.Getenv("CI_SERVER_VERSION_PATCH"))
	d.Set("ci_shared_environment", os.Getenv("CI_SHARED_ENVIRONMENT"))
	d.Set("get_sources_attempts", os.Getenv("GET_SOURCES_ATTEMPTS"))
	d.Set("gitlab_ci", os.Getenv("GITLAB_CI"))
	d.Set("gitlab_features", os.Getenv("GITLAB_FEATURES"))
	d.Set("gitlab_user_email", os.Getenv("GITLAB_USER_EMAIL"))
	d.Set("gitlab_user_id", os.Getenv("GITLAB_USER_ID"))
	d.Set("gitlab_user_login", os.Getenv("GITLAB_USER_LOGIN"))
	d.Set("gitlab_user_name", os.Getenv("GITLAB_USER_NAME"))
	d.Set("restore_cache_attempts", os.Getenv("RESTORE_CACHE_ATTEMPTS"))

	log.Printf("[TRACE] dataSourceGitlabCIEnvironmentRead() finished")
	return nil
}
