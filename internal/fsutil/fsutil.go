// Package fsutil contains functions for working with the filesystem.
package fsutil

import (
	"fmt"
	"io/ioutil"
	"os"
)

// IsDir returns true if the given path is a directory.
func IsDir(path string) bool {

	info, err := os.Stat(path)
	if err != nil {
		return false
	}

	return info.IsDir()
}

// IsFile returns true if the path is a file.
func IsFile(path string) bool {

	info, err := os.Stat(path)
	if err != nil {
		return false
	}

	return !info.IsDir()
}

// MakeTempDir returns a new temporary directory path and a cleanup function.
func MakeTempDir() (string, func()) {

	dir, err := ioutil.TempDir("", "cbtest-*")
	if err != nil {
		panic(fmt.Sprintf("could not create temporary directory: %s", err))
	}

	return dir, func() {
		os.RemoveAll(dir)
	}
}
