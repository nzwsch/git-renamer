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

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Failed to get home directory:", err)
		return
	}

	dirs, err := listAllPaths(homeDir)
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
