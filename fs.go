package cbtest

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	cp "github.com/otiai10/copy"
)

// FoldersToIgnore contains all the folders that are gonna be ignored during
// the merging process.
var FoldersToIgnore = []string{
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

// MakeTempDir returns a new temporary directory.
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

	if !IsDir(dest) {
		return fmt.Errorf("dest is not a directory: %s", dest)
	}

	options := cp.Options{Skip: shouldSkip}

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

// shouldSkip returns true for any folders that should not be merged.
func shouldSkip(src string) (bool, error) {

	for _, folder := range FoldersToIgnore {
		if strings.Contains(src, folder) {
			return true, nil
		}
	}

	return false, nil
}
