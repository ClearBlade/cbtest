package system

import (
	"github.com/clearblade/cbtest"
	"github.com/clearblade/cbtest/config"
	"github.com/clearblade/cbtest/provider"
	"github.com/stretchr/testify/require"
)

// UseOrImport tries to use an existing system (given by config or flags), if
// that fails, it tries to import the system instead.
// Panics on failure.
func UseOrImport(t cbtest.T, systemPath string, extraPaths ...string) *EphemeralSystem {
	t.Helper()
	system, err := UseOrImportE(t, systemPath, extraPaths...)
	require.NoError(t, err)
	return system
}

// UseOrImportE tries to use an existing system (given by config or flags), if
// that fails, it tries to import the system instead.
// Returns error on failure.
func UseOrImportE(t cbtest.T, systemPath string, extraPaths ...string) (*EphemeralSystem, error) {
	t.Helper()

	config, err := config.ObtainConfig(t)
	if err != nil {
		return nil, err
	}

	return UseOrImportWithConfigE(t, config, systemPath, extraPaths...)
}

// UseOrImportWithConfig tries to use an existing system (given by config), if
// that fails, it tries to import the system instead.
// Panics on failure.
func UseOrImportWithConfig(t cbtest.T, provider provider.Config, systemPath string, extraPaths ...string) *EphemeralSystem {
	t.Helper()
	system, err := UseOrImportWithConfigE(t, provider, systemPath, extraPaths...)
	require.NoError(t, err)
	return system
}

// UseOrImportWithConfigE tries to use an existing system (given by config), if
// that fails, it tries to import the system instead.
// Returns error on failure.
func UseOrImportWithConfigE(t cbtest.T, provider provider.Config, systemPath string, extraPaths ...string) (*EphemeralSystem, error) {
	t.Helper()

	system, err := UseWithConfigE(t, provider)
	if err == nil {
		return system, nil
	}

	t.Logf("Could not find system to use, falling back to import")
	return ImportWithConfigE(t, provider, systemPath, extraPaths...)
}
