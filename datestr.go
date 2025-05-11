package main

import (
	"fmt"
	"path/filepath"
	"time"
)

func convertToDate(dateStr string) (string, error) {
	date, err := time.Parse("2006-01-02 15:04:05 -0700", dateStr)
	if err != nil {
		return "", err
	}
	formatted := date.Format("20060102")
	return formatted, nil
}

func appendProjectToDate(project string, dateStr string) (string, error) {
	formatted, err := convertToDate(dateStr)
	if err != nil {
		return "", err
	}
	base := filepath.Base(project)
	return fmt.Sprintf("%s-%s", base, formatted), nil
}
