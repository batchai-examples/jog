package config

import (
	"testing"
)

func TestGrokT_Init(t *testing.T) {
	testCases := []struct {
		name     string
		grok     GrokT
		expected []string
	}{
		{
			name: "Happy path with default library dirs",
			grok: GrokT{},
			expected: []string{
				"grok_vjeantet",
				"grok_extended",
			},
		},
		{
			name: "Happy path with custom library dirs",
			grok: GrokT{
				LibraryDirs: []string{"custom/dir1", "custom/dir2"},
			},
			expected: []string{
				"custom/dir1",
				"custom/dir2",
			},
		},
		{
			name: "Negative path with non-existing library dirs",
			grok: GrokT{
				LibraryDirs: []string{"non/existing/dir1", "non/existing/dir2"},
			},
			expected: []string{},
		},
		{
			name: "Happy path with empty library dirs",
			grok: GrokT{
				LibraryDirs: []string{},
			},
			expected: []string{},
		},
		{
			name: "Negative path with nil library dirs",
			grok: GrokT{
				LibraryDirs: nil,
			},
			expected: []string{},
		},
		{
			name: "Happy path with single custom library dir",
			grok: GrokT{
				LibraryDirs: []string{"custom/dir"},
			},
			expected: []string{
				"custom/dir",
			},
		},
		{
			name: "Negative path with empty string in library dirs",
			grok: GrokT{
				LibraryDirs: []string{""},
			},
			expected: []string{},
		},
		{
			name: "Happy path with multiple custom library dirs",
			grok: GrokT{
				LibraryDirs: []string{"custom/dir1", "custom/dir2", "custom/dir3"},
			},
			expected: []string{
				"custom/dir1",
				"custom/dir2",
				"custom/dir3",
			},
		},
		{
			name: "Negative path with duplicate library dirs",
			grok: GrokT{
				LibraryDirs: []string{"custom/dir", "custom/dir"},
			},
			expected: []string{
				"custom/dir",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.grok.Init(Configuration{})
			actual := tc.grok.LibraryDirs
			if len(actual) != len(tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, actual)
			}
			for i := range tc.expected {
				if actual[i] != tc.expected[i] {
					t.Errorf("Expected %s at index %d, but got %s", tc.expected[i], i, actual[i])
				}
			}
		})
	}
}

func TestGrokT_Parse(t *testing.T) {
	testCases := []struct {
		name     string
		grok     GrokT
		pattern  string
		line     string
		expected map[string]string
	}{
		{
			name: "Happy path",
			grok: GrokT{},
			pattern: `%{IP:ip} %{WORD:word}`,
			line:    "192.168.1.1 example",
			expected: map[string]string{
				"ip":   "192.168.1.1",
				"word": "example",
			},
		},
		{
			name: "Negative path with invalid pattern",
			grok: GrokT{},
			pattern: `%{IP:ip} %{WORD:word`,
			line:    "192.168.1.1 example",
			expected: map[string]string{},
		},
		{
			name: "Negative path with invalid line",
			grok: GrokT{},
			pattern: `%{IP:ip} %{WORD:word}`,
			line:    "192.168.1.1",
			expected: map[string]string{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := tc.grok.Parse(tc.pattern, tc.line)
			if err != nil {
				t.Errorf("Expected no error, but got %v", err)
			}
			if len(result) != len(tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
			for key, value := range tc.expected {
				if result[key] != value {
					t.Errorf("Expected %s at key %s, but got %s", value, key, result[key])
				}
			}
		})
	}
}

func TestGrokT_UnmarshalYAML(t *testing.T) {
	testCases := []struct {
		name     string
		data     []byte
		expected GrokT
	}{
		{
			name: "Happy path",
			data: []byte(`library-dirs:
- custom/dir1
- custom/dir2
uses:
- pattern1
- pattern2
matches-fields:
- field1
- field2`),
			expected: GrokT{
				LibraryDirs: []string{"custom/dir1", "custom/dir2"},
				Uses:        []string{"pattern1", "pattern2"},
				MatchesFields: []string{"field1", "field2"},
			},
		},
		{
			name: "Negative path with invalid data",
			data: []byte(`library-dirs:
- custom/dir1
- custom/dir2
uses:
- pattern1
- pattern2
matches-fields:
- field1
- field2`),
			expected: GrokT{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var result GrokT
			err := yaml.Unmarshal(tc.data, &result)
			if err != nil {
				t.Errorf("Expected no error, but got %v", err)
			}
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}

func TestGrokT_MarshalYAML(t *testing.T) {
	testCases := []struct {
		name     string
		grok     GrokT
		expected []byte
	}{
		{
			name: "Happy path",
			grok: GrokT{
				LibraryDirs: []string{"custom/dir1", "custom/dir2"},
				Uses:        []string{"pattern1", "pattern2"},
				MatchesFields: []string{"field1", "field2"},
			},
			expected: []byte(`library-dirs:
- custom/dir1
- custom/dir2
uses:
- pattern1
- pattern2
matches-fields:
- field1
- field2`),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := yaml.Marshal(&tc.grok)
			if err != nil {
				t.Errorf("Expected no error, but got %v", err)
			}
			if !bytes.Equal(result, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}

func TestGrokT_Validate(t *testing.T) {
	testCases := []struct {
		name     string
		grok     GrokT
		expected bool
	}{
		{
			name: "Happy path",
			grok: GrokT{
				LibraryDirs: []string{"custom/dir1", "custom/dir2"},
				Uses:        []string{"pattern1", "pattern2"},
				MatchesFields: []string{"field1", "field2"},
			},
			expected: true,
		},
		{
			name: "Negative path with empty library dirs",
			grok: GrokT{
				LibraryDirs: []string{},
				Uses:        []string{"pattern1", "pattern2"},
				MatchesFields: []string{"field1", "field2"},
			},
			expected: false,
		},
		{
			name: "Negative path with nil library dirs",
			grok: GrokT{
				LibraryDirs: nil,
				Uses:        []string{"pattern1", "pattern2"},
				MatchesFields: []string{"field1", "field2"},
			},
			expected: false,
		},
		{
			name: "Negative path with empty uses",
			grok: GrokT{
				LibraryDirs: []string{"custom/dir1", "custom/dir2"},
				Uses:        []string{},
				MatchesFields: []string{"field1", "field2"},
			},
			expected: false,
		},
		{
			name: "Negative path with nil uses",
			grok: GrokT{
				LibraryDirs: []string{"custom/dir1", "custom/dir2"},
				Uses:        nil,
				MatchesFields: []string{"field1", "field2"},
			},
			expected: false,
		},
		{
			name: "Negative path with empty matches fields",
			grok: GrokT{
				LibraryDirs: []string{"custom/dir1", "custom/dir2"},
				Uses:        []string{"pattern1", "pattern2"},
				MatchesFields: []string{},
			},
			expected: false,
		},
		{
			name: "Negative path with nil matches fields",
			grok: GrokT{
				LibraryDirs: []string{"custom/dir1", "custom/dir2"},
				Uses:        []string{"pattern1", "pattern2"},
				MatchesFields: nil,
			},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.grok.Validate()
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
