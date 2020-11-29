package to

import (
	"github.com/clearblade/cbtest/modules/should/format"
	"github.com/clearblade/cbtest/modules/should/matcher"
)

type anything struct{}

func (a *anything) Match(actual interface{}) (bool, error) {
	return true, nil
}

func (a *anything) FailureMessage(actual interface{}) string {
	return format.Message(actual, "to match anything")
}

func (a *anything) NegatedFailureMessage(actual interface{}) string {
	return format.Message(actual, "not to match anything")
}

// Anything returns a matcher that always returns true.
func Anything() matcher.Matcher {
	return &anything{}
}
