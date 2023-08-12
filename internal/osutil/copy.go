package osutil

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func CopyFiles(src, dst string, overwrite bool) error {
	// Ensure destination directory exists
	if exists, err := DirExists(dst); err != nil {
		return err
	} else if !exists {
		return fmt.Errorf("destination directory %s does not exist", dst)
	}

	err := filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		dstPath := filepath.Join(dst, relPath)

		if info.IsDir() {
			if err := CreateDirs(dstPath); err != nil {
				return err
			}
		} else {
			if exists, err := FileExists(dstPath); err != nil {
				return err
			} else if exists && !overwrite {
				fmt.Printf("File %s already exists. Skipping...\n", dstPath)
				return nil
			}

			if err := copyFile(path, dstPath); err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Some files could not be written: %s\n", err)
		return err
	}

	return nil
}

func CopyFile(src, dst string, overwrite bool) error {
	// Ensure destination directory exists
	if exists, err := DirExists(dst); err != nil {
		return err
	} else if !exists {
		return fmt.Errorf("destination directory %s does not exist", dst)
	}

	if exists, err := FileExists(dst); err != nil {
		return err
	} else if exists && !overwrite {
		fmt.Printf("File %s already exists. Skipping...\n", dst)
		return nil
	}

	return copyFile(src, dst)
}

func copyFile(src, dst string) error {
	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)
	return err
}
