package matchers

import (
	"fmt"

	"github.com/clearblade/cbtest/modules/should/format"
	"github.com/clearblade/cbtest/modules/should/matcher"
	"github.com/clearblade/cbtest/modules/should/to"
)

// ServiceResponseMatcher checks that a service response object has the expected
// success value and results.
type ServiceResponseMatcher struct {
	ExpectedSuccess bool
	ExpectedResults interface{}
	success         bool
	results         interface{}
	matcher         matcher.Matcher
}

// Match returns true if actual matches the expected response.
func (sr *ServiceResponseMatcher) Match(actual interface{}) (bool, error) {

	var ok bool

	resp, ok := actual.(map[string]interface{})
	if !ok {
		return false, fmt.Errorf("actual must have type map[string]interface{}")
	}

	sr.success, ok = resp["success"].(bool)
	if !ok {
		return false, nil
	}

	sr.results, ok = resp["results"]
	if !ok {
		return false, nil
	}

	if sr.success != sr.ExpectedSuccess {
		return false, nil
	}

	sr.matcher, ok = sr.ExpectedResults.(matcher.Matcher)
	if !ok {
		sr.matcher = to.Equal(sr.ExpectedResults)
	}

	return sr.matcher.Match(sr.results)
}

// FailureMessage returns the failure message.
func (sr *ServiceResponseMatcher) FailureMessage(actual interface{}) string {
	return format.Message(sr.formatState(sr.success, sr.results), "to match", sr.formatState(sr.ExpectedSuccess, sr.ExpectedResults))
}

// NegatedFailureMessage returns the negated failure message.
func (sr *ServiceResponseMatcher) NegatedFailureMessage(actual interface{}) string {
	return format.Message(sr.formatState(sr.success, sr.results), "not to match", sr.formatState(sr.ExpectedSuccess, sr.ExpectedResults))
}

func (sr *ServiceResponseMatcher) formatState(success bool, results interface{}) string {
	return fmt.Sprintf("ServiceResponse<success: %t, results: %#v>", success, results)
}
