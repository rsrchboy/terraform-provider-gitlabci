module gitlab.com/rsrchboy/terraform-provider-gitlabci

go 1.13

require (
	github.com/armon/go-radix v1.0.0 // indirect
	github.com/elazarl/goproxy v0.0.0-20191011121108-aa519ddbe484 // indirect
	github.com/hashicorp/terraform-plugin-docs v0.5.1
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.10.1
	github.com/parnurzeal/gorequest v0.2.16
	github.com/posener/complete v1.2.1 // indirect
	github.com/smartystreets/goconvey v1.6.4 // indirect
	github.com/spf13/pflag v1.0.5
	moul.io/http2curl v1.0.0 // indirect
)

// replace github.com/docker/docker v1.4.2-0.20190822180741-9552f2b2fdde => github.com/docker/engine v1.4.2-0.20190822180741-9552f2b2fdde
replace github.com/docker/docker v1.13.1 => github.com/docker/engine v1.4.2-0.20191113042239-ea84732a7725

replace github.com/docker/docker v1.4.2-0.20190822180741-9552f2b2fdde => github.com/docker/engine v1.4.2-0.20191113042239-ea84732a7725

replace github.com/docker/engine v1.4.2-0.20190822180741-9552f2b2fdde => github.com/docker/engine v1.4.2-0.20191113042239-ea84732a7725

// github.com/docker/cli v0.0.0-20191105005515-99c5edceb48d
// replace github.com/docker/cli v0.0.0-20190814185437-1752eb3626e3 => github.com/docker/cli v0.0.0-20190814185437-1752eb3626e3

replace github.com/minio/go-homedir v0.0.0-20190425115525-017018655514 => gitlab.com/steveazz/go-homedir v0.0.0-20190425115525-017018655514
