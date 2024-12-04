package config

import (
	"testing"
)

func TestCompressPrefixActionString_HappyPath(t *testing.T) {
	// Test the String method for each CompressPrefixAction value
	testCases := []struct {
		name     string
		input    CompressPrefixAction
		expected string
	}{
		{"RemoveNonFirstLetter", CompressPrefixActionRemoveNonFirstLetter, "remove-non-first-letter"},
		{"Remove", CompressPrefixActionRemove, "remove"},
		{"Default", CompressPrefixActionDefault, "remove-non-first-letter"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.input.String()
			if result != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, result)
			}
		})
	}
}

func TestCompressPrefixActionString_NegativePath(t *testing.T) {
	// Test the String method with an invalid CompressPrefixAction value
	testCases := []struct {
		name     string
		input    CompressPrefixAction
		expected string
	}{
		{"InvalidValue", CompressPrefixAction(100), ""},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.input.String()
			if result != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, result)
			}
		})
	}
}

func TestParseCompressPrefixAction_HappyPath(t *testing.T) {
	// Test the ParseCompressPrefixAction function with valid input strings
	testCases := []struct {
		name     string
		input    string
		expected CompressPrefixAction
	}{
		{"RemoveNonFirstLetter", "remove-non-first-letter", CompressPrefixActionRemoveNonFirstLetter},
		{"Remove", "remove", CompressPrefixActionRemove},
		{"Default", "default", CompressPrefixActionDefault},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := ParseCompressPrefixAction(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}

func TestParseCompressPrefixAction_NegativePath(t *testing.T) {
	// Test the ParseCompressPrefixAction function with invalid input strings
	testCases := []struct {
		name     string
		input    string
		expected CompressPrefixAction
	}{
		{"Unknown", "unknown", CompressPrefixActionDefault},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := ParseCompressPrefixAction(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}

func TestCompressPrefixT_Init_HappyPath(t *testing.T) {
	// Test the Init method with default values
	testCases := []struct {
		name     string
		input    CompressPrefixT
		expected CompressPrefixT
	}{
		{"Default", CompressPrefixT{}, CompressPrefixT{Enabled: false, Separators: StringSet{}, WhiteList: StringSet{}, Action: CompressPrefixActionRemoveNonFirstLetter}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.input.Init()
			if tc.input.Enabled != tc.expected.Enabled || !tc.input.Separators.Equal(tc.expected.Separators) || !tc.input.WhiteList.Equal(tc.expected.WhiteList) || tc.input.Action != tc.expected.Action {
				t.Errorf("Expected %+v, got %+v", tc.expected, tc.input)
			}
		})
	}
}

func TestCompressPrefixT_Init_NegativePath(t *testing.T) {
	// Test the Init method with invalid values
	testCases := []struct {
		name     string
		input    CompressPrefixT
		expected CompressPrefixT
	}{
		{"InvalidAction", CompressPrefixT{Action: 100}, CompressPrefixT{Enabled: false, Separators: StringSet{}, WhiteList: StringSet{}, Action: CompressPrefixActionRemoveNonFirstLetter}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.input.Init()
			if tc.input.Enabled != tc.expected.Enabled || !tc.input.Separators.Equal(tc.expected.Separators) || !tc.input.WhiteList.Equal(tc.expected.WhiteList) || tc.input.Action != tc.expected.Action {
				t.Errorf("Expected %+v, got %+v", tc.expected, tc.input)
			}
		})
	}
}

func TestCompressPrefixT_Compress_HappyPath(t *testing.T) {
	// Test the Compress method with different input strings and actions
	testCases := []struct {
		name     string
		input    CompressPrefixT
		text     string
		expected string
	}{
		{"RemoveNonFirstLetter", CompressPrefixT{Action: CompressPrefixActionRemoveNonFirstLetter}, "hello-world", "h-w"},
		{"Remove", CompressPrefixT{Action: CompressPrefixActionRemove}, "hello-world", "world"},
		{"NoAction", CompressPrefixT{}, "hello-world", "hello-world"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.input.Compress(tc.text)
			if result != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, result)
			}
		})
	}
}

func TestCompressPrefixT_Compress_NegativePath(t *testing.T) {
	// Test the Compress method with invalid input strings
	testCases := []struct {
		name     string
		input    CompressPrefixT
		text     string
		expected string
	}{
		{"EmptyString", CompressPrefixT{Action: CompressPrefixActionRemoveNonFirstLetter}, "", ""},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.input.Compress(tc.text)
			if result != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, result)
			}
		})
	}
}

func TestCompressPrefixT_Compress_WithSeparators(t *testing.T) {
	// Test the Compress method with separators
	testCases := []struct {
		name     string
		input    CompressPrefixT
		text     string
		expected string
	}{
		{"RemoveNonFirstLetterWithSeparator", CompressPrefixT{Action: CompressPrefixActionRemoveNonFirstLetter, Separators: StringSet{"-"}}, "hello-world", "h-w"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.input.Compress(tc.text)
			if result != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, result)
			}
		})
	}
}

func TestCompressPrefixT_Compress_WithWhiteList(t *testing.T) {
	// Test the Compress method with white list
	testCases := []struct {
		name     string
		input    CompressPrefixT
		text     string
		expected string
	}{
		{"RemoveNonFirstLetterWithWhiteList", CompressPrefixT{Action: CompressPrefixActionRemoveNonFirstLetter, WhiteList: StringSet{"hello"}}, "hello-world", "h-w"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.input.Compress(tc.text)
			if result != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, result)
			}
		})
	}
}

func TestCompressPrefixT_Compress_WithBothSeparatorsAndWhiteList(t *testing.T) {
	// Test the Compress method with both separators and white list
	testCases := []struct {
		name     string
		input    CompressPrefixT
		text     string
		expected string
	}{
		{"RemoveNonFirstLetterWithBoth", CompressPrefixT{Action: CompressPrefixActionRemoveNonFirstLetter, Separators: StringSet{"-"}, WhiteList: StringSet{"hello"}}, "hello-world", "h-w"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.input.Compress(tc.text)
			if result != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, result)
			}
		})
	}
}
