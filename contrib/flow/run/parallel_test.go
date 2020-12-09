package run_test

import (
	"sync/atomic"
	"testing"

	"github.com/clearblade/cbtest/contrib/flow"
	"github.com/clearblade/cbtest/contrib/flow/run"
	"github.com/clearblade/cbtest/mocks"
	"github.com/stretchr/testify/assert"
)

func TestParallel_NoWorkers(t *testing.T) {

	total := int32(0)

	workflow := run.Parallel()

	mockT := &mocks.T{}
	mockT.On("Helper")
	flow.Run(mockT, workflow)

	mockT.AssertExpectations(t)
	assert.Equal(t, int32(0), total)
}

func TestParallel_OneWorker(t *testing.T) {

	total := int32(0)

	workflow := run.Parallel(
		func(t *flow.T, ctx flow.Context) {
			atomic.AddInt32(&total, 1)
		},
	)

	mockT := &mocks.T{}
	mockT.On("Helper")
	flow.Run(mockT, workflow)

	mockT.AssertExpectations(t)
	assert.Equal(t, int32(1), total)
}

func TestParallel_TwoWorkers(t *testing.T) {

	total := int32(0)

	workflow := run.Parallel(

		func(t *flow.T, ctx flow.Context) {
			atomic.AddInt32(&total, 1)
		},

		func(t *flow.T, ctx flow.Context) {
			atomic.AddInt32(&total, 2)
		},
	)

	mockT := &mocks.T{}
	mockT.On("Helper")
	flow.Run(mockT, workflow)

	mockT.AssertExpectations(t)
	assert.Equal(t, int32(3), total)
}

func TestParallel_ThreeWorkers(t *testing.T) {

	total := int32(0)

	workflow := run.Parallel(

		func(t *flow.T, ctx flow.Context) {
			atomic.AddInt32(&total, 1)
		},

		func(t *flow.T, ctx flow.Context) {
			atomic.AddInt32(&total, 2)
		},

		func(t *flow.T, ctx flow.Context) {
			atomic.AddInt32(&total, 3)
		},
	)

	mockT := &mocks.T{}
	mockT.On("Helper")
	flow.Run(mockT, workflow)

	mockT.AssertExpectations(t)
	assert.Equal(t, int32(6), total)
}

func TestParallel_FailedWorker(t *testing.T) {

	total := int32(0)

	workflow := run.Parallel(

		func(t *flow.T, ctx flow.Context) {
			atomic.AddInt32(&total, 1)
		},

		func(t *flow.T, ctx flow.Context) {
			atomic.AddInt32(&total, 2)
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
	assert.Equal(t, int32(3), total)
}
