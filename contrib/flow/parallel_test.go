package flow

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/clearblade/cbtest/contrib/flow/internal/testutil"
	"github.com/clearblade/cbtest/mocks"
	"github.com/stretchr/testify/assert"
)

func TestParallel_NoWorkers(t *testing.T) {

	total := int32(0)
	workflow := parallel()

	mockT := &mocks.T{}
	mockT.On("Helper")
	Run(mockT, workflow)

	mockT.AssertExpectations(t)
	assert.Equal(t, int32(0), total)
}

func TestParallel_OneWorker(t *testing.T) {

	total := int32(0)
	workflow := parallel(

		func(t *T, ctx Context) {
			atomic.AddInt32(&total, 1)
		},
	)

	mockT := &mocks.T{}
	mockT.On("Helper")
	Run(mockT, workflow)

	mockT.AssertExpectations(t)
	assert.Equal(t, int32(1), total)
}

func TestParallel_TwoWorkers(t *testing.T) {

	total := int32(0)
	workflow := parallel(

		func(t *T, ctx Context) {
			atomic.AddInt32(&total, 1)
		},

		func(t *T, ctx Context) {
			atomic.AddInt32(&total, 2)
		},
	)

	mockT := &mocks.T{}
	mockT.On("Helper")
	Run(mockT, workflow)

	mockT.AssertExpectations(t)
	assert.Equal(t, int32(3), total)
}

func TestParallel_ThreeWorkers(t *testing.T) {

	total := int32(0)
	workflow := parallel(

		func(t *T, ctx Context) {
			atomic.AddInt32(&total, 1)
		},

		func(t *T, ctx Context) {
			atomic.AddInt32(&total, 2)
		},

		func(t *T, ctx Context) {
			atomic.AddInt32(&total, 4)
		},
	)

	mockT := &mocks.T{}
	mockT.On("Helper")
	Run(mockT, workflow)

	mockT.AssertExpectations(t)
	assert.Equal(t, int32(7), total)
}

func TestParallel_ActuallyParallel(t *testing.T) {

	wg := sync.WaitGroup{}
	wg.Add(3)
	workflow := parallel(

		func(t *T, ctx Context) {
			wg.Done()
			wg.Wait()
		},

		func(t *T, ctx Context) {
			wg.Done()
			wg.Wait()
		},

		func(t *T, ctx Context) {
			wg.Done()
			wg.Wait()
		},
	)

	testutil.Timeout(t, time.Second, func() {
		mockT := &mocks.T{}
		mockT.On("Helper")
		Run(mockT, workflow)
	})
}

func TestParallel_FailedWorker(t *testing.T) {

	total := int32(0)
	workflow := parallel(

		func(t *T, ctx Context) {
			atomic.AddInt32(&total, 1)
		},

		func(t *T, ctx Context) {
			atomic.AddInt32(&total, 2)
		},

		func(t *T, ctx Context) {
			t.Errorf("Always fail")
		},
	)

	mockT := &mocks.T{}
	mockT.On("Helper")
	mockT.On("FailNow")
	Run(mockT, workflow)

	mockT.AssertExpectations(t)
	assert.Equal(t, int32(3), total)
}
