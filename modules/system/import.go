package system

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/clearblade/cblib"
	"github.com/clearblade/cbtest"
	"github.com/clearblade/cbtest/config"
	"github.com/clearblade/cbtest/internal/fsutil"
	"github.com/clearblade/cbtest/internal/merge"
	"github.com/clearblade/cbtest/modules/auth"
	"github.com/clearblade/cbtest/provider"
)

// checkSystem returns error if the given path does not contain a system.
func checkSystem(path string) error {

	systemJSONPath := filepath.Join(path, "system.json")

	_, err := os.Stat(systemJSONPath)
	if err != nil {
		return fmt.Errorf("not a system: %s", err)
	}

	return nil
}

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

	var err error

	err = checkSystem(systemPath)
	if err != nil {
		return nil, err
	}

	config := provider.Config(t)

	// our imported system root will be at a temporary directory
	tempdir, cleanupLocal := fsutil.MakeTempDir()
	system := NewImportedSystem(config, tempdir)

	// the system paths that are gonna be merged into the temporary directory
	mergePaths := make([]string, 0, 1+len(extraPaths))
	mergePaths = append(mergePaths, systemPath)
	mergePaths = append(mergePaths, extraPaths...)

	t.Log("Merging system folders...")
	err = merge.Folders(tempdir, mergePaths...)
	if err != nil {
		cleanupLocal()
		return nil, err
	}

	t.Log("Registering developer...")
	err = auth.RegisterDevE(t, system, config.Developer.Email, config.Developer.Password)
	if err != nil {
		cleanupLocal()
		return nil, err
	}

	t.Log("Importing system into platform...")
	_, err = cbImportSystem(t, system)
	if err != nil {
		cleanupLocal()
		return nil, err
	}

	t.Log("Import successful")
	t.Logf("System URL: %s", system.RemoteURL())
	return system, nil
}

// cbImportSystem imports the given system into a remote platform instance. Note
// that this function will modify the passed system and set its system key and
// secret.
// Returns stdout/stderr and error on failure.
func cbImportSystem(t cbtest.T, system *EphemeralSystem) (string, error) {
	t.Helper()

	importConfig := cbImportConfig(t, system)

	devClient, err := auth.LoginAsDevE(t, system)
	if err != nil {
		return "", err
	}

	result, err := cblib.ImportSystemUsingConfig(importConfig, system.LocalPath(), devClient)
	if err != nil {
		return "", err
	}

	system.config.SystemKey = result.SystemKey
	system.config.SystemSecret = result.SystemSecret
	return "", nil
}

// cbImportConfig returns a cblib.ImportConfig instance for importing the system.
func cbImportConfig(t cbtest.T, system *EphemeralSystem) cblib.ImportConfig {
	t.Helper()

	name := fmt.Sprintf("cbtest-%s", t.Name())
	nowstr := time.Now().UTC().Format(time.UnixDate)

	return cblib.ImportConfig{
		SystemName:             name,
		SystemDescription:      fmt.Sprintf("Created on %s", nowstr),
		ImportUsers:            system.config.Import.ImportUsers,
		ImportRows:             system.config.Import.ImportRows,
		DefaultUserPassword:    system.config.User.Password,
		DefaultDeviceActiveKey: system.config.Device.ActiveKey,
	}
}
