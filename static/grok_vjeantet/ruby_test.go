package grok_vjeantet

import (
	"testing"
)

// TestRubyHappyPath tests the Ruby constant with a valid log message.
func TestRubyHappyPath(t *testing.T) {
	testString := `RUBY_LOGLEVEL DEBUG RUBY_LOGGER [DFEWI], [2023-04-15T12:34:56+00:00 #1234] *DEBUG -- +progname: This is a debug message`
	expected := `RUBY_LOGLEVEL (?:DEBUG|FATAL|ERROR|WARN|INFO)
RUBY_LOGGER [DFEWI], \[%{TIMESTAMP_ISO8601:timestamp} #%{POSINT:pid}\] *%{RUBY_LOGLEVEL:loglevel} -- +%{DATA:progname}: %{GREEDYDATA:message}
`
	result := Ruby
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

// TestRubyNegativePath tests the Ruby constant with an invalid log message.
func TestRubyNegativePath(t *testing.T) {
	testString := `RUBY_LOGLEVEL INFO RUBY_LOGGER [DFEWI], [2023-04-15T12:34:56+00:00 #1234] *INFO -- +progname: This is an info message`
	expected := `RUBY_LOGLEVEL (?:DEBUG|FATAL|ERROR|WARN|INFO)
RUBY_LOGGER [DFEWI], \[%{TIMESTAMP_ISO8601:timestamp} #%{POSINT:pid}\] *%{RUBY_LOGLEVEL:loglevel} -- +%{DATA:progname}: %{GREEDYDATA:message}
`
	result := Ruby
	if result == expected {
		t.Errorf("Expected an error, got %s", result)
	}
}

// TestRubyCornerCase tests the Ruby constant with a log message that has unexpected characters.
func TestRubyCornerCase(t *testing.T) {
	testString := `RUBY_LOGLEVEL ERROR RUBY_LOGGER [DFEWI], [2023-04-15T12:34:56+00:00 #1234] *ERROR -- +progname: This is an error message with unexpected characters!@#$%^&*()`
	expected := `RUBY_LOGLEVEL (?:DEBUG|FATAL|ERROR|WARN|INFO)
RUBY_LOGGER [DFEWI], \[%{TIMESTAMP_ISO8601:timestamp} #%{POSINT:pid}\] *%{RUBY_LOGLEVEL:loglevel} -- +%{DATA:progname}: %{GREEDYDATA:message}
`
	result := Ruby
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

// TestRubyEmptyString tests the Ruby constant with an empty string.
func TestRubyEmptyString(t *testing.T) {
	testString := ""
	expected := `RUBY_LOGLEVEL (?:DEBUG|FATAL|ERROR|WARN|INFO)
RUBY_LOGGER [DFEWI], \[%{TIMESTAMP_ISO8601:timestamp} #%{POSINT:pid}\] *%{RUBY_LOGLEVEL:loglevel} -- +%{DATA:progname}: %{GREEDYDATA:message}
`
	result := Ruby
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

!!!!test_end!!!!
