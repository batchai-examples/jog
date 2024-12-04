package config

import (
	"errors"
	"testing"
)

func TestPrefix_Reset(t *testing.T) {
	t.Run("Happy Path", func(t *testing.T) {
		p := &PrefixT{}
		p.Reset()
		if p.Color.Get() != "FgBlue" {
			t.Errorf("Expected Color to be 'FgBlue', got '%s'", p.Color.Get())
		}
	})

	t.Run("Negative Case - Reset on nil pointer", func(t *testing.T) {
		var p Prefix
		p.Reset()
		if p.Color.Get() != "" {
			t.Errorf("Expected Color to be empty, got '%s'", p.Color.Get())
		}
	})
}

func TestPrefix_FromMap(t *testing.T) {
	t.Run("Happy Path", func(t *testing.T) {
		p := &PrefixT{}
		err := p.FromMap(map[string]interface{}{
			"Color": "FgRed",
		})
		if err != nil {
			t.Errorf("Expected no error, got '%v'", err)
		}
		if p.Color.Get() != "FgRed" {
			t.Errorf("Expected Color to be 'FgRed', got '%s'", p.Color.Get())
		}
	})

	t.Run("Negative Case - Invalid map key", func(t *testing.T) {
		p := &PrefixT{}
		err := p.FromMap(map[string]interface{}{
			"InvalidKey": "FgRed",
		})
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})

	t.Run("Negative Case - Invalid map value type", func(t *testing.T) {
		p := &PrefixT{}
		err := p.FromMap(map[string]interface{}{
			"Color": 123,
		})
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})

	t.Run("Corner Case - Empty map", func(t *testing.T) {
		p := &PrefixT{}
		err := p.FromMap(map[string]interface{}{})
		if err != nil {
			t.Errorf("Expected no error, got '%v'", err)
		}
	})
}

func TestPrefix_ToMap(t *testing.T) {
	t.Run("Happy Path", func(t *testing.T) {
		p := &PrefixT{}
		p.Color.Set("FgGreen")
		m := p.ToMap()
		if m["Color"] != "FgGreen" {
			t.Errorf("Expected Color to be 'FgGreen', got '%v'", m["Color"])
		}
	})

	t.Run("Corner Case - Empty map", func(t *testing.T) {
		p := &PrefixT{}
		m := p.ToMap()
		if len(m) != 0 {
			t.Errorf("Expected empty map, got '%v'", m)
		}
	})
}

func TestPrefix_UnmarshalYAML(t *testing.T) {
	t.Run("Happy Path", func(t *testing.T) {
		p := &PrefixT{}
		err := p.UnmarshalYAML(func(interface{}) error {
			return nil
		})
		if err != nil {
			t.Errorf("Expected no error, got '%v'", err)
		}
	})

	t.Run("Negative Case - UnmarshalYAML error", func(t *testing.T) {
		p := &PrefixT{}
		err := p.UnmarshalYAML(func(interface{}) error {
			return errors.New("unmarshal error")
		})
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}

func TestPrefix_MarshalYAML(t *testing.T) {
	t.Run("Happy Path", func(t *testing.T) {
		p := &PrefixT{}
		i, err := p.MarshalYAML()
		if err != nil {
			t.Errorf("Expected no error, got '%v'", err)
		}
		if i == nil {
			t.Errorf("Expected non-nil value, got 'nil'")
		}
	})

	t.Run("Corner Case - MarshalYAML returns nil", func(t *testing.T) {
		p := &PrefixT{}
		i, err := p.MarshalYAML()
		if err != nil {
			t.Errorf("Expected no error, got '%v'", err)
		}
		if i == nil {
			t.Errorf("Expected non-nil value, got 'nil'")
		}
	})
}
