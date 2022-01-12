// ____                 _     _             _____         _
// |  _ \ _ __ _____   _(_) __| | ___ _ __  |_   _|__  ___| |_ ___
// | |_) | '__/ _ \ \ / / |/ _` |/ _ \ '__|   | |/ _ \/ __| __/ __|
// |  __/| | | (_) \ V /| | (_| |  __/ |      | |  __/\__ \ |_\__ \
// |_|   |_|  \___/ \_/ |_|\__,_|\___|_|      |_|\___||___/\__|___/

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider()
	testAccProviders = map[string]terraform.ResourceProvider{
		"gitlabci": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ terraform.ResourceProvider = Provider()
}

func testAccPreCheck(t *testing.T) {
	// if v := os.Getenv("GITLAB_TOKEN"); v == "" {
	// 	t.Fatal("GITLAB_TOKEN must be set for acceptance tests")
	// }
}
