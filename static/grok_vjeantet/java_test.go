package grok_vjeantet

import (
	"testing"
)

func TestJavaRegex(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Happy path with Java stack trace",
			input:    "Jan 9, 2014 7:13:13 AM at com.example.service.ExampleService.method(com.example.service.ExampleService.java:123)",
			expected: true,
		},
		{
			name:     "Happy path with Java log message",
			input:    "Jan 9, 2014 7:13:13 AM com.example.service.ExampleService - something completely unexpected happened...",
			expected: true,
		},
		{
			name:     "Negative case: Missing class name in stack trace",
			input:    "Jan 9, 2014 7:13:13 AM at .method(com.example.service.ExampleService.java:123)",
			expected: false,
		},
		{
			name:     "Negative case: Missing method name in stack trace",
			input:    "Jan 9, 2014 7:13:13 AM at com.example.service.ExampleService.(com.example.service.ExampleService.java:123)",
			expected: false,
		},
		{
			name:     "Negative case: Missing file name in stack trace",
			input:    "Jan 9, 2014 7:13:13 AM at com.example.service.ExampleService.method(com.example.service..java:123)",
			expected: false,
		},
		{
			name:     "Negative case: Missing line number in stack trace",
			input:    "Jan 9, 2014 7:13:13 AM at com.example.service.ExampleService.method(com.example.service.ExampleService.java::)",
			expected: false,
		},
		{
			name:     "Negative case: Invalid date format in stack trace",
			input:    "Jan 9, 2014 7:13:13 PM at com.example.service.ExampleService.method(com.example.service.ExampleService.java:123)",
			expected: false,
		},
		{
			name:     "Negative case: Invalid log level in log message",
			input:    "Jan 9, 2014 7:13:13 AM | INVALID | com.example.service.ExampleService - something completely unexpected happened...",
			expected: false,
		},
		{
			name:     "Corner case: Empty input string",
			input:    "",
			expected: false,
		},
		{
			name:     "Corner case: Only date stamp in stack trace",
			input:    "Jan 9, 2014 7:13:13 AM",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := JavaRegex.MatchString(tt.input)
			if result != tt.expected {
				t.Errorf("JavaRegex(%s) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}
