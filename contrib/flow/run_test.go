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

func TestStuff(t *testing.T) {

	t.Run("the lazy dog", func(t *testing.T) {
		t.Parallel()

		t.Log("foo")
		t.Log("bar")
		t.Log("baz")

	})

	t.Run("Other", func(t *testing.T) {
		t.Parallel()

		t.Log("one")
		t.Log("two")
		t.Log("three")

	})
}
