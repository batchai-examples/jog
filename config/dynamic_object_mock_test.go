package config

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestMockDynamicObject_FromMap_HappyPath(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockObj := NewMockDynamicObject(ctrl)
	mockRecorder := mockObj.EXPECT()

	testMap := map[string]interface{}{
		"key1": "value1",
		"key2": 42,
	}

	mockRecorder.FromMap(testMap).Return(nil)

	err := mockObj.FromMap(testMap)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
}

func TestMockDynamicObject_FromMap_NegativePath(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockObj := NewMockDynamicObject(ctrl)
	mockRecorder := mockObj.EXPECT()

	testMap := map[string]interface{}{
		"key1": "value1",
		"key2": 42,
	}

	expectedErr := errors.New("test error")
	mockRecorder.FromMap(testMap).Return(expectedErr)

	err := mockObj.FromMap(testMap)
	if err == nil {
		t.Errorf("Expected an error, but got none")
	} else if err != expectedErr {
		t.Errorf("Expected error %v, but got %v", expectedErr, err)
	}
}

func TestMockDynamicObject_Init_HappyPath(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockObj := NewMockDynamicObject(ctrl)
	mockRecorder := mockObj.EXPECT()

	testConfig := Configuration{}

	mockRecorder.Init(testConfig)

	mockObj.Init(testConfig)
}

func TestMockDynamicObject_Reset_HappyPath(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockObj := NewMockDynamicObject(ctrl)
	mockRecorder := mockObj.EXPECT()

	mockRecorder.Reset()

	mockObj.Reset()
}

func TestMockDynamicObject_ToMap_HappyPath(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockObj := NewMockDynamicObject(ctrl)
	mockRecorder := mockObj.EXPECT()

	expectedMap := map[string]interface{}{
		"key1": "value1",
		"key2": 42,
	}

	mockRecorder.ToMap().Return(expectedMap)

	result := mockObj.ToMap()
	if !reflect.DeepEqual(result, expectedMap) {
		t.Errorf("Expected map %v, but got %v", expectedMap, result)
	}
}

func TestMockDynamicObject_ToMap_NegativePath(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockObj := NewMockDynamicObject(ctrl)
	mockRecorder := mockObj.EXPECT()

	expectedMap := map[string]interface{}{
		"key1": "value1",
		"key2": 42,
	}

	mockRecorder.ToMap().Return(nil)

	result := mockObj.ToMap()
	if result != nil {
		t.Errorf("Expected nil, but got %v", result)
	}
}

func TestMockDynamicObject_FromMap_CornerCase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockObj := NewMockDynamicObject(ctrl)
	mockRecorder := mockObj.EXPECT()

	testMap := map[string]interface{}{
		"key1": "value1",
		"key2": 42,
	}

	expectedErr := errors.New("test error")
	mockRecorder.FromMap(testMap).Return(expectedErr)

	err := mockObj.FromMap(testMap)
	if err == nil {
		t.Errorf("Expected an error, but got none")
	} else if err != expectedErr {
		t.Errorf("Expected error %v, but got %v", expectedErr, err)
	}
}

func TestMockDynamicObject_Init_CornerCase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockObj := NewMockDynamicObject(ctrl)
	mockRecorder := mockObj.EXPECT()

	testConfig := Configuration{}

	mockRecorder.Init(testConfig)

	mockObj.Init(testConfig)
}

func TestMockDynamicObject_Reset_CornerCase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockObj := NewMockDynamicObject(ctrl)
	mockRecorder := mockObj.EXPECT()

	mockRecorder.Reset()

	mockObj.Reset()
}
