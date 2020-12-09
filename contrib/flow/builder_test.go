package flow_test

import (
	"sync"
	"testing"
	"time"

	"github.com/clearblade/cbtest/contrib/flow"
	"github.com/clearblade/cbtest/contrib/flow/internal/testutil"
	"github.com/clearblade/cbtest/mocks"
	"github.com/stretchr/testify/assert"
)

func TestBuilder_WithName(t *testing.T) {

	name := ""

	workflow := flow.NewBuilder().WithName("overridden-name").Run(func(t *flow.T, ctx flow.Context) {
		name = t.Name()
	})

	mockT := &mocks.T{}
	mockT.On("Helper")
	flow.Run(mockT, workflow)

	assert.Equal(t, "overridden-name", name)
}

func TestBuilder_Sequence(t *testing.T) {

	numbers := []int{}

	workflow := flow.NewBuilder().Sequence(func(b *flow.Builder) {

		b.Run(func(t *flow.T, ctx flow.Context) {
			numbers = append(numbers, 1)
		})

		b.Run(func(t *flow.T, ctx flow.Context) {
			numbers = append(numbers, 2)
		})

		b.Run(func(t *flow.T, ctx flow.Context) {
			numbers = append(numbers, 4)
		})
	})

	mockT := &mocks.T{}
	mockT.On("Helper")
	flow.Run(mockT, workflow)

	assert.Equal(t, []int{1, 2, 4}, numbers)
}

func TestBuilder_Parallel(t *testing.T) {

	wg := sync.WaitGroup{}
	wg.Add(3)

	workflow := flow.NewBuilder().Parallel(func(b *flow.Builder) {

		b.Run(func(t *flow.T, ctx flow.Context) {
			wg.Done()
			wg.Wait()
		})

		b.Run(func(t *flow.T, ctx flow.Context) {
			wg.Done()
			wg.Wait()
		})

		b.Run(func(t *flow.T, ctx flow.Context) {
			wg.Done()
			wg.Wait()
		})
	})

	testutil.Timeout(t, time.Second, func() {
		mockT := &mocks.T{}
		mockT.On("Helper")
		flow.Run(mockT, workflow)
	})
}
