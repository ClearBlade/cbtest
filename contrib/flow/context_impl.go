package flow

import (
	"context"
	"fmt"
	"sync"
)

// wrapped is just a type alias to make the embedded context have a better name.
type wrapped = context.Context

// contextImpl implements the both the ContextBorrower and Context interfaces.
type contextImpl struct {
	wrapped
	identifier int
	borrowed   bool
	mu         sync.Mutex
}

// newContextBorrower returns a new ContextBorrower instance.
func newContextBorrower(wrapped context.Context, identifier int) ContextBorrower {
	return &contextImpl{wrapped, identifier, false, sync.Mutex{}}
}

func (ctx *contextImpl) Borrow() (Context, func(), error) {
	ctx.mu.Lock()
	defer ctx.mu.Unlock()

	if ctx.borrowed {
		return nil, nil, fmt.Errorf("context already borrowed")
	}

	release := func() {
		ctx.mu.Lock()
		defer ctx.mu.Unlock()
		ctx.borrowed = false
	}

	ctx.borrowed = true
	return ctx, release, nil
}

// newContext returns a new Context instance.
func newContext(wrapped context.Context, identifier int) Context {
	return &contextImpl{wrapped, identifier, false, sync.Mutex{}}
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
