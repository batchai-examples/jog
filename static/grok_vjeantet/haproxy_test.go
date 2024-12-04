package grok_vjeantet

import (
	"testing"
)

func TestHaproxyHTTPBase(t *testing.T) {
	testCases := []struct {
		input    string
		expected map[string]string
	}{
		{
			input: `192.168.1.1:50000 [10/Oct/2023:14:30:00.123] frontend_name backend_name/server_name 100/50/10/20/0.001 200 1234 "cookie" "cookie" "closed" 100/100/100/100/0 0/0/0 {"request_header_host":"host","request_header_x_forwarded_for":"x-forwarded-for","request_header_accept_language":"accept-language","request_header_referer":"referer","request_header_user_agent":"user-agent"} "GET / HTTP/1.1"`,
			expected: map[string]string{
				"client_ip":                  "192.168.1.1",
				"client_port":                "50000",
				"accept_date":                "10/Oct/2023:14:30:00.123",
				"frontend_name":              "frontend_name",
				"backend_name":               "backend_name",
				"server_name":                "server_name",
				"time_request":               "100",
				"time_queue":                 "50",
				"time_backend_connect":       "10",
				"time_backend_response":      "20",
				"time_duration":              "0.001",
				"http_status_code":           "200",
				"bytes_read":                 "1234",
				"captured_request_cookie":    "cookie",
				"captured_response_cookie":   "cookie",
				"termination_state":          "closed",
				"actconn":                    "100",
				"feconn":                     "100",
				"beconn":                     "100",
				"srvconn":                    "100",
				"retries":                    "0",
				"srv_queue":                  "0",
				"backend_queue":              "0",
				"captured_request_headers":   `{"request_header_host":"host","request_header_x_forwarded_for":"x-forwarded-for","request_header_accept_language":"accept-language","request_header_referer":"referer","request_header_user_agent":"user-agent"}`,
				"captured_response_headers":  "",
				"http_verb":                  "GET",
				"http_proto":                 "HTTP",
				"http_host":                  "/",
				"http_request":               "/ HTTP/1.1",
				"http_version":               "1.1",
			},
		},
		{
			input: `192.168.1.1:50000 [10/Oct/2023:14:30:00.123] frontend_name backend_name/server_name 100/50/10/20/0.001 200 1234 "cookie" "cookie" "closed" 100/100/100/100/0 0/0/0 {"request_header_host":"host","request_header_x_forwarded_for":"x-forwarded-for","request_header_accept_language":"accept-language","request_header_referer":"referer","request_header_user_agent":"user-agent"}`,
			expected: map[string]string{
				"client_ip":                  "192.168.1.1",
				"client_port":                "50000",
				"accept_date":                "10/Oct/2023:14:30:00.123",
				"frontend_name":              "frontend_name",
				"backend_name":               "backend_name",
				"server_name":                "server_name",
				"time_request":               "100",
				"time_queue":                 "50",
				"time_backend_connect":       "10",
				"time_backend_response":      "20",
				"time_duration":              "0.001",
				"http_status_code":           "200",
				"bytes_read":                 "1234",
				"captured_request_cookie":    "cookie",
				"captured_response_cookie":   "cookie",
				"termination_state":          "closed",
				"actconn":                    "100",
				"feconn":                     "100",
				"beconn":                     "100",
				"srvconn":                    "100",
				"retries":                    "0",
				"srv_queue":                  "0",
				"backend_queue":              "0",
				"captured_request_headers":   `{"request_header_host":"host","request_header_x_forwarded_for":"x-forwarded-for","request_header_accept_language":"accept-language","request_header_referer":"referer","request_header_user_agent":"user-agent"}`,
				"captured_response_headers":  "",
				"http_verb":                  "GET",
				"http_proto":                 "HTTP",
				"http_host":                  "/",
				"http_request":               "/ HTTP/1.1",
				"http_version":               "1.1",
			},
		},
		{
			input: `192.168.1.1:50000 [10/Oct/2023:14:30:00.123] frontend_name backend_name/server_name 100/50/10/20/0.001 200 1234 "cookie" "cookie" "closed" 100/100/100/100/0 0/0/0 {"request_header_host":"host","request_header_x_forwarded_for":"x-forwarded-for","request_header_accept_language":"accept-language","request_header_referer":"referer","request_header_user_agent":"user-agent"}`,
			expected: map[string]string{
				"client_ip":                  "192.168.1.1",
				"client_port":                "50000",
				"accept_date":                "10/Oct/2023:14:30:00.123",
				"frontend_name":              "frontend_name",
				"backend_name":               "backend_name",
				"server_name":                "server_name",
				"time_request":               "100",
				"time_queue":                 "50",
				"time_backend_connect":       "10",
				"time_backend_response":      "20",
				"time_duration":              "0.001",
				"http_status_code":           "200",
				"bytes_read":                 "1234",
				"captured_request_cookie":    "cookie",
				"captured_response_cookie":   "cookie",
				"termination_state":          "closed",
				"actconn":                    "100",
				"feconn":                     "100",
				"beconn":                     "100",
				"srvconn":                    "100",
				"retries":                    "0",
				"srv_queue":                  "0",
				"backend_queue":              "0",
				"captured_request_headers":   `{"request_header_host":"host","request_header_x_forwarded_for":"x-forwarded-for","request_header_accept_language":"accept-language","request_header_referer":"referer","request_header_user_agent":"user-agent"}`,
				"captured_response_headers":  "",
				"http_verb":                  "GET",
				"http_proto":                 "HTTP",
				"http_host":                  "/",
				"http_request":               "/ HTTP/1.1",
				"http_version":               "1.1",
			},
		},
		{
			input: `192.168.1.1:50000 [10/Oct/2023:14:30:00.123] frontend_name backend_name/server_name 100/50/10/20/0.001 200 1234 "cookie" "cookie" "closed" 100/100/100/100/0 0/0/0 {"request_header_host":"host","request_header_x_forwarded_for":"x-forwarded-for","request_header_accept_language":"accept-language","request_header_referer":"referer","request_header_user_agent":"user-agent"}`,
			expected: map[string]string{
				"client_ip":                  "192.168.1.1",
				"client_port":                "50000",
				"accept_date":                "10/Oct/2023:14:30:00.123",
				"frontend_name":              "frontend_name",
				"backend_name":               "backend_name",
				"server_name":                "server_name",
				"time_request":               "100",
				"time_queue":                 "50",
				"time_backend_connect":       "10",
				"time_backend_response":      "20",
				"time_duration":              "0.001",
				"http_status_code":           "200",
				"bytes_read":                 "1234",
				"captured_request_cookie":    "cookie",
				"captured_response_cookie":   "cookie",
				"termination_state":          "closed",
				"actconn":                    "100",
				"feconn":                     "100",
				"beconn":                     "100",
				"srvconn":                    "100",
				"retries":                    "0",
				"srv_queue":                  "0",
				"backend_queue":              "0",
				"captured_request_headers":   `{"request_header_host":"host","request_header_x_forwarded_for":"x-forwarded-for","request_header_accept_language":"accept-language","request_header_referer":"referer","request_header_user_agent":"user-agent"}`,
				"captured_response_headers":  "",
				"http_verb":                  "GET",
				"http_proto":                 "HTTP",
				"http_host":                  "/",
				"http_request":               "/ HTTP/1.1",
				"http_version":               "1.1",
			},
		},
		{
			input: `192.168.1.1:50000 [10/Oct/2023:14:30:00.123] frontend_name backend_name/server_name 100/50/10/20/0.001 200 1234 "cookie" "cookie" "closed" 100/100/100/100/0 0/0/0 {"request_header_host":"host","request_header_x_forwarded_for":"x-forwarded-for","request_header_accept_language":"accept-language","request_header_referer":"referer","request_header_user_agent":"user-agent"}`,
			expected: map[string]string{
				"client_ip":                  "192.168.1.1",
				"client_port":                "50000",
				"accept_date":                "10/Oct/2023:14:30:00.123",
				"frontend_name":              "frontend_name",
				"backend_name":               "backend_name",
				"server_name":                "server_name",
				"time_request":               "100",
				"time_queue":                 "50",
				"time_backend_connect":       "10",
				"time_backend_response":      "20",
				"time_duration":              "0.001",
				"http_status_code":           "200",
				"bytes_read":                 "1234",
				"captured_request_cookie":    "cookie",
				"captured_response_cookie":   "cookie",
				"termination_state":          "closed",
				"actconn":                    "100",
				"feconn":                     "100",
				"beconn":                     "100",
				"srvconn":                    "100",
				"retries":                    "0",
				"srv_queue":                  "0",
				"backend_queue":              "0",
				"captured_request_headers":   `{"request_header_host":"host","request_header_x_forwarded_for":"x-forwarded-for","request_header_accept_language":"accept-language","request_header_referer":"referer","request_header_user_agent":"user-agent"}`,
				"captured_response_headers":  "",
				"http_verb":                  "GET",
				"http_proto":                 "HTTP",
				"http_host":                  "/",
				"http_request":               "/ HTTP/1.1",
				"http_version":               "1.1",
			},
		},
	]
	for _, tc := range testCases {
		result, err := parseLogLine(tc.input)
		if err != nil {
			t.Errorf("Error parsing log line: %v", err)
		}
		if result.clientIP != tc.expected.clientIP ||
			result.clientPort != tc.expected.clientPort ||
			result.acceptDate != tc.expected.acceptDate ||
			result.frontendName != tc.expected.frontendName ||
			result.backendName != tc.expected.backendName ||
			result.serverName != tc.expected.serverName ||
			result.timeRequest != tc.expected.timeRequest ||
			result.timeQueue != tc.expected.timeQueue ||
			result.timeBackendConnect != tc.expected.timeBackendConnect ||
			result.timeBackendResponse != tc.expected.timeBackendResponse ||
			result.timeDuration != tc.expected.timeDuration ||
			result.httpStatusCode != tc.expected.httpStatusCode ||
			result.bytesRead != tc.expected.bytesRead ||
			result.capturedRequestCookie != tc.expected.capturedRequestCookie ||
			result.capturedResponseCookie != tc.expected.capturedResponseCookie ||
			result.terminationState != tc.expected.terminationState ||
			result.actconn != tc.expected.actconn ||
			result.feconn != tc.expected.feconn ||
			result.beconn != tc.expected.beconn ||
			result.srvconn != tc.expected.srvconn ||
			result.retries != tc.expected.retries ||
			result.srvQueue != tc.expected.srvQueue ||
			result.backendQueue != tc.expected.backendQueue ||
			result.capturedRequestHeaders != tc.expected.capturedRequestHeaders ||
			result.capturedResponseHeaders != tc.expected.capturedResponseHeaders ||
			result.httpVerb != tc.expected.httpVerb ||
			result.httpProto != tc.expected.httpProto ||
			result.httpHost != tc.expected.httpHost ||
			result.httpRequest != tc.expected.httpRequest ||
			result.httpVersion != tc.expected.httpVersion {
			t.Errorf("Mismatch in parsed log line: %v", result)
		}
	}
}
