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

	for _, dir := range dirs {
		fmt.Println(dir)
	}
}
