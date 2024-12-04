package jsonpath

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestGet_HappyPath(t *testing.T) {
	data := map[string]interface{}{
		"user": map[string]interface{}{
			"firstname": "seth",
			"lastname":  "rogen",
		},
		"age": 35,
		"filmography": map[string]interface{}{
			"movies": []string{
				"This Is The End",
				"Superbad",
				"Neighbors",
			},
		},
	}

	result, err := Get(data, "user.firstname")
	if err != nil {
		t.Errorf("failed to get user.firstname: %v", err)
	}
	if result != "seth" {
		t.Errorf("wrong get value, wanted %v, got %v", "seth", result)
	}

	result, err = Get(data, "filmography.movies[1]")
	if err != nil {
		t.Errorf("failed to get filmography.movies[1]: %v", err)
	}
	if result != "Superbad" {
		t.Errorf("wrong get value, wanted %v, got %v", "Superbad", result)
	}

	result, err = Get(data, "age")
	if err != nil {
		t.Errorf("failed to get age: %v", err)
	}
	if result != 35 {
		t.Errorf("wrong get value, wanted %v, got %v", 35, result)
	}
}

func TestGet_NonExistentField(t *testing.T) {
	data := map[string]interface{}{
		"user": map[string]interface{}{
			"firstname": "seth",
			"lastname":  "rogen",
		},
		"age": 35,
		"filmography": map[string]interface{}{
			"movies": []string{
				"This Is The End",
				"Superbad",
				"Neighbors",
			},
		},
	}

	result, err := Get(data, "where.is.this")
	if result != nil {
		t.Errorf("expected nil result for non-existent field, got %v", result)
	}
	if _, ok := err.(DoesNotExist); !ok {
		t.Errorf("expected DoesNotExist error for non-existent field, got %v", err)
	}
}

func TestGet_EmptyStringPath(t *testing.T) {
	data := map[string]interface{}{
		"user": map[string]interface{}{
			"firstname": "seth",
			"lastname":  "rogen",
		},
		"age": 35,
		"filmography": map[string]interface{}{
			"movies": []string{
				"This Is The End",
				"Superbad",
				"Neighbors",
			},
		},
	}

	result, err := Get(data, "")
	if result != nil {
		t.Errorf("expected nil result for empty string path, got %v", result)
	}
	if err != nil {
		t.Errorf("expected no error for empty string path, got %v", err)
	}
}

func TestGet_EmptyArrayPath(t *testing.T) {
	data := map[string]interface{}{
		"user": map[string]interface{}{
			"firstname": "seth",
			"lastname":  "rogen",
		},
		"age": 35,
		"filmography": map[string]interface{}{
			"movies": []string{
				"This Is The End",
				"Superbad",
				"Neighbors",
			},
		},
	}

	result, err := Get(data, "filmography.movies[]")
	if result != nil {
		t.Errorf("expected nil result for empty array path, got %v", result)
	}
	if err != nil {
		t.Errorf("expected no error for empty array path, got %v", err)
	}
}

func TestGet_ArrayIndexOutOfBounds(t *testing.T) {
	data := map[string]interface{}{
		"user": map[string]interface{}{
			"firstname": "seth",
			"lastname":  "rogen",
		},
		"age": 35,
		"filmography": map[string]interface{}{
			"movies": []string{
				"This Is The End",
				"Superbad",
				"Neighbors",
			},
		},
	}

	result, err := Get(data, "filmography.movies[10]")
	if result != nil {
		t.Errorf("expected nil result for out-of-bounds array index, got %v", result)
	}
	if _, ok := err.(DoesNotExist); !ok {
		t.Errorf("expected DoesNotExist error for out-of-bounds array index, got %v", err)
	}
}

func TestGet_NestedMap(t *testing.T) {
	data := map[string]interface{}{
		"user": map[string]interface{}{
			"firstname": "seth",
			"lastname":  "rogen",
			"address": map[string]string{
				"city":    "New York",
				"country": "USA",
			},
		},
	}

	result, err := Get(data, "user.address.city")
	if err != nil {
		t.Errorf("failed to get user.address.city: %v", err)
	}
	if result != "New York" {
		t.Errorf("wrong get value, wanted %v, got %v", "New York", result)
	}

	result, err = Get(data, "user.address.country")
	if err != nil {
		t.Errorf("failed to get user.address.country: %v", err)
	}
	if result != "USA" {
		t.Errorf("wrong get value, wanted %v, got %v", "USA", result)
	}
}

func TestGet_NestedMapNonExistentField(t *testing.T) {
	data := map[string]interface{}{
		"user": map[string]interface{}{
			"firstname": "seth",
			"lastname":  "rogen",
			"address": map[string]string{
				"city":    "New York",
				"country": "USA",
			},
		},
	}

	result, err := Get(data, "user.address.zip")
	if result != nil {
		t.Errorf("expected nil result for non-existent nested field, got %v", result)
	}
	if _, ok := err.(DoesNotExist); !ok {
		t.Errorf("expected DoesNotExist error for non-existent nested field, got %v", err)
	}
}

func TestGet_NestedMapEmptyStringPath(t *testing.T) {
	data := map[string]interface{}{
		"user": map[string]interface{}{
			"firstname": "seth",
			"lastname":  "rogen",
			"address": map[string]string{
				"city":    "New York",
				"country": "USA",
			},
		},
	}

	result, err := Get(data, "user.address.")
	if result != nil {
		t.Errorf("expected nil result for empty string path in nested map, got %v", result)
	}
	if err != nil {
		t.Errorf("expected no error for empty string path in nested map, got %v", err)
	}
}

func TestGet_NestedMapEmptyArrayPath(t *testing.T) {
	data := map[string]interface{}{
		"user": map[string]interface{}{
			"firstname": "seth",
			"lastname":  "rogen",
			"address": map[string]string{
				"city":    "New York",
				"country": "USA",
			},
		},
	}

	result, err := Get(data, "user.address[]")
	if result != nil {
		t.Errorf("expected nil result for empty array path in nested map, got %v", result)
	}
	if err != nil {
		t.Errorf("expected no error for empty array path in nested map, got %v", err)
	}
}

func TestGet_NestedMapArrayIndexOutOfBounds(t *testing.T) {
	data := map[string]interface{}{
		"user": map[string]interface{}{
			"firstname": "seth",
			"lastname":  "rogen",
			"address": map[string]string{
				"city":    "New York",
				"country": "USA",
			},
		},
	}

	result, err := Get(data, "user.address[10]")
	if result != nil {
		t.Errorf("expected nil result for out-of-bounds array index in nested map, got %v", result)
	}
	if _, ok := err.(DoesNotExist); !ok {
		t.Errorf("expected DoesNotExist error for out-of-bounds array index in nested map, got %v", err)
	}
}

func TestGet_NestedMapNestedArray(t *testing.T) {
	data := map[string]interface{}{
		"user": map[string]interface{}{
			"firstname": "seth",
			"lastname":  "rogen",
			"address": []map[string]string{
				{"city": "New York", "country": "USA"},
				{"city": "Los Angeles", "country": "USA"},
			},
		},
	}

	result, err := Get(data, "user.address[0].city")
	if err != nil {
		t.Errorf("failed to get user.address[0].city: %v", err)
	}
	if result != "New York" {
		t.Errorf("wrong get value, wanted %v, got %v", "New York", result)
	}

	result, err = Get(data, "user.address[1].country")
	if err != nil {
		t.Errorf("failed to get user.address[1].country: %v", err)
	}
	if result != "USA" {
		t.Errorf("wrong get value, wanted %v, got %v", "USA", result)
	}
}

func TestGet_NestedMapNestedArrayNonExistentField(t *testing.T) {
	data := map[string]interface{}{
		"user": map[string]interface{}{
			"firstname": "seth",
			"lastname":  "rogen",
			"address": []map[string]string{
				{"city": "New York", "country": "USA"},
				{"city": "Los Angeles", "country": "USA"},
			},
		},
	}

	result, err := Get(data, "user.address[0].zip")
	if result != nil {
		t.Errorf("expected nil result for non-existent nested field in nested array, got %v", result)
	}
	if _, ok := err.(DoesNotExist); !ok {
		t.Errorf("expected DoesNotExist error for non-existent nested field in nested array, got %v", err)
	}
}

func TestGet_NestedMapNestedArrayEmptyStringPath(t *testing.T) {
	data := map[string]interface{}{
		"user": map[string]interface{}{
			"firstname": "seth",
			"lastname":  "rogen",
			"address": []map[string]string{
				{"city": "New York", "country": "USA"},
				{"city": "Los Angeles", "country": "USA"},
			},
		},
	}

	result, err := Get(data, "user.address[0].")
	if result != nil {
		t.Errorf("expected nil result for empty string path in nested map and nested array, got %v", result)
	}
	if err != nil {
		t.Errorf("expected no error for empty string path in nested map and nested array, got %v", err)
	}
}

func TestGet_NestedMapNestedArrayEmptyArrayPath(t *testing.T) {
	data := map[string]interface{}{
		"user": map[string]interface{}{
			"firstname": "seth",
			"lastname":  "rogen",
			"address": []map[string]string{
				{"city": "New York", "country": "USA"},
				{"city": "Los Angeles", "country": "USA"},
			},
		},
	}

	result, err := Get(data, "user.address[].")
	if result != nil {
		t.Errorf("expected nil result for empty array path in nested map and nested array, got %v", result)
	}
	if err != nil {
		t.Errorf("expected no error for empty array path in nested map and nested array, got %v", err)
	}
}

func TestGet_NestedMapNestedArrayArrayIndexOutOfBounds(t *testing.T) {
	data := map[string]interface{}{
		"user": map[string]interface{}{
			"firstname": "seth",
			"lastname":  "rogen",
			"address": []map[string]string{
				{"city": "New York", "country": "USA"},
				{"city": "Los Angeles", "country": "USA"},
			},
		},
	}

	result, err := Get(data, "user.address[10].")
	if result != nil {
		t.Errorf("expected nil result for out-of-bounds array index in nested map and nested array, got %v", result)
	}
	if _, ok := err.(DoesNotExist); !ok {
		t.Errorf("expected DoesNotExist error for out-of-bounds array index in nested map and nested array, got %v", err)
	}
}

func TestGet_NestedMapNestedArrayNestedArray(t *testing.T) {
	data := map[string]interface{}{
		"user": map[string]interface{}{
			"firstname": "seth",
			"lastname":  "rogen",
			"address": []map[string]string{
				{"city": "New York", "country": "USA"},
				{"city": "Los Angeles", "country": "USA"},
			},
		},
	}

	result, err := Get(data, "user.address[0].coordinates[0]")
	if err != nil {
		t.Errorf("failed to get user.address[0].coordinates[0]: %v", err)
	}
	if result != 40.7128 {
		t.Errorf("wrong get value, wanted %v, got %v", 40.7128, result)
	}

	result, err = Get(data, "user.address[1].coordinates[1]")
	if err != nil {
		t.Errorf("failed to get user.address[1].coordinates[1]: %v", err)
	}
	if result != -118.2437 {
		t.Errorf("wrong get value, wanted %v, got %v", -118.2437, result)
	}
}

func TestGet_NestedMapNestedArrayNestedArrayNonExistentField(t *testing.T) {
	data := map[string]interface{}{
		"user": map[string]interface{}{
			"firstname": "seth",
			"lastname":  "rogen",
			"address": []map[string]string{
				{"city": "New York", "country": "USA"},
				{"city": "Los Angeles", "country": "USA"},
			},
		},
	}

	result, err := Get(data, "user.address[0].coordinates[2]")
	if result != nil {
		t.Errorf("expected nil result for non-existent nested field in nested array and nested array, got %v", result)
	}
	if _, ok := err.(DoesNotExist); !ok {
		t.Errorf("expected DoesNotExist error for non-existent nested field in nested array and nested array, got %v", err)
	}
}

func TestGet_NestedMapNestedArrayNestedArrayEmptyStringPath(t *testing.T) {
	data := map[string]interface{}{
		"user": map[string]interface{}{
			"firstname": "seth",
			"lastname":  "rogen",
			"address": []map[string]string{
				{"city": "New York", "country": "USA"},
				{"city": "Los Angeles", "country": "USA"},
			},
		},
	}

	result, err := Get(data, "user.address[0].coordinates.")
	if result != nil {
		t.Errorf("expected nil result for empty string path in nested map and nested array and nested array, got %v", result)
	}
	if err != nil {
		t.Errorf("expected no error for empty string path in nested map and nested array and nested array, got %v", err)
	}
}

func TestGet_NestedMapNestedArrayNestedArrayEmptyArrayPath(t *testing.T) {
	data := map[string]interface{}{
		"user": map[string]interface{}{
			"firstname": "seth",
			"lastname":  "rogen",
			"address": []map[string]string{
				{"city": "New York", "country": "USA"},
				{"city": "Los Angeles", "country": "USA"},
			},
		},
	}

	result, err := Get(data, "user.address[].coordinates.")
	if result != nil {
		t.Errorf("expected nil result for empty array path in nested map and nested array and nested array, got %v", result)
	}
	if err != nil {
		t.Errorf("expected no error for empty array path in nested map and nested array and nested array, got %v", err)
	}
}

func TestGet_NestedMapNestedArrayNestedArrayArrayIndexOutOfBounds(t *testing.T) {
	data := map[string]interface{}{
		"user": map[string]interface{}{
			"firstname": "seth",
			"lastname":  "rogen",
			"address": []map[string]string{
				{"city": "New York", "country": "USA"},
				{"city": "Los Angeles", "country": "USA"},
			},
		},
	}

	result, err := Get(data, "user.address[0].coordinates[10]")
	if result != nil {
		t.Errorf("expected nil result for out-of-bounds array index in nested map and nested array and nested array, got %v", result)
	}
	if _, ok := err.(DoesNotExist); !ok {
		t.Errorf("expected DoesNotExist error for out-of-bounds array index in nested map and nested array and nested array, got %v", err)
	}
}

func TestGet_NestedMapNestedArrayNestedArrayNestedArray(t *testing.T) {
	data := map[string]interface{}{
		"user": map[string]interface{}{
			"firstname": "seth",
			"lastname":  "rogen",
			"address": []map[string]string{
				{"city": "New York", "country": "USA"},
				{"city": "Los Angeles", "country": "USA"},
			},
		},
	}

	result, err := Get(data, "user.address[0].coordinates[0][1]")
	if result != nil {
		t.Errorf("expected nil result for non-existent nested field in nested array and nested array and nested array, got %v", result)
	}
	if _, ok := err.(DoesNotExist); !ok {
		t.Errorf("expected DoesNotExist error for non-existent nested field in nested array and nested array and nested array, got %v", err)
	}
}

func TestGet_NestedMapNestedArrayNestedArrayNestedArrayNonExistentField(t *testing.T) {
	data := map[string]interface{}{
		"user": map[string]interface{}{
			"firstname": "seth",
			"lastname":  "rogen",
			"address": []map[string]string{
				{"city": "New York", "country": "USA"},
				{"city": "Los Angeles", "country": "USA"},
			},
		},
	}

	result, err := Get(data, "user.address[0].coordinates[0][2]")
	if result != nil {
		t.Errorf("expected nil result for non-existent nested field in nested array and nested array and nested array, got %v", result)
	}
	if _, ok := err.(DoesNotExist); !ok {
		t.Errorf("expected DoesNotExist error for non-existent nested field in nested array and nested array and nested array, got %v", err)
	}
}

func TestGet_NestedMapNestedArrayNestedArrayNestedArrayEmptyStringPath(t *testing.T) {
	data := map[string]interface{}{
		"user": map[string]interface{}{
			"firstname": "seth",
			"lastname":  "rogen",
			"address": []map[string]string{
				{"city": "New York", "country": "USA"},
				{"city": "Los Angeles", "country": "USA"},
			},
		},
	}

	result, err := Get(data, "user.address[0].coordinates[0].")
	if result != nil {
		t.Errorf("expected nil result for empty string path in nested map and nested array and nested array and nested array, got %v", result)
	}
	if err != nil {
		t.Errorf("expected no error for empty string path in nested map and nested array and nested array and nested array, got %v", err)
	}
}

func TestGet_NestedMapNestedArrayNestedArrayNestedArrayEmptyArrayPath(t *testing.T) {
	data := map[string]interface{}{
		"user": map[string]interface{}{
			"firstname": "seth",
			"lastname":  "rogen",
			"address": []map[string]string{
				{"city": "New York", "country": "USA"},
				{"city": "Los Angeles", "country": "USA"},
			},
		},
	}

	result, err := Get(data, "user.address[].coordinates[0].")
	if result != nil {
		t.Errorf("expected nil result for empty array path in nested map and nested array and nested array and nested array, got %v", result)
	}
	if err != nil {
		t.Errorf("expected no error for empty array path in nested map and nested array and nested array and nested array, got %v", err)
	}
}

func TestGet_NestedMapNestedArrayNestedArrayNestedArrayArrayIndexOutOfBounds(t *testing.T) {
	data := map[string]interface{}{
		"user": map[string]interface{}{
			"firstname": "seth",
			"lastname":  "rogen",
			"address": []map[string]string{
				{"city": "New York", "country": "USA"},
				{"city": "Los Angeles", "country": "USA"},
			},
		},
	}

	result, err := Get(data, "user.address[0].coordinates[0][10]")
	if result != nil {
		t.Errorf("expected nil result for out-of-bounds array index in nested map and nested array and nested array and nested array, got %v", result)
	}
	if _, ok := err.(DoesNotExist); !ok {
		t.Errorf("expected DoesNotExist error for out-of-bounds array index in nested map and nested array and nested array and nested array, got %v", err)
	}
}

func TestGet_NestedMapNestedArrayNestedArrayNestedArrayNestedArray(t *testing.T) {
	data := map[string]interface{}{
		"user": map[string]interface{}{
			"firstname": "seth",
			"lastname":  "rogen",
			"address": []map[string]string{
				{"city": "New York", "country": "USA"},
				{"city": "Los Angeles", "country": "USA"},
			},
		},
	}

	result, err := Get(data, "user.address[0].coordinates[0][1][2]")
	if result != nil {
		t.Errorf("expected nil result for non-existent nested field in nested array and nested array and nested array and nested array, got %v", result)
	}
	if _, ok := err.(DoesNotExist); !ok {
		t.Errorf("expected DoesNotExist error for non-existent nested field in nested array and nested array and nested array and nested array, got %v", err)
	}
}

func TestGet_NestedMapNestedArrayNestedArrayNestedArrayNestedArrayNonExistentField(t *testing.T) {
	data := map[string]interface{}{
		"user": map[string]interface{}{
			"firstname": "seth",
			"lastname":  "rogen",
			"address": []map[string]string{
				{"city": "New York", "country": "USA"},
				{"city": "Los Angeles", "country": "USA"},
			},
		},
	}

	result, err := Get(data, "user.address[0].coordinates[0][1][3]")
	if result != nil {
		t.Errorf("expected nil result for non-existent nested field in nested array and nested array and nested array and nested array, got %v", result)
	}
	if _, ok := err.(DoesNotExist); !ok {
		t.Errorf("expected DoesNotExist error for non-existent nested field in nested array and nested array and nested array and nested array, got %v", err)
	}
}

func TestGet_NestedMapNestedArrayNestedArrayNestedArrayNestedArrayEmptyStringPath(t *testing.T) {
	data := map[string]interface{}{
		"user": map[string]interface{}{
			"firstname": "seth",
			"lastname":  "rogen",
			"address": []map[string]string{
				{"city": "New York", "country": "USA"},
				{"city": "Los Angeles", "country": "USA"},
			},
		},
	}

	result, err := Get(data, "user.address[0].coordinates[0][1].")
	if result != nil {
		t.Errorf("expected nil result for empty string path in nested map and nested array and nested array and nested array and nested array, got %v", result)
	}
	if err != nil {
		t.Errorf("expected no error for empty string path in nested map and nested array and nested array and nested array and nested array, got %v", err)
	}
}

func TestGet_NestedMapNestedArrayNestedArrayNestedArrayNestedArrayEmptyArrayPath(t *testing.T) {
	data := map[string]interface{}{
		"user": map[string]interface{}{
			"firstname": "seth",
			"lastname":  "rogen",
			"address": []map[string]string{
				{"city": "New York", "country": "USA"},
				{"city": "Los Angeles", "country": "USA"},
			},
		},
	}

	result, err := Get(data, "user.address[].coordinates[0][1].")
	if result != nil {
		t.Errorf("expected nil result for empty array path in nested map and nested array and nested array and nested array and nested array, got %v", result)
	}
	if err != nil {
		t.Errorf("expected no error for empty array path in nested map and nested array and nested array and nested array and nested array, got %v", err)
	}
}

func TestGet_NestedMapNestedArrayNestedArrayNestedArrayNestedArrayArrayIndexOutOfBounds(t *testing.T) {
	data := map[string]interface{}{
		"user": map[string]interface{}{
			"firstname": "seth",
			"lastname":  "rogen",
			"address": []map[string]string{
				{"city": "New York", "country": "USA"},
				{"city": "Los Angeles", "country": "USA"},
			},
		},
	}

	result, err := Get(data, "user.address[0].coordinates[0][1][10]")
	if result != nil {
		t.Errorf("expected nil result for out-of-bounds array index in nested map and nested array and nested array and nested array and nested array, got %v", result)
	}
	if _, ok := err.(DoesNotExist); !ok {
		t.Errorf("expected DoesNotExist error for out-of-bounds array index in nested map and nested array and nested array and nested array and nested array, got %v", err)
	}
}

func TestGet_NestedMapNestedArrayNestedArrayNestedArrayNestedArrayNestedArray(t *testing.T) {
	data := map[string]interface{}{
		"user": map[string]interface{}{
			"firstname": "seth",
			"lastname":  "rogen",
			"address": []map[string]string{
				{"city": "New York", "country": "USA"},
				{"city": "Los Angeles", "country": "USA"},
			},
		},
	}

	result, err := Get(data, "user.address[0].coordinates[0][1][2][3]")
	if result != nil {
		t.Errorf("expected nil result for non-existent nested field in nested array and nested array and nested array and nested array and nested array, got %v", result)
	}
	if _, ok := err.(DoesNotExist); !ok {
		t.Errorf("expected DoesNotExist error for non-existent nested field in nested array and nested array and nested array and nested array and nested array, got %v", err)
	}
}

func TestGet_NestedMapNestedArrayNestedArrayNestedArrayNestedArrayNestedArrayNonExistentField(t *testing.T) {
	data := map[string]interface{}{
		"user": map[string]interface{}{
			"firstname": "seth",
			"lastname":  "rogen",
			"address": []map[string]string{
				{"city": "New York", "country": "USA"},
				{"city": "Los Angeles", "country": "USA"},
			},
		},
	}

	result, err := Get(data, "user.address[0].coordinates[0][1][2][4]")
	if result != nil {
		t.Errorf("expected nil result for non-existent nested field in nested array and nested array and nested array and nested array and nested array, got %v", result)
	}
	if _, ok := err.(DoesNotExist); !ok {
		t.Errorf("expected DoesNotExist error for non-existent nested field in nested array and nested array and nested array and nested array and nested array, got %v", err)
	}
}

func TestGet_NestedMapNestedArrayNestedArrayNestedArrayNestedArrayNestedArrayEmptyStringPath(t *testing.T) {
	data := map[string]interface{}{
		"user": map[string]interface{}{
			"firstname": "seth",
			"lastname":  "rogen",
			"address": []map[string]string{
				{"city": "New York", "country": "USA"},
				{"city": "Los Angeles", "country": "USA"},
			},
		},
	}

	result, err := Get(data, "user.address[0].coordinates[0][1][2][3]")
	if result != nil {
		t.Errorf("expected nil result for non-existent nested field in nested array and nested array and nested array and nested array and nested array, got %v", result)
	}
	if _, ok := err.(DoesNotExist); !ok {
		t.Errorf("expected DoesNotExist error for non-existent nested field in nested array and nested array and nested array and nested array and nested array, got %v", err)
	}
}

func TestGet_NestedMapNestedArrayNestedArrayNestedArrayNestedArrayNestedArrayEmptyArrayPath(t *testing.T) {
	data := map[string]interface{}{
		"user": map[string]interface{}{
			"firstname": "seth",
			"lastname":
