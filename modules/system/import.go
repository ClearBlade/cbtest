package system

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/clearblade/cblib"

	"github.com/clearblade/cbtest/config"
	"github.com/clearblade/cbtest/internal/fsutil"
	"github.com/clearblade/cbtest/modules/auth"
)

// checkSystem returns true if the given path contains a system.
func checkSystem(path string) error {

	systemJSONPath := filepath.Join(path, "system.json")

	_, err := os.Stat(systemJSONPath)
	if err != nil {
		return fmt.Errorf("not a system: %s", err)
	}

	return nil
}

// ImportSystem imports the system given by merging the base system given by
// `systemPath` and the extra files given by each of the `extraPaths`.
// Panics on error.
func ImportSystem(t *testing.T, systemPath string, extraPaths ...string) *EphemeralSystem {
	t.Helper()
	system, err := ImportSystemE(t, systemPath, extraPaths...)
	require.NoError(t, err)
	return system
}

// ImportSystemE imports the system given by merging the base system given by
// `systemPath` and the extra files given by each of the `extraPaths`.
// Returns an error on failure.
func ImportSystemE(t *testing.T, systemPath string, extraPaths ...string) (*EphemeralSystem, error) {
	t.Helper()

	config, err := config.ObtainConfig()
	if err != nil {
		return nil, err
	}

	return ImportSystemWithConfigE(t, config, systemPath, extraPaths...)
}

// ImportSystemWithConfig imports the system given by merging the base system
// given by `systemPath` and the extra files given by each of the `extraPaths`
// into the platform instance given by the config.
// Panics on error.
func ImportSystemWithConfig(t *testing.T, config *config.Config, systemPath string, extraPaths ...string) *EphemeralSystem {
	t.Helper()

	system, err := ImportSystemWithConfigE(t, config, systemPath, extraPaths...)
	require.NoError(t, err)
	return system
}

// ImportSystemWithConfigE imports the system given by merging the base system
// given by `systemPath` and the extra files given by each of the `extraPaths`
// into the platform instance given by the config.
// Returns error on failure.
func ImportSystemWithConfigE(t *testing.T, config *config.Config, systemPath string, extraPaths ...string) (*EphemeralSystem, error) {
	t.Helper()

	var err error

	err = checkSystem(systemPath)
	if err != nil {
		return nil, err
	}

	// our imported system root will be at a temporary directory
	tempdir, cleanup := fsutil.MakeTempDir()
	system := NewImportedSystem(config, tempdir)

	// the system paths that are gonna be merged into the temporary directory
	merge := make([]string, 0, 1+len(extraPaths))
	merge = append(merge, systemPath)
	merge = append(merge, extraPaths...)

	t.Log("Merging system folders...")
	err = fsutil.MergeFolders(tempdir, merge...)
	if err != nil {
		cleanup()
		return nil, err
	}

	t.Log("Importing system into platform...")
	_, err = cbImportSystem(t, system)
	if err != nil {
		cleanup()
		return nil, err
	}

	t.Log("Registering developer...")
	err = auth.RegisterDevE(t, system, config.Developer.Email, config.Developer.Password)
	if err != nil {
		cleanup()
		return nil, err
	}

	t.Log("Registering user...")
	err = auth.RegisterUserE(t, system, config.User.Email, config.User.Password)
	if err != nil {
		cleanup()
		return nil, err
	}

	t.Logf("Import successful: %s", system.RemoteURL())
	return system, nil
}

// cbImportSystem imports the given system into a remote platform instance. Note
// that this function will modify the passed system and set its system key and
// secret.
// Returns stdout/stderr and error on failure.
func cbImportSystem(t *testing.T, system *EphemeralSystem) (string, error) {
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
func cbImportConfig(t *testing.T, system *EphemeralSystem) cblib.ImportConfig {
	t.Helper()

	name := fmt.Sprintf("cbtest-%s", t.Name())
	nowstr := time.Now().UTC().Format(time.UnixDate)

	return cblib.ImportConfig{
		SystemName:        name,
		SystemDescription: fmt.Sprintf("Created on %s", nowstr),
		ImportUsers:       system.config.Import.ImportUsers,
		ImportRows:        system.config.Import.ImportRows,
	}
}

// UseSystem uses the external (not managed by cbtest) system given
// by the config flag. External systems are never destroyed automatically.
// Panics on failure.
func UseSystem(t *testing.T) *EphemeralSystem {
	t.Helper()
	system, err := UseSystemE(t)
	require.NoError(t, err)
	return system
}

// UseSystemE uses the external (not managed by cbtest) system given
// by the config flag. External systems are never destroyed automatically.
// Returns error on failure.
func UseSystemE(t *testing.T) (*EphemeralSystem, error) {
	t.Helper()

	config, err := config.ObtainConfig()
	if err != nil {
		return nil, err
	}

	return UseSystemWithConfigE(t, config)
}

// UseSystemWithConfig uses the external (not managed by cbtest) system given
// by the config. External systems are never destroyed automatically.
// Panics on error.
func UseSystemWithConfig(t *testing.T, config *config.Config) *EphemeralSystem {
	t.Helper()
	system, err := UseSystemWithConfigE(t, config)
	require.NoError(t, err)
	return system
}

// UseSystemWithConfigE uses the external (not managed by cbtest) system given
// by the config. External systems are never destroyed automatically.
// Returns error on failure.
func UseSystemWithConfigE(t *testing.T, config *config.Config) (*EphemeralSystem, error) {
	t.Helper()
	return NewExternalSystem(config), nil
}
