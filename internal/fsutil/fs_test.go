package fsutil

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIsDirSucceeds(t *testing.T) {

	tempdir, cleanup := MakeTempDir()
	defer cleanup()

	f, err := ioutil.TempFile(tempdir, "file-*")
	require.NoError(t, err)

	assert.True(t, IsDir(tempdir))
	assert.False(t, IsDir(f.Name()))
}

func TestIsFileSucceeds(t *testing.T) {

	tempdir, cleanup := MakeTempDir()
	defer cleanup()

	f, err := ioutil.TempFile(tempdir, "file-*")
	require.NoError(t, err)

	assert.False(t, IsFile(tempdir))
	assert.True(t, IsFile(f.Name()))
}
