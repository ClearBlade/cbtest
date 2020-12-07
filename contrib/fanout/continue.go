package fanout

import (
	"fmt"
	"sync"

	"github.com/clearblade/cbtest"
	"github.com/stretchr/testify/require"
)

// Continue reuses a previously created group to continue execution in a new
// function as a new *Group.
// Panics on error.
func Continue(t cbtest.T, name string, group *Group, fn WorkerFunc) *Group {
	t.Helper()
	g, err := ContinueE(t, name, group, fn)
	require.NoError(t, err)
	return g
}

// ContinueE reuses a previously created group to continue execution in a new
// function as a new *Group.
// Returns error on failure.
func ContinueE(t cbtest.T, name string, group *Group, fn WorkerFunc) (*Group, error) {
	t.Helper()
	t.Logf("Continuing group \"%s\" from \"%s\"...", name, group.name)

	if !group.finished {
		return nil, makeGroupNotFinishedError(group.name)
	}

	numMembers := len(group.contexts)
	testingTs := make([]cbtest.T, 0, numMembers)
	contexts := make([]Context, 0, numMembers)

	wg := sync.WaitGroup{}
	wg.Add(numMembers)

	for idx := 0; idx < numMembers; idx++ {

		oldContext := group.contexts[idx]
		workerT := newFanoutT(t, name, fmt.Sprintf("%d", idx))
		workerContext := newContext(oldContext.Unwrap(), idx)

		testingTs = append(testingTs, workerT)
		contexts = append(contexts, workerContext)

		go workerRunner(&wg, fn, workerT, workerContext)
	}

	return &Group{t, name, testingTs, contexts, fn, &wg, false}, nil
}

func makeGroupNotFinishedError(groupName string) error {
	return fmt.Errorf("Group \"%s\" is not finished. Make sure you use fanout.Wait[E] before you call fanout.Continue[E]", groupName)
}
