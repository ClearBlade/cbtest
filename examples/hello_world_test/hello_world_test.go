package hello_world_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/clearblade/cbtest"
	"github.com/clearblade/cbtest/cbassert"
)

const (
	HelloWorldService = "helloWorld"
)

func TestHelloWorld(t *testing.T) {

	// import into new system
	system := cbtest.ImportSystem(t, "./extra")

	// destroy the system after the test
	defer cbtest.Destroy(t, system)

	// obtain developer client from the ephemeral system
	devClient := cbtest.LoginAsDev(t, system)

	// call the serice
	resp, err := devClient.CallService(system.SystemKey(), HelloWorldService, nil, false)
	require.NoError(t, err)

	// assert response from service
	cbassert.ServiceResponseEqual(t, "Hello, world!", resp)
}
