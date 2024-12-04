package grok_vjeantet

import (
	"testing"
)

// TestMcollectiveHappyPath tests the Mcollective constant with a happy path scenario.
func TestMcollectiveHappyPath(t *testing.T) {
	expected := `MCOLLECTIVEAUDIT %{TIMESTAMP_ISO8601:timestamp}:
`
	if Mcollective != expected {
		t.Errorf("Expected %s, got %s", expected, Mcollective)
	}
}

// TestMcollectivePositiveCase tests the Mcollective constant with a positive case scenario.
func TestMcollectivePositiveCase(t *testing.T) {
	expected := `MCOLLECTIVEAUDIT %{TIMESTAMP_ISO8601:timestamp}:
`
	if Mcollective != expected {
		t.Errorf("Expected %s, got %s", expected, Mcollective)
	}
}

// TestMcollectiveNegativeCase tests the Mcollective constant with a negative case scenario.
func TestMcollectiveNegativeCase(t *testing.T) {
	expected := `ANOTHER_VALUE`
	if Mcollective == expected {
		t.Errorf("Expected not to be %s", expected)
	}
}

// TestMcollectiveCornerCase tests the Mcollective constant with a corner case scenario.
func TestMcollectiveCornerCase(t *testing.T) {
	expected := ""
	if Mcollective != expected {
		t.Errorf("Expected %s, got %s", expected, Mcollective)
	}
}
