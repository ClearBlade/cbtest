package testutil

import (
	"time"

	"github.com/clearblade/cbtest"
)

// Timeout executes the given closure, and if it doesn't complete before the
// timeout expires, the test will fail.
func Timeout(t cbtest.T, after time.Duration, fn func()) {
	t.Helper()

	done := make(chan struct{})

	go func() {
		fn()
		done <- struct{}{}
	}()

	select {
	case <-time.After(after):
		t.Error("Operation timed out")
	case <-done:
	}
}
