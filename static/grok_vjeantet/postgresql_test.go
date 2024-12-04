package grok_vjeantet

import (
	"testing"
)

// TestPostgresqlHappyPath tests the Postgresql constant with a valid pattern.
func TestPostgresqlHappyPath(t *testing.T) {
	expected := `# Default postgresql pg_log format pattern
POSTGRESQL %{DATESTAMP:timestamp} %{TZ} %{DATA:user_id} %{GREEDYDATA:connection_id} %{POSINT:pid}

`
	if Postgresql != expected {
		t.Errorf("Expected %q, got %q", expected, Postgresql)
	}
}

// TestPostgresqlEmptyString tests the Postgresql constant with an empty string.
func TestPostgresqlEmptyString(t *testing.T) {
	expected := ""
	if Postgresql == expected {
		t.Errorf("Expected not to be %q, got %q", expected, Postgresql)
	}
}

// TestPostgresqlInvalidPattern tests the Postgresql constant with an invalid pattern.
func TestPostgresqlInvalidPattern(t *testing.T) {
	expected := `# Default postgresql pg_log format pattern
POSTGRESQL %{DATESTAMP:timestamp} %{TZ} %{DATA:user_id} %{GREEDYDATA:connection_id} %{POSINT:pid}
INVALID_PATTERN
`
	if Postgresql == expected {
		t.Errorf("Expected not to be %q, got %q", expected, Postgresql)
	}
}

// TestPostgresqlWhitespacePattern tests the Postgresql constant with a pattern containing only whitespace.
func TestPostgresqlWhitespacePattern(t *testing.T) {
	expected := "   "
	if Postgresql == expected {
		t.Errorf("Expected not to be %q, got %q", expected, Postgresql)
	}
}
