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
// However, some useful matchers are re-exported, so you can access them using
// check.MATCHER instead of gomega.MATCHER.
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
// Using with Gomega
//
// Any of the Gomega matchers should work fine:
//
//     ...
//     check.Verify(t, 10, gomega.BeNumerically(">", 5)) // true
//     ...
//
// Ordering of expected and actual
//
// Early testing frameworks originally used the order EXPECTED-ACTUAL, however,
// reversing the order to ACTUAL-EXPECTED allows for more fluent and expressive
// assertions. For instance, writing:
//
//     check.Verify(t, roses, gomega.Equal("red"))
//
// Reads much better than:
//
//     check.Verify(t, gomega.Equal("red"), roses)
//
package check
