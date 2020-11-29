package check

type anything struct{}

func (a *anything) Match(actual interface{}) (bool, error) {
	return true, nil
}

func (a *anything) FailureMessage(actual interface{}) string {
	return FormatMessage(actual, "to match anything")
}

func (a *anything) NegatedFailureMessage(actual interface{}) string {
	return FormatMessage(actual, "not to match anything")
}

// Anything returns a matcher that always returns true.
func Anything() Matcher {
	return &anything{}
}
