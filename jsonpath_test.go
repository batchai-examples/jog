package jsonpath

import (
	"testing"
)

func TestGet_HappyPath(t *testing.T) {
	data := map[string]interface{}{
		"name": "John",
		"age":  30,
	}

	result, err := Get(data, "name")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result.(string) != "John" {
		t.Errorf("Expected 'John', got %v", result)
	}
}

func TestGet_NegativePath(t *testing.T) {
	data := map[string]interface{}{
		"name": "John",
		"age":  30,
	}

	result, err := Get(data, "address")
	if err == nil {
		t.Errorf("Expected error, got none")
	}
	if result != nil {
		t.Errorf("Expected nil, got %v", result)
	}
}

func TestGet_InvalidObject(t *testing.T) {
	data := 123

	result, err := Get(data, "name")
	if err == nil {
		t.Errorf("Expected error, got none")
	}
	if result != nil {
		t.Errorf("Expected nil, got %v", result)
	}
}

func TestGet_EmptyPath(t *testing.T) {
	data := map[string]interface{}{
		"name": "John",
		"age":  30,
	}

	result, err := Get(data, "")
	if err == nil {
		t.Errorf("Expected error, got none")
	}
	if result != nil {
		t.Errorf("Expected nil, got %v", result)
	}
}

func TestSet_HappyPath(t *testing.T) {
	data := map[string]interface{}{
		"name": "John",
		"age":  30,
	}

	err := Set(data, "name", "Jane")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if data["name"].(string) != "Jane" {
		t.Errorf("Expected 'Jane', got %v", data["name"])
	}
}

func TestSet_NegativePath(t *testing.T) {
	data := map[string]interface{}{
		"name": "John",
		"age":  30,
	}

	err := Set(data, "address", "123 Main St")
	if err == nil {
		t.Errorf("Expected error, got none")
	}
	if data["address"] != nil {
		t.Errorf("Expected nil, got %v", data["address"])
	}
}

func TestSet_InvalidObject(t *testing.T) {
	data := 123

	err := Set(data, "name", "Jane")
	if err == nil {
		t.Errorf("Expected error, got none")
	}
	if data != 123 {
		t.Errorf("Expected 123, got %v", data)
	}
}

func TestSet_EmptyPath(t *testing.T) {
	data := map[string]interface{}{
		"name": "John",
		"age":  30,
	}

	err := Set(data, "", "Jane")
	if err == nil {
		t.Errorf("Expected error, got none")
	}
	if data != map[string]interface{}{"name": "John", "age": 30} {
		t.Errorf("Expected {\"name\": \"John\", \"age\": 30}, got %v", data)
	}
}

func TestSet_NestedPath(t *testing.T) {
	data := map[string]interface{}{
		"user": map[string]interface{}{
			"name": "John",
			"age":  30,
		},
	}

	err := Set(data, "user.name", "Jane")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if data["user"].(map[string]interface{})["name"].(string) != "Jane" {
		t.Errorf("Expected 'Jane', got %v", data["user"].(map[string]interface{})["name"])
	}
}

func TestSet_NestedPath_NotMap(t *testing.T) {
	data := map[string]interface{}{
		"user": 123,
	}

	err := Set(data, "user.name", "Jane")
	if err == nil {
		t.Errorf("Expected error, got none")
	}
	if data["user"] != 123 {
		t.Errorf("Expected 123, got %v", data["user"])
	}
}
