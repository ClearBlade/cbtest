package flow

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
