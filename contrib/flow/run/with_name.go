package run

import "github.com/clearblade/cbtest/contrib/flow"

// WithName returns the given worker but with the given name instead of
// the one it already has.
func WithName(name string, worker flow.Worker) flow.Worker {

	return func(t *flow.T, ctx flow.Context) {

		if t.Failed() {
			return
		}

		workerT := flow.NewSiblingT(t, name)
		flow.Work(nil, worker, workerT, ctx)
	}
}
