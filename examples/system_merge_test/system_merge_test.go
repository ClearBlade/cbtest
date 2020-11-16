// Package system_merge_test shows how to test against a system created from merging
// multiple systems together (foo_extra and bar_extra).
package system_merge_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/clearblade/cbtest"
	"github.com/clearblade/cbtest/modules/service"
	"github.com/clearblade/cbtest/modules/system"
)

const (
	FooService = "foo"
	BarService = "bar"
)

func TestSystemMerge(t *testing.T) {

	// import into new system
	s := system.ImportSystem(t, "./foo_extra", "./bar_extra")

	// destroy the system after the test
	defer cbtest.Destroy(t, s)

	// obtain developer client from the ephemeral system
	devClient := system.LoginAsDev(t, s)

	// call the foo serice
	foo, err := devClient.CallService(s.SystemKey(), FooService, nil, false)
	require.NoError(t, err)

	// assert response from service
	service.AssertResponseEqual(t, "foo", foo)

	// call the bar serice
	bar, err := devClient.CallService(s.SystemKey(), BarService, nil, false)
	require.NoError(t, err)

	// assert response from service
	service.AssertResponseEqual(t, "bar", bar)

}
