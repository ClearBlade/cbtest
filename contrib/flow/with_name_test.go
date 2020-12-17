package flow

import (
	"testing"

	"github.com/clearblade/cbtest/mocks"
	"github.com/stretchr/testify/assert"
)

func TestWithName(t *testing.T) {

	name := ""
	workflow := withName("overridden-name", func(t *T, ctx Context) {
		name = t.Name()
	})

	mockT := &mocks.T{}
	mockT.On("Helper")
	Run(mockT, workflow)

	assert.Equal(t, "overridden-name", name)
}
