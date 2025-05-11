package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
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
		if err := os.Chdir(dir); err != nil {
			fmt.Println("Failed to change directory:", err)
			return
		}

		cmd := exec.Command("git", "log", "--reverse", "--max-parents=0", "HEAD", "--format=%ci")
		output, err := cmd.Output()
		if err != nil {
			fmt.Println("Error executing git command:", err)
			return
		}
		scanner := bufio.NewScanner(strings.NewReader(string(output)))
		if scanner.Scan() {
			dateStr := scanner.Text()
			dateStr = strings.TrimSpace(dateStr)
			fmt.Println("First commit date:", dateStr)
		} else {
			fmt.Println("No commit found")
		}
	}
}
