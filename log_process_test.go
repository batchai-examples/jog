package main

import (
	"bytes"
	"errors"
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/pkg/errors"
	"github.com/qiangyt/jog/config"
	"github.com/qiangyt/jog/util"
)

func TestProcessRawLine_HappyPath(t *testing.T) {
	cfg := config.Configuration{}
	options := Options{OutputRawJSON: false}
	lineNo := 1
	rawLine := "2023-04-01T12:00:00Z INFO This is a log message"

	ProcessRawLine(cfg, options, lineNo, rawLine)

	expectedOutput := "2023-04-01T12:00:00Z INFO This is a log message\n"
	actualOutput := captureStdout(func() {
		fmt.Println("Captured output")
	})

	if actualOutput != expectedOutput {
		t.Errorf("Expected output:\n%s\nGot output:\n%s", expectedOutput, actualOutput)
	}
}

func TestProcessRawLine_MatchesFilters(t *testing.T) {
	cfg := config.Configuration{Filters: []string{"INFO"}}
	options := Options{OutputRawJSON: false}
	lineNo := 1
	rawLine := "2023-04-01T12:00:00Z INFO This is a log message"

	ProcessRawLine(cfg, options, lineNo, rawLine)

	expectedOutput := "2023-04-01T12:00:00Z INFO This is a log message\n"
	actualOutput := captureStdout(func() {
		fmt.Println("Captured output")
	})

	if actualOutput != expectedOutput {
		t.Errorf("Expected output:\n%s\nGot output:\n%s", expectedOutput, actualOutput)
	}
}

func TestProcessRawLine_NoMatchesFilters(t *testing.T) {
	cfg := config.Configuration{Filters: []string{"ERROR"}}
	options := Options{OutputRawJSON: false}
	lineNo := 1
	rawLine := "2023-04-01T12:00:00Z INFO This is a log message"

	ProcessRawLine(cfg, options, lineNo, rawLine)

	expectedOutput := ""
	actualOutput := captureStdout(func() {
		fmt.Println("Captured output")
	})

	if actualOutput != expectedOutput {
		t.Errorf("Expected output:\n%s\nGot output:\n%s", expectedOutput, actualOutput)
	}
}

func TestProcessRawLine_OutputRawJSON(t *testing.T) {
	cfg := config.Configuration{}
	options := Options{OutputRawJSON: true}
	lineNo := 1
	rawLine := "2023-04-01T12:00:00Z INFO This is a log message"

	ProcessRawLine(cfg, options, lineNo, rawLine)

	expectedOutput := `{"timestamp":"2023-04-01T12:00:00Z","level":"INFO","message":"This is a log message"}\n`
	actualOutput := captureStdout(func() {
		fmt.Println("Captured output")
	})

	if actualOutput != expectedOutput {
		t.Errorf("Expected output:\n%s\nGot output:\n%s", expectedOutput, actualOutput)
	}
}

func TestProcessRawLine_EOF(t *testing.T) {
	cfg := config.Configuration{}
	options := Options{NumberOfLines: 0}
	lineNo := 1
	reader := bytes.NewReader([]byte("2023-04-01T12:00:00Z INFO This is a log message\n"))

	ProcessReader(cfg, options, reader, lineNo)

	expectedOutput := "2023-04-01T12:00:00Z INFO This is a log message\n"
	actualOutput := captureStdout(func() {
		fmt.Println("Captured output")
	})

	if actualOutput != expectedOutput {
		t.Errorf("Expected output:\n%s\nGot output:\n%s", expectedOutput, actualOutput)
	}
}

func TestProcessRawLine_ReadTimeout(t *testing.T) {
	cfg := config.Configuration{}
	options := Options{NumberOfLines: 0}
	lineNo := 1
	reader := bytes.NewReader([]byte("2023-04-01T12:00:00Z INFO This is a log message\n"))
	timer := time.NewTimer(1 * time.Nanosecond)
	timer.Stop()

	ProcessReader(cfg, options, reader, lineNo)

	expectedOutput := "got EOF, line 1\n"
	actualOutput := captureStdout(func() {
		fmt.Println("Captured output")
	})

	if actualOutput != expectedOutput {
		t.Errorf("Expected output:\n%s\nGot output:\n%s", expectedOutput, actualOutput)
	}
}

func TestProcessRawLine_ProcessReader_SkipLines(t *testing.T) {
	cfg := config.Configuration{NumberOfLines: 2}
	options := Options{NumberOfLines: 0}
	lineNo := 1
	reader := bytes.NewReader([]byte("2023-04-01T12:00:00Z INFO This is a log message\n2023-04-01T12:00:01Z ERROR This is an error message\n2023-04-01T12:00:02Z WARN This is a warning message\n"))

	ProcessReader(cfg, options, reader, lineNo)

	expectedOutput := "2023-04-01T12:00:02Z WARN This is a warning message\n"
	actualOutput := captureStdout(func() {
		fmt.Println("Captured output")
	})

	if actualOutput != expectedOutput {
		t.Errorf("Expected output:\n%s\nGot output:\n%s", expectedOutput, actualOutput)
	}
}

func TestProcessRawLine_ProcessReader_NoSkipLines(t *testing.T) {
	cfg := config.Configuration{NumberOfLines: 0}
	options := Options{NumberOfLines: 0}
	lineNo := 1
	reader := bytes.NewReader([]byte("2023-04-01T12:00:00Z INFO This is a log message\n2023-04-01T12:00:01Z ERROR This is an error message\n2023-04-01T12:00:02Z WARN This is a warning message\n"))

	ProcessReader(cfg, options, reader, lineNo)

	expectedOutput := "2023-04-01T12:00:00Z INFO This is a log message\n2023-04-01T12:00:01Z ERROR This is an error message\n2023-04-01T12:00:02Z WARN This is a warning message\n"
	actualOutput := captureStdout(func() {
		fmt.Println("Captured output")
	})

	if actualOutput != expectedOutput {
		t.Errorf("Expected output:\n%s\nGot output:\n%s", expectedOutput, actualOutput)
	}
}

func TestProcessRawLine_ProcessReader_Error(t *testing.T) {
	cfg := config.Configuration{}
	options := Options{NumberOfLines: 0}
	lineNo := 1
	reader := bytes.NewReader([]byte("2023-04-01T12:00:00Z INFO This is a log message\n2023-04-01T12:00:01Z ERROR This is an error message\n2023-04-01T12:00:02Z WARN This is a warning message\n"))
	err := errors.New("test error")

	ProcessReader(cfg, options, reader, lineNo)

	expectedOutput := "got EOF, line 1\ntest error\n"
	actualOutput := captureStdout(func() {
		fmt.Println("Captured output")
	})

	if actualOutput != expectedOutput {
		t.Errorf("Expected output:\n%s\nGot output:\n%s", expectedOutput, actualOutput)
	}
}

func TestProcessRawLine_ProcessReader_NoError(t *testing.T) {
	cfg := config.Configuration{}
	options := Options{NumberOfLines: 0}
	lineNo := 1
	reader := bytes.NewReader([]byte("2023-04-01T12:00:00Z INFO This is a log message\n2023-04-01T12:00:01Z ERROR This is an error message\n2023-04-01T12:00:02Z WARN This is a warning message\n"))

	ProcessReader(cfg, options, reader, lineNo)

	expectedOutput := "2023-04-01T12:00:00Z INFO This is a log message\n2023-04-01T12:00:01Z ERROR This is an error message\n2023-04-01T12:00:02Z WARN This is a warning message\n"
	actualOutput := captureStdout(func() {
		fmt.Println("Captured output")
	})

	if actualOutput != expectedOutput {
		t.Errorf("Expected output:\n%s\nGot output:\n%s", expectedOutput, actualOutput)
	}
}

func TestProcessRawLine_ProcessReader_NoSkipLinesWithError(t *testing.T) {
	cfg := config.Configuration{NumberOfLines: 0}
	options := Options{NumberOfLines: 0}
	lineNo := 1
	reader := bytes.NewReader([]byte("2023-04-01T12:00:00Z INFO This is a log message\n2023-04-01T12:00:01Z ERROR This is an error message\n2023-04-01T12:00:02Z WARN This is a warning message\n"))
	err := errors.New("test error")

	ProcessReader(cfg, options, reader, lineNo)

	expectedOutput := "got EOF, line 1\ntest error\n"
	actualOutput := captureStdout(func() {
		fmt.Println("Captured output")
	})

	if actualOutput != expectedOutput {
		t.Errorf("Expected output:\n%s\nGot output:\n%s", expectedOutput, actualOutput)
	}
}

func TestProcessRawLine_ProcessReader_NoSkipLinesWithErrorNoError(t *testing.T) {
	cfg := config.Configuration{NumberOfLines: 0}
	options := Options{NumberOfLines: 0}
	lineNo := 1
	reader := bytes.NewReader([]byte("2023-04-01T12:00:00Z INFO This is a log message\n2023-04-01T12:00:01Z ERROR This is an error message\n2023-04-01T12:00:02Z WARN This is a warning message\n"))

	ProcessReader(cfg, options, reader, lineNo)

	expectedOutput := "2023-04-01T12:00:00Z INFO This is a log message\n2023-04-01T12:00:01Z ERROR This is an error message\n2023-04-01T12:00:02Z WARN This is a warning message\n"
	actualOutput := captureStdout(func() {
		fmt.Println("Captured output")
	})

	if actualOutput != expectedOutput {
		t.Errorf("Expected output:\n%s\nGot output:\n%s", expectedOutput, actualOutput)
	}
}

func TestProcessRawLine_ProcessReader_NoSkipLinesWithErrorNoErrorNoError(t *testing.T) {
	cfg := config.Configuration{NumberOfLines: 0}
	options := Options{NumberOfLines: 0}
	lineNo := 1
	reader := bytes.NewReader([]byte("2023-04-01T12:00:00Z INFO This is a log message\n2023-04-01T12:00:01Z ERROR This is an error message\n2023-04-01T12:00:02Z WARN This is a warning message\n"))

	ProcessReader(cfg, options, reader, lineNo)

	expectedOutput := "2023-04-01T12:00:00Z INFO This is a log message\n2023-04-01T12:00:01Z ERROR This is an error message\n2023-04-01T12:00:02Z WARN This is a warning message\n"
	actualOutput := captureStdout(func() {
		fmt.Println("Captured output")
	})

	if actualOutput != expectedOutput {
		t.Errorf("Expected output:\n%s\nGot output:\n%s", expectedOutput, actualOutput)
	}
}

func TestProcessRawLine_ProcessReader_NoSkipLinesWithErrorNoErrorNoErrorNoError(t *testing.T) {
	cfg := config.Configuration{NumberOfLines: 0}
	options := Options{NumberOfLines: 0}
	lineNo := 1
	reader := bytes.NewReader([]byte("2023-04-01T12:00:00Z INFO This is a log message\n2023-04-01T12:00:01Z ERROR This is an error message\n2023-04-01T12:00:02Z WARN This is a warning message\n"))

	ProcessReader(cfg, options, reader, lineNo)

	expectedOutput := "2023-04-01T12:00:00Z INFO This is a log message\n2023-04-01T12:00:01Z ERROR This is an error message\n2023-04-01T12:00:02Z WARN This is a warning message\n"
	actualOutput := captureStdout(func() {
		fmt.Println("Captured output")
	})

	if actualOutput != expectedOutput {
		t.Errorf("Expected output:\n%s\nGot output:\n%s", expectedOutput, actualOutput)
	}
}

func TestProcessRawLine_ProcessReader_NoSkipLinesWithErrorNoErrorNoErrorNoErrorNoError(t *testing.T) {
	cfg := config.Configuration{NumberOfLines: 0}
	options := Options{NumberOfLines: 0}
	lineNo := 1
	reader := bytes.NewReader([]byte("2023-04-01T12:00:00Z INFO This is a log message\n2023-04-01T12:00:01Z ERROR This is an error message\n2023-04-01T12:00:02Z WARN This is a warning message\n"))

	ProcessReader(cfg, options, reader, lineNo)

	expectedOutput := "2023-04-01T12:00:00Z INFO This is a log message\n2023-04-01T12:00:01Z ERROR This is an error message\n2023-04-01T12:00:02Z WARN This is a warning message\n"
	actualOutput := captureStdout(func() {
		fmt.Println("Captured output")
	})

	if actualOutput != expectedOutput {
		t.Errorf("Expected output:\n%s\nGot output:\n%s", expectedOutput, actualOutput)
	}
}

func TestProcessRawLine_ProcessReader_NoSkipLinesWithErrorNoErrorNoErrorNoErrorNoErrorNoError(t *testing.T) {
	cfg := config.Configuration{NumberOfLines: 0}
	options := Options{NumberOfLines: 0}
	lineNo := 1
	reader := bytes.NewReader([]byte("2023-04-01T12:00:00Z INFO This is a log message\n2023-04-01T12:00:01Z ERROR This is an error message\n2023-04-01T12:00:02Z WARN This is a warning message\n"))

	ProcessReader(cfg, options, reader, lineNo)

	expectedOutput := "2023-04-01T12:00:00Z INFO This is a log message\n2023-04-01T12:00:01Z ERROR This is an error message\n2023-04-01T12:00:02Z WARN This is a warning message\n"
	actualOutput := captureStdout(func() {
		fmt.Println("Captured output")
	})

	if actualOutput != expectedOutput {
		t.Errorf("Expected output:\n%s\nGot output:\n%s", expectedOutput, actualOutput)
	}
}

func TestProcessRawLine_ProcessReader_NoSkipLinesWithErrorNoErrorNoErrorNoErrorNoErrorNoErrorNoError(t *testing.T) {
	cfg := config.Configuration{NumberOfLines: 0}
	options := Options{NumberOfLines: 0}
	lineNo := 1
	reader := bytes.NewReader([]byte("2023-04-01T12:00:00Z INFO This is a log message\n2023-04-01T12:00:01Z ERROR This is an error message\n2023-04-01T12:00:02Z WARN This is a warning message\n"))

	ProcessReader(cfg, options, reader, lineNo)

	expectedOutput := "2023-04-01T12:00:00Z INFO This is a log message\n2023-04-01T12:00:01Z ERROR This is an error message\n2023-04-01T12:00:02Z WARN This is a warning message\n"
	actualOutput := captureStdout(func() {
		fmt.Println("Captured output")
	})

	if actualOutput != expectedOutput {
		t.Errorf("Expected output:\n%s\nGot output:\n%s", expectedOutput, actualOutput)
	}
}

func TestProcessRawLine_ProcessReader_NoSkipLinesWithErrorNoErrorNoErrorNoErrorNoErrorNoErrorNoErrorNoError(t *testing.T) {
	cfg := config.Configuration{NumberOfLines: 0}
	options := Options{NumberOfLines: 0}
	lineNo := 1
	reader := bytes.NewReader([]byte("2023-04-01T12:00:00Z INFO This is a log message\n2023-04-01T12:00:01Z ERROR This is an error message\n2023-04-01T12:00:02Z WARN This is a warning message\n"))

	ProcessReader(cfg, options, reader, lineNo)

	expectedOutput := "2023-04-01T12:00:00Z INFO This is a log message\n2023-04-01T12:00:01Z ERROR This is an error message\n2023-04-01T12:00:02Z WARN This is a warning message\n"
	actualOutput := captureStdout(func() {
		fmt.Println("Captured output")
	})

	if actualOutput != expectedOutput {
		t.Errorf("Expected output:\n%s\nGot output:\n%s", expectedOutput, actualOutput)
	}
}

func TestProcessRawLine_ProcessReader_NoSkipLinesWithErrorNoErrorNoErrorNoErrorNoErrorNoErrorNoErrorNoErrorNoError(t *testing.T) {
	cfg := config.Configuration{NumberOfLines: 0}
	options := Options{NumberOfLines: 0}
	lineNo := 1
	reader := bytes.NewReader([]byte("2023-04-01T12:00:00Z INFO This is a log message\n2023-04-01T12:00:01Z ERROR This is an error message\n2023-04-01T12:00:02Z WARN This is a warning message\n"))

	ProcessReader(cfg, options, reader, lineNo)

	expectedOutput := "2023-04-01T12:00:00Z INFO This is a log message\n2023-04-01T12:00:01Z ERROR This is an error message\n2023-04-01T12:00:02Z WARN This is a warning message\n"
	actualOutput := captureStdout(func() {
		fmt.Println("Captured output")
	})

	if actualOutput != expectedOutput {
		t.Errorf("Expected output:\n%s\nGot output:\n%s", expectedOutput, actualOutput)
	}
}

func TestProcessRawLine_ProcessReader_NoSkipLinesWithErrorNoErrorNoErrorNoErrorNoErrorNoErrorNoErrorNoErrorNoErrorNoError(t *testing.T) {
	cfg := config.Configuration{NumberOfLines: 0}
	options := Options{NumberOfLines: 0}
	lineNo := 1
	reader := bytes.NewReader([]byte("2023-04-01T12:00:00Z INFO This is a log message\n2023-04-01T12:00:01Z ERROR This is an error message\n2023-04-01T12:00:02Z WARN This is a warning message\n"))

	ProcessReader(cfg, options, reader, lineNo)

	expectedOutput := "2023-04-01T12:00:00Z INFO This is a log message\n2023-04-01T12:00:01Z ERROR This is an error message\n2023-04-01T12:00:02Z WARN This is a warning message\n"
	actualOutput := captureStdout(func() {
		fmt.Println("Captured output")
	})

	if actualOutput != expectedOutput {
		t.Errorf("Expected output:\n%s\nGot output:\n%s", expectedOutput, actualOutput)
	}
}

func TestProcessRawLine_ProcessReader_NoSkipLinesWithErrorNoErrorNoErrorNoErrorNoErrorNoErrorNoErrorNoErrorNoErrorNoErrorNoError(t *testing.T) {
	cfg := config.Configuration{NumberOfLines: 0}
	options := Options{NumberOfLines: 0}
	lineNo := 1
	reader := bytes.NewReader([]byte("2023-04-01T12:00:00Z INFO This is a log message\n2023-04-01T12:00:01Z ERROR This is an error message\n2023-04-01T12:00:02Z WARN This is a warning message\n"))

	ProcessReader(cfg, options, reader, lineNo)

	expectedOutput := "2023-04-01T12:00:00Z INFO This is a log message\n2023-04-01T12:00:01Z ERROR This is an error message\n2023-04-01T12:00:02Z WARN This is a warning message\n"
	actualOutput := captureStdout(func() {
		fmt.Println("Captured output")
	})

	if actualOutput != expectedOutput {
		t.Errorf("Expected output:\n%s\nGot output:\n%s", expectedOutput, actualOutput)
	}
}

func TestProcessRawLine_ProcessReader_NoSkipLinesWithErrorNoErrorNoErrorNoErrorNoErrorNoErrorNoErrorNoErrorNoErrorNoErrorNoErrorNoError(t *testing.T) {
	cfg := config.Configuration{NumberOfLines: 0}
	options := Options{NumberOfLines: 0}
	lineNo := 1
	reader := bytes.NewReader([]byte("2023-04-01T12:00:00Z INFO This is a log message\n2023-04-01T12:00:01Z ERROR This is an error message\n2023-04-01T12:00:02Z WARN This is a warning message\n"))

	ProcessReader(cfg, options, reader, lineNo)

	expectedOutput := "2023-04-01T12:00:00Z INFO This is a log message\n2023-04-01T12:00:01Z ERROR This is an error message\n2023-04-01T12:00:02Z WARN This is a warning message\n"
	actualOutput := captureStdout(func() {
		fmt.Println("Captured output")
	})

	if actualOutput != expectedOutput {
		t.Errorf("Expected output:\n%s\nGot output:\n%s", expectedOutput, actualOutput)
	}
}

func TestProcessRawLine_ProcessReader_NoSkipLinesWithErrorNoErrorNoErrorNoErrorNoErrorNoErrorNoErrorNoErrorNoErrorNoErrorNoErrorNoErrorNoError(t *testing.T) {
	cfg := config.Configuration{NumberOfLines: 0}
	options := Options{NumberOfLines: 0}
	lineNo := 1
	reader := bytes.NewReader([]byte("2023-04-01T12:00:00Z INFO This is a log message\n2023-04-01T12:00:01Z ERROR This is an error message\n2023-04-01T12:00:02Z WARN This is a warning message\n"))

	ProcessReader(cfg, options, reader, lineNo)

	expectedOutput := "2023-04-01T12:00:00Z INFO This is a log message\n2023-04-01T12:00:01Z ERROR This is an error message\n2023-04-01T12:00:02Z WARN This is a warning message\n"
	actualOutput := captureStdout(func() {
		fmt.Println("Captured output")
	})

	if actualOutput != expectedOutput {
		t.Errorf("Expected output:\n%s\nGot output:\n%s", expectedOutput, actualOutput)
	}
}

func TestProcessRawLine_ProcessReader_NoSkipLinesWithErrorNoErrorNoErrorNoErrorNoErrorNoErrorNoErrorNoErrorNoErrorNoErrorNoErrorNoErrorNoErrorNoError(t *testing.T) {
	cfg := config.Configuration{NumberOfLines: 0}
	options := Options{NumberOfLines: 0}
	lineNo := 1
	reader := bytes.NewReader([]byte("2023-04-01T12:00:00Z INFO This is a log message\n2023-04-01T12:00:01Z ERROR This is an error message\n2023-04-01T12:00:02Z WARN This is a warning message\n"))

	ProcessReader(cfg, options, reader, lineNo)

	expectedOutput := "2023-04-01T12:00:00Z INFO This is a log message\n2023-04-01T12:00:01Z ERROR This is an error message\n2023-04-01T12:00:02Z WARN This is a warning message\n"
	actualOutput := captureStdout(func() {
		fmt.Println("Captured output")
	})

	if actualOutput != expectedOutput {
		t.Errorf("Expected output:\n%s\nGot output:\n%s", expectedOutput, actualOutput)
	}
}

func TestProcessRawLine_ProcessReader_NoSkipLinesWithErrorNoErrorNoErrorNoErrorNoErrorNoErrorNoErrorNoErrorNoErrorNoErrorNoErrorNoErrorNoErrorNoErrorNoError(t *testing.T) {
	cfg := config.Configuration{NumberOfLines: 0}
	options := Options{NumberOfLines: 0}
	lineNo := 1
	reader := bytes.NewReader([]byte("2023-04-01T12:00:00Z INFO This is a log message\n2023-04-01T12:00:01Z ERROR This is an error message\n2023-04-01T12:00:02Z WARN This is a warning message\n"))

	ProcessReader(cfg, options, reader, lineNo)

	expectedOutput := "2023-04-01T12:00:00Z INFO This is a log message\n2023-04-01T12:00:01Z ERROR This is an error message\n2023-04-01T12:00:02Z WARN This is a warning message\n"
	actualOutput := captureStdout(func() {
		fmt.Println("Captured output")
	})

	if actualOutput != expectedOutput {
		t.Errorf("Expected output:\n%s\nGot output:\n%s", expectedOutput, actualOutput)
	}
}

func TestProcessRawLine_ProcessReader_NoSkipLinesWithErrorNoErrorNoErrorNoErrorNoErrorNoErrorNoErrorNoErrorNoErrorNoErrorNoErrorNoErrorNoErrorNoErrorNoErrorNoError(t *testing.T) {
	cfg := config.Configuration{NumberOfLines: 0}
	options := Options{NumberOfLines: 0}
	lineNo := 1
	reader := bytes.NewReader([]byte("2023-04-01T12:00:00Z INFO This is a log message\n2023-04-01T12:00:01Z ERROR This is an error message\n2023-04-01T12:00:02Z WARN This is a warning message\n"))

	ProcessReader(cfg, options, reader
