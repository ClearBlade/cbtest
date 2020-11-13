package cbtest

// Log specifies a function for logging output.
type Log interface {
	Log(args ...interface{})
}

// Logf specifies a functin for logging formatted output.
type Logf interface {
	Logf(format string, args ...interface{})
}
