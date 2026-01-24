package utils

import (
	"path/filepath"
	"regexp"
	"strings"
)

func ValidateModuleUrl(moduleUrl string) bool {
	if moduleUrl == "" {
		return false
	}

	if strings.ToLower(moduleUrl) != moduleUrl {
		return false
	}

	modulePathPattern := `^[a-z0-9][a-z0-9\-_.~/]*[a-z0-9]$`
	matched, err := regexp.MatchString(modulePathPattern, moduleUrl)

	if err != nil || !matched {
		return false
	}

	if !strings.Contains(moduleUrl, ".") {
		return false
	}

	invalidSequences := []string{"//", "..", ".-", "-.", "/_", "_/"}

	for _, seq := range invalidSequences {
		if strings.Contains(moduleUrl, seq) {
			return false
		}
	}

	return true
}

func IsGoModule(path string) (bool, error) {
	return FileExists(filepath.Join(path, "go.mod"))
}
