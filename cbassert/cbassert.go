package cbassert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// ServiceResponseEqual asserts that the given response is the expected value.
func ServiceResponseEqual(t *testing.T, expected interface{}, response map[string]interface{}, msgAndArgs ...interface{}) bool {
	return assert.Equal(t, expected, response["results"], msgAndArgs...)
}
