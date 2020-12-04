package fanout

import (
	"sync"

	"github.com/clearblade/cbtest"
)

// Group represents a collection of goroutines that run a common computation.
// Each goroutine has its own fanout.Context attached to it to make it easy
// to write highly-concurrent code without worrying about synchronization.
type Group struct {

	// parentT holds a reference to the parent test that is running this group.
	parentT cbtest.T

	// name holds the human-readable name for this group.
	name string

	// testingTs holds all the derived cbtest.T for each of the goroutines in this group.
	testingTs []cbtest.T

	// contexts holds the context for each one of the goroutines in this group.
	contexts []Context

	// fn is the handler function that executes for each of the goroutines.
	fn RunHandler

	// wg is the wait group reference that we use for waiting on each goroutine.
	wg *sync.WaitGroup

	// finished indicates whenever all goroutines in the group are finished.
	finished bool
}
