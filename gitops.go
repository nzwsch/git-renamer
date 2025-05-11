package main

import (
	"errors"
	"os/exec"
	"strings"
)

type CommandExecutor interface {
	Output(name string, args ...string) ([]byte, error)
}

// RealExecutor: Executes commands in reality
type RealExecutor struct{}

func (r RealExecutor) Output(name string, args ...string) ([]byte, error) {
	cmd := exec.Command(name, args...)
	return cmd.Output()
}

// getFirstCommitDate: Allows injection of an executor
func getFirstCommitDate(executor CommandExecutor, dir string) (string, error) {
	output, err := executor.Output("git", "-C", dir, "log", "--reverse", "--max-parents=0", "HEAD", "--format=%ci")
	if err != nil {
		return "", err
	}

	result := strings.TrimSpace(string(output))
	if result == "" {
		return "", errors.New("no commits found")
	}
	return result, nil
}
