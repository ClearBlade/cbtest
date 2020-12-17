package flow_test

import (
	"testing"

	"github.com/clearblade/cbtest/contrib/flow"
	"github.com/clearblade/cbtest/mocks"
)

func TestRun_FailedWorkerFailsTest(t *testing.T) {

	mockT := &mocks.T{}
	mockT.On("Helper")
	mockT.On("FailNow")

	flow.Run(mockT, func(t *flow.T, ctx flow.Context) {
		t.Errorf("Always fail")
	})

	mockT.AssertExpectations(t)
}
