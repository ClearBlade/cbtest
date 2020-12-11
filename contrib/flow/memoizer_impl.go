package flow

import (
	"context"
	"sync"
)

type memoizerImpl struct {
	borrowers map[interface{}]Borrower
	mu        sync.Mutex
}

// NewMemoizer returns a new Memoizer instance.
func NewMemoizer() Memoizer {
	return &memoizerImpl{
		borrowers: make(map[interface{}]Borrower),
	}
}

func (m *memoizerImpl) Get(key interface{}) Borrower {
	m.mu.Lock()
	defer m.mu.Unlock()

	if b, ok := m.borrowers[key]; ok {
		return b
	}

	b := newBorrower(newContext(context.TODO(), len(m.borrowers)))
	m.borrowers[key] = b
	return b
}
