// Package format includes useful functions for creating error messages.
package format

import "github.com/onsi/gomega/format"

// Message returns a formatted error message.
func Message(actual interface{}, message string, expected ...interface{}) string {
	return format.Message(actual, message, expected...)
}

// MessageWithDiff returns a formatted error message.
func MessageWithDiff(actual, message, expected string) string {
	return format.MessageWithDiff(actual, message, expected)
}

// Object returns the object parameter as a formatted string.
func Object(object interface{}, indentation uint) string {
	return format.Object(object, indentation)
}
