// Package hello_world_test showcases a test expects a response from a service.
package hello_world_test

import (
	"testing"

	"github.com/clearblade/cbtest/modules/auth"
	"github.com/clearblade/cbtest/modules/check"
	"github.com/clearblade/cbtest/modules/service"
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
	check.NoError(t, err)

	// assert response from service
	check.Verify(t, resp, service.ResponseSuccess("Hello, world!"))
}
