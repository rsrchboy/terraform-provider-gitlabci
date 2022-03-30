package config

// Code generated by centrifuge. DO NOT EDIT.

import (
	"time"
)

type Artifact struct {
	Name      string          `json:"name"`
	Untracked bool            `json:"untracked"`
	Paths     ArtifactPaths   `json:"paths"`
	Exclude   ArtifactExclude `json:"exclude"`
	When      ArtifactWhen    `json:"when"`
	Type      string          `json:"artifact_type"`
	Format    ArtifactFormat  `json:"artifact_format"`
	ExpireIn  string          `json:"expire_in"`
}

type ArtifactExclude []string

type ArtifactFormat string

type ArtifactPaths []string

type ArtifactWhen string

type Artifacts []Artifact

type ArtifactsOptions struct {
	BaseName string
	ExpireIn string
	Format   ArtifactFormat
	Type     string
}

type Cache struct {
	Key       string        `json:"key"`
	Untracked bool          `json:"untracked"`
	Policy    CachePolicy   `json:"policy"`
	Paths     ArtifactPaths `json:"paths"`
	When      CacheWhen     `json:"when"`
}

type CachePolicy string

type CacheWhen string

type Caches []Cache

type ConfigInfo struct {
	Gpus string `json:"gpus"`
}

type Credentials struct {
	Type     string `json:"type"`
	URL      string `json:"url"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Dependencies []Dependency

type Dependency struct {
	ID            int64                   `json:"id"`
	Token         string                  `json:"token"`
	Name          string                  `json:"name"`
	ArtifactsFile DependencyArtifactsFile `json:"artifacts_file"`
}

type DependencyArtifactsFile struct {
	Filename string `json:"filename"`
	Size     int64  `json:"size"`
}

type DownloadState int

type FeaturesInfo struct {
	Variables               bool `json:"variables"`
	Image                   bool `json:"image"`
	Services                bool `json:"services"`
	Artifacts               bool `json:"artifacts"`
	Cache                   bool `json:"cache"`
	Shared                  bool `json:"shared"`
	UploadMultipleArtifacts bool `json:"upload_multiple_artifacts"`
	UploadRawArtifacts      bool `json:"upload_raw_artifacts"`
	Session                 bool `json:"session"`
	Terminal                bool `json:"terminal"`
	Refspecs                bool `json:"refspecs"`
	Masking                 bool `json:"masking"`
	Proxy                   bool `json:"proxy"`
	RawVariables            bool `json:"raw_variables"`
	ArtifactsExclude        bool `json:"artifacts_exclude"`
	MultiBuildSteps         bool `json:"multi_build_steps"`
	TraceReset              bool `json:"trace_reset"`
	TraceChecksum           bool `json:"trace_checksum"`
	TraceSize               bool `json:"trace_size"`
	VaultSecrets            bool `json:"vault_secrets"`
	Cancelable              bool `json:"cancelable"`
	ReturnExitCode          bool `json:"return_exit_code"`
	ServiceVariables        bool `json:"service_variables"`
}

type GitInfo struct {
	RepoURL   string         `json:"repo_url"`
	Ref       string         `json:"ref"`
	Sha       string         `json:"sha"`
	BeforeSha string         `json:"before_sha"`
	RefType   GitInfoRefType `json:"ref_type"`
	Refspecs  []string       `json:"refspecs"`
	Depth     int            `json:"depth"`
}

type GitInfoRefType string

type GitlabFeatures struct {
	TraceSections  bool               `json:"trace_sections"`
	FailureReasons []JobFailureReason `json:"failure_reasons"`
}

type Image struct {
	Name       string       `json:"name"`
	Alias      string       `json:"alias,omitempty"`
	Command    []string     `json:"command,omitempty"`
	Entrypoint []string     `json:"entrypoint,omitempty"`
	Ports      []Port       `json:"ports,omitempty"`
	Variables  JobVariables `json:"variables,omitempty"`
}

type JobCredentials struct {
	ID          int64  `long:"id" env:"CI_JOB_ID" description:"The build ID to download and upload artifacts for"`
	Token       string `long:"token" env:"CI_JOB_TOKEN" required:"true" description:"Build token"`
	URL         string `long:"url" env:"CI_SERVER_URL" required:"true" description:"GitLab CI URL"`
	TLSCAFile   string `long:"tls-ca-file" env:"CI_SERVER_TLS_CA_FILE" description:"File containing the certificates to verify the peer when using HTTPS"`
	TLSCertFile string `long:"tls-cert-file" env:"CI_SERVER_TLS_CERT_FILE" description:"File containing certificate for TLS client auth with runner when using HTTPS"`
	TLSKeyFile  string `long:"tls-key-file" env:"CI_SERVER_TLS_KEY_FILE" description:"File containing private key for TLS client auth with runner when using HTTPS"`
}

type JobFailureReason string

type JobInfo struct {
	Name        string `json:"name"`
	Stage       string `json:"stage"`
	ProjectID   int64  `json:"project_id"`
	ProjectName string `json:"project_name"`
}

type JobRequest struct {
	Info       VersionInfo  `json:"info,omitempty"`
	Token      string       `json:"token,omitempty"`
	LastUpdate string       `json:"last_update,omitempty"`
	Session    *SessionInfo `json:"session,omitempty"`
}

type JobResponse struct {
	ID            int64          `json:"id"`
	Token         string         `json:"token"`
	AllowGitFetch bool           `json:"allow_git_fetch"`
	JobInfo       JobInfo        `json:"job_info"`
	GitInfo       GitInfo        `json:"git_info"`
	RunnerInfo    RunnerInfo     `json:"runner_info"`
	Variables     JobVariables   `json:"variables"`
	Steps         Steps          `json:"steps"`
	Image         Image          `json:"image"`
	Services      Services       `json:"services"`
	Artifacts     Artifacts      `json:"artifacts"`
	Cache         Caches         `json:"cache"`
	Credentials   []Credentials  `json:"credentials"`
	Dependencies  Dependencies   `json:"dependencies"`
	Features      GitlabFeatures `json:"features"`
	Secrets       Secrets        `json:"secrets,omitempty"`
	TLSCAChain    string         `json:"-"`
	TLSAuthCert   string         `json:"-"`
	TLSAuthKey    string         `json:"-"`
}

type JobState string

type JobTraceOutput struct {
	Checksum string `json:"checksum,omitempty"`
	Bytesize int    `json:"bytesize,omitempty"`
}

type PatchState int

type PatchTraceResult struct {
	SentOffset        int
	CancelRequested   bool
	State             PatchState
	NewUpdateInterval time.Duration
}

type Port struct {
	Number   int    `json:"number,omitempty"`
	Protocol string `json:"protocol,omitempty"`
	Name     string `json:"name,omitempty"`
}

type RegisterRunnerParameters struct {
	Description     string `json:"description,omitempty"`
	MaintenanceNote string `json:"maintenance_note,omitempty"`
	Tags            string `json:"tag_list,omitempty"`
	RunUntagged     bool   `json:"run_untagged"`
	Locked          bool   `json:"locked"`
	AccessLevel     string `json:"access_level,omitempty"`
	MaximumTimeout  int    `json:"maximum_timeout,omitempty"`
	Active          bool   `json:"active"`
}

type RegisterRunnerRequest struct {
	RegisterRunnerParameters
	Info  VersionInfo `json:"info,omitempty"`
	Token string      `json:"token,omitempty"`
}

type RegisterRunnerResponse struct {
	Token string `json:"token,omitempty"`
}

type RunnerInfo struct {
	Timeout int `json:"timeout"`
}

type Secret struct {
	Vault *VaultSecret `json:"vault,omitempty"`
	File  *bool        `json:"file,omitempty"`
}

type Secrets map[string]Secret

type Services []Image

type SessionInfo struct {
	URL           string `json:"url,omitempty"`
	Certificate   string `json:"certificate,omitempty"`
	Authorization string `json:"authorization,omitempty"`
}

type Step struct {
	Name         StepName   `json:"name"`
	Script       StepScript `json:"script"`
	Timeout      int        `json:"timeout"`
	When         StepWhen   `json:"when"`
	AllowFailure bool       `json:"allow_failure"`
}

type StepName string

type StepScript []string

type StepWhen string

type Steps []Step

type UnregisterRunnerRequest struct {
	Token string `json:"token,omitempty"`
}

type UpdateJobInfo struct {
	ID            int64
	State         JobState
	FailureReason JobFailureReason
	Output        JobTraceOutput
	ExitCode      int
}

type UpdateJobRequest struct {
	Info          VersionInfo      `json:"info,omitempty"`
	Token         string           `json:"token,omitempty"`
	State         JobState         `json:"state,omitempty"`
	FailureReason JobFailureReason `json:"failure_reason,omitempty"`
	Checksum      string           `json:"checksum,omitempty"`
	Output        JobTraceOutput   `json:"output,omitempty"`
	ExitCode      int              `json:"exit_code,omitempty"`
}

type UpdateJobResult struct {
	State             UpdateState
	CancelRequested   bool
	NewUpdateInterval time.Duration
}

type UpdateState int

type UploadState int

type VaultAuth struct {
	Name string        `json:"name"`
	Path string        `json:"path"`
	Data VaultAuthData `json:"data"`
}

type VaultAuthData map[string]interface{}

type VaultEngine struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

type VaultSecret struct {
	Server VaultServer `json:"server"`
	Engine VaultEngine `json:"engine"`
	Path   string      `json:"path"`
	Field  string      `json:"field"`
}

type VaultServer struct {
	URL       string    `json:"url"`
	Auth      VaultAuth `json:"auth"`
	Namespace string    `json:"namespace"`
}

type VerifyRunnerRequest struct {
	Token string `json:"token,omitempty"`
}

type VersionInfo struct {
	Name         string       `json:"name,omitempty"`
	Version      string       `json:"version,omitempty"`
	Revision     string       `json:"revision,omitempty"`
	Platform     string       `json:"platform,omitempty"`
	Architecture string       `json:"architecture,omitempty"`
	Executor     string       `json:"executor,omitempty"`
	Shell        string       `json:"shell,omitempty"`
	Features     FeaturesInfo `json:"features"`
	Config       ConfigInfo   `json:"config,omitempty"`
}

// vim: set nowrap :
