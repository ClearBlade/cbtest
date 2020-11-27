// Package check contains miscellaneous assertions and matchers for making the
// tests more expresive. The package is not named something like assert to avoid
// collision with existing packages (testify/assert, etc).
//
// Matcher interface
//
// This package defines a Matcher interface that other modules can use for
// defining their own assertions. The Matcher interface is also compatible
// with Gomega, so any matcher from the latter should work fine.
//
// Usage
//
// A matcher is an instance of an object that performs a check. The matcher
// is usually independent of the actual object being checked. We can apply
// a matcher to an actual object as follows:
//
//     check.Verify(t, ACTUAL, MATCHER)
//
// The code above will fail the test if the ACTUAL object fails the MATCHER. If
// you don't want to fail the test but instead get a boolean value indicating the
// result of the check you can do:
//
//     check.VerifyE(t, ACTUAL, MATCHER)
//
// Negating a check is also possible:
//
//     check.Refute(t, ACTUAL, MATCHER)
//     ...or
//     check.Verify(t, ACTUAL, check.Not(MATCHER))
//
package check
