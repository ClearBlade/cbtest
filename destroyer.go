package cbtest

import (
	"github.com/stretchr/testify/require"
)

// Destroyer specicies a destroy function that can be used for getting rid of a
// resource.
type Destroyer interface {
	Destroy(t T) error
}

// Destroy destroys the given destroyer, and fails the test if any error is
// returned.
func Destroy(t T, destroyer Destroyer) {
	t.Helper()
	err := destroyer.Destroy(t)
	require.NoError(t, err)
}
