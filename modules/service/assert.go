package service

import (
	"github.com/clearblade/cbtest"
	"github.com/stretchr/testify/assert"
)

// AssertResponseEqual asserts that the given response is the expected value.
func AssertResponseEqual(t cbtest.T, expected interface{}, response map[string]interface{}, msgAndArgs ...interface{}) bool {
	return assert.Equal(t, expected, response["results"], msgAndArgs...)
}

// AssertResponseNotEqual asserts that the given response is not the given value.
func AssertResponseNotEqual(t cbtest.T, value interface{}, response map[string]interface{}, msgAndArgs ...interface{}) bool {
	return assert.NotEqual(t, value, response["results"], msgAndArgs...)
}
