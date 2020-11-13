package cbtest

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	cb "github.com/clearblade/Go-SDK"
	"github.com/clearblade/cblib"
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
	system, err := ImportSystemE(t, systemPath, extraPaths...)
	require.NoError(t, err)
	return system
}

// ImportSystemE imports the system given by merging the base system given by
// `systemPath` and the extra files given by each of the `extraPaths`.
// Returns an error on failure.
func ImportSystemE(t *testing.T, systemPath string, extraPaths ...string) (*EphemeralSystem, error) {

	config, err := ObtainConfig()
	if err != nil {
		return nil, err
	}

	return ImportSystemWithConfigE(t, config, systemPath, extraPaths...)
}

// ImportSystemWithConfig imports the system given by merging the base system
// given by `systemPath` and the extra files given by each of the `extraPaths`
// into the platform instance given by the config.
// Panics on error.
func ImportSystemWithConfig(t *testing.T, config *Config, systemPath string, extraPaths ...string) *EphemeralSystem {
	system, err := ImportSystemWithConfigE(t, config, systemPath, extraPaths...)
	require.NoError(t, err)
	return system
}

// ImportSystemWithConfigE imports the system given by merging the base system
// given by `systemPath` and the extra files given by each of the `extraPaths`
// into the platform instance given by the config.
// Returns error on failure.
func ImportSystemWithConfigE(t *testing.T, config *Config, systemPath string, extraPaths ...string) (*EphemeralSystem, error) {

	var err error

	err = checkSystem(systemPath)
	if err != nil {
		return nil, err
	}

	// our imported system root will be at a temporary directory
	tempdir, cleanup := MakeTempDir()
	system := NewImportedSystem(config, tempdir)

	// the system paths that are gonna be merged into the temporary directory
	merge := make([]string, 0, 1+len(extraPaths))
	merge = append(merge, systemPath)
	merge = append(merge, extraPaths...)

	t.Log("Merging system folders...")
	err = MergeFolders(tempdir, merge...)
	if err != nil {
		cleanup()
		return nil, err
	}

	t.Log("Registering developer...")
	err = cbRegisterDeveloper(t, system)
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

	t.Logf("Import successful: %s", system.RemoteURL())
	return system, nil
}

// cbRegisterDeveloper registers a new developer in the system if it doesn't
// exists already.
func cbRegisterDeveloper(t *testing.T, system *EphemeralSystem) error {

	devClient := cb.NewDevClientWithAddrs(system.PlatformURL(), system.MessagingURL(), "", "")

	email := system.config.Developer.Email
	password := system.config.Developer.Password
	firstname := "cbtest"
	lastname := "cbtest"
	org := "ClearBlade, Inc."
	regkey := system.config.RegistrationKey

	_, err := devClient.RegisterDevUserWithKey(email, password, firstname, lastname, org, regkey)
	if err != nil && !strings.Contains(err.Error(), "already exists") {
		return err
	}

	return nil
}

// cbImportSystem imports the given system into a remote platform instance. Note
// that this function will modify the passed system and set its system key and
// secret.
// Returns stdout/stderr and error on failure.
func cbImportSystem(t *testing.T, system *EphemeralSystem) (string, error) {

	importConfig := cbImportConfig(t, system)

	devClient, err := LoginAsDevE(t, system)
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
	system, err := UseSystemE(t)
	require.NoError(t, err)
	return system
}

// UseSystemE uses the external (not managed by cbtest) system given
// by the config flag. External systems are never destroyed automatically.
// Returns error on failure.
func UseSystemE(t *testing.T) (*EphemeralSystem, error) {

	config, err := ObtainConfig()
	if err != nil {
		return nil, err
	}

	return UseSystemWithConfigE(t, config)
}

// UseSystemWithConfig uses the external (not managed by cbtest) system given
// by the config. External systems are never destroyed automatically.
// Panics on error.
func UseSystemWithConfig(t *testing.T, config *Config) *EphemeralSystem {
	system, err := UseSystemWithConfigE(t, config)
	require.NoError(t, err)
	return system
}

// UseSystemWithConfigE uses the external (not managed by cbtest) system given
// by the config. External systems are never destroyed automatically.
// Returns error on failure.
func UseSystemWithConfigE(t *testing.T, config *Config) (*EphemeralSystem, error) {
	return NewExternalSystem(config), nil
}

// LoginAsDev logs into the system as a Developer (given by config).
// Panics on failure.
func LoginAsDev(t *testing.T, system *EphemeralSystem) *cb.DevClient {
	devClient, err := LoginAsDevE(t, system)
	require.NoError(t, err)
	return devClient
}

// LoginAsDevE logs into the System as a Developer (given by config).
// Returns error on failure.
func LoginAsDevE(t *testing.T, system *EphemeralSystem) (*cb.DevClient, error) {
	return doLoginAsDev(system)
}

func doLoginAsDev(system *EphemeralSystem) (*cb.DevClient, error) {

	config := system.config

	if !config.HasDeveloper() {
		return nil, fmt.Errorf("config does not have developer information")
	}

	devClient := cb.NewDevClientWithAddrs(config.PlatformURL, config.MessagingURL, config.Developer.Email, config.Developer.Password)
	_, err := devClient.Authenticate()
	if err != nil {
		return nil, err
	}

	return devClient, nil
}

// LoginAsUser logs into the system as a User.
// Panics on failure.
func LoginAsUser(t *testing.T, system *EphemeralSystem) *cb.UserClient {
	userClient, err := LoginAsUserE(t, system)
	require.NoError(t, err)
	return userClient
}

// LoginAsUserE logs into the system as a User.
// Returns error on failure.
func LoginAsUserE(t *testing.T, system *EphemeralSystem) (*cb.UserClient, error) {
	return doLoginAsUser(system)
}

func doLoginAsUser(system *EphemeralSystem) (*cb.UserClient, error) {

	config := system.config

	if !config.HasUser() {
		return nil, fmt.Errorf("config does not have user information")
	}

	userClient := cb.NewUserClientWithAddrs(config.PlatformURL, config.MessagingURL, config.SystemKey, config.SystemSecret, config.User.Email, config.User.Password)
	_, err := userClient.Authenticate()
	if err != nil {
		return nil, err
	}

	return userClient, nil

}
