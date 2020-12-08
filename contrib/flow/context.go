package flow

import (
	"context"
	"time"
)

// Context interface wraps around context.Context to provide saving (stash)
// and laoding (unstash) of context values instead of just Value. For more
// information regarding each of the functions please refer to the context.Context
// docs in the standard library.
type Context interface {

	// Deadline returns the time when work on behalf of this context should
	// be cancalled.
	Deadline() (deadline time.Time, ok bool)

	// Done returs a channel that is closed when work on behalf of this context
	// should be cancelled.
	Done() <-chan struct{}

	// Err returns nil if Done is not closed. Otherwise, it returns an error
	// explaning why.
	Err() error

	// Indentifier returns an integer that uniquely identifies this context
	// life-cycle in a group (sequence, parallel, etc).
	Identifier() int

	// Stash saves a value for later use down the pipeline.
	Stash(key interface{}, value interface{})

	// Untash loads a previously stashed value.
	Unstash(key interface{}) interface{}

	// Unwrap returns the underlying context.Context.
	Unwrap() context.Context
}
