package gitlabci

import (
	"crypto/sha256"
	"fmt"
)

// copied from the gitlab provider
func hashSum(contents interface{}) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(contents.(string))))
}
