package collection

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// AssertHasLength returns true if the given collection data has the desired
// number of rows.
func AssertHasLength(t *testing.T, length int, data map[string]interface{}) bool {
	DATA, ok := data["DATA"].([]interface{})
	require.True(t, ok, "could not get DATA from collection")
	return assert.Len(t, DATA, length)
}
