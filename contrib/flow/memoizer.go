package flow

// Memoizer defines an interface for caching multiple Borrower(s). Useful for
// running workers with shared context(s).
type Memoizer interface {

	// Get returns the same context borrower for multiple invocations with the same key.
	Get(key interface{}) ContextBorrower
}
