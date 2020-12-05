package fanout

import (
	"context"
	"fmt"
	"sync"

	"github.com/clearblade/cbtest"
	"github.com/stretchr/testify/require"
)

// Run creates a new group with the given amount of members and starts executing
// them. The function will return once all members are up and running.
// Panics on failure.
func Run(t cbtest.T, name string, numMembers int, fn WorkerFunc) *Group {
	t.Helper()
	job, err := RunE(t, name, numMembers, fn)
	require.NoError(t, err)
	return job
}

// RunE creates a new group with the given amount of members and starts executing
// them. The function will return once all members are up and running.
// Returns error on failure.
func RunE(t cbtest.T, name string, numMembers int, fn WorkerFunc) (*Group, error) {
	t.Helper()
	t.Logf("Running group \"%s\"...", name)

	testingTs := make([]cbtest.T, 0, numMembers)
	contexts := make([]Context, 0, numMembers)

	wg := sync.WaitGroup{}
	wg.Add(numMembers)

	for idx := 0; idx < numMembers; idx++ {

		workerT := newFanoutT(t, name, fmt.Sprintf("%d", idx))
		workerContext := newContext(context.TODO(), idx)

		testingTs = append(testingTs, workerT)
		contexts = append(contexts, workerContext)

		go workerRunner(&wg, fn, workerT, workerContext)
	}

	return &Group{t, name, testingTs, contexts, fn, &wg, false}, nil
}
