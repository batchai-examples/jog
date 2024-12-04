package main

import (
	"testing"
	"time"

	"github.com/araddon/dateparse"
	"github.com/qiangyt/jog/config"
	"github.com/qiangyt/jog/util"
)

func TestNewFieldValue_HappyPath(t *testing.T) {
	cfg := config.Configuration{
		Field: config.Field{
			Type: config.FieldTypeString,
			Color: util.Color{
				R: 255,
				G: 0,
				B: 0,
			},
		},
	}
	options := Options{}
	fieldConfig := cfg.Field
	value := util.AnyValue{Text: "test"}

	expected := &FieldValueT{
		value:     value,
		enumValue: nil,
		timeValue: time.Time{},
		Output:    "test",
		Config:    fieldConfig,
	}

	result := NewFieldValue(cfg, options, fieldConfig, value)

	if result.value != expected.value {
		t.Errorf("Expected value to be %v, but got %v", expected.value, result.value)
	}
	if result.enumValue != expected.enumValue {
		t.Errorf("Expected enumValue to be %v, but got %v", expected.enumValue, result.enumValue)
	}
	if !result.timeValue.Equal(expected.timeValue) {
		t.Errorf("Expected timeValue to be %v, but got %v", expected.timeValue, result.timeValue)
	}
	if result.Output != expected.Output {
		t.Errorf("Expected Output to be %s, but got %s", expected.Output, result.Output)
	}
	if !result.Config.Equal(expected.Config) {
		t.Errorf("Expected Config to be %v, but got %v", expected.Config, result.Config)
	}
}

func TestNewFieldValue_Enum(t *testing.T) {
	cfg := config.Configuration{
		Field: config.Field{
			Type: config.FieldTypeEnum,
			Color: util.Color{
				R: 255,
				G: 0,
				B: 0,
			},
			Enums: config.Enums{
				Map: map[string]config.Enum{
					"test": {
						Name: "TEST",
						Color: util.Color{
							R: 0,
							G: 255,
							B: 0,
						},
					},
				},
			},
		},
	}
	options := Options{}
	fieldConfig := cfg.Field
	value := util.AnyValue{Text: "test"}

	expected := &FieldValueT{
		value:     value,
		enumValue: config.Enum{Name: "TEST", Color: util.Color{R: 0, G: 255, B: 0}},
		timeValue: time.Time{},
		Output:    "TEST",
		Config:    fieldConfig,
	}

	result := NewFieldValue(cfg, options, fieldConfig, value)

	if result.value != expected.value {
		t.Errorf("Expected value to be %v, but got %v", expected.value, result.value)
	}
	if !result.enumValue.Equal(expected.enumValue) {
		t.Errorf("Expected enumValue to be %v, but got %v", expected.enumValue, result.enumValue)
	}
	if !result.timeValue.Equal(expected.timeValue) {
		t.Errorf("Expected timeValue to be %v, but got %v", expected.timeValue, result.timeValue)
	}
	if result.Output != expected.Output {
		t.Errorf("Expected Output to be %s, but got %s", expected.Output, result.Output)
	}
	if !result.Config.Equal(expected.Config) {
		t.Errorf("Expected Config to be %v, but got %v", expected.Config, result.Config)
	}
}

func TestNewFieldValue_CompressPrefix(t *testing.T) {
	cfg := config.Configuration{
		Field: config.Field{
			Type: config.FieldTypeString,
			Color: util.Color{
				R: 255,
				G: 0,
				B: 0,
			},
			CompressPrefix: config.CompressPrefix{
				Enabled: true,
				Prefix:  "prefix_",
			},
		},
	}
	options := Options{}
	fieldConfig := cfg.Field
	value := util.AnyValue{Text: "test"}

	expected := &FieldValueT{
		value:     value,
		enumValue: nil,
		timeValue: time.Time{},
		Output:    "prefix_test",
		Config:    fieldConfig,
	}

	result := NewFieldValue(cfg, options, fieldConfig, value)

	if result.value != expected.value {
		t.Errorf("Expected value to be %v, but got %v", expected.value, result.value)
	}
	if result.enumValue != expected.enumValue {
		t.Errorf("Expected enumValue to be %v, but got %v", expected.enumValue, result.enumValue)
	}
	if !result.timeValue.Equal(expected.timeValue) {
		t.Errorf("Expected timeValue to be %v, but got %v", expected.timeValue, result.timeValue)
	}
	if result.Output != expected.Output {
		t.Errorf("Expected Output to be %s, but got %s", expected.Output, result.Output)
	}
	if !result.Config.Equal(expected.Config) {
		t.Errorf("Expected Config to be %v, but got %v", expected.Config, result.Config)
	}
}

func TestNewFieldValue_Time(t *testing.T) {
	cfg := config.Configuration{
		Field: config.Field{
			Type: config.FieldTypeTime,
			Color: util.Color{
				R: 255,
				G: 0,
				B: 0,
			},
		},
	}
	options := Options{}
	fieldConfig := cfg.Field
	value := util.AnyValue{Text: "2023-10-01T12:00:00Z"}

	expected := &FieldValueT{
		value:     value,
		enumValue: nil,
		timeValue: time.Date(2023, 10, 1, 12, 0, 0, 0, time.UTC),
		Output:    "2023-10-01T12:00:00Z",
		Config:    fieldConfig,
	}

	result := NewFieldValue(cfg, options, fieldConfig, value)

	if result.value != expected.value {
		t.Errorf("Expected value to be %v, but got %v", expected.value, result.value)
	}
	if result.enumValue != expected.enumValue {
		t.Errorf("Expected enumValue to be %v, but got %v", expected.enumValue, result.enumValue)
	}
	if !result.timeValue.Equal(expected.timeValue) {
		t.Errorf("Expected timeValue to be %v, but got %v", expected.timeValue, result.timeValue)
	}
	if result.Output != expected.Output {
		t.Errorf("Expected Output to be %s, but got %s", expected.Output, result.Output)
	}
	if !result.Config.Equal(expected.Config) {
		t.Errorf("Expected Config to be %v, but got %v", expected.Config, result.Config)
	}
}

func TestNewFieldValue_Error(t *testing.T) {
	cfg := config.Configuration{
		Field: config.Field{
			Type: config.FieldTypeString,
			Color: util.Color{
				R: 255,
				G: 0,
				B: 0,
			},
		},
	}
	options := Options{}
	fieldConfig := cfg.Field
	value := util.AnyValue{Text: "invalid"}

	expected := &FieldValueT{
		value:     value,
		enumValue: nil,
		timeValue: time.Time{},
		Output:    "",
		Config:    fieldConfig,
	}

	result := NewFieldValue(cfg, options, fieldConfig, value)

	if result.value != expected.value {
		t.Errorf("Expected value to be %v, but got %v", expected.value, result.value)
	}
	if result.enumValue != expected.enumValue {
		t.Errorf("Expected enumValue to be %v, but got %v", expected.enumValue, result.enumValue)
	}
	if !result.timeValue.Equal(expected.timeValue) {
		t.Errorf("Expected timeValue to be %v, but got %v", expected.timeValue, result.timeValue)
	}
	if result.Output != expected.Output {
		t.Errorf("Expected Output to be %s, but got %s", expected.Output, result.Output)
	}
	if !result.Config.Equal(expected.Config) {
		t.Errorf("Expected Config to be %v, but got %v", expected.Config, result.Config)
	}
}

func TestParseTime_HappyPath(t *testing.T) {
	input := "2023-10-01T12:00:00Z"
	expected := time.Date(2023, 10, 1, 12, 0, 0, 0, time.UTC)

	result := ParseTime(input)

	if !result.Equal(expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestParseTime_Error(t *testing.T) {
	input := "invalid"

	expected := time.Time{}

	result := ParseTime(input)

	if !result.Equal(expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestIsTimeValid_HappyPath(t *testing.T) {
	input := "2023-10-01T12:00:00Z"
	expected := true

	result := IsTimeValid(input)

	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestIsTimeValid_Error(t *testing.T) {
	input := "invalid"

	expected := false

	result := IsTimeValid(input)

	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestIsValidEnum_HappyPath(t *testing.T) {
	input := "valid"
	expected := true

	result := IsValidEnum(input)

	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestIsValidEnum_Error(t *testing.T) {
	input := "invalid"

	expected := false

	result := IsValidEnum(input)

	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestIsValidType_HappyPath(t *testing.T) {
	input := "string"
	expected := true

	result := IsValidType(input)

	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestIsValidType_Error(t *testing.T) {
	input := "invalid"

	expected := false

	result := IsValidType(input)

	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestIsValidColor_HappyPath(t *testing.T) {
	input := "#FF0000"
	expected := true

	result := IsValidColor(input)

	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestIsValidColor_Error(t *testing.T) {
	input := "invalid"

	expected := false

	result := IsValidColor(input)

	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestIsValidConfig_HappyPath(t *testing.T) {
	input := Config{
		Type:     "string",
		Color:    "#FF0000",
		Timezone: "UTC",
		Enum:     []string{"valid"},
	}

	expected := true

	result := IsValidConfig(input)

	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestIsValidConfig_Error(t *testing.T) {
	input := Config{
		Type:     "invalid",
		Color:    "#FF0000",
		Timezone: "UTC",
		Enum:     []string{"valid"},
	}

	expected := false

	result := IsValidConfig(input)

	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
