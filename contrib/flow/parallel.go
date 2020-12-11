package flow

import (
	"fmt"
	"sync"
)

// parallel returns a flow that runs the given workers in parallel, while
// waiting for all of them to finish.
func parallel(workers ...Worker) Worker {

	return func(t *T, ctx Context) {

		if t.Failed() {
			return
		}

		wg := sync.WaitGroup{}

		for idx, fn := range workers {

			workerFn := fn
			workerT := newChildT(t, fmt.Sprintf("par#%d", idx))
			workerCtx := newContext(ctx.Unwrap(), idx)

			wg.Add(1)
			go workerRunner(&wg, workerFn, workerT, workerCtx)
		}

		wg.Wait()
	}
}
