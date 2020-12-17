package flow

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBorrower_ReturnsSameContext(t *testing.T) {

	borrower := newContextBorrower(context.Background(), 0)

	ctx0, release, err := borrower.Borrow()
	require.NoError(t, err)
	release()

	ctx1, release, err := borrower.Borrow()
	require.NoError(t, err)
	release()

	assert.Same(t, ctx1, ctx0)
}

func TestBorrower_BorrowedReturnsError(t *testing.T) {

	borrower := newContextBorrower(context.Background(), 0)

	_, _, err := borrower.Borrow()
	require.NoError(t, err)

	_, _, err = borrower.Borrow()
	assert.Error(t, err)
}

func TestBorrower_Release(t *testing.T) {

	borrower := newContextBorrower(context.Background(), 0)

	_, release, err := borrower.Borrow()
	require.NoError(t, err)

	release()
	_, _, err = borrower.Borrow()
	assert.NoError(t, err)
}

func TestStashSucceeds(t *testing.T) {

	ctx := newContext(context.Background(), 0)

	ctx.Stash("one", 1)
	ctx.Stash("two", 2)
	ctx.Stash("three", 3)

	assert.Equal(t, 1, ctx.Unstash("one"))
	assert.Equal(t, 2, ctx.Unstash("two"))
	assert.Equal(t, 3, ctx.Unstash("three"))
}

func TestUnstash_FromParentContextSucceeds(t *testing.T) {

	type customKey string

	keyOne := customKey("one")
	keyTwo := customKey("two")
	keyThree := customKey("three")

	rawctx := context.Background()
	rawctx = context.WithValue(rawctx, keyOne, 1)
	rawctx = context.WithValue(rawctx, keyTwo, 2)
	rawctx = context.WithValue(rawctx, keyThree, 3)

	ctx := newContext(rawctx, 0)

	assert.Equal(t, 1, ctx.Unstash(keyOne))
	assert.Equal(t, 2, ctx.Unstash(keyTwo))
	assert.Equal(t, 3, ctx.Unstash(keyThree))
}
