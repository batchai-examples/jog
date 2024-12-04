package grok_vjeantet

import (
	"testing"
)

func TestLinuxSyslog(t *testing.T) {
	testCases := []struct {
		input    string
		expected map[string]string
	}{
		{
			input: "Jan 1 00:00:00 hostname syslog5424PRINTASCII message",
			expected: map[string]string{
				"timestamp":     "Jan 1 00:00:00",
				"logsource":   "hostname",
				"message":       "syslog5424PRINTASCII message",
			},
		},
		{
			input: "Jan 1 00:00:00 hostname syslog5424PRINTASCII message with special characters !@#$%^&*()_+",
			expected: map[string]string{
				"timestamp":     "Jan 1 00:00:00",
				"logsource":   "hostname",
				"message":       "syslog5424PRINTASCII message with special characters !@#$%^&*()_+",
			},
		},
		{
			input: "Jan 1 00:00:00 hostname syslog5424PRINTASCII message with spaces in it",
			expected: map[string]string{
				"timestamp":     "Jan 1 00:00:00",
				"logsource":   "hostname",
				"message":       "syslog5424PRINTASCII message with spaces in it",
			},
		},
		{
			input: "Jan 1 00:00:00 hostname syslog5424PRINTASCII message with special characters and spaces !@#$%^&*()_+ message with spaces in it",
			expected: map[string]string{
				"timestamp":     "Jan 1 00:00:00",
				"logsource":   "hostname",
				"message":       "syslog5424PRINTASCII message with special characters and spaces !@#$%^&*()_+ message with spaces in it",
			},
		},
		{
			input: "Jan 1 00:00:00 hostname syslog5424PRINTASCII message with empty string",
			expected: map[string]string{
				"timestamp":     "Jan 1 00:00:00",
				"logsource":   "hostname",
				"message":       "",
			},
		},
		{
			input: "Jan 1 00:00:00 hostname syslog5424PRINTASCII message with null character \x00",
			expected: map[string]string{
				"timestamp":     "Jan 1 00:00:00",
				"logsource":   "hostname",
				"message":       "",
			},
		},
		{
			input: "Jan 1 00:00:00 hostname syslog5424PRINTASCII message with non-printable characters \x01\x02\x03\x04\x05",
			expected: map[string]string{
				"timestamp":     "Jan 1 00:00:00",
				"logsource":   "hostname",
				"message":       "",
			},
		},
		{
			input: "Jan 1 00:00:00 hostname syslog5424PRINTASCII message with non-ascii characters éàùô",
			expected: map[string]string{
				"timestamp":     "Jan 1 00:00:00",
				"logsource":   "hostname",
				"message":       "syslog5424PRINTASCII message with non-ascii characters éàùô",
			},
		},
		{
			input: "Jan 1 00:00:00 hostname syslog5424PRINTASCII message with tabs \t\t",
			expected: map[string]string{
				"timestamp":     "Jan 1 00:00:00",
				"logsource":   "hostname",
				"message":       "syslog5424PRINTASCII message with tabs \t\t",
			},
		},
		{
			input: "Jan 1 00:00:00 hostname syslog5424PRINTASCII message with new lines \n\n",
			expected: map[string]string{
				"timestamp":     "Jan 1 00:00:00",
				"logsource":   "hostname",
				"message":       "",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := ParseLinuxSyslog(tc.input)
			if result != nil && len(result) > 0 {
				for key, value := range tc.expected {
					if result[key] != value {
						t.Errorf("Expected %s to be %s but got %s", key, value, result[key])
					}
				}
			} else {
				t.Errorf("Expected non-nil and non-empty map but got nil or empty map")
			}
		})
	}
}

func TestParseLinuxSyslog(t *testing.T) {
	testCases := []struct {
		input    string
		expected map[string]string
	}{
		{
			input: "Jan 1 00:00:00 hostname syslog5424PRINTASCII message",
			expected: map[string]string{
				"timestamp":     "Jan 1 00:00:00",
				"logsource":   "hostname",
				"message":       "syslog5424PRINTASCII message",
			},
		},
		{
			input: "Jan 1 00:00:00 hostname syslog5424PRINTASCII message with spaces in it",
			expected: map[string]string{
				"timestamp":     "Jan 1 00:00:00",
				"logsource":   "hostname",
				"message":       "syslog5424PRINTASCII message with spaces in it",
			},
		},
		{
			input: "Jan 1 00:00:00 hostname syslog5424PRINTASCII message with special characters and spaces !@#$%^&*()_+ message with spaces in it",
			expected: map[string]string{
				"timestamp":     "Jan 1 00:00:00",
				"logsource":   "hostname",
				"message":       "syslog5424PRINTASCII message with special characters and spaces !@#$%^&*()_+ message with spaces in it",
			},
		},
		{
			input: "Jan 1 00:00:00 hostname syslog5424PRINTASCII message with empty string",
			expected: map[string]string{
				"timestamp":     "Jan 1 00:00:00",
				"logsource":   "hostname",
				"message":       "",
			},
		},
		{
			input: "Jan 1 00:00:00 hostname syslog5424PRINTASCII message with null character \x00",
			expected: map[string]string{
				"timestamp":     "Jan 1 00:00:00",
				"logsource":   "hostname",
				"message":       "",
			},
		},
		{
			input: "Jan 1 00:00:00 hostname syslog5424PRINTASCII message with non-printable characters \x01\x02\x03\x04\x05",
			expected: map[string]string{
				"timestamp":     "Jan 1 00:00:00",
				"logsource":   "hostname",
				"message":       "",
			},
		},
		{
			input: "Jan 1 00:00:00 hostname syslog5424PRINTASCII message with non-ascii characters éàùô",
			expected: map[string]string{
				"timestamp":     "Jan 1 00:00:00",
				"logsource":   "hostname",
				"message":       "syslog5424PRINTASCII message with non-ascii characters éàùô",
			},
		},
		{
			input: "Jan 1 00:00:00 hostname syslog5424PRINTASCII message with tabs \t\t",
			expected: map[string]string{
				"timestamp":     "Jan 1 00:00:00",
				"logsource":   "hostname",
				"message":       "syslog5424PRINTASCII message with tabs \t\t",
			},
		},
		{
			input: "Jan 1 00:00:00 hostname syslog5424PRINTASCII message with new lines \n\n",
			expected: map[string]string{
				"timestamp":     "Jan 1 00:00:00",
				"logsource":   "hostname",
				"message":       "",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := ParseLinuxSyslog(tc.input)
			if result != nil && len(result) > 0 {
				for key, value := range tc.expected {
					if result[key] != value {
						t.Errorf("Expected %s to be %s but got %s", key, value, result[key])
					}
				}
			} else {
				t.Errorf("Expected non-nil and non-empty map but got nil or empty map")
			}
		})
	}
}
