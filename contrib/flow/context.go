package flow

import (
	"context"
	"time"
)

// ContextBorrower defines an interfaces for "borrowing" context(s). This is what we use
// for ensuring no workers are using the same context at the same time.
type ContextBorrower interface {

	// Borrow returns a context and a release function for releasing the borrow.
	// If the context was already borrowed, an error will be returned instead.
	Borrow() (Context, func(), error)
}

// Context simulates the interfaces exposed by context.Context but with storing
// (stash) and loading (unstash) of values. Note that users will never be able
// to create Context instances directly, and will need to go through the borrower
// instead.
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
