package system

import (
	"github.com/stretchr/testify/require"

	"github.com/clearblade/cbtest"
	"github.com/clearblade/cbtest/config"
	"github.com/clearblade/cbtest/provider"
)

// Import imports the system given by merging the base system given by
// `systemPath` and the extra files given by each of the `extraPaths`.
// Panics on failure.
func Import(t cbtest.T, systemPath string, extraPaths ...string) *EphemeralSystem {
	t.Helper()
	system, err := ImportE(t, systemPath, extraPaths...)
	require.NoError(t, err)
	return system
}

// ImportE imports the system given by merging the base system given by
// `systemPath` and the extra files given by each of the `extraPaths`.
// Returns error on failure.
func ImportE(t cbtest.T, systemPath string, extraPaths ...string) (*EphemeralSystem, error) {
	t.Helper()

	config, err := config.ObtainConfig(t)
	if err != nil {
		return nil, err
	}

	return ImportWithConfigE(t, config, systemPath, extraPaths...)
}

// ImportWithConfig imports the system given by merging the base system
// given by `systemPath` and the extra files given by each of the `extraPaths`
// into the platform instance given by the config.
// Panics on error.
func ImportWithConfig(t cbtest.T, provider provider.Config, systemPath string, extraPaths ...string) *EphemeralSystem {
	t.Helper()

	system, err := ImportWithConfigE(t, provider, systemPath, extraPaths...)
	require.NoError(t, err)
	return system
}

// ImportWithConfigE imports the system given by merging the base system
// given by `systemPath` and the extra files given by each of the `extraPaths`
// into the platform instance given by the config.
// Returns error on failure.
func ImportWithConfigE(t cbtest.T, provider provider.Config, systemPath string, extraPaths ...string) (*EphemeralSystem, error) {
	t.Helper()

	return runImportWorkflow(t, provider, &defaultImportSteps{
		path:  systemPath,
		extra: extraPaths,
	})
}

// runImportWorkflow is an unexported function that runs the import workflow.
func runImportWorkflow(t cbtest.T, provider provider.Config, steps importSteps) (*EphemeralSystem, error) {
	t.Helper()

	var err error
	config := provider.Config(t)
	systemPath := steps.Path()
	extraPaths := steps.Extra()

	// make sure the system is actually a system
	err = steps.CheckSystem(systemPath)
	if err != nil {
		return nil, err
	}

	// our imported system root will be at a temporary directory
	tempdir, tempcleanup := steps.MakeTempDir()

	// the system paths that are gonna be merged into the temporary directory
	mergePaths := make([]string, 0, 1+len(extraPaths))
	mergePaths = append(mergePaths, systemPath)
	mergePaths = append(mergePaths, extraPaths...)

	t.Log("Merging system folders...")
	err = steps.MergeFolders(tempdir, mergePaths...)
	if err != nil {
		steps.Cleanup(tempdir)
		tempcleanup()
		return nil, err
	}

	t.Log("Registering developer...")
	err = steps.RegisterDeveloper(t, config)
	if err != nil {
		steps.Cleanup(tempdir)
		tempcleanup()
		return nil, err
	}

	t.Log("Importing system into platform...")
	systemKey, systemSecret, err := steps.DoImport(t, config, tempdir)
	if err != nil {
		steps.Cleanup(tempdir)
		tempcleanup()
		return nil, err
	}

	system := newImportedSystem(config, tempdir)
	system.config.SystemKey = systemKey
	system.config.SystemSecret = systemSecret

	t.Log("Import successful")
	t.Logf("Platform URL: %s", system.PlatformURL())
	t.Logf("Messaging URL: %s", system.MessagingURL())
	t.Logf("System key: %s", system.SystemKey())
	t.Logf("System URL: %s", system.RemoteURL())
	return system, nil
}
