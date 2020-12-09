package run

import (
	"fmt"
	"sync"

	"github.com/clearblade/cbtest/contrib/flow"
)

// Parallel runs the given workers in parallel, and waits for all of them to finish.
func Parallel(workers ...flow.Worker) flow.Worker {

	return func(t *flow.T, ctx flow.Context) {

		if t.Failed() {
			return
		}

		wg := sync.WaitGroup{}

		for idx, fn := range workers {

			workerFn := fn
			workerT := flow.NewChildT(t, fmt.Sprintf("parallel-%d", idx))
			workerCtx := flow.NewContext(ctx.Unwrap(), idx)

			wg.Add(1)
			go flow.Work(&wg, workerFn, workerT, workerCtx)
		}

		wg.Wait()
	}
}
