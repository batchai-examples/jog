package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestField_Reset(t *testing.T) {
	field := &FieldT{
		Name:           "test",
		Alias:          util.NewMultiString("alias1", "alias2"),
		CaseSensitive:  true,
		TimeFormat:     "2006-01-02",
		Timezone:       "UTC",
		Type:           FieldTypeTime,
	}

	field.Reset()

	assert.Equal(t, "", field.Name)
	assert.Equal(t, util.NewMultiString(""), field.Alias)
	assert.False(t, field.CaseSensitive)
	assert.Equal(t, &CompressPrefixT{}, field.CompressPrefix)
	assert.Equal(t, &EnumMapT{}, field.Enums)
	assert.Equal(t, FieldTypeAuto, field.Type)
	assert.Equal(t, "", field.TimeFormat)
	assert.Equal(t, "", field.Timezone)
	assert.Nil(t, field.TimeLocation)
}

func TestField_UnmarshalYAML_HappyPath(t *testing.T) {
	field := &FieldT{}
	err := field.UnmarshalYAML(func(interface{}) error { return nil })
	assert.NoError(t, err)

	field = &FieldT{}
	err = field.UnmarshalYAML(func(interface{}) error {
		return fmt.Errorf("error unmarshalling YAML")
	})
	assert.Error(t, err)
}

func TestField_MarshalYAML_HappyPath(t *testing.T) {
	field := &FieldT{
		Name:           "test",
		Alias:          util.NewMultiString("alias1", "alias2"),
		CaseSensitive:  true,
		TimeFormat:     "2006-01-02",
		Timezone:       "UTC",
		Type:           FieldTypeTime,
	}

	result, err := field.MarshalYAML()
	assert.NoError(t, err)
	assert.NotNil(t, result)

	field = &FieldT{}
	result, err = field.MarshalYAML()
	assert.NoError(t, err)
	assert.Nil(t, result)
}

func TestField_Init_HappyPath(t *testing.T) {
	field := &FieldT{}
	cfg := Configuration{}
	field.Init(cfg)

	// Add more assertions as needed
}

func TestField_IsEnum_HappyPath(t *testing.T) {
	field := &FieldT{
		Enums: &EnumMapT{},
	}
	assert.True(t, field.IsEnum())

	field = &FieldT{}
	assert.False(t, field.IsEnum())
}

func TestField_ToMap_HappyPath(t *testing.T) {
	field := &FieldT{
		Name:           "test",
		Alias:          util.NewMultiString("alias1", "alias2"),
		CaseSensitive:  true,
		TimeFormat:     "2006-01-02",
		Timezone:       "UTC",
		Type:           FieldTypeTime,
	}

	result := field.ToMap()
	assert.Equal(t, map[string]interface{}{
		"name":         "test",
		"alias":        []interface{}{"alias1", "alias2"},
		"case-sensitive": true,
		"type":         "time",
		"time-format":  "2006-01-02",
		"timezone":     "UTC",
	}, result)

	field = &FieldT{}
	result = field.ToMap()
	assert.Equal(t, map[string]interface{}{
		"name":         "",
		"alias":        []interface{}{},
		"case-sensitive": false,
		"type":         "auto",
	}, result)
}

func TestField_FromMap_HappyPath(t *testing.T) {
	field := &FieldT{}
	m := map[string]interface{}{
		"name":         "test",
		"alias":        []interface{}{"alias1", "alias2"},
		"case-sensitive": true,
		"type":         "time",
		"time-format":  "2006-01-02",
		"timezone":     "UTC",
	}
	err := field.FromMap(m)
	assert.NoError(t, err)

	field = &FieldT{}
	m = map[string]interface{}{
		"name":         "",
		"alias":        []interface{}{},
		"case-sensitive": false,
		"type":         "auto",
	}
	err = field.FromMap(m)
	assert.NoError(t, err)

	field = &FieldT{}
	m = map[string]interface{}{
		"name":         "test",
		"alias":        []interface{}{"alias1", "alias2"},
		"case-sensitive": true,
		"type":         "unknown",
	}
	err = field.FromMap(m)
	assert.Error(t, err)

	field = &FieldT{}
	m = map[string]interface{}{
		"name":         "test",
		"alias":        []interface{}{"alias1", "alias2"},
		"case-sensitive": true,
		"time-format":  "invalid",
	}
	err = field.FromMap(m)
	assert.Error(t, err)

	field = &FieldT{}
	m = map[string]interface{}{
		"name":         "test",
		"alias":        []interface{}{"alias1", "alias2"},
		"case-sensitive": true,
		"timezone":     "invalid",
	}
	err = field.FromMap(m)
	assert.Error(t, err)
}

func TestField_GetColor_HappyPath(t *testing.T) {
	field := &FieldT{
		Enums: &EnumMapT{},
	}
	color := field.GetColor("value")
	assert.Equal(t, util.Color{}, color)

	field = &FieldT{}
	color = field.GetColor("value")
	assert.Equal(t, util.Color{}, color)
}

func TestField_GetColor_WithEnum(t *testing.T) {
	field := &FieldT{
		Enums: &EnumMapT{
			Map: map[string]*Enum{
				"value": &Enum{
					Color: util.Color{R: 255, G: 0, B: 0},
				},
			},
		},
	}
	color := field.GetColor("value")
	assert.Equal(t, util.Color{R: 255, G: 0, B: 0}, color)
}

func TestField_GetColor_WithNoEnum(t *testing.T) {
	field := &FieldT{}
	color := field.GetColor("value")
	assert.Equal(t, util.Color{}, color)
}
