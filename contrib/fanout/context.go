package fanout

import (
	"context"
	"sync"
	"time"
)

// Context interface wraps around context.Context to provide saving (stash)
// and laoding (unstash) of context values instead of just Value. For more
// information regarding each of the functions please refer to the context.Context
// docs in the standard library.
//
// Difference with context.Context
//
// Whereas context.Context is meant to be used in real applications, the context
// presented here is meant to be used in tests, where robustness and type-safety
// are not as important.
type Context interface {

	// Consolidate returns a new context.Context with the combined stashed values.
	Consolidate() context.Context

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
	// life-cycle in a Group.
	Identifier() int

	// Stash saves a value for later use down the pipeline.
	Stash(key interface{}, value interface{})

	// Untash loads a previously stashed value.
	Unstash(key interface{}) interface{}
}

type contextImpl struct {
	context.Context
	identifier int
	stash      sync.Map
}

func newContext(parent context.Context, identifier int) *contextImpl {
	return &contextImpl{parent, identifier, sync.Map{}}
}

func (ctx *contextImpl) Consolidate() context.Context {

	result := ctx.Context

	ctx.stash.Range(func(key, value interface{}) bool {
		result = context.WithValue(result, key, value)
		return true // keep iterating
	})

	return result
}

func (ctx *contextImpl) Identifier() int {
	return ctx.identifier
}

func (ctx *contextImpl) Stash(key interface{}, value interface{}) {
	ctx.stash.Store(key, value)
}

func (ctx *contextImpl) Unstash(key interface{}) interface{} {

	value, ok := ctx.stash.Load(key)
	if ok {
		return value
	}

	return ctx.Context.Value(key)
}
