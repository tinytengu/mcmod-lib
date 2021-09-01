package utils

import (
	"fmt"
	"strings"
)

func IsValidModEntry(entry string) bool {
	parts := strings.Split(entry, ":")
	return len(parts[0]) != 0 &&
		IsValidMcVersion(parts[1]) &&
		IsValidModType(parts[2]) &&
		IsValidModFile(parts[3])
}

func IsValidModsList(data string) (bool, error) {
	for _, line := range strings.Split(data, "\n") {
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}
		if !IsValidModEntry(line) {
			return false, fmt.Errorf("invalid mod entry: %v", line)
		}
	}
	return true, nil
}
