package utils

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/adm87/msgen/log"
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

func InitGoModule(moduleUrl, path string) error {
	log.Info("Initializing Go module:", "moduleUrl", moduleUrl)

	cmd := exec.Command("go", "mod", "init", moduleUrl)
	cmd.Dir = path

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Error("go mod init failed", "output", string(out))
		return err
	}

	return nil
}

func GoGet(path, moduleUrl string) error {
	log.Info("Running 'go get' for module:", "moduleUrl", moduleUrl)

	cmd := exec.Command("go", "get", moduleUrl)
	cmd.Dir = path

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Error("go get failed", "output", string(out))
		return err
	}

	return nil
}

func ModTidy(path string) error {
	log.Info("Running 'go mod tidy'")

	cmd := exec.Command("go", "mod", "tidy")
	cmd.Dir = path

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Error("go mod tidy failed", "output", string(out))
		return err
	}

	return nil
}

func GoFmt(path string) error {
	log.Info("Running 'go fmt'")

	cmd := exec.Command("go", "fmt", "./...")
	cmd.Dir = path

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Error("go fmt failed", "output", string(out))
		return err
	}

	return nil
}

func GetModuleUrl(path string) (string, error) {
	cmd := exec.Command("go", "list", "-m")
	cmd.Dir = path

	output, err := cmd.Output()

	if err != nil {
		return "", err
	}

	return string(output[:len(output)-1]), nil
}
