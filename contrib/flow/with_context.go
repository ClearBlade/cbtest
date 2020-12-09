package flow

// withContext returns the given worker but with the given context instead of
// the one it already has.
func withContext(ctx Context, worker Worker) Worker {

	return func(t *T, _ Context) {

		if t.Failed() {
			return
		}

		workerRunner(nil, worker, t, ctx)
	}
}
