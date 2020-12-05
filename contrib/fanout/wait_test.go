package fanout

import (
	"sync"
	"testing"
	"time"

	"github.com/clearblade/cbtest/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestWaitForGroup(t *testing.T) {

	mockT := &mocks.T{}
	mockT.On("Helper").Return()
	mockT.On("Logf", "Waiting for group \"%s\"...", "Group name").Return()

	wg := sync.WaitGroup{}
	err := waitForGroup(mockT, "Group name", &wg, time.Millisecond)
	assert.NoError(t, err)

	mockT.AssertExpectations(t)
}

func TestWaitForGroup_WaitGroupDone(t *testing.T) {

	mockT := &mocks.T{}
	mockT.On("Helper").Return()
	mockT.On("Logf", mock.Anything, mock.Anything).Return()

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		wg.Done()
	}()

	err := waitForGroup(mockT, "Group name", &wg, time.Millisecond)
	assert.NoError(t, err)
}

func TestWaitGroup_Timeout(t *testing.T) {

	mockT := &mocks.T{}
	mockT.On("Helper").Return()
	mockT.On("Logf", mock.Anything, mock.Anything).Return()
	mockT.On("Log", mock.Anything).Return()

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		time.Sleep(time.Millisecond * 2)
		wg.Done()
	}()

	err := waitForGroup(mockT, "Group name", &wg, time.Millisecond)
	assert.EqualError(t, err, "Timed out waiting for group \"Group name\" (1ms)")
}
