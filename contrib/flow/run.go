package flow

import (
	"context"
	"io"
	"os"
	"sync"

	"github.com/clearblade/cbtest"
)

// Run runs the given workflow and waits until it is done.
// Panics on failure.
func Run(t cbtest.T, worker Worker) {
	t.Helper()

	if !RunE(t, worker) {
		t.FailNow()
	}
}

// RunE runs the given workflow and waits until it is done.
// Returns boolean indicating success or failure.
func RunE(t cbtest.T, worker Worker) bool {
	t.Helper()
	return RunWithOutputE(t, worker, os.Stdout)
}

// RunWithOutput runs the given workflow and waits until it is done. All output
// is written to the given io.Writer.
// Panics on failure.
func RunWithOutput(t cbtest.T, worker Worker, output io.Writer) {
	t.Helper()

	if !RunWithOutputE(t, worker, output) {
		t.FailNow()
	}
}

// RunWithOutputE runs the given workflow and waits until it is done. All output
// is written to the given io.Writer.
// Returns boolean indicating success or failure.
func RunWithOutputE(t cbtest.T, worker Worker, output io.Writer) bool {
	t.Helper()

	wg := sync.WaitGroup{}
	workerFn := worker
	workerT := newTWithOutput("root", output)
	workerCtx := NewContext(context.TODO(), 0)

	wg.Add(1)
	go workerRunner(&wg, workerFn, workerT, workerCtx)
	wg.Wait()

	return !workerT.Failed()
}
