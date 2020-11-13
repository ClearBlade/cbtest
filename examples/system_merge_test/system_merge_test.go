package hello_world_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/clearblade/cbtest"
	"github.com/clearblade/cbtest/cbassert"
)

const (
	FooService = "foo"
	BarService = "bar"
)

func TestSystemMerge(t *testing.T) {

	// import into new system
	system := cbtest.ImportSystem(t, "./foo_extra", "./bar_extra")

	// destroy the system after the test
	defer cbtest.Destroy(t, system)

	// obtain developer client from the ephemeral system
	devClient := cbtest.LoginAsDev(t, system)

	// call the foo serice
	foo, err := devClient.CallService(system.SystemKey(), FooService, nil, false)
	require.NoError(t, err)

	// assert response from service
	cbassert.ServiceResponseEqual(t, "foo", foo)

	// call the bar serice
	bar, err := devClient.CallService(system.SystemKey(), BarService, nil, false)
	require.NoError(t, err)

	// assert response from service
	cbassert.ServiceResponseEqual(t, "bar", bar)

}
