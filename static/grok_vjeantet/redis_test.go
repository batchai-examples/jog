package grok_vjeantet

import (
	"testing"
)

// TestRedisHappyPath tests the Redis constant with a valid log entry.
func TestRedisHappyPath(t *testing.T) {
	expected := `REDISTIMESTAMP %{MONTHDAY} %{MONTH} %{TIME}
REDISLOG \[%{POSINT:pid}\] %{REDISTIMESTAMP:timestamp} \* 

`
	result := Redis
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

// TestRedisNegativePath tests the Redis constant with an invalid log entry.
func TestRedisNegativePath(t *testing.T) {
	expected := `INVALIDLOG %{MONTHDAY} %{MONTH} %{TIME}
INVALIDLOG \[%{POSINT:pid}\] %{REDISTIMESTAMP:timestamp} \* 

`
	result := Redis
	if result == expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

// TestRedisCornerCase tests the Redis constant with an empty log entry.
func TestRedisCornerCase(t *testing.T) {
	expected := `

`
	result := Redis
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

// TestRedisEmptyString tests the Redis constant with an empty string.
func TestRedisEmptyString(t *testing.T) {
	expected := ""
	result := Redis
	if result == expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
