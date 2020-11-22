// Package npm_test showcases a test that executes npm before importing the system.
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
	HelloWorldService = "helloWorld"
)

func TestHelloWorld(t *testing.T) {

	// executes npm before we import the system
	npm.Use(t, "./extra").Install().Run("build")

	// import into new system
	s := system.UseOrImport(t, "./extra")

	// destroy the system after the test
	defer system.Destroy(t, s)

	// obtain developer client from the ephemeral system
	devClient := auth.LoginAsDev(t, s)

	// call the service
	data := map[string]interface{}{"name": "npm!"}
	resp, err := devClient.CallService(s.SystemKey(), HelloWorldService, data, false)
	require.NoError(t, err)

	// assert response from service
	service.AssertResponseEqual(t, "Hello, npm!", resp)
}
