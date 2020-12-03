// Package fanout provides helpers for running and waiting for multiple *tests*
// (emphasis in that this package SHOULD NOT be used in code that is not tests).
//
// Why do we need this?
//
// The idea is to provide simple helper functions for spinning up many sub-tests
// in parallel and collect their results. Useful for people not familiar with Go
// concurrency and channels.
package fanout
