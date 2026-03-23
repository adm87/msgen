package cli

import (
	"os"
	"path/filepath"
)

func IsModule(path string) bool {
	goModPath := filepath.Join(path, "go.mod")
	if _, err := os.Stat(goModPath); os.IsNotExist(err) {
		return false
	}
	return true
}
