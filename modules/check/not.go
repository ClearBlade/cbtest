package check

type not struct {
	wrapped Matcher
}

func (n *not) Match(actual interface{}) (bool, error) {
	match, err := n.wrapped.Match(actual)
	return !match, err
}

func (n *not) FailureMessage(actual interface{}) string {
	return n.wrapped.NegatedFailureMessage(actual)
}

func (n *not) NegatedFailureMessage(actual interface{}) string {
	return n.wrapped.FailureMessage(actual)
}

// Not negates the given matcher and returns a new matcher.
func Not(matcher Matcher) Matcher {
	return &not{matcher}
}
