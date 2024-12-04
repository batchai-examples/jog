package config

import (
	"errors"
	"testing"
)

type MockDynamicObject struct {
	CalledReset    bool
	CalledFromMap  bool
	CalledToMap    bool
	CalledInit     bool
	Map            map[string]interface{}
	ErrFromMap     error
	ErrMarshalYAML error
}

func (m *MockDynamicObject) Reset() {
	m.CalledReset = true
}

func (m *MockDynamicObject) FromMap(m map[string]interface{}) error {
	m.CalledFromMap = true
	m.Map = m
	return m.ErrFromMap
}

func (m *MockDynamicObject) ToMap() map[string]interface{} {
	m.CalledToMap = true
	return m.Map
}

func (m *MockDynamicObject) Init(cfg Configuration) {
	m.CalledInit = true
}

func TestUnmarshalYAML(t *testing.T) {
	t.Run("Happy Path", func(t *testing.T) {
		mock := &MockDynamicObject{}
		unmarshalFunc := func(i interface{}) error {
			return nil
		}
		err := UnmarshalYAML(mock, unmarshalFunc)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if !mock.CalledReset || !mock.CalledFromMap || !mock.CalledToMap {
			t.Errorf("Expected all methods to be called")
		}
	})

	t.Run("Error from FromMap", func(t *testing.T) {
		mock := &MockDynamicObject{ErrFromMap: errors.New("error from FromMap")}
		unmarshalFunc := func(i interface{}) error {
			return nil
		}
		err := UnmarshalYAML(mock, unmarshalFunc)
		if err == nil {
			t.Errorf("Expected an error, got none")
		}
		if !mock.CalledReset || mock.CalledFromMap || mock.CalledToMap {
			t.Errorf("Expected Reset to be called and FromMap and ToMap not to be called")
		}
	})

	t.Run("Error from unmarshal", func(t *testing.T) {
		mock := &MockDynamicObject{}
		unmarshalFunc := func(i interface{}) error {
			return errors.New("error from unmarshal")
		}
		err := UnmarshalYAML(mock, unmarshalFunc)
		if err == nil {
			t.Errorf("Expected an error, got none")
		}
		if !mock.CalledReset || mock.CalledFromMap || mock.CalledToMap {
			t.Errorf("Expected Reset to be called and FromMap and ToMap not to be called")
		}
	})

	t.Run("Empty map", func(t *testing.T) {
		mock := &MockDynamicObject{}
		unmarshalFunc := func(i interface{}) error {
			return nil
		}
		err := UnmarshalYAML(mock, unmarshalFunc)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if !mock.CalledReset || !mock.CalledFromMap || !mock.CalledToMap {
			t.Errorf("Expected all methods to be called")
		}
	})
}

func TestMarshalYAML(t *testing.T) {
	t.Run("Happy Path", func(t *testing.T) {
		mock := &MockDynamicObject{Map: map[string]interface{}{"key": "value"}}
		result, err := MarshalYAML(mock)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if result.(map[string]interface{})["key"] != "value" {
			t.Errorf("Expected result to be map with key 'key' and value 'value'")
		}
	})

	t.Run("Error from ToMap", func(t *testing.T) {
		mock := &MockDynamicObject{ErrMarshalYAML: errors.New("error from ToMap")}
		result, err := MarshalYAML(mock)
		if err == nil {
			t.Errorf("Expected an error, got none")
		}
		if result != nil {
			t.Errorf("Expected result to be nil")
		}
	})
}
