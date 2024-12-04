package config

import (
	"testing"
)

func TestEnumT(t *testing.T) {
	t.Run("Happy Path", func(t *testing.T) {
		// Given
		enum := &EnumT{
			Name:  "example",
			Alias: util.NewMultiString("alias1", "alias2"),
			Color: util.NewColor(255, 0, 0),
		}

		// When
		m := enum.ToMap()

		// Then
		if m["name"] != "example" {
			t.Errorf("Expected name to be 'example', got %v", m["name"])
		}
		if len(m["alias"].([]interface{})) != 2 || m["alias"].([]interface{})[0] != "alias1" || m["alias"].([]interface{})[1] != "alias2" {
			t.Errorf("Expected alias to be ['alias1', 'alias2'], got %v", m["alias"])
		}
		if m["color"] != "rgb(255,0,0)" {
			t.Errorf("Expected color to be 'rgb(255,0,0)', got %v", m["color"])
		}
	})

	t.Run("Negative Path - Empty Name", func(t *testing.T) {
		// Given
		enum := &EnumT{
			Name:  "",
			Alias: util.NewMultiString("alias1"),
			Color: util.NewColor(0, 255, 0),
		}

		// When
		m := enum.ToMap()

		// Then
		if m["name"] != "" {
			t.Errorf("Expected name to be '', got %v", m["name"])
		}
		if len(m["alias"].([]interface{})) != 1 || m["alias"].([]interface{})[0] != "alias1" {
			t.Errorf("Expected alias to be ['alias1'], got %v", m["alias"])
		}
		if m["color"] != "rgb(0,255,0)" {
			t.Errorf("Expected color to be 'rgb(0,255,0)', got %v", m["color"])
		}
	})

	t.Run("Negative Path - Empty Alias", func(t *testing.T) {
		// Given
		enum := &EnumT{
			Name:  "example",
			Alias: util.NewMultiString(),
			Color: util.NewColor(0, 0, 255),
		}

		// When
		m := enum.ToMap()

		// Then
		if m["name"] != "example" {
			t.Errorf("Expected name to be 'example', got %v", m["name"])
		}
		if len(m["alias"].([]interface{})) != 0 {
			t.Errorf("Expected alias to be [], got %v", m["alias"])
		}
		if m["color"] != "rgb(0,0,255)" {
			t.Errorf("Expected color to be 'rgb(0,0,255)', got %v", m["color"])
		}
	})

	t.Run("Negative Path - Empty Color", func(t *testing.T) {
		// Given
		enum := &EnumT{
			Name:  "example",
			Alias: util.NewMultiString("alias1"),
			Color: util.NewColor(0, 0, 0),
		}

		// When
		m := enum.ToMap()

		// Then
		if m["name"] != "example" {
			t.Errorf("Expected name to be 'example', got %v", m["name"])
		}
		if len(m["alias"].([]interface{})) != 1 || m["alias"].([]interface{})[0] != "alias1" {
			t.Errorf("Expected alias to be ['alias1'], got %v", m["alias"])
		}
		if m["color"] != "rgb(0,0,0)" {
			t.Errorf("Expected color to be 'rgb(0,0,0)', got %v", m["color"])
		}
	})

	t.Run("Corner Case - Nil Alias", func(t *testing.T) {
		// Given
		enum := &EnumT{
			Name:  "example",
			Alias: nil,
			Color: util.NewColor(255, 255, 0),
		}

		// When
		m := enum.ToMap()

		// Then
		if m["name"] != "example" {
			t.Errorf("Expected name to be 'example', got %v", m["name"])
		}
		if len(m["alias"].([]interface{})) != 0 {
			t.Errorf("Expected alias to be [], got %v", m["alias"])
		}
		if m["color"] != "rgb(255,255,0)" {
			t.Errorf("Expected color to be 'rgb(255,255,0)', got %v", m["color"])
		}
	})

	t.Run("Corner Case - Nil Color", func(t *testing.T) {
		// Given
		enum := &EnumT{
			Name:  "example",
			Alias: util.NewMultiString("alias1"),
			Color: nil,
		}

		// When
		m := enum.ToMap()

		// Then
		if m["name"] != "example" {
			t.Errorf("Expected name to be 'example', got %v", m["name"])
		}
		if len(m["alias"].([]interface{})) != 1 || m["alias"].([]interface{})[0] != "alias1" {
			t.Errorf("Expected alias to be ['alias1'], got %v", m["alias"])
		}
		if m["color"] != "" {
			t.Errorf("Expected color to be '', got %v", m["color"])
		}
	})

	t.Run("Corner Case - Empty Map Input", func(t *testing.T) {
		// Given
		enum := &EnumT{}
		m := map[string]interface{}{}

		// When
		err := enum.FromMap(m)

		// Then
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if enum.Name != "" {
			t.Errorf("Expected name to be '', got %v", enum.Name)
		}
		if len(enum.Alias) != 0 {
			t.Errorf("Expected alias to be [], got %v", enum.Alias)
		}
		if enum.Color != nil {
			t.Errorf("Expected color to be nil, got %v", enum.Color)
		}
	})

	t.Run("Corner Case - Map with Only Name Input", func(t *testing.T) {
		// Given
		enum := &EnumT{}
		m := map[string]interface{}{
			"name": "example",
		}

		// When
		err := enum.FromMap(m)

		// Then
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if enum.Name != "example" {
			t.Errorf("Expected name to be 'example', got %v", enum.Name)
		}
		if len(enum.Alias) != 0 {
			t.Errorf("Expected alias to be [], got %v", enum.Alias)
		}
		if enum.Color != nil {
			t.Errorf("Expected color to be nil, got %v", enum.Color)
		}
	})

	t.Run("Corner Case - Map with Only Alias Input", func(t *testing.T) {
		// Given
		enum := &EnumT{}
		m := map[string]interface{}{
			"alias": []interface{}{"alias1"},
		}

		// When
		err := enum.FromMap(m)

		// Then
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if enum.Name != "" {
			t.Errorf("Expected name to be '', got %v", enum.Name)
		}
		if len(enum.Alias) != 1 || enum.Alias[0] != "alias1" {
			t.Errorf("Expected alias to be ['alias1'], got %v", enum.Alias)
		}
		if enum.Color != nil {
			t.Errorf("Expected color to be nil, got %v", enum.Color)
		}
	})

	t.Run("Corner Case - Map with Only Color Input", func(t *testing.T) {
		// Given
		enum := &EnumT{}
		m := map[string]interface{}{
			"color": []interface{}{255, 255, 0},
		}

		// When
		err := enum.FromMap(m)

		// Then
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if enum.Name != "" {
			t.Errorf("Expected name to be '', got %v", enum.Name)
		}
		if len(enum.Alias) != 0 {
			t.Errorf("Expected alias to be [], got %v", enum.Alias)
		}
		if enum.Color == nil || *enum.Color != util.Color{R: 255, G: 255, B: 0} {
			t.Errorf("Expected color to be rgb(255,255,0), got %v", enum.Color)
		}
	})

	t.Run("Corner Case - Map with All Inputs", func(t *testing.T) {
		// Given
		enum := &EnumT{}
		m := map[string]interface{}{
			"name":  "example",
			"alias": []interface{}{"alias1"},
			"color": []interface{}{255, 0, 0},
		}

		// When
		err := enum.FromMap(m)

		// Then
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if enum.Name != "example" {
			t.Errorf("Expected name to be 'example', got %v", enum.Name)
		}
		if len(enum.Alias) != 1 || enum.Alias[0] != "alias1" {
			t.Errorf("Expected alias to be ['alias1'], got %v", enum.Alias)
		}
		if enum.Color == nil || *enum.Color != util.Color{R: 255, G: 0, B: 0} {
			t.Errorf("Expected color to be rgb(255,0,0), got %v", enum.Color)
		}
	})

	t.Run("Corner Case - Map with Invalid Color Input", func(t *testing.T) {
		// Given
		enum := &EnumT{}
		m := map[string]interface{}{
			"name":  "example",
			"alias": []interface{}{"alias1"},
			"color": []interface{}{256, 0, 0},
		}

		// When
		err := enum.FromMap(m)

		// Then
		if err == nil {
			t.Errorf("Expected error for invalid color input")
		}
	})

	t.Run("Corner Case - Map with Invalid Alias Input", func(t *testing.T) {
		// Given
		enum := &EnumT{}
		m := map[string]interface{}{
			"name":  "example",
			"alias": []interface{}{"alias1", 2},
		}

		// When
		err := enum.FromMap(m)

		// Then
		if err == nil {
			t.Errorf("Expected error for invalid alias input")
		}
	})

	t.Run("Corner Case - Map with Invalid Name Input", func(t *testing.T) {
		// Given
		enum := &EnumT{}
		m := map[string]interface{}{
			"name":  123,
			"alias": []interface{}{"alias1"},
			"color": []interface{}{255, 0, 0},
		}

		// When
		err := enum.FromMap(m)

		// Then
		if err == nil {
			t.Errorf("Expected error for invalid name input")
		}
	})

	t.Run("Corner Case - Map with Missing Color Input", func(t *testing.T) {
		// Given
		enum := &EnumT{}
		m := map[string]interface{}{
			"name":  "example",
			"alias": []interface{}{"alias1"},
		}

		// When
		err := enum.FromMap(m)

		// Then
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})

	t.Run("Corner Case - Map with Missing Alias Input", func(t *testing.T) {
		// Given
		enum := &EnumT{}
		m := map[string]interface{}{
			"name":  "example",
			"color": []interface{}{255, 0, 0},
		}

		// When
		err := enum.FromMap(m)

		// Then
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})

	t.Run("Corner Case - Map with Missing Name Input", func(t *testing.T) {
		// Given
		enum := &EnumT{}
		m := map[string]interface{}{
			"alias": []interface{}{"alias1"},
			"color": []interface{}{255, 0, 0},
		}

		// When
		err := enum.FromMap(m)

		// Then
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})

	t.Run("Corner Case - Map with Empty Color Input", func(t *testing.T) {
		// Given
		enum := &EnumT{}
		m := map[string]interface{}{
			"name":  "example",
			"alias": []interface{}{"alias1"},
			"color": []interface{}{},
		}

		// When
		err := enum.FromMap(m)

		// Then
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})

	t.Run("Corner Case - Map with Empty Alias Input", func(t *testing.T) {
		// Given
		enum := &EnumT{}
		m := map[string]interface{}{
			"name":  "example",
			"alias": []interface{}{},
			"color": []interface{}{255, 0, 0},
		}

		// When
		err := enum.FromMap(m)

		// Then
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})

	t.Run("Corner Case - Map with Empty Name Input", func(t *testing.T) {
		// Given
		enum := &EnumT{}
		m := map[string]interface{}{
			"name":  "",
			"alias": []interface{}{"alias1"},
			"color": []interface{}{255, 0, 0},
		}

		// When
		err := enum.FromMap(m)

		// Then
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})

	t.Run("Corner Case - Map with Null Color Input", func(t *testing.T) {
		// Given
		enum := &EnumT{}
		m := map[string]interface{}{
			"name":  "example",
			"alias": []interface{}{"alias1"},
			"color": nil,
		}

		// When
		err := enum.FromMap(m)

		// Then
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})

	t.Run("Corner Case - Map with Null Alias Input", func(t *testing.T) {
		// Given
		enum := &EnumT{}
		m := map[string]interface{}{
			"name":  "example",
			"alias": nil,
			"color": []interface{}{255, 0, 0},
		}

		// When
		err := enum.FromMap(m)

		// Then
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})

	t.Run("Corner Case - Map with Null Name Input", func(t *testing.T) {
		// Given
		enum := &EnumT{}
		m := map[string]interface{}{
			"name":  nil,
			"alias": []interface{}{"alias1"},
			"color": []interface{}{255, 0, 0},
		}

		// When
		err := enum.FromMap(m)

		// Then
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})

	t.Run("Corner Case - Map with Empty String Color Input", func(t *testing.T) {
		// Given
		enum := &EnumT{}
		m := map[string]interface{}{
			"name":  "example",
			"alias": []interface{}{"alias1"},
			"color": "",
		}

		// When
		err := enum.FromMap(m)

		// Then
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})

	t.Run("Corner Case - Map with Empty String Alias Input", func(t *testing.T) {
		// Given
		enum := &EnumT{}
		m := map[string]interface{}{
			"name":  "example",
			"alias": "",
			"color": []interface{}{255, 0, 0},
		}

		// When
		err := enum.FromMap(m)

		// Then
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})

	t.Run("Corner Case - Map with Empty String Name Input", func(t *testing.T) {
		// Given
		enum := &EnumT{}
		m := map[string]interface{}{
			"name":  "",
			"alias": []interface{}{"alias1"},
			"color": []interface{}{255, 0, 0},
		}

		// When
		err := enum.FromMap(m)

		// Then
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})

	t.Run("Corner Case - Map with Invalid Color Format Input", func(t *testing.T) {
		// Given
		enum := &EnumT{}
		m := map[string]interface{}{
			"name":  "example",
			"alias": []interface{}{"alias1"},
			"color": []interface{}{256, 0, 0},
		}

		// When
		err := enum.FromMap(m)

		// Then
		if err == nil {
			t.Errorf("Expected error for invalid color format input")
		}
	})

	t.Run("Corner Case - Map with Invalid Alias Format Input", func(t *testing.T) {
		// Given
		enum := &EnumT{}
		m := map[string]interface{}{
			"name":  "example",
			"alias": []interface{}{"alias1", "invalid"},
		}

		// When
		err := enum.FromMap(m)

		// Then
		if err == nil {
			t.Errorf("Expected error for invalid alias format input")
		}
	})

	t.Run("Corner Case - Map with Invalid Name Format Input", func(t *testing.T) {
		// Given
		enum := &EnumT{}
		m := map[string]interface{}{
			"name":  []interface{}{"example"},
			"alias": []interface{}{"alias1"},
			"color": []interface{}{255, 0, 0},
		}

		// When
		err := enum.FromMap(m)

		// Then
		if err == nil {
			t.Errorf("Expected error for invalid name format input")
		}
	})

	t.Run("Corner Case - Map with Invalid Color Type Input", func(t *testing.T) {
		// Given
		enum := &EnumT{}
		m := map[string]interface{}{
			"name":  "example",
			"alias": []interface{}{"alias1"},
			"color": "invalid",
		}

		// When
		err := enum.FromMap(m)

		// Then
		if err == nil {
			t.Errorf("Expected error for invalid color type input")
		}
	})

	t.Run("Corner Case - Map with Invalid Alias Type Input", func(t *testing.T) {
		// Given
		enum := &EnumT{}
		m := map[string]interface{}{
			"name":  "example",
			"alias": "invalid",
			"color": []interface{}{255, 0, 0},
		}

		// When
		err := enum.FromMap(m)

		// Then
		if err == nil {
			t.Errorf("Expected error for invalid alias type input")
		}
	})

	t.Run("Corner Case - Map with Invalid Name Type Input", func(t *testing.T) {
		// Given
		enum := &EnumT{}
		m := map[string]interface{}{
			"name":  []interface{}{"example"},
			"alias": []interface{}{"alias1"},
			"color": []interface{}{255, 0, 0},
		}

		// When
		err := enum.FromMap(m)

		// Then
		if err == nil {
			t.Errorf("Expected error for invalid name type input")
		}
	})

	t.Run("Corner Case - Map with Invalid Color Length Input", func(t *testing.T) {
		// Given
		enum := &EnumT{}
		m := map[string]interface{}{
			"name":  "example",
			"alias": []interface{}{"alias1"},
			"color": []interface{}{255, 0},
		}

		// When
		err := enum.FromMap(m)

		// Then
		if err == nil {
			t.Errorf("Expected error for invalid color length input")
		}
	})

	t.Run("Corner Case - Map with Invalid Alias Length Input", func(t *testing.T) {
		// Given
		enum := &EnumT{}
		m := map[string]interface{}{
			"name":  "example",
			"alias": []interface{}{"alias1", "alias2", "alias3"},
			"color": []interface{}{255, 0, 0},
		}

		// When
		err := enum.FromMap(m)

		// Then
		if err == nil {
			t.Errorf("Expected error for invalid alias length input")
		}
	})

	t.Run("Corner Case - Map with Invalid Name Length Input", func(t *testing.T) {
		// Given
		enum := &EnumT{}
		m := map[string]interface{}{
			"name":  []interface{}{"example1", "example2"},
			"alias": []interface{}{"alias1"},
			"color": []interface{}{255, 0, 0},
		}

		// When
		err := enum.FromMap(m)

		// Then
		if err == nil {
			t.Errorf("Expected error for invalid name length input")
		}
	})

	t.Run("Corner Case - Map with Invalid Color Value Input", func(t *testing.T) {
		// Given
		enum := &EnumT{}
		m := map[string]interface{}{
			"name":  "example",
			"alias": []interface{}{"alias1"},
			"color": []interface{}{-1, 0, 0},
		}

		// When
		err := enum.FromMap(m)

		// Then
		if err == nil {
			t.Errorf("Expected error for invalid color value input")
		}
	})

	t.Run("Corner Case - Map with Invalid Alias Value Input", func(t *testing.T) {
		// Given
		enum := &EnumT{}
		m := map[string]interface{}{
			"name":  "example",
			"alias": []interface{}{"invalid"},
			"color": []interface{}{255, 0, 0},
		}

		// When
		err := enum.FromMap(m)

		// Then
		if err == nil {
			t.Errorf("Expected error for invalid alias value input")
		}
	})

	t.Run("Corner Case - Map with Invalid Name Value Input", func(t *testing.T) {
		// Given
		enum := &EnumT{}
		m := map[string]interface{}{
			"name":  []interface{}{"example"},
			"alias": []interface{}{"alias1"},
			"color": []interface{}{255, 0, 0},
		}

		// When
		err := enum.FromMap(m)

		// Then
		if err == nil {
			t.Errorf("Expected error for invalid name value input")
		}
	})

	t.Run("Corner Case - Map with Invalid Color Type Input", func(t *testing.T) {
		// Given
		enum := &EnumT{}
		m := map[string]interface{}{
			"name":  "example",
			"alias": []interface{}{"alias1"},
			"color": "invalid",
		}

		// When
		err := enum.FromMap(m)

		// Then
		if err == nil {
			t.Errorf("Expected error for invalid color type input")
		}
	})

	t.Run("Corner Case - Map with Invalid Alias Type Input", func(t *testing.T) {
		// Given
		enum := &EnumT{}
		m := map[string]interface{}{
			"name":  "example",
			"alias": "invalid",
			"color": []interface{}{255, 0, 0},
		}

		// When
		err := enum.FromMap(m)

		// Then
		if err == nil {
			t.Errorf("Expected error for invalid alias type input")
		}
	})

	t.Run("Corner Case - Map with Invalid Name Type Input", func(t *testing.T) {
		// Given
		enum := &EnumT{}
		m := map[string]interface{}{
			"name":  []interface{}{"example"},
			"alias": []interface{}{"alias1"},
			"color": []interface{}{255, 0, 0},
		}

		// When
		err := enum.FromMap(m)

		// Then
		if err == nil {
			t.Errorf("Expected error for invalid name type input")
		}
	})

	t.Run("Corner Case - Map with Invalid Color Length Input", func(t *testing.T) {
		// Given
		enum := &EnumT{}
		m := map[string]interface{}{
			"name":  "example",
			"alias": []interface{}{"alias1"},
			"color": []interface{}{255, 0},
		}

		// When
		err := enum.FromMap(m)

		// Then
		if err == nil {
			t.Errorf("Expected error for invalid color length input")
		}
	})

	t.Run("Corner Case - Map with Invalid Alias Length Input", func(t *testing.T) {
		// Given
		enum := &EnumT{}
		m := map[string]interface{}{
			"name":  "example",
			"alias": []interface{}{"alias1", "alias2", "alias3"},
			"color": []interface{}{255, 0, 0},
		}

		// When
		err := enum.FromMap(m)

		// Then
		if err == nil {
			t.Errorf("Expected error for invalid alias length input")
		}
	})

	t.Run("Corner Case - Map with Invalid Name Length Input", func(t *testing.T) {
		// Given
		enum := &EnumT{}
		m := map[string]interface{}{
			"name":  []interface{}{"example1", "example2"},
			"alias": []interface{}{"alias1"},
			"color": []interface{}{255, 0, 0},
		}

		// When
		err := enum.FromMap(m)

		// Then
		if err == nil {
			t.Errorf("Expected error for invalid name length input")
		}
	})

	t.Run("Corner Case - Map with Invalid Color Value Input", func(t *testing.T) {
		// Given
		enum := &EnumT{}
		m := map[string]interface{}{
			"name":  "example",
			"alias": []interface{}{"alias1"},
			"color": []interface{}{-1, 0, 0},
		}

		// When
		err := enum.FromMap(m)

		// Then
		if err == nil {
			t.Errorf("Expected error for invalid color value input")
		}
	})

	t.Run("Corner Case - Map with Invalid Alias Value Input", func(t *testing.T) {
		// Given
		enum := &EnumT{}
		m := map[string]interface{}{
			"name":  "example",
			"alias": []interface{}{"invalid"},
			"color": []interface{}{255, 0, 0},
		}

		// When
		err := enum.FromMap(m)

		// Then
		if err == nil {
			t.Errorf("Expected error for invalid alias value input")
		}
	})

	t.Run("Corner Case - Map with Invalid Name Value Input", func(t *testing.T) {
		// Given
		enum := &EnumT{}
		m := map[string]interface{}{
			"name":  []interface{}{"example"},
			"alias": []interface{}{"alias1"},
			"color": []interface{}{255, 0, 0},
		}

		// When
		err := enum.FromMap(m)

		// Then
