package main

import (
	"testing"
)

func TestConvertToDate(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    string
		expectError bool
	}{
		{
			name:        "Valid date string 1",
			input:       "2025-05-11 15:41:59 +0900",
			expected:    "20250511",
			expectError: false,
		},
		{
			name:        "Valid date string 2",
			input:       "2024-12-28 00:05:43 +0900",
			expected:    "20241228",
			expectError: false,
		},
		{
			name:        "Invalid date format",
			input:       "11-05-2025 15:41:59",
			expected:    "",
			expectError: true,
		},
		{
			name:        "Empty date string",
			input:       "",
			expected:    "",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := convertToDate(tt.input)

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

func TestAppendProjectToDate(t *testing.T) {
	tests := []struct {
		name        string
		project     string
		dateStr     string
		expected    string
		expectError bool
	}{
		{
			name:        "Valid project and date",
			project:     "/path/to/project",
			dateStr:     "2025-05-11 15:41:59 +0900",
			expected:    "project-20250511",
			expectError: false,
		},
		{
			name:        "Valid project with no path and valid date",
			project:     "project",
			dateStr:     "2024-12-28 00:05:43 +0900",
			expected:    "project-20241228",
			expectError: false,
		},
		{
			name:        "Invalid date format",
			project:     "/path/to/project",
			dateStr:     "11-05-2025 15:41:59",
			expected:    "",
			expectError: true,
		},
		{
			name:        "Empty date string",
			project:     "/path/to/project",
			dateStr:     "",
			expected:    "",
			expectError: true,
		},
		{
			name:        "Empty project path",
			project:     "",
			dateStr:     "2025-05-11 15:41:59 +0900",
			expected:    "",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := appendProjectToDate(tt.project, tt.dateStr)

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
