package merge

import (
	"fmt"
	"strings"

	cp "github.com/otiai10/copy"

	"github.com/clearblade/cbtest/internal/fsutil"
)

// Folders merges all the content of `srcs` folders into a single `dest` path.
func Folders(dest string, srcs ...string) error {
	options := DefaultOptions()
	return FoldersWithOptions(options, dest, srcs...)
}

// FoldersWithOptions merges all the contents of `srcs` folders into a single
// `dest` path using the given options.
func FoldersWithOptions(options *Options, dest string, srcs ...string) error {

	if !fsutil.IsDir(dest) {
		return fmt.Errorf("dest is not a directory: %s", dest)
	}

	cpOptions := cp.Options{Skip: options.Skip}

	for _, src := range srcs {

		if !fsutil.IsDir(src) {
			return fmt.Errorf("src is not a directory: %s", src)
		}

		// before merge hook
		if options.OnBeforeMerge != nil {
			options.OnBeforeMerge.OnBeforeMerge(dest, src)
		}

		// merge
		err := cp.Copy(src, dest, cpOptions)
		if err != nil {
			return err
		}

		// after merge hook
		if options.OnAfterMerge != nil {
			options.OnAfterMerge.OnAfterMerge(dest, src)
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
