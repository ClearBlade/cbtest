package check

type anything struct{}

func (a *anything) Match(actual interface{}) (bool, error) {
	return true, nil
}

func (a *anything) FailureMessage(actual interface{}) string {
	return "anything failure"
}

func (a *anything) NegatedFailureMessage(actual interface{}) string {
	return "negated anything failure"
}

// Anything returns a matcher that always returns true.
func Anything() Matcher {
	return &anything{}
}
