package flow

import (
	"context"
	"sync"
)

type memoizerImpl struct {
	borrowers map[interface{}]ContextBorrower
	mu        sync.Mutex
}

// NewMemoizer returns a new Memoizer instance that can be used for obtaining
// context borrower(s).
//
// Memoization
//
// The first time you call Get, a new borrower will be returned:
//
//     a := memo.Get("some-key")
//
// When you call it again with the same key, the *same* borrower will be returned:
//
//     b := memo.Get("some-key")
//
func NewMemoizer() Memoizer {
	return &memoizerImpl{
		borrowers: make(map[interface{}]ContextBorrower),
	}
}

func (m *memoizerImpl) Get(key interface{}) ContextBorrower {
	m.mu.Lock()
	defer m.mu.Unlock()

	if b, ok := m.borrowers[key]; ok {
		return b
	}

	b := newContextBorrower(context.TODO(), len(m.borrowers))
	m.borrowers[key] = b
	return b
}
