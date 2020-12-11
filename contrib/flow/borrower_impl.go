package flow

import (
	"fmt"
	"sync"
)

// borrowerImpl implements the Borrower interface.
type borrowerImpl struct {
	ctx      Context
	borrowed bool
	mu       sync.Mutex
}

// newBorrower returns a new Borrower instance with the given Context.
func newBorrower(ctx Context) *borrowerImpl {
	return &borrowerImpl{ctx, false, sync.Mutex{}}
}

func (b *borrowerImpl) Borrow() (Context, func(), error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.borrowed {
		return nil, nil, fmt.Errorf("context already borrowed; make sure you are not trying to borrow multiple times")
	}

	b.borrowed = true
	return b.ctx, b.release, nil
}

// release releases the current borrow lock.
func (b *borrowerImpl) release() {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.borrowed = false
}
