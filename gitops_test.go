package main

import (
	"bytes"
	"errors"
	"os/exec"
	"testing"
)

func TestGetFirstCommitDate(t *testing.T) {
	tests := []struct {
		name        string
		mockOutput  string
		mockError   error
		expected    string
		expectError bool
	}{
		{
			name:        "Valid commit date",
			mockOutput:  "2023-01-01 12:00:00 +0000\n",
			mockError:   nil,
			expected:    "2023-01-01 12:00:00 +0000",
			expectError: false,
		},
		{
			name:        "No commits",
			mockOutput:  "",
			mockError:   nil,
			expected:    "",
			expectError: true,
		},
		{
			name:        "Git command error",
			mockOutput:  "",
			mockError:   errors.New("git error"),
			expected:    "",
			expectError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Mock exec.Command
			execCommand = func(name string, arg ...string) *exec.Cmd {
				return &exec.Cmd{
					Path:   name,
					Args:   append([]string{name}, arg...),
					Stdout: bytes.NewBufferString(test.mockOutput),
					Stderr: bytes.NewBuffer(nil),
				}
			}

			// Mock the Run method to simulate error
			if test.mockError != nil {
				execCommand = func(name string, arg ...string) *exec.Cmd {
					return &exec.Cmd{
						Path: name,
						Args: append([]string{name}, arg...),
						Run: func() error {
							return test.mockError
						},
					}
				}
			}

			// Call getFirstCommitDate
			result, err := getFirstCommitDate("mockDir")

			// Check the result
			if test.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if result != test.expected {
					t.Errorf("Expected %q, got %q", test.expected, result)
				}
			}
		})
	}
}
