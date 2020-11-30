// Package should showcases the should module.
package should

import (
	"testing"

	"github.com/clearblade/cbtest/modules/auth"
	"github.com/clearblade/cbtest/modules/should"
	"github.com/clearblade/cbtest/modules/should/to"
	"github.com/clearblade/cbtest/modules/system"
)

const (
	FakerService = "faker"
)

func TestShould(t *testing.T) {

	// import into new system
	s := system.UseOrImport(t, "./extra")

	// close the system after the test
	defer system.Finish(t, s)

	// obtain developer client from the ephemeral system
	devClient := auth.LoginAsDev(t, s)

	// payload that we will send to the adder service
	payload := map[string]interface{}{"lhs": 3, "rhs": 4}

	// call the service
	resp, err := devClient.CallService(s.SystemKey(), FakerService, payload, false)
	should.NoError(t, err)

	// matcher for checking a number is in the [0, 100] range
	between0to100 := to.All(to.BeNumerically(">=", 0), to.BeNumerically("<=", 100))

	// assert the response expectations
	should.ExpectE(t, resp["success"], to.BeTrue())
	should.ExpectE(t, resp["results"], to.HaveKeyAndValue("foo", "foo"))
	should.ExpectE(t, resp["results"], to.HaveKeyAndValue("bar", "bar"))
	should.ExpectE(t, resp["results"], to.HaveKeyAndValue("baz", "baz"))
	should.ExpectE(t, resp["results"], to.HaveKeyAndValue("one", 1.0))
	should.ExpectE(t, resp["results"], to.HaveKeyAndValue("two", to.BeNumerically("==", 2)))
	should.ExpectE(t, resp["results"], to.HaveKeyAndValue("three", to.BeNumerically("~", 3, 0.1)))
	should.ExpectE(t, resp["results"], to.HaveKeyAndValue("random", between0to100))
	should.ExpectE(t, resp["results"], to.HaveKeyAndValue("message", to.ContainSubstring("lazy dog")))
}
