package gitlabci

import (
	"crypto/sha256"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// copied from the gitlab provider
func stringSetToStringSlice(stringSet *schema.Set) *[]string {
	ret := []string{}
	if stringSet == nil {
		return &ret
	}
	for _, envVal := range stringSet.List() {
		ret = append(ret, envVal.(string))
	}
	return &ret
}

// copied from the gitlab provider
func hashSum(contents interface{}) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(contents.(string))))
}
