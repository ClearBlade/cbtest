package flow

import (
	"fmt"
	"sync"
)

// Parallel runs the given workers in parallel, and waits for all of them to finish.
func Parallel(workers ...Worker) Worker {

	return func(t *T, ctx Context) {

		if t.failed {
			return
		}

		wg := sync.WaitGroup{}

		for idx, fn := range workers {

			wg.Add(1)

			workerFn := fn
			workerT := newChildFlowT(t, fmt.Sprintf("parallel-%d", idx))
			workerCtx := newContext(ctx.Unwrap(), idx)

			go workerRunner(&wg, workerFn, workerT, workerCtx)
		}

		wg.Wait()
	}
}
