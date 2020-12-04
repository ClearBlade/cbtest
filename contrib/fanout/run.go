package fanout

import (
	"fmt"
	"sync"
	"time"

	"github.com/clearblade/cbtest"
	"github.com/stretchr/testify/require"
)

// RunHandler is a function that represents a worker.
type RunHandler func(t cbtest.T, idx int)

// Group holds a reference to a fanout.Each call.
type Group struct {
	parentT     cbtest.T
	name        string
	numParallel int
	fn          RunHandler
	wg          *sync.WaitGroup
}

// Cancel tries to cancel the goroutine.
func (g *Group) Cancel() {
}

// Wait waits for every goroutine to finish.
func (g *Group) Wait(timeout ...time.Duration) {
	g.wg.Wait()
}

// Run runs each of the items in the given sequence in parallel. Note that the
// sequence must be a slice.
// Panics on failure.
func Run(t cbtest.T, name string, numParallel int, fn RunHandler) *Group {
	job, err := RunE(t, name, numParallel, fn)
	require.NoError(t, err)
	return job
}

// RunE runs each of the items in the given sequence in parallel. Note that the
// sequence must be a slice.
// Returns error on failure.
func RunE(t cbtest.T, name string, numParallel int, fn RunHandler) (*Group, error) {

	wg := sync.WaitGroup{}

	for idx := 0; idx < numParallel; idx++ {
		wg.Add(1)
		eachT := newFanoutT(t, name, fmt.Sprintf("%d", idx))
		go eachRunner(eachT, &wg, fn, idx)
	}

	return &Group{t, name, numParallel, fn, &wg}, nil
}

func eachRunner(t cbtest.T, wg *sync.WaitGroup, fn RunHandler, idx int) {

	// finished := false

	defer wg.Done()

	defer func() {
		// recover here and notify so we can panic on wait
	}()

	fn(t, idx)
	// finished = true
}
