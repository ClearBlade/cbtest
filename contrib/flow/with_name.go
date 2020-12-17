package flow

// withName returns the given worker but with the given name instead of
// the one it already has.
func withName(name string, worker Worker) Worker {

	return func(t *T, ctx Context) {

		if t.Failed() {
			return
		}

		workerT := newSiblingT(t, name)
		workerRunner(nil, worker, workerT, ctx)
	}
}
