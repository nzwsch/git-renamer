package main

import (
	"fmt"
	"os"
)

func main() {
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
		fmt.Println("Git directory:", dir)
		firstCommitDate, err := getFirstCommitDate(RealExecutor{}, dir)
		if err != nil {
			fmt.Println("Error getting first commit date:", err)
			return
		}
		fmt.Println("First commit date:", firstCommitDate)
	}
}
