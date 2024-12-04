package util

import (
	"strings"
	"testing"
)

// TestPrintable_IsEnabled tests the IsEnabled method of Printable interface.
func TestPrintable_IsEnabled(t *testing.T) {
	// Create a mock implementation of Printable interface
	mockPrintable := &mockPrintable{}

	// Test case 1: Happy path - IsEnabled returns true
	t.Run("IsEnabled returns true", func(t *testing.T) {
		mockPrintable.IsEnabledFunc = func() bool {
			return true
		}
		result := mockPrintable.IsEnabled()
		if result != true {
			t.Errorf("Expected true, got %v", result)
		}
	})

	// Test case 2: Negative case - IsEnabled returns false
	t.Run("IsEnabled returns false", func(t *testing.T) {
		mockPrintable.IsEnabledFunc = func() bool {
			return false
		}
		result := mockPrintable.IsEnabled()
		if result != false {
			t.Errorf("Expected false, got %v", result)
		}
	})
}

// TestPrintable_GetColor tests the GetColor method of Printable interface.
func TestPrintable_GetColor(t *testing.T) {
	// Create a mock implementation of Printable interface
	mockPrintable := &mockPrintable{}

	// Test case 1: Happy path - GetColor returns a valid Color
	t.Run("GetColor returns a valid Color", func(t *testing.T) {
		expectedColor := Red
		mockPrintable.GetColorFunc = func(value string) Color {
			return expectedColor
		}
		result := mockPrintable.GetColor("test")
		if result != expectedColor {
			t.Errorf("Expected %v, got %v", expectedColor, result)
		}
	})

	// Test case 2: Negative case - GetColor returns an invalid Color
	t.Run("GetColor returns an invalid Color", func(t *testing.T) {
		expectedColor := InvalidColor
		mockPrintable.GetColorFunc = func(value string) Color {
			return expectedColor
		}
		result := mockPrintable.GetColor("test")
		if result != expectedColor {
			t.Errorf("Expected %v, got %v", expectedColor, result)
		}
	})
}

// TestPrintable_PrintTo tests the PrintTo method of Printable interface.
func TestPrintable_PrintTo(t *testing.T) {
	// Create a mock implementation of Printable interface
	mockPrintable := &mockPrintable{}

	// Test case 1: Happy path - PrintTo appends the string to the builder with the specified color
	t.Run("PrintTo appends the string to the builder with the specified color", func(t *testing.T) {
		expectedColor := Red
		expectedString := "test"
		mockPrintable.PrintToFunc = func(color Color, builder *strings.Builder, a string) {
			builder.WriteString(a)
		}
		var builder strings.Builder
		mockPrintable.PrintTo(expectedColor, &builder, expectedString)
		result := builder.String()
		if result != expectedString {
			t.Errorf("Expected %v, got %v", expectedString, result)
		}
	})

	// Test case 2: Negative case - PrintTo does not append the string to the builder
	t.Run("PrintTo does not append the string to the builder", func(t *testing.T) {
		expectedColor := Red
		expectedString := "test"
		mockPrintable.PrintToFunc = func(color Color, builder *strings.Builder, a string) {
			// Do nothing
		}
		var builder strings.Builder
		mockPrintable.PrintTo(expectedColor, &builder, expectedString)
		result := builder.String()
		if result != "" {
			t.Errorf("Expected empty string, got %v", result)
		}
	})
}

// mockPrintable is a mock implementation of Printable interface for testing purposes.
type mockPrintable struct {
	IsEnabledFunc  func() bool
	GetColorFunc   func(value string) Color
	PrintToFunc    func(color Color, builder *strings.Builder, a string)
	isEnabled      bool
	getColor       Color
	printTo        func(Color, *strings.Builder, string)
}

// IsEnabled returns the result of IsEnabledFunc.
func (m *mockPrintable) IsEnabled() bool {
	if m.IsEnabledFunc != nil {
		return m.IsEnabledFunc()
	}
	return false
}

// GetColor returns the result of GetColorFunc.
func (m *mockPrintable) GetColor(value string) Color {
	if m.GetColorFunc != nil {
		return m.GetColorFunc(value)
	}
	return InvalidColor
}

// PrintTo calls the PrintToFunc with the provided arguments.
func (m *mockPrintable) PrintTo(color Color, builder *strings.Builder, a string) {
	if m.PrintToFunc != nil {
		m.PrintToFunc(color, builder, a)
	}
}
