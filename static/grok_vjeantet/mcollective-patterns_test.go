package grok_vjeantet

import (
	"testing"
)

// TestMcollectivePatternsHappyPath tests the Mcollective_patterns constant with a valid multi-line event.
func TestMcollectivePatternsHappyPath(t *testing.T) {
	expected := `# Remember, these can be multi-line events.
MCOLLECTIVE ., [%{TIMESTAMP_ISO8601:timestamp} #%{POSINT:pid}\]%{SPACE}%{LOGLEVEL:event_level}

MCOLLECTIVEAUDIT %{TIMESTAMP_ISO8601:timestamp}:
`
	if Mcollective_patterns != expected {
		t.Errorf("Expected %q, but got %q", expected, Mcollective_patterns)
	}
}

// TestMcollectivePatternsEmptyString tests the Mcollective_patterns constant with an empty string.
func TestMcollectivePatternsEmptyString(t *testing.T) {
	expected := ""
	if Mcollective_patterns == expected {
		t.Errorf("Expected non-empty string, but got %q", Mcollective_patterns)
	}
}

// TestMcollectivePatternsTrailingWhitespace tests the Mcollective_patterns constant with trailing whitespace.
func TestMcollectivePatternsTrailingWhitespace(t *testing.T) {
	expected := `# Remember, these can be multi-line events.
MCOLLECTIVE ., [%{TIMESTAMP_ISO8601:timestamp} #%{POSINT:pid}\]%{SPACE}%{LOGLEVEL:event_level}

MCOLLECTIVEAUDIT %{TIMESTAMP_ISO8601:timestamp}:
`
	Mcollective_patterns += " "
	if Mcollective_patterns != expected {
		t.Errorf("Expected %q, but got %q", expected, Mcollective_patterns)
	}
}

// TestMcollectivePatternsLeadingWhitespace tests the Mcollective_patterns constant with leading whitespace.
func TestMcollectivePatternsLeadingWhitespace(t *testing.T) {
	expected := `# Remember, these can be multi-line events.
MCOLLECTIVE ., [%{TIMESTAMP_ISO8601:timestamp} #%{POSINT:pid}\]%{SPACE}%{LOGLEVEL:event_level}

MCOLLECTIVEAUDIT %{TIMESTAMP_ISO8601:timestamp}:`
	Mcollective_patterns = " " + Mcollective_patterns
	if Mcollective_patterns != expected {
		t.Errorf("Expected %q, but got %q", expected, Mcollective_patterns)
	}
}
