package util

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnyValueFromRaw(t *testing.T) {
	testCases := []struct {
		lineNo     int
		raw        interface{}
		replace    map[string]string
		expected   string
	}{
		{
			lineNo: 1,
			raw:    nil,
			replace: map[string]string{},
			expected: "",
		},
		{
			lineNo: 2,
			raw:    "test",
			replace: map[string]string{},
			expected: "test",
		},
		{
			lineNo: 3,
			raw:    map[string]interface{}{"key": "value"},
			replace: map[string]string{},
			expected: `{
  "key": "value"
}`,
		},
		{
			lineNo: 4,
			raw:    []interface{}{"item1", "item2"},
			replace: map[string]string{},
			expected: `[  
  "item1",  
  "item2"  
]`,
		},
		{
			lineNo: 5,
			raw:    `{"key": "value"}`,
			replace: map[string]string{},
			expected: `{
  "key": "value"
}`,
		},
		{
			lineNo: 6,
			raw:    `"test"`,
			replace: map[string]string{},
			expected: "test",
		},
		{
			lineNo: 7,
			raw:    `'test'`,
			replace: map[string]string{},
			expected: "test",
		},
		{
			lineNo: 8,
			raw:    `{"key": "value"}`,
			replace: map[string]string{"old": "new"},
			expected: `{
  "key": "newValue"
}`,
		},
		{
			lineNo: 9,
			raw:    `{"key": "value"}`,
			replace: map[string]string{},
			expected: `{
  "key": "value"
}`,
		},
		{
			lineNo: 10,
			raw:    `{"key": "value"}`,
			replace: map[string]string{},
			expected: `{
  "key": "value"
}`,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("lineNo=%d,raw=%v", tc.lineNo, tc.raw), func(t *testing.T) {
			result := AnyValueFromRaw(tc.lineNo, tc.raw, tc.replace)
			assert.Equal(t, tc.expected, result.Text)
		})
	}
}
