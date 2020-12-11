package flow

// Borrower defines an interfaces for "borrowing" context(s). This is what we use
// for ensuring no workers are using the same context at the same time.
type Borrower interface {

	// Borrow returns a context and a release function for releasing the borrow.
	// If the context was already borrowed, an error will be returned instead.
	Borrow() (Context, func(), error)
}
