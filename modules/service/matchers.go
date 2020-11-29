package service

import (
	"github.com/clearblade/cbtest/modules/service/internal/matchers"
	"github.com/clearblade/cbtest/modules/should/matcher"
)

// Matcher is an alias to should.Matcher.
type Matcher matcher.Matcher

// ResponseEqual returns a matcher that checks whenever a code service response
// succeeded with the expected results.
func ResponseEqual(success bool, expectedResults interface{}) Matcher {
	return &matchers.ServiceResponseMatcher{ExpectedSuccess: success, ExpectedResults: expectedResults}
}

// ResponseSuccess returns a matcher that checks whenever a code service response
// was successful.
func ResponseSuccess(expectedResults interface{}) Matcher {
	return ResponseEqual(true, expectedResults)
}

// ResponseError returns a matcher that checks whenever a code service response
// was an error.
func ResponseError(expectedResults interface{}) Matcher {
	return ResponseEqual(false, expectedResults)
}
