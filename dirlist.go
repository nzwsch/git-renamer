package main

import (
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func getCurrentTargetDir() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return filepath.Abs(cwd)
}

func getTargetDir(args []string) (string, error) {
	if len(args) == 0 {
		return getCurrentTargetDir()
	}
	return filepath.Abs(args[0])
}

func listAllPaths(root string) ([]string, error) {
	var paths []string
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if isHiddenPath(path) {
			return nil
		}
		paths = append(paths, path)
		return nil
	})
	return paths, err
}

func isHiddenPath(path string) bool {
	cleaned := filepath.Clean(path)
	parts := strings.Split(cleaned, string(filepath.Separator))
	for _, part := range parts {
		if strings.HasPrefix(part, ".") && part != "." && part != ".." {
			return true
		}
	}
	return false
}

func onlyGitDirs(paths []string) []string {
	var dirs []string
	for _, path := range paths {
		info, err := os.Stat(path)
		if err != nil || !info.IsDir() {
			continue
		}

		if _, err := os.Stat(filepath.Join(path, ".git")); err != nil {
			continue
		}

		cmd := exec.Command("git", "rev-parse", "HEAD")
		cmd.Dir = path
		if err := cmd.Run(); err == nil {
			dirs = append(dirs, path)
		}
	}
	return dirs
}
