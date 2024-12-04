package static

import (
	"testing"
)

// TestAppVersion checks if the AppVersion constant is correctly defined.
func TestAppVersion(t *testing.T) {
	testCases := []struct {
		name     string
		expected string
	}{
		{
			name:     "Happy path",
			expected: `v1.0.3`,
		},
		{
			name:     "Negative case - empty version",
			expected: ``,
		},
		{
			name:     "Corner case - trailing newline",
			expected: `v1.0.3\n`,
		},
		{
			name:     "Corner case - leading whitespace",
			expected: `\tv1.0.3\n`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if AppVersion != tc.expected {
				t.Errorf("Expected %q, got %q", tc.expected, AppVersion)
			}
		})
	}
}
