package check

import (
	"fmt"

	"github.com/clearblade/cbtest"
)

// Expect checks whenever the given actual value passes the matcher.
// Panics on failure (just like testify/require).
func Expect(t cbtest.T, actual interface{}, matcher Matcher, labelAndArgs ...interface{}) {
	t.Helper()
	res := ExpectE(t, actual, matcher, labelAndArgs...)
	if !res {
		t.FailNow()
	}
}

// ExpectE checks whenever the given actual value passes the matcher.
// Returns boolean to indicate success or failure (just like testify/assert).
func ExpectE(t cbtest.T, actual interface{}, matcher Matcher, labelAndArgs ...interface{}) bool {
	t.Helper()

	label := buildFailureLabel(labelAndArgs...)
	match, err := matcher.Match(actual)

	if err != nil {
		t.Errorf("\n%s%s", label, err.Error())
		return false
	}

	if !match {
		t.Errorf("\n%s%s", label, matcher.FailureMessage(actual))
		return false
	}

	return true
}

// Refute checks whenever the given actual value fails the matcher.
// Panics on failure (just like testify/require).
func Refute(t cbtest.T, matcher Matcher, actual interface{}, labelAndArgs ...interface{}) {
	t.Helper()
	Expect(t, actual, Not(matcher), labelAndArgs...)
}

// RefuteE checks whenever the given actual value fails the matcher.
// Returns boolean to indicate success or failure (just like testify/assert).
func RefuteE(t cbtest.T, matcher Matcher, actual interface{}, labelAndArgs ...interface{}) bool {
	t.Helper()
	return ExpectE(t, actual, Not(matcher), labelAndArgs...)
}

// buildFailureLabel returns the failure label (if passed).
func buildFailureLabel(labelAndArgs ...interface{}) string {

	if len(labelAndArgs) == 0 {
		return ""
	}

	return fmt.Sprintf(labelAndArgs[0].(string), labelAndArgs[1:]...) + "\n"
}
