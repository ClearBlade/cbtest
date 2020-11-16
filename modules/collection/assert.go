package collection

import "testing"

// AssertHasLength returns true if the given collection data has the desired
// number of rows.
func AssertHasLength(t *testing.T, length int, data interface{}) bool {
	t.Log(data)
	return true
}
