package hello_world_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/clearblade/cbtest"
	"github.com/clearblade/cbtest/cbassert"
)

const (
	AdderService = "adder"
)

func TestAdder(t *testing.T) {

	// import into new system
	system := cbtest.ImportSystem(t, "./extra")

	// destroy the system after the test
	defer cbtest.Destroy(t, system)

	// obtain developer client from the ephemeral system
	devClient := cbtest.LoginAsDev(t, system)

	// payload that we will send to the adder service
	payload := map[string]interface{}{"lhs": 3, "rhs": 4}

	// call the serice
	resp, err := devClient.CallService(system.SystemKey(), AdderService, payload, false)
	require.NoError(t, err)

	// assert response from service
	cbassert.ServiceResponseEqual(t, 7.0, resp)
}
