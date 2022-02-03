package config

import (
	"fmt"
)

// StringOrArray is pulled from gitlab-runner/common

// // StringOrArray implements UnmarshalTOML to unmarshal either a string or array of strings.
// type StringOrArray []string

func (p *StringOrArray) UnmarshalTOML(data interface{}) error {
	switch v := data.(type) {
	case string:
		*p = StringOrArray{v}
	case []interface{}:
		for _, vv := range v {
			switch item := vv.(type) {
			case string:
				*p = append(*p, item)
			default:
				return fmt.Errorf("unexpected data type: %v", item)
			}
		}
	default:
		return fmt.Errorf("unexpected data type: %v", v)
	}

	return nil
}
