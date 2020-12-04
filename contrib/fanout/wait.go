package fanout

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/clearblade/cbtest"
	"github.com/stretchr/testify/require"
)

// Wait waits for the given group to finish execution. An optional timeout
// can be provided to issue an error.
// Panics on failure.
func Wait(t cbtest.T, group *Group, timeout ...time.Duration) {
	t.Helper()
	err := WaitE(t, group, timeout...)
	require.NoError(t, err)
}

// WaitE waits for the given group to finish execution. An optional timeout
// can be provided to issue an error (defaults to 30 seconds).
// Returns error on failure.
func WaitE(t cbtest.T, group *Group, timeout ...time.Duration) error {
	t.Helper()

	waitTimeout := time.Second * 30
	if len(timeout) > 0 {
		waitTimeout = timeout[0]
	}

	return waitForGroup(t, group.name, group.wg, waitTimeout)
}

// waitForGroup waits for the given named group to finish. Returns error if the
// timeout is eached before the group finished executing.
func waitForGroup(t cbtest.T, name string, wg *sync.WaitGroup, timeout time.Duration) error {
	t.Helper()

	t.Logf("Waiting for group \"%s\"...", name)

	select {
	case <-waitForWaitGroup(wg):
		return nil
	case <-time.After(timeout):
		msg := fmt.Sprintf("Timed out waiting for group \"%s\"", name)
		t.Log(msg)
		return errors.New(msg)
	}
}

// waitForWaitGroup returns channel that is closed when the given wg.WaitGroup
// returns from the Wait call. Note that the goroutine created in this function
// will leak if the wait group never returns.
func waitForWaitGroup(wg *sync.WaitGroup) chan struct{} {
	c := make(chan struct{})
	go func() {
		wg.Wait()
		close(c)
	}()
	return c
}
