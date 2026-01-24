package utils

import "os"

func FSNodeExists(path string) (bool, error) {
	_, err := os.Stat(path)
	return err == nil || !os.IsNotExist(err), err
}

func FileExists(path string) (bool, error) {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false, nil
	}
	return err == nil && info.Mode().IsRegular(), err
}

func DirExists(path string) (bool, error) {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false, nil
	}
	return err == nil && info.IsDir(), err
}
