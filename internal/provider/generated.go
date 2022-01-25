package provider

// generated file. DO NOT EDIT!

import (
	"context"
	"fmt"

	"github.com/giantswarm/to"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"gitlab.com/rsrchboy/terraform-provider-gitlabci/third_party/gitlab/runner/config"
	"gitlab.com/rsrchboy/terraform-provider-gitlabci/third_party/gitlab/runner/config/docker"
	"gitlab.com/rsrchboy/terraform-provider-gitlabci/third_party/gitlab/runner/config/referees"
	"gitlab.com/rsrchboy/terraform-provider-gitlabci/third_party/gitlab/runner/config/ssh"
)

var configDataSourceRawSchema = map[string]*schema.Schema{
	"listen_address": {
		Computed:    false,
		Optional:    true,
		Description: "",
		Type:        schema.TypeString,
	},
	"session_server": {
		Computed:    false,
		Optional:    true,
		Description: "",
		Type:        schema.TypeList,
		MinItems:    0,
		MaxItems:    1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"listen_address": {
					Computed:    false,
					Optional:    true,
					Description: "Address that the runner will communicate directly with",
					Type:        schema.TypeString,
				},
				"advertise_address": {
					Computed:    false,
					Optional:    true,
					Description: "Address the runner will expose to the world to connect to the session server",
					Type:        schema.TypeString,
				},
				"session_timeout": {
					Computed:    false,
					Optional:    true,
					Description: "How long a terminal session can be active after a build completes, in seconds",
					Type:        schema.TypeInt,
				},
			},
		},
	},
	"concurrent": {
		Computed:    false,
		Optional:    true,
		Description: "",
		Type:        schema.TypeInt,
	},
	"check_interval": {
		Computed:    false,
		Optional:    true,
		Description: "Define active checking interval of jobs",
		Type:        schema.TypeInt,
	},
	"log_level": {
		Computed:    false,
		Optional:    true,
		Description: "Define log level (one of: panic, fatal, error, warning, info, debug)",
		Type:        schema.TypeString,
	},
	"log_format": {
		Computed:    false,
		Optional:    true,
		Description: "Define log format (one of: runner, text, json)",
		Type:        schema.TypeString,
	},
	"user": {
		Computed:    false,
		Optional:    true,
		Description: "",
		Type:        schema.TypeString,
	},
	"runners": {
		Computed:    false,
		Optional:    true,
		Description: "",
		Type:        schema.TypeList,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"name": {
					Computed:    false,
					Optional:    true,
					Description: "Runner name",
					Type:        schema.TypeString,
				},
				"limit": {
					Computed:    false,
					Optional:    true,
					Description: "Maximum number of builds processed by this runner",
					Type:        schema.TypeInt,
				},
				"output_limit": {
					Computed:    false,
					Optional:    true,
					Description: "Maximum build trace size in kilobytes",
					Type:        schema.TypeInt,
				},
				"request_concurrency": {
					Computed:    false,
					Optional:    true,
					Description: "Maximum concurrency for job requests",
					Type:        schema.TypeInt,
				},
				"url": {
					Computed:    false,
					Optional:    true,
					Description: "Runner URL",
					Type:        schema.TypeString,
				},
				"token": {
					Computed:    false,
					Optional:    true,
					Description: "Runner token",
					Type:        schema.TypeString,
				},
				"tls_ca_file": {
					Computed:    false,
					Optional:    true,
					Description: "File containing the certificates to verify the peer when using HTTPS",
					Type:        schema.TypeString,
				},
				"tls_cert_file": {
					Computed:    false,
					Optional:    true,
					Description: "File containing certificate for TLS client auth when using HTTPS",
					Type:        schema.TypeString,
				},
				"tls_key_file": {
					Computed:    false,
					Optional:    true,
					Description: "File containing private key for TLS client auth when using HTTPS",
					Type:        schema.TypeString,
				},
				"executor": {
					Computed:    false,
					Optional:    true,
					Description: "Select executor, eg. shell, docker, etc.",
					Type:        schema.TypeString,
				},
				"builds_dir": {
					Computed:    false,
					Optional:    true,
					Description: "Directory where builds are stored",
					Type:        schema.TypeString,
				},
				"cache_dir": {
					Computed:    false,
					Optional:    true,
					Description: "Directory where build cache is stored",
					Type:        schema.TypeString,
				},
				"clone_url": {
					Computed:    false,
					Optional:    true,
					Description: "Overwrite the default URL used to clone or fetch the git ref",
					Type:        schema.TypeString,
				},
				"environment": {
					Computed:    false,
					Optional:    true,
					Description: "Custom environment variables injected to build environment",
					Type:        schema.TypeList,
					Elem:        &schema.Schema{Type: schema.TypeString},
				},
				"pre_clone_script": {
					Computed:    false,
					Optional:    true,
					Description: "Runner-specific command script executed before code is pulled",
					Type:        schema.TypeString,
				},
				"pre_build_script": {
					Computed:    false,
					Optional:    true,
					Description: "Runner-specific command script executed after code is pulled, just before build executes",
					Type:        schema.TypeString,
				},
				"post_build_script": {
					Computed:    false,
					Optional:    true,
					Description: "Runner-specific command script executed after code is pulled and just after build executes",
					Type:        schema.TypeString,
				},
				"debug_trace_disabled": {
					Computed:    false,
					Optional:    true,
					Description: "When set to true Runner will disable the possibility of using the CI_DEBUG_TRACE feature",
					Type:        schema.TypeBool,
				},
				"shell": {
					Computed:    false,
					Optional:    true,
					Description: "Select bash, cmd, pwsh or powershell",
					Type:        schema.TypeString,
				},
				"custom_build_dir": {
					Computed:    false,
					Optional:    true,
					Description: "",
					Type:        schema.TypeList,
					MinItems:    0,
					MaxItems:    1,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"enabled": {
								Computed:    false,
								Optional:    true,
								Description: "Enable job specific build directories",
								Type:        schema.TypeBool,
							},
						},
					},
				},
				"referees": {
					Computed:    false,
					Optional:    true,
					Description: "",
					Type:        schema.TypeList,
					MinItems:    0,
					MaxItems:    1,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"metrics": {
								Computed:    false,
								Optional:    true,
								Description: "",
								Type:        schema.TypeList,
								MinItems:    0,
								MaxItems:    1,
								Elem: &schema.Resource{
									Schema: map[string]*schema.Schema{
										"prometheus_address": {
											Computed:    false,
											Optional:    true,
											Description: "A host:port to a prometheus metrics server",
											Type:        schema.TypeString,
										},
										"query_interval": {
											Computed:    false,
											Optional:    true,
											Description: "Query interval (in seconds)",
											Type:        schema.TypeInt,
										},
										"queries": {
											Computed:    false,
											Optional:    true,
											Description: "A list of metrics to query (in PromQL)",
											Type:        schema.TypeList,
											Elem:        &schema.Schema{Type: schema.TypeString},
										},
									},
								},
							},
						},
					},
				},
				"cache": {
					Computed:    false,
					Optional:    true,
					Description: "",
					Type:        schema.TypeList,
					MinItems:    0,
					MaxItems:    1,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"type": {
								Computed:    false,
								Optional:    true,
								Description: "Select caching method",
								Type:        schema.TypeString,
							},
							"path": {
								Computed:    false,
								Optional:    true,
								Description: "Name of the path to prepend to the cache URL",
								Type:        schema.TypeString,
							},
							"shared": {
								Computed:    false,
								Optional:    true,
								Description: "Enable cache sharing between runners.",
								Type:        schema.TypeBool,
							},
							"s3": {
								Computed:    false,
								Optional:    true,
								Description: "",
								Type:        schema.TypeList,
								MinItems:    0,
								MaxItems:    1,
								Elem: &schema.Resource{
									Schema: map[string]*schema.Schema{
										"server_address": {
											Computed:    false,
											Optional:    true,
											Description: "A host:port to the used S3-compatible server",
											Type:        schema.TypeString,
										},
										"access_key": {
											Computed:    false,
											Optional:    true,
											Description: "S3 Access Key",
											Type:        schema.TypeString,
										},
										"secret_key": {
											Computed:    false,
											Optional:    true,
											Description: "S3 Secret Key",
											Type:        schema.TypeString,
										},
										"bucket_name": {
											Computed:    false,
											Optional:    true,
											Description: "Name of the bucket where cache will be stored",
											Type:        schema.TypeString,
										},
										"bucket_location": {
											Computed:    false,
											Optional:    true,
											Description: "Name of S3 region",
											Type:        schema.TypeString,
										},
										"insecure": {
											Computed:    false,
											Optional:    true,
											Description: "Use insecure mode (without https)",
											Type:        schema.TypeBool,
										},
										"authentication_type": {
											Computed:    false,
											Optional:    true,
											Description: "IAM or credentials",
											Type:        schema.TypeString,
										},
									},
								},
							},
							"gcs": {
								Computed:    false,
								Optional:    true,
								Description: "",
								Type:        schema.TypeList,
								MinItems:    0,
								MaxItems:    1,
								Elem: &schema.Resource{
									Schema: map[string]*schema.Schema{
										"access_id": {
											Computed:    false,
											Optional:    true,
											Description: "ID of GCP Service Account used to access the storage",
											Type:        schema.TypeString,
										},
										"private_key": {
											Computed:    false,
											Optional:    true,
											Description: "Private key used to sign GCS requests",
											Type:        schema.TypeString,
										},
										"credentials_file": {
											Computed:    false,
											Optional:    true,
											Description: "File with GCP credentials, containing AccessID and PrivateKey",
											Type:        schema.TypeString,
										},
										"bucket_name": {
											Computed:    false,
											Optional:    true,
											Description: "Name of the bucket where cache will be stored",
											Type:        schema.TypeString,
										},
									},
								},
							},
							"azure": {
								Computed:    false,
								Optional:    true,
								Description: "",
								Type:        schema.TypeList,
								MinItems:    0,
								MaxItems:    1,
								Elem: &schema.Resource{
									Schema: map[string]*schema.Schema{
										"account_name": {
											Computed:    false,
											Optional:    true,
											Description: "Account name for Azure Blob Storage",
											Type:        schema.TypeString,
										},
										"account_key": {
											Computed:    false,
											Optional:    true,
											Description: "Access key for Azure Blob Storage",
											Type:        schema.TypeString,
										},
										"container_name": {
											Computed:    false,
											Optional:    true,
											Description: "Name of the Azure container where cache will be stored",
											Type:        schema.TypeString,
										},
										"storage_domain": {
											Computed:    false,
											Optional:    true,
											Description: "Domain name of the Azure storage (e.g. blob.core.windows.net)",
											Type:        schema.TypeString,
										},
									},
								},
							},
						},
					},
				},
				"feature_flags": {
					Computed:    false,
					Optional:    true,
					Description: "Enable/Disable feature flags https://docs.gitlab.com/runner/configuration/feature-flags.html",
					Type:        schema.TypeMap,
					Elem:        &schema.Schema{Type: schema.TypeBool},
				},
				"ssh": {
					Computed:    false,
					Optional:    true,
					Description: "",
					Type:        schema.TypeList,
					MinItems:    0,
					MaxItems:    1,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"user": {
								Computed:    false,
								Optional:    true,
								Description: "User name",
								Type:        schema.TypeString,
							},
							"password": {
								Computed:    false,
								Optional:    true,
								Description: "User password",
								Type:        schema.TypeString,
							},
							"host": {
								Computed:    false,
								Optional:    true,
								Description: "Remote host",
								Type:        schema.TypeString,
							},
							"port": {
								Computed:    false,
								Optional:    true,
								Description: "Remote host port",
								Type:        schema.TypeString,
							},
							"identity_file": {
								Computed:    false,
								Optional:    true,
								Description: "Identity file to be used",
								Type:        schema.TypeString,
							},
							"disable_strict_host_key_checking": {
								Computed:    false,
								Optional:    true,
								Description: "Disable SSH strict host key checking",
								Type:        schema.TypeBool,
							},
							"known_hosts_file": {
								Computed:    false,
								Optional:    true,
								Description: "Location of known_hosts file. Defaults to ~/.ssh/known_hosts",
								Type:        schema.TypeString,
							},
						},
					},
				},
				"docker": {
					Computed:    false,
					Optional:    true,
					Description: "",
					Type:        schema.TypeList,
					MinItems:    0,
					MaxItems:    1,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"host": {
								Computed:    false,
								Optional:    true,
								Description: "Docker daemon address",
								Type:        schema.TypeString,
							},
							"tls_cert_path": {
								Computed:    false,
								Optional:    true,
								Description: "Certificate path",
								Type:        schema.TypeString,
							},
							"tls_verify": {
								Computed:    false,
								Optional:    true,
								Description: "Use TLS and verify the remote",
								Type:        schema.TypeBool,
							},
							"hostname": {
								Computed:    false,
								Optional:    true,
								Description: "Custom container hostname",
								Type:        schema.TypeString,
							},
							"image": {
								Computed:    false,
								Optional:    true,
								Description: "Docker image to be used",
								Type:        schema.TypeString,
							},
							"runtime": {
								Computed:    false,
								Optional:    true,
								Description: "Docker runtime to be used",
								Type:        schema.TypeString,
							},
							"memory": {
								Computed:    false,
								Optional:    true,
								Description: "Memory limit (format: <number>[<unit>]). Unit can be one of b, k, m, or g. Minimum is 4M.",
								Type:        schema.TypeString,
							},
							"memory_swap": {
								Computed:    false,
								Optional:    true,
								Description: "Total memory limit (memory + swap, format: <number>[<unit>]). Unit can be one of b, k, m, or g.",
								Type:        schema.TypeString,
							},
							"memory_reservation": {
								Computed:    false,
								Optional:    true,
								Description: "Memory soft limit (format: <number>[<unit>]). Unit can be one of b, k, m, or g.",
								Type:        schema.TypeString,
							},
							"cpuset_cpus": {
								Computed:    false,
								Optional:    true,
								Description: "String value containing the cgroups CpusetCpus to use",
								Type:        schema.TypeString,
							},
							"cpus": {
								Computed:    false,
								Optional:    true,
								Description: "Number of CPUs",
								Type:        schema.TypeString,
							},
							"cpu_shares": {
								Computed:    false,
								Optional:    true,
								Description: "Number of CPU shares",
								Type:        schema.TypeInt,
							},
							"dns": {
								Computed:    false,
								Optional:    true,
								Description: "A list of DNS servers for the container to use",
								Type:        schema.TypeList,
								Elem:        &schema.Schema{Type: schema.TypeString},
							},
							"dns_search": {
								Computed:    false,
								Optional:    true,
								Description: "A list of DNS search domains",
								Type:        schema.TypeList,
								Elem:        &schema.Schema{Type: schema.TypeString},
							},
							"privileged": {
								Computed:    false,
								Optional:    true,
								Description: "Give extended privileges to container",
								Type:        schema.TypeBool,
							},
							"disable_entrypoint_overwrite": {
								Computed:    false,
								Optional:    true,
								Description: "Disable the possibility for a container to overwrite the default image entrypoint",
								Type:        schema.TypeBool,
							},
							"userns_mode": {
								Computed:    false,
								Optional:    true,
								Description: "User namespace to use",
								Type:        schema.TypeString,
							},
							"cap_add": {
								Computed:    false,
								Optional:    true,
								Description: "Add Linux capabilities",
								Type:        schema.TypeList,
								Elem:        &schema.Schema{Type: schema.TypeString},
							},
							"cap_drop": {
								Computed:    false,
								Optional:    true,
								Description: "Drop Linux capabilities",
								Type:        schema.TypeList,
								Elem:        &schema.Schema{Type: schema.TypeString},
							},
							"oom_kill_disable": {
								Computed:    false,
								Optional:    true,
								Description: "Do not kill processes in a container if an out-of-memory (OOM) error occurs",
								Type:        schema.TypeBool,
							},
							"oom_score_adjust": {
								Computed:    false,
								Optional:    true,
								Description: "Adjust OOM score",
								Type:        schema.TypeInt,
							},
							"security_opt": {
								Computed:    false,
								Optional:    true,
								Description: "Security Options",
								Type:        schema.TypeList,
								Elem:        &schema.Schema{Type: schema.TypeString},
							},
							"devices": {
								Computed:    false,
								Optional:    true,
								Description: "Add a host device to the container",
								Type:        schema.TypeList,
								Elem:        &schema.Schema{Type: schema.TypeString},
							},
							"gpus": {
								Computed:    false,
								Optional:    true,
								Description: "Request GPUs to be used by Docker",
								Type:        schema.TypeString,
							},
							"disable_cache": {
								Computed:    false,
								Optional:    true,
								Description: "Disable all container caching",
								Type:        schema.TypeBool,
							},
							"volumes": {
								Computed:    false,
								Optional:    true,
								Description: "Bind-mount a volume and create it if it doesn't exist prior to mounting. Can be specified multiple times once per mountpoint, e.g. --docker-volumes 'test0:/test0' --docker-volumes 'test1:/test1'",
								Type:        schema.TypeList,
								Elem:        &schema.Schema{Type: schema.TypeString},
							},
							"volume_driver": {
								Computed:    false,
								Optional:    true,
								Description: "Volume driver to be used",
								Type:        schema.TypeString,
							},
							"cache_dir": {
								Computed:    false,
								Optional:    true,
								Description: "Directory where to store caches",
								Type:        schema.TypeString,
							},
							"extra_hosts": {
								Computed:    false,
								Optional:    true,
								Description: "Add a custom host-to-IP mapping",
								Type:        schema.TypeList,
								Elem:        &schema.Schema{Type: schema.TypeString},
							},
							"volumes_from": {
								Computed:    false,
								Optional:    true,
								Description: "A list of volumes to inherit from another container",
								Type:        schema.TypeList,
								Elem:        &schema.Schema{Type: schema.TypeString},
							},
							"network_mode": {
								Computed:    false,
								Optional:    true,
								Description: "Add container to a custom network",
								Type:        schema.TypeString,
							},
							"links": {
								Computed:    false,
								Optional:    true,
								Description: "Add link to another container",
								Type:        schema.TypeList,
								Elem:        &schema.Schema{Type: schema.TypeString},
							},
							"services": {
								Computed:    false,
								Optional:    true,
								Description: "Add service that is started with container",
								Type:        schema.TypeList,
								Elem: &schema.Resource{
									Schema: map[string]*schema.Schema{
										"name": {
											Computed:    false,
											Optional:    true,
											Description: "The image path for the service",
											Type:        schema.TypeString,
										},
										"alias": {
											Computed:    false,
											Optional:    true,
											Description: "The alias of the service",
											Type:        schema.TypeString,
										},
										"command": {
											Computed:    false,
											Optional:    true,
											Description: "Command or script that should be used as the container’s command. Syntax is similar to https://docs.docker.com/engine/reference/builder/#cmd",
											Type:        schema.TypeList,
											Elem:        &schema.Schema{Type: schema.TypeString},
										},
										"entrypoint": {
											Computed:    false,
											Optional:    true,
											Description: "Command or script that should be executed as the container’s entrypoint. syntax is similar to https://docs.docker.com/engine/reference/builder/#entrypoint",
											Type:        schema.TypeList,
											Elem:        &schema.Schema{Type: schema.TypeString},
										},
									},
								},
							},
							"wait_for_services_timeout": {
								Computed:    false,
								Optional:    true,
								Description: "How long to wait for service startup",
								Type:        schema.TypeInt,
							},
							"allowed_images": {
								Computed:    false,
								Optional:    true,
								Description: "Image allowlist",
								Type:        schema.TypeList,
								Elem:        &schema.Schema{Type: schema.TypeString},
							},
							"allowed_services": {
								Computed:    false,
								Optional:    true,
								Description: "Service allowlist",
								Type:        schema.TypeList,
								Elem:        &schema.Schema{Type: schema.TypeString},
							},
							"pull_policy": {
								Computed:    false,
								Optional:    true,
								Description: "Image pull policy: never, if-not-present, always",
								Type:        schema.TypeList,
								Elem:        &schema.Schema{Type: schema.TypeString},
							},
							"shm_size": {
								Computed:    false,
								Optional:    true,
								Description: "Shared memory size for docker images (in bytes)",
								Type:        schema.TypeInt,
							},
							"tmpfs": {
								Computed:    false,
								Optional:    true,
								Description: "A toml table/json object with the format key=values. When set this will mount the specified path in the key as a tmpfs volume in the main container, using the options specified as key. For the supported options, see the documentation for the unix 'mount' command",
								Type:        schema.TypeMap,
								Elem:        &schema.Schema{Type: schema.TypeString},
							},
							"services_tmpfs": {
								Computed:    false,
								Optional:    true,
								Description: "A toml table/json object with the format key=values. When set this will mount the specified path in the key as a tmpfs volume in all the service containers, using the options specified as key. For the supported options, see the documentation for the unix 'mount' command",
								Type:        schema.TypeMap,
								Elem:        &schema.Schema{Type: schema.TypeString},
							},
							"sysctls": {
								Computed:    false,
								Optional:    true,
								Description: "Sysctl options, a toml table/json object of key=value. Value is expected to be a string.",
								Type:        schema.TypeMap,
								Elem:        &schema.Schema{Type: schema.TypeString},
							},
							"helper_image": {
								Computed:    false,
								Optional:    true,
								Description: "[ADVANCED] Override the default helper image used to clone repos and upload artifacts",
								Type:        schema.TypeString,
							},
							"helper_image_flavor": {
								Computed:    false,
								Optional:    true,
								Description: "Set helper image flavor (alpine, ubuntu), defaults to alpine",
								Type:        schema.TypeString,
							},
						},
					},
				},
				"parallels": {
					Computed:    false,
					Optional:    true,
					Description: "",
					Type:        schema.TypeList,
					MinItems:    0,
					MaxItems:    1,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"base_name": {
								Computed:    false,
								Optional:    true,
								Description: "VM name to be used",
								Type:        schema.TypeString,
							},
							"template_name": {
								Computed:    false,
								Optional:    true,
								Description: "VM template to be created",
								Type:        schema.TypeString,
							},
							"disable_snapshots": {
								Computed:    false,
								Optional:    true,
								Description: "Disable snapshoting to speedup VM creation",
								Type:        schema.TypeBool,
							},
							"time_server": {
								Computed:    false,
								Optional:    true,
								Description: "Timeserver to sync the guests time from. Defaults to time.apple.com",
								Type:        schema.TypeString,
							},
							"allowed_images": {
								Computed:    false,
								Optional:    true,
								Description: "Image (base_name) allowlist",
								Type:        schema.TypeList,
								Elem:        &schema.Schema{Type: schema.TypeString},
							},
						},
					},
				},
				"virtualbox": {
					Computed:    false,
					Optional:    true,
					Description: "",
					Type:        schema.TypeList,
					MinItems:    0,
					MaxItems:    1,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"base_name": {
								Computed:    false,
								Optional:    true,
								Description: "VM name to be used",
								Type:        schema.TypeString,
							},
							"base_snapshot": {
								Computed:    false,
								Optional:    true,
								Description: "Name or UUID of a specific VM snapshot to clone",
								Type:        schema.TypeString,
							},
							"base_folder": {
								Computed:    false,
								Optional:    true,
								Description: "Folder in which to save the new VM. If empty, uses VirtualBox default",
								Type:        schema.TypeString,
							},
							"disable_snapshots": {
								Computed:    false,
								Optional:    true,
								Description: "Disable snapshoting to speedup VM creation",
								Type:        schema.TypeBool,
							},
							"allowed_images": {
								Computed:    false,
								Optional:    true,
								Description: "Image allowlist",
								Type:        schema.TypeList,
								Elem:        &schema.Schema{Type: schema.TypeString},
							},
						},
					},
				},
				"machine": {
					Computed:    false,
					Optional:    true,
					Description: "",
					Type:        schema.TypeList,
					MinItems:    0,
					MaxItems:    1,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"max_growth_rate": {
								Computed:    false,
								Optional:    true,
								Description: "Maximum machines being provisioned concurrently, set to 0 for unlimited",
								Type:        schema.TypeInt,
							},
							"idle_count": {
								Computed:    false,
								Optional:    true,
								Description: "Maximum idle machines",
								Type:        schema.TypeInt,
							},
							"idle_scale_factor": {
								Computed:    false,
								Optional:    true,
								Description: "(Experimental) Defines what factor of in-use machines should be used as current idle value, but never more then defined IdleCount. 0.0 means use IdleCount as a static number (defaults to 0.0). Must be defined as float number.",
								Type:        schema.TypeFloat,
							},
							"idle_count_min": {
								Computed:    false,
								Optional:    true,
								Description: "Minimal number of idle machines when IdleScaleFactor is in use. Defaults to 1.",
								Type:        schema.TypeInt,
							},
							"idle_time": {
								Computed:    false,
								Optional:    true,
								Description: "Minimum time after node can be destroyed",
								Type:        schema.TypeInt,
							},
							"max_builds": {
								Computed:    false,
								Optional:    true,
								Description: "Maximum number of builds processed by machine",
								Type:        schema.TypeInt,
							},
							"machine_driver": {
								Computed:    false,
								Optional:    true,
								Description: "The driver to use when creating machine",
								Type:        schema.TypeString,
							},
							"machine_name": {
								Computed:    false,
								Optional:    true,
								Description: "The template for machine name (needs to include %s)",
								Type:        schema.TypeString,
							},
							"machine_options": {
								Computed:    false,
								Optional:    true,
								Description: "Additional machine creation options",
								Type:        schema.TypeList,
								Elem:        &schema.Schema{Type: schema.TypeString},
							},
							"off_peak_periods": {
								Computed:    false,
								Optional:    true,
								Description: "Time periods when the scheduler is in the OffPeak mode. DEPRECATED",
								Type:        schema.TypeList,
								Elem:        &schema.Schema{Type: schema.TypeString},
							},
							"off_peak_timezone": {
								Computed:    false,
								Optional:    true,
								Description: "Timezone for the OffPeak periods (defaults to Local). DEPRECATED",
								Type:        schema.TypeString,
							},
							"off_peak_idle_count": {
								Computed:    false,
								Optional:    true,
								Description: "Maximum idle machines when the scheduler is in the OffPeak mode. DEPRECATED",
								Type:        schema.TypeInt,
							},
							"off_peak_idle_time": {
								Computed:    false,
								Optional:    true,
								Description: "Minimum time after machine can be destroyed when the scheduler is in the OffPeak mode. DEPRECATED",
								Type:        schema.TypeInt,
							},
							"autoscaling": {
								Computed:    false,
								Optional:    true,
								Description: "Ordered list of configurations for autoscaling periods (last match wins)",
								Type:        schema.TypeList,
								Elem: &schema.Resource{
									Schema: map[string]*schema.Schema{
										"periods": {
											Computed:    false,
											Optional:    true,
											Description: "List of crontab expressions for this autoscaling configuration",
											Type:        schema.TypeList,
											Elem:        &schema.Schema{Type: schema.TypeString},
										},
										"timezone": {
											Computed:    false,
											Optional:    true,
											Description: "Timezone for the periods (defaults to Local)",
											Type:        schema.TypeString,
										},
										"idle_count": {
											Computed:    false,
											Optional:    true,
											Description: "Maximum idle machines when this configuration is active",
											Type:        schema.TypeInt,
										},
										"idle_scale_factor": {
											Computed:    false,
											Optional:    true,
											Description: "(Experimental) Defines what factor of in-use machines should be used as current idle value, but never more then defined IdleCount. 0.0 means use IdleCount as a static number (defaults to 0.0). Must be defined as float number.",
											Type:        schema.TypeFloat,
										},
										"idle_count_min": {
											Computed:    false,
											Optional:    true,
											Description: "Minimal number of idle machines when IdleScaleFactor is in use. Defaults to 1.",
											Type:        schema.TypeInt,
										},
										"idle_time": {
											Computed:    false,
											Optional:    true,
											Description: "Minimum time after which and idle machine can be destroyed when this configuration is active",
											Type:        schema.TypeInt,
										},
									},
								},
							},
						},
					},
				},
				"kubernetes": {
					Computed:    false,
					Optional:    true,
					Description: "",
					Type:        schema.TypeList,
					MinItems:    0,
					MaxItems:    1,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"host": {
								Computed:    false,
								Optional:    true,
								Description: "Optional Kubernetes master host URL (auto-discovery attempted if not specified)",
								Type:        schema.TypeString,
							},
							"cert_file": {
								Computed:    false,
								Optional:    true,
								Description: "Optional Kubernetes master auth certificate",
								Type:        schema.TypeString,
							},
							"key_file": {
								Computed:    false,
								Optional:    true,
								Description: "Optional Kubernetes master auth private key",
								Type:        schema.TypeString,
							},
							"ca_file": {
								Computed:    false,
								Optional:    true,
								Description: "Optional Kubernetes master auth ca certificate",
								Type:        schema.TypeString,
							},
							"bearer_token_overwrite_allowed": {
								Computed:    false,
								Optional:    true,
								Description: "Bool to authorize builds to specify their own bearer token for creation.",
								Type:        schema.TypeBool,
							},
							"bearer_token": {
								Computed:    false,
								Optional:    true,
								Description: "Optional Kubernetes service account token used to start build pods.",
								Type:        schema.TypeString,
							},
							"image": {
								Computed:    false,
								Optional:    true,
								Description: "Default docker image to use for builds when none is specified",
								Type:        schema.TypeString,
							},
							"namespace": {
								Computed:    false,
								Optional:    true,
								Description: "Namespace to run Kubernetes jobs in",
								Type:        schema.TypeString,
							},
							"namespace_overwrite_allowed": {
								Computed:    false,
								Optional:    true,
								Description: "Regex to validate 'KUBERNETES_NAMESPACE_OVERWRITE' value",
								Type:        schema.TypeString,
							},
							"privileged": {
								Computed:    false,
								Optional:    true,
								Description: "Run all containers with the privileged flag enabled",
								Type:        schema.TypeBool,
							},
							"allow_privilege_escalation": {
								Computed:    false,
								Optional:    true,
								Description: "Run all containers with the security context allowPrivilegeEscalation flag enabled. When empty, it does not define the allowPrivilegeEscalation flag in the container SecurityContext and allows Kubernetes to use the default privilege escalation behavior.",
								Type:        schema.TypeBool,
							},
							"cpu_limit": {
								Computed:    false,
								Optional:    true,
								Description: "The CPU allocation given to build containers",
								Type:        schema.TypeString,
							},
							"cpu_limit_overwrite_max_allowed": {
								Computed:    false,
								Optional:    true,
								Description: "If set, the max amount the cpu limit can be set to. Used with the KUBERNETES_CPU_LIMIT variable in the build.",
								Type:        schema.TypeString,
							},
							"cpu_request": {
								Computed:    false,
								Optional:    true,
								Description: "The CPU allocation requested for build containers",
								Type:        schema.TypeString,
							},
							"cpu_request_overwrite_max_allowed": {
								Computed:    false,
								Optional:    true,
								Description: "If set, the max amount the cpu request can be set to. Used with the KUBERNETES_CPU_REQUEST variable in the build.",
								Type:        schema.TypeString,
							},
							"memory_limit": {
								Computed:    false,
								Optional:    true,
								Description: "The amount of memory allocated to build containers",
								Type:        schema.TypeString,
							},
							"memory_limit_overwrite_max_allowed": {
								Computed:    false,
								Optional:    true,
								Description: "If set, the max amount the memory limit can be set to. Used with the KUBERNETES_MEMORY_LIMIT variable in the build.",
								Type:        schema.TypeString,
							},
							"memory_request": {
								Computed:    false,
								Optional:    true,
								Description: "The amount of memory requested from build containers",
								Type:        schema.TypeString,
							},
							"memory_request_overwrite_max_allowed": {
								Computed:    false,
								Optional:    true,
								Description: "If set, the max amount the memory request can be set to. Used with the KUBERNETES_MEMORY_REQUEST variable in the build.",
								Type:        schema.TypeString,
							},
							"ephemeral_storage_limit": {
								Computed:    false,
								Optional:    true,
								Description: "The amount of ephemeral storage allocated to build containers",
								Type:        schema.TypeString,
							},
							"ephemeral_storage_limit_overwrite_max_allowed": {
								Computed:    false,
								Optional:    true,
								Description: "If set, the max amount the ephemeral limit can be set to. Used with the KUBERNETES_EPHEMERAL_STORAGE_LIMIT variable in the build.",
								Type:        schema.TypeString,
							},
							"ephemeral_storage_request": {
								Computed:    false,
								Optional:    true,
								Description: "The amount of ephemeral storage requested from build containers",
								Type:        schema.TypeString,
							},
							"ephemeral_storage_request_overwrite_max_allowed": {
								Computed:    false,
								Optional:    true,
								Description: "If set, the max amount the ephemeral storage request can be set to. Used with the KUBERNETES_EPHEMERAL_STORAGE_REQUEST variable in the build.",
								Type:        schema.TypeString,
							},
							"service_cpu_limit": {
								Computed:    false,
								Optional:    true,
								Description: "The CPU allocation given to build service containers",
								Type:        schema.TypeString,
							},
							"service_cpu_limit_overwrite_max_allowed": {
								Computed:    false,
								Optional:    true,
								Description: "If set, the max amount the service cpu limit can be set to. Used with the KUBERNETES_SERVICE_CPU_LIMIT variable in the build.",
								Type:        schema.TypeString,
							},
							"service_cpu_request": {
								Computed:    false,
								Optional:    true,
								Description: "The CPU allocation requested for build service containers",
								Type:        schema.TypeString,
							},
							"service_cpu_request_overwrite_max_allowed": {
								Computed:    false,
								Optional:    true,
								Description: "If set, the max amount the service cpu request can be set to. Used with the KUBERNETES_SERVICE_CPU_REQUEST variable in the build.",
								Type:        schema.TypeString,
							},
							"service_memory_limit": {
								Computed:    false,
								Optional:    true,
								Description: "The amount of memory allocated to build service containers",
								Type:        schema.TypeString,
							},
							"service_memory_limit_overwrite_max_allowed": {
								Computed:    false,
								Optional:    true,
								Description: "If set, the max amount the service memory limit can be set to. Used with the KUBERNETES_SERVICE_MEMORY_LIMIT variable in the build.",
								Type:        schema.TypeString,
							},
							"service_memory_request": {
								Computed:    false,
								Optional:    true,
								Description: "The amount of memory requested for build service containers",
								Type:        schema.TypeString,
							},
							"service_memory_request_overwrite_max_allowed": {
								Computed:    false,
								Optional:    true,
								Description: "If set, the max amount the service memory request can be set to. Used with the KUBERNETES_SERVICE_MEMORY_REQUEST variable in the build.",
								Type:        schema.TypeString,
							},
							"service_ephemeral_storage_limit": {
								Computed:    false,
								Optional:    true,
								Description: "The amount of ephemeral storage allocated to build service containers",
								Type:        schema.TypeString,
							},
							"service_ephemeral_storage_limit_overwrite_max_allowed": {
								Computed:    false,
								Optional:    true,
								Description: "If set, the max amount the service ephemeral storage limit can be set to. Used with the KUBERNETES_SERVICE_EPHEMERAL_STORAGE_LIMIT variable in the build.",
								Type:        schema.TypeString,
							},
							"service_ephemeral_storage_request": {
								Computed:    false,
								Optional:    true,
								Description: "The amount of ephemeral storage requested for build service containers",
								Type:        schema.TypeString,
							},
							"service_ephemeral_storage_request_overwrite_max_allowed": {
								Computed:    false,
								Optional:    true,
								Description: "If set, the max amount the service ephemeral storage request can be set to. Used with the KUBERNETES_SERVICE_EPHEMERAL_STORAGE_REQUEST variable in the build.",
								Type:        schema.TypeString,
							},
							"helper_cpu_limit": {
								Computed:    false,
								Optional:    true,
								Description: "The CPU allocation given to build helper containers",
								Type:        schema.TypeString,
							},
							"helper_cpu_limit_overwrite_max_allowed": {
								Computed:    false,
								Optional:    true,
								Description: "If set, the max amount the helper cpu limit can be set to. Used with the KUBERNETES_HELPER_CPU_LIMIT variable in the build.",
								Type:        schema.TypeString,
							},
							"helper_cpu_request": {
								Computed:    false,
								Optional:    true,
								Description: "The CPU allocation requested for build helper containers",
								Type:        schema.TypeString,
							},
							"helper_cpu_request_overwrite_max_allowed": {
								Computed:    false,
								Optional:    true,
								Description: "If set, the max amount the helper cpu request can be set to. Used with the KUBERNETES_HELPER_CPU_REQUEST variable in the build.",
								Type:        schema.TypeString,
							},
							"helper_memory_limit": {
								Computed:    false,
								Optional:    true,
								Description: "The amount of memory allocated to build helper containers",
								Type:        schema.TypeString,
							},
							"helper_memory_limit_overwrite_max_allowed": {
								Computed:    false,
								Optional:    true,
								Description: "If set, the max amount the helper memory limit can be set to. Used with the KUBERNETES_HELPER_MEMORY_LIMIT variable in the build.",
								Type:        schema.TypeString,
							},
							"helper_memory_request": {
								Computed:    false,
								Optional:    true,
								Description: "The amount of memory requested for build helper containers",
								Type:        schema.TypeString,
							},
							"helper_memory_request_overwrite_max_allowed": {
								Computed:    false,
								Optional:    true,
								Description: "If set, the max amount the helper memory request can be set to. Used with the KUBERNETES_HELPER_MEMORY_REQUEST variable in the build.",
								Type:        schema.TypeString,
							},
							"helper_ephemeral_storage_limit": {
								Computed:    false,
								Optional:    true,
								Description: "The amount of ephemeral storage allocated to build helper containers",
								Type:        schema.TypeString,
							},
							"helper_ephemeral_storage_limit_overwrite_max_allowed": {
								Computed:    false,
								Optional:    true,
								Description: "If set, the max amount the helper ephemeral storage limit can be set to. Used with the KUBERNETES_HELPER_EPHEMERAL_STORAGE_LIMIT variable in the build.",
								Type:        schema.TypeString,
							},
							"helper_ephemeral_storage_request": {
								Computed:    false,
								Optional:    true,
								Description: "The amount of ephemeral storage requested for build helper containers",
								Type:        schema.TypeString,
							},
							"helper_ephemeral_storage_request_overwrite_max_allowed": {
								Computed:    false,
								Optional:    true,
								Description: "If set, the max amount the helper ephemeral storage request can be set to. Used with the KUBERNETES_HELPER_EPHEMERAL_STORAGE_REQUEST variable in the build.",
								Type:        schema.TypeString,
							},
							"allowed_images": {
								Computed:    false,
								Optional:    true,
								Description: "Image allowlist",
								Type:        schema.TypeList,
								Elem:        &schema.Schema{Type: schema.TypeString},
							},
							"allowed_services": {
								Computed:    false,
								Optional:    true,
								Description: "Service allowlist",
								Type:        schema.TypeList,
								Elem:        &schema.Schema{Type: schema.TypeString},
							},
							"pull_policy": {
								Computed:    false,
								Optional:    true,
								Description: "Policy for if/when to pull a container image (never, if-not-present, always). The cluster default will be used if not set",
								Type:        schema.TypeList,
								Elem:        &schema.Schema{Type: schema.TypeString},
							},
							"node_selector": {
								Computed:    false,
								Optional:    true,
								Description: "A toml table/json object of key:value. Value is expected to be a string. When set this will create pods on k8s nodes that match all the key:value pairs. Only one selector is supported through environment variable configuration.",
								Type:        schema.TypeMap,
								Elem:        &schema.Schema{Type: schema.TypeString},
							},
							"node_tolerations": {
								Computed:    false,
								Optional:    true,
								Description: "A toml table/json object of key=value:effect. Value and effect are expected to be strings. When set, pods will tolerate the given taints. Only one toleration is supported through environment variable configuration.",
								Type:        schema.TypeMap,
								Elem:        &schema.Schema{Type: schema.TypeString},
							},
							"affinity": {
								Computed:    false,
								Optional:    true,
								Description: "Kubernetes Affinity setting that is used to select the node that spawns a pod",
								Type:        schema.TypeList,
								MinItems:    0,
								MaxItems:    1,
								Elem: &schema.Resource{
									Schema: map[string]*schema.Schema{
										"node_affinity": {
											Computed:    false,
											Optional:    true,
											Description: "Node affinity is conceptually similar to nodeSelector -- it allows you to constrain which nodes your pod is eligible to be scheduled on, based on labels on the node.",
											Type:        schema.TypeList,
											MinItems:    0,
											MaxItems:    1,
											Elem: &schema.Resource{
												Schema: map[string]*schema.Schema{
													"required_during_scheduling_ignored_during_execution": {
														Computed:    false,
														Optional:    true,
														Description: "",
														Type:        schema.TypeList,
														MinItems:    0,
														MaxItems:    1,
														Elem: &schema.Resource{
															Schema: map[string]*schema.Schema{
																"node_selector_terms": {
																	Computed:    false,
																	Optional:    true,
																	Description: "",
																	Type:        schema.TypeList,
																	Elem: &schema.Resource{
																		Schema: map[string]*schema.Schema{
																			"match_expressions": {
																				Computed:    false,
																				Optional:    true,
																				Description: "",
																				Type:        schema.TypeList,
																				Elem: &schema.Resource{
																					Schema: map[string]*schema.Schema{
																						"key": {
																							Computed:    false,
																							Optional:    true,
																							Description: "",
																							Type:        schema.TypeString,
																						},
																						"operator": {
																							Computed:    false,
																							Optional:    true,
																							Description: "",
																							Type:        schema.TypeString,
																						},
																						"values": {
																							Computed:    false,
																							Optional:    true,
																							Description: "",
																							Type:        schema.TypeList,
																							Elem:        &schema.Schema{Type: schema.TypeString},
																						},
																					},
																				},
																			},
																			"match_fields": {
																				Computed:    false,
																				Optional:    true,
																				Description: "",
																				Type:        schema.TypeList,
																				Elem: &schema.Resource{
																					Schema: map[string]*schema.Schema{
																						"key": {
																							Computed:    false,
																							Optional:    true,
																							Description: "",
																							Type:        schema.TypeString,
																						},
																						"operator": {
																							Computed:    false,
																							Optional:    true,
																							Description: "",
																							Type:        schema.TypeString,
																						},
																						"values": {
																							Computed:    false,
																							Optional:    true,
																							Description: "",
																							Type:        schema.TypeList,
																							Elem:        &schema.Schema{Type: schema.TypeString},
																						},
																					},
																				},
																			},
																		},
																	},
																},
															},
														},
													},
													"preferred_during_scheduling_ignored_during_execution": {
														Computed:    false,
														Optional:    true,
														Description: "",
														Type:        schema.TypeList,
														Elem: &schema.Resource{
															Schema: map[string]*schema.Schema{
																"weight": {
																	Computed:    false,
																	Optional:    true,
																	Description: "",
																	Type:        schema.TypeInt,
																},
																"preference": {
																	Computed:    false,
																	Optional:    true,
																	Description: "",
																	Type:        schema.TypeList,
																	MinItems:    0,
																	MaxItems:    1,
																	Elem: &schema.Resource{
																		Schema: map[string]*schema.Schema{
																			"match_expressions": {
																				Computed:    false,
																				Optional:    true,
																				Description: "",
																				Type:        schema.TypeList,
																				Elem: &schema.Resource{
																					Schema: map[string]*schema.Schema{
																						"key": {
																							Computed:    false,
																							Optional:    true,
																							Description: "",
																							Type:        schema.TypeString,
																						},
																						"operator": {
																							Computed:    false,
																							Optional:    true,
																							Description: "",
																							Type:        schema.TypeString,
																						},
																						"values": {
																							Computed:    false,
																							Optional:    true,
																							Description: "",
																							Type:        schema.TypeList,
																							Elem:        &schema.Schema{Type: schema.TypeString},
																						},
																					},
																				},
																			},
																			"match_fields": {
																				Computed:    false,
																				Optional:    true,
																				Description: "",
																				Type:        schema.TypeList,
																				Elem: &schema.Resource{
																					Schema: map[string]*schema.Schema{
																						"key": {
																							Computed:    false,
																							Optional:    true,
																							Description: "",
																							Type:        schema.TypeString,
																						},
																						"operator": {
																							Computed:    false,
																							Optional:    true,
																							Description: "",
																							Type:        schema.TypeString,
																						},
																						"values": {
																							Computed:    false,
																							Optional:    true,
																							Description: "",
																							Type:        schema.TypeList,
																							Elem:        &schema.Schema{Type: schema.TypeString},
																						},
																					},
																				},
																			},
																		},
																	},
																},
															},
														},
													},
												},
											},
										},
										"pod_affinity": {
											Computed:    false,
											Optional:    true,
											Description: "Pod affinity allows to constrain which nodes your pod is eligible to be scheduled on based on the labels on other pods.",
											Type:        schema.TypeList,
											MinItems:    0,
											MaxItems:    1,
											Elem: &schema.Resource{
												Schema: map[string]*schema.Schema{
													"required_during_scheduling_ignored_during_execution": {
														Computed:    false,
														Optional:    true,
														Description: "",
														Type:        schema.TypeList,
														Elem: &schema.Resource{
															Schema: map[string]*schema.Schema{
																"label_selector": {
																	Computed:    false,
																	Optional:    true,
																	Description: "",
																	Type:        schema.TypeList,
																	MinItems:    0,
																	MaxItems:    1,
																	Elem: &schema.Resource{
																		Schema: map[string]*schema.Schema{
																			"match_labels": {
																				Computed:    false,
																				Optional:    true,
																				Description: "",
																				Type:        schema.TypeMap,
																				Elem:        &schema.Schema{Type: schema.TypeString},
																			},
																			"match_expressions": {
																				Computed:    false,
																				Optional:    true,
																				Description: "",
																				Type:        schema.TypeList,
																				Elem: &schema.Resource{
																					Schema: map[string]*schema.Schema{
																						"key": {
																							Computed:    false,
																							Optional:    true,
																							Description: "",
																							Type:        schema.TypeString,
																						},
																						"operator": {
																							Computed:    false,
																							Optional:    true,
																							Description: "",
																							Type:        schema.TypeString,
																						},
																						"values": {
																							Computed:    false,
																							Optional:    true,
																							Description: "",
																							Type:        schema.TypeList,
																							Elem:        &schema.Schema{Type: schema.TypeString},
																						},
																					},
																				},
																			},
																		},
																	},
																},
																"namespaces": {
																	Computed:    false,
																	Optional:    true,
																	Description: "",
																	Type:        schema.TypeList,
																	Elem:        &schema.Schema{Type: schema.TypeString},
																},
																"topology_key": {
																	Computed:    false,
																	Optional:    true,
																	Description: "",
																	Type:        schema.TypeString,
																},
																"namespace_selector": {
																	Computed:    false,
																	Optional:    true,
																	Description: "",
																	Type:        schema.TypeList,
																	MinItems:    0,
																	MaxItems:    1,
																	Elem: &schema.Resource{
																		Schema: map[string]*schema.Schema{
																			"match_labels": {
																				Computed:    false,
																				Optional:    true,
																				Description: "",
																				Type:        schema.TypeMap,
																				Elem:        &schema.Schema{Type: schema.TypeString},
																			},
																			"match_expressions": {
																				Computed:    false,
																				Optional:    true,
																				Description: "",
																				Type:        schema.TypeList,
																				Elem: &schema.Resource{
																					Schema: map[string]*schema.Schema{
																						"key": {
																							Computed:    false,
																							Optional:    true,
																							Description: "",
																							Type:        schema.TypeString,
																						},
																						"operator": {
																							Computed:    false,
																							Optional:    true,
																							Description: "",
																							Type:        schema.TypeString,
																						},
																						"values": {
																							Computed:    false,
																							Optional:    true,
																							Description: "",
																							Type:        schema.TypeList,
																							Elem:        &schema.Schema{Type: schema.TypeString},
																						},
																					},
																				},
																			},
																		},
																	},
																},
															},
														},
													},
													"preferred_during_scheduling_ignored_during_execution": {
														Computed:    false,
														Optional:    true,
														Description: "",
														Type:        schema.TypeList,
														Elem: &schema.Resource{
															Schema: map[string]*schema.Schema{
																"weight": {
																	Computed:    false,
																	Optional:    true,
																	Description: "",
																	Type:        schema.TypeInt,
																},
																"pod_affinity_term": {
																	Computed:    false,
																	Optional:    true,
																	Description: "",
																	Type:        schema.TypeList,
																	MinItems:    0,
																	MaxItems:    1,
																	Elem: &schema.Resource{
																		Schema: map[string]*schema.Schema{
																			"label_selector": {
																				Computed:    false,
																				Optional:    true,
																				Description: "",
																				Type:        schema.TypeList,
																				MinItems:    0,
																				MaxItems:    1,
																				Elem: &schema.Resource{
																					Schema: map[string]*schema.Schema{
																						"match_labels": {
																							Computed:    false,
																							Optional:    true,
																							Description: "",
																							Type:        schema.TypeMap,
																							Elem:        &schema.Schema{Type: schema.TypeString},
																						},
																						"match_expressions": {
																							Computed:    false,
																							Optional:    true,
																							Description: "",
																							Type:        schema.TypeList,
																							Elem: &schema.Resource{
																								Schema: map[string]*schema.Schema{
																									"key": {
																										Computed:    false,
																										Optional:    true,
																										Description: "",
																										Type:        schema.TypeString,
																									},
																									"operator": {
																										Computed:    false,
																										Optional:    true,
																										Description: "",
																										Type:        schema.TypeString,
																									},
																									"values": {
																										Computed:    false,
																										Optional:    true,
																										Description: "",
																										Type:        schema.TypeList,
																										Elem:        &schema.Schema{Type: schema.TypeString},
																									},
																								},
																							},
																						},
																					},
																				},
																			},
																			"namespaces": {
																				Computed:    false,
																				Optional:    true,
																				Description: "",
																				Type:        schema.TypeList,
																				Elem:        &schema.Schema{Type: schema.TypeString},
																			},
																			"topology_key": {
																				Computed:    false,
																				Optional:    true,
																				Description: "",
																				Type:        schema.TypeString,
																			},
																			"namespace_selector": {
																				Computed:    false,
																				Optional:    true,
																				Description: "",
																				Type:        schema.TypeList,
																				MinItems:    0,
																				MaxItems:    1,
																				Elem: &schema.Resource{
																					Schema: map[string]*schema.Schema{
																						"match_labels": {
																							Computed:    false,
																							Optional:    true,
																							Description: "",
																							Type:        schema.TypeMap,
																							Elem:        &schema.Schema{Type: schema.TypeString},
																						},
																						"match_expressions": {
																							Computed:    false,
																							Optional:    true,
																							Description: "",
																							Type:        schema.TypeList,
																							Elem: &schema.Resource{
																								Schema: map[string]*schema.Schema{
																									"key": {
																										Computed:    false,
																										Optional:    true,
																										Description: "",
																										Type:        schema.TypeString,
																									},
																									"operator": {
																										Computed:    false,
																										Optional:    true,
																										Description: "",
																										Type:        schema.TypeString,
																									},
																									"values": {
																										Computed:    false,
																										Optional:    true,
																										Description: "",
																										Type:        schema.TypeList,
																										Elem:        &schema.Schema{Type: schema.TypeString},
																									},
																								},
																							},
																						},
																					},
																				},
																			},
																		},
																	},
																},
															},
														},
													},
												},
											},
										},
										"pod_anti_affinity": {
											Computed:    false,
											Optional:    true,
											Description: "Pod anti-affinity allows to constrain which nodes your pod is eligible to be scheduled on based on the labels on other pods.",
											Type:        schema.TypeList,
											MinItems:    0,
											MaxItems:    1,
											Elem: &schema.Resource{
												Schema: map[string]*schema.Schema{
													"required_during_scheduling_ignored_during_execution": {
														Computed:    false,
														Optional:    true,
														Description: "",
														Type:        schema.TypeList,
														Elem: &schema.Resource{
															Schema: map[string]*schema.Schema{
																"label_selector": {
																	Computed:    false,
																	Optional:    true,
																	Description: "",
																	Type:        schema.TypeList,
																	MinItems:    0,
																	MaxItems:    1,
																	Elem: &schema.Resource{
																		Schema: map[string]*schema.Schema{
																			"match_labels": {
																				Computed:    false,
																				Optional:    true,
																				Description: "",
																				Type:        schema.TypeMap,
																				Elem:        &schema.Schema{Type: schema.TypeString},
																			},
																			"match_expressions": {
																				Computed:    false,
																				Optional:    true,
																				Description: "",
																				Type:        schema.TypeList,
																				Elem: &schema.Resource{
																					Schema: map[string]*schema.Schema{
																						"key": {
																							Computed:    false,
																							Optional:    true,
																							Description: "",
																							Type:        schema.TypeString,
																						},
																						"operator": {
																							Computed:    false,
																							Optional:    true,
																							Description: "",
																							Type:        schema.TypeString,
																						},
																						"values": {
																							Computed:    false,
																							Optional:    true,
																							Description: "",
																							Type:        schema.TypeList,
																							Elem:        &schema.Schema{Type: schema.TypeString},
																						},
																					},
																				},
																			},
																		},
																	},
																},
																"namespaces": {
																	Computed:    false,
																	Optional:    true,
																	Description: "",
																	Type:        schema.TypeList,
																	Elem:        &schema.Schema{Type: schema.TypeString},
																},
																"topology_key": {
																	Computed:    false,
																	Optional:    true,
																	Description: "",
																	Type:        schema.TypeString,
																},
																"namespace_selector": {
																	Computed:    false,
																	Optional:    true,
																	Description: "",
																	Type:        schema.TypeList,
																	MinItems:    0,
																	MaxItems:    1,
																	Elem: &schema.Resource{
																		Schema: map[string]*schema.Schema{
																			"match_labels": {
																				Computed:    false,
																				Optional:    true,
																				Description: "",
																				Type:        schema.TypeMap,
																				Elem:        &schema.Schema{Type: schema.TypeString},
																			},
																			"match_expressions": {
																				Computed:    false,
																				Optional:    true,
																				Description: "",
																				Type:        schema.TypeList,
																				Elem: &schema.Resource{
																					Schema: map[string]*schema.Schema{
																						"key": {
																							Computed:    false,
																							Optional:    true,
																							Description: "",
																							Type:        schema.TypeString,
																						},
																						"operator": {
																							Computed:    false,
																							Optional:    true,
																							Description: "",
																							Type:        schema.TypeString,
																						},
																						"values": {
																							Computed:    false,
																							Optional:    true,
																							Description: "",
																							Type:        schema.TypeList,
																							Elem:        &schema.Schema{Type: schema.TypeString},
																						},
																					},
																				},
																			},
																		},
																	},
																},
															},
														},
													},
													"preferred_during_scheduling_ignored_during_execution": {
														Computed:    false,
														Optional:    true,
														Description: "",
														Type:        schema.TypeList,
														Elem: &schema.Resource{
															Schema: map[string]*schema.Schema{
																"weight": {
																	Computed:    false,
																	Optional:    true,
																	Description: "",
																	Type:        schema.TypeInt,
																},
																"pod_affinity_term": {
																	Computed:    false,
																	Optional:    true,
																	Description: "",
																	Type:        schema.TypeList,
																	MinItems:    0,
																	MaxItems:    1,
																	Elem: &schema.Resource{
																		Schema: map[string]*schema.Schema{
																			"label_selector": {
																				Computed:    false,
																				Optional:    true,
																				Description: "",
																				Type:        schema.TypeList,
																				MinItems:    0,
																				MaxItems:    1,
																				Elem: &schema.Resource{
																					Schema: map[string]*schema.Schema{
																						"match_labels": {
																							Computed:    false,
																							Optional:    true,
																							Description: "",
																							Type:        schema.TypeMap,
																							Elem:        &schema.Schema{Type: schema.TypeString},
																						},
																						"match_expressions": {
																							Computed:    false,
																							Optional:    true,
																							Description: "",
																							Type:        schema.TypeList,
																							Elem: &schema.Resource{
																								Schema: map[string]*schema.Schema{
																									"key": {
																										Computed:    false,
																										Optional:    true,
																										Description: "",
																										Type:        schema.TypeString,
																									},
																									"operator": {
																										Computed:    false,
																										Optional:    true,
																										Description: "",
																										Type:        schema.TypeString,
																									},
																									"values": {
																										Computed:    false,
																										Optional:    true,
																										Description: "",
																										Type:        schema.TypeList,
																										Elem:        &schema.Schema{Type: schema.TypeString},
																									},
																								},
																							},
																						},
																					},
																				},
																			},
																			"namespaces": {
																				Computed:    false,
																				Optional:    true,
																				Description: "",
																				Type:        schema.TypeList,
																				Elem:        &schema.Schema{Type: schema.TypeString},
																			},
																			"topology_key": {
																				Computed:    false,
																				Optional:    true,
																				Description: "",
																				Type:        schema.TypeString,
																			},
																			"namespace_selector": {
																				Computed:    false,
																				Optional:    true,
																				Description: "",
																				Type:        schema.TypeList,
																				MinItems:    0,
																				MaxItems:    1,
																				Elem: &schema.Resource{
																					Schema: map[string]*schema.Schema{
																						"match_labels": {
																							Computed:    false,
																							Optional:    true,
																							Description: "",
																							Type:        schema.TypeMap,
																							Elem:        &schema.Schema{Type: schema.TypeString},
																						},
																						"match_expressions": {
																							Computed:    false,
																							Optional:    true,
																							Description: "",
																							Type:        schema.TypeList,
																							Elem: &schema.Resource{
																								Schema: map[string]*schema.Schema{
																									"key": {
																										Computed:    false,
																										Optional:    true,
																										Description: "",
																										Type:        schema.TypeString,
																									},
																									"operator": {
																										Computed:    false,
																										Optional:    true,
																										Description: "",
																										Type:        schema.TypeString,
																									},
																									"values": {
																										Computed:    false,
																										Optional:    true,
																										Description: "",
																										Type:        schema.TypeList,
																										Elem:        &schema.Schema{Type: schema.TypeString},
																									},
																								},
																							},
																						},
																					},
																				},
																			},
																		},
																	},
																},
															},
														},
													},
												},
											},
										},
									},
								},
							},
							"image_pull_secrets": {
								Computed:    false,
								Optional:    true,
								Description: "A list of image pull secrets that are used for pulling docker image",
								Type:        schema.TypeList,
								Elem:        &schema.Schema{Type: schema.TypeString},
							},
							"helper_image": {
								Computed:    false,
								Optional:    true,
								Description: "[ADVANCED] Override the default helper image used to clone repos and upload artifacts",
								Type:        schema.TypeString,
							},
							"helper_image_flavor": {
								Computed:    false,
								Optional:    true,
								Description: "Set helper image flavor (alpine, ubuntu), defaults to alpine",
								Type:        schema.TypeString,
							},
							"termination_grace_period_seconds": {
								Computed:    false,
								Optional:    true,
								Description: "Duration after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal.DEPRECATED: use KUBERNETES_POD_TERMINATION_GRACE_PERIOD_SECONDS and KUBERNETES_CLEANUP_GRACE_PERIOD_SECONDS instead.",
								Type:        schema.TypeInt,
							},
							"pod_termination_grace_period_seconds": {
								Computed:    false,
								Optional:    true,
								Description: "Pod-level setting which determines the duration in seconds which the pod has to terminate gracefully. After this, the processes are forcibly halted with a kill signal. Ignored if KUBERNETES_TERMINATIONGRACEPERIODSECONDS is specified.",
								Type:        schema.TypeInt,
							},
							"cleanup_grace_period_seconds": {
								Computed:    false,
								Optional:    true,
								Description: "When cleaning up a pod on completion of a job, the duration in seconds which the pod has to terminate gracefully. After this, the processes are forcibly halted with a kill signal. Ignored if KUBERNETES_TERMINATIONGRACEPERIODSECONDS is specified.",
								Type:        schema.TypeInt,
							},
							"poll_interval": {
								Computed:    false,
								Optional:    true,
								Description: "How frequently, in seconds, the runner will poll the Kubernetes pod it has just created to check its status",
								Type:        schema.TypeInt,
							},
							"poll_timeout": {
								Computed:    false,
								Optional:    true,
								Description: "The total amount of time, in seconds, that needs to pass before the runner will timeout attempting to connect to the pod it has just created (useful for queueing more builds that the cluster can handle at a time)",
								Type:        schema.TypeInt,
							},
							"pod_labels": {
								Computed:    false,
								Optional:    true,
								Description: "A toml table/json object of key-value. Value is expected to be a string. When set, this will create pods with the given pod labels. Environment variables will be substituted for values here.",
								Type:        schema.TypeMap,
								Elem:        &schema.Schema{Type: schema.TypeString},
							},
							"service_account": {
								Computed:    false,
								Optional:    true,
								Description: "Executor pods will use this Service Account to talk to kubernetes API",
								Type:        schema.TypeString,
							},
							"service_account_overwrite_allowed": {
								Computed:    false,
								Optional:    true,
								Description: "Regex to validate 'KUBERNETES_SERVICE_ACCOUNT' value",
								Type:        schema.TypeString,
							},
							"pod_annotations": {
								Computed:    false,
								Optional:    true,
								Description: "A toml table/json object of key-value. Value is expected to be a string. When set, this will create pods with the given annotations. Can be overwritten in build with KUBERNETES_POD_ANNOTATION_* variables",
								Type:        schema.TypeMap,
								Elem:        &schema.Schema{Type: schema.TypeString},
							},
							"pod_annotations_overwrite_allowed": {
								Computed:    false,
								Optional:    true,
								Description: "Regex to validate 'KUBERNETES_POD_ANNOTATIONS_*' values",
								Type:        schema.TypeString,
							},
							"pod_security_context": {
								Computed:    false,
								Optional:    true,
								Description: "A security context attached to each build pod",
								Type:        schema.TypeList,
								MinItems:    0,
								MaxItems:    1,
								Elem: &schema.Resource{
									Schema: map[string]*schema.Schema{
										"fs_group": {
											Computed:    false,
											Optional:    true,
											Description: "A special supplemental group that applies to all containers in a pod",
											Type:        schema.TypeInt,
										},
										"run_as_group": {
											Computed:    false,
											Optional:    true,
											Description: "The GID to run the entrypoint of the container process",
											Type:        schema.TypeInt,
										},
										"run_as_non_root": {
											Computed:    false,
											Optional:    true,
											Description: "Indicates that the container must run as a non-root user",
											Type:        schema.TypeBool,
										},
										"run_as_user": {
											Computed:    false,
											Optional:    true,
											Description: "The UID to run the entrypoint of the container process",
											Type:        schema.TypeInt,
										},
										"supplemental_groups": {
											Computed:    false,
											Optional:    true,
											Description: "A list of groups applied to the first process run in each container, in addition to the container's primary GID",
											Type:        schema.TypeList,
											Elem:        &schema.Schema{Type: schema.TypeInt},
										},
									},
								},
							},
							"build_container_security_context": {
								Computed:    false,
								Optional:    true,
								Description: "A security context attached to the build container inside the build pod",
								Type:        schema.TypeList,
								MinItems:    0,
								MaxItems:    1,
								Elem: &schema.Resource{
									Schema: map[string]*schema.Schema{
										"capabilities": {
											Computed:    false,
											Optional:    true,
											Description: "The capabilities to add/drop when running the container",
											Type:        schema.TypeList,
											MinItems:    0,
											MaxItems:    1,
											Elem: &schema.Resource{
												Schema: map[string]*schema.Schema{
													"add": {
														Computed:    false,
														Optional:    true,
														Description: "List of capabilities to add to the build container",
														Type:        schema.TypeList,
														Elem:        &schema.Schema{Type: schema.TypeString},
													},
													"drop": {
														Computed:    false,
														Optional:    true,
														Description: "List of capabilities to drop from the build container",
														Type:        schema.TypeList,
														Elem:        &schema.Schema{Type: schema.TypeString},
													},
												},
											},
										},
										"privileged": {
											Computed:    false,
											Optional:    true,
											Description: "Run container in privileged mode",
											Type:        schema.TypeBool,
										},
										"run_as_user": {
											Computed:    false,
											Optional:    true,
											Description: "The UID to run the entrypoint of the container process",
											Type:        schema.TypeInt,
										},
										"run_as_group": {
											Computed:    false,
											Optional:    true,
											Description: "The GID to run the entrypoint of the container process",
											Type:        schema.TypeInt,
										},
										"run_as_non_root": {
											Computed:    false,
											Optional:    true,
											Description: "Indicates that the container must run as a non-root user",
											Type:        schema.TypeBool,
										},
										"read_only_root_filesystem": {
											Computed:    false,
											Optional:    true,
											Description: " Whether this container has a read-only root filesystem.",
											Type:        schema.TypeBool,
										},
										"allow_privilege_escalation": {
											Computed:    false,
											Optional:    true,
											Description: "AllowPrivilegeEscalation controls whether a process can gain more privileges than its parent process",
											Type:        schema.TypeBool,
										},
									},
								},
							},
							"helper_container_security_context": {
								Computed:    false,
								Optional:    true,
								Description: "A security context attached to the helper container inside the build pod",
								Type:        schema.TypeList,
								MinItems:    0,
								MaxItems:    1,
								Elem: &schema.Resource{
									Schema: map[string]*schema.Schema{
										"capabilities": {
											Computed:    false,
											Optional:    true,
											Description: "The capabilities to add/drop when running the container",
											Type:        schema.TypeList,
											MinItems:    0,
											MaxItems:    1,
											Elem: &schema.Resource{
												Schema: map[string]*schema.Schema{
													"add": {
														Computed:    false,
														Optional:    true,
														Description: "List of capabilities to add to the build container",
														Type:        schema.TypeList,
														Elem:        &schema.Schema{Type: schema.TypeString},
													},
													"drop": {
														Computed:    false,
														Optional:    true,
														Description: "List of capabilities to drop from the build container",
														Type:        schema.TypeList,
														Elem:        &schema.Schema{Type: schema.TypeString},
													},
												},
											},
										},
										"privileged": {
											Computed:    false,
											Optional:    true,
											Description: "Run container in privileged mode",
											Type:        schema.TypeBool,
										},
										"run_as_user": {
											Computed:    false,
											Optional:    true,
											Description: "The UID to run the entrypoint of the container process",
											Type:        schema.TypeInt,
										},
										"run_as_group": {
											Computed:    false,
											Optional:    true,
											Description: "The GID to run the entrypoint of the container process",
											Type:        schema.TypeInt,
										},
										"run_as_non_root": {
											Computed:    false,
											Optional:    true,
											Description: "Indicates that the container must run as a non-root user",
											Type:        schema.TypeBool,
										},
										"read_only_root_filesystem": {
											Computed:    false,
											Optional:    true,
											Description: " Whether this container has a read-only root filesystem.",
											Type:        schema.TypeBool,
										},
										"allow_privilege_escalation": {
											Computed:    false,
											Optional:    true,
											Description: "AllowPrivilegeEscalation controls whether a process can gain more privileges than its parent process",
											Type:        schema.TypeBool,
										},
									},
								},
							},
							"service_container_security_context": {
								Computed:    false,
								Optional:    true,
								Description: "A security context attached to the service containers inside the build pod",
								Type:        schema.TypeList,
								MinItems:    0,
								MaxItems:    1,
								Elem: &schema.Resource{
									Schema: map[string]*schema.Schema{
										"capabilities": {
											Computed:    false,
											Optional:    true,
											Description: "The capabilities to add/drop when running the container",
											Type:        schema.TypeList,
											MinItems:    0,
											MaxItems:    1,
											Elem: &schema.Resource{
												Schema: map[string]*schema.Schema{
													"add": {
														Computed:    false,
														Optional:    true,
														Description: "List of capabilities to add to the build container",
														Type:        schema.TypeList,
														Elem:        &schema.Schema{Type: schema.TypeString},
													},
													"drop": {
														Computed:    false,
														Optional:    true,
														Description: "List of capabilities to drop from the build container",
														Type:        schema.TypeList,
														Elem:        &schema.Schema{Type: schema.TypeString},
													},
												},
											},
										},
										"privileged": {
											Computed:    false,
											Optional:    true,
											Description: "Run container in privileged mode",
											Type:        schema.TypeBool,
										},
										"run_as_user": {
											Computed:    false,
											Optional:    true,
											Description: "The UID to run the entrypoint of the container process",
											Type:        schema.TypeInt,
										},
										"run_as_group": {
											Computed:    false,
											Optional:    true,
											Description: "The GID to run the entrypoint of the container process",
											Type:        schema.TypeInt,
										},
										"run_as_non_root": {
											Computed:    false,
											Optional:    true,
											Description: "Indicates that the container must run as a non-root user",
											Type:        schema.TypeBool,
										},
										"read_only_root_filesystem": {
											Computed:    false,
											Optional:    true,
											Description: " Whether this container has a read-only root filesystem.",
											Type:        schema.TypeBool,
										},
										"allow_privilege_escalation": {
											Computed:    false,
											Optional:    true,
											Description: "AllowPrivilegeEscalation controls whether a process can gain more privileges than its parent process",
											Type:        schema.TypeBool,
										},
									},
								},
							},
							"volumes": {
								Computed:    false,
								Optional:    true,
								Description: "",
								Type:        schema.TypeList,
								MinItems:    0,
								MaxItems:    1,
								Elem: &schema.Resource{
									Schema: map[string]*schema.Schema{
										"host_path": {
											Computed:    false,
											Optional:    true,
											Description: "The host paths which will be mounted",
											Type:        schema.TypeList,
											Elem: &schema.Resource{
												Schema: map[string]*schema.Schema{
													"name": {
														Computed:    false,
														Optional:    true,
														Description: "The name of the volume",
														Type:        schema.TypeString,
													},
													"mount_path": {
														Computed:    false,
														Optional:    true,
														Description: "Path where volume should be mounted inside of container",
														Type:        schema.TypeString,
													},
													"sub_path": {
														Computed:    false,
														Optional:    true,
														Description: "The sub-path of the volume to mount (defaults to volume root)",
														Type:        schema.TypeString,
													},
													"read_only": {
														Computed:    false,
														Optional:    true,
														Description: "If this volume should be mounted read only",
														Type:        schema.TypeBool,
													},
													"host_path": {
														Computed:    false,
														Optional:    true,
														Description: "Path from the host that should be mounted as a volume",
														Type:        schema.TypeString,
													},
												},
											},
										},
										"pvc": {
											Computed:    false,
											Optional:    true,
											Description: "The persistent volume claims that will be mounted",
											Type:        schema.TypeList,
											Elem: &schema.Resource{
												Schema: map[string]*schema.Schema{
													"name": {
														Computed:    false,
														Optional:    true,
														Description: "The name of the volume and PVC to use",
														Type:        schema.TypeString,
													},
													"mount_path": {
														Computed:    false,
														Optional:    true,
														Description: "Path where volume should be mounted inside of container",
														Type:        schema.TypeString,
													},
													"sub_path": {
														Computed:    false,
														Optional:    true,
														Description: "The sub-path of the volume to mount (defaults to volume root)",
														Type:        schema.TypeString,
													},
													"read_only": {
														Computed:    false,
														Optional:    true,
														Description: "If this volume should be mounted read only",
														Type:        schema.TypeBool,
													},
												},
											},
										},
										"config_map": {
											Computed:    false,
											Optional:    true,
											Description: "The config maps which will be mounted as volumes",
											Type:        schema.TypeList,
											Elem: &schema.Resource{
												Schema: map[string]*schema.Schema{
													"name": {
														Computed:    false,
														Optional:    true,
														Description: "The name of the volume and ConfigMap to use",
														Type:        schema.TypeString,
													},
													"mount_path": {
														Computed:    false,
														Optional:    true,
														Description: "Path where volume should be mounted inside of container",
														Type:        schema.TypeString,
													},
													"sub_path": {
														Computed:    false,
														Optional:    true,
														Description: "The sub-path of the volume to mount (defaults to volume root)",
														Type:        schema.TypeString,
													},
													"read_only": {
														Computed:    false,
														Optional:    true,
														Description: "If this volume should be mounted read only",
														Type:        schema.TypeBool,
													},
													"items": {
														Computed:    false,
														Optional:    true,
														Description: "Key-to-path mapping for keys from the config map that should be used.",
														Type:        schema.TypeMap,
														Elem:        &schema.Schema{Type: schema.TypeString},
													},
												},
											},
										},
										"secret": {
											Computed:    false,
											Optional:    true,
											Description: "The secret maps which will be mounted",
											Type:        schema.TypeList,
											Elem: &schema.Resource{
												Schema: map[string]*schema.Schema{
													"name": {
														Computed:    false,
														Optional:    true,
														Description: "The name of the volume and Secret to use",
														Type:        schema.TypeString,
													},
													"mount_path": {
														Computed:    false,
														Optional:    true,
														Description: "Path where volume should be mounted inside of container",
														Type:        schema.TypeString,
													},
													"sub_path": {
														Computed:    false,
														Optional:    true,
														Description: "The sub-path of the volume to mount (defaults to volume root)",
														Type:        schema.TypeString,
													},
													"read_only": {
														Computed:    false,
														Optional:    true,
														Description: "If this volume should be mounted read only",
														Type:        schema.TypeBool,
													},
													"items": {
														Computed:    false,
														Optional:    true,
														Description: "Key-to-path mapping for keys from the secret that should be used.",
														Type:        schema.TypeMap,
														Elem:        &schema.Schema{Type: schema.TypeString},
													},
												},
											},
										},
										"empty_dir": {
											Computed:    false,
											Optional:    true,
											Description: "The empty dirs which will be mounted",
											Type:        schema.TypeList,
											Elem: &schema.Resource{
												Schema: map[string]*schema.Schema{
													"name": {
														Computed:    false,
														Optional:    true,
														Description: "The name of the volume and EmptyDir to use",
														Type:        schema.TypeString,
													},
													"mount_path": {
														Computed:    false,
														Optional:    true,
														Description: "Path where volume should be mounted inside of container",
														Type:        schema.TypeString,
													},
													"sub_path": {
														Computed:    false,
														Optional:    true,
														Description: "The sub-path of the volume to mount (defaults to volume root)",
														Type:        schema.TypeString,
													},
													"medium": {
														Computed:    false,
														Optional:    true,
														Description: "Set to 'Memory' to have a tmpfs",
														Type:        schema.TypeString,
													},
												},
											},
										},
										"csi": {
											Computed:    false,
											Optional:    true,
											Description: "The CSI volumes which will be mounted",
											Type:        schema.TypeList,
											Elem: &schema.Resource{
												Schema: map[string]*schema.Schema{
													"name": {
														Computed:    false,
														Optional:    true,
														Description: "The name of the CSI volume and volumeMount to use",
														Type:        schema.TypeString,
													},
													"mount_path": {
														Computed:    false,
														Optional:    true,
														Description: "Path where volume should be mounted inside of container",
														Type:        schema.TypeString,
													},
													"sub_path": {
														Computed:    false,
														Optional:    true,
														Description: "The sub-path of the volume to mount (defaults to volume root)",
														Type:        schema.TypeString,
													},
													"driver": {
														Computed:    false,
														Optional:    true,
														Description: "A string value that specifies the name of the volume driver to use.",
														Type:        schema.TypeString,
													},
													"fs_type": {
														Computed:    false,
														Optional:    true,
														Description: "Filesystem type to mount. If not provided, the empty value is passed to the associated CSI driver which will determine the default filesystem to apply.",
														Type:        schema.TypeString,
													},
													"read_only": {
														Computed:    false,
														Optional:    true,
														Description: "If this volume should be mounted read only",
														Type:        schema.TypeBool,
													},
													"volume_attributes": {
														Computed:    false,
														Optional:    true,
														Description: "Key-value pair mapping for attributes of the CSI volume.",
														Type:        schema.TypeMap,
														Elem:        &schema.Schema{Type: schema.TypeString},
													},
												},
											},
										},
									},
								},
							},
							"host_aliases": {
								Computed:    false,
								Optional:    true,
								Description: "Add a custom host-to-IP mapping",
								Type:        schema.TypeList,
								Elem: &schema.Resource{
									Schema: map[string]*schema.Schema{
										"ip": {
											Computed:    false,
											Optional:    true,
											Description: "The IP address you want to attach hosts to",
											Type:        schema.TypeString,
										},
										"hostnames": {
											Computed:    false,
											Optional:    true,
											Description: "A list of hostnames that will be attached to the IP",
											Type:        schema.TypeList,
											Elem:        &schema.Schema{Type: schema.TypeString},
										},
									},
								},
							},
							"services": {
								Computed:    false,
								Optional:    true,
								Description: "Add service that is started with container",
								Type:        schema.TypeList,
								Elem: &schema.Resource{
									Schema: map[string]*schema.Schema{
										"name": {
											Computed:    false,
											Optional:    true,
											Description: "The image path for the service",
											Type:        schema.TypeString,
										},
										"alias": {
											Computed:    false,
											Optional:    true,
											Description: "The alias of the service",
											Type:        schema.TypeString,
										},
										"command": {
											Computed:    false,
											Optional:    true,
											Description: "Command or script that should be used as the container’s command. Syntax is similar to https://docs.docker.com/engine/reference/builder/#cmd",
											Type:        schema.TypeList,
											Elem:        &schema.Schema{Type: schema.TypeString},
										},
										"entrypoint": {
											Computed:    false,
											Optional:    true,
											Description: "Command or script that should be executed as the container’s entrypoint. syntax is similar to https://docs.docker.com/engine/reference/builder/#entrypoint",
											Type:        schema.TypeList,
											Elem:        &schema.Schema{Type: schema.TypeString},
										},
									},
								},
							},
							"cap_add": {
								Computed:    false,
								Optional:    true,
								Description: "Add Linux capabilities",
								Type:        schema.TypeList,
								Elem:        &schema.Schema{Type: schema.TypeString},
							},
							"cap_drop": {
								Computed:    false,
								Optional:    true,
								Description: "Drop Linux capabilities",
								Type:        schema.TypeList,
								Elem:        &schema.Schema{Type: schema.TypeString},
							},
							"dns_policy": {
								Computed:    false,
								Optional:    true,
								Description: "How Kubernetes should try to resolve DNS from the created pods. If unset, Kubernetes will use the default 'ClusterFirst'. Valid values are: none, default, cluster-first, cluster-first-with-host-net",
								Type:        schema.TypeString,
							},
							"dns_config": {
								Computed:    false,
								Optional:    true,
								Description: "Pod DNS config",
								Type:        schema.TypeList,
								MinItems:    0,
								MaxItems:    1,
								Elem: &schema.Resource{
									Schema: map[string]*schema.Schema{
										"nameservers": {
											Computed:    false,
											Optional:    true,
											Description: "A list of IP addresses that will be used as DNS servers for the Pod.",
											Type:        schema.TypeList,
											Elem:        &schema.Schema{Type: schema.TypeString},
										},
										"options": {
											Computed:    false,
											Optional:    true,
											Description: "An optional list of objects where each object may have a name property (required) and a value property (optional).",
											Type:        schema.TypeList,
											Elem: &schema.Resource{
												Schema: map[string]*schema.Schema{
													"name": {
														Computed:    false,
														Optional:    true,
														Description: "",
														Type:        schema.TypeString,
													},
													"value": {
														Computed:    false,
														Optional:    true,
														Description: "",
														Type:        schema.TypeString,
													},
												},
											},
										},
										"searches": {
											Computed:    false,
											Optional:    true,
											Description: "A list of DNS search domains for hostname lookup in the Pod.",
											Type:        schema.TypeList,
											Elem:        &schema.Schema{Type: schema.TypeString},
										},
									},
								},
							},
							"container_lifecycle": {
								Computed:    false,
								Optional:    true,
								Description: "Actions that the management system should take in response to container lifecycle events",
								Type:        schema.TypeList,
								MinItems:    0,
								MaxItems:    1,
								Elem: &schema.Resource{
									Schema: map[string]*schema.Schema{
										"post_start": {
											Computed:    false,
											Optional:    true,
											Description: "PostStart is called immediately after a container is created. If the handler fails, the container is terminated and restarted according to its restart policy. Other management of the container blocks until the hook completes",
											Type:        schema.TypeList,
											MinItems:    0,
											MaxItems:    1,
											Elem: &schema.Resource{
												Schema: map[string]*schema.Schema{
													"exec": {
														Computed:    false,
														Optional:    true,
														Description: "Exec specifies the action to take",
														Type:        schema.TypeList,
														MinItems:    0,
														MaxItems:    1,
														Elem: &schema.Resource{
															Schema: map[string]*schema.Schema{
																"command": {
																	Computed:    false,
																	Optional:    true,
																	Description: "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy",
																	Type:        schema.TypeList,
																	Elem:        &schema.Schema{Type: schema.TypeString},
																},
															},
														},
													},
													"http_get": {
														Computed:    false,
														Optional:    true,
														Description: "HTTPGet specifies the http request to perform.",
														Type:        schema.TypeList,
														MinItems:    0,
														MaxItems:    1,
														Elem: &schema.Resource{
															Schema: map[string]*schema.Schema{
																"host": {
																	Computed:    false,
																	Optional:    true,
																	Description: "Host name to connect to, defaults to the pod IP. You probably want to set \"Host\" in httpHeaders instead",
																	Type:        schema.TypeString,
																},
																"http_headers": {
																	Computed:    false,
																	Optional:    true,
																	Description: "Custom headers to set in the request. HTTP allows repeated headers",
																	Type:        schema.TypeList,
																	Elem: &schema.Resource{
																		Schema: map[string]*schema.Schema{
																			"name": {
																				Computed:    false,
																				Optional:    true,
																				Description: "The header field name",
																				Type:        schema.TypeString,
																			},
																			"value": {
																				Computed:    false,
																				Optional:    true,
																				Description: "The header field value",
																				Type:        schema.TypeString,
																			},
																		},
																	},
																},
																"path": {
																	Computed:    false,
																	Optional:    true,
																	Description: "Path to access on the HTTP server",
																	Type:        schema.TypeString,
																},
																"port": {
																	Computed:    false,
																	Optional:    true,
																	Description: "Number of the port to access on the container. Number must be in the range 1 to 65535",
																	Type:        schema.TypeInt,
																},
																"scheme": {
																	Computed:    false,
																	Optional:    true,
																	Description: "Scheme to use for connecting to the host. Defaults to HTTP",
																	Type:        schema.TypeString,
																},
															},
														},
													},
													"tcp_socket": {
														Computed:    false,
														Optional:    true,
														Description: "TCPSocket specifies an action involving a TCP port",
														Type:        schema.TypeList,
														MinItems:    0,
														MaxItems:    1,
														Elem: &schema.Resource{
															Schema: map[string]*schema.Schema{
																"host": {
																	Computed:    false,
																	Optional:    true,
																	Description: "Host name to connect to, defaults to the pod IP. You probably want to set \"Host\" in httpHeaders instead",
																	Type:        schema.TypeString,
																},
																"port": {
																	Computed:    false,
																	Optional:    true,
																	Description: "Number of the port to access on the container. Number must be in the range 1 to 65535",
																	Type:        schema.TypeInt,
																},
															},
														},
													},
												},
											},
										},
										"pre_stop": {
											Computed:    false,
											Optional:    true,
											Description: "PreStop is called immediately before a container is terminated due to an API request or management event such as liveness/startup probe failure, preemption, resource contention, etc. The handler is not called if the container crashes or exits. The reason for termination is passed to the handler. The Pod's termination grace period countdown begins before the PreStop hooked is executed. Regardless of the outcome of the handler, the container will eventually terminate within the Pod's termination grace period. Other management of the container blocks until the hook completes or until the termination grace period is reached",
											Type:        schema.TypeList,
											MinItems:    0,
											MaxItems:    1,
											Elem: &schema.Resource{
												Schema: map[string]*schema.Schema{
													"exec": {
														Computed:    false,
														Optional:    true,
														Description: "Exec specifies the action to take",
														Type:        schema.TypeList,
														MinItems:    0,
														MaxItems:    1,
														Elem: &schema.Resource{
															Schema: map[string]*schema.Schema{
																"command": {
																	Computed:    false,
																	Optional:    true,
																	Description: "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy",
																	Type:        schema.TypeList,
																	Elem:        &schema.Schema{Type: schema.TypeString},
																},
															},
														},
													},
													"http_get": {
														Computed:    false,
														Optional:    true,
														Description: "HTTPGet specifies the http request to perform.",
														Type:        schema.TypeList,
														MinItems:    0,
														MaxItems:    1,
														Elem: &schema.Resource{
															Schema: map[string]*schema.Schema{
																"host": {
																	Computed:    false,
																	Optional:    true,
																	Description: "Host name to connect to, defaults to the pod IP. You probably want to set \"Host\" in httpHeaders instead",
																	Type:        schema.TypeString,
																},
																"http_headers": {
																	Computed:    false,
																	Optional:    true,
																	Description: "Custom headers to set in the request. HTTP allows repeated headers",
																	Type:        schema.TypeList,
																	Elem: &schema.Resource{
																		Schema: map[string]*schema.Schema{
																			"name": {
																				Computed:    false,
																				Optional:    true,
																				Description: "The header field name",
																				Type:        schema.TypeString,
																			},
																			"value": {
																				Computed:    false,
																				Optional:    true,
																				Description: "The header field value",
																				Type:        schema.TypeString,
																			},
																		},
																	},
																},
																"path": {
																	Computed:    false,
																	Optional:    true,
																	Description: "Path to access on the HTTP server",
																	Type:        schema.TypeString,
																},
																"port": {
																	Computed:    false,
																	Optional:    true,
																	Description: "Number of the port to access on the container. Number must be in the range 1 to 65535",
																	Type:        schema.TypeInt,
																},
																"scheme": {
																	Computed:    false,
																	Optional:    true,
																	Description: "Scheme to use for connecting to the host. Defaults to HTTP",
																	Type:        schema.TypeString,
																},
															},
														},
													},
													"tcp_socket": {
														Computed:    false,
														Optional:    true,
														Description: "TCPSocket specifies an action involving a TCP port",
														Type:        schema.TypeList,
														MinItems:    0,
														MaxItems:    1,
														Elem: &schema.Resource{
															Schema: map[string]*schema.Schema{
																"host": {
																	Computed:    false,
																	Optional:    true,
																	Description: "Host name to connect to, defaults to the pod IP. You probably want to set \"Host\" in httpHeaders instead",
																	Type:        schema.TypeString,
																},
																"port": {
																	Computed:    false,
																	Optional:    true,
																	Description: "Number of the port to access on the container. Number must be in the range 1 to 65535",
																	Type:        schema.TypeInt,
																},
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
				"custom": {
					Computed:    false,
					Optional:    true,
					Description: "",
					Type:        schema.TypeList,
					MinItems:    0,
					MaxItems:    1,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"config_exec": {
								Computed:    false,
								Optional:    true,
								Description: "Executable that allows to inject configuration values to the executor",
								Type:        schema.TypeString,
							},
							"config_args": {
								Computed:    false,
								Optional:    true,
								Description: "Arguments for the config executable",
								Type:        schema.TypeList,
								Elem:        &schema.Schema{Type: schema.TypeString},
							},
							"config_exec_timeout": {
								Computed:    false,
								Optional:    true,
								Description: "Timeout for the config executable (in seconds)",
								Type:        schema.TypeInt,
							},
							"prepare_exec": {
								Computed:    false,
								Optional:    true,
								Description: "Executable that prepares executor",
								Type:        schema.TypeString,
							},
							"prepare_args": {
								Computed:    false,
								Optional:    true,
								Description: "Arguments for the prepare executable",
								Type:        schema.TypeList,
								Elem:        &schema.Schema{Type: schema.TypeString},
							},
							"prepare_exec_timeout": {
								Computed:    false,
								Optional:    true,
								Description: "Timeout for the prepare executable (in seconds)",
								Type:        schema.TypeInt,
							},
							"run_exec": {
								Computed:    false,
								Optional:    true,
								Description: "Executable that runs the job script in executor",
								Type:        schema.TypeString,
							},
							"run_args": {
								Computed:    false,
								Optional:    true,
								Description: "Arguments for the run executable",
								Type:        schema.TypeList,
								Elem:        &schema.Schema{Type: schema.TypeString},
							},
							"cleanup_exec": {
								Computed:    false,
								Optional:    true,
								Description: "Executable that cleanups after executor run",
								Type:        schema.TypeString,
							},
							"cleanup_args": {
								Computed:    false,
								Optional:    true,
								Description: "Arguments for the cleanup executable",
								Type:        schema.TypeList,
								Elem:        &schema.Schema{Type: schema.TypeString},
							},
							"cleanup_exec_timeout": {
								Computed:    false,
								Optional:    true,
								Description: "Timeout for the cleanup executable (in seconds)",
								Type:        schema.TypeInt,
							},
							"graceful_kill_timeout": {
								Computed:    false,
								Optional:    true,
								Description: "Graceful timeout for scripts execution after SIGTERM is sent to the process (in seconds). This limits the time given for scripts to perform the cleanup before exiting",
								Type:        schema.TypeInt,
							},
							"force_kill_timeout": {
								Computed:    false,
								Optional:    true,
								Description: "Force timeout for scripts execution (in seconds). Counted from the force kill call; if process will be not terminated, Runner will abandon process termination and log an error",
								Type:        schema.TypeInt,
							},
						},
					},
				},
			},
		},
	},
	"sentry_dsn": {
		Computed:    false,
		Optional:    true,
		Description: "",
		Type:        schema.TypeString,
	},
}

func dataSourceGitlabCIRunnerConfigReadNEW(d *schema.ResourceData, meta interface{}) error {

	// c := config.Config{}

	// ListenAddress: listen_address -- string, string
	// if v, ok := d.GetOk("listen_address"); ok {
	//c.ListenAddress = v.(FIXME type)
	// }
	// SessionServer: session_server -- SessionServer, config.SessionServer
	// if v, ok := d.GetOk("session_server"); ok {
	//c.SessionServer = v.(FIXME type)
	// }
	// Concurrent: concurrent -- int, int
	// if v, ok := d.GetOk("concurrent"); ok {
	//c.Concurrent = v.(FIXME type)
	// }
	// CheckInterval: check_interval -- int, int
	// if v, ok := d.GetOk("check_interval"); ok {
	//c.CheckInterval = v.(FIXME type)
	// }
	// LogLevel: log_level -- , *string
	// if v, ok := d.GetOk("log_level"); ok {
	//c.LogLevel = v.(FIXME type)
	// }
	// LogFormat: log_format -- , *string
	// if v, ok := d.GetOk("log_format"); ok {
	//c.LogFormat = v.(FIXME type)
	// }
	// User: user -- string, string
	// if v, ok := d.GetOk("user"); ok {
	//c.User = v.(FIXME type)
	// }
	// Runners: runners -- , []*config.RunnerConfig
	// if v, ok := d.GetOk("runners"); ok {
	//c.Runners = v.(FIXME type)
	// }
	// SentryDSN: sentry_dsn -- , *string
	// if v, ok := d.GetOk("sentry_dsn"); ok {
	//c.SentryDSN = v.(FIXME type)
	// }

	return nil
}

// config.CacheAzureConfig

func dsRunnerConfigReadStructConfigCacheAzureConfig(ctx context.Context, prefix string, d *schema.ResourceData) (config.CacheAzureConfig, error) {

	if prefix != "" {
		prefix = prefix + "."
	}

	tflog.Debug(ctx, fmt.Sprintf("beginning dsRunnerConfigReadStructConfigCacheAzureConfig run; prefix is '%s'", prefix))

	val := config.CacheAzureConfig{}

	// AccountName: account_name -- string, string
	if v, ok := d.GetOk(prefix + "account_name"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "account_name"))
		val.AccountName = v.(string)
	}

	// AccountKey: account_key -- string, string
	if v, ok := d.GetOk(prefix + "account_key"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "account_key"))
		val.AccountKey = v.(string)
	}

	// ContainerName: container_name -- string, string
	if v, ok := d.GetOk(prefix + "container_name"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "container_name"))
		val.ContainerName = v.(string)
	}

	// StorageDomain: storage_domain -- string, string
	if v, ok := d.GetOk(prefix + "storage_domain"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "storage_domain"))
		val.StorageDomain = v.(string)
	}

	return val, nil
}

// config.CacheAzureCredentials

func dsRunnerConfigReadStructConfigCacheAzureCredentials(ctx context.Context, prefix string, d *schema.ResourceData) (config.CacheAzureCredentials, error) {

	if prefix != "" {
		prefix = prefix + "."
	}

	tflog.Debug(ctx, fmt.Sprintf("beginning dsRunnerConfigReadStructConfigCacheAzureCredentials run; prefix is '%s'", prefix))

	val := config.CacheAzureCredentials{}

	// AccountName: account_name -- string, string
	if v, ok := d.GetOk(prefix + "account_name"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "account_name"))
		val.AccountName = v.(string)
	}

	// AccountKey: account_key -- string, string
	if v, ok := d.GetOk(prefix + "account_key"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "account_key"))
		val.AccountKey = v.(string)
	}

	return val, nil
}

// config.CacheConfig

func dsRunnerConfigReadStructConfigCacheConfig(ctx context.Context, prefix string, d *schema.ResourceData) (config.CacheConfig, error) {

	if prefix != "" {
		prefix = prefix + "."
	}

	tflog.Debug(ctx, fmt.Sprintf("beginning dsRunnerConfigReadStructConfigCacheConfig run; prefix is '%s'", prefix))

	val := config.CacheConfig{}

	// Type: type -- string, string
	if v, ok := d.GetOk(prefix + "type"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "type"))
		val.Type = v.(string)
	}

	// Path: path -- string, string
	if v, ok := d.GetOk(prefix + "path"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "path"))
		val.Path = v.(string)
	}

	// Shared: shared -- bool, bool
	if v, ok := d.GetOk(prefix + "shared"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "shared"))
		val.Shared = v.(bool)
	}

	// S3: s3 -- , *config.CacheS3Config

	if _, ok := d.GetOk(prefix + "s3.0"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "s3"))
		thing, err := dsRunnerConfigReadStructConfigCacheS3Config(ctx, prefix+"s3.0", d)
		if err != nil {
			return val, err
		}
		val.S3 = &thing
	} else {
		tflog.Debug(ctx, fmt.Sprintf("not set: %s", prefix+"s3"))
	}

	// GCS: gcs -- , *config.CacheGCSConfig

	if _, ok := d.GetOk(prefix + "gcs.0"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "gcs"))
		thing, err := dsRunnerConfigReadStructConfigCacheGCSConfig(ctx, prefix+"gcs.0", d)
		if err != nil {
			return val, err
		}
		val.GCS = &thing
	} else {
		tflog.Debug(ctx, fmt.Sprintf("not set: %s", prefix+"gcs"))
	}

	// Azure: azure -- , *config.CacheAzureConfig

	if _, ok := d.GetOk(prefix + "azure.0"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "azure"))
		thing, err := dsRunnerConfigReadStructConfigCacheAzureConfig(ctx, prefix+"azure.0", d)
		if err != nil {
			return val, err
		}
		val.Azure = &thing
	} else {
		tflog.Debug(ctx, fmt.Sprintf("not set: %s", prefix+"azure"))
	}

	return val, nil
}

// config.CacheGCSConfig

func dsRunnerConfigReadStructConfigCacheGCSConfig(ctx context.Context, prefix string, d *schema.ResourceData) (config.CacheGCSConfig, error) {

	if prefix != "" {
		prefix = prefix + "."
	}

	tflog.Debug(ctx, fmt.Sprintf("beginning dsRunnerConfigReadStructConfigCacheGCSConfig run; prefix is '%s'", prefix))

	val := config.CacheGCSConfig{}

	// AccessID: access_id -- string, string
	if v, ok := d.GetOk(prefix + "access_id"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "access_id"))
		val.AccessID = v.(string)
	}

	// PrivateKey: private_key -- string, string
	if v, ok := d.GetOk(prefix + "private_key"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "private_key"))
		val.PrivateKey = v.(string)
	}

	// CredentialsFile: credentials_file -- string, string
	if v, ok := d.GetOk(prefix + "credentials_file"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "credentials_file"))
		val.CredentialsFile = v.(string)
	}

	// BucketName: bucket_name -- string, string
	if v, ok := d.GetOk(prefix + "bucket_name"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "bucket_name"))
		val.BucketName = v.(string)
	}

	return val, nil
}

// config.CacheGCSCredentials

func dsRunnerConfigReadStructConfigCacheGCSCredentials(ctx context.Context, prefix string, d *schema.ResourceData) (config.CacheGCSCredentials, error) {

	if prefix != "" {
		prefix = prefix + "."
	}

	tflog.Debug(ctx, fmt.Sprintf("beginning dsRunnerConfigReadStructConfigCacheGCSCredentials run; prefix is '%s'", prefix))

	val := config.CacheGCSCredentials{}

	// AccessID: access_id -- string, string
	if v, ok := d.GetOk(prefix + "access_id"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "access_id"))
		val.AccessID = v.(string)
	}

	// PrivateKey: private_key -- string, string
	if v, ok := d.GetOk(prefix + "private_key"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "private_key"))
		val.PrivateKey = v.(string)
	}

	return val, nil
}

// config.CacheS3Config

func dsRunnerConfigReadStructConfigCacheS3Config(ctx context.Context, prefix string, d *schema.ResourceData) (config.CacheS3Config, error) {

	if prefix != "" {
		prefix = prefix + "."
	}

	tflog.Debug(ctx, fmt.Sprintf("beginning dsRunnerConfigReadStructConfigCacheS3Config run; prefix is '%s'", prefix))

	val := config.CacheS3Config{}

	// ServerAddress: server_address -- string, string
	if v, ok := d.GetOk(prefix + "server_address"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "server_address"))
		val.ServerAddress = v.(string)
	}

	// AccessKey: access_key -- string, string
	if v, ok := d.GetOk(prefix + "access_key"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "access_key"))
		val.AccessKey = v.(string)
	}

	// SecretKey: secret_key -- string, string
	if v, ok := d.GetOk(prefix + "secret_key"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "secret_key"))
		val.SecretKey = v.(string)
	}

	// BucketName: bucket_name -- string, string
	if v, ok := d.GetOk(prefix + "bucket_name"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "bucket_name"))
		val.BucketName = v.(string)
	}

	// BucketLocation: bucket_location -- string, string
	if v, ok := d.GetOk(prefix + "bucket_location"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "bucket_location"))
		val.BucketLocation = v.(string)
	}

	// Insecure: insecure -- bool, bool
	if v, ok := d.GetOk(prefix + "insecure"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "insecure"))
		val.Insecure = v.(bool)
	}

	// AuthenticationType: authentication_type -- string, string
	if v, ok := d.GetOk(prefix + "authentication_type"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "authentication_type"))
		val.AuthenticationType = v.(string)
	}

	return val, nil
}

// config.Config

func dsRunnerConfigReadStructConfigConfig(ctx context.Context, prefix string, d *schema.ResourceData) (config.Config, error) {

	if prefix != "" {
		prefix = prefix + "."
	}

	tflog.Debug(ctx, fmt.Sprintf("beginning dsRunnerConfigReadStructConfigConfig run; prefix is '%s'", prefix))

	val := config.Config{}

	// ListenAddress: listen_address -- string, string
	if v, ok := d.GetOk(prefix + "listen_address"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "listen_address"))
		val.ListenAddress = v.(string)
	}

	// SessionServer: session_server -- SessionServer, config.SessionServer

	if _, ok := d.GetOk(prefix + "session_server.0"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "session_server"))
		thing, err := dsRunnerConfigReadStructConfigSessionServer(ctx, prefix+"session_server.0", d)
		if err != nil {
			return val, err
		}
		val.SessionServer = thing
	}

	// Concurrent: concurrent -- int, int
	if v, ok := d.GetOk(prefix + "concurrent"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "concurrent"))
		val.Concurrent = v.(int)
	}

	// CheckInterval: check_interval -- int, int
	if v, ok := d.GetOk(prefix + "check_interval"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "check_interval"))
		val.CheckInterval = v.(int)
	}

	// LogLevel: log_level -- , *string
	if v, ok := d.GetOk(prefix + "log_level"); ok {
		tflog.Debug(ctx, "set: %s.%s", prefix, "log_level")
		val.LogLevel = to.StringP(v.(string))
	}

	// LogFormat: log_format -- , *string
	if v, ok := d.GetOk(prefix + "log_format"); ok {
		tflog.Debug(ctx, "set: %s.%s", prefix, "log_format")
		val.LogFormat = to.StringP(v.(string))
	}

	// User: user -- string, string
	if v, ok := d.GetOk(prefix + "user"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "user"))
		val.User = v.(string)
	}

	// Runners: runners -- , []*config.RunnerConfig

	// HERE
	if _, ok := d.GetOk(prefix + "runners"); ok {
		tflog.Debug(ctx, fmt.Sprintf("key is set: %s%s", prefix, "runners"))
		i := 0
		val.Runners = []*config.RunnerConfig{}
		for {
			pfx := fmt.Sprintf("%s%s.%d", prefix, "runners", i)
			if _, ok := d.GetOk(pfx); ok {
				tflog.Debug(ctx, fmt.Sprintf("key is set: %s", pfx))
				thing, err := dsRunnerConfigReadStructConfigRunnerConfig(ctx, pfx, d)
				if err != nil {
					return val, err
				}
				val.Runners = append(val.Runners, &thing)
				i++
			} else {
				tflog.Debug(ctx, fmt.Sprintf("not set: %s", pfx))
				break
			}
		}
	}

	// SentryDSN: sentry_dsn -- , *string
	if v, ok := d.GetOk(prefix + "sentry_dsn"); ok {
		tflog.Debug(ctx, "set: %s.%s", prefix, "sentry_dsn")
		val.SentryDSN = to.StringP(v.(string))
	}

	return val, nil
}

// config.CustomBuildDir

func dsRunnerConfigReadStructConfigCustomBuildDir(ctx context.Context, prefix string, d *schema.ResourceData) (config.CustomBuildDir, error) {

	if prefix != "" {
		prefix = prefix + "."
	}

	tflog.Debug(ctx, fmt.Sprintf("beginning dsRunnerConfigReadStructConfigCustomBuildDir run; prefix is '%s'", prefix))

	val := config.CustomBuildDir{}

	// Enabled: enabled -- bool, bool
	if v, ok := d.GetOk(prefix + "enabled"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "enabled"))
		val.Enabled = v.(bool)
	}

	return val, nil
}

// config.CustomConfig

func dsRunnerConfigReadStructConfigCustomConfig(ctx context.Context, prefix string, d *schema.ResourceData) (config.CustomConfig, error) {

	if prefix != "" {
		prefix = prefix + "."
	}

	tflog.Debug(ctx, fmt.Sprintf("beginning dsRunnerConfigReadStructConfigCustomConfig run; prefix is '%s'", prefix))

	val := config.CustomConfig{}

	// ConfigExec: config_exec -- string, string
	if v, ok := d.GetOk(prefix + "config_exec"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "config_exec"))
		val.ConfigExec = v.(string)
	}

	// ConfigArgs: config_args -- , []string

	// HERE
	if _, ok := d.GetOk(prefix + "config_args"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "config_args"))
		i := 0
		val.ConfigArgs = []string{}
		for {
			pfx := fmt.Sprintf("%s.%s.%d", prefix, "config_args", i)
			if v, ok := d.GetOk(pfx); ok {
				tflog.Debug(ctx, fmt.Sprintf("key is set: %s", pfx))
				val.ConfigArgs = append(val.ConfigArgs, v.(string))
				i++
			} else {
				tflog.Debug(ctx, "not set: %s", pfx)
				break
			}
		}
	}

	// ConfigExecTimeout: config_exec_timeout -- , *int
	if v, ok := d.GetOk(prefix + "config_exec_timeout"); ok {
		tflog.Debug(ctx, "set: %s.%s", prefix, "config_exec_timeout")
		val.ConfigExecTimeout = to.IntP(v.(int))
	}

	// PrepareExec: prepare_exec -- string, string
	if v, ok := d.GetOk(prefix + "prepare_exec"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "prepare_exec"))
		val.PrepareExec = v.(string)
	}

	// PrepareArgs: prepare_args -- , []string

	// HERE
	if _, ok := d.GetOk(prefix + "prepare_args"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "prepare_args"))
		i := 0
		val.PrepareArgs = []string{}
		for {
			pfx := fmt.Sprintf("%s.%s.%d", prefix, "prepare_args", i)
			if v, ok := d.GetOk(pfx); ok {
				tflog.Debug(ctx, fmt.Sprintf("key is set: %s", pfx))
				val.PrepareArgs = append(val.PrepareArgs, v.(string))
				i++
			} else {
				tflog.Debug(ctx, "not set: %s", pfx)
				break
			}
		}
	}

	// PrepareExecTimeout: prepare_exec_timeout -- , *int
	if v, ok := d.GetOk(prefix + "prepare_exec_timeout"); ok {
		tflog.Debug(ctx, "set: %s.%s", prefix, "prepare_exec_timeout")
		val.PrepareExecTimeout = to.IntP(v.(int))
	}

	// RunExec: run_exec -- string, string
	if v, ok := d.GetOk(prefix + "run_exec"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "run_exec"))
		val.RunExec = v.(string)
	}

	// RunArgs: run_args -- , []string

	// HERE
	if _, ok := d.GetOk(prefix + "run_args"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "run_args"))
		i := 0
		val.RunArgs = []string{}
		for {
			pfx := fmt.Sprintf("%s.%s.%d", prefix, "run_args", i)
			if v, ok := d.GetOk(pfx); ok {
				tflog.Debug(ctx, fmt.Sprintf("key is set: %s", pfx))
				val.RunArgs = append(val.RunArgs, v.(string))
				i++
			} else {
				tflog.Debug(ctx, "not set: %s", pfx)
				break
			}
		}
	}

	// CleanupExec: cleanup_exec -- string, string
	if v, ok := d.GetOk(prefix + "cleanup_exec"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "cleanup_exec"))
		val.CleanupExec = v.(string)
	}

	// CleanupArgs: cleanup_args -- , []string

	// HERE
	if _, ok := d.GetOk(prefix + "cleanup_args"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "cleanup_args"))
		i := 0
		val.CleanupArgs = []string{}
		for {
			pfx := fmt.Sprintf("%s.%s.%d", prefix, "cleanup_args", i)
			if v, ok := d.GetOk(pfx); ok {
				tflog.Debug(ctx, fmt.Sprintf("key is set: %s", pfx))
				val.CleanupArgs = append(val.CleanupArgs, v.(string))
				i++
			} else {
				tflog.Debug(ctx, "not set: %s", pfx)
				break
			}
		}
	}

	// CleanupExecTimeout: cleanup_exec_timeout -- , *int
	if v, ok := d.GetOk(prefix + "cleanup_exec_timeout"); ok {
		tflog.Debug(ctx, "set: %s.%s", prefix, "cleanup_exec_timeout")
		val.CleanupExecTimeout = to.IntP(v.(int))
	}

	// GracefulKillTimeout: graceful_kill_timeout -- , *int
	if v, ok := d.GetOk(prefix + "graceful_kill_timeout"); ok {
		tflog.Debug(ctx, "set: %s.%s", prefix, "graceful_kill_timeout")
		val.GracefulKillTimeout = to.IntP(v.(int))
	}

	// ForceKillTimeout: force_kill_timeout -- , *int
	if v, ok := d.GetOk(prefix + "force_kill_timeout"); ok {
		tflog.Debug(ctx, "set: %s.%s", prefix, "force_kill_timeout")
		val.ForceKillTimeout = to.IntP(v.(int))
	}

	return val, nil
}

// config.DockerConfig

func dsRunnerConfigReadStructConfigDockerConfig(ctx context.Context, prefix string, d *schema.ResourceData) (config.DockerConfig, error) {

	if prefix != "" {
		prefix = prefix + "."
	}

	tflog.Debug(ctx, fmt.Sprintf("beginning dsRunnerConfigReadStructConfigDockerConfig run; prefix is '%s'", prefix))

	val := config.DockerConfig{}

	// Host: host -- string, string
	if v, ok := d.GetOk(prefix + "host"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "host"))
		val.Host = v.(string)
	}

	// CertPath: tls_cert_path -- string, string
	if v, ok := d.GetOk(prefix + "tls_cert_path"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "tls_cert_path"))
		val.CertPath = v.(string)
	}

	// TLSVerify: tls_verify -- bool, bool
	if v, ok := d.GetOk(prefix + "tls_verify"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "tls_verify"))
		val.TLSVerify = v.(bool)
	}

	// Hostname: hostname -- string, string
	if v, ok := d.GetOk(prefix + "hostname"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "hostname"))
		val.Hostname = v.(string)
	}

	// Image: image -- string, string
	if v, ok := d.GetOk(prefix + "image"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "image"))
		val.Image = v.(string)
	}

	// Runtime: runtime -- string, string
	if v, ok := d.GetOk(prefix + "runtime"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "runtime"))
		val.Runtime = v.(string)
	}

	// Memory: memory -- string, string
	if v, ok := d.GetOk(prefix + "memory"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "memory"))
		val.Memory = v.(string)
	}

	// MemorySwap: memory_swap -- string, string
	if v, ok := d.GetOk(prefix + "memory_swap"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "memory_swap"))
		val.MemorySwap = v.(string)
	}

	// MemoryReservation: memory_reservation -- string, string
	if v, ok := d.GetOk(prefix + "memory_reservation"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "memory_reservation"))
		val.MemoryReservation = v.(string)
	}

	// CPUSetCPUs: cpuset_cpus -- string, string
	if v, ok := d.GetOk(prefix + "cpuset_cpus"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "cpuset_cpus"))
		val.CPUSetCPUs = v.(string)
	}

	// CPUS: cpus -- string, string
	if v, ok := d.GetOk(prefix + "cpus"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "cpus"))
		val.CPUS = v.(string)
	}

	// CPUShares: cpu_shares -- int64, int64
	if v, ok := d.GetOk(prefix + "cpu_shares"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "cpu_shares"))
		val.CPUShares = v.(int64)
	}

	// DNS: dns -- , []string

	// HERE
	if _, ok := d.GetOk(prefix + "dns"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "dns"))
		i := 0
		val.DNS = []string{}
		for {
			pfx := fmt.Sprintf("%s.%s.%d", prefix, "dns", i)
			if v, ok := d.GetOk(pfx); ok {
				tflog.Debug(ctx, fmt.Sprintf("key is set: %s", pfx))
				val.DNS = append(val.DNS, v.(string))
				i++
			} else {
				tflog.Debug(ctx, "not set: %s", pfx)
				break
			}
		}
	}

	// DNSSearch: dns_search -- , []string

	// HERE
	if _, ok := d.GetOk(prefix + "dns_search"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "dns_search"))
		i := 0
		val.DNSSearch = []string{}
		for {
			pfx := fmt.Sprintf("%s.%s.%d", prefix, "dns_search", i)
			if v, ok := d.GetOk(pfx); ok {
				tflog.Debug(ctx, fmt.Sprintf("key is set: %s", pfx))
				val.DNSSearch = append(val.DNSSearch, v.(string))
				i++
			} else {
				tflog.Debug(ctx, "not set: %s", pfx)
				break
			}
		}
	}

	// Privileged: privileged -- bool, bool
	if v, ok := d.GetOk(prefix + "privileged"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "privileged"))
		val.Privileged = v.(bool)
	}

	// DisableEntrypointOverwrite: disable_entrypoint_overwrite -- bool, bool
	if v, ok := d.GetOk(prefix + "disable_entrypoint_overwrite"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "disable_entrypoint_overwrite"))
		val.DisableEntrypointOverwrite = v.(bool)
	}

	// UsernsMode: userns_mode -- string, string
	if v, ok := d.GetOk(prefix + "userns_mode"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "userns_mode"))
		val.UsernsMode = v.(string)
	}

	// CapAdd: cap_add -- , []string

	// HERE
	if _, ok := d.GetOk(prefix + "cap_add"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "cap_add"))
		i := 0
		val.CapAdd = []string{}
		for {
			pfx := fmt.Sprintf("%s.%s.%d", prefix, "cap_add", i)
			if v, ok := d.GetOk(pfx); ok {
				tflog.Debug(ctx, fmt.Sprintf("key is set: %s", pfx))
				val.CapAdd = append(val.CapAdd, v.(string))
				i++
			} else {
				tflog.Debug(ctx, "not set: %s", pfx)
				break
			}
		}
	}

	// CapDrop: cap_drop -- , []string

	// HERE
	if _, ok := d.GetOk(prefix + "cap_drop"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "cap_drop"))
		i := 0
		val.CapDrop = []string{}
		for {
			pfx := fmt.Sprintf("%s.%s.%d", prefix, "cap_drop", i)
			if v, ok := d.GetOk(pfx); ok {
				tflog.Debug(ctx, fmt.Sprintf("key is set: %s", pfx))
				val.CapDrop = append(val.CapDrop, v.(string))
				i++
			} else {
				tflog.Debug(ctx, "not set: %s", pfx)
				break
			}
		}
	}

	// OomKillDisable: oom_kill_disable -- bool, bool
	if v, ok := d.GetOk(prefix + "oom_kill_disable"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "oom_kill_disable"))
		val.OomKillDisable = v.(bool)
	}

	// OomScoreAdjust: oom_score_adjust -- int, int
	if v, ok := d.GetOk(prefix + "oom_score_adjust"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "oom_score_adjust"))
		val.OomScoreAdjust = v.(int)
	}

	// SecurityOpt: security_opt -- , []string

	// HERE
	if _, ok := d.GetOk(prefix + "security_opt"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "security_opt"))
		i := 0
		val.SecurityOpt = []string{}
		for {
			pfx := fmt.Sprintf("%s.%s.%d", prefix, "security_opt", i)
			if v, ok := d.GetOk(pfx); ok {
				tflog.Debug(ctx, fmt.Sprintf("key is set: %s", pfx))
				val.SecurityOpt = append(val.SecurityOpt, v.(string))
				i++
			} else {
				tflog.Debug(ctx, "not set: %s", pfx)
				break
			}
		}
	}

	// Devices: devices -- , []string

	// HERE
	if _, ok := d.GetOk(prefix + "devices"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "devices"))
		i := 0
		val.Devices = []string{}
		for {
			pfx := fmt.Sprintf("%s.%s.%d", prefix, "devices", i)
			if v, ok := d.GetOk(pfx); ok {
				tflog.Debug(ctx, fmt.Sprintf("key is set: %s", pfx))
				val.Devices = append(val.Devices, v.(string))
				i++
			} else {
				tflog.Debug(ctx, "not set: %s", pfx)
				break
			}
		}
	}

	// Gpus: gpus -- string, string
	if v, ok := d.GetOk(prefix + "gpus"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "gpus"))
		val.Gpus = v.(string)
	}

	// DisableCache: disable_cache -- bool, bool
	if v, ok := d.GetOk(prefix + "disable_cache"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "disable_cache"))
		val.DisableCache = v.(bool)
	}

	// Volumes: volumes -- , []string

	// HERE
	if _, ok := d.GetOk(prefix + "volumes"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "volumes"))
		i := 0
		val.Volumes = []string{}
		for {
			pfx := fmt.Sprintf("%s.%s.%d", prefix, "volumes", i)
			if v, ok := d.GetOk(pfx); ok {
				tflog.Debug(ctx, fmt.Sprintf("key is set: %s", pfx))
				val.Volumes = append(val.Volumes, v.(string))
				i++
			} else {
				tflog.Debug(ctx, "not set: %s", pfx)
				break
			}
		}
	}

	// VolumeDriver: volume_driver -- string, string
	if v, ok := d.GetOk(prefix + "volume_driver"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "volume_driver"))
		val.VolumeDriver = v.(string)
	}

	// CacheDir: cache_dir -- string, string
	if v, ok := d.GetOk(prefix + "cache_dir"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "cache_dir"))
		val.CacheDir = v.(string)
	}

	// ExtraHosts: extra_hosts -- , []string

	// HERE
	if _, ok := d.GetOk(prefix + "extra_hosts"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "extra_hosts"))
		i := 0
		val.ExtraHosts = []string{}
		for {
			pfx := fmt.Sprintf("%s.%s.%d", prefix, "extra_hosts", i)
			if v, ok := d.GetOk(pfx); ok {
				tflog.Debug(ctx, fmt.Sprintf("key is set: %s", pfx))
				val.ExtraHosts = append(val.ExtraHosts, v.(string))
				i++
			} else {
				tflog.Debug(ctx, "not set: %s", pfx)
				break
			}
		}
	}

	// VolumesFrom: volumes_from -- , []string

	// HERE
	if _, ok := d.GetOk(prefix + "volumes_from"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "volumes_from"))
		i := 0
		val.VolumesFrom = []string{}
		for {
			pfx := fmt.Sprintf("%s.%s.%d", prefix, "volumes_from", i)
			if v, ok := d.GetOk(pfx); ok {
				tflog.Debug(ctx, fmt.Sprintf("key is set: %s", pfx))
				val.VolumesFrom = append(val.VolumesFrom, v.(string))
				i++
			} else {
				tflog.Debug(ctx, "not set: %s", pfx)
				break
			}
		}
	}

	// NetworkMode: network_mode -- string, string
	if v, ok := d.GetOk(prefix + "network_mode"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "network_mode"))
		val.NetworkMode = v.(string)
	}

	// Links: links -- , []string

	// HERE
	if _, ok := d.GetOk(prefix + "links"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "links"))
		i := 0
		val.Links = []string{}
		for {
			pfx := fmt.Sprintf("%s.%s.%d", prefix, "links", i)
			if v, ok := d.GetOk(pfx); ok {
				tflog.Debug(ctx, fmt.Sprintf("key is set: %s", pfx))
				val.Links = append(val.Links, v.(string))
				i++
			} else {
				tflog.Debug(ctx, "not set: %s", pfx)
				break
			}
		}
	}

	// Services: services -- , []config.Service

	// HERE
	if _, ok := d.GetOk(prefix + "services"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "services"))
		i := 0
		val.Services = []config.Service{}
		for {
			pfx := fmt.Sprintf("%s.%s.%d", prefix, "services", i)
			if _, ok := d.GetOk(pfx); ok {
				tflog.Debug(ctx, "set: %s", pfx)
				thing, err := dsRunnerConfigReadStructConfigService(ctx, pfx, d)
				if err != nil {
					return val, err
				}
				val.Services = append(val.Services, thing)
				i++
			} else {
				tflog.Debug(ctx, "not set: %s", pfx)
				break
			}
		}
	}

	// WaitForServicesTimeout: wait_for_services_timeout -- int, int
	if v, ok := d.GetOk(prefix + "wait_for_services_timeout"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "wait_for_services_timeout"))
		val.WaitForServicesTimeout = v.(int)
	}

	// AllowedImages: allowed_images -- , []string

	// HERE
	if _, ok := d.GetOk(prefix + "allowed_images"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "allowed_images"))
		i := 0
		val.AllowedImages = []string{}
		for {
			pfx := fmt.Sprintf("%s.%s.%d", prefix, "allowed_images", i)
			if v, ok := d.GetOk(pfx); ok {
				tflog.Debug(ctx, fmt.Sprintf("key is set: %s", pfx))
				val.AllowedImages = append(val.AllowedImages, v.(string))
				i++
			} else {
				tflog.Debug(ctx, "not set: %s", pfx)
				break
			}
		}
	}

	// AllowedServices: allowed_services -- , []string

	// HERE
	if _, ok := d.GetOk(prefix + "allowed_services"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "allowed_services"))
		i := 0
		val.AllowedServices = []string{}
		for {
			pfx := fmt.Sprintf("%s.%s.%d", prefix, "allowed_services", i)
			if v, ok := d.GetOk(pfx); ok {
				tflog.Debug(ctx, fmt.Sprintf("key is set: %s", pfx))
				val.AllowedServices = append(val.AllowedServices, v.(string))
				i++
			} else {
				tflog.Debug(ctx, "not set: %s", pfx)
				break
			}
		}
	}

	// PullPolicy: pull_policy -- StringOrArray, config.StringOrArray
	if v, ok := d.GetOk(prefix + "pull_policy"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "pull_policy"))
		val.PullPolicy = v.([]string)
	}

	// ShmSize: shm_size -- int64, int64
	if v, ok := d.GetOk(prefix + "shm_size"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "shm_size"))
		val.ShmSize = v.(int64)
	}

	// Tmpfs: tmpfs -- , map[string]string
	if v, ok := d.GetOk(prefix + "tmpfs"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "tmpfs"))
		val.Tmpfs = v.(map[string]string)
	}

	// ServicesTmpfs: services_tmpfs -- , map[string]string
	if v, ok := d.GetOk(prefix + "services_tmpfs"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "services_tmpfs"))
		val.ServicesTmpfs = v.(map[string]string)
	}

	// SysCtls: sysctls -- , map[string]string
	if v, ok := d.GetOk(prefix + "sysctls"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "sysctls"))
		val.SysCtls = v.(map[string]string)
	}

	// HelperImage: helper_image -- string, string
	if v, ok := d.GetOk(prefix + "helper_image"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "helper_image"))
		val.HelperImage = v.(string)
	}

	// HelperImageFlavor: helper_image_flavor -- string, string
	if v, ok := d.GetOk(prefix + "helper_image_flavor"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "helper_image_flavor"))
		val.HelperImageFlavor = v.(string)
	}

	return val, nil
}

// config.DockerMachine

func dsRunnerConfigReadStructConfigDockerMachine(ctx context.Context, prefix string, d *schema.ResourceData) (config.DockerMachine, error) {

	if prefix != "" {
		prefix = prefix + "."
	}

	tflog.Debug(ctx, fmt.Sprintf("beginning dsRunnerConfigReadStructConfigDockerMachine run; prefix is '%s'", prefix))

	val := config.DockerMachine{}

	// MaxGrowthRate: max_growth_rate -- int, int
	if v, ok := d.GetOk(prefix + "max_growth_rate"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "max_growth_rate"))
		val.MaxGrowthRate = v.(int)
	}

	// IdleCount: idle_count -- int, int
	if v, ok := d.GetOk(prefix + "idle_count"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "idle_count"))
		val.IdleCount = v.(int)
	}

	// IdleScaleFactor: idle_scale_factor -- float64, float64
	if v, ok := d.GetOk(prefix + "idle_scale_factor"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "idle_scale_factor"))
		val.IdleScaleFactor = v.(float64)
	}

	// IdleCountMin: idle_count_min -- int, int
	if v, ok := d.GetOk(prefix + "idle_count_min"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "idle_count_min"))
		val.IdleCountMin = v.(int)
	}

	// IdleTime: idle_time -- int, int
	if v, ok := d.GetOk(prefix + "idle_time"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "idle_time"))
		val.IdleTime = v.(int)
	}

	// MaxBuilds: max_builds -- int, int
	if v, ok := d.GetOk(prefix + "max_builds"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "max_builds"))
		val.MaxBuilds = v.(int)
	}

	// MachineDriver: machine_driver -- string, string
	if v, ok := d.GetOk(prefix + "machine_driver"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "machine_driver"))
		val.MachineDriver = v.(string)
	}

	// MachineName: machine_name -- string, string
	if v, ok := d.GetOk(prefix + "machine_name"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "machine_name"))
		val.MachineName = v.(string)
	}

	// MachineOptions: machine_options -- , []string

	// HERE
	if _, ok := d.GetOk(prefix + "machine_options"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "machine_options"))
		i := 0
		val.MachineOptions = []string{}
		for {
			pfx := fmt.Sprintf("%s.%s.%d", prefix, "machine_options", i)
			if v, ok := d.GetOk(pfx); ok {
				tflog.Debug(ctx, fmt.Sprintf("key is set: %s", pfx))
				val.MachineOptions = append(val.MachineOptions, v.(string))
				i++
			} else {
				tflog.Debug(ctx, "not set: %s", pfx)
				break
			}
		}
	}

	// OffPeakPeriods: off_peak_periods -- , []string

	// HERE
	if _, ok := d.GetOk(prefix + "off_peak_periods"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "off_peak_periods"))
		i := 0
		val.OffPeakPeriods = []string{}
		for {
			pfx := fmt.Sprintf("%s.%s.%d", prefix, "off_peak_periods", i)
			if v, ok := d.GetOk(pfx); ok {
				tflog.Debug(ctx, fmt.Sprintf("key is set: %s", pfx))
				val.OffPeakPeriods = append(val.OffPeakPeriods, v.(string))
				i++
			} else {
				tflog.Debug(ctx, "not set: %s", pfx)
				break
			}
		}
	}

	// OffPeakTimezone: off_peak_timezone -- string, string
	if v, ok := d.GetOk(prefix + "off_peak_timezone"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "off_peak_timezone"))
		val.OffPeakTimezone = v.(string)
	}

	// OffPeakIdleCount: off_peak_idle_count -- int, int
	if v, ok := d.GetOk(prefix + "off_peak_idle_count"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "off_peak_idle_count"))
		val.OffPeakIdleCount = v.(int)
	}

	// OffPeakIdleTime: off_peak_idle_time -- int, int
	if v, ok := d.GetOk(prefix + "off_peak_idle_time"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "off_peak_idle_time"))
		val.OffPeakIdleTime = v.(int)
	}

	// AutoscalingConfigs: autoscaling -- , []*config.DockerMachineAutoscaling

	// HERE
	if _, ok := d.GetOk(prefix + "autoscaling"); ok {
		tflog.Debug(ctx, fmt.Sprintf("key is set: %s%s", prefix, "autoscaling"))
		i := 0
		val.AutoscalingConfigs = []*config.DockerMachineAutoscaling{}
		for {
			pfx := fmt.Sprintf("%s%s.%d", prefix, "autoscaling", i)
			if _, ok := d.GetOk(pfx); ok {
				tflog.Debug(ctx, fmt.Sprintf("key is set: %s", pfx))
				thing, err := dsRunnerConfigReadStructConfigDockerMachineAutoscaling(ctx, pfx, d)
				if err != nil {
					return val, err
				}
				val.AutoscalingConfigs = append(val.AutoscalingConfigs, &thing)
				i++
			} else {
				tflog.Debug(ctx, fmt.Sprintf("not set: %s", pfx))
				break
			}
		}
	}

	return val, nil
}

// config.DockerMachineAutoscaling

func dsRunnerConfigReadStructConfigDockerMachineAutoscaling(ctx context.Context, prefix string, d *schema.ResourceData) (config.DockerMachineAutoscaling, error) {

	if prefix != "" {
		prefix = prefix + "."
	}

	tflog.Debug(ctx, fmt.Sprintf("beginning dsRunnerConfigReadStructConfigDockerMachineAutoscaling run; prefix is '%s'", prefix))

	val := config.DockerMachineAutoscaling{}

	// Periods: periods -- , []string

	// HERE
	if _, ok := d.GetOk(prefix + "periods"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "periods"))
		i := 0
		val.Periods = []string{}
		for {
			pfx := fmt.Sprintf("%s.%s.%d", prefix, "periods", i)
			if v, ok := d.GetOk(pfx); ok {
				tflog.Debug(ctx, fmt.Sprintf("key is set: %s", pfx))
				val.Periods = append(val.Periods, v.(string))
				i++
			} else {
				tflog.Debug(ctx, "not set: %s", pfx)
				break
			}
		}
	}

	// Timezone: timezone -- string, string
	if v, ok := d.GetOk(prefix + "timezone"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "timezone"))
		val.Timezone = v.(string)
	}

	// IdleCount: idle_count -- int, int
	if v, ok := d.GetOk(prefix + "idle_count"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "idle_count"))
		val.IdleCount = v.(int)
	}

	// IdleScaleFactor: idle_scale_factor -- float64, float64
	if v, ok := d.GetOk(prefix + "idle_scale_factor"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "idle_scale_factor"))
		val.IdleScaleFactor = v.(float64)
	}

	// IdleCountMin: idle_count_min -- int, int
	if v, ok := d.GetOk(prefix + "idle_count_min"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "idle_count_min"))
		val.IdleCountMin = v.(int)
	}

	// IdleTime: idle_time -- int, int
	if v, ok := d.GetOk(prefix + "idle_time"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "idle_time"))
		val.IdleTime = v.(int)
	}

	return val, nil
}

// config.KubernetesAffinity

func dsRunnerConfigReadStructConfigKubernetesAffinity(ctx context.Context, prefix string, d *schema.ResourceData) (config.KubernetesAffinity, error) {

	if prefix != "" {
		prefix = prefix + "."
	}

	tflog.Debug(ctx, fmt.Sprintf("beginning dsRunnerConfigReadStructConfigKubernetesAffinity run; prefix is '%s'", prefix))

	val := config.KubernetesAffinity{}

	// NodeAffinity: node_affinity -- , *config.KubernetesNodeAffinity

	if _, ok := d.GetOk(prefix + "node_affinity.0"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "node_affinity"))
		thing, err := dsRunnerConfigReadStructConfigKubernetesNodeAffinity(ctx, prefix+"node_affinity.0", d)
		if err != nil {
			return val, err
		}
		val.NodeAffinity = &thing
	} else {
		tflog.Debug(ctx, fmt.Sprintf("not set: %s", prefix+"node_affinity"))
	}

	// PodAffinity: pod_affinity -- , *config.KubernetesPodAffinity

	if _, ok := d.GetOk(prefix + "pod_affinity.0"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "pod_affinity"))
		thing, err := dsRunnerConfigReadStructConfigKubernetesPodAffinity(ctx, prefix+"pod_affinity.0", d)
		if err != nil {
			return val, err
		}
		val.PodAffinity = &thing
	} else {
		tflog.Debug(ctx, fmt.Sprintf("not set: %s", prefix+"pod_affinity"))
	}

	// PodAntiAffinity: pod_anti_affinity -- , *config.KubernetesPodAntiAffinity

	if _, ok := d.GetOk(prefix + "pod_anti_affinity.0"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "pod_anti_affinity"))
		thing, err := dsRunnerConfigReadStructConfigKubernetesPodAntiAffinity(ctx, prefix+"pod_anti_affinity.0", d)
		if err != nil {
			return val, err
		}
		val.PodAntiAffinity = &thing
	} else {
		tflog.Debug(ctx, fmt.Sprintf("not set: %s", prefix+"pod_anti_affinity"))
	}

	return val, nil
}

// config.KubernetesCSI

func dsRunnerConfigReadStructConfigKubernetesCSI(ctx context.Context, prefix string, d *schema.ResourceData) (config.KubernetesCSI, error) {

	if prefix != "" {
		prefix = prefix + "."
	}

	tflog.Debug(ctx, fmt.Sprintf("beginning dsRunnerConfigReadStructConfigKubernetesCSI run; prefix is '%s'", prefix))

	val := config.KubernetesCSI{}

	// Name: name -- string, string
	if v, ok := d.GetOk(prefix + "name"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "name"))
		val.Name = v.(string)
	}

	// MountPath: mount_path -- string, string
	if v, ok := d.GetOk(prefix + "mount_path"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "mount_path"))
		val.MountPath = v.(string)
	}

	// SubPath: sub_path -- string, string
	if v, ok := d.GetOk(prefix + "sub_path"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "sub_path"))
		val.SubPath = v.(string)
	}

	// Driver: driver -- string, string
	if v, ok := d.GetOk(prefix + "driver"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "driver"))
		val.Driver = v.(string)
	}

	// FSType: fs_type -- string, string
	if v, ok := d.GetOk(prefix + "fs_type"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "fs_type"))
		val.FSType = v.(string)
	}

	// ReadOnly: read_only -- bool, bool
	if v, ok := d.GetOk(prefix + "read_only"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "read_only"))
		val.ReadOnly = v.(bool)
	}

	// VolumeAttributes: volume_attributes -- , map[string]string
	if v, ok := d.GetOk(prefix + "volume_attributes"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "volume_attributes"))
		val.VolumeAttributes = v.(map[string]string)
	}

	return val, nil
}

// config.KubernetesConfig

func dsRunnerConfigReadStructConfigKubernetesConfig(ctx context.Context, prefix string, d *schema.ResourceData) (config.KubernetesConfig, error) {

	if prefix != "" {
		prefix = prefix + "."
	}

	tflog.Debug(ctx, fmt.Sprintf("beginning dsRunnerConfigReadStructConfigKubernetesConfig run; prefix is '%s'", prefix))

	val := config.KubernetesConfig{}

	// Host: host -- string, string
	if v, ok := d.GetOk(prefix + "host"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "host"))
		val.Host = v.(string)
	}

	// CertFile: cert_file -- string, string
	if v, ok := d.GetOk(prefix + "cert_file"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "cert_file"))
		val.CertFile = v.(string)
	}

	// KeyFile: key_file -- string, string
	if v, ok := d.GetOk(prefix + "key_file"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "key_file"))
		val.KeyFile = v.(string)
	}

	// CAFile: ca_file -- string, string
	if v, ok := d.GetOk(prefix + "ca_file"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "ca_file"))
		val.CAFile = v.(string)
	}

	// BearerTokenOverwriteAllowed: bearer_token_overwrite_allowed -- bool, bool
	if v, ok := d.GetOk(prefix + "bearer_token_overwrite_allowed"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "bearer_token_overwrite_allowed"))
		val.BearerTokenOverwriteAllowed = v.(bool)
	}

	// BearerToken: bearer_token -- string, string
	if v, ok := d.GetOk(prefix + "bearer_token"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "bearer_token"))
		val.BearerToken = v.(string)
	}

	// Image: image -- string, string
	if v, ok := d.GetOk(prefix + "image"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "image"))
		val.Image = v.(string)
	}

	// Namespace: namespace -- string, string
	if v, ok := d.GetOk(prefix + "namespace"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "namespace"))
		val.Namespace = v.(string)
	}

	// NamespaceOverwriteAllowed: namespace_overwrite_allowed -- string, string
	if v, ok := d.GetOk(prefix + "namespace_overwrite_allowed"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "namespace_overwrite_allowed"))
		val.NamespaceOverwriteAllowed = v.(string)
	}

	// Privileged: privileged -- , *bool
	if v, ok := d.GetOk(prefix + "privileged"); ok {
		tflog.Debug(ctx, "set: %s.%s", prefix, "privileged")
		val.Privileged = to.BoolP(v.(bool))
	}

	// AllowPrivilegeEscalation: allow_privilege_escalation -- , *bool
	if v, ok := d.GetOk(prefix + "allow_privilege_escalation"); ok {
		tflog.Debug(ctx, "set: %s.%s", prefix, "allow_privilege_escalation")
		val.AllowPrivilegeEscalation = to.BoolP(v.(bool))
	}

	// CPULimit: cpu_limit -- string, string
	if v, ok := d.GetOk(prefix + "cpu_limit"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "cpu_limit"))
		val.CPULimit = v.(string)
	}

	// CPULimitOverwriteMaxAllowed: cpu_limit_overwrite_max_allowed -- string, string
	if v, ok := d.GetOk(prefix + "cpu_limit_overwrite_max_allowed"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "cpu_limit_overwrite_max_allowed"))
		val.CPULimitOverwriteMaxAllowed = v.(string)
	}

	// CPURequest: cpu_request -- string, string
	if v, ok := d.GetOk(prefix + "cpu_request"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "cpu_request"))
		val.CPURequest = v.(string)
	}

	// CPURequestOverwriteMaxAllowed: cpu_request_overwrite_max_allowed -- string, string
	if v, ok := d.GetOk(prefix + "cpu_request_overwrite_max_allowed"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "cpu_request_overwrite_max_allowed"))
		val.CPURequestOverwriteMaxAllowed = v.(string)
	}

	// MemoryLimit: memory_limit -- string, string
	if v, ok := d.GetOk(prefix + "memory_limit"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "memory_limit"))
		val.MemoryLimit = v.(string)
	}

	// MemoryLimitOverwriteMaxAllowed: memory_limit_overwrite_max_allowed -- string, string
	if v, ok := d.GetOk(prefix + "memory_limit_overwrite_max_allowed"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "memory_limit_overwrite_max_allowed"))
		val.MemoryLimitOverwriteMaxAllowed = v.(string)
	}

	// MemoryRequest: memory_request -- string, string
	if v, ok := d.GetOk(prefix + "memory_request"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "memory_request"))
		val.MemoryRequest = v.(string)
	}

	// MemoryRequestOverwriteMaxAllowed: memory_request_overwrite_max_allowed -- string, string
	if v, ok := d.GetOk(prefix + "memory_request_overwrite_max_allowed"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "memory_request_overwrite_max_allowed"))
		val.MemoryRequestOverwriteMaxAllowed = v.(string)
	}

	// EphemeralStorageLimit: ephemeral_storage_limit -- string, string
	if v, ok := d.GetOk(prefix + "ephemeral_storage_limit"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "ephemeral_storage_limit"))
		val.EphemeralStorageLimit = v.(string)
	}

	// EphemeralStorageLimitOverwriteMaxAllowed: ephemeral_storage_limit_overwrite_max_allowed -- string, string
	if v, ok := d.GetOk(prefix + "ephemeral_storage_limit_overwrite_max_allowed"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "ephemeral_storage_limit_overwrite_max_allowed"))
		val.EphemeralStorageLimitOverwriteMaxAllowed = v.(string)
	}

	// EphemeralStorageRequest: ephemeral_storage_request -- string, string
	if v, ok := d.GetOk(prefix + "ephemeral_storage_request"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "ephemeral_storage_request"))
		val.EphemeralStorageRequest = v.(string)
	}

	// EphemeralStorageRequestOverwriteMaxAllowed: ephemeral_storage_request_overwrite_max_allowed -- string, string
	if v, ok := d.GetOk(prefix + "ephemeral_storage_request_overwrite_max_allowed"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "ephemeral_storage_request_overwrite_max_allowed"))
		val.EphemeralStorageRequestOverwriteMaxAllowed = v.(string)
	}

	// ServiceCPULimit: service_cpu_limit -- string, string
	if v, ok := d.GetOk(prefix + "service_cpu_limit"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "service_cpu_limit"))
		val.ServiceCPULimit = v.(string)
	}

	// ServiceCPULimitOverwriteMaxAllowed: service_cpu_limit_overwrite_max_allowed -- string, string
	if v, ok := d.GetOk(prefix + "service_cpu_limit_overwrite_max_allowed"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "service_cpu_limit_overwrite_max_allowed"))
		val.ServiceCPULimitOverwriteMaxAllowed = v.(string)
	}

	// ServiceCPURequest: service_cpu_request -- string, string
	if v, ok := d.GetOk(prefix + "service_cpu_request"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "service_cpu_request"))
		val.ServiceCPURequest = v.(string)
	}

	// ServiceCPURequestOverwriteMaxAllowed: service_cpu_request_overwrite_max_allowed -- string, string
	if v, ok := d.GetOk(prefix + "service_cpu_request_overwrite_max_allowed"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "service_cpu_request_overwrite_max_allowed"))
		val.ServiceCPURequestOverwriteMaxAllowed = v.(string)
	}

	// ServiceMemoryLimit: service_memory_limit -- string, string
	if v, ok := d.GetOk(prefix + "service_memory_limit"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "service_memory_limit"))
		val.ServiceMemoryLimit = v.(string)
	}

	// ServiceMemoryLimitOverwriteMaxAllowed: service_memory_limit_overwrite_max_allowed -- string, string
	if v, ok := d.GetOk(prefix + "service_memory_limit_overwrite_max_allowed"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "service_memory_limit_overwrite_max_allowed"))
		val.ServiceMemoryLimitOverwriteMaxAllowed = v.(string)
	}

	// ServiceMemoryRequest: service_memory_request -- string, string
	if v, ok := d.GetOk(prefix + "service_memory_request"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "service_memory_request"))
		val.ServiceMemoryRequest = v.(string)
	}

	// ServiceMemoryRequestOverwriteMaxAllowed: service_memory_request_overwrite_max_allowed -- string, string
	if v, ok := d.GetOk(prefix + "service_memory_request_overwrite_max_allowed"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "service_memory_request_overwrite_max_allowed"))
		val.ServiceMemoryRequestOverwriteMaxAllowed = v.(string)
	}

	// ServiceEphemeralStorageLimit: service_ephemeral_storage_limit -- string, string
	if v, ok := d.GetOk(prefix + "service_ephemeral_storage_limit"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "service_ephemeral_storage_limit"))
		val.ServiceEphemeralStorageLimit = v.(string)
	}

	// ServiceEphemeralStorageLimitOverwriteMaxAllowed: service_ephemeral_storage_limit_overwrite_max_allowed -- string, string
	if v, ok := d.GetOk(prefix + "service_ephemeral_storage_limit_overwrite_max_allowed"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "service_ephemeral_storage_limit_overwrite_max_allowed"))
		val.ServiceEphemeralStorageLimitOverwriteMaxAllowed = v.(string)
	}

	// ServiceEphemeralStorageRequest: service_ephemeral_storage_request -- string, string
	if v, ok := d.GetOk(prefix + "service_ephemeral_storage_request"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "service_ephemeral_storage_request"))
		val.ServiceEphemeralStorageRequest = v.(string)
	}

	// ServiceEphemeralStorageRequestOverwriteMaxAllowed: service_ephemeral_storage_request_overwrite_max_allowed -- string, string
	if v, ok := d.GetOk(prefix + "service_ephemeral_storage_request_overwrite_max_allowed"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "service_ephemeral_storage_request_overwrite_max_allowed"))
		val.ServiceEphemeralStorageRequestOverwriteMaxAllowed = v.(string)
	}

	// HelperCPULimit: helper_cpu_limit -- string, string
	if v, ok := d.GetOk(prefix + "helper_cpu_limit"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "helper_cpu_limit"))
		val.HelperCPULimit = v.(string)
	}

	// HelperCPULimitOverwriteMaxAllowed: helper_cpu_limit_overwrite_max_allowed -- string, string
	if v, ok := d.GetOk(prefix + "helper_cpu_limit_overwrite_max_allowed"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "helper_cpu_limit_overwrite_max_allowed"))
		val.HelperCPULimitOverwriteMaxAllowed = v.(string)
	}

	// HelperCPURequest: helper_cpu_request -- string, string
	if v, ok := d.GetOk(prefix + "helper_cpu_request"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "helper_cpu_request"))
		val.HelperCPURequest = v.(string)
	}

	// HelperCPURequestOverwriteMaxAllowed: helper_cpu_request_overwrite_max_allowed -- string, string
	if v, ok := d.GetOk(prefix + "helper_cpu_request_overwrite_max_allowed"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "helper_cpu_request_overwrite_max_allowed"))
		val.HelperCPURequestOverwriteMaxAllowed = v.(string)
	}

	// HelperMemoryLimit: helper_memory_limit -- string, string
	if v, ok := d.GetOk(prefix + "helper_memory_limit"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "helper_memory_limit"))
		val.HelperMemoryLimit = v.(string)
	}

	// HelperMemoryLimitOverwriteMaxAllowed: helper_memory_limit_overwrite_max_allowed -- string, string
	if v, ok := d.GetOk(prefix + "helper_memory_limit_overwrite_max_allowed"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "helper_memory_limit_overwrite_max_allowed"))
		val.HelperMemoryLimitOverwriteMaxAllowed = v.(string)
	}

	// HelperMemoryRequest: helper_memory_request -- string, string
	if v, ok := d.GetOk(prefix + "helper_memory_request"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "helper_memory_request"))
		val.HelperMemoryRequest = v.(string)
	}

	// HelperMemoryRequestOverwriteMaxAllowed: helper_memory_request_overwrite_max_allowed -- string, string
	if v, ok := d.GetOk(prefix + "helper_memory_request_overwrite_max_allowed"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "helper_memory_request_overwrite_max_allowed"))
		val.HelperMemoryRequestOverwriteMaxAllowed = v.(string)
	}

	// HelperEphemeralStorageLimit: helper_ephemeral_storage_limit -- string, string
	if v, ok := d.GetOk(prefix + "helper_ephemeral_storage_limit"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "helper_ephemeral_storage_limit"))
		val.HelperEphemeralStorageLimit = v.(string)
	}

	// HelperEphemeralStorageLimitOverwriteMaxAllowed: helper_ephemeral_storage_limit_overwrite_max_allowed -- string, string
	if v, ok := d.GetOk(prefix + "helper_ephemeral_storage_limit_overwrite_max_allowed"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "helper_ephemeral_storage_limit_overwrite_max_allowed"))
		val.HelperEphemeralStorageLimitOverwriteMaxAllowed = v.(string)
	}

	// HelperEphemeralStorageRequest: helper_ephemeral_storage_request -- string, string
	if v, ok := d.GetOk(prefix + "helper_ephemeral_storage_request"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "helper_ephemeral_storage_request"))
		val.HelperEphemeralStorageRequest = v.(string)
	}

	// HelperEphemeralStorageRequestOverwriteMaxAllowed: helper_ephemeral_storage_request_overwrite_max_allowed -- string, string
	if v, ok := d.GetOk(prefix + "helper_ephemeral_storage_request_overwrite_max_allowed"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "helper_ephemeral_storage_request_overwrite_max_allowed"))
		val.HelperEphemeralStorageRequestOverwriteMaxAllowed = v.(string)
	}

	// AllowedImages: allowed_images -- , []string

	// HERE
	if _, ok := d.GetOk(prefix + "allowed_images"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "allowed_images"))
		i := 0
		val.AllowedImages = []string{}
		for {
			pfx := fmt.Sprintf("%s.%s.%d", prefix, "allowed_images", i)
			if v, ok := d.GetOk(pfx); ok {
				tflog.Debug(ctx, fmt.Sprintf("key is set: %s", pfx))
				val.AllowedImages = append(val.AllowedImages, v.(string))
				i++
			} else {
				tflog.Debug(ctx, "not set: %s", pfx)
				break
			}
		}
	}

	// AllowedServices: allowed_services -- , []string

	// HERE
	if _, ok := d.GetOk(prefix + "allowed_services"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "allowed_services"))
		i := 0
		val.AllowedServices = []string{}
		for {
			pfx := fmt.Sprintf("%s.%s.%d", prefix, "allowed_services", i)
			if v, ok := d.GetOk(pfx); ok {
				tflog.Debug(ctx, fmt.Sprintf("key is set: %s", pfx))
				val.AllowedServices = append(val.AllowedServices, v.(string))
				i++
			} else {
				tflog.Debug(ctx, "not set: %s", pfx)
				break
			}
		}
	}

	// PullPolicy: pull_policy -- StringOrArray, config.StringOrArray
	if v, ok := d.GetOk(prefix + "pull_policy"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "pull_policy"))
		val.PullPolicy = v.([]string)
	}

	// NodeSelector: node_selector -- , map[string]string
	if v, ok := d.GetOk(prefix + "node_selector"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "node_selector"))
		val.NodeSelector = v.(map[string]string)
	}

	// NodeTolerations: node_tolerations -- , map[string]string
	if v, ok := d.GetOk(prefix + "node_tolerations"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "node_tolerations"))
		val.NodeTolerations = v.(map[string]string)
	}

	// Affinity: affinity -- KubernetesAffinity, config.KubernetesAffinity

	if _, ok := d.GetOk(prefix + "affinity.0"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "affinity"))
		thing, err := dsRunnerConfigReadStructConfigKubernetesAffinity(ctx, prefix+"affinity.0", d)
		if err != nil {
			return val, err
		}
		val.Affinity = thing
	}

	// ImagePullSecrets: image_pull_secrets -- , []string

	// HERE
	if _, ok := d.GetOk(prefix + "image_pull_secrets"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "image_pull_secrets"))
		i := 0
		val.ImagePullSecrets = []string{}
		for {
			pfx := fmt.Sprintf("%s.%s.%d", prefix, "image_pull_secrets", i)
			if v, ok := d.GetOk(pfx); ok {
				tflog.Debug(ctx, fmt.Sprintf("key is set: %s", pfx))
				val.ImagePullSecrets = append(val.ImagePullSecrets, v.(string))
				i++
			} else {
				tflog.Debug(ctx, "not set: %s", pfx)
				break
			}
		}
	}

	// HelperImage: helper_image -- string, string
	if v, ok := d.GetOk(prefix + "helper_image"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "helper_image"))
		val.HelperImage = v.(string)
	}

	// HelperImageFlavor: helper_image_flavor -- string, string
	if v, ok := d.GetOk(prefix + "helper_image_flavor"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "helper_image_flavor"))
		val.HelperImageFlavor = v.(string)
	}

	// TerminationGracePeriodSeconds: termination_grace_period_seconds -- , *int64
	if v, ok := d.GetOk(prefix + "termination_grace_period_seconds"); ok {
		tflog.Debug(ctx, "set: %s.%s", prefix, "termination_grace_period_seconds")
		val.TerminationGracePeriodSeconds = to.Int64P(v.(int64))
	}

	// PodTerminationGracePeriodSeconds: pod_termination_grace_period_seconds -- , *int64
	if v, ok := d.GetOk(prefix + "pod_termination_grace_period_seconds"); ok {
		tflog.Debug(ctx, "set: %s.%s", prefix, "pod_termination_grace_period_seconds")
		val.PodTerminationGracePeriodSeconds = to.Int64P(v.(int64))
	}

	// CleanupGracePeriodSeconds: cleanup_grace_period_seconds -- , *int64
	if v, ok := d.GetOk(prefix + "cleanup_grace_period_seconds"); ok {
		tflog.Debug(ctx, "set: %s.%s", prefix, "cleanup_grace_period_seconds")
		val.CleanupGracePeriodSeconds = to.Int64P(v.(int64))
	}

	// PollInterval: poll_interval -- int, int
	if v, ok := d.GetOk(prefix + "poll_interval"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "poll_interval"))
		val.PollInterval = v.(int)
	}

	// PollTimeout: poll_timeout -- int, int
	if v, ok := d.GetOk(prefix + "poll_timeout"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "poll_timeout"))
		val.PollTimeout = v.(int)
	}

	// PodLabels: pod_labels -- , map[string]string
	if v, ok := d.GetOk(prefix + "pod_labels"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "pod_labels"))
		val.PodLabels = v.(map[string]string)
	}

	// ServiceAccount: service_account -- string, string
	if v, ok := d.GetOk(prefix + "service_account"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "service_account"))
		val.ServiceAccount = v.(string)
	}

	// ServiceAccountOverwriteAllowed: service_account_overwrite_allowed -- string, string
	if v, ok := d.GetOk(prefix + "service_account_overwrite_allowed"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "service_account_overwrite_allowed"))
		val.ServiceAccountOverwriteAllowed = v.(string)
	}

	// PodAnnotations: pod_annotations -- , map[string]string
	if v, ok := d.GetOk(prefix + "pod_annotations"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "pod_annotations"))
		val.PodAnnotations = v.(map[string]string)
	}

	// PodAnnotationsOverwriteAllowed: pod_annotations_overwrite_allowed -- string, string
	if v, ok := d.GetOk(prefix + "pod_annotations_overwrite_allowed"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "pod_annotations_overwrite_allowed"))
		val.PodAnnotationsOverwriteAllowed = v.(string)
	}

	// PodSecurityContext: pod_security_context -- KubernetesPodSecurityContext, config.KubernetesPodSecurityContext

	if _, ok := d.GetOk(prefix + "pod_security_context.0"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "pod_security_context"))
		thing, err := dsRunnerConfigReadStructConfigKubernetesPodSecurityContext(ctx, prefix+"pod_security_context.0", d)
		if err != nil {
			return val, err
		}
		val.PodSecurityContext = thing
	}

	// BuildContainerSecurityContext: build_container_security_context -- KubernetesContainerSecurityContext, config.KubernetesContainerSecurityContext

	if _, ok := d.GetOk(prefix + "build_container_security_context.0"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "build_container_security_context"))
		thing, err := dsRunnerConfigReadStructConfigKubernetesContainerSecurityContext(ctx, prefix+"build_container_security_context.0", d)
		if err != nil {
			return val, err
		}
		val.BuildContainerSecurityContext = thing
	}

	// HelperContainerSecurityContext: helper_container_security_context -- KubernetesContainerSecurityContext, config.KubernetesContainerSecurityContext

	if _, ok := d.GetOk(prefix + "helper_container_security_context.0"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "helper_container_security_context"))
		thing, err := dsRunnerConfigReadStructConfigKubernetesContainerSecurityContext(ctx, prefix+"helper_container_security_context.0", d)
		if err != nil {
			return val, err
		}
		val.HelperContainerSecurityContext = thing
	}

	// ServiceContainerSecurityContext: service_container_security_context -- KubernetesContainerSecurityContext, config.KubernetesContainerSecurityContext

	if _, ok := d.GetOk(prefix + "service_container_security_context.0"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "service_container_security_context"))
		thing, err := dsRunnerConfigReadStructConfigKubernetesContainerSecurityContext(ctx, prefix+"service_container_security_context.0", d)
		if err != nil {
			return val, err
		}
		val.ServiceContainerSecurityContext = thing
	}

	// Volumes: volumes -- KubernetesVolumes, config.KubernetesVolumes

	if _, ok := d.GetOk(prefix + "volumes.0"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "volumes"))
		thing, err := dsRunnerConfigReadStructConfigKubernetesVolumes(ctx, prefix+"volumes.0", d)
		if err != nil {
			return val, err
		}
		val.Volumes = thing
	}

	// HostAliases: host_aliases -- , []config.KubernetesHostAliases

	// HERE
	if _, ok := d.GetOk(prefix + "host_aliases"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "host_aliases"))
		i := 0
		val.HostAliases = []config.KubernetesHostAliases{}
		for {
			pfx := fmt.Sprintf("%s.%s.%d", prefix, "host_aliases", i)
			if _, ok := d.GetOk(pfx); ok {
				tflog.Debug(ctx, "set: %s", pfx)
				thing, err := dsRunnerConfigReadStructConfigKubernetesHostAliases(ctx, pfx, d)
				if err != nil {
					return val, err
				}
				val.HostAliases = append(val.HostAliases, thing)
				i++
			} else {
				tflog.Debug(ctx, "not set: %s", pfx)
				break
			}
		}
	}

	// Services: services -- , []config.Service

	// HERE
	if _, ok := d.GetOk(prefix + "services"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "services"))
		i := 0
		val.Services = []config.Service{}
		for {
			pfx := fmt.Sprintf("%s.%s.%d", prefix, "services", i)
			if _, ok := d.GetOk(pfx); ok {
				tflog.Debug(ctx, "set: %s", pfx)
				thing, err := dsRunnerConfigReadStructConfigService(ctx, pfx, d)
				if err != nil {
					return val, err
				}
				val.Services = append(val.Services, thing)
				i++
			} else {
				tflog.Debug(ctx, "not set: %s", pfx)
				break
			}
		}
	}

	// CapAdd: cap_add -- , []string

	// HERE
	if _, ok := d.GetOk(prefix + "cap_add"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "cap_add"))
		i := 0
		val.CapAdd = []string{}
		for {
			pfx := fmt.Sprintf("%s.%s.%d", prefix, "cap_add", i)
			if v, ok := d.GetOk(pfx); ok {
				tflog.Debug(ctx, fmt.Sprintf("key is set: %s", pfx))
				val.CapAdd = append(val.CapAdd, v.(string))
				i++
			} else {
				tflog.Debug(ctx, "not set: %s", pfx)
				break
			}
		}
	}

	// CapDrop: cap_drop -- , []string

	// HERE
	if _, ok := d.GetOk(prefix + "cap_drop"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "cap_drop"))
		i := 0
		val.CapDrop = []string{}
		for {
			pfx := fmt.Sprintf("%s.%s.%d", prefix, "cap_drop", i)
			if v, ok := d.GetOk(pfx); ok {
				tflog.Debug(ctx, fmt.Sprintf("key is set: %s", pfx))
				val.CapDrop = append(val.CapDrop, v.(string))
				i++
			} else {
				tflog.Debug(ctx, "not set: %s", pfx)
				break
			}
		}
	}

	// DNSPolicy: dns_policy -- KubernetesDNSPolicy, config.KubernetesDNSPolicy
	// FIXME unhandled type config.KubernetesDNSPolicy

	// DNSConfig: dns_config -- KubernetesDNSConfig, config.KubernetesDNSConfig

	if _, ok := d.GetOk(prefix + "dns_config.0"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "dns_config"))
		thing, err := dsRunnerConfigReadStructConfigKubernetesDNSConfig(ctx, prefix+"dns_config.0", d)
		if err != nil {
			return val, err
		}
		val.DNSConfig = thing
	}

	// ContainerLifecycle: container_lifecycle -- KubernetesContainerLifecyle, config.KubernetesContainerLifecyle

	if _, ok := d.GetOk(prefix + "container_lifecycle.0"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "container_lifecycle"))
		thing, err := dsRunnerConfigReadStructConfigKubernetesContainerLifecyle(ctx, prefix+"container_lifecycle.0", d)
		if err != nil {
			return val, err
		}
		val.ContainerLifecycle = thing
	}

	return val, nil
}

// config.KubernetesConfigMap

func dsRunnerConfigReadStructConfigKubernetesConfigMap(ctx context.Context, prefix string, d *schema.ResourceData) (config.KubernetesConfigMap, error) {

	if prefix != "" {
		prefix = prefix + "."
	}

	tflog.Debug(ctx, fmt.Sprintf("beginning dsRunnerConfigReadStructConfigKubernetesConfigMap run; prefix is '%s'", prefix))

	val := config.KubernetesConfigMap{}

	// Name: name -- string, string
	if v, ok := d.GetOk(prefix + "name"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "name"))
		val.Name = v.(string)
	}

	// MountPath: mount_path -- string, string
	if v, ok := d.GetOk(prefix + "mount_path"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "mount_path"))
		val.MountPath = v.(string)
	}

	// SubPath: sub_path -- string, string
	if v, ok := d.GetOk(prefix + "sub_path"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "sub_path"))
		val.SubPath = v.(string)
	}

	// ReadOnly: read_only -- bool, bool
	if v, ok := d.GetOk(prefix + "read_only"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "read_only"))
		val.ReadOnly = v.(bool)
	}

	// Items: items -- , map[string]string
	if v, ok := d.GetOk(prefix + "items"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "items"))
		val.Items = v.(map[string]string)
	}

	return val, nil
}

// config.KubernetesContainerCapabilities

func dsRunnerConfigReadStructConfigKubernetesContainerCapabilities(ctx context.Context, prefix string, d *schema.ResourceData) (config.KubernetesContainerCapabilities, error) {

	if prefix != "" {
		prefix = prefix + "."
	}

	tflog.Debug(ctx, fmt.Sprintf("beginning dsRunnerConfigReadStructConfigKubernetesContainerCapabilities run; prefix is '%s'", prefix))

	val := config.KubernetesContainerCapabilities{}

	// Add: add -- , []string

	// HERE
	if _, ok := d.GetOk(prefix + "add"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "add"))
		i := 0
		val.Add = []string{}
		for {
			pfx := fmt.Sprintf("%s.%s.%d", prefix, "add", i)
			if v, ok := d.GetOk(pfx); ok {
				tflog.Debug(ctx, fmt.Sprintf("key is set: %s", pfx))
				val.Add = append(val.Add, v.(string))
				i++
			} else {
				tflog.Debug(ctx, "not set: %s", pfx)
				break
			}
		}
	}

	// Drop: drop -- , []string

	// HERE
	if _, ok := d.GetOk(prefix + "drop"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "drop"))
		i := 0
		val.Drop = []string{}
		for {
			pfx := fmt.Sprintf("%s.%s.%d", prefix, "drop", i)
			if v, ok := d.GetOk(pfx); ok {
				tflog.Debug(ctx, fmt.Sprintf("key is set: %s", pfx))
				val.Drop = append(val.Drop, v.(string))
				i++
			} else {
				tflog.Debug(ctx, "not set: %s", pfx)
				break
			}
		}
	}

	return val, nil
}

// config.KubernetesContainerLifecyle

func dsRunnerConfigReadStructConfigKubernetesContainerLifecyle(ctx context.Context, prefix string, d *schema.ResourceData) (config.KubernetesContainerLifecyle, error) {

	if prefix != "" {
		prefix = prefix + "."
	}

	tflog.Debug(ctx, fmt.Sprintf("beginning dsRunnerConfigReadStructConfigKubernetesContainerLifecyle run; prefix is '%s'", prefix))

	val := config.KubernetesContainerLifecyle{}

	// PostStart: post_start -- , *config.KubernetesLifecycleHandler

	if _, ok := d.GetOk(prefix + "post_start.0"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "post_start"))
		thing, err := dsRunnerConfigReadStructConfigKubernetesLifecycleHandler(ctx, prefix+"post_start.0", d)
		if err != nil {
			return val, err
		}
		val.PostStart = &thing
	} else {
		tflog.Debug(ctx, fmt.Sprintf("not set: %s", prefix+"post_start"))
	}

	// PreStop: pre_stop -- , *config.KubernetesLifecycleHandler

	if _, ok := d.GetOk(prefix + "pre_stop.0"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "pre_stop"))
		thing, err := dsRunnerConfigReadStructConfigKubernetesLifecycleHandler(ctx, prefix+"pre_stop.0", d)
		if err != nil {
			return val, err
		}
		val.PreStop = &thing
	} else {
		tflog.Debug(ctx, fmt.Sprintf("not set: %s", prefix+"pre_stop"))
	}

	return val, nil
}

// config.KubernetesContainerSecurityContext

func dsRunnerConfigReadStructConfigKubernetesContainerSecurityContext(ctx context.Context, prefix string, d *schema.ResourceData) (config.KubernetesContainerSecurityContext, error) {

	if prefix != "" {
		prefix = prefix + "."
	}

	tflog.Debug(ctx, fmt.Sprintf("beginning dsRunnerConfigReadStructConfigKubernetesContainerSecurityContext run; prefix is '%s'", prefix))

	val := config.KubernetesContainerSecurityContext{}

	// Capabilities: capabilities -- , *config.KubernetesContainerCapabilities

	if _, ok := d.GetOk(prefix + "capabilities.0"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "capabilities"))
		thing, err := dsRunnerConfigReadStructConfigKubernetesContainerCapabilities(ctx, prefix+"capabilities.0", d)
		if err != nil {
			return val, err
		}
		val.Capabilities = &thing
	} else {
		tflog.Debug(ctx, fmt.Sprintf("not set: %s", prefix+"capabilities"))
	}

	// Privileged: privileged -- , *bool
	if v, ok := d.GetOk(prefix + "privileged"); ok {
		tflog.Debug(ctx, "set: %s.%s", prefix, "privileged")
		val.Privileged = to.BoolP(v.(bool))
	}

	// RunAsUser: run_as_user -- , *int64
	if v, ok := d.GetOk(prefix + "run_as_user"); ok {
		tflog.Debug(ctx, "set: %s.%s", prefix, "run_as_user")
		val.RunAsUser = to.Int64P(v.(int64))
	}

	// RunAsGroup: run_as_group -- , *int64
	if v, ok := d.GetOk(prefix + "run_as_group"); ok {
		tflog.Debug(ctx, "set: %s.%s", prefix, "run_as_group")
		val.RunAsGroup = to.Int64P(v.(int64))
	}

	// RunAsNonRoot: run_as_non_root -- , *bool
	if v, ok := d.GetOk(prefix + "run_as_non_root"); ok {
		tflog.Debug(ctx, "set: %s.%s", prefix, "run_as_non_root")
		val.RunAsNonRoot = to.BoolP(v.(bool))
	}

	// ReadOnlyRootFilesystem: read_only_root_filesystem -- , *bool
	if v, ok := d.GetOk(prefix + "read_only_root_filesystem"); ok {
		tflog.Debug(ctx, "set: %s.%s", prefix, "read_only_root_filesystem")
		val.ReadOnlyRootFilesystem = to.BoolP(v.(bool))
	}

	// AllowPrivilegeEscalation: allow_privilege_escalation -- , *bool
	if v, ok := d.GetOk(prefix + "allow_privilege_escalation"); ok {
		tflog.Debug(ctx, "set: %s.%s", prefix, "allow_privilege_escalation")
		val.AllowPrivilegeEscalation = to.BoolP(v.(bool))
	}

	return val, nil
}

// config.KubernetesDNSConfig

func dsRunnerConfigReadStructConfigKubernetesDNSConfig(ctx context.Context, prefix string, d *schema.ResourceData) (config.KubernetesDNSConfig, error) {

	if prefix != "" {
		prefix = prefix + "."
	}

	tflog.Debug(ctx, fmt.Sprintf("beginning dsRunnerConfigReadStructConfigKubernetesDNSConfig run; prefix is '%s'", prefix))

	val := config.KubernetesDNSConfig{}

	// Nameservers: nameservers -- , []string

	// HERE
	if _, ok := d.GetOk(prefix + "nameservers"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "nameservers"))
		i := 0
		val.Nameservers = []string{}
		for {
			pfx := fmt.Sprintf("%s.%s.%d", prefix, "nameservers", i)
			if v, ok := d.GetOk(pfx); ok {
				tflog.Debug(ctx, fmt.Sprintf("key is set: %s", pfx))
				val.Nameservers = append(val.Nameservers, v.(string))
				i++
			} else {
				tflog.Debug(ctx, "not set: %s", pfx)
				break
			}
		}
	}

	// Options: options -- , []config.KubernetesDNSConfigOption

	// HERE
	if _, ok := d.GetOk(prefix + "options"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "options"))
		i := 0
		val.Options = []config.KubernetesDNSConfigOption{}
		for {
			pfx := fmt.Sprintf("%s.%s.%d", prefix, "options", i)
			if _, ok := d.GetOk(pfx); ok {
				tflog.Debug(ctx, "set: %s", pfx)
				thing, err := dsRunnerConfigReadStructConfigKubernetesDNSConfigOption(ctx, pfx, d)
				if err != nil {
					return val, err
				}
				val.Options = append(val.Options, thing)
				i++
			} else {
				tflog.Debug(ctx, "not set: %s", pfx)
				break
			}
		}
	}

	// Searches: searches -- , []string

	// HERE
	if _, ok := d.GetOk(prefix + "searches"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "searches"))
		i := 0
		val.Searches = []string{}
		for {
			pfx := fmt.Sprintf("%s.%s.%d", prefix, "searches", i)
			if v, ok := d.GetOk(pfx); ok {
				tflog.Debug(ctx, fmt.Sprintf("key is set: %s", pfx))
				val.Searches = append(val.Searches, v.(string))
				i++
			} else {
				tflog.Debug(ctx, "not set: %s", pfx)
				break
			}
		}
	}

	return val, nil
}

// config.KubernetesDNSConfigOption

func dsRunnerConfigReadStructConfigKubernetesDNSConfigOption(ctx context.Context, prefix string, d *schema.ResourceData) (config.KubernetesDNSConfigOption, error) {

	if prefix != "" {
		prefix = prefix + "."
	}

	tflog.Debug(ctx, fmt.Sprintf("beginning dsRunnerConfigReadStructConfigKubernetesDNSConfigOption run; prefix is '%s'", prefix))

	val := config.KubernetesDNSConfigOption{}

	// Name: name -- string, string
	if v, ok := d.GetOk(prefix + "name"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "name"))
		val.Name = v.(string)
	}

	// Value: value -- , *string
	if v, ok := d.GetOk(prefix + "value"); ok {
		tflog.Debug(ctx, "set: %s.%s", prefix, "value")
		val.Value = to.StringP(v.(string))
	}

	return val, nil
}

// config.KubernetesEmptyDir

func dsRunnerConfigReadStructConfigKubernetesEmptyDir(ctx context.Context, prefix string, d *schema.ResourceData) (config.KubernetesEmptyDir, error) {

	if prefix != "" {
		prefix = prefix + "."
	}

	tflog.Debug(ctx, fmt.Sprintf("beginning dsRunnerConfigReadStructConfigKubernetesEmptyDir run; prefix is '%s'", prefix))

	val := config.KubernetesEmptyDir{}

	// Name: name -- string, string
	if v, ok := d.GetOk(prefix + "name"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "name"))
		val.Name = v.(string)
	}

	// MountPath: mount_path -- string, string
	if v, ok := d.GetOk(prefix + "mount_path"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "mount_path"))
		val.MountPath = v.(string)
	}

	// SubPath: sub_path -- string, string
	if v, ok := d.GetOk(prefix + "sub_path"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "sub_path"))
		val.SubPath = v.(string)
	}

	// Medium: medium -- string, string
	if v, ok := d.GetOk(prefix + "medium"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "medium"))
		val.Medium = v.(string)
	}

	return val, nil
}

// config.KubernetesHostAliases

func dsRunnerConfigReadStructConfigKubernetesHostAliases(ctx context.Context, prefix string, d *schema.ResourceData) (config.KubernetesHostAliases, error) {

	if prefix != "" {
		prefix = prefix + "."
	}

	tflog.Debug(ctx, fmt.Sprintf("beginning dsRunnerConfigReadStructConfigKubernetesHostAliases run; prefix is '%s'", prefix))

	val := config.KubernetesHostAliases{}

	// IP: ip -- string, string
	if v, ok := d.GetOk(prefix + "ip"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "ip"))
		val.IP = v.(string)
	}

	// Hostnames: hostnames -- , []string

	// HERE
	if _, ok := d.GetOk(prefix + "hostnames"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "hostnames"))
		i := 0
		val.Hostnames = []string{}
		for {
			pfx := fmt.Sprintf("%s.%s.%d", prefix, "hostnames", i)
			if v, ok := d.GetOk(pfx); ok {
				tflog.Debug(ctx, fmt.Sprintf("key is set: %s", pfx))
				val.Hostnames = append(val.Hostnames, v.(string))
				i++
			} else {
				tflog.Debug(ctx, "not set: %s", pfx)
				break
			}
		}
	}

	return val, nil
}

// config.KubernetesHostPath

func dsRunnerConfigReadStructConfigKubernetesHostPath(ctx context.Context, prefix string, d *schema.ResourceData) (config.KubernetesHostPath, error) {

	if prefix != "" {
		prefix = prefix + "."
	}

	tflog.Debug(ctx, fmt.Sprintf("beginning dsRunnerConfigReadStructConfigKubernetesHostPath run; prefix is '%s'", prefix))

	val := config.KubernetesHostPath{}

	// Name: name -- string, string
	if v, ok := d.GetOk(prefix + "name"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "name"))
		val.Name = v.(string)
	}

	// MountPath: mount_path -- string, string
	if v, ok := d.GetOk(prefix + "mount_path"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "mount_path"))
		val.MountPath = v.(string)
	}

	// SubPath: sub_path -- string, string
	if v, ok := d.GetOk(prefix + "sub_path"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "sub_path"))
		val.SubPath = v.(string)
	}

	// ReadOnly: read_only -- bool, bool
	if v, ok := d.GetOk(prefix + "read_only"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "read_only"))
		val.ReadOnly = v.(bool)
	}

	// HostPath: host_path -- string, string
	if v, ok := d.GetOk(prefix + "host_path"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "host_path"))
		val.HostPath = v.(string)
	}

	return val, nil
}

// config.KubernetesLifecycleExecAction

func dsRunnerConfigReadStructConfigKubernetesLifecycleExecAction(ctx context.Context, prefix string, d *schema.ResourceData) (config.KubernetesLifecycleExecAction, error) {

	if prefix != "" {
		prefix = prefix + "."
	}

	tflog.Debug(ctx, fmt.Sprintf("beginning dsRunnerConfigReadStructConfigKubernetesLifecycleExecAction run; prefix is '%s'", prefix))

	val := config.KubernetesLifecycleExecAction{}

	// Command: command -- , []string

	// HERE
	if _, ok := d.GetOk(prefix + "command"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "command"))
		i := 0
		val.Command = []string{}
		for {
			pfx := fmt.Sprintf("%s.%s.%d", prefix, "command", i)
			if v, ok := d.GetOk(pfx); ok {
				tflog.Debug(ctx, fmt.Sprintf("key is set: %s", pfx))
				val.Command = append(val.Command, v.(string))
				i++
			} else {
				tflog.Debug(ctx, "not set: %s", pfx)
				break
			}
		}
	}

	return val, nil
}

// config.KubernetesLifecycleHTTPGet

func dsRunnerConfigReadStructConfigKubernetesLifecycleHTTPGet(ctx context.Context, prefix string, d *schema.ResourceData) (config.KubernetesLifecycleHTTPGet, error) {

	if prefix != "" {
		prefix = prefix + "."
	}

	tflog.Debug(ctx, fmt.Sprintf("beginning dsRunnerConfigReadStructConfigKubernetesLifecycleHTTPGet run; prefix is '%s'", prefix))

	val := config.KubernetesLifecycleHTTPGet{}

	// Host: host -- string, string
	if v, ok := d.GetOk(prefix + "host"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "host"))
		val.Host = v.(string)
	}

	// HTTPHeaders: http_headers -- , []config.KubernetesLifecycleHTTPGetHeader

	// HERE
	if _, ok := d.GetOk(prefix + "http_headers"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "http_headers"))
		i := 0
		val.HTTPHeaders = []config.KubernetesLifecycleHTTPGetHeader{}
		for {
			pfx := fmt.Sprintf("%s.%s.%d", prefix, "http_headers", i)
			if _, ok := d.GetOk(pfx); ok {
				tflog.Debug(ctx, "set: %s", pfx)
				thing, err := dsRunnerConfigReadStructConfigKubernetesLifecycleHTTPGetHeader(ctx, pfx, d)
				if err != nil {
					return val, err
				}
				val.HTTPHeaders = append(val.HTTPHeaders, thing)
				i++
			} else {
				tflog.Debug(ctx, "not set: %s", pfx)
				break
			}
		}
	}

	// Path: path -- string, string
	if v, ok := d.GetOk(prefix + "path"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "path"))
		val.Path = v.(string)
	}

	// Port: port -- int, int
	if v, ok := d.GetOk(prefix + "port"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "port"))
		val.Port = v.(int)
	}

	// Scheme: scheme -- string, string
	if v, ok := d.GetOk(prefix + "scheme"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "scheme"))
		val.Scheme = v.(string)
	}

	return val, nil
}

// config.KubernetesLifecycleHTTPGetHeader

func dsRunnerConfigReadStructConfigKubernetesLifecycleHTTPGetHeader(ctx context.Context, prefix string, d *schema.ResourceData) (config.KubernetesLifecycleHTTPGetHeader, error) {

	if prefix != "" {
		prefix = prefix + "."
	}

	tflog.Debug(ctx, fmt.Sprintf("beginning dsRunnerConfigReadStructConfigKubernetesLifecycleHTTPGetHeader run; prefix is '%s'", prefix))

	val := config.KubernetesLifecycleHTTPGetHeader{}

	// Name: name -- string, string
	if v, ok := d.GetOk(prefix + "name"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "name"))
		val.Name = v.(string)
	}

	// Value: value -- string, string
	if v, ok := d.GetOk(prefix + "value"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "value"))
		val.Value = v.(string)
	}

	return val, nil
}

// config.KubernetesLifecycleHandler

func dsRunnerConfigReadStructConfigKubernetesLifecycleHandler(ctx context.Context, prefix string, d *schema.ResourceData) (config.KubernetesLifecycleHandler, error) {

	if prefix != "" {
		prefix = prefix + "."
	}

	tflog.Debug(ctx, fmt.Sprintf("beginning dsRunnerConfigReadStructConfigKubernetesLifecycleHandler run; prefix is '%s'", prefix))

	val := config.KubernetesLifecycleHandler{}

	// Exec: exec -- , *config.KubernetesLifecycleExecAction

	if _, ok := d.GetOk(prefix + "exec.0"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "exec"))
		thing, err := dsRunnerConfigReadStructConfigKubernetesLifecycleExecAction(ctx, prefix+"exec.0", d)
		if err != nil {
			return val, err
		}
		val.Exec = &thing
	} else {
		tflog.Debug(ctx, fmt.Sprintf("not set: %s", prefix+"exec"))
	}

	// HTTPGet: http_get -- , *config.KubernetesLifecycleHTTPGet

	if _, ok := d.GetOk(prefix + "http_get.0"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "http_get"))
		thing, err := dsRunnerConfigReadStructConfigKubernetesLifecycleHTTPGet(ctx, prefix+"http_get.0", d)
		if err != nil {
			return val, err
		}
		val.HTTPGet = &thing
	} else {
		tflog.Debug(ctx, fmt.Sprintf("not set: %s", prefix+"http_get"))
	}

	// TCPSocket: tcp_socket -- , *config.KubernetesLifecycleTCPSocket

	if _, ok := d.GetOk(prefix + "tcp_socket.0"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "tcp_socket"))
		thing, err := dsRunnerConfigReadStructConfigKubernetesLifecycleTCPSocket(ctx, prefix+"tcp_socket.0", d)
		if err != nil {
			return val, err
		}
		val.TCPSocket = &thing
	} else {
		tflog.Debug(ctx, fmt.Sprintf("not set: %s", prefix+"tcp_socket"))
	}

	return val, nil
}

// config.KubernetesLifecycleTCPSocket

func dsRunnerConfigReadStructConfigKubernetesLifecycleTCPSocket(ctx context.Context, prefix string, d *schema.ResourceData) (config.KubernetesLifecycleTCPSocket, error) {

	if prefix != "" {
		prefix = prefix + "."
	}

	tflog.Debug(ctx, fmt.Sprintf("beginning dsRunnerConfigReadStructConfigKubernetesLifecycleTCPSocket run; prefix is '%s'", prefix))

	val := config.KubernetesLifecycleTCPSocket{}

	// Host: host -- string, string
	if v, ok := d.GetOk(prefix + "host"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "host"))
		val.Host = v.(string)
	}

	// Port: port -- int, int
	if v, ok := d.GetOk(prefix + "port"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "port"))
		val.Port = v.(int)
	}

	return val, nil
}

// config.KubernetesNodeAffinity

func dsRunnerConfigReadStructConfigKubernetesNodeAffinity(ctx context.Context, prefix string, d *schema.ResourceData) (config.KubernetesNodeAffinity, error) {

	if prefix != "" {
		prefix = prefix + "."
	}

	tflog.Debug(ctx, fmt.Sprintf("beginning dsRunnerConfigReadStructConfigKubernetesNodeAffinity run; prefix is '%s'", prefix))

	val := config.KubernetesNodeAffinity{}

	// RequiredDuringSchedulingIgnoredDuringExecution: required_during_scheduling_ignored_during_execution -- , *config.NodeSelector

	if _, ok := d.GetOk(prefix + "required_during_scheduling_ignored_during_execution.0"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "required_during_scheduling_ignored_during_execution"))
		thing, err := dsRunnerConfigReadStructConfigNodeSelector(ctx, prefix+"required_during_scheduling_ignored_during_execution.0", d)
		if err != nil {
			return val, err
		}
		val.RequiredDuringSchedulingIgnoredDuringExecution = &thing
	} else {
		tflog.Debug(ctx, fmt.Sprintf("not set: %s", prefix+"required_during_scheduling_ignored_during_execution"))
	}

	// PreferredDuringSchedulingIgnoredDuringExecution: preferred_during_scheduling_ignored_during_execution -- , []config.PreferredSchedulingTerm

	// HERE
	if _, ok := d.GetOk(prefix + "preferred_during_scheduling_ignored_during_execution"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "preferred_during_scheduling_ignored_during_execution"))
		i := 0
		val.PreferredDuringSchedulingIgnoredDuringExecution = []config.PreferredSchedulingTerm{}
		for {
			pfx := fmt.Sprintf("%s.%s.%d", prefix, "preferred_during_scheduling_ignored_during_execution", i)
			if _, ok := d.GetOk(pfx); ok {
				tflog.Debug(ctx, "set: %s", pfx)
				thing, err := dsRunnerConfigReadStructConfigPreferredSchedulingTerm(ctx, pfx, d)
				if err != nil {
					return val, err
				}
				val.PreferredDuringSchedulingIgnoredDuringExecution = append(val.PreferredDuringSchedulingIgnoredDuringExecution, thing)
				i++
			} else {
				tflog.Debug(ctx, "not set: %s", pfx)
				break
			}
		}
	}

	return val, nil
}

// config.KubernetesPVC

func dsRunnerConfigReadStructConfigKubernetesPVC(ctx context.Context, prefix string, d *schema.ResourceData) (config.KubernetesPVC, error) {

	if prefix != "" {
		prefix = prefix + "."
	}

	tflog.Debug(ctx, fmt.Sprintf("beginning dsRunnerConfigReadStructConfigKubernetesPVC run; prefix is '%s'", prefix))

	val := config.KubernetesPVC{}

	// Name: name -- string, string
	if v, ok := d.GetOk(prefix + "name"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "name"))
		val.Name = v.(string)
	}

	// MountPath: mount_path -- string, string
	if v, ok := d.GetOk(prefix + "mount_path"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "mount_path"))
		val.MountPath = v.(string)
	}

	// SubPath: sub_path -- string, string
	if v, ok := d.GetOk(prefix + "sub_path"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "sub_path"))
		val.SubPath = v.(string)
	}

	// ReadOnly: read_only -- bool, bool
	if v, ok := d.GetOk(prefix + "read_only"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "read_only"))
		val.ReadOnly = v.(bool)
	}

	return val, nil
}

// config.KubernetesPodAffinity

func dsRunnerConfigReadStructConfigKubernetesPodAffinity(ctx context.Context, prefix string, d *schema.ResourceData) (config.KubernetesPodAffinity, error) {

	if prefix != "" {
		prefix = prefix + "."
	}

	tflog.Debug(ctx, fmt.Sprintf("beginning dsRunnerConfigReadStructConfigKubernetesPodAffinity run; prefix is '%s'", prefix))

	val := config.KubernetesPodAffinity{}

	// RequiredDuringSchedulingIgnoredDuringExecution: required_during_scheduling_ignored_during_execution -- , []config.PodAffinityTerm

	// HERE
	if _, ok := d.GetOk(prefix + "required_during_scheduling_ignored_during_execution"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "required_during_scheduling_ignored_during_execution"))
		i := 0
		val.RequiredDuringSchedulingIgnoredDuringExecution = []config.PodAffinityTerm{}
		for {
			pfx := fmt.Sprintf("%s.%s.%d", prefix, "required_during_scheduling_ignored_during_execution", i)
			if _, ok := d.GetOk(pfx); ok {
				tflog.Debug(ctx, "set: %s", pfx)
				thing, err := dsRunnerConfigReadStructConfigPodAffinityTerm(ctx, pfx, d)
				if err != nil {
					return val, err
				}
				val.RequiredDuringSchedulingIgnoredDuringExecution = append(val.RequiredDuringSchedulingIgnoredDuringExecution, thing)
				i++
			} else {
				tflog.Debug(ctx, "not set: %s", pfx)
				break
			}
		}
	}

	// PreferredDuringSchedulingIgnoredDuringExecution: preferred_during_scheduling_ignored_during_execution -- , []config.WeightedPodAffinityTerm

	// HERE
	if _, ok := d.GetOk(prefix + "preferred_during_scheduling_ignored_during_execution"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "preferred_during_scheduling_ignored_during_execution"))
		i := 0
		val.PreferredDuringSchedulingIgnoredDuringExecution = []config.WeightedPodAffinityTerm{}
		for {
			pfx := fmt.Sprintf("%s.%s.%d", prefix, "preferred_during_scheduling_ignored_during_execution", i)
			if _, ok := d.GetOk(pfx); ok {
				tflog.Debug(ctx, "set: %s", pfx)
				thing, err := dsRunnerConfigReadStructConfigWeightedPodAffinityTerm(ctx, pfx, d)
				if err != nil {
					return val, err
				}
				val.PreferredDuringSchedulingIgnoredDuringExecution = append(val.PreferredDuringSchedulingIgnoredDuringExecution, thing)
				i++
			} else {
				tflog.Debug(ctx, "not set: %s", pfx)
				break
			}
		}
	}

	return val, nil
}

// config.KubernetesPodAntiAffinity

func dsRunnerConfigReadStructConfigKubernetesPodAntiAffinity(ctx context.Context, prefix string, d *schema.ResourceData) (config.KubernetesPodAntiAffinity, error) {

	if prefix != "" {
		prefix = prefix + "."
	}

	tflog.Debug(ctx, fmt.Sprintf("beginning dsRunnerConfigReadStructConfigKubernetesPodAntiAffinity run; prefix is '%s'", prefix))

	val := config.KubernetesPodAntiAffinity{}

	// RequiredDuringSchedulingIgnoredDuringExecution: required_during_scheduling_ignored_during_execution -- , []config.PodAffinityTerm

	// HERE
	if _, ok := d.GetOk(prefix + "required_during_scheduling_ignored_during_execution"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "required_during_scheduling_ignored_during_execution"))
		i := 0
		val.RequiredDuringSchedulingIgnoredDuringExecution = []config.PodAffinityTerm{}
		for {
			pfx := fmt.Sprintf("%s.%s.%d", prefix, "required_during_scheduling_ignored_during_execution", i)
			if _, ok := d.GetOk(pfx); ok {
				tflog.Debug(ctx, "set: %s", pfx)
				thing, err := dsRunnerConfigReadStructConfigPodAffinityTerm(ctx, pfx, d)
				if err != nil {
					return val, err
				}
				val.RequiredDuringSchedulingIgnoredDuringExecution = append(val.RequiredDuringSchedulingIgnoredDuringExecution, thing)
				i++
			} else {
				tflog.Debug(ctx, "not set: %s", pfx)
				break
			}
		}
	}

	// PreferredDuringSchedulingIgnoredDuringExecution: preferred_during_scheduling_ignored_during_execution -- , []config.WeightedPodAffinityTerm

	// HERE
	if _, ok := d.GetOk(prefix + "preferred_during_scheduling_ignored_during_execution"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "preferred_during_scheduling_ignored_during_execution"))
		i := 0
		val.PreferredDuringSchedulingIgnoredDuringExecution = []config.WeightedPodAffinityTerm{}
		for {
			pfx := fmt.Sprintf("%s.%s.%d", prefix, "preferred_during_scheduling_ignored_during_execution", i)
			if _, ok := d.GetOk(pfx); ok {
				tflog.Debug(ctx, "set: %s", pfx)
				thing, err := dsRunnerConfigReadStructConfigWeightedPodAffinityTerm(ctx, pfx, d)
				if err != nil {
					return val, err
				}
				val.PreferredDuringSchedulingIgnoredDuringExecution = append(val.PreferredDuringSchedulingIgnoredDuringExecution, thing)
				i++
			} else {
				tflog.Debug(ctx, "not set: %s", pfx)
				break
			}
		}
	}

	return val, nil
}

// config.KubernetesPodSecurityContext

func dsRunnerConfigReadStructConfigKubernetesPodSecurityContext(ctx context.Context, prefix string, d *schema.ResourceData) (config.KubernetesPodSecurityContext, error) {

	if prefix != "" {
		prefix = prefix + "."
	}

	tflog.Debug(ctx, fmt.Sprintf("beginning dsRunnerConfigReadStructConfigKubernetesPodSecurityContext run; prefix is '%s'", prefix))

	val := config.KubernetesPodSecurityContext{}

	// FSGroup: fs_group -- , *int64
	if v, ok := d.GetOk(prefix + "fs_group"); ok {
		tflog.Debug(ctx, "set: %s.%s", prefix, "fs_group")
		val.FSGroup = to.Int64P(v.(int64))
	}

	// RunAsGroup: run_as_group -- , *int64
	if v, ok := d.GetOk(prefix + "run_as_group"); ok {
		tflog.Debug(ctx, "set: %s.%s", prefix, "run_as_group")
		val.RunAsGroup = to.Int64P(v.(int64))
	}

	// RunAsNonRoot: run_as_non_root -- , *bool
	if v, ok := d.GetOk(prefix + "run_as_non_root"); ok {
		tflog.Debug(ctx, "set: %s.%s", prefix, "run_as_non_root")
		val.RunAsNonRoot = to.BoolP(v.(bool))
	}

	// RunAsUser: run_as_user -- , *int64
	if v, ok := d.GetOk(prefix + "run_as_user"); ok {
		tflog.Debug(ctx, "set: %s.%s", prefix, "run_as_user")
		val.RunAsUser = to.Int64P(v.(int64))
	}

	// SupplementalGroups: supplemental_groups -- , []int64

	// HERE
	if _, ok := d.GetOk(prefix + "supplemental_groups"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "supplemental_groups"))
		i := 0
		val.SupplementalGroups = []int64{}
		for {
			pfx := fmt.Sprintf("%s.%s.%d", prefix, "supplemental_groups", i)
			if v, ok := d.GetOk(pfx); ok {
				tflog.Debug(ctx, fmt.Sprintf("key is set: %s", pfx))
				val.SupplementalGroups = append(val.SupplementalGroups, v.(int64))
				i++
			} else {
				tflog.Debug(ctx, "not set: %s", pfx)
				break
			}
		}
	}

	return val, nil
}

// config.KubernetesSecret

func dsRunnerConfigReadStructConfigKubernetesSecret(ctx context.Context, prefix string, d *schema.ResourceData) (config.KubernetesSecret, error) {

	if prefix != "" {
		prefix = prefix + "."
	}

	tflog.Debug(ctx, fmt.Sprintf("beginning dsRunnerConfigReadStructConfigKubernetesSecret run; prefix is '%s'", prefix))

	val := config.KubernetesSecret{}

	// Name: name -- string, string
	if v, ok := d.GetOk(prefix + "name"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "name"))
		val.Name = v.(string)
	}

	// MountPath: mount_path -- string, string
	if v, ok := d.GetOk(prefix + "mount_path"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "mount_path"))
		val.MountPath = v.(string)
	}

	// SubPath: sub_path -- string, string
	if v, ok := d.GetOk(prefix + "sub_path"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "sub_path"))
		val.SubPath = v.(string)
	}

	// ReadOnly: read_only -- bool, bool
	if v, ok := d.GetOk(prefix + "read_only"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "read_only"))
		val.ReadOnly = v.(bool)
	}

	// Items: items -- , map[string]string
	if v, ok := d.GetOk(prefix + "items"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "items"))
		val.Items = v.(map[string]string)
	}

	return val, nil
}

// config.KubernetesVolumes

func dsRunnerConfigReadStructConfigKubernetesVolumes(ctx context.Context, prefix string, d *schema.ResourceData) (config.KubernetesVolumes, error) {

	if prefix != "" {
		prefix = prefix + "."
	}

	tflog.Debug(ctx, fmt.Sprintf("beginning dsRunnerConfigReadStructConfigKubernetesVolumes run; prefix is '%s'", prefix))

	val := config.KubernetesVolumes{}

	// HostPaths: host_path -- , []config.KubernetesHostPath

	// HERE
	if _, ok := d.GetOk(prefix + "host_path"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "host_path"))
		i := 0
		val.HostPaths = []config.KubernetesHostPath{}
		for {
			pfx := fmt.Sprintf("%s.%s.%d", prefix, "host_path", i)
			if _, ok := d.GetOk(pfx); ok {
				tflog.Debug(ctx, "set: %s", pfx)
				thing, err := dsRunnerConfigReadStructConfigKubernetesHostPath(ctx, pfx, d)
				if err != nil {
					return val, err
				}
				val.HostPaths = append(val.HostPaths, thing)
				i++
			} else {
				tflog.Debug(ctx, "not set: %s", pfx)
				break
			}
		}
	}

	// PVCs: pvc -- , []config.KubernetesPVC

	// HERE
	if _, ok := d.GetOk(prefix + "pvc"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "pvc"))
		i := 0
		val.PVCs = []config.KubernetesPVC{}
		for {
			pfx := fmt.Sprintf("%s.%s.%d", prefix, "pvc", i)
			if _, ok := d.GetOk(pfx); ok {
				tflog.Debug(ctx, "set: %s", pfx)
				thing, err := dsRunnerConfigReadStructConfigKubernetesPVC(ctx, pfx, d)
				if err != nil {
					return val, err
				}
				val.PVCs = append(val.PVCs, thing)
				i++
			} else {
				tflog.Debug(ctx, "not set: %s", pfx)
				break
			}
		}
	}

	// ConfigMaps: config_map -- , []config.KubernetesConfigMap

	// HERE
	if _, ok := d.GetOk(prefix + "config_map"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "config_map"))
		i := 0
		val.ConfigMaps = []config.KubernetesConfigMap{}
		for {
			pfx := fmt.Sprintf("%s.%s.%d", prefix, "config_map", i)
			if _, ok := d.GetOk(pfx); ok {
				tflog.Debug(ctx, "set: %s", pfx)
				thing, err := dsRunnerConfigReadStructConfigKubernetesConfigMap(ctx, pfx, d)
				if err != nil {
					return val, err
				}
				val.ConfigMaps = append(val.ConfigMaps, thing)
				i++
			} else {
				tflog.Debug(ctx, "not set: %s", pfx)
				break
			}
		}
	}

	// Secrets: secret -- , []config.KubernetesSecret

	// HERE
	if _, ok := d.GetOk(prefix + "secret"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "secret"))
		i := 0
		val.Secrets = []config.KubernetesSecret{}
		for {
			pfx := fmt.Sprintf("%s.%s.%d", prefix, "secret", i)
			if _, ok := d.GetOk(pfx); ok {
				tflog.Debug(ctx, "set: %s", pfx)
				thing, err := dsRunnerConfigReadStructConfigKubernetesSecret(ctx, pfx, d)
				if err != nil {
					return val, err
				}
				val.Secrets = append(val.Secrets, thing)
				i++
			} else {
				tflog.Debug(ctx, "not set: %s", pfx)
				break
			}
		}
	}

	// EmptyDirs: empty_dir -- , []config.KubernetesEmptyDir

	// HERE
	if _, ok := d.GetOk(prefix + "empty_dir"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "empty_dir"))
		i := 0
		val.EmptyDirs = []config.KubernetesEmptyDir{}
		for {
			pfx := fmt.Sprintf("%s.%s.%d", prefix, "empty_dir", i)
			if _, ok := d.GetOk(pfx); ok {
				tflog.Debug(ctx, "set: %s", pfx)
				thing, err := dsRunnerConfigReadStructConfigKubernetesEmptyDir(ctx, pfx, d)
				if err != nil {
					return val, err
				}
				val.EmptyDirs = append(val.EmptyDirs, thing)
				i++
			} else {
				tflog.Debug(ctx, "not set: %s", pfx)
				break
			}
		}
	}

	// CSIs: csi -- , []config.KubernetesCSI

	// HERE
	if _, ok := d.GetOk(prefix + "csi"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "csi"))
		i := 0
		val.CSIs = []config.KubernetesCSI{}
		for {
			pfx := fmt.Sprintf("%s.%s.%d", prefix, "csi", i)
			if _, ok := d.GetOk(pfx); ok {
				tflog.Debug(ctx, "set: %s", pfx)
				thing, err := dsRunnerConfigReadStructConfigKubernetesCSI(ctx, pfx, d)
				if err != nil {
					return val, err
				}
				val.CSIs = append(val.CSIs, thing)
				i++
			} else {
				tflog.Debug(ctx, "not set: %s", pfx)
				break
			}
		}
	}

	return val, nil
}

// config.LabelSelector

func dsRunnerConfigReadStructConfigLabelSelector(ctx context.Context, prefix string, d *schema.ResourceData) (config.LabelSelector, error) {

	if prefix != "" {
		prefix = prefix + "."
	}

	tflog.Debug(ctx, fmt.Sprintf("beginning dsRunnerConfigReadStructConfigLabelSelector run; prefix is '%s'", prefix))

	val := config.LabelSelector{}

	// MatchLabels: match_labels -- , map[string]string
	if v, ok := d.GetOk(prefix + "match_labels"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "match_labels"))
		val.MatchLabels = v.(map[string]string)
	}

	// MatchExpressions: match_expressions -- , []config.NodeSelectorRequirement

	// HERE
	if _, ok := d.GetOk(prefix + "match_expressions"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "match_expressions"))
		i := 0
		val.MatchExpressions = []config.NodeSelectorRequirement{}
		for {
			pfx := fmt.Sprintf("%s.%s.%d", prefix, "match_expressions", i)
			if _, ok := d.GetOk(pfx); ok {
				tflog.Debug(ctx, "set: %s", pfx)
				thing, err := dsRunnerConfigReadStructConfigNodeSelectorRequirement(ctx, pfx, d)
				if err != nil {
					return val, err
				}
				val.MatchExpressions = append(val.MatchExpressions, thing)
				i++
			} else {
				tflog.Debug(ctx, "not set: %s", pfx)
				break
			}
		}
	}

	return val, nil
}

// config.NodeSelector

func dsRunnerConfigReadStructConfigNodeSelector(ctx context.Context, prefix string, d *schema.ResourceData) (config.NodeSelector, error) {

	if prefix != "" {
		prefix = prefix + "."
	}

	tflog.Debug(ctx, fmt.Sprintf("beginning dsRunnerConfigReadStructConfigNodeSelector run; prefix is '%s'", prefix))

	val := config.NodeSelector{}

	// NodeSelectorTerms: node_selector_terms -- , []config.NodeSelectorTerm

	// HERE
	if _, ok := d.GetOk(prefix + "node_selector_terms"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "node_selector_terms"))
		i := 0
		val.NodeSelectorTerms = []config.NodeSelectorTerm{}
		for {
			pfx := fmt.Sprintf("%s.%s.%d", prefix, "node_selector_terms", i)
			if _, ok := d.GetOk(pfx); ok {
				tflog.Debug(ctx, "set: %s", pfx)
				thing, err := dsRunnerConfigReadStructConfigNodeSelectorTerm(ctx, pfx, d)
				if err != nil {
					return val, err
				}
				val.NodeSelectorTerms = append(val.NodeSelectorTerms, thing)
				i++
			} else {
				tflog.Debug(ctx, "not set: %s", pfx)
				break
			}
		}
	}

	return val, nil
}

// config.NodeSelectorRequirement

func dsRunnerConfigReadStructConfigNodeSelectorRequirement(ctx context.Context, prefix string, d *schema.ResourceData) (config.NodeSelectorRequirement, error) {

	if prefix != "" {
		prefix = prefix + "."
	}

	tflog.Debug(ctx, fmt.Sprintf("beginning dsRunnerConfigReadStructConfigNodeSelectorRequirement run; prefix is '%s'", prefix))

	val := config.NodeSelectorRequirement{}

	// Key: key -- string, string
	if v, ok := d.GetOk(prefix + "key"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "key"))
		val.Key = v.(string)
	}

	// Operator: operator -- string, string
	if v, ok := d.GetOk(prefix + "operator"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "operator"))
		val.Operator = v.(string)
	}

	// Values: values -- , []string

	// HERE
	if _, ok := d.GetOk(prefix + "values"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "values"))
		i := 0
		val.Values = []string{}
		for {
			pfx := fmt.Sprintf("%s.%s.%d", prefix, "values", i)
			if v, ok := d.GetOk(pfx); ok {
				tflog.Debug(ctx, fmt.Sprintf("key is set: %s", pfx))
				val.Values = append(val.Values, v.(string))
				i++
			} else {
				tflog.Debug(ctx, "not set: %s", pfx)
				break
			}
		}
	}

	return val, nil
}

// config.NodeSelectorTerm

func dsRunnerConfigReadStructConfigNodeSelectorTerm(ctx context.Context, prefix string, d *schema.ResourceData) (config.NodeSelectorTerm, error) {

	if prefix != "" {
		prefix = prefix + "."
	}

	tflog.Debug(ctx, fmt.Sprintf("beginning dsRunnerConfigReadStructConfigNodeSelectorTerm run; prefix is '%s'", prefix))

	val := config.NodeSelectorTerm{}

	// MatchExpressions: match_expressions -- , []config.NodeSelectorRequirement

	// HERE
	if _, ok := d.GetOk(prefix + "match_expressions"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "match_expressions"))
		i := 0
		val.MatchExpressions = []config.NodeSelectorRequirement{}
		for {
			pfx := fmt.Sprintf("%s.%s.%d", prefix, "match_expressions", i)
			if _, ok := d.GetOk(pfx); ok {
				tflog.Debug(ctx, "set: %s", pfx)
				thing, err := dsRunnerConfigReadStructConfigNodeSelectorRequirement(ctx, pfx, d)
				if err != nil {
					return val, err
				}
				val.MatchExpressions = append(val.MatchExpressions, thing)
				i++
			} else {
				tflog.Debug(ctx, "not set: %s", pfx)
				break
			}
		}
	}

	// MatchFields: match_fields -- , []config.NodeSelectorRequirement

	// HERE
	if _, ok := d.GetOk(prefix + "match_fields"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "match_fields"))
		i := 0
		val.MatchFields = []config.NodeSelectorRequirement{}
		for {
			pfx := fmt.Sprintf("%s.%s.%d", prefix, "match_fields", i)
			if _, ok := d.GetOk(pfx); ok {
				tflog.Debug(ctx, "set: %s", pfx)
				thing, err := dsRunnerConfigReadStructConfigNodeSelectorRequirement(ctx, pfx, d)
				if err != nil {
					return val, err
				}
				val.MatchFields = append(val.MatchFields, thing)
				i++
			} else {
				tflog.Debug(ctx, "not set: %s", pfx)
				break
			}
		}
	}

	return val, nil
}

// config.ParallelsConfig

func dsRunnerConfigReadStructConfigParallelsConfig(ctx context.Context, prefix string, d *schema.ResourceData) (config.ParallelsConfig, error) {

	if prefix != "" {
		prefix = prefix + "."
	}

	tflog.Debug(ctx, fmt.Sprintf("beginning dsRunnerConfigReadStructConfigParallelsConfig run; prefix is '%s'", prefix))

	val := config.ParallelsConfig{}

	// BaseName: base_name -- string, string
	if v, ok := d.GetOk(prefix + "base_name"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "base_name"))
		val.BaseName = v.(string)
	}

	// TemplateName: template_name -- string, string
	if v, ok := d.GetOk(prefix + "template_name"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "template_name"))
		val.TemplateName = v.(string)
	}

	// DisableSnapshots: disable_snapshots -- bool, bool
	if v, ok := d.GetOk(prefix + "disable_snapshots"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "disable_snapshots"))
		val.DisableSnapshots = v.(bool)
	}

	// TimeServer: time_server -- string, string
	if v, ok := d.GetOk(prefix + "time_server"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "time_server"))
		val.TimeServer = v.(string)
	}

	// AllowedImages: allowed_images -- , []string

	// HERE
	if _, ok := d.GetOk(prefix + "allowed_images"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "allowed_images"))
		i := 0
		val.AllowedImages = []string{}
		for {
			pfx := fmt.Sprintf("%s.%s.%d", prefix, "allowed_images", i)
			if v, ok := d.GetOk(pfx); ok {
				tflog.Debug(ctx, fmt.Sprintf("key is set: %s", pfx))
				val.AllowedImages = append(val.AllowedImages, v.(string))
				i++
			} else {
				tflog.Debug(ctx, "not set: %s", pfx)
				break
			}
		}
	}

	return val, nil
}

// config.PodAffinityTerm

func dsRunnerConfigReadStructConfigPodAffinityTerm(ctx context.Context, prefix string, d *schema.ResourceData) (config.PodAffinityTerm, error) {

	if prefix != "" {
		prefix = prefix + "."
	}

	tflog.Debug(ctx, fmt.Sprintf("beginning dsRunnerConfigReadStructConfigPodAffinityTerm run; prefix is '%s'", prefix))

	val := config.PodAffinityTerm{}

	// LabelSelector: label_selector -- , *config.LabelSelector

	if _, ok := d.GetOk(prefix + "label_selector.0"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "label_selector"))
		thing, err := dsRunnerConfigReadStructConfigLabelSelector(ctx, prefix+"label_selector.0", d)
		if err != nil {
			return val, err
		}
		val.LabelSelector = &thing
	} else {
		tflog.Debug(ctx, fmt.Sprintf("not set: %s", prefix+"label_selector"))
	}

	// Namespaces: namespaces -- , []string

	// HERE
	if _, ok := d.GetOk(prefix + "namespaces"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "namespaces"))
		i := 0
		val.Namespaces = []string{}
		for {
			pfx := fmt.Sprintf("%s.%s.%d", prefix, "namespaces", i)
			if v, ok := d.GetOk(pfx); ok {
				tflog.Debug(ctx, fmt.Sprintf("key is set: %s", pfx))
				val.Namespaces = append(val.Namespaces, v.(string))
				i++
			} else {
				tflog.Debug(ctx, "not set: %s", pfx)
				break
			}
		}
	}

	// TopologyKey: topology_key -- string, string
	if v, ok := d.GetOk(prefix + "topology_key"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "topology_key"))
		val.TopologyKey = v.(string)
	}

	// NamespaceSelector: namespace_selector -- , *config.LabelSelector

	if _, ok := d.GetOk(prefix + "namespace_selector.0"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "namespace_selector"))
		thing, err := dsRunnerConfigReadStructConfigLabelSelector(ctx, prefix+"namespace_selector.0", d)
		if err != nil {
			return val, err
		}
		val.NamespaceSelector = &thing
	} else {
		tflog.Debug(ctx, fmt.Sprintf("not set: %s", prefix+"namespace_selector"))
	}

	return val, nil
}

// config.PreferredSchedulingTerm

func dsRunnerConfigReadStructConfigPreferredSchedulingTerm(ctx context.Context, prefix string, d *schema.ResourceData) (config.PreferredSchedulingTerm, error) {

	if prefix != "" {
		prefix = prefix + "."
	}

	tflog.Debug(ctx, fmt.Sprintf("beginning dsRunnerConfigReadStructConfigPreferredSchedulingTerm run; prefix is '%s'", prefix))

	val := config.PreferredSchedulingTerm{}

	// Weight: weight -- int32, int32
	if v, ok := d.GetOk(prefix + "weight"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "weight"))
		val.Weight = v.(int32)
	}

	// Preference: preference -- NodeSelectorTerm, config.NodeSelectorTerm

	if _, ok := d.GetOk(prefix + "preference.0"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "preference"))
		thing, err := dsRunnerConfigReadStructConfigNodeSelectorTerm(ctx, prefix+"preference.0", d)
		if err != nil {
			return val, err
		}
		val.Preference = thing
	}

	return val, nil
}

// config.RunnerConfig

func dsRunnerConfigReadStructConfigRunnerConfig(ctx context.Context, prefix string, d *schema.ResourceData) (config.RunnerConfig, error) {

	if prefix != "" {
		prefix = prefix + "."
	}

	tflog.Debug(ctx, fmt.Sprintf("beginning dsRunnerConfigReadStructConfigRunnerConfig run; prefix is '%s'", prefix))

	val := config.RunnerConfig{}

	// Name: name -- string, string
	if v, ok := d.GetOk(prefix + "name"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "name"))
		val.Name = v.(string)
	}

	// Limit: limit -- int, int
	if v, ok := d.GetOk(prefix + "limit"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "limit"))
		val.Limit = v.(int)
	}

	// OutputLimit: output_limit -- int, int
	if v, ok := d.GetOk(prefix + "output_limit"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "output_limit"))
		val.OutputLimit = v.(int)
	}

	// RequestConcurrency: request_concurrency -- int, int
	if v, ok := d.GetOk(prefix + "request_concurrency"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "request_concurrency"))
		val.RequestConcurrency = v.(int)
	}

	// URL: url -- string, string
	if v, ok := d.GetOk(prefix + "url"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "url"))
		val.URL = v.(string)
	}

	// Token: token -- string, string
	if v, ok := d.GetOk(prefix + "token"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "token"))
		val.Token = v.(string)
	}

	// TLSCAFile: tls_ca_file -- string, string
	if v, ok := d.GetOk(prefix + "tls_ca_file"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "tls_ca_file"))
		val.TLSCAFile = v.(string)
	}

	// TLSCertFile: tls_cert_file -- string, string
	if v, ok := d.GetOk(prefix + "tls_cert_file"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "tls_cert_file"))
		val.TLSCertFile = v.(string)
	}

	// TLSKeyFile: tls_key_file -- string, string
	if v, ok := d.GetOk(prefix + "tls_key_file"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "tls_key_file"))
		val.TLSKeyFile = v.(string)
	}

	// Executor: executor -- string, string
	if v, ok := d.GetOk(prefix + "executor"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "executor"))
		val.Executor = v.(string)
	}

	// BuildsDir: builds_dir -- string, string
	if v, ok := d.GetOk(prefix + "builds_dir"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "builds_dir"))
		val.BuildsDir = v.(string)
	}

	// CacheDir: cache_dir -- string, string
	if v, ok := d.GetOk(prefix + "cache_dir"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "cache_dir"))
		val.CacheDir = v.(string)
	}

	// CloneURL: clone_url -- string, string
	if v, ok := d.GetOk(prefix + "clone_url"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "clone_url"))
		val.CloneURL = v.(string)
	}

	// Environment: environment -- , []string

	// HERE
	if _, ok := d.GetOk(prefix + "environment"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "environment"))
		i := 0
		val.Environment = []string{}
		for {
			pfx := fmt.Sprintf("%s.%s.%d", prefix, "environment", i)
			if v, ok := d.GetOk(pfx); ok {
				tflog.Debug(ctx, fmt.Sprintf("key is set: %s", pfx))
				val.Environment = append(val.Environment, v.(string))
				i++
			} else {
				tflog.Debug(ctx, "not set: %s", pfx)
				break
			}
		}
	}

	// PreCloneScript: pre_clone_script -- string, string
	if v, ok := d.GetOk(prefix + "pre_clone_script"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "pre_clone_script"))
		val.PreCloneScript = v.(string)
	}

	// PreBuildScript: pre_build_script -- string, string
	if v, ok := d.GetOk(prefix + "pre_build_script"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "pre_build_script"))
		val.PreBuildScript = v.(string)
	}

	// PostBuildScript: post_build_script -- string, string
	if v, ok := d.GetOk(prefix + "post_build_script"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "post_build_script"))
		val.PostBuildScript = v.(string)
	}

	// DebugTraceDisabled: debug_trace_disabled -- bool, bool
	if v, ok := d.GetOk(prefix + "debug_trace_disabled"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "debug_trace_disabled"))
		val.DebugTraceDisabled = v.(bool)
	}

	// Shell: shell -- string, string
	if v, ok := d.GetOk(prefix + "shell"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "shell"))
		val.Shell = v.(string)
	}

	// CustomBuildDir: custom_build_dir -- , *config.CustomBuildDir

	if _, ok := d.GetOk(prefix + "custom_build_dir.0"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "custom_build_dir"))
		thing, err := dsRunnerConfigReadStructConfigCustomBuildDir(ctx, prefix+"custom_build_dir.0", d)
		if err != nil {
			return val, err
		}
		val.CustomBuildDir = &thing
	} else {
		tflog.Debug(ctx, fmt.Sprintf("not set: %s", prefix+"custom_build_dir"))
	}

	// Referees: referees -- Config, referees.Config

	if _, ok := d.GetOk(prefix + "referees.0"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "referees"))
		thing, err := dsRunnerConfigReadStructRefereesConfig(ctx, prefix+"referees.0", d)
		if err != nil {
			return val, err
		}
		val.Referees = thing
	}

	// Cache: cache -- , *config.CacheConfig

	if _, ok := d.GetOk(prefix + "cache.0"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "cache"))
		thing, err := dsRunnerConfigReadStructConfigCacheConfig(ctx, prefix+"cache.0", d)
		if err != nil {
			return val, err
		}
		val.Cache = &thing
	} else {
		tflog.Debug(ctx, fmt.Sprintf("not set: %s", prefix+"cache"))
	}

	// FeatureFlags: feature_flags -- , map[string]bool
	if v, ok := d.GetOk(prefix + "feature_flags"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "feature_flags"))
		val.FeatureFlags = v.(map[string]bool)
	}

	// SSH: ssh -- , *ssh.Config

	if _, ok := d.GetOk(prefix + "ssh.0"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "ssh"))
		thing, err := dsRunnerConfigReadStructSshConfig(ctx, prefix+"ssh.0", d)
		if err != nil {
			return val, err
		}
		val.SSH = &thing
	} else {
		tflog.Debug(ctx, fmt.Sprintf("not set: %s", prefix+"ssh"))
	}

	// Docker: docker -- , *config.DockerConfig

	if _, ok := d.GetOk(prefix + "docker.0"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "docker"))
		thing, err := dsRunnerConfigReadStructConfigDockerConfig(ctx, prefix+"docker.0", d)
		if err != nil {
			return val, err
		}
		val.Docker = &thing
	} else {
		tflog.Debug(ctx, fmt.Sprintf("not set: %s", prefix+"docker"))
	}

	// Parallels: parallels -- , *config.ParallelsConfig

	if _, ok := d.GetOk(prefix + "parallels.0"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "parallels"))
		thing, err := dsRunnerConfigReadStructConfigParallelsConfig(ctx, prefix+"parallels.0", d)
		if err != nil {
			return val, err
		}
		val.Parallels = &thing
	} else {
		tflog.Debug(ctx, fmt.Sprintf("not set: %s", prefix+"parallels"))
	}

	// VirtualBox: virtualbox -- , *config.VirtualBoxConfig

	if _, ok := d.GetOk(prefix + "virtualbox.0"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "virtualbox"))
		thing, err := dsRunnerConfigReadStructConfigVirtualBoxConfig(ctx, prefix+"virtualbox.0", d)
		if err != nil {
			return val, err
		}
		val.VirtualBox = &thing
	} else {
		tflog.Debug(ctx, fmt.Sprintf("not set: %s", prefix+"virtualbox"))
	}

	// Machine: machine -- , *config.DockerMachine

	if _, ok := d.GetOk(prefix + "machine.0"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "machine"))
		thing, err := dsRunnerConfigReadStructConfigDockerMachine(ctx, prefix+"machine.0", d)
		if err != nil {
			return val, err
		}
		val.Machine = &thing
	} else {
		tflog.Debug(ctx, fmt.Sprintf("not set: %s", prefix+"machine"))
	}

	// Kubernetes: kubernetes -- , *config.KubernetesConfig

	if _, ok := d.GetOk(prefix + "kubernetes.0"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "kubernetes"))
		thing, err := dsRunnerConfigReadStructConfigKubernetesConfig(ctx, prefix+"kubernetes.0", d)
		if err != nil {
			return val, err
		}
		val.Kubernetes = &thing
	} else {
		tflog.Debug(ctx, fmt.Sprintf("not set: %s", prefix+"kubernetes"))
	}

	// Custom: custom -- , *config.CustomConfig

	if _, ok := d.GetOk(prefix + "custom.0"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "custom"))
		thing, err := dsRunnerConfigReadStructConfigCustomConfig(ctx, prefix+"custom.0", d)
		if err != nil {
			return val, err
		}
		val.Custom = &thing
	} else {
		tflog.Debug(ctx, fmt.Sprintf("not set: %s", prefix+"custom"))
	}

	return val, nil
}

// config.RunnerCredentials

func dsRunnerConfigReadStructConfigRunnerCredentials(ctx context.Context, prefix string, d *schema.ResourceData) (config.RunnerCredentials, error) {

	if prefix != "" {
		prefix = prefix + "."
	}

	tflog.Debug(ctx, fmt.Sprintf("beginning dsRunnerConfigReadStructConfigRunnerCredentials run; prefix is '%s'", prefix))

	val := config.RunnerCredentials{}

	// URL: url -- string, string
	if v, ok := d.GetOk(prefix + "url"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "url"))
		val.URL = v.(string)
	}

	// Token: token -- string, string
	if v, ok := d.GetOk(prefix + "token"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "token"))
		val.Token = v.(string)
	}

	// TLSCAFile: tls_ca_file -- string, string
	if v, ok := d.GetOk(prefix + "tls_ca_file"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "tls_ca_file"))
		val.TLSCAFile = v.(string)
	}

	// TLSCertFile: tls_cert_file -- string, string
	if v, ok := d.GetOk(prefix + "tls_cert_file"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "tls_cert_file"))
		val.TLSCertFile = v.(string)
	}

	// TLSKeyFile: tls_key_file -- string, string
	if v, ok := d.GetOk(prefix + "tls_key_file"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "tls_key_file"))
		val.TLSKeyFile = v.(string)
	}

	return val, nil
}

// config.RunnerSettings

func dsRunnerConfigReadStructConfigRunnerSettings(ctx context.Context, prefix string, d *schema.ResourceData) (config.RunnerSettings, error) {

	if prefix != "" {
		prefix = prefix + "."
	}

	tflog.Debug(ctx, fmt.Sprintf("beginning dsRunnerConfigReadStructConfigRunnerSettings run; prefix is '%s'", prefix))

	val := config.RunnerSettings{}

	// Executor: executor -- string, string
	if v, ok := d.GetOk(prefix + "executor"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "executor"))
		val.Executor = v.(string)
	}

	// BuildsDir: builds_dir -- string, string
	if v, ok := d.GetOk(prefix + "builds_dir"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "builds_dir"))
		val.BuildsDir = v.(string)
	}

	// CacheDir: cache_dir -- string, string
	if v, ok := d.GetOk(prefix + "cache_dir"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "cache_dir"))
		val.CacheDir = v.(string)
	}

	// CloneURL: clone_url -- string, string
	if v, ok := d.GetOk(prefix + "clone_url"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "clone_url"))
		val.CloneURL = v.(string)
	}

	// Environment: environment -- , []string

	// HERE
	if _, ok := d.GetOk(prefix + "environment"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "environment"))
		i := 0
		val.Environment = []string{}
		for {
			pfx := fmt.Sprintf("%s.%s.%d", prefix, "environment", i)
			if v, ok := d.GetOk(pfx); ok {
				tflog.Debug(ctx, fmt.Sprintf("key is set: %s", pfx))
				val.Environment = append(val.Environment, v.(string))
				i++
			} else {
				tflog.Debug(ctx, "not set: %s", pfx)
				break
			}
		}
	}

	// PreCloneScript: pre_clone_script -- string, string
	if v, ok := d.GetOk(prefix + "pre_clone_script"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "pre_clone_script"))
		val.PreCloneScript = v.(string)
	}

	// PreBuildScript: pre_build_script -- string, string
	if v, ok := d.GetOk(prefix + "pre_build_script"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "pre_build_script"))
		val.PreBuildScript = v.(string)
	}

	// PostBuildScript: post_build_script -- string, string
	if v, ok := d.GetOk(prefix + "post_build_script"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "post_build_script"))
		val.PostBuildScript = v.(string)
	}

	// DebugTraceDisabled: debug_trace_disabled -- bool, bool
	if v, ok := d.GetOk(prefix + "debug_trace_disabled"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "debug_trace_disabled"))
		val.DebugTraceDisabled = v.(bool)
	}

	// Shell: shell -- string, string
	if v, ok := d.GetOk(prefix + "shell"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "shell"))
		val.Shell = v.(string)
	}

	// CustomBuildDir: custom_build_dir -- , *config.CustomBuildDir

	if _, ok := d.GetOk(prefix + "custom_build_dir.0"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "custom_build_dir"))
		thing, err := dsRunnerConfigReadStructConfigCustomBuildDir(ctx, prefix+"custom_build_dir.0", d)
		if err != nil {
			return val, err
		}
		val.CustomBuildDir = &thing
	} else {
		tflog.Debug(ctx, fmt.Sprintf("not set: %s", prefix+"custom_build_dir"))
	}

	// Referees: referees -- Config, referees.Config

	if _, ok := d.GetOk(prefix + "referees.0"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "referees"))
		thing, err := dsRunnerConfigReadStructRefereesConfig(ctx, prefix+"referees.0", d)
		if err != nil {
			return val, err
		}
		val.Referees = thing
	}

	// Cache: cache -- , *config.CacheConfig

	if _, ok := d.GetOk(prefix + "cache.0"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "cache"))
		thing, err := dsRunnerConfigReadStructConfigCacheConfig(ctx, prefix+"cache.0", d)
		if err != nil {
			return val, err
		}
		val.Cache = &thing
	} else {
		tflog.Debug(ctx, fmt.Sprintf("not set: %s", prefix+"cache"))
	}

	// FeatureFlags: feature_flags -- , map[string]bool
	if v, ok := d.GetOk(prefix + "feature_flags"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "feature_flags"))
		val.FeatureFlags = v.(map[string]bool)
	}

	// SSH: ssh -- , *ssh.Config

	if _, ok := d.GetOk(prefix + "ssh.0"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "ssh"))
		thing, err := dsRunnerConfigReadStructSshConfig(ctx, prefix+"ssh.0", d)
		if err != nil {
			return val, err
		}
		val.SSH = &thing
	} else {
		tflog.Debug(ctx, fmt.Sprintf("not set: %s", prefix+"ssh"))
	}

	// Docker: docker -- , *config.DockerConfig

	if _, ok := d.GetOk(prefix + "docker.0"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "docker"))
		thing, err := dsRunnerConfigReadStructConfigDockerConfig(ctx, prefix+"docker.0", d)
		if err != nil {
			return val, err
		}
		val.Docker = &thing
	} else {
		tflog.Debug(ctx, fmt.Sprintf("not set: %s", prefix+"docker"))
	}

	// Parallels: parallels -- , *config.ParallelsConfig

	if _, ok := d.GetOk(prefix + "parallels.0"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "parallels"))
		thing, err := dsRunnerConfigReadStructConfigParallelsConfig(ctx, prefix+"parallels.0", d)
		if err != nil {
			return val, err
		}
		val.Parallels = &thing
	} else {
		tflog.Debug(ctx, fmt.Sprintf("not set: %s", prefix+"parallels"))
	}

	// VirtualBox: virtualbox -- , *config.VirtualBoxConfig

	if _, ok := d.GetOk(prefix + "virtualbox.0"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "virtualbox"))
		thing, err := dsRunnerConfigReadStructConfigVirtualBoxConfig(ctx, prefix+"virtualbox.0", d)
		if err != nil {
			return val, err
		}
		val.VirtualBox = &thing
	} else {
		tflog.Debug(ctx, fmt.Sprintf("not set: %s", prefix+"virtualbox"))
	}

	// Machine: machine -- , *config.DockerMachine

	if _, ok := d.GetOk(prefix + "machine.0"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "machine"))
		thing, err := dsRunnerConfigReadStructConfigDockerMachine(ctx, prefix+"machine.0", d)
		if err != nil {
			return val, err
		}
		val.Machine = &thing
	} else {
		tflog.Debug(ctx, fmt.Sprintf("not set: %s", prefix+"machine"))
	}

	// Kubernetes: kubernetes -- , *config.KubernetesConfig

	if _, ok := d.GetOk(prefix + "kubernetes.0"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "kubernetes"))
		thing, err := dsRunnerConfigReadStructConfigKubernetesConfig(ctx, prefix+"kubernetes.0", d)
		if err != nil {
			return val, err
		}
		val.Kubernetes = &thing
	} else {
		tflog.Debug(ctx, fmt.Sprintf("not set: %s", prefix+"kubernetes"))
	}

	// Custom: custom -- , *config.CustomConfig

	if _, ok := d.GetOk(prefix + "custom.0"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "custom"))
		thing, err := dsRunnerConfigReadStructConfigCustomConfig(ctx, prefix+"custom.0", d)
		if err != nil {
			return val, err
		}
		val.Custom = &thing
	} else {
		tflog.Debug(ctx, fmt.Sprintf("not set: %s", prefix+"custom"))
	}

	return val, nil
}

// config.Service

func dsRunnerConfigReadStructConfigService(ctx context.Context, prefix string, d *schema.ResourceData) (config.Service, error) {

	if prefix != "" {
		prefix = prefix + "."
	}

	tflog.Debug(ctx, fmt.Sprintf("beginning dsRunnerConfigReadStructConfigService run; prefix is '%s'", prefix))

	val := config.Service{}

	// Name: name -- string, string
	if v, ok := d.GetOk(prefix + "name"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "name"))
		val.Name = v.(string)
	}

	// Alias: alias -- string, string
	if v, ok := d.GetOk(prefix + "alias"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "alias"))
		val.Alias = v.(string)
	}

	// Command: command -- , []string

	// HERE
	if _, ok := d.GetOk(prefix + "command"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "command"))
		i := 0
		val.Command = []string{}
		for {
			pfx := fmt.Sprintf("%s.%s.%d", prefix, "command", i)
			if v, ok := d.GetOk(pfx); ok {
				tflog.Debug(ctx, fmt.Sprintf("key is set: %s", pfx))
				val.Command = append(val.Command, v.(string))
				i++
			} else {
				tflog.Debug(ctx, "not set: %s", pfx)
				break
			}
		}
	}

	// Entrypoint: entrypoint -- , []string

	// HERE
	if _, ok := d.GetOk(prefix + "entrypoint"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "entrypoint"))
		i := 0
		val.Entrypoint = []string{}
		for {
			pfx := fmt.Sprintf("%s.%s.%d", prefix, "entrypoint", i)
			if v, ok := d.GetOk(pfx); ok {
				tflog.Debug(ctx, fmt.Sprintf("key is set: %s", pfx))
				val.Entrypoint = append(val.Entrypoint, v.(string))
				i++
			} else {
				tflog.Debug(ctx, "not set: %s", pfx)
				break
			}
		}
	}

	return val, nil
}

// config.SessionServer

func dsRunnerConfigReadStructConfigSessionServer(ctx context.Context, prefix string, d *schema.ResourceData) (config.SessionServer, error) {

	if prefix != "" {
		prefix = prefix + "."
	}

	tflog.Debug(ctx, fmt.Sprintf("beginning dsRunnerConfigReadStructConfigSessionServer run; prefix is '%s'", prefix))

	val := config.SessionServer{}

	// ListenAddress: listen_address -- string, string
	if v, ok := d.GetOk(prefix + "listen_address"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "listen_address"))
		val.ListenAddress = v.(string)
	}

	// AdvertiseAddress: advertise_address -- string, string
	if v, ok := d.GetOk(prefix + "advertise_address"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "advertise_address"))
		val.AdvertiseAddress = v.(string)
	}

	// SessionTimeout: session_timeout -- int, int
	if v, ok := d.GetOk(prefix + "session_timeout"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "session_timeout"))
		val.SessionTimeout = v.(int)
	}

	return val, nil
}

// config.VirtualBoxConfig

func dsRunnerConfigReadStructConfigVirtualBoxConfig(ctx context.Context, prefix string, d *schema.ResourceData) (config.VirtualBoxConfig, error) {

	if prefix != "" {
		prefix = prefix + "."
	}

	tflog.Debug(ctx, fmt.Sprintf("beginning dsRunnerConfigReadStructConfigVirtualBoxConfig run; prefix is '%s'", prefix))

	val := config.VirtualBoxConfig{}

	// BaseName: base_name -- string, string
	if v, ok := d.GetOk(prefix + "base_name"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "base_name"))
		val.BaseName = v.(string)
	}

	// BaseSnapshot: base_snapshot -- string, string
	if v, ok := d.GetOk(prefix + "base_snapshot"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "base_snapshot"))
		val.BaseSnapshot = v.(string)
	}

	// BaseFolder: base_folder -- string, string
	if v, ok := d.GetOk(prefix + "base_folder"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "base_folder"))
		val.BaseFolder = v.(string)
	}

	// DisableSnapshots: disable_snapshots -- bool, bool
	if v, ok := d.GetOk(prefix + "disable_snapshots"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "disable_snapshots"))
		val.DisableSnapshots = v.(bool)
	}

	// AllowedImages: allowed_images -- , []string

	// HERE
	if _, ok := d.GetOk(prefix + "allowed_images"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "allowed_images"))
		i := 0
		val.AllowedImages = []string{}
		for {
			pfx := fmt.Sprintf("%s.%s.%d", prefix, "allowed_images", i)
			if v, ok := d.GetOk(pfx); ok {
				tflog.Debug(ctx, fmt.Sprintf("key is set: %s", pfx))
				val.AllowedImages = append(val.AllowedImages, v.(string))
				i++
			} else {
				tflog.Debug(ctx, "not set: %s", pfx)
				break
			}
		}
	}

	return val, nil
}

// config.WeightedPodAffinityTerm

func dsRunnerConfigReadStructConfigWeightedPodAffinityTerm(ctx context.Context, prefix string, d *schema.ResourceData) (config.WeightedPodAffinityTerm, error) {

	if prefix != "" {
		prefix = prefix + "."
	}

	tflog.Debug(ctx, fmt.Sprintf("beginning dsRunnerConfigReadStructConfigWeightedPodAffinityTerm run; prefix is '%s'", prefix))

	val := config.WeightedPodAffinityTerm{}

	// Weight: weight -- int32, int32
	if v, ok := d.GetOk(prefix + "weight"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "weight"))
		val.Weight = v.(int32)
	}

	// PodAffinityTerm: pod_affinity_term -- PodAffinityTerm, config.PodAffinityTerm

	if _, ok := d.GetOk(prefix + "pod_affinity_term.0"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "pod_affinity_term"))
		thing, err := dsRunnerConfigReadStructConfigPodAffinityTerm(ctx, prefix+"pod_affinity_term.0", d)
		if err != nil {
			return val, err
		}
		val.PodAffinityTerm = thing
	}

	return val, nil
}

// docker.Credentials

func dsRunnerConfigReadStructDockerCredentials(ctx context.Context, prefix string, d *schema.ResourceData) (docker.Credentials, error) {

	if prefix != "" {
		prefix = prefix + "."
	}

	tflog.Debug(ctx, fmt.Sprintf("beginning dsRunnerConfigReadStructDockerCredentials run; prefix is '%s'", prefix))

	val := docker.Credentials{}

	// Host: host -- string, string
	if v, ok := d.GetOk(prefix + "host"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "host"))
		val.Host = v.(string)
	}

	// CertPath: tls_cert_path -- string, string
	if v, ok := d.GetOk(prefix + "tls_cert_path"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "tls_cert_path"))
		val.CertPath = v.(string)
	}

	// TLSVerify: tls_verify -- bool, bool
	if v, ok := d.GetOk(prefix + "tls_verify"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "tls_verify"))
		val.TLSVerify = v.(bool)
	}

	return val, nil
}

// referees.Config

func dsRunnerConfigReadStructRefereesConfig(ctx context.Context, prefix string, d *schema.ResourceData) (referees.Config, error) {

	if prefix != "" {
		prefix = prefix + "."
	}

	tflog.Debug(ctx, fmt.Sprintf("beginning dsRunnerConfigReadStructRefereesConfig run; prefix is '%s'", prefix))

	val := referees.Config{}

	// Metrics: metrics -- , *referees.MetricsRefereeConfig

	if _, ok := d.GetOk(prefix + "metrics.0"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "metrics"))
		thing, err := dsRunnerConfigReadStructRefereesMetricsRefereeConfig(ctx, prefix+"metrics.0", d)
		if err != nil {
			return val, err
		}
		val.Metrics = &thing
	} else {
		tflog.Debug(ctx, fmt.Sprintf("not set: %s", prefix+"metrics"))
	}

	return val, nil
}

// referees.MetricsRefereeConfig

func dsRunnerConfigReadStructRefereesMetricsRefereeConfig(ctx context.Context, prefix string, d *schema.ResourceData) (referees.MetricsRefereeConfig, error) {

	if prefix != "" {
		prefix = prefix + "."
	}

	tflog.Debug(ctx, fmt.Sprintf("beginning dsRunnerConfigReadStructRefereesMetricsRefereeConfig run; prefix is '%s'", prefix))

	val := referees.MetricsRefereeConfig{}

	// PrometheusAddress: prometheus_address -- string, string
	if v, ok := d.GetOk(prefix + "prometheus_address"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "prometheus_address"))
		val.PrometheusAddress = v.(string)
	}

	// QueryInterval: query_interval -- int, int
	if v, ok := d.GetOk(prefix + "query_interval"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "query_interval"))
		val.QueryInterval = v.(int)
	}

	// Queries: queries -- , []string

	// HERE
	if _, ok := d.GetOk(prefix + "queries"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "queries"))
		i := 0
		val.Queries = []string{}
		for {
			pfx := fmt.Sprintf("%s.%s.%d", prefix, "queries", i)
			if v, ok := d.GetOk(pfx); ok {
				tflog.Debug(ctx, fmt.Sprintf("key is set: %s", pfx))
				val.Queries = append(val.Queries, v.(string))
				i++
			} else {
				tflog.Debug(ctx, "not set: %s", pfx)
				break
			}
		}
	}

	return val, nil
}

// ssh.Config

func dsRunnerConfigReadStructSshConfig(ctx context.Context, prefix string, d *schema.ResourceData) (ssh.Config, error) {

	if prefix != "" {
		prefix = prefix + "."
	}

	tflog.Debug(ctx, fmt.Sprintf("beginning dsRunnerConfigReadStructSshConfig run; prefix is '%s'", prefix))

	val := ssh.Config{}

	// User: user -- string, string
	if v, ok := d.GetOk(prefix + "user"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "user"))
		val.User = v.(string)
	}

	// Password: password -- string, string
	if v, ok := d.GetOk(prefix + "password"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "password"))
		val.Password = v.(string)
	}

	// Host: host -- string, string
	if v, ok := d.GetOk(prefix + "host"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "host"))
		val.Host = v.(string)
	}

	// Port: port -- string, string
	if v, ok := d.GetOk(prefix + "port"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "port"))
		val.Port = v.(string)
	}

	// IdentityFile: identity_file -- string, string
	if v, ok := d.GetOk(prefix + "identity_file"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "identity_file"))
		val.IdentityFile = v.(string)
	}

	// DisableStrictHostKeyChecking: disable_strict_host_key_checking -- , *bool
	if v, ok := d.GetOk(prefix + "disable_strict_host_key_checking"); ok {
		tflog.Debug(ctx, "set: %s.%s", prefix, "disable_strict_host_key_checking")
		val.DisableStrictHostKeyChecking = to.BoolP(v.(bool))
	}

	// KnownHostsFile: known_hosts_file -- string, string
	if v, ok := d.GetOk(prefix + "known_hosts_file"); ok {
		tflog.Debug(ctx, fmt.Sprintf("set: %s%s", prefix, "known_hosts_file"))
		val.KnownHostsFile = v.(string)
	}

	return val, nil
}
