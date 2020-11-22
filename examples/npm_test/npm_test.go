// Package npm_test showcases a test that executes npm before importing the system.
// Useful for systems that are written in TypeScript and need to be transpiled
// beforehand.
package npm_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/clearblade/cbtest/contrib/npm"
	"github.com/clearblade/cbtest/modules/auth"
	"github.com/clearblade/cbtest/modules/service"
	"github.com/clearblade/cbtest/modules/system"
)

const (
	SayHelloService = "sayHello"
)

func TestNPMBasedSystem(t *testing.T) {

	// executes npm before we import the system
	npm.Use(t, "./extra").Install().Run("build")

	// import transpiled dist into new system
	s := system.UseOrImport(t, "./extra/dist")

	// destroy the system after the test
	defer system.Destroy(t, s)

	// obtain developer client from the ephemeral system
	devClient := auth.LoginAsDev(t, s)

	// call the service
	data := map[string]interface{}{"name": "npm!"}
	resp, err := devClient.CallService(s.SystemKey(), SayHelloService, data, false)
	require.NoError(t, err)

	// assert response from service
	service.AssertResponseEqual(t, "Hello, npm!", resp)
}
