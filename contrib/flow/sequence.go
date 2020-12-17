package flow

import (
	"fmt"
	"sync"
)

// sequence returns a flow that runs the given workers sequentially, while
// waiting for all of them to finish.
func sequence(workers ...Worker) Worker {

	return func(t *T, ctx Context) {

		wg := sync.WaitGroup{}

		for idx, fn := range workers {

			if t.Failed() {
				return
			}

			workerFn := fn
			workerT := newChildT(t, fmt.Sprintf("seq#%d", idx))
			workerCtx := newContext(ctx.Unwrap(), idx)

			wg.Add(1)
			workerRunner(&wg, workerFn, workerT, workerCtx)
			wg.Wait()
		}

	}
}
