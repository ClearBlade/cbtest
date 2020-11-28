package collection

import "github.com/clearblade/cbtest/modules/check"

// Matcher is an alias to check.Matcher.
type Matcher check.Matcher

// HaveTotal checks whenever the actual value contains a count field with the
// expected count.
func HaveTotal(count int) Matcher {

	return check.WithTransform(func(actual interface{}) int {

		data, ok := actual.(map[string]interface{})
		if !ok {
			return -1
		}

		count, ok := data["count"].(float64) // NOTE: response comes as float64
		if !ok {
			return -1
		}

		return int(count)

	}, check.Equal(count))

}
