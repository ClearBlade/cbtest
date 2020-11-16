package fsutil

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIsDirSucceeds(t *testing.T) {
	assert.True(t, IsDir("testdata"))

	assert.True(t, IsDir("testdata/foo"))
	assert.True(t, IsDir("testdata/bar"))
	assert.True(t, IsDir("testdata/baz"))

	assert.False(t, IsDir("testdata/foo/foo.txt"))
	assert.False(t, IsDir("testdata/bar/bar.txt"))
	assert.False(t, IsDir("testdata/baz/baz.txt"))
}

func TestIsFileSucceeds(t *testing.T) {
	assert.False(t, IsFile("testdata"))

	assert.False(t, IsFile("testdata/foo"))
	assert.False(t, IsFile("testdata/bar"))
	assert.False(t, IsFile("testdata/baz"))

	assert.True(t, IsFile("testdata/foo/foo.txt"))
	assert.True(t, IsFile("testdata/bar/bar.txt"))
	assert.True(t, IsFile("testdata/baz/baz.txt"))
}

func listFiles(path string) []string {

	files := make([]string, 0)

	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		files = append(files, path)
		return nil
	})

	return files
}

func TestMergeFoldersWithNoSourceSucceeds(t *testing.T) {

	tempdir, cleanup := MakeTempDir()
	defer cleanup()

	err := MergeFolders(tempdir)
	require.NoError(t, err)

	expected := []string{
		tempdir,
	}

	assert.Equal(t, expected, listFiles(tempdir))
}

func TestMergeFoldersWithSingleSourceSucceeds(t *testing.T) {

	tempdir, cleanup := MakeTempDir()
	defer cleanup()

	err := MergeFolders(tempdir, "testdata/foo")
	require.NoError(t, err)

	expected := []string{
		tempdir,
		fmt.Sprintf("%s/foo.txt", tempdir),
	}

	assert.Equal(t, expected, listFiles(tempdir))
}

func TestMergeFoldersWithTwoSourcesSucceeds(t *testing.T) {

	tempdir, cleanup := MakeTempDir()
	defer cleanup()

	err := MergeFolders(tempdir, "testdata/foo", "testdata/bar")
	require.NoError(t, err)

	expected := []string{
		tempdir,
		fmt.Sprintf("%s/bar.txt", tempdir),
		fmt.Sprintf("%s/foo.txt", tempdir),
	}

	assert.Equal(t, expected, listFiles(tempdir))
}

func TestMergeFoldersWithThreeSourcesSucceeds(t *testing.T) {

	tempdir, cleanup := MakeTempDir()
	defer cleanup()

	err := MergeFolders(tempdir, "testdata/foo", "testdata/bar", "testdata/baz")
	require.NoError(t, err)

	expected := []string{
		tempdir,
		fmt.Sprintf("%s/bar.txt", tempdir),
		fmt.Sprintf("%s/baz.txt", tempdir),
		fmt.Sprintf("%s/foo.txt", tempdir),
	}

	assert.Equal(t, expected, listFiles(tempdir))
}

func TestMergeFoldersWithSingleNestedSourceSucceeds(t *testing.T) {

	tempdir, cleanup := MakeTempDir()
	defer cleanup()

	err := MergeFolders(tempdir, "testdata")
	require.NoError(t, err)

	expected := []string{
		tempdir,
		fmt.Sprintf("%s/bar", tempdir),
		fmt.Sprintf("%s/bar/bar.txt", tempdir),
		fmt.Sprintf("%s/baz", tempdir),
		fmt.Sprintf("%s/baz/baz.txt", tempdir),
		fmt.Sprintf("%s/foo", tempdir),
		fmt.Sprintf("%s/foo/foo.txt", tempdir),
		fmt.Sprintf("%s/qux", tempdir),
	}

	assert.Equal(t, expected, listFiles(tempdir))
}
