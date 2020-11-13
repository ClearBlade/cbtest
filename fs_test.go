package cbtest

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIsDirSucceeds(t *testing.T) {
	assert.True(t, IsDir("golden"))

	assert.True(t, IsDir("golden/foo"))
	assert.True(t, IsDir("golden/bar"))
	assert.True(t, IsDir("golden/baz"))

	assert.False(t, IsDir("golden/foo/foo.txt"))
	assert.False(t, IsDir("golden/bar/bar.txt"))
	assert.False(t, IsDir("golden/baz/baz.txt"))
}

func TestIsFileSucceeds(t *testing.T) {
	assert.False(t, IsFile("golden"))

	assert.False(t, IsFile("golden/foo"))
	assert.False(t, IsFile("golden/bar"))
	assert.False(t, IsFile("golden/baz"))

	assert.True(t, IsFile("golden/foo/foo.txt"))
	assert.True(t, IsFile("golden/bar/bar.txt"))
	assert.True(t, IsFile("golden/baz/baz.txt"))
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

	err := MergeFolders(tempdir, "golden/foo")
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

	err := MergeFolders(tempdir, "golden/foo", "golden/bar")
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

	err := MergeFolders(tempdir, "golden/foo", "golden/bar", "golden/baz")
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

	err := MergeFolders(tempdir, "golden")
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
		fmt.Sprintf("%s/sys", tempdir),
		fmt.Sprintf("%s/sys/system.json", tempdir),
	}

	assert.Equal(t, expected, listFiles(tempdir))
}
