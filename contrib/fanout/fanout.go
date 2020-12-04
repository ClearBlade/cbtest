// Package fanout provides helpers for running and waiting for "worker" routines
// not managed by parent testing framework.
//
// Why do we need this?
//
// The idea is to provide simple helper functions for spinning up many workers
// in parallel for testing purposes (consumers, producers, etc).
package fanout
