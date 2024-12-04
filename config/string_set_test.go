package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringSet_Parse_HappyPath(t *testing.T) {
	set := &StringSetT{CaseSensitive: true}
	set.Parse([]string{"apple", "banana", "cherry"})
	assert.Equal(t, map[string]bool{"apple": true, "banana": true, "cherry": true}, set.ValueMap)
}

func TestStringSet_Parse_EmptyInput(t *testing.T) {
	set := &StringSetT{CaseSensitive: true}
	set.Parse("")
	assert.Empty(t, set.ValueMap)
}

func TestStringSet_Parse_InvalidInput(t *testing.T) {
	set := &StringSetT{CaseSensitive: true}
	set.Parse(123)
	assert.Empty(t, set.ValueMap)
}

func TestStringSet_Reset_HappyPath(t *testing.T) {
	set := &StringSetT{
		ValueMap: map[string]bool{"apple": true, "banana": true},
	}
	set.Reset()
	assert.Empty(t, set.ValueMap)
}

func TestStringSet_IsEmpty_EmptySet(t *testing.T) {
	set := &StringSetT{}
	assert.True(t, set.IsEmpty())
}

func TestStringSet_IsEmpty_NonEmptySet(t *testing.T) {
	set := &StringSetT{ValueMap: map[string]bool{"apple": true}}
	assert.False(t, set.IsEmpty())
}

func TestStringSet_Contains_HappyPath(t *testing.T) {
	set := &StringSetT{
		ValueMap: map[string]bool{"apple": true, "banana": true},
	}
	assert.True(t, set.Contains("apple"))
	assert.False(t, set.Contains("Banana"))
}

func TestStringSet_Contains_CaseSensitive(t *testing.T) {
	set := &StringSetT{
		ValueMap: map[string]bool{"apple": true, "banana": true},
		CaseSensitive: true,
	}
	assert.True(t, set.Contains("apple"))
	assert.False(t, set.Contains("Banana"))
}

func TestStringSet_ContainsPrefixOf_HappyPath(t *testing.T) {
	set := &StringSetT{
		ValueMap: map[string]bool{"apple": true, "banana": true},
	}
	assert.True(t, set.ContainsPrefixOf("app"))
	assert.False(t, set.ContainsPrefixOf("Banana"))
}

func TestStringSet_ContainsPrefixOf_CaseSensitive(t *testing.T) {
	set := &StringSetT{
		ValueMap: map[string]bool{"apple": true, "banana": true},
		CaseSensitive: true,
	}
	assert.True(t, set.ContainsPrefixOf("app"))
	assert.False(t, set.ContainsPrefixOf("Banana"))
}

func TestStringSet_String_HappyPath(t *testing.T) {
	set := &StringSetT{
		ValueMap: map[string]bool{"apple": true, "banana": true},
	}
	result := set.String()
	assert.Equal(t, "apple, banana", result)
}

func TestStringSet_UnmarshalYAML_HappyPath(t *testing.T) {
	set := &StringSetT{CaseSensitive: true}
	err := set.UnmarshalYAML(func(interface{}) error {
		return nil
	})
	assert.NoError(t, err)
	set.Parse([]string{"apple", "banana", "cherry"})
	assert.Equal(t, map[string]bool{"apple": true, "banana": true, "cherry": true}, set.ValueMap)
}

func TestStringSet_MarshalYAML_HappyPath(t *testing.T) {
	set := &StringSetT{
		ValueMap: map[string]bool{"apple": true, "banana": true},
	}
	result, err := set.MarshalYAML()
	assert.NoError(t, err)
	assert.Equal(t, []interface{}{"apple", "banana"}, result)
}
