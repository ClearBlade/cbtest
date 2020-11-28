package check

import (
	"github.com/onsi/gomega"
	"github.com/onsi/gomega/types"
)

// Matchers re-exported from gomega (alphabetical order)

// All checks whenever actual satifies all the passed matchers.
func All(matchers ...Matcher) Matcher {
	gmatchers := make([]types.GomegaMatcher, 0, len(matchers))
	for _, matcher := range matchers {
		gmatchers = append(gmatchers, matcher)
	}
	return gomega.SatisfyAll(gmatchers...)
}

// Any checks whenever actual satifies any of the passed matchers.
func Any(matchers ...Matcher) Matcher {
	gmatchers := make([]types.GomegaMatcher, 0, len(matchers))
	for _, matcher := range matchers {
		gmatchers = append(gmatchers, matcher)
	}
	return gomega.SatisfyAny(gmatchers...)
}

// BeEmpty checks whenever actual is empty (string, slice, map, etc).
func BeEmpty() Matcher {
	return gomega.BeEmpty()
}

// BeFalse checks whenever actual is false.
func BeFalse() Matcher {
	return gomega.BeFalse()
}

// BeNil checks whenever actual is nil.
func BeNil() Matcher {
	return gomega.BeNil()
}

// BeNumerically checks whenever actual satisfies the comparison.
func BeNumerically(comparator string, compareTo ...interface{}) Matcher {
	return gomega.BeNumerically(comparator, compareTo...)
}

// BeTrue checks whenever actual is true.
func BeTrue() Matcher {
	return gomega.BeTrue()
}

// BeZero checks whenever actual is the zero value of its type.
func BeZero() Matcher {
	return gomega.BeZero()
}

// ConsistOf checks whenever actual consists of the expected elements (no more,
// no less).
func ConsistOf(elements ...interface{}) Matcher {
	return gomega.ConsistOf(elements...)
}

// ContainElements checks if actual contains all the given elements. Ordering
// does not matter, and you can nest other matchers for the elements.
func ContainElements(elements ...interface{}) Matcher {
	return gomega.ContainElements(elements...)
}

// ContainSubstring checks that actual contains the given substring, additional
// arguments can be passed to construct a formatted string with fmt.Sprintf.
func ContainSubstring(substr string, args ...interface{}) Matcher {
	return gomega.ContainSubstring(substr, args...)
}

// Equal uses deep-equal to compare against the expected element.
func Equal(expected interface{}) Matcher {
	return gomega.Equal(expected)
}

// Equivalent to is like Equal, but a little bit more flexible when checking
// for equality, for example, int(3) is equivalent to float(3.0).
func Equivalent(expected interface{}) Matcher {
	return gomega.BeEquivalentTo(expected)
}

// HaveLen checks that actual as the given length (it has to be either a string,
// slice, map, channel, etc).
func HaveLen(count int) Matcher {
	return gomega.HaveLen(count)
}

// HaveKey checks whenever actual is a map and contains the expected key.
func HaveKey(key interface{}) Matcher {
	return gomega.HaveKey(key)
}

// HaveKeyAndValue checks whenever actual is a map and contains the expected key
// and value.
func HaveKeyAndValue(key, value interface{}) Matcher {
	return gomega.HaveKeyWithValue(key, value)
}

// MatchRegexp checks that actual matches the given regular expression, additional
// arguments can be passed to construct a formatted regex with fmt.Sprintf.
func MatchRegexp(regexp string, args ...interface{}) Matcher {
	return gomega.MatchRegexp(regexp, args...)
}

// Succeed checks whenever actual is not an error.
func Succeed() Matcher {
	return gomega.Succeed()
}

// WithTransform applies the given transform function to the value being matched
// before calling the wrapped matcher. The transform function must take one value
// and return one value.
func WithTransform(transform interface{}, matcher Matcher) Matcher {
	return gomega.WithTransform(transform, matcher)
}
