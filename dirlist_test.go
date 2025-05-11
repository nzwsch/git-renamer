package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestListAllPaths(t *testing.T) {
	// Create a temporary directory for testing
	tempDir := t.TempDir()

	// Create test files and directories
	testStructure := []string{
		"file1.txt",
		"dir1/file2.txt",
		"dir1/.hiddenfile",
		".hiddendir/file3.txt",
		"dir2/.git/config",
	}
	for _, path := range testStructure {
		fullPath := filepath.Join(tempDir, path)
		if strings.HasSuffix(path, "/") {
			os.MkdirAll(fullPath, 0755)
		} else {
			os.MkdirAll(filepath.Dir(fullPath), 0755)
			os.WriteFile(fullPath, []byte("test"), 0644)
		}
	}

	// Call listAllPaths
	paths, err := listAllPaths(tempDir)
	if err != nil {
		t.Fatalf("listAllPaths returned an error: %v", err)
	}

	// Expected paths (excluding hidden files and directories)
	expectedPaths := []string{
		tempDir,
		filepath.Join(tempDir, "file1.txt"),
		filepath.Join(tempDir, "dir1"),
		filepath.Join(tempDir, "dir1/file2.txt"),
		filepath.Join(tempDir, "dir2"),
	}

	// Check if the returned paths match the expected paths
	expectedSet := make(map[string]struct{})
	for _, p := range expectedPaths {
		expectedSet[p] = struct{}{}
	}

	for _, p := range paths {
		if _, exists := expectedSet[p]; !exists {
			t.Errorf("Unexpected path: %s", p)
		}
		delete(expectedSet, p)
	}

	for p := range expectedSet {
		t.Errorf("Missing expected path: %s", p)
	}
}

func TestIsHiddenPath(t *testing.T) {
	tests := []struct {
		path     string
		expected bool
	}{
		{path: "file.txt", expected: false},
		{path: ".hiddenfile", expected: true},
		{path: "dir/.hiddenfile", expected: true},
		{path: "dir/file.txt", expected: false},
		{path: ".hiddendir/file.txt", expected: true},
		{path: "dir/.hiddendir/file.txt", expected: true},
		{path: ".", expected: false},
		{path: "..", expected: false},
		{path: "./.hiddenfile", expected: true},
		{path: "../.hiddenfile", expected: true},
	}

	for _, test := range tests {
		result := isHiddenPath(test.path)
		if result != test.expected {
			t.Errorf("isHiddenPath(%q) = %v; want %v", test.path, result, test.expected)
		}
	}
}

func TestOnlyGitDirs(t *testing.T) {
	// Create a temporary directory for testing
	tempDir := t.TempDir()

	// Create test directories and files
	testStructure := []string{
		"dir1/.git/config",
		"dir2/file.txt",
		"dir3/.git/config",
		"dir4/",
	}
	for _, path := range testStructure {
		fullPath := filepath.Join(tempDir, path)
		if strings.HasSuffix(path, "/") || strings.Contains(path, ".git") {
			os.MkdirAll(fullPath, 0755)
		} else {
			os.MkdirAll(filepath.Dir(fullPath), 0755)
			os.WriteFile(fullPath, []byte("test"), 0644)
		}
	}

	// List all paths in the temporary directory
	allPaths, err := listAllPaths(tempDir)
	if err != nil {
		t.Fatalf("listAllPaths returned an error: %v", err)
	}

	// Call onlyGitDirs
	gitDirs := onlyGitDirs(allPaths)

	// Expected git directories
	expectedGitDirs := []string{
		filepath.Join(tempDir, "dir1"),
		filepath.Join(tempDir, "dir3"),
	}

	// Check if the returned git directories match the expected directories
	expectedSet := make(map[string]struct{})
	for _, dir := range expectedGitDirs {
		expectedSet[dir] = struct{}{}
	}

	for _, dir := range gitDirs {
		if _, exists := expectedSet[dir]; !exists {
			t.Errorf("Unexpected git directory: %s", dir)
		}
		delete(expectedSet, dir)
	}

	for dir := range expectedSet {
		t.Errorf("Missing expected git directory: %s", dir)
	}
}
