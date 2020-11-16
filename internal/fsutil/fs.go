package fsutil

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	cp "github.com/otiai10/copy"
)

// defaultIgnore contains all the folders that are gonna be ignored during
// the merging process.
var defaultIgnore = []string{
	"node_modules",
}

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

// MergeFolders merges all the content of `srcs` folders into a single `dest` path.
func MergeFolders(dest string, srcs ...string) error {
	skip := MakeSkip(defaultIgnore)
	return MergeFoldersWithSkip(skip, dest, srcs...)
}

// MergeFoldersWithSkip merges all the contents of `srcs` folders into a single
// `dest` path using the given skip function.
func MergeFoldersWithSkip(skip func(string) (bool, error), dest string, srcs ...string) error {

	if !IsDir(dest) {
		return fmt.Errorf("dest is not a directory: %s", dest)
	}

	options := cp.Options{Skip: skip}

	for _, src := range srcs {

		if !IsDir(src) {
			return fmt.Errorf("src is not a directory: %s", src)
		}

		err := cp.Copy(src, dest, options)
		if err != nil {
			return err
		}
	}

	return nil

}

// MakeSkip returns a skip function that skips any of the given names.
func MakeSkip(names []string) func(string) (bool, error) {
	return func(src string) (bool, error) {
		for _, name := range names {
			if strings.Contains(src, name) {
				return true, nil
			}
		}
		return false, nil
	}
}
