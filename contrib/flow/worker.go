package flow

import (
	"runtime/debug"
	"sync"
)

// Worker is a function that represents a worker with a context.
type Worker func(t *T, ctx Context)

// workerRunner executes the given worker and captures any panic.
func workerRunner(wg *sync.WaitGroup, worker Worker, workerT *T, workerCtx Context) {

	defer func() {
		if wg != nil {
			wg.Done()
		}
	}()

	defer func() {
		if e := recover(); e != nil {
			workerT.Errorf("PANIC:\n%s\n%s", e, debug.Stack())
		}
	}()

	worker(workerT, workerCtx)
}
