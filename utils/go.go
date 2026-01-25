package utils

import (
	"errors"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	ErrMustNotBeEmpty     = errors.New("value must not be empty")
	ErrMustBeLowercase    = errors.New("value must be lowercase")
	ErrInvalidModulePath  = errors.New("value contains invalid characters for a module path")
	ErrMustContainDot     = errors.New("module path must contain at least one dot (.)")
	ErrContainsInvalidSeq = errors.New("module path contains invalid sequences")
)

func ValidateModuleUrl(moduleUrl string) error {
	if moduleUrl == "" {
		return ErrMustNotBeEmpty
	}

	if strings.ToLower(moduleUrl) != moduleUrl {
		return ErrMustBeLowercase
	}

	modulePathPattern := `^[a-z0-9][a-z0-9\-_.~/]*[a-z0-9]$`
	matched, err := regexp.MatchString(modulePathPattern, moduleUrl)

	if err != nil || !matched {
		return ErrInvalidModulePath
	}

	if !strings.Contains(moduleUrl, ".") {
		return ErrMustContainDot
	}

	invalidSequences := []string{"//", "..", ".-", "-.", "/_", "_/"}

	for _, seq := range invalidSequences {
		if strings.Contains(moduleUrl, seq) {
			return ErrContainsInvalidSeq
		}
	}

	return nil
}

func IsGoModule(path string) (bool, error) {
	exists, err := FileExists(filepath.Join(path, "go.mod"))

	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return false, err
	}

	return exists, nil
}
