package merge

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/clearblade/cbtest/internal/fsutil"
)

type hookMock struct {
	mock.Mock
}

func (hm *hookMock) OnBeforeMerge(dest, src string) error {
	args := hm.Called(dest, src)
	return args.Error(0)
}

func (hm *hookMock) OnAfterMerge(dest, src string) error {
	args := hm.Called(dest, src)
	return args.Error(0)
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

func TestMergeFolders_WithNoSourceSucceeds(t *testing.T) {

	tempdir, cleanup := fsutil.MakeTempDir()
	defer cleanup()

	err := Folders(tempdir)
	require.NoError(t, err)

	expected := []string{
		tempdir,
	}

	assert.Equal(t, expected, listFiles(tempdir))
}

func TestMergeFolders_WithSingleSourceSucceeds(t *testing.T) {

	tempdir, cleanup := fsutil.MakeTempDir()
	defer cleanup()

	err := Folders(tempdir, "testdata/foo")
	require.NoError(t, err)

	expected := []string{
		tempdir,
		fmt.Sprintf("%s/foo.txt", tempdir),
	}

	assert.Equal(t, expected, listFiles(tempdir))
}

func TestMergeFolders_WithTwoSourcesSucceeds(t *testing.T) {

	tempdir, cleanup := fsutil.MakeTempDir()
	defer cleanup()

	err := Folders(tempdir, "testdata/foo", "testdata/bar")
	require.NoError(t, err)

	expected := []string{
		tempdir,
		fmt.Sprintf("%s/bar.txt", tempdir),
		fmt.Sprintf("%s/foo.txt", tempdir),
	}

	assert.Equal(t, expected, listFiles(tempdir))
}

func TestMergeFolders_WithThreeSourcesSucceeds(t *testing.T) {

	tempdir, cleanup := fsutil.MakeTempDir()
	defer cleanup()

	err := Folders(tempdir, "testdata/foo", "testdata/bar", "testdata/baz")
	require.NoError(t, err)

	expected := []string{
		tempdir,
		fmt.Sprintf("%s/bar.txt", tempdir),
		fmt.Sprintf("%s/baz.txt", tempdir),
		fmt.Sprintf("%s/foo.txt", tempdir),
	}

	assert.Equal(t, expected, listFiles(tempdir))
}

func TestMergeFolders_WithSingleNestedSourceSucceeds(t *testing.T) {

	tempdir, cleanup := fsutil.MakeTempDir()
	defer cleanup()

	err := Folders(tempdir, "testdata")
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

func TestMergeFoldersWithOptions_OnBeforeMergeSuceeeds(t *testing.T) {

	tempdir, cleanup := fsutil.MakeTempDir()
	defer cleanup()

	hook := &hookMock{}
	hook.On("OnBeforeMerge", mock.Anything, "testdata/foo").Return(nil)

	options := DefaultOptions()
	options.OnBeforeMerge = hook
	err := FoldersWithOptions(options, tempdir, "testdata/foo")
	require.NoError(t, err)

	hook.AssertExpectations(t)
}

func TestMergeFoldersWithOptions_OnAfterMergeSucceeds(t *testing.T) {

	tempdir, cleanup := fsutil.MakeTempDir()
	defer cleanup()

	hook := &hookMock{}
	hook.On("OnAfterMerge", mock.Anything, "testdata/foo").Return(nil)

	options := DefaultOptions()
	options.OnAfterMerge = hook
	err := FoldersWithOptions(options, tempdir, "testdata/foo")
	require.NoError(t, err)

	hook.AssertExpectations(t)
}
