package matchers

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/clearblade/cbtest/modules/should/to"
)

func TestResponseMatcherSucceeds(t *testing.T) {

	matcher := &ServiceResponseMatcher{ExpectedSuccess: true, ExpectedResults: 9}

	actual := map[string]interface{}{"success": true, "results": 9}
	success, err := matcher.Match(actual)

	assert.True(t, success)
	assert.NoError(t, err)
}

func TestResponseMatcher_NoMatch(t *testing.T) {

	matcher := &ServiceResponseMatcher{ExpectedSuccess: true, ExpectedResults: 9}

	actual := map[string]interface{}{"success": false, "results": 0}
	success, err := matcher.Match(actual)

	assert.False(t, success)
	assert.NoError(t, err)
	assert.Regexp(t, ".*", matcher.FailureMessage(actual))
}

func TestResponseMatcher_ResultsMatcher(t *testing.T) {

	matcher := &ServiceResponseMatcher{ExpectedSuccess: true, ExpectedResults: to.BeNumerically(">", 5)}

	actual := map[string]interface{}{"success": true, "results": 9}
	success, err := matcher.Match(actual)

	assert.True(t, success)
	assert.NoError(t, err)
}
