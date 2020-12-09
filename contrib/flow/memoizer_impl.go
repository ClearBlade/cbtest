package flow

import (
	"context"
	"sync"
)

type memoizerImpl struct {
	contexts map[interface{}]Context
	mu       sync.Mutex
}

// NewMemoizer returns a new Memoizer instance.
func NewMemoizer() Memoizer {
	return &memoizerImpl{
		contexts: make(map[interface{}]Context),
	}
}

func (m *memoizerImpl) Get(key interface{}) Context {
	m.mu.Lock()
	defer m.mu.Unlock()

	if ctx, ok := m.contexts[key]; ok {
		return ctx
	}

	ctx := NewContext(context.TODO(), len(m.contexts))
	m.contexts[key] = ctx
	return ctx
}
