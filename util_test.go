package util

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExeDirectory(t *testing.T) {
	t.Run("Happy Path", func(t *testing.T) {
		exePath := os.Args[0]
		expected := filepath.Dir(exePath)
		result := ExeDirectory()
		assert.Equal(t, expected, result)
	})

	t.Run("Negative Path - Non-Existent File", func(t *testing.T) {
		exePath := "/non-existent-file"
		defer os.Remove(exePath)
		os.Create(exePath)
		expected := filepath.Dir(exePath)
		result := ExeDirectory()
		assert.Equal(t, expected, result)
	})

	t.Run("Negative Path - Invalid File Path", func(t *testing.T) {
		exePath := "invalid/path"
		defer os.Remove(exePath)
		os.Create(exePath)
		expected := filepath.Dir(exePath)
		result := ExeDirectory()
		assert.Equal(t, expected, result)
	})

	t.Run("Negative Path - Empty File Path", func(t *testing.T) {
		exePath := ""
		defer os.Remove(exePath)
		os.Create(exePath)
		expected := filepath.Dir(exePath)
		result := ExeDirectory()
		assert.Equal(t, expected, result)
	})
}

func TestFileStat(t *testing.T) {
	t.Run("Happy Path - File Exists", func(t *testing.T) {
		filePath := "testfile.txt"
		defer os.Remove(filePath)
		os.Create(filePath)
		result := FileStat(filePath, true)
		assert.NotNil(t, result)
	})

	t.Run("Negative Path - File Does Not Exist", func(t *testing.T) {
		filePath := "/non-existent-file"
		result := FileStat(filePath, false)
		assert.Nil(t, result)
	})

	t.Run("Negative Path - Invalid File Path", func(t *testing.T) {
		filePath := "invalid/path"
		result := FileStat(filePath, false)
		assert.Nil(t, result)
	})
}

func TestFileExists(t *testing.T) {
	t.Run("Happy Path - File Exists", func(t *testing.T) {
		filePath := "testfile.txt"
		defer os.Remove(filePath)
		os.Create(filePath)
		result := FileExists(filePath)
		assert.True(t, result)
	})

	t.Run("Negative Path - File Does Not Exist", func(t *testing.T) {
		filePath := "/non-existent-file"
		result := FileExists(filePath)
		assert.False(t, result)
	})

	t.Run("Negative Path - Invalid File Path", func(t *testing.T) {
		filePath := "invalid/path"
		result := FileExists(filePath)
		assert.False(t, result)
	})
}

func TestDirExists(t *testing.T) {
	t.Run("Happy Path - Directory Exists", func(t *testing.T) {
		dirPath := "testdir"
		defer os.RemoveAll(dirPath)
		os.MkdirAll(dirPath, os.ModePerm)
		result := DirExists(dirPath)
		assert.True(t, result)
	})

	t.Run("Negative Path - Directory Does Not Exist", func(t *testing.T) {
		dirPath := "/non-existent-dir"
		result := DirExists(dirPath)
		assert.False(t, result)
	})

	t.Run("Negative Path - Invalid Directory Path", func(t *testing.T) {
		dirPath := "invalid/path"
		result := DirExists(dirPath)
		assert.False(t, result)
	})
}

func TestRemoveFile(t *testing.T) {
	t.Run("Happy Path - File Exists", func(t *testing.T) {
		filePath := "testfile.txt"
		defer os.Remove(filePath)
		os.Create(filePath)
		err := RemoveFile(filePath)
		assert.NoError(t, err)
	})

	t.Run("Negative Path - File Does Not Exist", func(t *testing.T) {
		filePath := "/non-existent-file"
		err := RemoveFile(filePath)
		assert.Error(t, err)
	})

	t.Run("Negative Path - Invalid File Path", func(t *testing.T) {
		filePath := "invalid/path"
		err := RemoveFile(filePath)
		assert.Error(t, err)
	})
}

func TestRemoveDir(t *testing.T) {
	t.Run("Happy Path - Directory Exists", func(t *testing.T) {
		dirPath := "testdir"
		defer os.RemoveAll(dirPath)
		os.MkdirAll(dirPath, os.ModePerm)
		err := RemoveDir(dirPath)
		assert.NoError(t, err)
	})

	t.Run("Negative Path - Directory Does Not Exist", func(t *testing.T) {
		dirPath := "/non-existent-dir"
		err := RemoveDir(dirPath)
		assert.Error(t, err)
	})

	t.Run("Negative Path - Invalid Directory Path", func(t *testing.T) {
		dirPath := "invalid/path"
		err := RemoveDir(dirPath)
		assert.Error(t, err)
	})
}

func TestMkdirAll(t *testing.T) {
	t.Run("Happy Path - Directory Does Not Exist", func(t *testing.T) {
		dirPath := "testdir"
		defer os.RemoveAll(dirPath)
		err := MkdirAll(dirPath)
		assert.NoError(t, err)
	})

	t.Run("Negative Path - Directory Already Exists", func(t *testing.T) {
		dirPath := "testdir"
		defer os.RemoveAll(dirPath)
		os.MkdirAll(dirPath, os.ModePerm)
		err := MkdirAll(dirPath)
		assert.NoError(t, err)
	})

	t.Run("Negative Path - Invalid Directory Path", func(t *testing.T) {
		dirPath := "invalid/path"
		err := MkdirAll(dirPath)
		assert.Error(t, err)
	})
}

func TestMustStringSlice(t *testing.T) {
	t.Run("Happy Path - String Slice", func(t *testing.T) {
		raw := []string{"a", "b", "c"}
		result, err := MustStringSlice(raw)
		assert.NoError(t, err)
		assert.Equal(t, raw, result)
	})

	t.Run("Negative Path - Non-String Slice", func(t *testing.T) {
		raw := []interface{}{1, 2, 3}
		_, err := MustStringSlice(raw)
		assert.Error(t, err)
	})
}

func TestExtractStringSliceFromMap(t *testing.T) {
	t.Run("Happy Path - String Slice", func(t *testing.T) {
		m := map[string]interface{}{
			"key": []string{"a", "b", "c"},
		}
		result, err := ExtractStringSliceFromMap(m, "key")
		assert.NoError(t, err)
		assert.Equal(t, []string{"a", "b", "c"}, result)
	})

	t.Run("Negative Path - Non-String Slice", func(t *testing.T) {
		m := map[string]interface{}{
			"key": []interface{}{1, 2, 3},
		}
		_, err := ExtractStringSliceFromMap(m, "key")
		assert.Error(t, err)
	})

	t.Run("Negative Path - Key Does Not Exist", func(t *testing.T) {
		m := map[string]interface{}{
			"key": []string{"a", "b", "c"},
		}
		result, err := ExtractStringSliceFromMap(m, "non-existent-key")
		assert.NoError(t, err)
		assert.Equal(t, []string{}, result)
	})
}

func TestExtractStringFromMap(t *testing.T) {
	t.Run("Happy Path - String Value", func(t *testing.T) {
		m := map[string]interface{}{
			"key": "value",
		}
		result := ExtractStringFromMap(m, "key")
		assert.Equal(t, "value", result)
	})

	t.Run("Negative Path - Key Does Not Exist", func(t *testing.T) {
		m := map[string]interface{}{
			"key": "value",
		}
		result := ExtractStringFromMap(m, "non-existent-key")
		assert.Equal(t, "", result)
	})
}

func TestExtractIntFromMap(t *testing.T) {
	t.Run("Happy Path - Int Value", func(t *testing.T) {
		m := map[string]interface{}{
			"key": 123,
		}
		result := ExtractIntFromMap(m, "key")
		assert.Equal(t, 123, result)
	})

	t.Run("Negative Path - Key Does Not Exist", func(t *testing.T) {
		m := map[string]interface{}{
			"key": 123,
		}
		result := ExtractIntFromMap(m, "non-existent-key")
		assert.Equal(t, 0, result)
	})
}

func TestExtractBoolFromMap(t *testing.T) {
	t.Run("Happy Path - Bool Value", func(t *testing.T) {
		m := map[string]interface{}{
			"key": true,
		}
		result := ExtractBoolFromMap(m, "key")
		assert.Equal(t, true, result)
	})

	t.Run("Negative Path - Key Does Not Exist", func(t *testing.T) {
		m := map[string]interface{}{
			"key": true,
		}
		result := ExtractBoolFromMap(m, "non-existent-key")
		assert.Equal(t, false, result)
	})
}

func TestExtractFloat64FromMap(t *testing.T) {
	t.Run("Happy Path - Float64 Value", func(t *testing.T) {
		m := map[string]interface{}{
			"key": 123.45,
		}
		result := ExtractFloat64FromMap(m, "key")
		assert.Equal(t, 123.45, result)
	})

	t.Run("Negative Path - Key Does Not Exist", func(t *testing.T) {
		m := map[string]interface{}{
			"key": 123.45,
		}
		result := ExtractFloat64FromMap(m, "non-existent-key")
		assert.Equal(t, 0.0, result)
	})
}

func TestExtractTimeFromMap(t *testing.T) {
	t.Run("Happy Path - Time Value", func(t *testing.T) {
		m := map[string]interface{}{
			"key": time.Now(),
		}
		result := ExtractTimeFromMap(m, "key")
		assert.Equal(t, time.Now(), result)
	})

	t.Run("Negative Path - Key Does Not Exist", func(t *testing.T) {
		m := map[string]interface{}{
			"key": time.Now(),
		}
		result := ExtractTimeFromMap(m, "non-existent-key")
		assert.Equal(t, time.Time{}, result)
	})
}

func TestExtractDurationFromMap(t *testing.T) {
	t.Run("Happy Path - Duration Value", func(t *testing.T) {
		m := map[string]interface{}{
			"key": 10 * time.Second,
		}
		result := ExtractDurationFromMap(m, "key")
		assert.Equal(t, 10*time.Second, result)
	})

	t.Run("Negative Path - Key Does Not Exist", func(t *testing.T) {
		m := map[string]interface{}{
			"key": 10 * time.Second,
		}
		result := ExtractDurationFromMap(m, "non-existent-key")
		assert.Equal(t, 0*time.Second, result)
	})
}

func TestExtractInterfaceFromMap(t *testing.T) {
	t.Run("Happy Path - Interface Value", func(t *testing.T) {
		m := map[string]interface{}{
			"key": "value",
		}
		result := ExtractInterfaceFromMap(m, "key")
		assert.Equal(t, "value", result)
	})

	t.Run("Negative Path - Key Does Not Exist", func(t *testing.T) {
		m := map[string]interface{}{
			"key": "value",
		}
		result := ExtractInterfaceFromMap(m, "non-existent-key")
		assert.Nil(t, result)
	})
}

func TestExtractSliceFromMap(t *testing.T) {
	t.Run("Happy Path - Slice Value", func(t *testing.T) {
		m := map[string]interface{}{
			"key": []string{"a", "b", "c"},
		}
		result := ExtractSliceFromMap(m, "key")
		assert.Equal(t, []string{"a", "b", "c"}, result)
	})

	t.Run("Negative Path - Key Does Not Exist", func(t *testing.T) {
		m := map[string]interface{}{
			"key": []string{"a", "b", "c"},
		}
		result := ExtractSliceFromMap(m, "non-existent-key")
		assert.Nil(t, result)
	})
}

func TestExtractMapFromMap(t *testing.T) {
	t.Run("Happy Path - Map Value", func(t *testing.T) {
		m := map[string]interface{}{
			"key": map[string]string{"a": "b", "c": "d"},
		}
		result := ExtractMapFromMap(m, "key")
		assert.Equal(t, map[string]string{"a": "b", "c": "d"}, result)
	})

	t.Run("Negative Path - Key Does Not Exist", func(t *testing.T) {
		m := map[string]interface{}{
			"key": map[string]string{"a": "b", "c": "d"},
		}
		result := ExtractMapFromMap(m, "non-existent-key")
		assert.Nil(t, result)
	})
}

func TestExtractStringSliceFromMap(t *testing.T) {
	t.Run("Happy Path - String Slice Value", func(t *testing.T) {
		m := map[string]interface{}{
			"key": []string{"a", "b", "c"},
		}
		result := ExtractStringSliceFromMap(m, "key")
		assert.Equal(t, []string{"a", "b", "c"}, result)
	})

	t.Run("Negative Path - Key Does Not Exist", func(t *testing.T) {
		m := map[string]interface{}{
			"key": []string{"a", "b", "c"},
		}
		result := ExtractStringSliceFromMap(m, "non-existent-key")
		assert.Nil(t, result)
	})
}

func TestExtractIntSliceFromMap(t *testing.T) {
	t.Run("Happy Path - Int Slice Value", func(t *testing.T) {
		m := map[string]interface{}{
			"key": []int{1, 2, 3},
		}
		result := ExtractIntSliceFromMap(m, "key")
		assert.Equal(t, []int{1, 2, 3}, result)
	})

	t.Run("Negative Path - Key Does Not Exist", func(t *testing.T) {
		m := map[string]interface{}{
			"key": []int{1, 2, 3},
		}
		result := ExtractIntSliceFromMap(m, "non-existent-key")
		assert.Nil(t, result)
	})
}

func TestExtractBoolSliceFromMap(t *testing.T) {
	t.Run("Happy Path - Bool Slice Value", func(t *testing.T) {
		m := map[string]interface{}{
			"key": []bool{true, false, true},
		}
		result := ExtractBoolSliceFromMap(m, "key")
		assert.Equal(t, []bool{true, false, true}, result)
	})

	t.Run("Negative Path - Key Does Not Exist", func(t *testing.T) {
		m := map[string]interface{}{
			"key": []bool{true, false, true},
		}
		result := ExtractBoolSliceFromMap(m, "non-existent-key")
		assert.Nil(t, result)
	})
}

func TestExtractFloat64SliceFromMap(t *testing.T) {
	t.Run("Happy Path - Float64 Slice Value", func(t *testing.T) {
		m := map[string]interface{}{
			"key": []float64{1.0, 2.0, 3.0},
		}
		result := ExtractFloat64SliceFromMap(m, "key")
		assert.Equal(t, []float64{1.0, 2.0, 3.0}, result)
	})

	t.Run("Negative Path - Key Does Not Exist", func(t *testing.T) {
		m := map[string]interface{}{
			"key": []float64{1.0, 2.0, 3.0},
		}
		result := ExtractFloat64SliceFromMap(m, "non-existent-key")
		assert.Nil(t, result)
	})
}

func TestExtractDurationSliceFromMap(t *testing.T) {
	t.Run("Happy Path - Duration Slice Value", func(t *testing.T) {
		m := map[string]interface{}{
			"key": []time.Duration{10 * time.Second, 20 * time.Second},
		}
		result := ExtractDurationSliceFromMap(m, "key")
		assert.Equal(t, []time.Duration{10*time.Second, 20*time.Second}, result)
	})

	t.Run("Negative Path - Key Does Not Exist", func(t *testing.T) {
		m := map[string]interface{}{
			"key": []time.Duration{10 * time.Second, 20 * time.Second},
		}
		result := ExtractDurationSliceFromMap(m, "non-existent-key")
		assert.Nil(t, result)
	})
}

func TestExtractTimeSliceFromMap(t *testing.T) {
	t.Run("Happy Path - Time Slice Value", func(t *testing.T) {
		m := map[string]interface{}{
			"key": []time.Time{time.Now(), time.Now().Add(1 * time.Hour)},
		}
		result := ExtractTimeSliceFromMap(m, "key")
		assert.Equal(t, []time.Time{time.Now(), time.Now().Add(1 * time.Hour)}, result)
	})

	t.Run("Negative Path - Key Does Not Exist", func(t *testing.T) {
		m := map[string]interface{}{
			"key": []time.Time{time.Now(), time.Now().Add(1 * time.Hour)},
		}
		result := ExtractTimeSliceFromMap(m, "non-existent-key")
		assert.Nil(t, result)
	})
}

func TestExtractStringMapFromMap(t *testing.T) {
	t.Run("Happy Path - String Map Value", func(t *testing.T) {
		m := map[string]interface{}{
			"key": map[string]string{"a": "b", "c": "d"},
		}
		result := ExtractStringMapFromMap(m, "key")
		assert.Equal(t, map[string]string{"a": "b", "c": "d"}, result)
	})

	t.Run("Negative Path - Key Does Not Exist", func(t *testing.T) {
		m := map[string]interface{}{
			"key": map[string]string{"a": "b", "c": "d"},
		}
		result := ExtractStringMapFromMap(m, "non-existent-key")
		assert.Nil(t, result)
	})
}

func TestExtractIntMapFromMap(t *testing.T) {
	t.Run("Happy Path - Int Map Value", func(t *testing.T) {
		m := map[string]interface{}{
			"key": map[string]int{"a": 1, "b": 2},
		}
		result := ExtractIntMapFromMap(m, "key")
		assert.Equal(t, map[string]int{"a": 1, "b": 2}, result)
	})

	t.Run("Negative Path - Key Does Not Exist", func(t *testing.T) {
		m := map[string]interface{}{
			"key": map[string]int{"a": 1, "b": 2},
		}
		result := ExtractIntMapFromMap(m, "non-existent-key")
		assert.Nil(t, result)
	})
}

func TestExtractBoolMapFromMap(t *testing.T) {
	t.Run("Happy Path - Bool Map Value", func(t *testing.T) {
		m := map[string]interface{}{
			"key": map[string]bool{"a": true, "b": false},
		}
		result := ExtractBoolMapFromMap(m, "key")
		assert.Equal(t, map[string]bool{"a": true, "b": false}, result)
	})

	t.Run("Negative Path - Key Does Not Exist", func(t *testing.T) {
		m := map[string]interface{}{
			"key": map[string]bool{"a": true, "b": false},
		}
		result := ExtractBoolMapFromMap(m, "non-existent-key")
		assert.Nil(t, result)
	})
}

func TestExtractFloat64MapFromMap(t *testing.T) {
	t.Run("Happy Path - Float64 Map Value", func(t *testing.T) {
		m := map[string]interface{}{
			"key": map[string]float64{"a": 1.0, "b": 2.0},
		}
		result := ExtractFloat64MapFromMap(m, "key")
		assert.Equal(t, map[string]float64{"a": 1.0, "b": 2.0}, result)
	})

	t.Run("Negative Path - Key Does Not Exist", func(t *testing.T) {
		m := map[string]interface{}{
			"key": map[string]float64{"a": 1.0, "b": 2.0},
		}
		result := ExtractFloat64MapFromMap(m, "non-existent-key")
		assert.Nil(t, result)
	})
}

func TestExtractDurationMapFromMap(t *testing.T) {
	t.Run("Happy Path - Duration Map Value", func(t *testing.T) {
		m := map[string]interface{}{
			"key": map[string]time.Duration{"a": 10 * time.Second, "b": 20 * time.Second},
		}
		result := ExtractDurationMapFromMap(m, "key")
		assert.Equal(t, map[string]time.Duration{"a": 10*time.Second, "b": 20*time.Second}, result)
	})

	t.Run("Negative Path - Key Does Not Exist", func(t *testing.T) {
		m := map[string]interface{}{
			"key": map[string]time.Duration{"a": 10 * time.Second, "b": 20 * time.Second},
		}
		result := ExtractDurationMapFromMap(m, "non-existent-key")
		assert.Nil(t, result)
	})
}

func TestExtractTimeMapFromMap(t *testing.T) {
	t.Run("Happy Path - Time Map Value", func(t *testing.T) {
		m := map[string]interface{}{
			"key": map[string]time.Time{"a": time.Now(), "b": time.Now().Add(1 * time.Hour)},
		}
		result := ExtractTimeMapFromMap(m, "key")
		assert.Equal(t, map[string]time.Time{"a": time.Now(), "b": time.Now().Add(1 * time.Hour)}, result)
	})

	t.Run("Negative Path - Key Does Not Exist", func(t *testing.T) {
		m := map[string]interface{}{
			"key": map[string]time.Time{"a": time.Now(), "b": time.Now().Add(1 * time.Hour)},
		}
		result := ExtractTimeMapFromMap(m, "non-existent-key")
		assert.Nil(t, result)
	})
}

func TestExtractStringArrayFromMap(t *testing.T) {
	t.Run("Happy Path - String Array Value", func(t *testing.T) {
		m := map[string]interface{}{
			"key": []string{"a", "b", "c"},
		}
		result := ExtractStringArrayFromMap(m, "key")
		assert.Equal(t, []string{"a", "b", "c"}, result)
	})

	t.Run("Negative Path - Key Does Not Exist", func(t *testing.T) {
		m := map[string]interface{}{
			"key": []string{"a", "b", "c"},
		}
		result := ExtractStringArrayFromMap(m, "non-existent-key")
		assert.Nil(t, result)
	})
}

func TestExtractIntArrayFromMap(t *testing.T) {
	t.Run("Happy Path - Int Array Value", func(t *testing.T) {
		m := map[string]interface{}{
			"key": []int{1, 2, 3},
		}
		result := ExtractIntArrayFromMap(m, "key")
		assert.Equal(t, []int{1, 2, 3}, result)
	})

	t.Run("Negative Path - Key Does Not Exist", func(t *testing.T) {
		m := map[string]interface{}{
			"key": []int{1, 2, 3},
		}
		result := ExtractIntArrayFromMap(m, "non-existent-key")
		assert.Nil(t, result)
	})
}

func TestExtractBoolArrayFromMap(t *testing.T) {
	t.Run("Happy Path - Bool Array Value", func(t *testing.T) {
		m := map[string]interface{}{
			"key": []bool{true, false, true},
		}
		result := ExtractBoolArrayFromMap(m, "key")
		assert.Equal(t, []bool{true, false, true}, result)
	})

	t.Run("Negative Path - Key Does Not Exist", func(t *testing.T) {
		m := map[string]interface{}{
			"key": []bool{true, false, true},
		}
		result := ExtractBoolArrayFromMap(m, "non-existent-key")
		assert.Nil(t, result)
	})
}

func TestExtractFloat64ArrayFromMap(t *testing.T) {
	t.Run("Happy Path - Float64 Array Value", func(t *testing.T) {
		m := map[string]interface{}{
			"key": []float64{1.0, 2.0, 3.0},
		}
		result := ExtractFloat64ArrayFromMap(m, "key")
		assert.Equal(t, []float64{1.0, 2.0, 3.0}, result)
	})

	t.Run("Negative Path - Key Does Not Exist", func(t *testing.T) {
		m := map[string]interface{}{
			"key": []float64{1.0, 2.0, 3.0},
		}
		result := ExtractFloat64ArrayFromMap(m, "non-existent-key")
		assert.Nil(t, result)
	})
}

func TestExtractDurationArrayFromMap(t *testing.T) {
	t.Run("Happy Path - Duration Array Value", func(t *testing.T) {
		m := map[string]interface{}{
			"key": []time.Duration{10 * time.Second, 20 * time.Second},
		}
		result := ExtractDurationArrayFromMap(m, "key")
		assert.Equal(t, []time.Duration{10*time.Second, 20*time.Second}, result)
	})

	t.Run("Negative Path - Key Does Not Exist", func(t *testing.T) {
		m := map[string]interface{}{
			"key": []time.Duration{10 * time.Second, 20 * time.Second},
		}
		result := ExtractDurationArrayFromMap(m, "non-existent-key")
		assert.Nil(t, result)
	})
}

func TestExtractTimeArrayFromMap(t *testing.T) {
	t.Run("Happy Path - Time Array Value", func(t *testing.T) {
		m := map[string]interface{}{
			"key": []time.Time{time.Now(), time.Now().Add(1 * time.Hour)},
		}
		result := ExtractTimeArrayFromMap(m, "key")
		assert.Equal(t, []time.Time{time.Now(), time.Now().Add(1 * time.Hour)}, result)
	})

	t.Run("Negative Path - Key Does Not Exist", func(t *testing.T) {
		m := map[string]interface{}{
			"key": []time.Time{time.Now(), time.Now().Add(1 * time.Hour)},
		}
		result := ExtractTimeArrayFromMap(m, "non-existent-key")
		assert.Nil(t, result)
	})
}

func TestExtractStringMapFromMap(t *testing.T) {
	t.Run("Happy Path - String Map Value", func(t *testing.T) {
		m := map[string]interface{}{
			"key": map[string]string{"a": "1", "b": "2"},
		}
		result := ExtractStringMapFromMap(m, "key")
		assert.Equal(t, map[string]string{"a": "1", "b": "2"}, result)
	})

	t.Run("Negative Path - Key Does Not Exist", func(t *testing.T) {
		m := map[string]interface{}{
			"key": map[string]string{"a": "1", "b": "2"},
		}
		result := ExtractStringMapFromMap(m, "non-existent-key")
		assert.Nil(t, result)
	})
}

func TestExtractIntMapFromMap(t *testing.T) {
	t.Run("Happy Path - Int Map Value", func(t *testing.T) {
		m := map[string]interface{}{
			"key": map[string]int{"a": 1, "b": 2},
		}
		result := ExtractIntMapFromMap(m, "key")
		assert.Equal(t, map[string]int{"a": 1, "b": 2}, result)
	})

	t.Run("Negative Path - Key Does Not Exist", func(t *testing.T) {
		m := map[string]interface{}{
			"key": map[string]int{"a": 1, "b": 2},
		}
		result := ExtractIntMapFromMap(m, "non-existent-key")
		assert.Nil(t, result)
	})
}

func TestExtractBoolMapFromMap(t *testing.T) {
	t.Run("Happy Path - Bool Map Value", func(t *testing.T) {
		m := map[string]interface{}{
			"key": map[string]bool{"a": true, "b": false},
		}
		result := ExtractBoolMapFromMap(m, "key")
		assert.Equal(t, map[string]bool{"a": true, "b": false}, result)
	})

	t.Run("Negative Path - Key Does Not Exist", func(t *testing.T) {
		m := map[string]interface{}{
			"key": map[string]bool{"a": true, "b": false},
		}
		result := ExtractBoolMapFromMap(m, "non-existent-key")
		assert.Nil(t, result)
	})
}

func TestExtractFloat64MapFromMap(t *testing.T) {
	t.Run("Happy Path - Float64 Map Value", func(t *testing.T) {
		m := map[string]interface{}{
			"key": map[string]float64{"a": 1.0, "b": 2.0},
