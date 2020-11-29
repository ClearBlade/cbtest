// Package should contains miscellaneous assertions and matchers for making the
// tests more expresive. The package is not named something like assert to avoid
// collision with existing packages.
//
// Matchers
//
// This package defines a Matcher interface that other modules can use for
// defining their own assertions. The Matcher interface is also compatible
// with Gomega, so any matcher from the latter should work fine.
//
// Essential matchers can be found in the should/to package, so you can access
// them using to.MATCHER.
//
// Usage
//
// A matcher is an instance of an object that performs a check. The matcher
// is usually independent of the actual object being matched; if you want
// to apply a matcher on an actual object you can do it as follows:
//
//     should.Expect(t, ACTUAL, MATCHER)
//
// The code above will fail the test if the ACTUAL object fails the MATCHER. If
// you don't want to fail the test but instead get a boolean value indicating the
// result of the match you can do:
//
//     should.ExpectE(t, ACTUAL, MATCHER)
//
// Negating a matcher is also possible:
//
//     should.Refute(t, ACTUAL, MATCHER)
//     ...or
//     should.Expect(t, ACTUAL, to.Not(MATCHER))
//
// Using with Gomega
//
// Any of the Gomega matchers should work fine:
//
//     should.Expect(t, 10, gomega.BeNumerically(">", 5)) // true
//
// Ordering of expected and actual
//
// Early testing frameworks originally used the order EXPECTED-ACTUAL, however,
// reversing the order to ACTUAL-EXPECTED allows for more fluent and expressive
// assertions. For instance, writing:
//
//     should.Expect(t, roses, to.Equal("red"))
//
// Reads much better than:
//
//     should.Expect(t, to.Equal("red"), roses)
//
package should
