package flow_test

import (
	"testing"

	"github.com/clearblade/cbtest/contrib/flow"
	"github.com/clearblade/cbtest/mocks"
	"github.com/stretchr/testify/assert"
)

func TestSequence_NoWorkers(t *testing.T) {

	numbers := []int{}

	workflow := flow.Sequence()

	mockT := &mocks.T{}
	mockT.On("Helper")
	flow.Run(mockT, workflow)

	mockT.AssertExpectations(t)
	assert.Equal(t, []int{}, numbers)
}

func TestSequence_OneWorker(t *testing.T) {

	numbers := []int{}

	workflow := flow.Sequence(
		func(t *flow.T, ctx flow.Context) {
			numbers = append(numbers, 1)
		},
	)

	mockT := &mocks.T{}
	mockT.On("Helper")
	flow.Run(mockT, workflow)

	mockT.AssertExpectations(t)
	assert.Equal(t, []int{1}, numbers)
}

func TestSequence_TwoWorkers(t *testing.T) {

	numbers := []int{}

	workflow := flow.Sequence(

		func(t *flow.T, ctx flow.Context) {
			numbers = append(numbers, 1)
		},

		func(t *flow.T, ctx flow.Context) {
			numbers = append(numbers, 2)
		},
	)

	mockT := &mocks.T{}
	mockT.On("Helper")
	flow.Run(mockT, workflow)

	mockT.AssertExpectations(t)
	assert.Equal(t, []int{1, 2}, numbers)
}

func TestSequence_ThreeWorkers(t *testing.T) {

	numbers := []int{}

	workflow := flow.Sequence(

		func(t *flow.T, ctx flow.Context) {
			numbers = append(numbers, 1)
		},

		func(t *flow.T, ctx flow.Context) {
			numbers = append(numbers, 2)
		},

		func(t *flow.T, ctx flow.Context) {
			numbers = append(numbers, 3)
		},
	)

	mockT := &mocks.T{}
	mockT.On("Helper")
	flow.Run(mockT, workflow)

	mockT.AssertExpectations(t)
	assert.Equal(t, []int{1, 2, 3}, numbers)
}

func TestSequence_FailedWorker(t *testing.T) {

	numbers := []int{}

	workflow := flow.Sequence(

		func(t *flow.T, ctx flow.Context) {
			numbers = append(numbers, 1)
		},

		func(t *flow.T, ctx flow.Context) {
			numbers = append(numbers, 2)
		},

		func(t *flow.T, ctx flow.Context) {
			t.Errorf("Always fail")
		},
	)

	mockT := &mocks.T{}
	mockT.On("Helper")
	mockT.On("FailNow")
	flow.Run(mockT, workflow)

	mockT.AssertExpectations(t)
	assert.Equal(t, []int{1, 2}, numbers)
}
