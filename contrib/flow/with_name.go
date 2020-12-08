package flow

// WithName returns the given worker but with the given name instead of
// the one it already has.
func WithName(name string, worker Worker) Worker {

	return func(t *T, ctx Context) {

		if t.Failed() {
			return
		}

		t.name = name
		workerRunner(nil, worker, t, ctx)
	}
}
