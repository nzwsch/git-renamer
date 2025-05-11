package main

import (
	"io/fs"
	"os"
	"path/filepath"
)

func listAllPaths(root string) ([]string, error) {
	var paths []string

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		paths = append(paths, path)
		return nil
	})

	return paths, err
}

func onlyGitDirs(paths []string) []string {
	var dirs []string
	for _, path := range paths {
		info, err := os.Stat(path)
		if err != nil {
			continue
		}
		if !info.IsDir() {
			continue
		}
		if _, err := os.Stat(filepath.Join(path, ".git")); err == nil {
			dirs = append(dirs, path)
		}
	}
	return dirs
}
