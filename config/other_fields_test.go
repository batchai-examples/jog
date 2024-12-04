package config

import (
	"testing"
)

func TestReset(t *testing.T) {
	// Arrange
	i := &OtherFieldsT{
		Name:      &ElementT{},
		Separator: &SeparatorFieldT{},
		Value:     &ElementT{},
	}

	// Act
	i.Reset()

	// Assert
	if i.Name == nil || !i.Name.IsReset() {
		t.Errorf("Expected Name to be reset, but got %v", i.Name)
	}
	if i.Separator == nil || !i.Separator.IsReset() {
		t.Errorf("Expected Separator to be reset, but got %v", i.Separator)
	}
	if i.Value == nil || !i.Value.IsReset() {
		t.Errorf("Expected Value to be reset, but got %v", i.Value)
	}
}

func TestToMap(t *testing.T) {
	// Arrange
	i := &OtherFieldsT{
		Name:      &ElementT{Name: "testName"},
		Separator: &SeparatorFieldT{Value: "-"},
		Value:     &ElementT{Value: "testValue"},
	}

	// Act
	result := i.ToMap()

	// Assert
	if result["name"] != "testName" {
		t.Errorf("Expected name to be 'testName', but got %v", result["name"])
	}
	if result["separator"] != "-" {
		t.Errorf("Expected separator to be '-', but got %v", result["separator"])
	}
	if result["value"] != "testValue" {
		t.Errorf("Expected value to be 'testValue', but got %v", result["value"])
	}
}

func TestResetWithNilFields(t *testing.T) {
	// Arrange
	i := &OtherFieldsT{}

	// Act
	i.Reset()

	// Assert
	if i.Name != nil || !i.Name.IsReset() {
		t.Errorf("Expected Name to be reset, but got %v", i.Name)
	}
	if i.Separator != nil || !i.Separator.IsReset() {
		t.Errorf("Expected Separator to be reset, but got %v", i.Separator)
	}
	if i.Value != nil || !i.Value.IsReset() {
		t.Errorf("Expected Value to be reset, but got %v", i.Value)
	}
}

func TestToMapWithNilFields(t *testing.T) {
	// Arrange
	i := &OtherFieldsT{}

	// Act
	result := i.ToMap()

	// Assert
	if result["name"] != nil {
		t.Errorf("Expected name to be nil, but got %v", result["name"])
	}
	if result["separator"] != nil {
		t.Errorf("Expected separator to be nil, but got %v", result["separator"])
	}
	if result["value"] != nil {
		t.Errorf("Expected value to be nil, but got %v", result["value"])
	}
}

func TestResetWithEmptyFields(t *testing.T) {
	// Arrange
	i := &OtherFieldsT{
		Name:      &ElementT{Name: ""},
		Separator: &SeparatorFieldT{Value: ""},
		Value:     &ElementT{Value: ""},
	}

	// Act
	i.Reset()

	// Assert
	if i.Name == nil || !i.Name.IsReset() {
		t.Errorf("Expected Name to be reset, but got %v", i.Name)
	}
	if i.Separator == nil || !i.Separator.IsReset() {
		t.Errorf("Expected Separator to be reset, but got %v", i.Separator)
	}
	if i.Value == nil || !i.Value.IsReset() {
		t.Errorf("Expected Value to be reset, but got %v", i.Value)
	}
}

func TestToMapWithEmptyFields(t *testing.T) {
	// Arrange
	i := &OtherFieldsT{
		Name:      &ElementT{Name: ""},
		Separator: &SeparatorFieldT{Value: ""},
		Value:     &ElementT{Value: ""},
	}

	// Act
	result := i.ToMap()

	// Assert
	if result["name"] != "" {
		t.Errorf("Expected name to be '', but got %v", result["name"])
	}
	if result["separator"] != "" {
		t.Errorf("Expected separator to be '', but got %v", result["separator"])
	}
	if result["value"] != "" {
		t.Errorf("Expected value to be '', but got %v", result["value"])
	}
}
