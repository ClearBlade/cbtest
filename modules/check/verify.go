package check

import (
	"fmt"

	"github.com/clearblade/cbtest"
)

// Verify checks whenever the given actual value passes the matcher.
// Panics on failure.
func Verify(t cbtest.T, actual interface{}, matcher Matcher, labelAndArgs ...interface{}) {
	t.Helper()
	res := VerifyE(t, actual, matcher, labelAndArgs...)
	if !res {
		t.Error("Verify failed")
		t.FailNow()
	}
}

// VerifyE checks whenever the given actual value passes the matcher.
// Returns error on failure.
func VerifyE(t cbtest.T, actual interface{}, matcher Matcher, labelAndArgs ...interface{}) bool {
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
// Panics on failure.
func Refute(t cbtest.T, matcher Matcher, actual interface{}, labelAndArgs ...interface{}) {
	t.Helper()
	Verify(t, actual, Not(matcher), labelAndArgs...)
}

// RefuteE checks whenever the given actual value fails the matcher.
// Returns error on failure.
func RefuteE(t cbtest.T, matcher Matcher, actual interface{}, labelAndArgs ...interface{}) bool {
	t.Helper()
	return VerifyE(t, actual, Not(matcher), labelAndArgs...)
}

func buildFailureLabel(labelAndArgs ...interface{}) string {

	if len(labelAndArgs) == 0 {
		return ""
	}

	return fmt.Sprintf(labelAndArgs[0].(string), labelAndArgs[1:]...) + "\n"
}
