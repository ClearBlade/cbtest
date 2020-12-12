package flow

// withContext borrows the context from the given Borrower and passes it to the
// worker. Note that the worker will fail if another worker has already borrowed
// the same context.
func withContext(borrower ContextBorrower, worker Worker) Worker {

	return func(t *T, _ Context) {

		if t.Failed() {
			return
		}

		ctx, release, err := borrower.Borrow()
		if err != nil {
			t.Errorf("%s", err)
			t.FailNow()
		}

		defer release()
		workerRunner(nil, worker, t, ctx)
	}
}
