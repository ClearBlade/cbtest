package should

import (
	"github.com/clearblade/cbtest"
	"github.com/clearblade/cbtest/modules/should/to"
)

// NoError fails the test if the given value is an error. Useful for ensuring
// no errors instead of using if err != nil, etc.
func NoError(t cbtest.T, err interface{}) {
	t.Helper()
	res := NoErrorE(t, err)
	if !res {
		t.FailNow()
	}
}

// NoErrorE checks and returns whenever the given value is not an error.
func NoErrorE(t cbtest.T, err interface{}) bool {
	t.Helper()
	return ExpectE(t, err, to.Success())
}
