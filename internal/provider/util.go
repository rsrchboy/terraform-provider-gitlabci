package provider

import (
	"crypto/sha256"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type stringMap map[string]string

func toStringMap(key string, d *schema.ResourceData) stringMap {
	imap := d.Get(key).(map[string]interface{})
	// log.Printf("iMapToStringMap: %T, %s", imap, spew.Sdump(imap))
	smap := make(stringMap, len(imap))

	for k, v := range imap {
		smap[k] = v.(string)
	}

	return smap
}

func stringList(key string, d *schema.ResourceData) []string {
	stringsI := d.Get(key).([]interface{})
	strings := make([]string, len(stringsI))

	// I'm hopeful there's a better way I'm simply unaware of as of yet
	for i, str := range stringsI {
		strings[i] = str.(string)
	}

	return strings
}

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
