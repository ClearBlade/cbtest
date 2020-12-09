package flow_test

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/clearblade/cbtest/contrib/flow"
	"github.com/clearblade/cbtest/contrib/flow/internal/testutil"
	"github.com/clearblade/cbtest/mocks"
	"github.com/stretchr/testify/assert"
)

func TestBuilder_WithName(t *testing.T) {

	name := "overridden-name"
	workflow := flow.NewBuilder().WithName(name).Run(func(t *flow.T, ctx flow.Context) {
		assert.Equal(t, name, t.Name())
	})

	flow.Run(t, workflow)
}

func TestBuilder_WithContext(t *testing.T) {

	override := flow.NewContext(context.Background(), 0)
	workflow := flow.NewBuilder().WithContext(override).Run(func(t *flow.T, ctx flow.Context) {
		assert.Same(t, override, ctx)
	})

	flow.Run(t, workflow)
}

func TestBuilder_WithName_WithContext(t *testing.T) {

	name := "overridden-name"
	override := flow.NewContext(context.Background(), 0)
	workflow := flow.NewBuilder().WithName(name).WithContext(override).Run(func(t *flow.T, ctx flow.Context) {
		assert.Equal(t, name, t.Name())
		assert.Same(t, override, ctx)
	})

	flow.Run(t, workflow)
}

func TestBuilder_Middleware_Reset(t *testing.T) {

	name := "overridden-name"
	override := flow.NewContext(context.Background(), 0)
	b := flow.NewBuilder()

	flow.Run(
		t,
		b.WithName(name).WithContext(override).Run(func(t *flow.T, ctx flow.Context) {
			assert.Equal(t, name, t.Name())
			assert.Same(t, override, ctx)
		}),
	)

	flow.Run(
		t,
		b.Run(func(t *flow.T, ctx flow.Context) {
			assert.NotEqual(t, name, t.Name())
			assert.NotSame(t, override, ctx)
		}),
	)
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
