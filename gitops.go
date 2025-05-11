package main

import (
	"bufio"
	"errors"
	"os/exec"
	"strings"
)

// 実行可能なコマンドを生成する関数型
var execCommand = exec.Command

func getFirstCommitDate(dir string) (string, error) {
	cmd := execCommand("git", "-C", dir, "log", "--reverse", "--max-parents=0", "HEAD", "--format=%ci")
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
	return "", errors.New("no commits found")
}
