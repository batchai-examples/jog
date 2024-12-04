!!!!test_begin!!!!

package config

import (
	"testing"
)

func TestStartupLine_UnmarshalYAML(t *testing.T) {
	t.Run("Happy Path", func(t *testing.T) {
		// Given
		yamlData := `contains: "Started Application in"`
		var startupLine StartupLineT

		// When
		err := startupLine.UnmarshalYAML(func(v interface{}) error {
			return unmarshalYAML(yamlData, v)
		})

		// Then
		if err != nil {
			t.Errorf("Expected no error, but got: %v", err)
		}
		if startupLine.Contains != "Started Application in" {
			t.Errorf("Expected Contains to be 'Started Application in', but got: %s", startupLine.Contains)
		}
	})

	t.Run("Negative Case - Invalid YAML", func(t *testing.T) {
		// Given
		yamlData := `contains: "Started Application in"`
		var startupLine StartupLineT

		// When
		err := startupLine.UnmarshalYAML(func(v interface{}) error {
			return unmarshalYAML("invalid yaml data", v)
		})

		// Then
		if err == nil {
			t.Errorf("Expected an error, but got none")
		}
	})

	t.Run("Corner Case - Empty YAML", func(t *testing.T) {
		// Given
		yamlData := ""
		var startupLine StartupLineT

		// When
		err := startupLine.UnmarshalYAML(func(v interface{}) error {
			return unmarshalYAML(yamlData, v)
		})

		// Then
		if err == nil {
			t.Errorf("Expected an error, but got none")
		}
	})

	t.Run("Happy Path - Contains Field Missing", func(t *testing.T) {
		// Given
		yamlData := `otherField: "some value"`
		var startupLine StartupLineT

		// When
		err := startupLine.UnmarshalYAML(func(v interface{}) error {
			return unmarshalYAML(yamlData, v)
		})

		// Then
		if err != nil {
			t.Errorf("Expected no error, but got: %v", err)
		}
		if startupLine.Contains != "Started Application in" {
			t.Errorf("Expected Contains to be 'Started Application in', but got: %s", startupLine.Contains)
		}
	})

	t.Run("Negative Case - Invalid Contains Value", func(t *testing.T) {
		// Given
		yamlData := `contains: 123`
		var startupLine StartupLineT

		// When
		err := startupLine.UnmarshalYAML(func(v interface{}) error {
			return unmarshalYAML(yamlData, v)
		})

		// Then
		if err == nil {
			t.Errorf("Expected an error, but got none")
		}
	})

	t.Run("Happy Path - Contains Field Empty", func(t *testing.T) {
		// Given
		yamlData := `contains: ""`
		var startupLine StartupLineT

		// When
		err := startupLine.UnmarshalYAML(func(v interface{}) error {
			return unmarshalYAML(yamlData, v)
		})

		// Then
		if err != nil {
			t.Errorf("Expected no error, but got: %v", err)
		}
		if startupLine.Contains != "" {
			t.Errorf("Expected Contains to be '', but got: %s", startupLine.Contains)
		}
	})

	t.Run("Negative Case - Contains Field Null", func(t *testing.T) {
		// Given
		yamlData := `contains: null`
		var startupLine StartupLineT

		// When
		err := startupLine.UnmarshalYAML(func(v interface{}) error {
			return unmarshalYAML(yamlData, v)
		})

		// Then
		if err == nil {
			t.Errorf("Expected an error, but got none")
		}
	})

	t.Run("Happy Path - Contains Field with Leading and Trailing Spaces", func(t *testing.T) {
		// Given
		yamlData := `contains: "  Started Application in  "`
		var startupLine StartupLineT

		// When
		err := startupLine.UnmarshalYAML(func(v interface{}) error {
			return unmarshalYAML(yamlData, v)
		})

		// Then
		if err != nil {
			t.Errorf("Expected no error, but got: %v", err)
		}
		if startupLine.Contains != "Started Application in" {
			t.Errorf("Expected Contains to be 'Started Application in', but got: %s", startupLine.Contains)
		}
	})

	t.Run("Negative Case - Contains Field with Special Characters", func(t *testing.T) {
		// Given
		yamlData := `contains: "Started!Application@in"`
		var startupLine StartupLineT

		// When
		err := startupLine.UnmarshalYAML(func(v interface{}) error {
			return unmarshalYAML(yamlData, v)
		})

		// Then
		if err != nil {
			t.Errorf("Expected no error, but got: %v", err)
		}
		if startupLine.Contains != "Started!Application@in" {
			t.Errorf("Expected Contains to be 'Started!Application@in', but got: %s", startupLine.Contains)
		}
	})
}

func unmarshalYAML(data string, v interface{}) error {
	return yaml.Unmarshal([]byte(data), v)
}
!!!!test_end!!!!
