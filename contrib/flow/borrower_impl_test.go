package flow

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBorrower(t *testing.T) {

	ctx := newContext(context.Background(), 0)
	ctx.Stash("foo", "bar")
	borrower := newBorrower(ctx)

	ctx, _, err := borrower.Borrow()
	require.NoError(t, err)
	assert.Equal(t, "bar", ctx.Unstash("foo"))
}

func TestBorrower_BorrowedReturnsError(t *testing.T) {

	ctx := newContext(context.Background(), 0)
	borrower := newBorrower(ctx)

	_, _, err := borrower.Borrow()
	require.NoError(t, err)

	_, _, err = borrower.Borrow()
	assert.Error(t, err)
}

func TestBorrower_Release(t *testing.T) {

	ctx := newContext(context.Background(), 0)
	borrower := newBorrower(ctx)

	_, release, err := borrower.Borrow()
	require.NoError(t, err)

	release()
	_, _, err = borrower.Borrow()
	assert.NoError(t, err)
}
