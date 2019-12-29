module gitlab.com/rsrchboy/terraform-provider-gitlabci

go 1.13

require (
	github.com/Azure/go-ansiterm v0.0.0-20170929234023-d6e3b3328b78 // indirect
	github.com/BurntSushi/toml v0.3.1
	github.com/Microsoft/hcsshim v0.8.7 // indirect
	github.com/containerd/containerd v1.3.2 // indirect
	github.com/containerd/continuity v0.0.0-20191214063359-1097c8bae83b // indirect
	github.com/docker/cli v0.0.0-20190814185437-1752eb3626e3 // indirect
	github.com/docker/distribution v2.7.1+incompatible // indirect
	github.com/docker/docker v1.13.1 // indirect
	github.com/docker/docker-credential-helpers v0.6.3 // indirect
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-metrics v0.0.1 // indirect
	github.com/docker/libtrust v0.0.0-20160708172513-aabc10ec26b7 // indirect
	github.com/docker/machine v0.16.2 // indirect
	github.com/gopherjs/gopherjs v0.0.0-20181017120253-0766667cb4d1 // indirect
	github.com/gorhill/cronexpr v0.0.0-20180427100037-88b0669f7d75 // indirect
	github.com/gorilla/mux v1.7.3 // indirect
	github.com/gorilla/websocket v1.4.0 // indirect
	github.com/hashicorp/terraform-plugin-sdk v1.4.1
	github.com/imdario/mergo v0.3.8
	github.com/jtolds/gls v4.2.1+incompatible // indirect
	github.com/morikuni/aec v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.0.1 // indirect
	github.com/opencontainers/runc v0.1.1 // indirect
	github.com/parnurzeal/gorequest v0.2.16
	github.com/rsrchboy/structs v1.1.1
	github.com/smartystreets/assertions v0.0.0-20180927180507-b2de0cb4f26d // indirect
	github.com/smartystreets/goconvey v0.0.0-20180222194500-ef6db91d284a // indirect
	github.com/stoewer/go-strcase v1.1.0
	github.com/tevino/abool v0.0.0-20170917061928-9b9efcf221b5 // indirect
	github.com/urfave/cli v1.22.2 // indirect
	gitlab.com/ayufan/golang-cli-helpers v0.0.0-20171103152739-a7cf72d604cd // indirect
	gitlab.com/gitlab-org/gitlab-runner v12.5.0+incompatible
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
