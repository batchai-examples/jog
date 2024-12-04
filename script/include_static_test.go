package main

import (
	"testing"
)

func TestIncludeDir(t *testing.T) {
	// Test case: Happy path - include a directory with files and subdirectories
	staticGoParentDir := "./static_go"
	staticFileParentDir := "./static_files"
	dirName := "test_dir"

	includeDir(staticGoParentDir, staticFileParentDir, dirName)

	// Check if the directories are created
	if _, err := os.Stat(filepath.Join(staticGoParentDir, dirName)); os.IsNotExist(err) {
		t.Errorf("Directory %s not created", filepath.Join(staticGoParentDir, dirName))
	}

	// Check if the files are included
	if _, err := os.Stat(filepath.Join(staticGoParentDir, dirName, "file1.go")); os.IsNotExist(err) {
		t.Errorf("File file1.go not included")
	}
	if _, err := os.Stat(filepath.Join(staticGoParentDir, dirName, "subdir", "file2.go")); os.IsNotExist(err) {
		t.Errorf("File file2.go not included in subdir")
	}

	// Clean up
	os.RemoveAll(staticGoParentDir)
}

func TestIncludeFile(t *testing.T) {
	// Test case: Happy path - include a single file
	staticGoParentDir := "./static_go"
	staticFileParentDir := "./static_files"
	fName := "file1.txt"

	includeFile(staticGoParentDir, staticFileParentDir, fName)

	// Check if the file is included
	if _, err := os.Stat(filepath.Join(staticGoParentDir, fName+".go")); os.IsNotExist(err) {
		t.Errorf("File %s.go not included", fName)
	}

	// Clean up
	os.RemoveAll(staticGoParentDir)
}

func TestIncludeEmptyDir(t *testing.T) {
	// Test case: Corner case - include an empty directory
	staticGoParentDir := "./static_go"
	staticFileParentDir := "./static_files"
	dirName := "empty_dir"

	includeDir(staticGoParentDir, staticFileParentDir, dirName)

	// Check if the directory is created
	if _, err := os.Stat(filepath.Join(staticGoParentDir, dirName)); os.IsNotExist(err) {
		t.Errorf("Directory %s not created", filepath.Join(staticGoParentDir, dirName))
	}

	// Clean up
	os.RemoveAll(staticGoParentDir)
}

func TestIncludeFileWithExt(t *testing.T) {
	// Test case: Happy path - include a file with extension
	staticGoParentDir := "./static_go"
	staticFileParentDir := "./static_files"
	fName := "file1.txt"

	includeFile(staticGoParentDir, staticFileParentDir, fName)

	// Check if the file is included with correct variable name
	if _, err := os.Stat(filepath.Join(staticGoParentDir, "File1_txt.go")); os.IsNotExist(err) {
		t.Errorf("File File1_txt.go not included")
	}

	// Clean up
	os.RemoveAll(staticGoParentDir)
}

func TestIncludeFileWithDash(t *testing.T) {
	// Test case: Happy path - include a file with dash in name
	staticGoParentDir := "./static_go"
	staticFileParentDir := "./static_files"
	fName := "file-1.txt"

	includeFile(staticGoParentDir, staticFileParentDir, fName)

	// Check if the file is included with correct variable name
	if _, err := os.Stat(filepath.Join(staticGoParentDir, "File_1_txt.go")); os.IsNotExist(err) {
		t.Errorf("File File_1_txt.go not included")
	}

	// Clean up
	os.RemoveAll(staticGoParentDir)
}

func TestIncludeFileWithSpecialChars(t *testing.T) {
	// Test case: Happy path - include a file with special characters in name
	staticGoParentDir := "./static_go"
	staticFileParentDir := "./static_files"
	fName := "file!@#$.txt"

	includeFile(staticGoParentDir, staticFileParentDir, fName)

	// Check if the file is included with correct variable name
	if _, err := os.Stat(filepath.Join(staticGoParentDir, "File_!_@___$.txt.go")); os.IsNotExist(err) {
		t.Errorf("File File_!_@___$.txt.go not included")
	}

	// Clean up
	os.RemoveAll(staticGoParentDir)
}

func TestIncludeFileWithQuotes(t *testing.T) {
	// Test case: Happy path - include a file with quotes in content
	staticGoParentDir := "./static_go"
	staticFileParentDir := "./static_files"
	fName := "file_with_quotes.txt"

	ioutil.WriteFile(filepath.Join(staticFileParentDir, fName), []byte(`"Hello, world!"`), os.ModePerm)
	includeFile(staticGoParentDir, staticFileParentDir, fName)

	// Check if the file is included with correct content
	contentBytes, _ := ioutil.ReadFile(filepath.Join(staticGoParentDir, fName+".go"))
	content := string(contentBytes)
	if !strings.Contains(content, "`Hello, world!`") {
		t.Errorf("Content of %s.go does not match", fName)
	}

	// Clean up
	os.RemoveAll(staticGoParentDir)
}

func TestIncludeFileWithBackticks(t *testing.T) {
	// Test case: Happy path - include a file with backticks in content
	staticGoParentDir := "./static_go"
	staticFileParentDir := "./static_files"
	fName := "file_with_backticks.txt"

	ioutil.WriteFile(filepath.Join(staticFileParentDir, fName), []byte("`Hello, world!`"), os.ModePerm)
	includeFile(staticGoParentDir, staticFileParentDir, fName)

	// Check if the file is included with correct content
	contentBytes, _ := ioutil.ReadFile(filepath.Join(staticGoParentDir, fName+".go"))
	content := string(contentBytes)
	if !strings.Contains(content, "`Hello, world!`") {
		t.Errorf("Content of %s.go does not match", fName)
	}

	// Clean up
	os.RemoveAll(staticGoParentDir)
}

func TestIncludeFileWithNewLines(t *testing.T) {
	// Test case: Happy path - include a file with new lines in content
	staticGoParentDir := "./static_go"
	staticFileParentDir := "./static_files"
	fName := "file_with_new_lines.txt"

	ioutil.WriteFile(filepath.Join(staticFileParentDir, fName), []byte("Hello,\nworld!"), os.ModePerm)
	includeFile(staticGoParentDir, staticFileParentDir, fName)

	// Check if the file is included with correct content
	contentBytes, _ := ioutil.ReadFile(filepath.Join(staticGoParentDir, fName+".go"))
	content := string(contentBytes)
	if !strings.Contains(content, "`Hello,\nworld!`") {
		t.Errorf("Content of %s.go does not match", fName)
	}

	// Clean up
	os.RemoveAll(staticGoParentDir)
}
