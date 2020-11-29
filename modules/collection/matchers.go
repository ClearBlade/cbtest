package collection

import (
	"github.com/clearblade/cbtest/modules/should/matcher"
	"github.com/clearblade/cbtest/modules/should/to"
)

// Matcher is an alias to matcher.Matcher.
type Matcher matcher.Matcher

// HaveTotal checks whenever the actual value contains a count field with the
// expected count.
func HaveTotal(count int) Matcher {

	return to.WithTransform(func(actual interface{}) int {

		data, ok := actual.(map[string]interface{})
		if !ok {
			return -1
		}

		count, ok := data["count"].(float64) // NOTE: response comes as float64
		if !ok {
			return -1
		}

		return int(count)

	}, to.Equal(count))

}
