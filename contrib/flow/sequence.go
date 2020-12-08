package flow

import (
	"fmt"
	"sync"
)

// Sequence runs the given workers sequentially, and waits for all of them to finish.
func Sequence(workers ...Worker) Worker {

	return func(t *T, ctx Context) {

		if t.failed {
			return
		}

		wg := sync.WaitGroup{}

		for idx, fn := range workers {

			workerFn := fn
			workerT := newChildFlowT(t, fmt.Sprintf("sequence-%d", idx))
			workerCtx := newContext(ctx.Unwrap(), idx)

			wg.Add(1)
			workerRunner(&wg, workerFn, workerT, workerCtx)
			wg.Wait()
		}

	}
}
