// ____                 _     _             _____         _
// |  _ \ _ __ _____   _(_) __| | ___ _ __  |_   _|__  ___| |_ ___
// | |_) | '__/ _ \ \ / / |/ _` |/ _ \ '__|   | |/ _ \/ __| __/ __|
// |  __/| | | (_) \ V /| | (_| |  __/ |      | |  __/\__ \ |_\__ \
// |_|   |_|  \___/ \_/ |_|\__,_|\___|_|      |_|\___||___/\__|___/

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load("../../.env")
}

// providerFactories are used to instantiate a provider during acceptance
// testing.  The factory function will be invoked for every Terraform CLI
// command executed to create a provider server to which the CLI can reattach.
var providerFactories = map[string]func() (*schema.Provider, error){
	"gitlabci": func() (*schema.Provider, error) {
		return NewProvider("dev", "deadb33f")(), nil
	},
}

func TestProvider(t *testing.T) {
	if err := NewProvider("dev", "deadb22f")().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func testAccPreCheck(t *testing.T) {
	// You can add code here to run prior to any test case execution, for
	// example assertions about the appropriate environment variables being
	// set are common to see in a pre-check function.
	//
	// if v := os.Getenv("GITLAB_TOKEN"); v == "" {
	// 	t.Fatal("GITLAB_TOKEN must be set for acceptance tests")
	// }
}
