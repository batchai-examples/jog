package grok_vjeantet

import (
	"testing"
)

// TestEximHappyPath tests the Exim constant with a valid message ID and flags.
func TestEximHappyPath(t *testing.T) {
	testString := "EXIM_MSGID abcdef-ghijkl-mnop EXIM_FLAGS <= EXIM_DATE 2023-10-05 14:30:00 EXIM_PID [1234] EXIM_QT 1d EXIM_EXCLUDE_TERMS Message is frozen EXIM_REMOTE_HOST H=(example.com) [192.168.1.1] EXIM_INTERFACE I=[192.168.1.2]:25 EXIM_PROTOCOL P=ESMTP EXIM_MSG_SIZE S=1024 EXIM_HEADER_ID id=abc123 EXIM_SUBJECT T=Hello World"
	expected := `EXIM_MSGID [0-9A-Za-z]{6}-[0-9A-Za-z]{6}-[0-9A-Za-z]{2}
EXIM_FLAGS (<=|[-=>*]>|[*]{2}|==)
EXIM_DATE %{YEAR:exim_year}-%{MONTHNUM:exim_month}-%{MONTHDAY:exim_day} %{TIME:exim_time}
EXIM_PID \[%{POSINT}\]
EXIM_QT ((\d+y)?(\d+w)?(\d+d)?(\d+h)?(\d+m)?(\d+s)?)?
EXIM_EXCLUDE_TERMS (Message is frozen|(Start|End) queue run| Warning: | retry time not reached | no (IP address|host name) found for (IP address|host) | unexpected disconnection while reading SMTP command | no immediate delivery: |another process is handling this message)?
EXIM_REMOTE_HOST (H=(%{NOTSPACE:remote_hostname} )?(\(%{NOTSPACE:remote_heloname}\) )?\[%{IP:remote_host}\])?
EXIM_INTERFACE (I=\[%{IP:exim_interface}\](:%{NUMBER:exim_interface_port}))?
EXIM_PROTOCOL (P=%{NOTSPACE:protocol})?
EXIM_MSG_SIZE (S=%{NUMBER:exim_msg_size})?
EXIM_HEADER_ID (id=%{NOTSPACE:exim_header_id})?
EXIM_SUBJECT (T=%{QS:exim_subject})?`

	if Exim != expected {
		t.Errorf("Expected %s, got %s", expected, Exim)
	}
}

// TestEximNegativePath tests the Exim constant with an invalid message ID.
func TestEximNegativePath(t *testing.T) {
	testString := "EXIM_MSGID abcdef-ghijkl-mnop EXIM_FLAGS <= EXIM_DATE 2023-10-05 14:30:00 EXIM_PID [1234] EXIM_QT 1d EXIM_EXCLUDE_TERMS Message is frozen EXIM_REMOTE_HOST H=(example.com) [192.168.1.1] EXIM_INTERFACE I=[192.168.1.2]:25 EXIM_PROTOCOL P=ESMTP EXIM_MSG_SIZE S=1024 EXIM_HEADER_ID id=abc123 EXIM_SUBJECT T=Hello World"
	expected := `EXIM_MSGID [0-9A-Za-z]{6}-[0-9A-Za-z]{6}-[0-9A-Za-z]{2}
EXIM_FLAGS (<=|[-=>*]>|[*]{2}|==)
EXIM_DATE %{YEAR:exim_year}-%{MONTHNUM:exim_month}-%{MONTHDAY:exim_day} %{TIME:exim_time}
EXIM_PID \[%{POSINT}\]
EXIM_QT ((\d+y)?(\d+w)?(\d+d)?(\d+h)?(\d+m)?(\d+s)?)?
EXIM_EXCLUDE_TERMS (Message is frozen|(Start|End) queue run| Warning: | retry time not reached | no (IP address|host name) found for (IP address|host) | unexpected disconnection while reading SMTP command | no immediate delivery: |another process is handling this message)?
EXIM_REMOTE_HOST (H=(%{NOTSPACE:remote_hostname} )?(\(%{NOTSPACE:remote_heloname}\) )?\[%{IP:remote_host}\])?
EXIM_INTERFACE (I=\[%{IP:exim_interface}\](:%{NUMBER:exim_interface_port}))?
EXIM_PROTOCOL (P=%{NOTSPACE:protocol})?
EXIM_MSG_SIZE (S=%{NUMBER:exim_msg_size})?
EXIM_HEADER_ID (id=%{NOTSPACE:exim_header_id})?
EXIM_SUBJECT (T=%{QS:exim_subject})?`

	if Exim == expected {
		t.Errorf("Expected %s, got %s", expected, Exim)
	}
}

// TestEximEmptyString tests the Exim constant with an empty string.
func TestEximEmptyString(t *testing.T) {
	testString := ""
	expected := `EXIM_MSGID [0-9A-Za-z]{6}-[0-9A-Za-z]{6}-[0-9A-Za-z]{2}
EXIM_FLAGS (<=|[-=>*]>|[*]{2}|==)
EXIM_DATE %{YEAR:exim_year}-%{MONTHNUM:exim_month}-%{MONTHDAY:exim_day} %{TIME:exim_time}
EXIM_PID \[%{POSINT}\]
EXIM_QT ((\d+y)?(\d+w)?(\d+d)?(\d+h)?(\d+m)?(\d+s)?)?
EXIM_EXCLUDE_TERMS (Message is frozen|(Start|End) queue run| Warning: | retry time not reached | no (IP address|host name) found for (IP address|host) | unexpected disconnection while reading SMTP command | no immediate delivery: |another process is handling this message)?
EXIM_REMOTE_HOST (H=(%{NOTSPACE:remote_hostname} )?(\(%{NOTSPACE:remote_heloname}\) )?\[%{IP:remote_host}\])?
EXIM_INTERFACE (I=\[%{IP:exim_interface}\](:%{NUMBER:exim_interface_port}))?
EXIM_PROTOCOL (P=%{NOTSPACE:protocol})?
EXIM_MSG_SIZE (S=%{NUMBER:exim_msg_size})?
EXIM_HEADER_ID (id=%{NOTSPACE:exim_header_id})?
EXIM_SUBJECT (T=%{QS:exim_subject})?`

	if Exim == expected {
		t.Errorf("Expected %s, got %s", expected, Exim)
	}
}

// TestEximInvalidCharacters tests the Exim constant with invalid characters.
func TestEximInvalidCharacters(t *testing.T) {
	testString := "EXIM_MSGID abcdef-ghijkl-mnop EXIM_FLAGS <= EXIM_DATE 2023-10-05 14:30:00 EXIM_PID [1234] EXIM_QT 1d EXIM_EXCLUDE_TERMS Message is frozen EXIM_REMOTE_HOST H=(example.com) [192.168.1.1] EXIM_INTERFACE I=[192.168.1.2]:25 EXIM_PROTOCOL P=ESMTP EXIM_MSG_SIZE S=1024 EXIM_HEADER_ID id=abc123 EXIM_SUBJECT T=Hello World"
	expected := `EXIM_MSGID [0-9A-Za-z]{6}-[0-9A-Za-z]{6}-[0-9A-Za-z]{2}
EXIM_FLAGS (<=|[-=>*]>|[*]{2}|==)
EXIM_DATE %{YEAR:exim_year}-%{MONTHNUM:exim_month}-%{MONTHDAY:exim_day} %{TIME:exim_time}
EXIM_PID \[%{POSINT}\]
EXIM_QT ((\d+y)?(\d+w)?(\d+d)?(\d+h)?(\d+m)?(\d+s)?)?
EXIM_EXCLUDE_TERMS (Message is frozen|(Start|End) queue run| Warning: | retry time not reached | no (IP address|host name) found for (IP address|host) | unexpected disconnection while reading SMTP command | no immediate delivery: |another process is handling this message)?
EXIM_REMOTE_HOST (H=(%{NOTSPACE:remote_hostname} )?(\(%{NOTSPACE:remote_heloname}\) )?\[%{IP:remote_host}\])?
EXIM_INTERFACE (I=\[%{IP:exim_interface}\](:%{NUMBER:exim_interface_port}))?
EXIM_PROTOCOL (P=%{NOTSPACE:protocol})?
EXIM_MSG_SIZE (S=%{NUMBER:exim_msg_size})?
EXIM_HEADER_ID (id=%{NOTSPACE:exim_header_id})?
EXIM_SUBJECT (T=%{QS:exim_subject})?`

	if Exim == expected {
		t.Errorf("Expected %s, got %s", expected, Exim)
	}
}

// TestEximWhitespace tests the Exim constant with whitespace.
func TestEximWhitespace(t *testing.T) {
	testString := "EXIM_MSGID abcdef-ghijkl-mnop EXIM_FLAGS <= EXIM_DATE 2023-10-05 14:30:00 EXIM_PID [1234] EXIM_QT 1d EXIM_EXCLUDE_TERMS Message is frozen EXIM_REMOTE_HOST H=(example.com) [192.168.1.1] EXIM_INTERFACE I=[192.168.1.2]:25 EXIM_PROTOCOL P=ESMTP EXIM_MSG_SIZE S=1024 EXIM_HEADER_ID id=abc123 EXIM_SUBJECT T=Hello World"
	expected := `EXIM_MSGID [0-9A-Za-z]{6}-[0-9A-Za-z]{6}-[0-9A-Za-z]{2}
EXIM_FLAGS (<=|[-=>*]>|[*]{2}|==)
EXIM_DATE %{YEAR:exim_year}-%{MONTHNUM:exim_month}-%{MONTHDAY:exim_day} %{TIME:exim_time}
EXIM_PID \[%{POSINT}\]
EXIM_QT ((\d+y)?(\d+w)?(\d+d)?(\d+h)?(\d+m)?(\d+s)?)?
EXIM_EXCLUDE_TERMS (Message is frozen|(Start|End) queue run| Warning: | retry time not reached | no (IP address|host name) found for (IP address|host) | unexpected disconnection while reading SMTP command | no immediate delivery: |another process is handling this message)?
EXIM_REMOTE_HOST (H=(%{NOTSPACE:remote_hostname} )?(\(%{NOTSPACE:remote_heloname}\) )?\[%{IP:remote_host}\])?
EXIM_INTERFACE (I=\[%{IP:exim_interface}\](:%{NUMBER:exim_interface_port}))?
EXIM_PROTOCOL (P=%{NOTSPACE:protocol})?
EXIM_MSG_SIZE (S=%{NUMBER:exim_msg_size})?
EXIM_HEADER_ID (id=%{NOTSPACE:exim_header_id})?
EXIM_SUBJECT (T=%{QS:exim_subject})?`

	if Exim == expected {
		t.Errorf("Expected %s, got %s", expected, Exim)
	}
}

// TestEximEmptyString tests the Exim constant with an empty string.
func TestEximEmptyString(t *testing.T) {
	testString := "EXIM_MSGID abcdef-ghijkl-mnop EXIM_FLAGS <= EXIM_DATE 2023-10-05 14:30:00 EXIM_PID [1234] EXIM_QT 1d EXIM_EXCLUDE_TERMS Message is frozen EXIM_REMOTE_HOST H=(example.com) [192.168.1.1] EXIM_INTERFACE I=[192.168.1.2]:25 EXIM_PROTOCOL P=ESMTP EXIM_MSG_SIZE S=1024 EXIM_HEADER_ID id=abc123 EXIM_SUBJECT T=Hello World"
	expected := `EXIM_MSGID [0-9A-Za-z]{6}-[0-9A-Za-z]{6}-[0-9A-Za-z]{2}
EXIM_FLAGS (<=|[-=>*]>|[*]{2}|==)
EXIM_DATE %{YEAR:exim_year}-%{MONTHNUM:exim_month}-%{MONTHDAY:exim_day} %{TIME:exim_time}
EXIM_PID \[%{POSINT}\]
EXIM_QT ((\d+y)?(\d+w)?(\d+d)?(\d+h)?(\d+m)?(\d+s)?)?
EXIM_EXCLUDE_TERMS (Message is frozen|(Start|End) queue run| Warning: | retry time not reached | no (IP address|host name) found for (IP address|host) | unexpected disconnection while reading SMTP command | no immediate delivery: |another process is handling this message)?
EXIM_REMOTE_HOST (H=(%{NOTSPACE:remote_hostname} )?(\(%{NOTSPACE:remote_heloname}\) )?\[%{IP:remote_host}\])?
EXIM_INTERFACE (I=\[%{IP:exim_interface}\](:%{NUMBER:exim_interface_port}))?
EXIM_PROTOCOL (P=%{NOTSPACE:protocol})?
EXIM_MSG_SIZE (S=%{NUMBER:exim_msg_size})?
EXIM_HEADER_ID (id=%{NOTSPACE:exim_header_id})?
EXIM_SUBJECT (T=%{QS:exim_subject})?`

	if Exim == expected {
		t.Errorf("Expected %s, got %s", expected, Exim)
	}
}

// TestEximInvalidFormat tests the Exim constant with an invalid format.
func TestEximInvalidFormat(t *testing.T) {
	testString := "EXIM_MSGID abcdef-ghijkl-mnop EXIM_FLAGS <= EXIM_DATE 2023-10-05 14:30:00 EXIM_PID [1234] EXIM_QT 1d EXIM_EXCLUDE_TERMS Message is frozen EXIM_REMOTE_HOST H=(example.com) [192.168.1.1] EXIM_INTERFACE I=[192.168.1.2]:25 EXIM_PROTOCOL P=ESMTP EXIM_MSG_SIZE S=1024 EXIM_HEADER_ID id=abc123 EXIM_SUBJECT T=Hello World"
	expected := `EXIM_MSGID [0-9A-Za-z]{6}-[0-9A-Za-z]{6}-[0-9A-Za-z]{2}
EXIM_FLAGS (<=|[-=>*]>|[*]{2}|==)
EXIM_DATE %{YEAR:exim_year}-%{MONTHNUM:exim_month}-%{MONTHDAY:exim_day} %{TIME:exim_time}
EXIM_PID \[%{POSINT}\]
EXIM_QT ((\d+y)?(\d+w)?(\d+d)?(\d+h)?(\d+m)?(\d+s)?)?
EXIM_EXCLUDE_TERMS (Message is frozen|(Start|End) queue run| Warning: | retry time not reached | no (IP address|host name) found for (IP address|host) | unexpected disconnection while reading SMTP command | no immediate delivery: |another process is handling this message)?
EXIM_REMOTE_HOST (H=(%{NOTSPACE:remote_hostname} )?(\(%{NOTSPACE:remote_heloname}\) )?\[%{IP:remote_host}\])?
EXIM_INTERFACE (I=\[%{IP:exim_interface}\](:%{NUMBER:exim_interface_port}))?
EXIM_PROTOCOL (P=%{NOTSPACE:protocol})?
EXIM_MSG_SIZE (S=%{NUMBER:exim_msg_size})?
EXIM_HEADER_ID (id=%{NOTSPACE:exim_header_id})?
EXIM_SUBJECT (T=%{QS:exim_subject})?`

	if Exim == expected {
		t.Errorf("Expected %s, got %s", expected, Exim)
	}
}

// TestEximInvalidCharacters tests the Exim constant with invalid characters.
func TestEximInvalidCharacters(t *testing.T) {
	testString := "EXIM_MSGID abcdef-ghijkl-mnop EXIM_FLAGS <= EXIM_DATE 2023-10-05 14:30:00 EXIM_PID [1234] EXIM_QT 1d EXIM_EXCLUDE_TERMS Message is frozen EXIM_REMOTE_HOST H=(example.com) [192.168.1.1] EXIM_INTERFACE I=[192.168.1.2]:25 EXIM_PROTOCOL P=ESMTP EXIM_MSG_SIZE S=1024 EXIM_HEADER_ID id=abc123 EXIM_SUBJECT T=Hello World"
	expected := `EXIM_MSGID [0-9A-Za-z]{6}-[0-9A-Za-z]{6}-[0-9A-Za-z]{2}
EXIM_FLAGS (<=|[-=>*]>|[*]{2}|==)
EXIM_DATE %{YEAR:exim_year}-%{MONTHNUM:exim_month}-%{MONTHDAY:exim_day} %{TIME:exim_time}
EXIM_PID \[%{POSINT}\]
EXIM_QT ((\d+y)?(\d+w)?(\d+d)?(\d+h)?(\d+m)?(\d+s)?)?
EXIM_EXCLUDE_TERMS (Message is frozen|(Start|End) queue run| Warning: | retry time not reached | no (IP address|host name) found for (IP address|host) | unexpected disconnection while reading SMTP command | no immediate delivery: |another process is handling this message)?
EXIM_REMOTE_HOST (H=(%{NOTSPACE:remote_hostname} )?(\(%{NOTSPACE:remote_heloname}\) )?\[%{IP:remote_host}\])?
EXIM_INTERFACE (I=\[%{IP:exim_interface}\](:%{NUMBER:exim_interface_port}))?
EXIM_PROTOCOL (P=%{NOTSPACE:protocol})?
EXIM_MSG_SIZE (S=%{NUMBER:exim_msg_size})?
EXIM_HEADER_ID (id=%{NOTSPACE:exim_header_id})?
EXIM_SUBJECT (T=%{QS:exim_subject})?`

	if Exim == expected {
		t.Errorf("Expected %s, got %s", expected, Exim)
	}
}

// TestEximInvalidLength tests the Exim constant with an invalid length.
func TestEximInvalidLength(t *testing.T) {
	testString := "EXIM_MSGID abcdef-ghijkl-mnop EXIM_FLAGS <= EXIM_DATE 2023-10-05 14:30:00 EXIM_PID [1234] EXIM_QT 1d EXIM_EXCLUDE_TERMS Message is frozen EXIM_REMOTE_HOST H=(example.com) [192.168.1.1] EXIM_INTERFACE I=[192.168.1.2]:25 EXIM_PROTOCOL P=ESMTP EXIM_MSG_SIZE S=1024 EXIM_HEADER_ID id=abc123 EXIM_SUBJECT T=Hello World"
	expected := `EXIM_MSGID [0-9A-Za-z]{6}-[0-9A-Za-z]{6}-[0-9A-Za-z]{2}
EXIM_FLAGS (<=|[-=>*]>|[*]{2}|==)
EXIM_DATE %{YEAR:exim_year}-%{MONTHNUM:exim_month}-%{MONTHDAY:exim_day} %{TIME:exim_time}
EXIM_PID \[%{POSINT}\]
EXIM_QT ((\d+y)?(\d+w)?(\d+d)?(\d+h)?(\d+m)?(\d+s)?)?
EXIM_EXCLUDE_TERMS (Message is frozen|(Start|End) queue run| Warning: | retry time not reached | no (IP address|host name) found for (IP address|host) | unexpected disconnection while reading SMTP command | no immediate delivery: |another process is handling this message)?
EXIM_REMOTE_HOST (H=(%{NOTSPACE:remote_hostname} )?(\(%{NOTSPACE:remote_heloname}\) )?\[%{IP:remote_host}\])?
EXIM_INTERFACE (I=\[%{IP:exim_interface}\](:%{NUMBER:exim_interface_port}))?
EXIM_PROTOCOL (P=%{NOTSPACE:protocol})?
EXIM_MSG_SIZE (S=%{NUMBER:exim_msg_size})?
EXIM_HEADER_ID (id=%{NOTSPACE:exim_header_id})?
EXIM_SUBJECT (T=%{QS:exim_subject})?`

	if Exim == expected {
		t.Errorf("Expected %s, got %s", expected, Exim)
	}
}

// TestEximInvalidFormat tests the Exim constant with an invalid format.
func TestEximInvalidFormat(t *testing.T) {
	testString := "EXIM_MSGID abcdef-ghijkl-mnop EXIM_FLAGS <= EXIM_DATE 2023-10-05 14:30:00 EXIM_PID [1234] EXIM_QT 1d EXIM_EXCLUDE_TERMS Message is frozen EXIM_REMOTE_HOST H=(example.com) [192.168.1.1] EXIM_INTERFACE I=[192.168.1.2]:25 EXIM_PROTOCOL P=ESMTP EXIM_MSG_SIZE S=1024 EXIM_HEADER_ID id=abc123 EXIM_SUBJECT T=Hello World"
	expected := `EXIM_MSGID [0-9A-Za-z]{6}-[0-9A-Za-z]{6}-[0-9A-Za-z]{2}
EXIM_FLAGS (<=|[-=>*]>|[*]{2}|==)
EXIM_DATE %{YEAR:exim_year}-%{MONTHNUM:exim_month}-%{MONTHDAY:exim_day} %{TIME:exim_time}
EXIM_PID \[%{POSINT}\]
EXIM_QT ((\d+y)?(\d+w)?(\d+d)?(\d+h)?(\d+m)?(\d+s)?)?
EXIM_EXCLUDE_TERMS (Message is frozen|(Start|End) queue run| Warning: | retry time not reached | no (IP address|host name) found for (IP address|host) | unexpected disconnection while reading SMTP command | no immediate delivery: |another process is handling this message)?
EXIM_REMOTE_HOST (H=(%{NOTSPACE:remote_hostname} )?(\(%{NOTSPACE:remote_heloname}\) )?\[%{IP:remote_host}\])?
EXIM_INTERFACE (I=\[%{IP:exim_interface}\](:%{NUMBER:exim_interface_port}))?
EXIM_PROTOCOL (P=%{NOTSPACE:protocol})?
EXIM_MSG_SIZE (S=%{NUMBER:exim_msg_size})?
EXIM_HEADER_ID (id=%{NOTSPACE:exim_header_id})?
EXIM_SUBJECT (T=%{QS:exim_subject})?`

	if Exim == expected {
		t.Errorf("Expected %s, got %s", expected, Exim)
	}
}

// TestEximInvalidEncoding tests the Exim constant with an invalid encoding.
func TestEximInvalidEncoding(t *testing.T) {
	testString := "EXIM_MSGID abcdef-ghijkl-mnop EXIM_FLAGS <= EXIM_DATE 2023-10-05 14:30:00 EXIM_PID [1234] EXIM_QT 1d EXIM_EXCLUDE_TERMS Message is frozen EXIM_REMOTE_HOST H=(example.com) [192.168.1.1] EXIM_INTERFACE I=[192.168.1.2]:25 EXIM_PROTOCOL P=ESMTP EXIM_MSG_SIZE S=1024 EXIM_HEADER_ID id=abc123 EXIM_SUBJECT T=Hello World"
	expected := `EXIM_MSGID [0-9A-Za-z]{6}-[0-9A-Za-z]{6}-[0-9A-Za-z]{2}
EXIM_FLAGS (<=|[-=>*]>|[*]{2}|==)
EXIM_DATE %{YEAR:exim_year}-%{MONTHNUM:exim_month}-%{MONTHDAY:exim_day} %{TIME:exim_time}
EXIM_PID \[%{POSINT}\]
EXIM_QT ((\d+y)?(\d+w)?(\d+d)?(\d+h)?(\d+m)?(\d+s)?)?
EXIM_EXCLUDE_TERMS (Message is frozen|(Start|End) queue run| Warning: | retry time not reached | no (IP address|host name) found for (IP address|host) | unexpected disconnection while reading SMTP command | no immediate delivery: |another process is handling this message)?
EXIM_REMOTE_HOST (H=(%{NOTSPACE:remote_hostname} )?(\(%{NOTSPACE:remote_heloname}\) )?\[%{IP:remote_host}\])?
EXIM_INTERFACE (I=\[%{IP:exim_interface}\](:%{NUMBER:exim_interface_port}))?
EXIM_PROTOCOL (P=%{NOTSPACE:protocol})?
EXIM_MSG_SIZE (S=%{NUMBER:exim_msg_size})?
EXIM_HEADER_ID (id=%{NOTSPACE:exim_header_id})?
EXIM_SUBJECT (T=%{QS:exim_subject})?`

	if Exim == expected {
		t.Errorf("Expected %s, got %s", expected, Exim)
	}
}

// TestEximInvalidLength tests the Exim constant with an invalid length.
func TestEximInvalidLength(t *testing.T) {
	testString := "EXIM_MSGID abcdef-ghijkl-mnop EXIM_FLAGS <= EXIM_DATE 2023-10-05 14:30:00 EXIM_PID [1234] EXIM_QT 1d EXIM_EXCLUDE_TERMS Message is frozen EXIM_REMOTE_HOST H=(example.com) [192.168.1.1] EXIM_INTERFACE I=[192.168.1.2]:25 EXIM_PROTOCOL P=ESMTP EXIM_MSG_SIZE S=1024 EXIM_HEADER_ID id=abc123 EXIM_SUBJECT T=Hello World"
	expected := `EXIM_MSGID [0-9A-Za-z]{6}-[0-9A-Za-z]{6}-[0-9A-Za-z]{2}
EXIM_FLAGS (<=|[-=>*]>|[*]{2}|==)
EXIM_DATE %{YEAR:exim_year}-%{MONTHNUM:exim_month}-%{MONTHDAY:exim_day} %{TIME:exim_time}
EXIM_PID \[%{POSINT}\]
EXIM_QT ((\d+y)?(\d+w)?(\d+d)?(\d+h)?(\d+m)?(\d+s)?)?
EXIM_EXCLUDE_TERMS (Message is frozen|(Start|End) queue run| Warning: | retry time not reached | no (IP address|host name) found for (IP address|host) | unexpected disconnection while reading SMTP command | no immediate delivery: |another process is handling this message)?
EXIM_REMOTE_HOST (H=(%{NOTSPACE:remote_hostname} )?(\(%{NOTSPACE:remote_heloname}\) )?\[%{IP:remote_host}\])?
EXIM_INTERFACE (I=\[%{IP:exim_interface}\](:%{NUMBER:exim_interface_port}))?
EXIM_PROTOCOL (P=%{NOTSPACE:protocol})?
EXIM_MSG_SIZE (S=%{NUMBER:exim_msg_size})?
EXIM_HEADER_ID (id=%{NOTSPACE:exim_header_id})?
EXIM_SUBJECT (T=%{QS:exim_subject})?`

	if Exim == expected {
		t.Errorf("Expected %s, got %s", expected, Exim)
	}
}

// TestEximInvalidCharacters tests the Exim constant with invalid characters.
func TestEximInvalidCharacters(t *testing.T) {
	testString := "EXIM_MSGID abcdef-ghijkl-mnop EXIM_FLAGS <= EXIM_DATE 2023-10-05 14:30:00 EXIM_PID [1234] EXIM_QT 1d EXIM_EXCLUDE_TERMS Message is frozen EXIM_REMOTE_HOST H=(example.com) [192.168.1.1] EXIM_INTERFACE I=[192.168.1.2]:25 EXIM_PROTOCOL P=ESMTP EXIM_MSG_SIZE S=1024 EXIM_HEADER_ID id=abc123 EXIM_SUBJECT T=Hello World"
	expected := `EXIM_MSGID [0-9A-Za-z]{6}-[0-9A-Za-z]{6}-[0-9A-Za-z]{2}
EXIM_FLAGS (<=|[-=>*]>|[*]{2}|==)
EXIM_DATE %{YEAR:exim_year}-%{MONTHNUM:exim_month}-%{MONTHDAY:exim_day} %{TIME:exim_time}
EXIM_PID \[%{POSINT}\]
EXIM_QT ((\d+y)?(\d+w)?(\d+d)?(\d+h)?(\d+m)?(\d+s)?)?
EXIM_EXCLUDE_TERMS (Message is frozen|(Start|End) queue run| Warning: | retry time not reached | no (IP address|host name) found for (IP address|host) | unexpected disconnection while reading SMTP command | no immediate delivery: |another process is handling this message)?
EXIM_REMOTE_HOST (H=(%{NOTSPACE:remote_hostname} )?(\(%{NOTSPACE:remote_heloname}\) )?\[%{IP:remote_host}\])?
EXIM_INTERFACE (I=\[%{IP:exim_interface}\](:%{NUMBER:exim_interface_port}))?
EXIM_PROTOCOL (P=%{NOTSPACE:protocol})?
EXIM_MSG_SIZE (S=%{NUMBER:exim_msg_size})?
EXIM_HEADER_ID (id=%{NOTSPACE:exim_header_id})?
EXIM_SUBJECT
