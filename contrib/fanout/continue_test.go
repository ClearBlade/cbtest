package fanout_test

import (
	"testing"

	"github.com/clearblade/cbtest"
	"github.com/clearblade/cbtest/contrib/fanout"
	"github.com/clearblade/cbtest/mocks"
)

func TestContinue(t *testing.T) {

	// mockT expectations

	mockT := &mocks.T{}
	mockT.On("Helper").Return()
	mockT.On("Logf", "Running group \"%s\"...", "Create numbers")
	mockT.On("Logf", "Waiting for group \"%s\"...", "Create numbers")
	mockT.On("Logf", "Continuing group \"%s\" from \"%s\"...", "Show numbers", "Create numbers")
	mockT.On("Logf", "%s Number is: %d", "fanout/Show_numbers/0:", 0)
	mockT.On("Logf", "%s Number is: %d", "fanout/Show_numbers/1:", 10)
	mockT.On("Logf", "%s Number is: %d", "fanout/Show_numbers/2:", 20)
	mockT.On("Logf", "Waiting for group \"%s\"...", "Show numbers")

	// stage 0 - Create numbers

	createNumbers := fanout.Run(mockT, "Create numbers", 3, func(t cbtest.T, ctx fanout.Context) {
		idx := ctx.Identifier()
		num := idx * 10
		ctx.Stash(idx, num)
	})

	fanout.Wait(mockT, createNumbers)

	// stage 1 - Show numbers

	showNumbers := fanout.Continue(mockT, "Show numbers", createNumbers, func(t cbtest.T, ctx fanout.Context) {
		idx := ctx.Identifier()
		num := ctx.Unstash(idx).(int)
		t.Logf("Number is: %d", num)
	})

	fanout.Wait(mockT, showNumbers)

	// assert

	mockT.AssertExpectations(t)
}
