package utils

import (
	"path/filepath"
	"regexp"
)

// moduleUrlRegex matches valid Go module URLs. TODO - Consider more explicit validation for better readability.
var moduleUrlRegex = regexp.MustCompile(`^(?![a-zA-Z][a-zA-Z0-9+.-]*://)[a-z0-9](?:[a-z0-9-]*[a-z0-9])?(?:\.[a-z0-9](?:[a-z0-9-]*[a-z0-9])?)+(?:/[a-z0-9._~\-]+)+$`)

func ValidateModuleUrl(moduleUrl string) bool {
	return moduleUrlRegex.MatchString(moduleUrl)
}

func IsGoModule(path string) bool {
	return FileExists(filepath.Join(path, "go.mod"))
}
