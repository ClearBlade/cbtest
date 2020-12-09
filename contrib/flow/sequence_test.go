package flow

import (
	"testing"

	"github.com/clearblade/cbtest/mocks"
	"github.com/stretchr/testify/assert"
)

func TestSequence_NoWorkers(t *testing.T) {

	numbers := []int{}
	workflow := sequence()

	mockT := &mocks.T{}
	mockT.On("Helper")
	Run(mockT, workflow)

	mockT.AssertExpectations(t)
	assert.Equal(t, []int{}, numbers)
}

func TestSequence_OneWorker(t *testing.T) {

	numbers := []int{}
	workflow := sequence(
		func(t *T, ctx Context) {
			numbers = append(numbers, 1)
		},
	)

	mockT := &mocks.T{}
	mockT.On("Helper")
	Run(mockT, workflow)

	mockT.AssertExpectations(t)
	assert.Equal(t, []int{1}, numbers)
}

func TestSequence_TwoWorkers(t *testing.T) {

	numbers := []int{}
	workflow := sequence(

		func(t *T, ctx Context) {
			numbers = append(numbers, 1)
		},

		func(t *T, ctx Context) {
			numbers = append(numbers, 2)
		},
	)

	mockT := &mocks.T{}
	mockT.On("Helper")
	Run(mockT, workflow)

	mockT.AssertExpectations(t)
	assert.Equal(t, []int{1, 2}, numbers)
}

func TestSequence_ThreeWorkers(t *testing.T) {

	numbers := []int{}
	workflow := sequence(

		func(t *T, ctx Context) {
			numbers = append(numbers, 1)
		},

		func(t *T, ctx Context) {
			numbers = append(numbers, 2)
		},

		func(t *T, ctx Context) {
			numbers = append(numbers, 3)
		},
	)

	mockT := &mocks.T{}
	mockT.On("Helper")
	Run(mockT, workflow)

	mockT.AssertExpectations(t)
	assert.Equal(t, []int{1, 2, 3}, numbers)
}

func TestSequence_FailedWorkerInterruptsSequence(t *testing.T) {

	numbers := []int{}
	workflow := sequence(

		func(t *T, ctx Context) {
			numbers = append(numbers, 1)
		},

		func(t *T, ctx Context) {
			t.Errorf("Always fail")
		},

		func(t *T, ctx Context) {
			numbers = append(numbers, 4)
		},
	)

	mockT := &mocks.T{}
	mockT.On("Helper")
	mockT.On("FailNow")
	Run(mockT, workflow)

	mockT.AssertExpectations(t)
	assert.Equal(t, []int{1}, numbers)
}
