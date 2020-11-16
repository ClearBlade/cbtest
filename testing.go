package cbtest

// T is an interface wrapper around *testing.T
type T interface {
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	FailNow()
	Helper()
	Name() string
	Log(args ...interface{})
	Logf(format string, args ...interface{})
}
