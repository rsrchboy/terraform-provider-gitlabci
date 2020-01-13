package gitlabci

import (
	"bytes"
	"fmt"
	"log"
	"strconv"

	"github.com/BurntSushi/toml"
	"github.com/davecgh/go-spew/spew"
	"github.com/giantswarm/to"
	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/rsrchboy/structs"
	rcommon "gitlab.com/gitlab-org/gitlab-runner/common"
	rdhelpers "gitlab.com/gitlab-org/gitlab-runner/helpers/docker"
	"gitlab.com/gitlab-org/gitlab-runner/helpers/ssh"
	"gitlab.com/gitlab-org/gitlab-runner/referees"
	"gitlab.com/rsrchboy/terraform-provider-gitlabci/helper/runner"
)

type schemaMap map[string]*schema.Schema
type iMap map[string]interface{}
type stringMap map[string]string

func dataSourceGitlabCIRunnerConfig() *schema.Resource {

	log.SetFlags(log.Lshortfile)

	structs.DefaultTagName = "toml"

	optionalString := &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
	}

	schema := &schema.Resource{
		Read: dataSourceGitlabCIRunnerConfigRead,

		Schema: map[string]*schema.Schema{
			// generated
			"config": {
				Type:     schema.TypeString,
				Computed: true,
			},
			// input
			"concurrent": {
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      1,
				ValidateFunc: validation.IntAtLeast(1),
			},
			"check_interval": {
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      0,
				ValidateFunc: validation.IntAtLeast(0),
			},
			"log_level": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"panic", "fatal", "error", "warning", "info", "debug"}, false),
			},
			"log_format": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"runner", "text", "json"}, false),
			},
			"session_server": {
				Type:     schema.TypeList,
				Optional: true,
				MinItems: 0,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"listen_address":    optionalString,
						"advertise_address": optionalString,
						"session_timeout":   {Type: schema.TypeInt, Optional: true},
					},
				},
			},
			"runners": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Resource{Schema: runner.RunnerConfigToTerraformSchema()},
			},
		},
	}

	// log.Printf("[TRACE] generated schema is: %s", spew.Sdump(schema))
	return schema
}

func dataSourceGitlabCIRunnerConfigRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[TRACE] dataSourceGitlabCIRunnerConfigRead() (mark III)")

	c := rcommon.NewConfig()

	c.CheckInterval = d.Get("check_interval").(int)
	c.Concurrent = d.Get("concurrent").(int)
	c.LogFormat = to.StringP(d.Get("log_format").(string))
	c.LogLevel = to.StringP(d.Get("log_level").(string))

	if sessionServers, _ := d.GetOk("session_server"); len(sessionServers.([]interface{})) > 0 {
		sessionServer := sessionServers.([]interface{})[0]
		log.Printf("hasSessionServer: %T, %#v", sessionServer, sessionServer)
		log.Printf("hasSessionServer: %s", spew.Sdump(sessionServer))
		c.SessionServer.ListenAddress = d.Get("session_server.0.listen_address").(string)
		c.SessionServer.AdvertiseAddress = d.Get("session_server.0.advertise_address").(string)
		c.SessionServer.SessionTimeout = d.Get("session_server.0.session_timeout").(int)
	}

	if runners, hasRunners := d.GetOk("runners"); hasRunners {
		c.Runners = make([]*rcommon.RunnerConfig, len(runners.([]interface{})))
		for i, _ := range runners.([]interface{}) {
			log.Printf("working on runner block #%d", i)
			pfx := fmt.Sprintf("runners.%d.", i)
			log.Printf("i is %d, pfx is %s", i, pfx)
			r := rcommon.RunnerConfig{
				Name:               d.Get(pfx + "name").(string),
				Limit:              d.Get(pfx + "limit").(int),
				OutputLimit:        d.Get(pfx + "output_limit").(int),
				RequestConcurrency: d.Get(pfx + "request_concurrency").(int),
				RunnerCredentials: rcommon.RunnerCredentials{
					URL:         d.Get(pfx + "url").(string),
					Token:       d.Get(pfx + "token").(string),
					TLSCAFile:   d.Get(pfx + "tls_ca_file").(string),
					TLSCertFile: d.Get(pfx + "tls_cert_file").(string),
					TLSKeyFile:  d.Get(pfx + "tls_key_file").(string),
				},
				RunnerSettings: rcommon.RunnerSettings{
					Executor:           d.Get(pfx + "executor").(string),
					BuildsDir:          d.Get(pfx + "builds_dir").(string),
					CacheDir:           d.Get(pfx + "cache_dir").(string),
					CloneURL:           d.Get(pfx + "clone_url").(string),
					Environment:        stringList(pfx+"environment", d),
					PreCloneScript:     d.Get(pfx + "pre_clone_script").(string),
					PreBuildScript:     d.Get(pfx + "pre_build_script").(string),
					PostBuildScript:    d.Get(pfx + "post_build_script").(string),
					DebugTraceDisabled: d.Get(pfx + "debug_trace_disabled").(bool),
					Shell:              d.Get(pfx + "shell").(string),
					CustomBuildDir:     customBuildDirStructs(pfx, d),
					SSH:                sshStructs(pfx, d),
					Docker:             dockerConfigStructs(pfx, d),
					Parallels:          parallelsStructs(pfx, d),
					VirtualBox:         virtualBoxStructs(pfx, d),
					Cache:              cacheConfigStructs(pfx, d),
					Machine:            dockerMachineStructs(pfx, d), // d.Get(pfx + "machine").(*DockerMachine),
					Kubernetes:         k8sStructs(pfx, d),
					Custom:             customStructs(pfx, d),
					Referees:           refereeConfigStructs(pfx, d),
				},
			}

			c.Runners[i] = &r
		}

	}

	// generate toml config
	buf := new(bytes.Buffer)
	if err := toml.NewEncoder(buf).Encode(c); err != nil {
		return err
	}

	d.Set("config", fmt.Sprintf("%s", buf.String()))
	log.Printf("[INFO] runner c:\n\n%s", buf.String())

	configString := buf.String()

	d.SetId(strconv.Itoa(hashcode.String(configString)))

	log.Printf("[TRACE] dataSourceGitlabCIRunnerConfigRead() finished")
	return nil
}

func k8sStructs(prefix string, d *schema.ResourceData) *rcommon.KubernetesConfig {
	key := "kubernetes"
	pfx := prefix + key + ".0."
	if _, hasIt := d.GetOkExists(prefix + "kubernetes"); !hasIt {
		// if we don't have this key, we don't have any of it.  return!
		return nil
	}

	k := rcommon.KubernetesConfig{
		Host:                        d.Get(pfx + "host").(string),
		CertFile:                    d.Get(pfx + "cert_file").(string),
		KeyFile:                     d.Get(pfx + "key_file").(string),
		CAFile:                      d.Get(pfx + "ca_file").(string),
		BearerTokenOverwriteAllowed: d.Get(pfx + "bearer_token_overwrite_allowed").(bool),
		BearerToken:                 d.Get(pfx + "bearer_token").(string),
		Image:                       d.Get(pfx + "image").(string),
		Namespace:                   d.Get(pfx + "namespace").(string),
		NamespaceOverwriteAllowed:   d.Get(pfx + "namespace_overwrite_allowed").(string),
		Privileged:                  d.Get(pfx + "privileged").(bool),
		CPULimit:                    d.Get(pfx + "cpu_limit").(string),
		MemoryLimit:                 d.Get(pfx + "memory_limit").(string),
		ServiceCPULimit:             d.Get(pfx + "service_cpu_limit").(string),
		ServiceMemoryLimit:          d.Get(pfx + "service_memory_limit").(string),
		HelperCPULimit:              d.Get(pfx + "helper_cpu_limit").(string),
		HelperMemoryLimit:           d.Get(pfx + "helper_memory_limit").(string),
		CPURequest:                  d.Get(pfx + "cpu_request").(string),
		MemoryRequest:               d.Get(pfx + "memory_request").(string),
		ServiceCPURequest:           d.Get(pfx + "service_cpu_request").(string),
		ServiceMemoryRequest:        d.Get(pfx + "service_memory_request").(string),
		HelperCPURequest:            d.Get(pfx + "helper_cpu_request").(string),
		HelperMemoryRequest:         d.Get(pfx + "helper_memory_request").(string),
		// PullPolicy:
		NodeSelector:                   toStringMap(pfx+"node_selector", d),
		NodeTolerations:                toStringMap(pfx+"node_tolerations", d),
		ImagePullSecrets:               stringList(pfx+"image_pull_secrets", d),
		HelperImage:                    d.Get(pfx + "helper_image").(string),
		TerminationGracePeriodSeconds:  int64(d.Get(pfx + "termination_grace_period_seconds").(int)),
		PollInterval:                   d.Get(pfx + "poll_interval").(int),
		PollTimeout:                    d.Get(pfx + "poll_timeout").(int),
		PodLabels:                      toStringMap(pfx+"pod_labels", d),
		ServiceAccount:                 d.Get(pfx + "service_account").(string),
		ServiceAccountOverwriteAllowed: d.Get(pfx + "service_account_overwrite_allowed").(string),
		PodAnnotations:                 toStringMap(pfx+"pod_annotations", d),
		PodAnnotationsOverwriteAllowed: d.Get(pfx + "pod_annotations_overwrite_allowed").(string),
		// PodSecurityContext:             d.Get(pfx + "pod_security_context").(KubernetesPodSecurityContext),
		// Volumes:                        d.Get(pfx + "volumes").(KubernetesVolumes),
		// Services:                       d.Get(pfx + "services").([]Service),
	}

	switch d.Get(pfx + "pull_policy").(string) {
	case "":
		k.PullPolicy = ""
	case "Always":
		k.PullPolicy = rcommon.PullPolicyAlways
	case "Never":
		k.PullPolicy = rcommon.PullPolicyNever
	case "IfNotPresent":
		k.PullPolicy = rcommon.PullPolicyIfNotPresent
	}

	if _, hasIt := d.GetOkExists(pfx + "volumes"); hasIt {
		pfx := pfx + "volumes.0."
		k.Volumes = rcommon.KubernetesVolumes{
			// HostPaths  []KubernetesHostPath  `toml:"host_path" description:"The host paths which will be mounted" json:"host_paths" tf:"host_path"`
			// PVCs       []KubernetesPVC       `toml:"pvc" description:"The persistent volume claims that will be mounted" json:"pv_cs" tf:"pvc"`
			// ConfigMaps []KubernetesConfigMap `toml:"config_map" description:"The config maps which will be mounted as volumes" json:"config_maps" tf:"config_map"`
			// Secrets    []KubernetesSecret    `toml:"secret" description:"The secret maps which will be mounted" json:"secrets" tf:"secret"`
			// EmptyDirs  []KubernetesEmptyDir  `toml:"empty_dir" description:"The empty dirs which will be mounted" json:"empty_dirs" tf:"empty_dir"
		}

		subKey := "host_path"
		if it, hasIt := d.GetOkExists(pfx + subKey); hasIt {
			vols := make([]rcommon.KubernetesHostPath, len(it.([]interface{})))
			for i, _ := range it.([]interface{}) {
				pfx := fmt.Sprintf("%s%s.%d.", pfx, subKey, i)
				vols[i] = rcommon.KubernetesHostPath{
					Name:      d.Get(pfx + "name").(string),
					MountPath: d.Get(pfx + "mount_path").(string),
					ReadOnly:  d.Get(pfx + "read_only").(bool),
					HostPath:  d.Get(pfx + "host_path").(string),
				}
			}
			k.Volumes.HostPaths = vols
		}

		subKey = "pvc"
		if it, hasIt := d.GetOkExists(pfx + subKey); hasIt {
			vols := make([]rcommon.KubernetesPVC, len(it.([]interface{})))
			for i, _ := range it.([]interface{}) {
				pfx := fmt.Sprintf("%s%s.%d.", pfx, subKey, i)
				vols[i] = rcommon.KubernetesPVC{
					Name:      d.Get(pfx + "name").(string),
					MountPath: d.Get(pfx + "mount_path").(string),
					ReadOnly:  d.Get(pfx + "read_only").(bool),
				}
			}
			k.Volumes.PVCs = vols
		}

		subKey = "config_map"
		if it, hasIt := d.GetOkExists(pfx + subKey); hasIt {
			vols := make([]rcommon.KubernetesConfigMap, len(it.([]interface{})))
			for i, _ := range it.([]interface{}) {
				pfx := fmt.Sprintf("%s%s.%d.", pfx, subKey, i)
				vols[i] = rcommon.KubernetesConfigMap{
					Name:      d.Get(pfx + "name").(string),
					MountPath: d.Get(pfx + "mount_path").(string),
					ReadOnly:  d.Get(pfx + "read_only").(bool),
					Items:     toStringMap(pfx+"items", d),
				}
			}
			k.Volumes.ConfigMaps = vols
		}

		subKey = "secret"
		if it, hasIt := d.GetOkExists(pfx + subKey); hasIt {
			vols := make([]rcommon.KubernetesSecret, len(it.([]interface{})))
			for i, _ := range it.([]interface{}) {
				pfx := fmt.Sprintf("%s%s.%d.", pfx, subKey, i)
				vols[i] = rcommon.KubernetesSecret{
					Name:      d.Get(pfx + "name").(string),
					MountPath: d.Get(pfx + "mount_path").(string),
					ReadOnly:  d.Get(pfx + "read_only").(bool),
					Items:     toStringMap(pfx+"items", d),
				}
			}
			k.Volumes.Secrets = vols
		}

		subKey = "empty_dir"
		if it, hasIt := d.GetOkExists(pfx + subKey); hasIt {
			vols := make([]rcommon.KubernetesEmptyDir, len(it.([]interface{})))
			for i, _ := range it.([]interface{}) {
				pfx := fmt.Sprintf("%s%s.%d.", pfx, subKey, i)
				vols[i] = rcommon.KubernetesEmptyDir{
					Name:      d.Get(pfx + "name").(string),
					MountPath: d.Get(pfx + "mount_path").(string),
					Medium:    d.Get(pfx + "medium").(string),
				}
			}
			k.Volumes.EmptyDirs = vols
		}

	}

	if _, hasIt := d.GetOkExists(pfx + "pod_security_context"); hasIt {
		pfx := pfx + "pod_security_context.0."
		k.PodSecurityContext = rcommon.KubernetesPodSecurityContext{
			FSGroup:            d.Get(pfx + "fs_group").(*int64),
			RunAsGroup:         d.Get(pfx + "run_as_group").(*int64),
			RunAsNonRoot:       d.Get(pfx + "run_as_non_root").(*bool),
			RunAsUser:          d.Get(pfx + "run_as_user").(*int64),
			SupplementalGroups: d.Get(pfx + "supplemental_groups").([]int64),
		}
	}

	if servicesI, hasIt := d.GetOkExists(pfx + "services"); hasIt {
		services := make([]rcommon.Service, len(servicesI.([]interface{})))
		for i, v := range servicesI.([]interface{}) {
			services[i] = rcommon.Service{
				Name: v.(string),
			}
			k.Services = services
		}
	}

	return &k
}

func customStructs(prefix string, d *schema.ResourceData) *rcommon.CustomConfig {
	key := "custom"
	pfx := prefix + key + ".0."
	if _, hasIt := d.GetOkExists(prefix + key); !hasIt {
		return nil
	}

	cc := rcommon.CustomConfig{
		ConfigExec:          d.Get(pfx + "config_exec").(string),
		ConfigArgs:          stringList(pfx+"config_args", d),
		ConfigExecTimeout:   to.IntP(d.Get(pfx + "config_exec_timeout").(int)),
		PrepareExec:         d.Get(pfx + "prepare_exec").(string),
		PrepareArgs:         stringList(pfx+"prepare_args", d),
		PrepareExecTimeout:  to.IntP(d.Get(pfx + "prepare_exec_timeout").(int)),
		RunExec:             d.Get(pfx + "run_exec").(string),
		RunArgs:             stringList(pfx+"run_args", d),
		CleanupExec:         d.Get(pfx + "cleanup_exec").(string),
		CleanupArgs:         stringList(pfx+"cleanup_args", d),
		CleanupExecTimeout:  to.IntP(d.Get(pfx + "cleanup_exec_timeout").(int)),
		GracefulKillTimeout: to.IntP(d.Get(pfx + "graceful_kill_timeout").(int)),
		ForceKillTimeout:    to.IntP(d.Get(pfx + "force_kill_timeout").(int)),
	}

	return &cc
}

func parallelsStructs(prefix string, d *schema.ResourceData) *rcommon.ParallelsConfig {
	key := "parallels"
	pfx := prefix + key + ".0."
	if _, hasIt := d.GetOkExists(prefix + key); !hasIt {
		return nil
	}

	p := rcommon.ParallelsConfig{
		BaseName:         d.Get(pfx + "base_name").(string),
		TemplateName:     d.Get(pfx + "template_name").(string),
		DisableSnapshots: d.Get(pfx + "disable_snapshots").(bool),
		TimeServer:       d.Get(pfx + "time_server").(string),
	}

	return &p
}

func virtualBoxStructs(prefix string, d *schema.ResourceData) *rcommon.VirtualBoxConfig {
	key := "virtual_box"
	pfx := prefix + key + ".0."
	if _, hasIt := d.GetOkExists(prefix + key); !hasIt {
		return nil
	}

	vb := rcommon.VirtualBoxConfig{
		BaseName:         d.Get(pfx + "base_name").(string),
		BaseSnapshot:     d.Get(pfx + "base_snapshot").(string),
		DisableSnapshots: d.Get(pfx + "disable_snapshots").(bool),
	}

	return &vb
}

func sshStructs(prefix string, d *schema.ResourceData) *ssh.Config {

	pfx := prefix + "ssh.0."
	if _, hasIt := d.GetOkExists(prefix + "ssh"); !hasIt {
		return nil
	}

	ssh := ssh.Config{
		User:         d.Get(pfx + "user").(string),
		Password:     d.Get(pfx + "password").(string),
		Host:         d.Get(pfx + "host").(string),
		Port:         d.Get(pfx + "port").(string),
		IdentityFile: d.Get(pfx + "identity_file").(string),
	}

	return &ssh
}

func customBuildDirStructs(prefix string, d *schema.ResourceData) *rcommon.CustomBuildDir {

	pfx := prefix + "custom_build_dir.0."
	if _, hasIt := d.GetOkExists(prefix + "custom_build_dir"); !hasIt {
		return nil
	}

	cbd := rcommon.CustomBuildDir{
		Enabled: d.Get(pfx + "enabled").(bool),
	}

	return &cbd
}

func cacheConfigStructs(prefix string, d *schema.ResourceData) *rcommon.CacheConfig {

	pfx := prefix + "cache.0."
	if _, hasIt := d.GetOkExists(prefix + "cache"); !hasIt {
		return nil
	}
	cache := rcommon.CacheConfig{
		Type:   d.Get(pfx + "type").(string),
		Path:   d.Get(pfx + "path").(string),
		Shared: d.Get(pfx + "shared").(bool),
		// S3:  d.Get(pfx + "s3").(*CacheS3Config),
		// GCS: d.Get(pfx + "gcs").(*CacheGCSConfig),
	}

	if _, hasIt := d.GetOkExists(pfx + "gcs"); hasIt {
		pfx := pfx + "gcs.0."
		cache.GCS = &rcommon.CacheGCSConfig{
			// CacheGCSCredentials `tf:"cache_gcs_credentials"`
			CredentialsFile: d.Get(pfx + "credentials_file").(string),
			BucketName:      d.Get(pfx + "bucket_name").(string),
		}

		if _, hasCreds := d.GetOkExists(pfx + "gcs"); hasCreds {
			pfx := pfx + "cache_gcs_credentials.0."
			cache.GCS.CacheGCSCredentials = rcommon.CacheGCSCredentials{
				AccessID:   d.Get(pfx + "access_id").(string),
				PrivateKey: d.Get(pfx + "private_key").(string),
			}
		}
	}

	// pfx = prefix + "0.cache.0.s3.0."
	log.Printf("============================> pfx is: %s", pfx)
	if _, hasS3 := d.GetOkExists(pfx + "s3"); hasS3 {
		pfx := pfx + "s3.0."
		cache.S3 = &rcommon.CacheS3Config{
			ServerAddress:  d.Get(pfx + "server_address").(string),
			AccessKey:      d.Get(pfx + "access_key").(string),
			SecretKey:      d.Get(pfx + "secret_key").(string),
			BucketName:     d.Get(pfx + "bucket_name").(string),
			BucketLocation: d.Get(pfx + "bucket_location").(string),
			Insecure:       d.Get(pfx + "insecure").(bool),
		}
	}

	return &cache
}

func refereeConfigStructs(prefix string, d *schema.ResourceData) *referees.Config {

	pfx := prefix + "referee.0."
	if _, hasIt := d.GetOkExists(prefix + "referee"); !hasIt {
		// if we don't have this key, we don't have any of it.  return!
		return nil
	}

	ref := referees.Config{
		// Metrics	*MetricsRefereeConfig `toml:"metrics,omitempty" json:"metrics" namespace:"metrics" tf:"metrics"`
	}

	if _, hasIt := d.GetOkExists(pfx + "metrics"); hasIt {
		pfx := pfx + "metrics.0."

		ref.Metrics = &referees.MetricsRefereeConfig{
			PrometheusAddress: d.Get(pfx + "prometheus_address").(string),
			QueryInterval:     d.Get(pfx + "query_interval").(int),
			Queries:           stringList(pfx+"queries", d),
		}
	}

	return &ref
}

func dockerConfigStructs(prefix string, d *schema.ResourceData) *rcommon.DockerConfig {
	pfx := prefix + "docker.0."
	if _, hasIt := d.GetOkExists(prefix + "docker"); !hasIt {
		// if we don't have this key, we don't have any of it.  return!
		return nil
	}
	dkr := rcommon.DockerConfig{
		DockerCredentials: rdhelpers.DockerCredentials{
			Host:      d.Get(pfx + "host").(string),
			CertPath:  d.Get(pfx + "cert_path").(string),
			TLSVerify: d.Get(pfx + "tls_verify").(bool),
		},
		Hostname:                   d.Get(pfx + "hostname").(string),
		Image:                      d.Get(pfx + "image").(string),
		Runtime:                    d.Get(pfx + "runtime").(string),
		Memory:                     d.Get(pfx + "memory").(string),
		MemorySwap:                 d.Get(pfx + "memory_swap").(string),
		MemoryReservation:          d.Get(pfx + "memory_reservation").(string),
		CPUSetCPUs:                 d.Get(pfx + "cpu_set_cp_us").(string),
		CPUS:                       d.Get(pfx + "cpus").(string),
		CPUShares:                  int64(d.Get(pfx + "cpu_shares").(int)),
		DNS:                        stringList(pfx+"dns", d),
		DNSSearch:                  stringList(pfx+"dns_search", d),
		Privileged:                 d.Get(pfx + "privileged").(bool),
		DisableEntrypointOverwrite: d.Get(pfx + "disable_entrypoint_overwrite").(bool),
		UsernsMode:                 d.Get(pfx + "userns_mode").(string),
		CapAdd:                     stringList(pfx+"cap_add", d),
		CapDrop:                    stringList(pfx+"cap_drop", d),
		OomKillDisable:             d.Get(pfx + "oom_kill_disable").(bool),
		OomScoreAdjust:             d.Get(pfx + "oom_score_adjust").(int),
		SecurityOpt:                stringList(pfx+"security_opt", d),
		Devices:                    stringList(pfx+"devices", d),
		DisableCache:               d.Get(pfx + "disable_cache").(bool),
		Volumes:                    stringList(pfx+"volumes", d),
		VolumeDriver:               d.Get(pfx + "volume_driver").(string),
		CacheDir:                   d.Get(pfx + "cache_dir").(string),
		ExtraHosts:                 stringList(pfx+"extra_hosts", d),
		VolumesFrom:                stringList(pfx+"volumes_from", d),
		NetworkMode:                d.Get(pfx + "network_mode").(string),
		Links:                      stringList(pfx+"links", d),
		WaitForServicesTimeout:     d.Get(pfx + "wait_for_services_timeout").(int),
		AllowedImages:              stringList(pfx+"allowed_images", d),
		AllowedServices:            stringList(pfx+"allowed_services", d),
		PullPolicy:                 stringToDockerPullPolicy(d.Get(pfx + "pull_policy").(string)),
		ShmSize:                    int64(d.Get(pfx + "shm_size").(int)),
		Tmpfs:                      toStringMap(pfx+"tmpfs", d),
		ServicesTmpfs:              toStringMap(pfx+"services_tmpfs", d),
		SysCtls:                    toDockerSysCtls(pfx+"sys_ctls", d),
		HelperImage:                d.Get(pfx + "helper_image").(string),
	}

	subKey := "services"
	if it, hasIt := d.GetOkExists(pfx + subKey); hasIt {
		svcs := make([]*rcommon.DockerService, len(it.([]interface{})))
		for i, _ := range it.([]interface{}) {
			pfx := fmt.Sprintf("%s%s.%d.", pfx, subKey, i)
			svcs[i] = &rcommon.DockerService{
				Alias: d.Get(pfx + "alias").(string),
				// embedded via Services
				Service: rcommon.Service{
					Name: d.Get(pfx + "name").(string),
				},
			}
		}
		dkr.Services = svcs
	}

	type Service struct {
		Name string `toml:"name" long:"name" description:"The image path for the service" json:"name" tf:"name"`
	}

	type DockerService struct {
		Service
		Alias string `toml:"alias,omitempty" long:"alias" description:"The alias of the service"`
	}

	return &dkr
}

func toDockerSysCtls(key string, d *schema.ResourceData) rcommon.DockerSysCtls {
	sm := toStringMap(key, d)
	dsc := rcommon.DockerSysCtls{}

	for k, v := range sm {
		dsc[k] = v
	}

	return dsc
}

func stringToDockerPullPolicy(s string) rcommon.DockerPullPolicy {

	// FIXME this is going to blow up terribly if we don't match
	switch s {
	case "always":
		return rcommon.PullPolicyAlways
	case "never":
		return rcommon.PullPolicyNever
	case "if-not-present":
		return rcommon.PullPolicyIfNotPresent
	}

	// may as well return the default...
	return rcommon.PullPolicyAlways
}

func dockerMachineStructs(prefix string, d *schema.ResourceData) *rcommon.DockerMachine {
	pfx := prefix + "machine.0."
	if _, hasIt := d.GetOkExists(prefix + "machine"); !hasIt {
		// if we don't have this key, we don't have any of it.  return!
		return nil
	}

	m := rcommon.DockerMachine{
		IdleCount:        d.Get(pfx + "idle_count").(int),
		IdleTime:         d.Get(pfx + "idle_time").(int),
		MaxBuilds:        d.Get(pfx + "max_builds").(int),
		MachineDriver:    d.Get(pfx + "machine_driver").(string),
		MachineName:      d.Get(pfx + "machine_name").(string),
		MachineOptions:   stringList(pfx+"machine_options", d),
		OffPeakPeriods:   stringList(pfx+"off_peak_periods", d),
		OffPeakTimezone:  d.Get(pfx + "off_peak_timezone").(string),
		OffPeakIdleCount: d.Get(pfx + "off_peak_idle_count").(int),
		OffPeakIdleTime:  d.Get(pfx + "off_peak_idle_time").(int),
		// unexported
		// offPeakTimePeriods *timeperiod.TimePeriod
	}

	return &m
}
