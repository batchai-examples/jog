package grok_vjeantet

import (
	"testing"
)

func TestJunosPatternRT_FLOW_EVENT(t *testing.T) {
	testCases := []struct {
		input    string
		expected map[string]string
	}{
		{
			input: `RT_FLOW_SESSION_CREATE: session created 192.168.1.1/5000->192.168.1.2/6000 SSH 10.0.0.1/5001->10.0.0.2/6001 NAT-RULE-OUTSIDE NAT-RULE-INSIDE 6 HTTP outside inside 12345 sent(100) received(200) 10 .`,
			expected: map[string]string{
				"event":                "RT_FLOW_SESSION_CREATE",
				"close-reason":         "",
				"src-ip":               "192.168.1.1",
				"src-port":             "5000",
				"dst-ip":               "192.168.1.2",
				"dst-port":             "6000",
				"service":              "SSH",
				"nat-src-ip":           "10.0.0.1",
				"nat-src-port":         "5001",
				"nat-dst-ip":           "10.0.0.2",
				"nat-dst-port":         "6001",
				"src-nat-rule-name":    "NAT-RULE-OUTSIDE",
				"dst-nat-rule-name":    "NAT-RULE-INSIDE",
				"protocol-id":          "6",
				"policy-name":          "HTTP",
				"from-zone":            "outside",
				"to-zone":              "inside",
				"session-id":           "12345",
				"sent":                 "100",
				"received":             "200",
				"elapsed-time":         "10",
			},
		},
		{
			input: `RT_FLOW_SESSION_CLOSE: session closed 192.168.1.3/5002->192.168.1.4/6002 SSH 10.0.0.3/5003->10.0.0.4/6003 NAT-RULE-OUTSIDE NAT-RULE-INSIDE 6 HTTP outside inside 12346 sent(150) received(250) 15 .`,
			expected: map[string]string{
				"event":                "RT_FLOW_SESSION_CLOSE",
				"close-reason":         "",
				"src-ip":               "192.168.1.3",
				"src-port":             "5002",
				"dst-ip":               "192.168.1.4",
				"dst-port":             "6002",
				"service":              "SSH",
				"nat-src-ip":           "10.0.0.3",
				"nat-src-port":         "5003",
				"nat-dst-ip":           "10.0.0.4",
				"nat-dst-port":         "6003",
				"src-nat-rule-name":    "NAT-RULE-OUTSIDE",
				"dst-nat-rule-name":    "NAT-RULE-INSIDE",
				"protocol-id":          "6",
				"policy-name":          "HTTP",
				"from-zone":            "outside",
				"to-zone":              "inside",
				"session-id":           "12346",
				"sent":                 "150",
				"received":             "250",
				"elapsed-time":         "15",
			},
		},
		{
			input: `RT_FLOW_SESSION_DENY: session denied 192.168.1.5/5004->192.168.1.6/6004 SSH 10.0.0.5/5005->10.0.0.6/6005 NAT-RULE-OUTSIDE NAT-RULE-INSIDE 6 HTTP outside inside .`,
			expected: map[string]string{
				"event":                "RT_FLOW_SESSION_DENY",
				"close-reason":         "",
				"src-ip":               "192.168.1.5",
				"src-port":             "5004",
				"dst-ip":               "192.168.1.6",
				"dst-port":             "6004",
				"service":              "SSH",
				"nat-src-ip":           "10.0.0.5",
				"nat-src-port":         "5005",
				"nat-dst-ip":           "10.0.0.6",
				"nat-dst-port":         "6005",
				"src-nat-rule-name":    "NAT-RULE-OUTSIDE",
				"dst-nat-rule-name":    "NAT-RULE-INSIDE",
				"protocol-id":          "6",
				"policy-name":          "HTTP",
				"from-zone":            "outside",
				"to-zone":              "inside",
			},
		},
		{
			input: `RT_FLOW1 RT_FLOW_SESSION_CREATE: session created 192.168.1.7/5006->192.168.1.8/6006 SSH 10.0.0.7/5007->10.0.0.8/6007 NAT-RULE-OUTSIDE NAT-RULE-INSIDE 6 HTTP outside inside 12347 sent(200) received(300) 20 .`,
			expected: map[string]string{
				"event":                "RT_FLOW_SESSION_CREATE",
				"close-reason":         "",
				"src-ip":               "192.168.1.7",
				"src-port":             "5006",
				"dst-ip":               "192.168.1.8",
				"dst-port":             "6006",
				"service":              "SSH",
				"nat-src-ip":           "10.0.0.7",
				"nat-src-port":         "5007",
				"nat-dst-ip":           "10.0.0.8",
				"nat-dst-port":         "6007",
				"src-nat-rule-name":    "NAT-RULE-OUTSIDE",
				"dst-nat-rule-name":    "NAT-RULE-INSIDE",
				"protocol-id":          "6",
				"policy-name":          "HTTP",
				"from-zone":            "outside",
				"to-zone":              "inside",
				"session-id":           "12347",
				"sent":                 "200",
				"received":             "300",
				"elapsed-time":         "20",
			},
		},
		{
			input: `RT_FLOW1 RT_FLOW_SESSION_CLOSE: session closed 192.168.1.9/5008->192.168.1.10/6008 SSH 10.0.0.9/5009->10.0.0.10/6009 NAT-RULE-OUTSIDE NAT-RULE-INSIDE 6 HTTP outside inside 12348 sent(250) received(350) 25 .`,
			expected: map[string]string{
				"event":                "RT_FLOW_SESSION_CLOSE",
				"close-reason":         "",
				"src-ip":               "192.168.1.9",
				"src-port":             "5008",
				"dst-ip":               "192.168.1.10",
				"dst-port":             "6008",
				"service":              "SSH",
				"nat-src-ip":           "10.0.0.9",
				"nat-src-port":         "5009",
				"nat-dst-ip":           "10.0.0.10",
				"nat-dst-port":         "6009",
				"src-nat-rule-name":    "NAT-RULE-OUTSIDE",
				"dst-nat-rule-name":    "NAT-RULE-INSIDE",
				"protocol-id":          "6",
				"policy-name":          "HTTP",
				"from-zone":            "outside",
				"to-zone":              "inside",
				"session-id":           "12348",
				"sent":                 "250",
				"received":             "350",
				"elapsed-time":         "25",
			},
		},
		{
			input: `RT_FLOW1 RT_FLOW_SESSION_DENY: session denied 192.168.1.11/5010->192.168.1.12/6010 SSH 10.0.0.11/5011->10.0.0.12/6011 NAT-RULE-OUTSIDE NAT-RULE-INSIDE 6 HTTP outside inside .`,
			expected: map[string]string{
				"event":                "RT_FLOW_SESSION_DENY",
				"close-reason":         "",
				"src-ip":               "192.168.1.11",
				"src-port":             "5010",
				"dst-ip":               "192.168.1.12",
				"dst-port":             "6010",
				"service":              "SSH",
				"nat-src-ip":           "10.0.0.11",
				"nat-src-port":         "5011",
				"nat-dst-ip":           "10.0.0.12",
				"nat-dst-port":         "6011",
				"src-nat-rule-name":    "NAT-RULE-OUTSIDE",
				"dst-nat-rule-name":    "NAT-RULE-INSIDE",
				"protocol-id":          "6",
				"policy-name":          "HTTP",
				"from-zone":            "outside",
				"to-zone":              "inside",
				"session-id":           "",
				"sent":                 "",
				"received":             "",
				"elapsed-time":         "",
			},
		},
	}

	for _, tc := range testCases {
		result, err := parseFlow(tc.input)
		if err != nil {
			t.Errorf("parseFlow(%q) error: %v", tc.input, err)
		} else if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("parseFlow(%q) = %v; want %v", tc.input, result, tc.expected)
		}
	}
}
