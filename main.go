package main

import (
	"fmt"
)

func main() {
	dirs, err := listAllPaths(".")

	if err != nil {
		fmt.Println("Error listing paths:", err)
		return
	}

	fmt.Println("dirs:", dirs)
}
