module gitlab.com/rsrchboy/terraform-provider-gitlabci

go 1.13

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/DiSiqueira/GoTree v1.0.0
	github.com/Microsoft/hcsshim v0.8.7 // indirect
	github.com/containerd/containerd v1.3.2 // indirect
	github.com/containerd/continuity v0.0.0-20191214063359-1097c8bae83b // indirect
	github.com/davecgh/go-spew v1.1.1
	github.com/docker/cli v0.0.0-20190814185437-1752eb3626e3 // indirect
	github.com/docker/distribution v2.7.1+incompatible // indirect
	github.com/docker/docker v1.13.1 // indirect
	github.com/docker/docker-credential-helpers v0.6.3 // indirect
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-metrics v0.0.1 // indirect
	github.com/docker/machine v0.16.2 // indirect
	github.com/giantswarm/to v0.0.0-20191022113953-f2078541ec95
	github.com/gorhill/cronexpr v0.0.0-20180427100037-88b0669f7d75 // indirect
	github.com/gorilla/mux v1.7.3 // indirect
	github.com/gorilla/websocket v1.4.1 // indirect
	github.com/hashicorp/terraform-plugin-sdk v1.4.1
	github.com/imdario/mergo v0.3.8
	github.com/morikuni/aec v1.0.0 // indirect
	github.com/parnurzeal/gorequest v0.2.16
	github.com/pkg/errors v0.8.1
	github.com/stoewer/go-strcase v1.1.0
	github.com/tevino/abool v0.0.0-20170917061928-9b9efcf221b5 // indirect
	github.com/urfave/cli v1.22.2 // indirect
	gitlab.com/gitlab-org/gitlab-runner v1.11.1-0.20200110014830-8e5683f82927
	k8s.io/api v0.17.0 // indirect
	moul.io/http2curl v1.0.0 // indirect
)

// replace github.com/docker/docker v1.4.2-0.20190822180741-9552f2b2fdde => github.com/docker/engine v1.4.2-0.20190822180741-9552f2b2fdde
replace github.com/docker/docker v1.13.1 => github.com/docker/engine v1.4.2-0.20191113042239-ea84732a7725

replace github.com/docker/docker v1.4.2-0.20190822180741-9552f2b2fdde => github.com/docker/engine v1.4.2-0.20191113042239-ea84732a7725

replace github.com/docker/engine v1.4.2-0.20190822180741-9552f2b2fdde => github.com/docker/engine v1.4.2-0.20191113042239-ea84732a7725

// github.com/docker/cli v0.0.0-20191105005515-99c5edceb48d
// replace github.com/docker/cli v0.0.0-20190814185437-1752eb3626e3 => github.com/docker/cli v0.0.0-20190814185437-1752eb3626e3

replace github.com/minio/go-homedir v0.0.0-20190425115525-017018655514 => gitlab.com/steveazz/go-homedir v0.0.0-20190425115525-017018655514
