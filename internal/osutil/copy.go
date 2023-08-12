package osutil

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func CopyFromTo(src, dst string, overwrite bool) error {
	err := filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, _ := filepath.Rel(src, path)
		dstPath := filepath.Join(dst, relPath)

		if info.IsDir() {
			return os.MkdirAll(dstPath, info.Mode())
		}

		if _, err := os.Stat(dstPath); err == nil && !overwrite {
			fmt.Printf("File %s already exists. Skipping...\n", dstPath)
			return nil
		}

		source, err := os.Open(path)
		if err != nil {
			return err
		}
		defer source.Close()

		destination, err := os.Create(dstPath)
		if err != nil {
			return err
		}
		defer destination.Close()

		_, err = io.Copy(destination, source)
		return err
	})

	if err != nil {
		fmt.Printf("Some files could not be written: %s\n", err)
		return err
	}

	return nil
}