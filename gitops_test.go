package main

import (
	"errors"
	"testing"
)

// MockExecutor : コマンド実行結果を差し替える
type MockExecutor struct {
	OutputFunc func(name string, args ...string) ([]byte, error)
}

func (m MockExecutor) Output(name string, args ...string) ([]byte, error) {
	return m.OutputFunc(name, args...)
}

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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := MockExecutor{
				OutputFunc: func(name string, args ...string) ([]byte, error) {
					return []byte(tt.mockOutput), tt.mockError
				},
			}

			result, err := getFirstCommitDate(mock, "/dummy/path")

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error, got nil")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if result != tt.expected {
					t.Errorf("Expected %q, got %q", tt.expected, result)
				}
			}
		})
	}
}
