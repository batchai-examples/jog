package grok_vjeantet

import (
	"testing"
)

func TestAwsRegex(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{
			input:    "GET /path HTTP/1.1",
			expected: true,
		},
		{
			input:    "PUT /bucket/key HTTP/1.0",
			expected: true,
		},
		{
			input:    "POST /data HTTP/2.0",
			expected: true,
		},
		{
			input:    "DELETE /resource HTTP/1.1",
			expected: true,
		},
		{
			input:    "HEAD /file.txt HTTP/1.0",
			expected: true,
		},
		{
			input:    "GET /path?param=value HTTP/1.1",
			expected: true,
		},
		{
			input:    "GET /path#fragment HTTP/1.1",
			expected: false,
		},
		{
			input:    "GET /path HTTP/0.9",
			expected: false,
		},
		{
			input:    "GET /path HTTP/3.0",
			expected: false,
		},
		{
			input:    "GET /path HTTP/",
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Input: %s", tc.input), func(t *testing.T) {
			result := AwsRegex.MatchString(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %v, got %v for input '%s'", tc.expected, result, tc.input)
			}
		})
	}
}

func TestElbRegex(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{
			input:    "2023-04-15T12:34:56Z elb1 192.168.1.1:8080 (10.0.0.1:80) 0.001 0.002 0.003 200 200 1024 512",
			expected: true,
		},
		{
			input:    "2023-04-15T12:34:56Z elb2 192.168.1.2:8080 - 0.001 0.002 0.003 200 200 1024 512",
			expected: true,
		},
		{
			input:    "2023-04-15T12:34:56Z elb3 192.168.1.3:8080 (10.0.0.2:80) - 0.001 0.002 0.003 200 200 1024 512",
			expected: true,
		},
		{
			input:    "2023-04-15T12:34:56Z elb4 192.168.1.4:8080 (10.0.0.3:80) - 0.001 0.002 0.003 200 200 1024 512",
			expected: true,
		},
		{
			input:    "2023-04-15T12:34:56Z elb5 192.168.1.5:8080 (10.0.0.4:80) - 0.001 0.002 0.003 200 200 1024 512",
			expected: true,
		},
		{
			input:    "2023-04-15T12:34:56Z elb6 192.168.1.6:8080 (10.0.0.5:80) - 0.001 0.002 0.003 200 200 1024 512",
			expected: true,
		},
		{
			input:    "2023-04-15T12:34:56Z elb7 192.168.1.7:8080 (10.0.0.6:80) - 0.001 0.002 0.003 200 200 1024 512",
			expected: true,
		},
		{
			input:    "2023-04-15T12:34:56Z elb8 192.168.1.8:8080 (10.0.0.7:80) - 0.001 0.002 0.003 200 200 1024 512",
			expected: true,
		},
		{
			input:    "2023-04-15T12:34:56Z elb9 192.168.1.9:8080 (10.0.0.8:80) - 0.001 0.002 0.003 200 200 1024 512",
			expected: true,
		},
		{
			input:    "2023-04-15T12:34:56Z elb10 192.168.1.10:8080 (10.0.0.9:80) - 0.001 0.002 0.003 200 200 1024 512",
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Input: %s", tc.input), func(t *testing.T) {
			result := ElbRegex.MatchString(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %v, got %v for input '%s'", tc.expected, result, tc.input)
			}
		})
	}
}

func TestS3Regex(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{
			input:    "my-bucket/my-object",
			expected: true,
		},
		{
			input:    "my-bucket/my-folder/my-subfolder/my-object",
			expected: true,
		},
		{
			input:    "my-bucket/my-object.txt",
			expected: true,
		},
		{
			input:    "my-bucket/my-object/",
			expected: false,
		},
		{
			input:    "my-bucket//my-object",
			expected: false,
		},
		{
			input:    "my-bucket/my-object..txt",
			expected: false,
		},
		{
			input:    "my-bucket/my-object.txt.",
			expected: false,
		},
		{
			input:    "my-bucket/my-object.txt..",
			expected: false,
		},
		{
			input:    "my-bucket/my-object.txt...",
			expected: false,
		},
		{
			input:    "my-bucket/my-object.txt....",
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Input: %s", tc.input), func(t *testing.T) {
			result := S3Regex.MatchString(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %v, got %v for input '%s'", tc.expected, result, tc.input)
			}
		})
	}
}

func TestDynamoDBRegex(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{
			input:    "my-table/my-item",
			expected: true,
		},
		{
			input:    "my-table/my-folder/my-subfolder/my-item",
			expected: true,
		},
		{
			input:    "my-table/my-item.txt",
			expected: true,
		},
		{
			input:    "my-table/my-item/",
			expected: false,
		},
		{
			input:    "my-table//my-item",
			expected: false,
		},
		{
			input:    "my-table/my-item..txt",
			expected: false,
		},
		{
			input:    "my-table/my-item.txt.",
			expected: false,
		},
		{
			input:    "my-table/my-item.txt..",
			expected: false,
		},
		{
			input:    "my-table/my-item.txt...",
			expected: false,
		},
		{
			input:    "my-table/my-item.txt....",
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Input: %s", tc.input), func(t *testing.T) {
			result := DynamoDBRegex.MatchString(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %v, got %v for input '%s'", tc.expected, result, tc.input)
			}
		})
	}
}

func TestRDSRegex(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{
			input:    "my-db/my-instance",
			expected: true,
		},
		{
			input:    "my-db/my-folder/my-subfolder/my-instance",
			expected: true,
		},
		{
			input:    "my-db/my-instance.txt",
			expected: true,
		},
		{
			input:    "my-db/my-instance/",
			expected: false,
		},
		{
			input:    "my-db//my-instance",
			expected: false,
		},
		{
			input:    "my-db/my-instance..txt",
			expected: false,
		},
		{
			input:    "my-db/my-instance.txt.",
			expected: false,
		},
		{
			input:    "my-db/my-instance.txt..",
			expected: false,
		},
		{
			input:    "my-db/my-instance.txt...",
			expected: false,
		},
		{
			input:    "my-db/my-instance.txt....",
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Input: %s", tc.input), func(t *testing.T) {
			result := RDSRegex.MatchString(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %v, got %v for input '%s'", tc.expected, result, tc.input)
			}
		})
	}
}

func TestElasticBeanstalkRegex(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{
			input:    "my-app/my-env",
			expected: true,
		},
		{
			input:    "my-app/my-folder/my-subfolder/my-env",
			expected: true,
		},
		{
			input:    "my-app/my-env.txt",
			expected: true,
		},
		{
			input:    "my-app/my-env/",
			expected: false,
		},
		{
			input:    "my-app//my-env",
			expected: false,
		},
		{
			input:    "my-app/my-env..txt",
			expected: false,
		},
		{
			input:    "my-app/my-env.txt.",
			expected: false,
		},
		{
			input:    "my-app/my-env.txt..",
			expected: false,
		},
		{
			input:    "my-app/my-env.txt...",
			expected: false,
		},
		{
			input:    "my-app/my-env.txt....",
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Input: %s", tc.input), func(t *testing.T) {
			result := ElasticBeanstalkRegex.MatchString(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %v, got %v for input '%s'", tc.expected, result, tc.input)
			}
		})
	}
}

func TestElasticCacheRegex(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{
			input:    "my-cache/my-instance",
			expected: true,
		},
		{
			input:    "my-cache/my-folder/my-subfolder/my-instance",
			expected: true,
		},
		{
			input:    "my-cache/my-instance.txt",
			expected: true,
		},
		{
			input:    "my-cache/my-instance/",
			expected: false,
		},
		{
			input:    "my-cache//my-instance",
			expected: false,
		},
		{
			input:    "my-cache/my-instance..txt",
			expected: false,
		},
		{
			input:    "my-cache/my-instance.txt.",
			expected: false,
		},
		{
			input:    "my-cache/my-instance.txt..",
			expected: false,
		},
		{
			input:    "my-cache/my-instance.txt...",
			expected: false,
		},
		{
			input:    "my-cache/my-instance.txt....",
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Input: %s", tc.input), func(t *testing.T) {
			result := ElasticCacheRegex.MatchString(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %v, got %v for input '%s'", tc.expected, result, tc.input)
			}
		})
	}
}

func TestElasticSearchRegex(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{
			input:    "my-es/my-instance",
			expected: true,
		},
		{
			input:    "my-es/my-folder/my-subfolder/my-instance",
			expected: true,
		},
		{
			input:    "my-es/my-instance.txt",
			expected: true,
		},
		{
			input:    "my-es/my-instance/",
			expected: false,
		},
		{
			input:    "my-es//my-instance",
			expected: false,
		},
		{
			input:    "my-es/my-instance..txt",
			expected: false,
		},
		{
			input:    "my-es/my-instance.txt.",
			expected: false,
		},
		{
			input:    "my-es/my-instance.txt..",
			expected: false,
		},
		{
			input:    "my-es/my-instance.txt...",
			expected: false,
		},
		{
			input:    "my-es/my-instance.txt....",
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Input: %s", tc.input), func(t *testing.T) {
			result := ElasticSearchRegex.MatchString(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %v, got %v for input '%s'", tc.expected, result, tc.input)
			}
		})
	}
}

func TestRDSInstanceRegex(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{
			input:    "my-rds/my-instance",
			expected: true,
		},
		{
			input:    "my-rds/my-folder/my-subfolder/my-instance",
			expected: true,
		},
		{
			input:    "my-rds/my-instance.txt",
			expected: true,
		},
		{
			input:    "my-rds/my-instance/",
			expected: false,
		},
		{
			input:    "my-rds//my-instance",
			expected: false,
		},
		{
			input:    "my-rds/my-instance..txt",
			expected: false,
		},
		{
			input:    "my-rds/my-instance.txt.",
			expected: false,
		},
		{
			input:    "my-rds/my-instance.txt..",
			expected: false,
		},
		{
			input:    "my-rds/my-instance.txt...",
			expected: false,
		},
		{
			input:    "my-rds/my-instance.txt....",
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Input: %s", tc.input), func(t *testing.T) {
			result := RDSInstanceRegex.MatchString(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %v, got %v for input '%s'", tc.expected, result, tc.input)
			}
		})
	}
}

func TestRDSDBInstanceRegex(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{
			input:    "my-rds/my-db-instance",
			expected: true,
		},
		{
			input:    "my-rds/my-folder/my-subfolder/my-db-instance",
			expected: true,
		},
		{
			input:    "my-rds/my-db-instance.txt",
			expected: true,
		},
		{
			input:    "my-rds/my-db-instance/",
			expected: false,
		},
		{
			input:    "my-rds//my-db-instance",
			expected: false,
		},
		{
			input:    "my-rds/my-db-instance..txt",
			expected: false,
		},
		{
			input:    "my-rds/my-db-instance.txt.",
			expected: false,
		},
		{
			input:    "my-rds/my-db-instance.txt..",
			expected: false,
		},
		{
			input:    "my-rds/my-db-instance.txt...",
			expected: false,
		},
		{
			input:    "my-rds/my-db-instance.txt....",
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Input: %s", tc.input), func(t *testing.T) {
			result := RDSDBInstanceRegex.MatchString(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %v, got %v for input '%s'", tc.expected, result, tc.input)
			}
		})
	}
}

func TestRDSDBParameterGroupRegex(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{
			input:    "my-rds/my-db-parameter-group",
			expected: true,
		},
		{
			input:    "my-rds/my-folder/my-subfolder/my-db-parameter-group",
			expected: true,
		},
		{
			input:    "my-rds/my-db-parameter-group.txt",
			expected: true,
		},
		{
			input:    "my-rds/my-db-parameter-group/",
			expected: false,
		},
		{
			input:    "my-rds//my-db-parameter-group",
			expected: false,
		},
		{
			input:    "my-rds/my-db-parameter-group..txt",
			expected: false,
		},
		{
			input:    "my-rds/my-db-parameter-group.txt.",
			expected: false,
		},
		{
			input:    "my-rds/my-db-parameter-group.txt..",
			expected: false,
		},
		{
			input:    "my-rds/my-db-parameter-group.txt...",
			expected: false,
		},
		{
			input:    "my-rds/my-db-parameter-group.txt....",
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Input: %s", tc.input), func(t *testing.T) {
			result := RDSDBParameterGroupRegex.MatchString(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %v, got %v for input '%s'", tc.expected, result, tc.input)
			}
		})
	}
}

func TestRDSDBSubnetGroupRegex(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{
			input:    "my-rds/my-db-subnet-group",
			expected: true,
		},
		{
			input:    "my-rds/my-folder/my-subfolder/my-db-subnet-group",
			expected: true,
		},
		{
			input:    "my-rds/my-db-subnet-group.txt",
			expected: true,
		},
		{
			input:    "my-rds/my-db-subnet-group/",
			expected: false,
		},
		{
			input:    "my-rds//my-db-subnet-group",
			expected: false,
		},
		{
			input:    "my-rds/my-db-subnet-group..txt",
			expected: false,
		},
		{
			input:    "my-rds/my-db-subnet-group.txt.",
			expected: false,
		},
		{
			input:    "my-rds/my-db-subnet-group.txt..",
			expected: false,
		},
		{
			input:    "my-rds/my-db-subnet-group.txt...",
			expected: false,
		},
		{
			input:    "my-rds/my-db-subnet-group.txt....",
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Input: %s", tc.input), func(t *testing.T) {
			result := RDSDBSubnetGroupRegex.MatchString(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %v, got %v for input '%s'", tc.expected, result, tc.input)
			}
		})
	}
}

func TestRDSDBSnapshotRegex(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{
			input:    "my-rds/my-db-snapshot",
			expected: true,
		},
		{
			input:    "my-rds/my-folder/my-subfolder/my-db-snapshot",
			expected: true,
		},
		{
			input:    "my-rds/my-db-snapshot.txt",
			expected: true,
		},
		{
			input:    "my-rds/my-db-snapshot/",
			expected: false,
		},
		{
			input:    "my-rds//my-db-snapshot",
			expected: false,
		},
		{
			input:    "my-rds/my-db-snapshot..txt",
			expected: false,
		},
		{
			input:    "my-rds/my-db-snapshot.txt.",
			expected: false,
		},
		{
			input:    "my-rds/my-db-snapshot.txt..",
			expected: false,
		},
		{
			input:    "my-rds/my-db-snapshot.txt...",
			expected: false,
		},
		{
			input:    "my-rds/my-db-snapshot.txt....",
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Input: %s", tc.input), func(t *testing.T) {
			result := RDSDBSnapshotRegex.MatchString(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %v, got %v for input '%s'", tc.expected, result, tc.input)
			}
		})
	}
}

func TestRDSDBClusterRegex(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{
			input:    "my-rds/my-db-cluster",
			expected: true,
		},
		{
			input:    "my-rds/my-folder/my-subfolder/my-db-cluster",
			expected: true,
		},
		{
			input:    "my-rds/my-db-cluster.txt",
			expected: true,
		},
		{
			input:    "my-rds/my-db-cluster/",
			expected: false,
		},
		{
			input:    "my-rds//my-db-cluster",
			expected: false,
		},
		{
			input:    "my-rds/my-db-cluster..txt",
			expected: false,
		},
		{
			input:    "my-rds/my-db-cluster.txt.",
			expected: false,
		},
		{
			input:    "my-rds/my-db-cluster.txt..",
			expected: false,
		},
		{
			input:    "my-rds/my-db-cluster.txt...",
			expected: false,
		},
		{
			input:    "my-rds/my-db-cluster.txt....",
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Input: %s", tc.input), func(t *testing.T) {
			result := RDSDBClusterRegex.MatchString(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %v, got %v for input '%s'", tc.expected, result, tc.input)
			}
		})
	}
}

func TestRDSDBClusterSnapshotRegex(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{
			input:    "my-rds/my-db-cluster-snapshot",
			expected: true,
		},
		{
			input:    "my-rds/my-folder/my-subfolder/my-db-cluster-snapshot",
			expected: true,
		},
		{
			input:    "my-rds/my-db-cluster-snapshot.txt",
			expected: true,
		},
		{
			input:    "my-rds/my-db-cluster-snapshot/",
			expected: false,
		},
		{
			input:    "my-rds//my-db-cluster-snapshot",
			expected: false,
		},
		{
			input:    "my-rds/my-db-cluster-snapshot..txt",
			expected: false,
		},
		{
			input:    "my-rds/my-db-cluster-snapshot.txt.",
			expected: false,
		},
		{
			input:    "my-rds/my-db-cluster-snapshot.txt..",
			expected: false,
		},
		{
			input:    "my-rds/my-db-cluster-snapshot.txt...",
			expected: false,
		},
		{
			input:    "my-rds/my-db-cluster-snapshot.txt....",
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Input: %s", tc.input), func(t *testing.T) {
			result := RDSDBClusterSnapshotRegex.MatchString(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %v, got %v for input '%s'", tc.expected, result, tc.input)
			}
		})
	}
}

func TestRDSDBProxyRegex(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{
			input:    "my-rds/my-db-proxy",
			expected: true,
		},
		{
			input:    "my-rds/my-folder/my-subfolder/my-db-proxy",
			expected: true,
		},
		{
			input:    "my-rds/my-db-proxy.txt",
			expected: true,
		},
		{
			input:    "my-rds/my-db-proxy/",
			expected: false,
		},
		{
			input:    "my-rds//my-db-proxy",
			expected: false,
		},
		{
			input:    "my-rds/my-db-proxy..txt",
			expected: false,
		},
		{
			input:    "my-rds/my-db-proxy.txt.",
			expected: false,
		},
		{
			input:    "my-rds/my-db-proxy.txt..",
			expected: false,
		},
		{
			input:    "my-rds/my-db-proxy.txt...",
			expected: false,
		},
		{
			input:    "my-rds/my-db-proxy.txt....",
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Input: %s", tc.input), func(t *testing.T) {
			result := RDSDBProxyRegex.MatchString(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %v, got %v for input '%s'", tc.expected, result, tc.input)
			}
		})
	}
}

func TestRDSDBProxyEndpointRegex(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{
			input:    "my-rds/my-db-proxy-endpoint",
			expected: true,
		},
		{
			input:    "my-rds/my-folder/my-subfolder
