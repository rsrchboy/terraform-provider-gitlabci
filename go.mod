module gitlab.com/rsrchboy/terraform-provider-gitlabci

go 1.13

require (
	cloud.google.com/go v0.49.0 // indirect
	github.com/elazarl/goproxy v0.0.0-20191011121108-aa519ddbe484 // indirect
	github.com/hashicorp/terraform-plugin-sdk v1.4.1
	github.com/parnurzeal/gorequest v0.2.16
	github.com/pkg/errors v0.8.1 // indirect
	github.com/smartystreets/goconvey v1.6.4 // indirect
	github.com/stretchr/testify v1.4.0 // indirect
	golang.org/x/lint v0.0.0-20191125180803-fdd1cda4f05f // indirect
	golang.org/x/sys v0.0.0-20190916202348-b4ddaad3f8a3 // indirect
	golang.org/x/tools v0.0.0-20191126055441-b0650ceb63d9 // indirect
	gopkg.in/yaml.v2 v2.2.4 // indirect
	moul.io/http2curl v1.0.0 // indirect
)

// replace github.com/docker/docker v1.4.2-0.20190822180741-9552f2b2fdde => github.com/docker/engine v1.4.2-0.20190822180741-9552f2b2fdde
replace github.com/docker/docker v1.13.1 => github.com/docker/engine v1.4.2-0.20191113042239-ea84732a7725

replace github.com/docker/docker v1.4.2-0.20190822180741-9552f2b2fdde => github.com/docker/engine v1.4.2-0.20191113042239-ea84732a7725

replace github.com/docker/engine v1.4.2-0.20190822180741-9552f2b2fdde => github.com/docker/engine v1.4.2-0.20191113042239-ea84732a7725

// github.com/docker/cli v0.0.0-20191105005515-99c5edceb48d
// replace github.com/docker/cli v0.0.0-20190814185437-1752eb3626e3 => github.com/docker/cli v0.0.0-20190814185437-1752eb3626e3

replace github.com/minio/go-homedir v0.0.0-20190425115525-017018655514 => gitlab.com/steveazz/go-homedir v0.0.0-20190425115525-017018655514
