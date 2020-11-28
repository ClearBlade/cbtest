package matchers

import (
	"fmt"

	"github.com/clearblade/cbtest/modules/check"
)

// ServiceResponseMatcher checks that a service response object has the expected
// success value and results.
type ServiceResponseMatcher struct {
	ExpectedSuccess bool
	ExpectedResults interface{}
	isResponse      bool
	success         bool
	results         interface{}
}

// Match returns true if actual matches the expected response.
func (sr *ServiceResponseMatcher) Match(actual interface{}) (bool, error) {

	resp, ok := actual.(map[string]interface{})
	if !ok {
		return false, fmt.Errorf("could not handle actual type (%T)", actual)
	}

	sr.success, sr.isResponse = resp["success"].(bool)
	sr.results, _ = resp["results"]

	if !sr.isResponse || sr.success != sr.ExpectedSuccess {
		return false, nil
	}

	matcher, ok := sr.ExpectedResults.(check.Matcher)
	if ok {
		return matcher.Match(sr.results)
	}
	return check.Equal(sr.ExpectedResults).Match(sr.results)
}

// FailureMessage returns the failure message.
func (sr *ServiceResponseMatcher) FailureMessage(actual interface{}) string {
	switch {
	case !sr.isResponse:
		return check.FormatMessage(actual, "to be a service response")
	case sr.success != sr.ExpectedSuccess:
		return check.FormatMessage(actual, "to have success as", sr.ExpectedSuccess)
	default:
		return check.FormatMessage(sr.results, "to have results matching", sr.ExpectedResults)
	}
}

// NegatedFailureMessage returns the negated failure message.
func (sr *ServiceResponseMatcher) NegatedFailureMessage(actual interface{}) string {
	switch {
	case !sr.isResponse:
		return check.FormatMessage(actual, "to be a service response")
	case sr.success != sr.ExpectedSuccess:
		return check.FormatMessage(actual, "to have success as", sr.ExpectedSuccess)
	default:
		return check.FormatMessage(sr.results, "to have results not matching", sr.ExpectedResults)
	}
}
