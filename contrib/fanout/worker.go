package fanout

import (
	"sync"

	"github.com/clearblade/cbtest"
)

// WorkerFunc is a function that represents a worker with a context.
type WorkerFunc func(t cbtest.T, ctx Context)

// workerRunner executes the given worker function with the given arguments.
func workerRunner(wg *sync.WaitGroup, fn WorkerFunc, t cbtest.T, ctx Context) {
	// finished := false

	defer wg.Done()

	defer func() {
		// TODO: recover and better error notification
	}()

	fn(t, ctx)
	// finished = true
}
