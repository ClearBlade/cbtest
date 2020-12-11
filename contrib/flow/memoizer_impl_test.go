package flow_test

import (
	"fmt"

	"github.com/clearblade/cbtest/contrib/flow"
	"github.com/clearblade/cbtest/mocks"
)

func ExampleNewMemoizer() {

	mockT := &mocks.T{}
	mockT.On("Helper")

	// ignore above this line

	memo := flow.NewMemoizer()

	workflow := flow.NewBuilder().Sequence(func(b *flow.Builder) {

		b.WithContext(memo.Get(0)).Run(func(t *flow.T, ctx flow.Context) {
			ctx.Stash("message", "Hello, world!")
		})

		b.WithContext(memo.Get(0)).Run(func(t *flow.T, ctx flow.Context) {
			fmt.Println(ctx.Unstash("message"))
		})

	})

	flow.Run(mockT, workflow)

	// Output:
	// Hello, world!
}
