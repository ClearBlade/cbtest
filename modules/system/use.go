package system

import (
	"fmt"

	"github.com/stretchr/testify/require"

	"github.com/clearblade/cbtest"
	"github.com/clearblade/cbtest/config"
	"github.com/clearblade/cbtest/provider"
)

// Use uses the external (not managed by cbtest) system given
// by the config flag. External systems are never destroyed automatically.
// Panics on failure.
func Use(t cbtest.T) *EphemeralSystem {
	t.Helper()
	system, err := UseE(t)
	require.NoError(t, err)
	return system
}

// UseE uses the external (not managed by cbtest) system given
// by the config flag. External systems are never destroyed automatically.
// Returns error on failure.
func UseE(t cbtest.T) (*EphemeralSystem, error) {
	t.Helper()

	config, err := config.ObtainConfig(t)
	if err != nil {
		return nil, err
	}

	return UseWithConfigE(t, config)
}

// UseWithConfig uses the external (not managed by cbtest) system given
// by the config. External systems are never destroyed automatically.
// Panics on error.
func UseWithConfig(t cbtest.T, provider provider.Config) *EphemeralSystem {
	t.Helper()
	system, err := UseWithConfigE(t, provider)
	require.NoError(t, err)
	return system
}

// UseWithConfigE uses the external (not managed by cbtest) system given
// by the config. External systems are never destroyed automatically.
// Returns error on failure.
func UseWithConfigE(t cbtest.T, provider provider.Config) (*EphemeralSystem, error) {
	t.Helper()

	config := provider.Config(t)

	if !config.HasSystem() {
		return nil, fmt.Errorf("external system configuration did not specify a system key or secret")
	}

	system := newExternalSystem(config)
	t.Logf("Using existing system")
	t.Logf("Platform URL: %s", system.PlatformURL())
	t.Logf("Messaging URL: %s", system.MessagingURL())
	t.Logf("System key: %s", system.SystemKey())
	t.Logf("System URL: %s", system.RemoteURL())
	return system, nil
}
