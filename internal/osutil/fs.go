package osutil

import (
	"errors"
	"fmt"
	"os"
)

func FileExists(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false, nil
		}

		return false, err
	}

	if info.IsDir() {
		return false, fmt.Errorf("%v: expected file, got dir", path)
	}

	return true, nil
}

func DirExists(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false, nil
		}

		return false, err
	}

	if !info.IsDir() {
		return false, fmt.Errorf("%v: expected dir, got file", path)
	}

	return true, nil
}

func CreateDirs(paths ...string) error {
	for _, path := range paths {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return err
		}
	}

	return nil
}