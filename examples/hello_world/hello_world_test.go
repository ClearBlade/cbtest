// Package hello_world showcases a test expects a response from a service.
package hello_world

import (
	"testing"

	"github.com/clearblade/cbtest/modules/auth"
	"github.com/clearblade/cbtest/modules/service"
	"github.com/clearblade/cbtest/modules/should"
	"github.com/clearblade/cbtest/modules/system"
)

const (
	HelloWorldService = "helloWorld"
)

func TestHelloWorld(t *testing.T) {

	// import into new system
	s := system.UseOrImport(t, "./extra")

	// close the system after the test
	defer system.Finish(t, s)

	// obtain developer client from the ephemeral system
	devClient := auth.LoginAsDev(t, s)

	// call the service
	resp, err := devClient.CallService(s.SystemKey(), HelloWorldService, nil, false)
	should.NoError(t, err)

	// assert response from service
	should.Expect(t, resp, service.ResponseSuccess("Hello, world!"))
}
