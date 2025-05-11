package main

import (
	"bufio"
	"fmt"
	"os/exec"
	"strings"
)

func getFirstCommitDate(dir string) (string, error) {
	cmd := exec.Command("git", "log", "--reverse", "--max-parents=0", "HEAD", "--format=%ci")
	cmd.Dir = dir
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	scanner := bufio.NewScanner(strings.NewReader(string(output)))
	if scanner.Scan() {
		dateStr := scanner.Text()
		dateStr = strings.TrimSpace(dateStr)
		return dateStr, nil
	}
	return "", fmt.Errorf("no commit found")
}
