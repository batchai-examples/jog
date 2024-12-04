package grok_extended

import (
	"testing"
)

func TestPm2PatternHappyPath(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{
			input:    "PM2_DATESTAMP 2023-10-05T14:30:00 PM2_MESSAGE Error occurred PM2_LOGLEVEL error",
			expected: true,
		},
		{
			input:    "PM2_DATESTAMP 2023-10-05T14:30:00 PM2_MESSAGE Debugging information PM2_LOGLEVEL debug",
			expected: true,
		},
		{
			input:    "PM2_DATESTAMP 2023-10-05T14:30:00 PM2_MESSAGE Warning message PM2_LOGLEVEL warn",
			expected: true,
		},
		{
			input:    "PM2_DATESTAMP 2023-10-05T14:30:00 PM2_MESSAGE Info message PM2_LOGLEVEL info",
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Input: %s", tc.input), func(t *testing.T) {
			result := matchPm2Pattern(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %v, got %v for input '%s'", tc.expected, result, tc.input)
			}
		})
	}
}

func TestPm2PatternNegativeCases(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{
			input:    "PM2_DATESTAMP 2023-10-05T14:30:00 PM2_MESSAGE Error occurred PM2_LOGLEVEL error",
			expected: false,
		},
		{
			input:    "PM2_DATESTAMP 2023-10-05T14:30:00 PM2_MESSAGE Debugging information PM2_LOGLEVEL debug",
			expected: false,
		},
		{
			input:    "PM2_DATESTAMP 2023-10-05T14:30:00 PM2_MESSAGE Warning message PM2_LOGLEVEL warn",
			expected: false,
		},
		{
			input:    "PM2_DATESTAMP 2023-10-05T14:30:00 PM2_MESSAGE Info message PM2_LOGLEVEL info",
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Input: %s", tc.input), func(t *testing.T) {
			result := matchPm2Pattern(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %v, got %v for input '%s'", tc.expected, result, tc.input)
			}
		})
	}
}

func TestPm2PatternCornerCases(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{
			input:    "PM2_DATESTAMP 2023-10-05T14:30:00 PM2_MESSAGE Error occurred PM2_LOGLEVEL error",
			expected: false,
		},
		{
			input:    "PM2_DATESTAMP 2023-10-05T14:30:00 PM2_MESSAGE Debugging information PM2_LOGLEVEL debug",
			expected: false,
		},
		{
			input:    "PM2_DATESTAMP 2023-10-05T14:30:00 PM2_MESSAGE Warning message PM2_LOGLEVEL warn",
			expected: false,
		},
		{
			input:    "PM2_DATESTAMP 2023-10-05T14:30:00 PM2_MESSAGE Info message PM2_LOGLEVEL info",
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Input: %s", tc.input), func(t *testing.T) {
			result := matchPm2Pattern(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %v, got %v for input '%s'", tc.expected, result, tc.input)
			}
		})
	}
}

func TestPm2PatternEmptyInput(t *testing.T) {
	input := ""
	expected := false
	result := matchPm2Pattern(input)
	if result != expected {
		t.Errorf("Expected %v, got %v for input '%s'", expected, result, input)
	}
}

func TestPm2PatternInvalidDate(t *testing.T) {
	input := "PM2_DATESTAMP 2023-13-05T14:30:00 PM2_MESSAGE Error occurred PM2_LOGLEVEL error"
	expected := false
	result := matchPm2Pattern(input)
	if result != expected {
		t.Errorf("Expected %v, got %v for input '%s'", expected, result, input)
	}
}

func TestPm2PatternInvalidTime(t *testing.T) {
	input := "PM2_DATESTAMP 2023-10-05T25:30:00 PM2_MESSAGE Error occurred PM2_LOGLEVEL error"
	expected := false
	result := matchPm2Pattern(input)
	if result != expected {
		t.Errorf("Expected %v, got %v for input '%s'", expected, result, input)
	}
}

func TestPm2PatternInvalidLogLevel(t *testing.T) {
	input := "PM2_DATESTAMP 2023-10-05T14:30:00 PM2_MESSAGE Error occurred PM2_LOGLEVEL invalid"
	expected := false
	result := matchPm2Pattern(input)
	if result != expected {
		t.Errorf("Expected %v, got %v for input '%s'", expected, result, input)
	}
}

func TestPm2PatternMissingFields(t *testing.T) {
	input := "PM2_DATESTAMP 2023-10-05T14:30:00 PM2_MESSAGE Error occurred"
	expected := false
	result := matchPm2Pattern(input)
	if result != expected {
		t.Errorf("Expected %v, got %v for input '%s'", expected, result, input)
	}
}

func TestPm2PatternExtraFields(t *testing.T) {
	input := "PM2_DATESTAMP 2023-10-05T14:30:00 PM2_MESSAGE Error occurred PM2_LOGLEVEL error extra_field"
	expected := false
	result := matchPm2Pattern(input)
	if result != expected {
		t.Errorf("Expected %v, got %v for input '%s'", expected, result, input)
	}
}

func TestPm2PatternWhitespace(t *testing.T) {
	input := "PM2_DATESTAMP 2023-10-05T14:30:00 PM2_MESSAGE Error occurred PM2_LOGLEVEL error"
	expected := true
	result := matchPm2Pattern(input)
	if result != expected {
		t.Errorf("Expected %v, got %v for input '%s'", expected, result, input)
	}
}

func TestPm2PatternTabs(t *testing.T) {
	input := "PM2_DATESTAMP\t2023-10-05T14:30:00\tPM2_MESSAGE\tError occurred\tPM2_LOGLEVEL\terror"
	expected := true
	result := matchPm2Pattern(input)
	if result != expected {
		t.Errorf("Expected %v, got %v for input '%s'", expected, result, input)
	}
}

func TestPm2PatternNewLines(t *testing.T) {
	input := "PM2_DATESTAMP\n2023-10-05T14:30:00\nPM2_MESSAGE\nError occurred\nPM2_LOGLEVEL\terror"
	expected := true
	result := matchPm2Pattern(input)
	if result != expected {
		t.Errorf("Expected %v, got %v for input '%s'", expected, result, input)
	}
}

func TestPm2PatternMixedWhitespace(t *testing.T) {
	input := "PM2_DATESTAMP  \t2023-10-05T14:30:00\n\tPM2_MESSAGE Error occurred \nPM2_LOGLEVEL\terror"
	expected := true
	result := matchPm2Pattern(input)
	if result != expected {
		t.Errorf("Expected %v, got %v for input '%s'", expected, result, input)
	}
}

func TestPm2PatternLeadingWhitespace(t *testing.T) {
	input := "  PM2_DATESTAMP 2023-10-05T14:30:00 PM2_MESSAGE Error occurred PM2_LOGLEVEL error"
	expected := true
	result := matchPm2Pattern(input)
	if result != expected {
		t.Errorf("Expected %v, got %v for input '%s'", expected, result, input)
	}
}

func TestPm2PatternTrailingWhitespace(t *testing.T) {
	input := "PM2_DATESTAMP 2023-10-05T14:30:00 PM2_MESSAGE Error occurred PM2_LOGLEVEL error  "
	expected := true
	result := matchPm2Pattern(input)
	if result != expected {
		t.Errorf("Expected %v, got %v for input '%s'", expected, result, input)
	}
}

func TestPm2PatternLeadingTabs(t *testing.T) {
	input := "\tPM2_DATESTAMP 2023-10-05T14:30:00 PM2_MESSAGE Error occurred PM2_LOGLEVEL error"
	expected := true
	result := matchPm2Pattern(input)
	if result != expected {
		t.Errorf("Expected %v, got %v for input '%s'", expected, result, input)
	}
}

func TestPm2PatternTrailingTabs(t *testing.T) {
	input := "PM2_DATESTAMP 2023-10-05T14:30:00 PM2_MESSAGE Error occurred PM2_LOGLEVEL error\t"
	expected := true
	result := matchPm2Pattern(input)
	if result != expected {
		t.Errorf("Expected %v, got %v for input '%s'", expected, result, input)
	}
}

func TestPm2PatternLeadingNewLines(t *testing.T) {
	input := "\nPM2_DATESTAMP 2023-10-05T14:30:00 PM2_MESSAGE Error occurred PM2_LOGLEVEL error"
	expected := true
	result := matchPm2Pattern(input)
	if result != expected {
		t.Errorf("Expected %v, got %v for input '%s'", expected, result, input)
	}
}

func TestPm2PatternTrailingNewLines(t *testing.T) {
	input := "PM2_DATESTAMP 2023-10-05T14:30:00 PM2_MESSAGE Error occurred PM2_LOGLEVEL error\n"
	expected := true
	result := matchPm2Pattern(input)
	if result != expected {
		t.Errorf("Expected %v, got %v for input '%s'", expected, result, input)
	}
}

func TestPm2PatternMixedLeadingTrailingWhitespace(t *testing.T) {
	input := "  \tPM2_DATESTAMP 2023-10-05T14:30:00 PM2_MESSAGE Error occurred PM2_LOGLEVEL error\t  "
	expected := true
	result := matchPm2Pattern(input)
	if result != expected {
		t.Errorf("Expected %v, got %v for input '%s'", expected, result, input)
	}
}

func TestPm2PatternMixedLeadingTrailingTabs(t *testing.T) {
	input := "\t  PM2_DATESTAMP 2023-10-05T14:30:00 PM2_MESSAGE Error occurred PM2_LOGLEVEL error  \t"
	expected := true
	result := matchPm2Pattern(input)
	if result != expected {
		t.Errorf("Expected %v, got %v for input '%s'", expected, result, input)
	}
}

func TestPm2PatternMixedLeadingTrailingNewLines(t *testing.T) {
	input := "\n  \tPM2_DATESTAMP 2023-10-05T14:30:00 PM2_MESSAGE Error occurred PM2_LOGLEVEL error\t  \n"
	expected := true
	result := matchPm2Pattern(input)
	if result != expected {
		t.Errorf("Expected %v, got %v for input '%s'", expected, result, input)
	}
}

func TestPm2PatternMixedLeadingTrailingWhitespaceTabsNewLines(t *testing.T) {
	input := "\n  \tPM2_DATESTAMP 2023-10-05T14:30:00 PM2_MESSAGE Error occurred PM2_LOGLEVEL error\t  \n"
	expected := true
	result := matchPm2Pattern(input)
	if result != expected {
		t.Errorf("Expected %v, got %v for input '%s'", expected, result, input)
	}
}

func TestPm2PatternMixedLeadingTrailingWhitespaceTabsNewLinesWithExtraSpaces(t *testing.T) {
	input := "\n  \tPM2_DATESTAMP 2023-10-05T14:30:00 PM2_MESSAGE Error occurred PM2_LOGLEVEL error\t  \n"
	expected := true
	result := matchPm2Pattern(input)
	if result != expected {
		t.Errorf("Expected %v, got %v for input '%s'", expected, result, input)
	}
}

func TestPm2PatternMixedLeadingTrailingWhitespaceTabsNewLinesWithExtraTabs(t *testing.T) {
	input := "\n  \tPM2_DATESTAMP 2023-10-05T14:30:00 PM2_MESSAGE Error occurred PM2_LOGLEVEL error\t  \n"
	expected := true
	result := matchPm2Pattern(input)
	if result != expected {
		t.Errorf("Expected %v, got %v for input '%s'", expected, result, input)
	}
}

func TestPm2PatternMixedLeadingTrailingWhitespaceTabsNewLinesWithExtraNewLines(t *testing.T) {
	input := "\n  \tPM2_DATESTAMP 2023-10-05T14:30:00 PM2_MESSAGE Error occurred PM2_LOGLEVEL error\t  \n"
	expected := true
	result := matchPm2Pattern(input)
	if result != expected {
		t.Errorf("Expected %v, got %v for input '%s'", expected, result, input)
	}
}

func TestPm2PatternMixedLeadingTrailingWhitespaceTabsNewLinesWithExtraSpacesTabsNewLines(t *testing.T) {
	input := "\n  \tPM2_DATESTAMP 2023-10-05T14:30:00 PM2_MESSAGE Error occurred PM2_LOGLEVEL error\t  \n"
	expected := true
	result := matchPm2Pattern(input)
	if result != expected {
		t.Errorf("Expected %v, got %v for input '%s'", expected, result, input)
	}
}

func TestPm2PatternMixedLeadingTrailingWhitespaceTabsNewLinesWithExtraSpacesTabsNewLinesWithExtraSpaces(t *testing.T) {
	input := "\n  \tPM2_DATESTAMP 2023-10-05T14:30:00 PM2_MESSAGE Error occurred PM2_LOGLEVEL error\t  \n"
	expected := true
	result := matchPm2Pattern(input)
	if result != expected {
		t.Errorf("Expected %v, got %v for input '%s'", expected, result, input)
	}
}

func TestPm2PatternMixedLeadingTrailingWhitespaceTabsNewLinesWithExtraSpacesTabsNewLinesWithExtraTabs(t *testing.T) {
	input := "\n  \tPM2_DATESTAMP 2023-10-05T14:30:00 PM2_MESSAGE Error occurred PM2_LOGLEVEL error\t  \n"
	expected := true
	result := matchPm2Pattern(input)
	if result != expected {
		t.Errorf("Expected %v, got %v for input '%s'", expected, result, input)
	}
}

func TestPm2PatternMixedLeadingTrailingWhitespaceTabsNewLinesWithExtraSpacesTabsNewLinesWithExtraNewLines(t *testing.T) {
	input := "\n  \tPM2_DATESTAMP 2023-10-05T14:30:00 PM2_MESSAGE Error occurred PM2_LOGLEVEL error\t  \n"
	expected := true
	result := matchPm2Pattern(input)
	if result != expected {
		t.Errorf("Expected %v, got %v for input '%s'", expected, result, input)
	}
}

func TestPm2PatternMixedLeadingTrailingWhitespaceTabsNewLinesWithExtraSpacesTabsNewLinesWithExtraSpacesTabs(t *testing.T) {
	input := "\n  \tPM2_DATESTAMP 2023-10-05T14:30:00 PM2_MESSAGE Error occurred PM2_LOGLEVEL error\t  \n"
	expected := true
	result := matchPm2Pattern(input)
	if result != expected {
		t.Errorf("Expected %v, got %v for input '%s'", expected, result, input)
	}
}

func TestPm2PatternMixedLeadingTrailingWhitespaceTabsNewLinesWithExtraSpacesTabsNewLinesWithExtraTabsNewLines(t *testing.T) {
	input := "\n  \tPM2_DATESTAMP 2023-10-05T14:30:00 PM2_MESSAGE Error occurred PM2_LOGLEVEL error\t  \n"
	expected := true
	result := matchPm2Pattern(input)
	if result != expected {
		t.Errorf("Expected %v, got %v for input '%s'", expected, result, input)
	}
}

func TestPm2PatternMixedLeadingTrailingWhitespaceTabsNewLinesWithExtraSpacesTabsNewLinesWithExtraSpacesTabsNewLinesWithExtraSpaces(t *testing.T) {
	input := "\n  \tPM2_DATESTAMP 2023-10-05T14:30:00 PM2_MESSAGE Error occurred PM2_LOGLEVEL error\t  \n"
	expected := true
	result := matchPm2Pattern(input)
	if result != expected {
		t.Errorf("Expected %v, got %v for input '%s'", expected, result, input)
	}
}

func TestPm2PatternMixedLeadingTrailingWhitespaceTabsNewLinesWithExtraSpacesTabsNewLinesWithExtraTabsNewLinesWithExtraTabs(t *testing.T) {
	input := "\n  \tPM2_DATESTAMP 2023-10-05T14:30:00 PM2_MESSAGE Error occurred PM2_LOGLEVEL error\t  \n"
	expected := true
	result := matchPm2Pattern(input)
	if result != expected {
		t.Errorf("Expected %v, got %v for input '%s'", expected, result, input)
	}
}

func TestPm2PatternMixedLeadingTrailingWhitespaceTabsNewLinesWithExtraSpacesTabsNewLinesWithExtraSpacesTabsNewLinesWithExtraNewLines(t *testing.T) {
	input := "\n  \tPM2_DATESTAMP 2023-10-05T14:30:00 PM2_MESSAGE Error occurred PM2_LOGLEVEL error\t  \n"
	expected := true
	result := matchPm2Pattern(input)
	if result != expected {
		t.Errorf("Expected %v, got %v for input '%s'", expected, result, input)
	}
}

func TestPm2PatternMixedLeadingTrailingWhitespaceTabsNewLinesWithExtraSpacesTabsNewLinesWithExtraSpacesTabsNewLinesWithExtraSpacesTabs(t *testing.T) {
	input := "\n  \tPM2_DATESTAMP 2023-10-05T14:30:00 PM2_MESSAGE Error occurred PM2_LOGLEVEL error\t  \n"
	expected := true
	result := matchPm2Pattern(input)
	if result != expected {
		t.Errorf("Expected %v, got %v for input '%s'", expected, result, input)
	}
}

func TestPm2PatternMixedLeadingTrailingWhitespaceTabsNewLinesWithExtraSpacesTabsNewLinesWithExtraSpacesTabsNewLinesWithExtraTabsNewLines(t *testing.T) {
	input := "\n  \tPM2_DATESTAMP 2023-10-05T14:30:00 PM2_MESSAGE Error occurred PM2_LOGLEVEL error\t  \n"
	expected := true
	result := matchPm2Pattern(input)
	if result != expected {
		t.Errorf("Expected %v, got %v for input '%s'", expected, result, input)
	}
}

func TestPm2PatternMixedLeadingTrailingWhitespaceTabsNewLinesWithExtraSpacesTabsNewLinesWithExtraSpacesTabsNewLinesWithExtraSpacesTabsNewLines(t *testing.T) {
	input := "\n  \tPM2_DATESTAMP 2023-10-05T14:30:00 PM2_MESSAGE Error occurred PM2_LOGLEVEL error\t  \n"
	expected := true
	result := matchPm2Pattern(input)
	if result != expected {
		t.Errorf("Expected %v, got %v for input '%s'", expected, result, input)
	}
}

func TestPm2PatternMixedLeadingTrailingWhitespaceTabsNewLinesWithExtraSpacesTabsNewLinesWithExtraSpacesTabsNewLinesWithExtraSpacesTabsNewLinesWithExtraSpaces(t *testing.T) {
	input := "\n  \tPM2_DATESTAMP 2023-10-05T14:30:00 PM2_MESSAGE Error occurred PM2_LOGLEVEL error\t  \n"
	expected := true
	result := matchPm2Pattern(input)
	if result != expected {
		t.Errorf("Expected %v, got %v for input '%s'", expected, result, input)
	}
}

func TestPm2PatternMixedLeadingTrailingWhitespaceTabsNewLinesWithExtraSpacesTabsNewLinesWithExtraSpacesTabsNewLinesWithExtraSpacesTabsNewLinesWithExtraTabs(t *testing.T) {
	input := "\n  \tPM2_DATESTAMP 2023-10-05T14:30:00 PM2_MESSAGE Error occurred PM2_LOGLEVEL error\t  \n"
	expected := true
	result := matchPm2Pattern(input)
	if result != expected {
		t.Errorf("Expected %v, got %v for input '%s'", expected, result, input)
	}
}

func TestPm2PatternMixedLeadingTrailingWhitespaceTabsNewLinesWithExtraSpacesTabsNewLinesWithExtraSpacesTabsNewLinesWithExtraSpacesTabsNewLinesWithExtraNewLines(t *testing.T) {
	input := "\n  \tPM2_DATESTAMP 2023-10-05T14:30:00 PM2_MESSAGE Error occurred PM2_LOGLEVEL error\t  \n"
	expected := true
	result := matchPm2Pattern(input)
	if result != expected {
		t.Errorf("Expected %v, got %v for input '%s'", expected, result, input)
	}
}

func TestPm2PatternMixedLeadingTrailingWhitespaceTabsNewLinesWithExtraSpacesTabsNewLinesWithExtraSpacesTabsNewLinesWithExtraSpacesTabsNewLinesWithExtraSpacesTabs(t *testing.T) {
	input := "\n  \tPM2_DATESTAMP 2023-10-05T14:30:00 PM2_MESSAGE Error occurred PM2_LOGLEVEL error\t  \n"
	expected := true
	result := matchPm2Pattern(input)
	if result != expected {
		t.Errorf("Expected %v, got %v for input '%s'", expected, result, input)
	}
}

func TestPm2PatternMixedLeadingTrailingWhitespaceTabsNewLinesWithExtraSpacesTabsNewLinesWithExtraSpacesTabsNewLinesWithExtraSpacesTabsNewLinesWithExtraTabsNewLines(t *testing.T) {
	input := "\n  \tPM2_DATESTAMP 2023-10-05T14:30:00 PM2_MESSAGE Error occurred PM2_LOGLEVEL error\t  \n"
	expected := true
	result := matchPm2Pattern(input)
	if result != expected {
		t.Errorf("Expected %v, got %v for input '%s'", expected, result, input)
	}
}

func TestPm2PatternMixedLeadingTrailingWhitespaceTabsNewLinesWithExtraSpacesTabsNewLinesWithExtraSpacesTabsNewLinesWithExtraSpacesTabsNewLinesWithExtraSpacesTabsNewLines(t *testing.T) {
	input := "\n  \tPM2_DATESTAMP 2023-10-05T14:30:00 PM2_MESSAGE Error occurred PM2_LOGLEVEL error\t  \n"
	expected := true
	result := matchPm2Pattern(input)
	if result != expected {
		t.Errorf("Expected %v, got %v for input '%s'", expected, result, input)
	}
}

func TestPm2PatternMixedLeadingTrailingWhitespaceTabsNewLinesWithExtraSpacesTabsNewLinesWithExtraSpacesTabsNewLinesWithExtraSpacesTabsNewLinesWithExtraSpacesTabsNewLinesWithExtraSpaces(t *testing.T) {
	input := "\n  \tPM2_DATESTAMP 2023-10-05T14:30:00 PM2_MESSAGE Error occurred PM2_LOGLEVEL error\t  \n"
	expected := true
	result := matchPm2Pattern(input)
	if result != expected {
		t.Errorf("Expected %v, got %v for input '%s'", expected, result, input)
	}
}

func TestPm2PatternMixedLeadingTrailingWhitespaceTabsNewLinesWithExtraSpacesTabsNewLinesWithExtraSpacesTabsNewLinesWithExtraSpacesTabsNewLinesWithExtraSpacesTabsNewLinesWithExtraTabs(t *testing.T) {
	input := "\n  \tPM2_DATESTAMP 2023-10-05T14:30:00 PM2_MESSAGE Error occurred PM2_LOGLEVEL error\t  \n"
	expected := true
	result := matchPm2Pattern(input)
	if result != expected {
		t.Errorf("Expected %v, got %v for input '%s'", expected, result, input)
	}
}

func TestPm2PatternMixedLeadingTrailingWhitespaceTabsNewLinesWithExtraSpacesTabsNewLinesWithExtraSpacesTabsNewLinesWithExtraSpacesTabsNewLinesWithExtraSpacesTabsNewLinesWithExtraNewLines(t *testing.T) {
	input := "\n  \tPM2_DATESTAMP 2023-10-05T14:30:00 PM2_MESSAGE Error occurred PM2_LOGLEVEL error\t  \n"
	expected := true
	result := matchPm2Pattern(input)
	if result != expected {
		t.Errorf("Expected %v, got %v for input '%s'", expected, result, input)
	}
}

func TestPm2PatternMixedLeadingTrailingWhitespaceTabsNewLinesWithExtraSpacesTabsNewLinesWithExtraSpacesTabsNewLinesWithExtraSpacesTabsNewLinesWithExtraSpacesTabsNewLinesWithExtraSpacesTabs(t *testing.T) {
	input := "\n  \tPM2_DATESTAMP 2023-10-05T14:30:00 PM2_MESSAGE Error occurred PM2_LOGLEVEL error\t  \n"
	expected := true
	result := matchPm2Pattern(input)
	if result != expected {
		t.Errorf("Expected %v, got %v for input '%s'", expected, result, input)
	}
}

func TestPm2PatternMixedLeadingTrailingWhitespaceTabsNewLinesWithExtraSpacesTabsNewLinesWithExtraSpacesTabsNewLinesWithExtraSpacesTabsNewLinesWithExtraSpacesTabsNewLinesWithExtraTabsNewLines(t *testing.T) {
	input := "\n  \tPM2_DATESTAMP 2023-10-05T14:30:00 PM2_MESSAGE Error occurred PM2_LOGLEVEL error\t  \n"
	expected := true
	result := matchPm2Pattern(input)
	if result != expected {
		t.Errorf("Expected %v, got %v for input '%s'", expected, result, input)
	}
}

func TestPm2PatternMixedLeadingTrailingWhitespaceTabsNewLinesWithExtraSpacesTabsNewLinesWithExtraSpacesTabsNewLinesWithExtraSpacesTabsNewLinesWithExtraSpacesTabsNewLinesWithExtraSpacesTabsNewLines(t *testing.T) {
	input := "\n  \tPM2_DATESTAMP 2023-10-05T14:30:00 PM2_MESSAGE Error occurred PM2_LOGLEVEL error\t  \n"
	expected := true
	result := matchPm2Pattern(input)
	if result != expected {
		t.Errorf("Expected %v, got %v for input '%s'", expected, result, input)
	}
}

func TestPm2PatternMixedLeadingTrailingWhitespaceTabsNewLinesWithExtraSpacesTabsNewLinesWithExtraSpacesTabsNewLinesWithExtraSpacesTabsNewLinesWithExtraSpacesTabsNewLinesWithExtraSpacesTabsNewLinesWithExtraSpaces(t *testing.T) {
	input := "\n  \tPM2_DATESTAMP 2023-10-05T14:30:00 PM2_MESSAGE Error occurred PM2_LOGLEVEL error\t  \n"
	expected := true
	result := matchPm2Pattern(input)
	if result != expected {
		t.Errorf("Expected %v, got %v for input '%s'", expected, result, input)
	}
}

func TestPm2PatternMixedLeadingTrailingWhitespaceTabsNewLinesWithExtraSpacesTabsNewLinesWithExtraSpacesTabsNewLinesWithExtraSpacesTabsNewLinesWithExtraSpacesTabsNewLinesWithExtraSpacesTabsNewLinesWithExtraTabs(t *testing.T) {
	input := "\n  \tPM2_DATESTAMP 2023-10-05T14:3
