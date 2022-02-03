package config

// source: https://gitlab.com/gitlab-org/gitlab-runner/blob/main/common/config_test.go

import (
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/stretchr/testify/assert"
)

func TestStringOrArray_UnmarshalTOML(t *testing.T) {
	tests := map[string]struct {
		toml           string
		expectedResult StringOrArray
		expectedErr    bool
	}{
		"no fields": {
			toml:           "",
			expectedResult: nil,
			expectedErr:    false,
		},
		"empty string_or_array": {
			toml:           `string_or_array = ""`,
			expectedResult: StringOrArray{""},
			expectedErr:    false,
		},
		"string": {
			toml:           `string_or_array = "always"`,
			expectedResult: StringOrArray{"always"},
			expectedErr:    false,
		},
		"slice with invalid single value": {
			toml:        `string_or_array = 10`,
			expectedErr: true,
		},
		"valid slice with multiple values": {
			toml:           `string_or_array = ["unknown", "always"]`,
			expectedResult: StringOrArray{"unknown", "always"},
			expectedErr:    false,
		},
		"slice with mixed values": {
			toml:        `string_or_array = ["unknown", 10]`,
			expectedErr: true,
		},
		"slice with invalid values": {
			toml:        `string_or_array = [true, false]`,
			expectedErr: true,
		},
	}

	for tn, tt := range tests {
		t.Run(tn, func(t *testing.T) {
			type Config struct {
				StringOrArray StringOrArray `toml:"string_or_array"`
			}

			var result Config
			_, err := toml.Decode(tt.toml, &result)

			if tt.expectedErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.expectedResult, result.StringOrArray)
		})
	}
}
