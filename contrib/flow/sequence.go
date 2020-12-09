package flow

import (
	"fmt"
	"sync"
)

// sequence returns a flow that runs the given workers sequentially, while
// waiting for all of them to finish.
func sequence(workers ...Worker) Worker {

	return func(t *T, ctx Context) {

		if t.Failed() {
			return
		}

		wg := sync.WaitGroup{}

		for idx, fn := range workers {

			workerFn := fn
			workerT := newChildT(t, fmt.Sprintf("seq#%d", idx))
			workerCtx := NewContext(ctx.Unwrap(), idx)

			wg.Add(1)
			workerRunner(&wg, workerFn, workerT, workerCtx)
			wg.Wait()
		}

	}
}
