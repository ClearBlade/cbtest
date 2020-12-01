package system

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/clearblade/cblib"
	"github.com/clearblade/cbtest"
	"github.com/clearblade/cbtest/internal/fsutil"
	"github.com/clearblade/cbtest/internal/merge"
	"github.com/clearblade/cbtest/modules/auth"
	"github.com/clearblade/cbtest/provider"
)

//go:generate mockery --name importSteps --inpackage --testonly

// importSteps defines a set of functions that we can use for importing a system.
type importSteps interface {
	Path() string
	Extra() []string
	CheckSystem(path string) error
	MakeTempDir() (string, func())
	MergeFolders(dest string, srcs ...string) error
	RegisterDeveloper(t cbtest.T, provider provider.Config) error
	DoImport(t cbtest.T, provider provider.Config, localPath string) (systemKey string, systemSecret string, err error)
	Cleanup(tempdir string)
}

type defaultImportSteps struct {
	path  string
	extra []string
}

func (s *defaultImportSteps) Path() string {
	return s.path
}

func (s *defaultImportSteps) Extra() []string {
	return s.extra
}

func (s *defaultImportSteps) CheckSystem(path string) error {
	return cbCheckSystem(path)
}

func (s *defaultImportSteps) MakeTempDir() (string, func()) {
	return fsutil.MakeTempDir()
}

func (s *defaultImportSteps) MergeFolders(dest string, srcs ...string) error {
	return merge.Folders(dest, srcs...)
}

func (s *defaultImportSteps) RegisterDeveloper(t cbtest.T, provider provider.Config) error {
	return cbRegisterDeveloper(t, provider)
}

func (s *defaultImportSteps) DoImport(t cbtest.T, provider provider.Config, localPath string) (string, string, error) {
	return cbImportSystem(t, provider, localPath)
}

func (s *defaultImportSteps) Cleanup(tempdir string) {
	os.RemoveAll(tempdir)
}

// cbCheckSystem returns error if the given path does not contain a system.
func cbCheckSystem(path string) error {

	systemJSONPath := filepath.Join(path, "system.json")

	_, err := os.Stat(systemJSONPath)
	if err != nil {
		return fmt.Errorf("not a system: %s", err)
	}

	return nil
}

// cbRegisterDeveloper uses the developer information in the config to register
// a developer in the system also given by the config.
func cbRegisterDeveloper(t cbtest.T, provider provider.Config) error {
	config := provider.Config(t)
	email := config.Developer.Email
	password := config.Developer.Password
	return auth.RegisterDevE(t, provider, email, password)
}

// cbImportSystem imports the given system into a remote platform instance.
// Returns system key, system secret, and error on failure.
func cbImportSystem(t cbtest.T, provider provider.Config, localPath string) (string, string, error) {
	t.Helper()

	importConfig := cbImportConfig(t, provider)

	devClient, err := auth.LoginAsDevE(t, provider)
	if err != nil {
		return "", "", err
	}

	result, err := cblib.ImportSystemUsingConfig(importConfig, localPath, devClient)
	if err != nil {
		return "", "", err
	}

	return result.SystemKey, result.SystemSecret, nil
}

// cbImportConfig returns a cblib.ImportConfig instance for importing the system.
func cbImportConfig(t cbtest.T, provider provider.Config) cblib.ImportConfig {
	t.Helper()

	config := provider.Config(t)
	name := fmt.Sprintf("cbtest-%s", t.Name())
	nowstr := time.Now().UTC().Format(time.UnixDate)

	return cblib.ImportConfig{
		SystemName:             name,
		SystemDescription:      fmt.Sprintf("Created on %s", nowstr),
		ImportUsers:            config.Import.ImportUsers,
		ImportRows:             config.Import.ImportRows,
		DefaultUserPassword:    config.User.Password,
		DefaultDeviceActiveKey: config.Device.ActiveKey,
	}
}
