package flow

// Memoizer defines an interface for caching Context(s) values. Useful for
// running workers with shared context.
type Memoizer interface {

	// Get returns the same context for multiple invocations with the same key.
	Get(key interface{}) Context
}
