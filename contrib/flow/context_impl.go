package flow

import (
	"context"
	"sync"
)

// wrapped is just a type alias to make the embedded context have a better name.
type wrapped = context.Context

// contextImpl implements the Context interface.
type contextImpl struct {
	wrapped
	identifier int
	mu         sync.Mutex
}

// newContext returns a new flow.Context instance.
func newContext(wrapped context.Context, identifier int) Context {
	return &contextImpl{wrapped, identifier, sync.Mutex{}}
}

func (ctx *contextImpl) Identifier() int {
	return ctx.identifier
}

func (ctx *contextImpl) Stash(key interface{}, value interface{}) {
	ctx.mu.Lock()
	defer ctx.mu.Unlock()
	ctx.wrapped = context.WithValue(ctx.wrapped, key, value)
}

func (ctx *contextImpl) Unstash(key interface{}) interface{} {
	return ctx.wrapped.Value(key)
}

func (ctx *contextImpl) Unwrap() context.Context {
	return ctx.wrapped
}
