package utils

import (
	"errors"
	"os"
)

type FSNodeType int

const (
	NodeAny FSNodeType = iota
	NodeDirectory
	NodeFile
)

const (
	DirPerm  os.FileMode = 0755
	FilePerm os.FileMode = 0644
)

var (
	ErrInvalidFSNodeType = errors.New("invalid filesystem node type")
)

func FSNodeExists(path string, expectedType FSNodeType) (bool, error) {
	info, err := os.Stat(path)

	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}

		return false, err
	}

	switch expectedType {
	case NodeDirectory:
		return info.IsDir(), nil

	case NodeFile:
		return info.Mode().IsRegular(), nil

	case NodeAny:
		return info.Mode().IsRegular() || info.IsDir(), nil

	default:
		return false, ErrInvalidFSNodeType
	}
}

func FileExists(path string) (bool, error) {
	return FSNodeExists(path, NodeFile)
}

func DirectoryExists(path string) (bool, error) {
	return FSNodeExists(path, NodeDirectory)
}

func MakeDirectory(path string) error {
	return os.MkdirAll(path, DirPerm)
}

func WriteFile(path string, content string) error {
	return os.WriteFile(path, []byte(content), FilePerm)
}
