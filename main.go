package main

import (
	"context"
	"flag"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/spf13/pflag"
	"gitlab.com/rsrchboy/terraform-provider-gitlabci/internal/provider"
)

// Run "go generate" to format example terraform files and generate the docs
// for the registry/website

// If you do not have terraform installed, you can remove the formatting
// command, but its suggested to ensure the documentation is formatted
// properly.
//go:generate terraform fmt -recursive ./examples/

// Run the docs generation tool, check its repository for more information on
// how it works and how docs can be customized.
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

// func main() {
// 	plugin.Serve(&plugin.ServeOpts{
// 		ProviderFunc: func() terraform.ResourceProvider {
// 			return gitlabci.Provider()
// 		},
// 	})
// }

var (
	// these will be set by the goreleaser configuration
	// to appropriate values for the compiled binary
	version string = "dev"
	commit  string = ""
)

func main() {
	var debugMode bool

	pflag.BoolVar(&debugMode, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := &plugin.ServeOpts{ProviderFunc: provider.NewProvider(version, commit)}

	if debugMode {
		// TODO: update this string with the full name of your provider as used in your configs
		err := plugin.Debug(context.Background(), "registry.terraform.io/rsrchboy/gitlabci", opts)
		if err != nil {
			log.Fatal(err.Error())
		}
		return
	}

	plugin.Serve(opts)
}
