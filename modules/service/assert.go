package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// AssertResponseEqual asserts that the given response is the expected value.
func AssertResponseEqual(t *testing.T, expected interface{}, response map[string]interface{}, msgAndArgs ...interface{}) bool {
	return assert.Equal(t, expected, response["results"], msgAndArgs...)
}

// AssertResponseNotEqual asserts that the given response is not the given value.
func AssertResponseNotEqual(t *testing.T, value interface{}, response map[string]interface{}, msgAndArgs ...interface{}) bool {
	return assert.NotEqual(t, value, response["results"], msgAndArgs...)
}
