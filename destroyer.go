package cbtest

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Destroyer specicies a destroy function that can be used for getting rid of a
// resource.
type Destroyer interface {
	Destroy(t *testing.T) error
}

// Destroy destroys the given destroyer, and fails the test if any error is
// returned.
func Destroy(t *testing.T, destroyer Destroyer) {
	t.Helper()
	err := destroyer.Destroy(t)
	require.NoError(t, err)
}
