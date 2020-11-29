package matcher

// Matcher defines an interface for matching values. It purposefully mimicks the
// interface defined by Gomega (for compatibility).
type Matcher interface {
	Match(actual interface{}) (success bool, err error)
	FailureMessage(actual interface{}) (message string)
	NegatedFailureMessage(actual interface{}) (message string)
}
