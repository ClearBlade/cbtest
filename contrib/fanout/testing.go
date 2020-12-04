package fanout

import (
	"fmt"
	"regexp"

	"github.com/clearblade/cbtest"
)

// fanoutT acts as a proxy to cbtest.T in order to inject fanout information
// in the output and error functions.
type fanoutT struct {
	cbtest.T
	parentName string
	name       string
	failed     bool
	terminated bool
}

func newFanoutT(parentT cbtest.T, parentName, name string) *fanoutT {
	return &fanoutT{T: parentT, parentName: parentName, name: name}
}

func (t *fanoutT) prefix() string {
	formatted := fmt.Sprintf("fanout/%s/%s", t.parentName, t.name)
	re := regexp.MustCompile(`\s+`)
	return re.ReplaceAllString(formatted, "_") + ":"
}

func (t *fanoutT) forwardArgs(args []interface{}) []interface{} {
	forwarded := make([]interface{}, 0, len(args)+1)
	forwarded = append(forwarded, t.prefix())
	forwarded = append(forwarded, args...)
	return forwarded
}

func (t *fanoutT) Error(args ...interface{}) {
	t.T.Helper()
	t.failed = true
	forward := t.forwardArgs(args)
	t.T.Error(forward...)
}

func (t *fanoutT) Errorf(format string, args ...interface{}) {
	t.T.Helper()
	t.failed = true
	forward := t.forwardArgs(args)
	t.T.Errorf("%s"+format, forward...)
}

func (t *fanoutT) FailNow() {
	t.T.Helper()
	t.failed = true
	t.terminated = true
	t.T.FailNow()
}

func (t *fanoutT) Name() string {
	t.T.Helper()
	return fmt.Sprintf("%s/%s/%s", t.T.Name(), t.parentName, t.name)
}

func (t *fanoutT) Log(args ...interface{}) {
	t.T.Helper()
	forward := t.forwardArgs(args)
	t.T.Log(forward...) // space is added between each argument
}

func (t *fanoutT) Logf(format string, args ...interface{}) {
	t.T.Helper()
	forward := t.forwardArgs(args)
	t.T.Logf("%s "+format, forward...) // add space after first argument
}
