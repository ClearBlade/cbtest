package cbtest

// TestingT is an interface wrapper around *testing.T
type TestingT interface {
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	FailNow()
	Helper()
	Log(args ...interface{})
	Logf(format string, args ...interface{})
}
