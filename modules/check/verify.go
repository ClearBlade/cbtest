package check

import (
	"github.com/clearblade/cbtest"
	"github.com/stretchr/testify/require"
)

// Verify checks whenever the given actual value passes the matcher.
// Panics on failure.
func Verify(t cbtest.T, actual interface{}, matcher Matcher) {
	t.Helper()
	res := VerifyE(t, actual, matcher)
	require.True(t, res)
}

// VerifyE checks whenever the given actual value passes the matcher.
// Returns error on failure.
func VerifyE(t cbtest.T, actual interface{}, matcher Matcher) bool {
	t.Helper()

	match, err := matcher.Match(actual)

	if err != nil || !match {
		t.Errorf("\n%s", matcher.FailureMessage(actual))
		return false
	}

	return true
}

// Refute checks whenever the given actual value fails the matcher.
// Panics on failure.
func Refute(t cbtest.T, matcher Matcher, actual interface{}) {
	t.Helper()
	Verify(t, actual, Not(matcher))
}

// RefuteE checks whenever the given actual value fails the matcher.
// Returns error on failure.
func RefuteE(t cbtest.T, matcher Matcher, actual interface{}) bool {
	t.Helper()
	return VerifyE(t, actual, Not(matcher))
}
