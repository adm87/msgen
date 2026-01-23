package utils

import "path/filepath"

func IsGoModule(path string) bool {
	return FileExists(filepath.Join(path, "go.mod"))
}
