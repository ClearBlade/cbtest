package flow

import "context"

type wrapped = context.Context

type contextImpl struct {
	wrapped
	identifier int
}

// NewContext returns a new flow.Context instance.
func NewContext(wrapped context.Context, identifier int) Context {
	return &contextImpl{wrapped, identifier}
}

func (ctx *contextImpl) Identifier() int {
	return ctx.identifier
}

func (ctx *contextImpl) Stash(key interface{}, value interface{}) {
	ctx.wrapped = context.WithValue(ctx.wrapped, key, value)
}

func (ctx *contextImpl) Unstash(key interface{}) interface{} {
	return ctx.wrapped.Value(key)
}

func (ctx *contextImpl) Unwrap() context.Context {
	return ctx.wrapped
}
