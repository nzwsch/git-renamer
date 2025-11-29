package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

const (
	colorGreen = "\033[32m"
	colorReset = "\033[0m"
)

func main() {
	noColor := flag.Bool("no-color", false, "Disable colored output")
	flag.Parse()

	var targetDir string
	args := flag.Args()

	if len(args) == 0 {
		// No argument provided, use home directory
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Failed to get home directory:", err)
			return
		}
		targetDir = homeDir
	} else {
		// Use the provided directory argument
		targetDir = args[0]
	}

	// Convert to absolute path
	absPath, err := filepath.Abs(targetDir)
	if err != nil {
		fmt.Println("Failed to get absolute path:", err)
		return
	}
	targetDir = absPath

	dirs, err := listAllPaths(targetDir)
	if err != nil {
		fmt.Println("Error listing paths:", err)
		return
	}

	dirs = onlyGitDirs(dirs)
	for _, dir := range dirs {
		firstCommitDate, err := getFirstCommitDate(RealExecutor{}, dir)
		if err != nil {
			fmt.Println("Error getting first commit date:", err)
			return
		}

		appended, err := appendProjectToDate(dir, firstCommitDate)
		if err != nil {
			fmt.Println("Error appending project to date:", err)
			return
		}

		if *noColor {
			fmt.Println(filepath.Dir(dir)+":", appended)
		} else {
			fmt.Println(filepath.Dir(dir)+":", colorGreen+appended+colorReset)
		}
	}
}
