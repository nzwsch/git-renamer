package main

import (
	"os"
	"os/exec"
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
		"dir2/file.txt",
		"dir4/",
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

	// Initialize actual git repos in dir1 and dir3
	for _, dir := range []string{"dir1", "dir3"} {
		dirPath := filepath.Join(tempDir, dir)
		
		// Ensure directory exists
		if err := os.MkdirAll(dirPath, 0755); err != nil {
			t.Fatalf("Failed to create directory %s: %v", dirPath, err)
		}
		
		cmd := exec.Command("git", "init")
		cmd.Dir = dirPath
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("Failed to initialize git repo in %s: %v, output: %s", dirPath, err, output)
		}
		
		// Create an initial commit so HEAD exists
		cmd = exec.Command("git", "config", "user.email", "test@example.com")
		cmd.Dir = dirPath
		cmd.Run()
		cmd = exec.Command("git", "config", "user.name", "Test User")
		cmd.Dir = dirPath
		cmd.Run()
		
		// Create a file and commit it
		testFile := filepath.Join(dirPath, "test.txt")
		os.WriteFile(testFile, []byte("test"), 0644)
		cmd = exec.Command("git", "add", "test.txt")
		cmd.Dir = dirPath
		cmd.Run()
		cmd = exec.Command("git", "commit", "-m", "Initial commit")
		cmd.Dir = dirPath
		if err := cmd.Run(); err != nil {
			t.Fatalf("Failed to create initial commit in %s: %v", dirPath, err)
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
