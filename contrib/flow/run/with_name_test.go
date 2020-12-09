package run_test

import (
	"testing"

	"github.com/clearblade/cbtest/contrib/flow"
	"github.com/clearblade/cbtest/contrib/flow/run"
	"github.com/clearblade/cbtest/mocks"
	"github.com/stretchr/testify/assert"
)

func TestWithName(t *testing.T) {

	name := ""

	workflow := run.WithName("overridden-name", func(t *flow.T, ctx flow.Context) {
		name = t.Name()
	})

	mockT := &mocks.T{}
	mockT.On("Helper")
	flow.Run(mockT, workflow)

	assert.Equal(t, "overridden-name", name)
}

func TestWithName_Nested(t *testing.T) {

	name := ""

	workflow := run.Sequence(
		run.WithName("overridden-sequence", func(t *flow.T, ctx flow.Context) {
			name = t.Name()
		}),
	)

	mockT := &mocks.T{}
	mockT.On("Helper")
	flow.Run(mockT, workflow)

	assert.Equal(t, "root/overridden-sequence", name)
}

func TestWithName_MoreNested(t *testing.T) {

	name := ""

	workflow := run.Sequence(
		run.Parallel(
			run.WithName("overridden-parallel", func(t *flow.T, ctx flow.Context) {
				name = t.Name()
			}),
		),
	)

	mockT := &mocks.T{}
	mockT.On("Helper")
	flow.Run(mockT, workflow)

	assert.Equal(t, "root/sequence-0/overridden-parallel", name)
}
