package grok_vjeantet

import (
	"testing"
)

func TestRailsPattern(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{
			input:    "Started GET \"/users/show\" for 127.0.0.1 at 2023-04-15 12:34:56 +0000 UTC",
			expected: true,
		},
		{
			input:    "Processing by UsersController as HTML Parameters: {:action=>\"show\", :controller=>\"users\"}",
			expected: true,
		},
		{
			input:    "Completed 200 OK in 15ms (Views: 5ms | ActiveRecord: 10ms)",
			expected: true,
		},
		{
			input:    "Started POST \"/users/create\" for 192.168.1.1 at 2023-04-15 12:34:57 +0000 UTC",
			expected: true,
		},
		{
			input:    "Processing by UsersController as JSON Parameters: {:action=>\"create\", :controller=>\"users\"}",
			expected: true,
		},
		{
			input:    "Completed 404 Not Found in 20ms (Views: 10ms | ActiveRecord: 10ms)",
			expected: true,
		},
		{
			input:    "Started GET \"/users\" for 127.0.0.1 at 2023-04-15 12:34:58 +0000 UTC",
			expected: true,
		},
		{
			input:    "Processing by UsersController as HTML Parameters: {:action=>\"index\", :controller=>\"users\"}",
			expected: true,
		},
		{
			input:    "Completed 200 OK in 10ms (Views: 5ms | ActiveRecord: 5ms)",
			expected: true,
		},
		{
			input:    "Started DELETE \"/users/1\" for 192.168.1.1 at 2023-04-15 12:34:59 +0000 UTC",
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Input: %s", tc.input), func(t *testing.T) {
			result := Rails.MatchString(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %v, got %v for input '%s'", tc.expected, result, tc.input)
			}
		})
	}
}
