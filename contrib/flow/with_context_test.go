package flow

import (
	"context"
	"testing"

	"github.com/clearblade/cbtest/mocks"
	"github.com/stretchr/testify/assert"
)

func TestWithContext(t *testing.T) {

	number := 0
	override := NewContext(context.TODO(), 0)
	override.Stash("overridden-number", 1)
	workflow := withContext(override, func(t *T, ctx Context) {
		number = ctx.Unstash("overridden-number").(int)
	})

	mockT := &mocks.T{}
	mockT.On("Helper")
	Run(mockT, workflow)

	assert.Equal(t, 1, number)
}
