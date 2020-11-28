package check

import "github.com/onsi/gomega/format"

// FormatMessage returns a formatted error message.
func FormatMessage(actual interface{}, message string, expected ...interface{}) string {
	return format.Message(actual, message, expected...)
}

// FormatMessageWithDiff returns a formatted error message.
func FormatMessageWithDiff(actual, message, expected string) string {
	return format.MessageWithDiff(actual, message, expected)
}

// FormatObject returns the object parameter as a formatted string.
func FormatObject(object interface{}, indentation uint) string {
	return format.Object(object, indentation)
}