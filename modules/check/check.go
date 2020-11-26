// Package check contains miscellaneous assertions and matchers for making the
// tests more expresive. Package is not named something like assert to avoid
// collision with existing packages.
//
// Matchers
//
// This package defines a Matcher interface that is compatible with Gomega.
// Therefore, any Gomega matcher should be compatible. It also provides the
// Verify[E] helper functions for making integration with Gomega more clean.
package check
