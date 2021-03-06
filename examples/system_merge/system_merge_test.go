// Package system_merge shows how to test against a system created from merging
// multiple systems together (foo_extra and bar_extra).
package system_merge

import (
	"testing"

	"github.com/clearblade/cbtest/modules/auth"
	"github.com/clearblade/cbtest/modules/service"
	"github.com/clearblade/cbtest/modules/should"
	"github.com/clearblade/cbtest/modules/system"
)

const (
	FooService = "foo"
	BarService = "bar"
)

func TestSystemMerge(t *testing.T) {

	// import into new system
	s := system.UseOrImport(t, "./foo_extra", "./bar_extra")

	// close the system after the test
	defer system.Finish(t, s)

	// obtain developer client from the ephemeral system
	devClient := auth.LoginAsDev(t, s)

	// call the foo service
	foo, err := devClient.CallService(s.SystemKey(), FooService, nil, false)
	should.NoError(t, err)

	// assert response from service
	should.Expect(t, foo, service.ResponseSuccess("foo"))

	// call the bar service
	bar, err := devClient.CallService(s.SystemKey(), BarService, nil, false)
	should.NoError(t, err)

	// assert response from service
	should.Expect(t, bar, service.ResponseSuccess("bar"))
}
