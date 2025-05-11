package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

// ANSI color codes
func colorGreen(text string) string {
	return fmt.Sprintf("\033[32m%s\033[0m", text)
}

func colorYellow(text string) string {
	return fmt.Sprintf("\033[33m%s\033[0m", text)
}

func main() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Failed to get home directory:", err)
		return
	}

	codesDirsPattern := filepath.Join(homeDir, "code", "*")
	thisDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println("Failed to get current directory:", err)
		return
	}

	matches, err := filepath.Glob(codesDirsPattern)
	if err != nil {
		fmt.Println("Glob failed:", err)
		return
	}

	for _, dir := range matches {
		if dir == thisDir {
			continue
		}

		if err := os.Chdir(dir); err != nil {
			fmt.Println(colorYellow(filepath.Base(dir)))
			continue
		}

		cmd := exec.Command("git", "log", "--reverse", "--max-parents=0", "HEAD", "--format=%ci")
		output, err := cmd.Output()
		if err != nil {
			fmt.Println(colorYellow(filepath.Base(dir)))
			continue
		}

		scanner := bufio.NewScanner(strings.NewReader(string(output)))
		if scanner.Scan() {
			dateStr := scanner.Text()
			t, err := time.Parse("2006-01-02 15:04:05 -0700", dateStr)
			if err != nil {
				fmt.Println(colorYellow(filepath.Base(dir)))
				continue
			}
			formatted := t.Format("060102")
			projectName := fmt.Sprintf("%s-%s", filepath.Base(dir), formatted)
			fmt.Println(colorGreen(projectName))
		} else {
			fmt.Println(colorYellow(filepath.Base(dir)))
		}
	}
}
