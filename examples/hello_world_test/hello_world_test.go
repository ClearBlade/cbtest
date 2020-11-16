// Package hello_world_test showcases a test expects a response from a service.
package hello_world_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/clearblade/cbtest"
	"github.com/clearblade/cbtest/modules/service"
	"github.com/clearblade/cbtest/modules/system"
)

const (
	HelloWorldService = "helloWorld"
)

func TestHelloWorld(t *testing.T) {

	// import into new system
	s := system.ImportSystem(t, "./extra")

	// destroy the system after the test
	defer cbtest.Destroy(t, s)

	// obtain developer client from the ephemeral system
	devClient := system.LoginAsDev(t, s)

	// call the service
	resp, err := devClient.CallService(s.SystemKey(), HelloWorldService, nil, false)
	require.NoError(t, err)

	// assert response from service
	service.AssertResponseEqual(t, "Hello, world!", resp)
}
