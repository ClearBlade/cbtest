// Package adder showcases a test that passes parameters to a code service.
package adder

import (
	"testing"

	"github.com/clearblade/cbtest/modules/auth"
	"github.com/clearblade/cbtest/modules/service"
	"github.com/clearblade/cbtest/modules/should"
	"github.com/clearblade/cbtest/modules/system"
)

const (
	AdderService = "adder"
)

func TestAdder(t *testing.T) {

	// import into new system
	s := system.UseOrImport(t, "./extra")

	// close the system after the test
	defer system.Finish(t, s)

	// obtain developer client from the ephemeral system
	devClient := auth.LoginAsDev(t, s)

	// payload that we will send to the adder service
	payload := map[string]interface{}{"lhs": 3, "rhs": 4}

	// call the service
	resp, err := devClient.CallService(s.SystemKey(), AdderService, payload, false)
	should.NoError(t, err)

	// assert response from service
	should.Expect(t, resp, service.ResponseSuccess(7.0))
}
