package config

import (
	"fmt"

	// "encoding/json"
	// "github.com/BurntSushi/toml"
)

// type JSONPayload struct {
// 	*Configuration
// }

// func (c JSONPayload) MarshalJSON() ([]byte, error) {
// 	if c.Configuration == nil {
// 		return nil, nil
// 	}
// 	return json.Marshal(c.Configuration)
// }

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
				return fmt.Errorf("unexpected data type: %!v(MISSING)", item)
			}
		}
	default:
		return fmt.Errorf("unexpected data type: %!v(MISSING)", v)
	}

	return nil
}
