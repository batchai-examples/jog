package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseConfigExpression(t *testing.T) {
	testCases := []struct {
		expr       string
		expectedKey  string
		expectedValue string
		expectError bool
	}{
		{"key=value", "key", "value", false},
		{"invalid_expr", "", "", true},
		{"another_key=another_value", "another_key", "another_value", false},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("expr=%s", tc.expr), func(t *testing.T) {
			key, value, err := ParseConfigExpression(tc.expr)
			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedKey, key)
				assert.Equal(t, tc.expectedValue, value)
			}
		})
	}
}

func TestReadConfig(t *testing.T) {
	testCases := []struct {
		configFilePath string
		expected     config.Configuration
	}{
		{"", config.WithDefaultYamlFile()},
		{"path/to/config.yaml", config.WithYamlFile("path/to/config.yaml")},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("configFilePath=%s", tc.configFilePath), func(t *testing.T) {
			result := ReadConfig(tc.configFilePath)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestPrintConfigItem(t *testing.T) {
	testCases := []struct {
		configItemPath string
		expected       string
	}{
		{"key", "value"},
		{"nonexistent_key", ""},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("configItemPath=%s", tc.configItemPath), func(t *testing.T) {
			m := map[string]interface{}{
				"key": "value",
			}
			var output strings.Builder
			oldPrint := fmt.Print
			defer func() { fmt.Print = oldPrint }()
			fmt.Print = func(s string) {
				output.WriteString(s)
			}

			PrintConfigItem(m, tc.configItemPath)

			if tc.expected == "" {
				assert.Empty(t, output.String())
			} else {
				assert.Equal(t, tc.expected, output.String())
			}
		})
	}
}

func TestSetConfigItem(t *testing.T) {
	testCases := []struct {
		configItemPath string
		configItemValue string
		expected       map[string]interface{}
	}{
		{"key", "new_value", map[string]interface{}{"key": "new_value"}},
		{"nonexistent_key", "value", map[string]interface{}{"nonexistent_key": "value"}},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("configItemPath=%s, configItemValue=%s", tc.configItemPath, tc.configItemValue), func(t *testing.T) {
			cfg := config.Configuration{}
			m := map[string]interface{}{
				"key": "value",
			}

			SetConfigItem(cfg, m, tc.configItemPath, tc.configItemValue)

			assert.Equal(t, tc.expected, m)
		})
	}
}
