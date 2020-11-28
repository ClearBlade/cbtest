package check

import "strings"

type withMatch struct {
	Run     func(actual interface{}) (success bool, err error)
	success bool
	err     error
}

// WithMatch lets you write matchers in a simpler way than implementing the
// Matcher interface manually. Error messages are expected to be in a form
// that can be used in the following sentences.
//
//     [Expected | Did not expect]
//         ACTUAL
//     ERROR
//
// For example, the following function checks that actual is nil:
//
//     check.WithMatch(func(actual interface{}) (bool, error) {
//         if actual != nil {
//             return false, fmt.Errorf("to be nil")
//         }
//         return true, nil
//     })
//
// The error for values that are not nil will show up as:
//
//     Expected
//         ACTUAL
//     to be nil
//
func WithMatch(match func(actual interface{}) (success bool, err error)) Matcher {
	return &withMatch{Run: match}
}

// Match returns whenever actual passes the matcher.
func (wm *withMatch) Match(actual interface{}) (bool, error) {
	wm.success, wm.err = wm.Run(actual)
	return wm.success, wm.err
}

// NegatedFailureMessage returns the failure message.
func (wm *withMatch) FailureMessage(actual interface{}) string {
	return FormatMessage(actual, wm.err.Error())
}

// NegatedFailureMessage returns the negated failure message.
func (wm *withMatch) NegatedFailureMessage(actual interface{}) string {
	msg := FormatMessage(actual, wm.err.Error())
	return strings.Replace(msg, "Expected", "Did not expect", 1)
}
