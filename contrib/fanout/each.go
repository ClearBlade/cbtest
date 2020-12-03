package fanout

// eachHandler is an alias for raw interface{}. Used as an alias since it looks
// better and more clear in the documentation.
type eachHandler interface{}

// Each runs each of the items in the given sequence in parallel. Note that the
// sequence must be a slice.
func Each(t tWithRunParallel, name string, sequence interface{}, fn eachHandler) error {
	return nil
}
