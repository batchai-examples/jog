package util

import (
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/agiledragon/gomonkey"
	"github.com/stretchr/testify/assert"
)

func Test_ExtractStringSliceFromMap_Error(t *testing.T) {
	assert := assert.New(t)
	m := map[string]interface{}{
		"k": "not slice",
	}

	v, err := ExtractStringSliceFromMap(m, "k")
	assert.Error(err)
	assert.Nil(v)
	_, has := m["k"]
	assert.True(has, "should be still there")
}
