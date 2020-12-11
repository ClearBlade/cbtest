package flow

import (
	"context"
	"testing"

	"github.com/clearblade/cbtest/mocks"
	"github.com/stretchr/testify/assert"
)

func TestWithContext(t *testing.T) {

	number := 0
	ctx := newContext(context.Background(), 0)
	ctx.Stash("overridden-number", 1)
	borrower := newBorrower(ctx)

	workflow := withContext(borrower, func(t *T, ctx Context) {
		number = ctx.Unstash("overridden-number").(int)
	})

	mockT := &mocks.T{}
	mockT.On("Helper")
	Run(mockT, workflow)

	mockT.AssertExpectations(t)
	assert.Equal(t, 1, number)
}

func TestWithContext_AlreadyBorrowedFails(t *testing.T) {

	ctx := newContext(context.Background(), 0)
	borrower := newBorrower(ctx)
	_, _, _ = borrower.Borrow()

	workflow := withContext(borrower, func(t *T, ctx Context) {})

	mockT := &mocks.T{}
	mockT.On("Helper")
	mockT.On("FailNow")
	Run(mockT, workflow)

	mockT.AssertExpectations(t)
}
