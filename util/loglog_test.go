package util

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLogFileT_Write(t *testing.T) {
	t.Run("Happy Path", func(t *testing.T) {
		tmpDir := t.TempDir()
		logFile := LogFileT{path: filepath.Join(tmpDir, "test.log")}
		logFile.Open()

		data := []byte("Test data")
		n, err := logFile.Write(data)
		assert.NoError(t, err)
		assert.Equal(t, len(data), n)

		logFile.Close()
	})

	t.Run("Negative Path - File Not Found", func(t *testing.T) {
		tmpDir := t.TempDir()
		logFile := LogFileT{path: filepath.Join(tmpDir, "nonexistent.log")}
		err := logFile.Write([]byte("Test data"))
		assert.Error(t, err)
	})

	t.Run("Corner Case - Large File", func(t *testing.T) {
		tmpDir := t.TempDir()
		logFile := LogFileT{path: filepath.Join(tmpDir, "largefile.log")}
		logFile.Open()

		for i := 0; i < 15*1024*1024; i++ {
			data := []byte("a")
			n, err := logFile.Write(data)
			assert.NoError(t, err)
			assert.Equal(t, len(data), n)
		}

		logFile.Close()
	})

	t.Run("Negative Path - File Too Large", func(t *testing.T) {
		tmpDir := t.TempDir()
		logFile := LogFileT{path: filepath.Join(tmpDir, "largefile.log")}
		logFile.Open()

		for i := 0; i < 120*1024*1024; i++ {
			data := []byte("a")
			n, err := logFile.Write(data)
			assert.NoError(t, err)
			assert.Equal(t, len(data), n)
		}

		logFile.Close()
	})

	t.Run("Happy Path - Existing File", func(t *testing.T) {
		tmpDir := t.TempDir()
		logFile := LogFileT{path: filepath.Join(tmpDir, "existing.log")}
		logFile.Open()

		data := []byte("Existing data")
		n, err := logFile.Write(data)
		assert.NoError(t, err)
		assert.Equal(t, len(data), n)

		logFile.Close()
	})

	t.Run("Negative Path - Existing File Too Large", func(t *testing.T) {
		tmpDir := t.TempDir()
		logFile := LogFileT{path: filepath.Join(tmpDir, "existing.log")}
		logFile.Open()

		for i := 0; i < 120*1024*1024; i++ {
			data := []byte("a")
			n, err := logFile.Write(data)
			assert.NoError(t, err)
			assert.Equal(t, len(data), n)
		}

		logFile.Close()
	})

	t.Run("Happy Path - InitLogger", func(t *testing.T) {
		tmpDir := t.TempDir()
		logger := InitLogger(tmpDir)

		data := []byte("Test data")
		n, err := logger.Write(data)
		assert.NoError(t, err)
		assert.Equal(t, len(data), n)

		logger.Close()
	})

	t.Run("Negative Path - InitLogger with Nonexistent Directory", func(t *testing.T) {
		tmpDir := filepath.Join(os.TempDir(), "nonexistent")
		err := os.MkdirAll(tmpDir, 0755)
		assert.NoError(t, err)
		defer os.RemoveAll(tmpDir)

		logger := InitLogger(tmpDir)
		data := []byte("Test data")
		n, err := logger.Write(data)
		assert.Error(t, err)
		assert.Equal(t, 0, n)
	})

	t.Run("Happy Path - Close", func(t *testing.T) {
		tmpDir := t.TempDir()
		logFile := LogFileT{path: filepath.Join(tmpDir, "close.log")}
		logFile.Open()

		err := logFile.Close()
		assert.NoError(t, err)

		if _, err := os.Stat(logFile.path); !os.IsNotExist(err) {
			t.Errorf("Log file should be closed")
		}
	})

	t.Run("Negative Path - Close with Unopened File", func(t *testing.T) {
		logFile := LogFileT{path: filepath.Join(os.TempDir(), "unopened.log")}
		err := logFile.Close()
		assert.NoError(t, err)
	})
}

func TestLogFileT_Open(t *testing.T) {
	t.Run("Happy Path", func(t *testing.T) {
		tmpDir := t.TempDir()
		logFile := LogFileT{path: filepath.Join(tmpDir, "test.log")}
		logFile.Open()

		if _, err := os.Stat(logFile.path); os.IsNotExist(err) {
			t.Errorf("Log file should be created")
		}

		logFile.Close()
	})

	t.Run("Negative Path - File Not Found", func(t *testing.T) {
		tmpDir := t.TempDir()
		logFile := LogFileT{path: filepath.Join(tmpDir, "nonexistent.log")}
		err := logFile.Open()
		assert.Error(t, err)
	})

	t.Run("Corner Case - Large File", func(t *testing.T) {
		tmpDir := t.TempDir()
		logFile := LogFileT{path: filepath.Join(tmpDir, "largefile.log")}
		logFile.Open()

		for i := 0; i < 15*1024*1024; i++ {
			data := []byte("a")
			n, err := logFile.Write(data)
			assert.NoError(t, err)
			assert.Equal(t, len(data), n)
		}

		logFile.Close()
	})

	t.Run("Negative Path - File Too Large", func(t *testing.T) {
		tmpDir := t.TempDir()
		logFile := LogFileT{path: filepath.Join(tmpDir, "largefile.log")}
		logFile.Open()

		for i := 0; i < 120*1024*1024; i++ {
			data := []byte("a")
			n, err := logFile.Write(data)
			assert.NoError(t, err)
			assert.Equal(t, len(data), n)
		}

		logFile.Close()
	})

	t.Run("Happy Path - Existing File", func(t *testing.T) {
		tmpDir := t.TempDir()
		logFile := LogFileT{path: filepath.Join(tmpDir, "existing.log")}
		logFile.Open()

		data := []byte("Existing data")
		n, err := logFile.Write(data)
		assert.NoError(t, err)
		assert.Equal(t, len(data), n)

		logFile.Close()
	})

	t.Run("Negative Path - Existing File Too Large", func(t *testing.T) {
		tmpDir := t.TempDir()
		logFile := LogFileT{path: filepath.Join(tmpDir, "existing.log")}
		logFile.Open()

		for i := 0; i < 120*1024*1024; i++ {
			data := []byte("a")
			n, err := logFile.Write(data)
			assert.NoError(t, err)
			assert.Equal(t, len(data), n)
		}

		logFile.Close()
	})

	t.Run("Happy Path - InitLogger", func(t *testing.T) {
		tmpDir := t.TempDir()
		logger := InitLogger(tmpDir)

		data := []byte("Test data")
		n, err := logger.Write(data)
		assert.NoError(t, err)
		assert.Equal(t, len(data), n)

		logger.Close()
	})

	t.Run("Negative Path - InitLogger with Nonexistent Directory", func(t *testing.T) {
		tmpDir := filepath.Join(os.TempDir(), "nonexistent")
		err := os.MkdirAll(tmpDir, 0755)
		assert.NoError(t, err)
		defer os.RemoveAll(tmpDir)

		logger := InitLogger(tmpDir)
		data := []byte("Test data")
		n, err := logger.Write(data)
		assert.Error(t, err)
		assert.Equal(t, 0, n)
	})

	t.Run("Happy Path - Close", func(t *testing.T) {
		tmpDir := t.TempDir()
		logFile := LogFileT{path: filepath.Join(tmpDir, "close.log")}
		logFile.Open()

		err := logFile.Close()
		assert.NoError(t, err)

		if _, err := os.Stat(logFile.path); !os.IsNotExist(err) {
			t.Errorf("Log file should be closed")
		}
	})

	t.Run("Negative Path - Close with Unopened File", func(t *testing.T) {
		logFile := LogFileT{path: filepath.Join(os.TempDir(), "unopened.log")}
		err := logFile.Close()
		assert.NoError(t, err)
	})
}

func TestLogFileT_Close(t *testing.T) {
	t.Run("Happy Path", func(t *testing.T) {
		tmpDir := t.TempDir()
		logFile := LogFileT{path: filepath.Join(tmpDir, "close.log")}
		logFile.Open()

		err := logFile.Close()
		assert.NoError(t, err)

		if _, err := os.Stat(logFile.path); !os.IsNotExist(err) {
			t.Errorf("Log file should be closed")
		}
	})

	t.Run("Negative Path - Close with Unopened File", func(t *testing.T) {
		logFile := LogFileT{path: filepath.Join(os.TempDir(), "unopened.log")}
		err := logFile.Close()
		assert.NoError(t, err)
	})
}

func TestLogFileT_InitLogger(t *testing.T) {
	t.Run("Happy Path", func(t *testing.T) {
		tmpDir := t.TempDir()
		logger := InitLogger(tmpDir)

		data := []byte("Test data")
		n, err := logger.Write(data)
		assert.NoError(t, err)
		assert.Equal(t, len(data), n)

		logger.Close()
	})

	t.Run("Negative Path - InitLogger with Nonexistent Directory", func(t *testing.T) {
		tmpDir := filepath.Join(os.TempDir(), "nonexistent")
		err := os.MkdirAll(tmpDir, 0755)
		assert.NoError(t, err)
		defer os.RemoveAll(tmpDir)

		logger := InitLogger(tmpDir)
		data := []byte("Test data")
		n, err := logger.Write(data)
		assert.Error(t, err)
		assert.Equal(t, 0, n)
	})
}
