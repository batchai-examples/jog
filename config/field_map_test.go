package config

import (
	"testing"
)

func TestFieldMap_Reset(t *testing.T) {
	i := &FieldMapT{}
	i.Reset()

	if i.Others == nil || !i.Others.IsEmpty() {
		t.Errorf("Expected Others to be reset")
	}

	if len(i.Standards) != 0 {
		t.Errorf("Expected Standards to be reset")
	}

	if len(i.StandardsWithAllAliases) != 0 {
		t.Errorf("Expected StandardsWithAllAliases to be reset")
	}
}

func TestFieldMap_UnmarshalYAML(t *testing.T) {
	i := &FieldMapT{}
	data := []byte(`others: {}
standards:
  field1:
    name: field1
    case_sensitive: true
    alias:
      values: []
      lowercased_values: []
  field2:
    name: field2
    case_sensitive: false
    alias:
      values: [alias2]
      lowercased_values: []`)

	if err := i.UnmarshalYAML(func(interface{}) error { return nil }); err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(i.Standards) != 2 {
		t.Errorf("Expected 2 standards fields, got %d", len(i.Standards))
	}

	if len(i.StandardsWithAllAliases) != 3 {
		t.Errorf("Expected 3 standards with all aliases fields, got %d", len(i.StandardsWithAllAliases))
	}
}

func TestFieldMap_MarshalYAML(t *testing.T) {
	i := &FieldMapT{}
	data := []byte(`others: {}
standards:
  field1:
    name: field1
    case_sensitive: true
    alias:
      values: []
      lowercased_values: []
  field2:
    name: field2
    case_sensitive: false
    alias:
      values: [alias2]
      lowercased_values: []`)

	if err := i.UnmarshalYAML(func(interface{}) error { return nil }); err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	m, err := i.MarshalYAML()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	expected := `{"others":{},"standards":{"field1":{"name":"field1","case_sensitive":true,"alias":{"values":[],"lowercased_values":[]}},"field2":{"name":"field2","case_sensitive":false,"alias":{"values":["alias2"],"lowercased_values":[]}}}}`
	if !strings.EqualFold(fmt.Sprintf("%v", m), expected) {
		t.Errorf("Expected %s, got %s", expected, fmt.Sprintf("%v", m))
	}
}

func TestFieldMap_Init(t *testing.T) {
	i := &FieldMapT{}
	cfg := Configuration{}

	i.Init(cfg)

	if i.Others == nil || !i.Others.IsEmpty() {
		t.Errorf("Expected Others to be initialized")
	}

	if len(i.Standards) != 0 {
		t.Errorf("Expected Standards to be initialized")
	}

	if len(i.StandardsWithAllAliases) != 0 {
		t.Errorf("Expected StandardsWithAllAliases to be initialized")
	}
}

func TestFieldMap_FromMap(t *testing.T) {
	i := &FieldMapT{}
	m := map[string]interface{}{
		"others": map[string]interface{}{},
		"standards": map[string]interface{}{
			"field1": map[string]interface{}{
				"name":        "field1",
				"case_sensitive": true,
				"alias": map[string]interface{}{
					"values": []interface{}{},
					"lowercased_values": []interface{}{},
				},
			},
			"field2": map[string]interface{}{
				"name":        "field2",
				"case_sensitive": false,
				"alias": map[string]interface{}{
					"values": []interface{}{"alias2"},
					"lowercased_values": []interface{}{},
				},
			},
		},
	}

	if err := i.FromMap(m); err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(i.Standards) != 2 {
		t.Errorf("Expected 2 standards fields, got %d", len(i.Standards))
	}

	if len(i.StandardsWithAllAliases) != 3 {
		t.Errorf("Expected 3 standards with all aliases fields, got %d", len(i.StandardsWithAllAliases))
	}
}

func TestFieldMap_ToMap(t *testing.T) {
	i := &FieldMapT{}
	data := []byte(`others: {}
standards:
  field1:
    name: field1
    case_sensitive: true
    alias:
      values: []
      lowercased_values: []
  field2:
    name: field2
    case_sensitive: false
    alias:
      values: [alias2]
      lowercased_values: []`)

	if err := i.UnmarshalYAML(func(interface{}) error { return nil }); err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	m := i.ToMap()

	expected := map[string]interface{}{
		"others": map[string]interface{}{},
		"standards": map[string]interface{}{
			"field1": map[string]interface{}{
				"name":        "field1",
				"case_sensitive": true,
				"alias": map[string]interface{}{
					"values": []interface{}{},
					"lowercased_values": []interface{}{},
				},
			},
			"field2": map[string]interface{}{
				"name":        "field2",
				"case_sensitive": false,
				"alias": map[string]interface{}{
					"values": []interface{}{"alias2"},
					"lowercased_values": []interface{}{},
				},
			},
		},
	}

	if !reflect.DeepEqual(m, expected) {
		t.Errorf("Expected %v, got %v", expected, m)
	}
}

func TestFieldMap_FromMap_DuplicateAlias(t *testing.T) {
	i := &FieldMapT{}
	m := map[string]interface{}{
		"others": map[string]interface{}{},
		"standards": map[string]interface{}{
			"field1": map[string]interface{}{
				"name":        "field1",
				"case_sensitive": true,
				"alias": map[string]interface{}{
					"values": []interface{}{"alias1"},
					"lowercased_values": []interface{}{},
				},
			},
			"field2": map[string]interface{}{
				"name":        "field2",
				"case_sensitive": false,
				"alias": map[string]interface{}{
					"values": []interface{}{"alias1"},
					"lowercased_values": []interface{}{},
				},
			},
		},
	}

	if err := i.FromMap(m); err == nil {
		t.Errorf("Expected error, got none")
	}
}

func TestFieldMap_FromMap_InvalidAlias(t *testing.T) {
	i := &FieldMapT{}
	m := map[string]interface{}{
		"others": map[string]interface{}{},
		"standards": map[string]interface{}{
			"field1": map[string]interface{}{
				"name":        "field1",
				"case_sensitive": true,
				"alias": map[string]interface{}{
					"values": []interface{}{"alias1"},
					"lowercased_values": []interface{}{},
				},
			},
			"field2": map[string]interface{}{
				"name":        "field2",
				"case_sensitive": false,
				"alias": map[string]interface{}{
					"values": []interface{}{"invalid_alias"},
					"lowercased_values": []interface{}{},
				},
			},
		},
	}

	if err := i.FromMap(m); err == nil {
		t.Errorf("Expected error, got none")
	}
}

func TestFieldMap_FromMap_InvalidCaseSensitive(t *testing.T) {
	i := &FieldMapT{}
	m := map[string]interface{}{
		"others": map[string]interface{}{},
		"standards": map[string]interface{}{
			"field1": map[string]interface{}{
				"name":        "field1",
				"case_sensitive": "invalid_value",
				"alias": map[string]interface{}{
					"values": []interface{}{"alias1"},
					"lowercased_values": []interface{}{},
				},
			},
		},
	}

	if err := i.FromMap(m); err == nil {
		t.Errorf("Expected error, got none")
	}
}
