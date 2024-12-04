package util

import (
	"testing"
)

func TestMultiString_Set(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected map[string]bool
	}{
		{
			name:     "Happy path",
			input:    "apple,banana,cherry",
			expected: map[string]bool{"apple": true, "banana": true, "cherry": true},
		},
		{
			name:     "Empty string",
			input:    "",
			expected: map[string]bool{},
		},
		{
			name:     "Single value",
			input:    "apple",
			expected: map[string]bool{"apple": true},
		},
		{
			name:     "Whitespace around values",
			input:    "  apple, banana , cherry ",
			expected: map[string]bool{"apple": true, "banana": true, "cherry": true},
		},
		{
			name:     "Duplicate values",
			input:    "apple,apple,banana",
			expected: map[string]bool{"apple": true, "banana": true},
		},
		{
			name:     "Mixed case values",
			input:    "Apple,Banana,cherry",
			expected: map[string]bool{"apple": true, "banana": true, "cherry": true},
		},
		{
			name:     "Special characters in values",
			input:    "apple!@#banana$%^cherry&*()",
			expected: map[string]bool{"apple!@#banana$%^cherry&*()": true},
		},
		{
			name:     "Negative case - non-string input",
			input:    "123,456",
			expected: map[string]bool{},
		},
		{
			name:     "Negative case - empty values",
			input:    ",,,",
			expected: map[string]bool{},
		},
		{
			name:     "Negative case - only whitespace",
			input:    "   ",
			expected: map[string]bool{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			i := &MultiStringT{}
			i.Set(tc.input)

			if !reflect.DeepEqual(i.Values, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, i.Values)
			}
		})
	}
}

func TestMultiString_Contains(t *testing.T) {
	testCases := []struct {
		name     string
		input    MultiString
		value    string
		caseSensitive bool
		expected bool
	}{
		{
			name:     "Happy path - case sensitive",
			input:    &MultiStringT{Values: map[string]bool{"apple": true, "banana": true}},
			value:    "apple",
			caseSensitive: true,
			expected: true,
		},
		{
			name:     "Happy path - case insensitive",
			input:    &MultiStringT{LowercasedValues: map[string]bool{"apple": true, "banana": true}},
			value:    "Apple",
			caseSensitive: false,
			expected: true,
		},
		{
			name:     "Negative case - value not found",
			input:    &MultiStringT{Values: map[string]bool{"apple": true, "banana": true}},
			value:    "cherry",
			caseSensitive: true,
			expected: false,
		},
		{
			name:     "Negative case - empty MultiString",
			input:    &MultiStringT{},
			value:    "apple",
			caseSensitive: true,
			expected: false,
		},
		{
			name:     "Negative case - non-string input",
			input:    &MultiStringT{Values: map[string]bool{"123": true}},
			value:    "456",
			caseSensitive: true,
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.input.Contains(tc.value, tc.caseSensitive)

			if result != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}

func TestMultiString_Reset(t *testing.T) {
	testCases := []struct {
		name     string
		input    MultiString
		expected map[string]bool
	}{
		{
			name:     "Happy path",
			input:    &MultiStringT{Values: map[string]bool{"apple": true, "banana": true}},
			expected: map[string]bool{},
		},
		{
			name:     "Negative case - empty MultiString",
			input:    &MultiStringT{},
			expected: map[string]bool{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.input.Reset()

			if !reflect.DeepEqual(tc.input.Values, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, tc.input.Values)
			}
		})
	}
}

func TestMultiString_UnmarshalYAML(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected map[string]bool
	}{
		{
			name:     "Happy path",
			input:    "apple,banana,cherry",
			expected: map[string]bool{"apple": true, "banana": true, "cherry": true},
		},
		{
			name:     "Empty string",
			input:    "",
			expected: map[string]bool{},
		},
		{
			name:     "Single value",
			input:    "apple",
			expected: map[string]bool{"apple": true},
		},
		{
			name:     "Whitespace around values",
			input:    "  apple, banana , cherry ",
			expected: map[string]bool{"apple": true, "banana": true, "cherry": true},
		},
		{
			name:     "Duplicate values",
			input:    "apple,apple,banana",
			expected: map[string]bool{"apple": true, "banana": true},
		},
		{
			name:     "Mixed case values",
			input:    "Apple,Banana,cherry",
			expected: map[string]bool{"apple": true, "banana": true, "cherry": true},
		},
		{
			name:     "Special characters in input",
			input:    "apple!@#banana$%^cherry&*()",
			expected: map[string]bool{"apple!@#": true, "banana$%^": true, "cherry&*()": true},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			i := &MultiStringT{}
			err := yaml.Unmarshal([]byte(tc.input), i)

			if err != nil {
				t.Errorf("Expected no error, got %v", err)
			}

			if !reflect.DeepEqual(i.Values, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, i.Values)
			}
		})
	}
}

func TestMultiString_MarshalYAML(t *testing.T) {
	testCases := []struct {
		name     string
		input    MultiString
		expected string
	}{
		{
			name:     "Happy path",
			input:    &MultiStringT{Values: map[string]bool{"apple": true, "banana": true}},
			expected: "apple,banana",
		},
		{
			name:     "Empty MultiString",
			input:    &MultiStringT{},
			expected: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := yaml.Marshal(tc.input)

			if err != nil {
				t.Errorf("Expected no error, got %v", err)
			}

			if string(result) != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, result)
			}
		})
	}
}
