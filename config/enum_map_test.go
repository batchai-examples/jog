package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnumMapT_Reset(t *testing.T) {
	i := &EnumMapT{
		CaseSensitive: true,
		Default:       "default",
		values:        map[string]Enum{"key1": {Name: "key1", Alias: Alias{Values: map[string]bool{"alias1": true}}}},
		allMap:        map[string]Enum{"key1": {Name: "key1", Alias: Alias{Values: map[string]bool{"alias1": true}}}},
	}

	i.Reset()

	assert.False(t, i.CaseSensitive)
	assert.Empty(t, i.Default)
	assert.Empty(t, i.values)
	assert.Empty(t, i.allMap)
}

func TestEnumMapT_UnmarshalYAML(t *testing.T) {
	var i EnumMap
	err := i.UnmarshalYAML(func(interface{}) error { return nil })
	assert.NoError(t, err)

	i = &EnumMapT{}
	err = i.UnmarshalYAML(func(interface{}) error {
		return errors.New("unmarshal error")
	})
	assert.Error(t, err)
}

func TestEnumMapT_MarshalYAML(t *testing.T) {
	i := &EnumMapT{
		CaseSensitive: true,
		Default:       "default",
		values:        map[string]Enum{"key1": {Name: "key1", Alias: Alias{Values: map[string]bool{"alias1": true}}}},
		allMap:        map[string]Enum{"key1": {Name: "key1", Alias: Alias{Values: map[string]bool{"alias1": true}}}},
	}

	data, err := i.MarshalYAML()
	assert.NoError(t, err)
	assert.NotNil(t, data)

	i = &EnumMapT{}
	data, err = i.MarshalYAML()
	assert.NoError(t, err)
	assert.Nil(t, data)
}

func TestEnumMapT_Init(t *testing.T) {
	var i EnumMap
	i.Init(Configuration{})
}

func TestEnumMapT_IsEmpty(t *testing.T) {
	i := &EnumMapT{
		values: map[string]Enum{},
	}
	assert.True(t, i.IsEmpty())

	i = &EnumMapT{
		values: map[string]Enum{"key1": {Name: "key1", Alias: Alias{Values: map[string]bool{"alias1": true}}}},
	}
	assert.False(t, i.IsEmpty())
}

func TestEnumMapT_GetEnum(t *testing.T) {
	i := &EnumMapT{
		CaseSensitive: false,
		Default:       "default",
		values:        map[string]Enum{"key1": {Name: "key1", Alias: Alias{Values: map[string]bool{"alias1": true}}}},
		allMap:        map[string]Enum{"key1": {Name: "key1", Alias: Alias{Values: map[string]bool{"alias1": true}}}},
	}

	enum := i.GetEnum("Key1")
	assert.Equal(t, "key1", enum.Name)

	i.Default = "default2"
	enum = i.GetEnum("Key1")
	assert.Equal(t, "default2", enum.Name)
}

func TestEnumMapT_FromMap(t *testing.T) {
	var i EnumMap
	err := i.FromMap(map[string]interface{}{
		"case-sensitive": true,
		"default":        "default",
		"key1":           map[string]interface{}{"name": "key1", "alias": map[string]bool{"alias1": true}},
	})
	assert.NoError(t, err)
	assert.True(t, i.CaseSensitive)
	assert.Equal(t, "default", i.Default)
	assert.Len(t, i.values, 1)
	assert.Len(t, i.allMap, 2)

	err = i.FromMap(map[string]interface{}{
		"case-sensitive": true,
	})
	assert.Error(t, err)

	err = i.FromMap(map[string]interface{}{
		"default": "default",
	})
	assert.NoError(t, err)
}

func TestEnumMapT_ToMap(t *testing.T) {
	i := &EnumMapT{
		CaseSensitive: true,
		Default:       "default",
		values:        map[string]Enum{"key1": {Name: "key1", Alias: Alias{Values: map[string]bool{"alias1": true}}}},
		allMap:        map[string]Enum{"key1": {Name: "key1", Alias: Alias{Values: map[string]bool{"alias1": true}}}},
	}

	data := i.ToMap()
	assert.Equal(t, true, data["case-sensitive"])
	assert.Equal(t, "default", data["default"])
	assert.Len(t, data, 2)
}
!!!!test_end!!!!
