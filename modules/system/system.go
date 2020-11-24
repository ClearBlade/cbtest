// Package system contains functions using, importing, saving, and destroying
// systems. Each tests usually consists of use/create and close/destroy operations,
// this module makes that easier.
//
// Using existing systems
//
// Either Use or UseOrImport (when system info is provided in the config or the
// flags) will return a reference to an existing system.
//
// Importing into new system
//
// Either Import or UseOrImport (when no system info is provided) will import and
// return a reference to the new system.
//
// Destroying systems
//
// Destroy will destroy the system when called, therefore, is mostly useful when
// used together with a defer statement:
//
//     defer system.Destroy(t, s)
//
// Saving systems
//
// Save will create write a config file with the system info in it for later use:
//
//     defer system.Save(t, s)
//
// Polymorphic destroying and saving
//
// To avoid having to change the test code each time we want to destroy and/dor
// save a system, the Close function will take care of deciding what to do based
// on the provided configuration and flags:
//
//     defer system.Close(t, s)
//
package system

import (
	"fmt"
	"net/url"
	"sync"

	cb "github.com/clearblade/Go-SDK"
	"github.com/clearblade/cbtest/config"
)

// EphemeralSystem represents a system that exists solely for testing.
type EphemeralSystem struct {
	config     *config.Config
	localPath  string
	isExternal bool
	client     *cb.DevClient
	clientLock sync.Mutex
}

// NewImportedSystem creates a new *EphemeralSystem from an imported system.
func NewImportedSystem(config *config.Config, localPath string) *EphemeralSystem {
	return &EphemeralSystem{
		config:     config,
		localPath:  localPath,
		isExternal: false,
	}
}

// NewExternalSystem creates anew *EphemeralSystem from an external system.
func NewExternalSystem(config *config.Config) *EphemeralSystem {
	return &EphemeralSystem{
		config:     config,
		localPath:  "",
		isExternal: true,
	}
}

// Provide returns the config that was used for creating this system. This
// function implements the config.Provider interface.
func (es *EphemeralSystem) Provide() *config.Config {
	return es.config
}

// PlatformURL returns the platform url that hosts this system.
func (es *EphemeralSystem) PlatformURL() string {
	return es.config.PlatformURL
}

// MessagingURL returns the messaging url that hosts this system.
func (es *EphemeralSystem) MessagingURL() string {
	return es.config.MessagingURL
}

// SystemKey returns the system key corresponding to this ephemeral system.
func (es *EphemeralSystem) SystemKey() string {
	return es.config.SystemKey
}

// SystemSecret returns the system secret corresponding to this ephemeral system.
func (es *EphemeralSystem) SystemSecret() string {
	return es.config.SystemSecret
}

// LocalPath returns the path where the local system is located.
func (es *EphemeralSystem) LocalPath() string {
	if es.IsExternal() {
		return "<external>"
	}
	return es.localPath
}

// IsExternal returns if the ephemeral system is an external one.
func (es *EphemeralSystem) IsExternal() bool {
	return es.isExternal
}

// RemoteURL returns the URL where the remote system is running.
func (es *EphemeralSystem) RemoteURL() string {
	rawurl := fmt.Sprintf("%s/console/system/%s/detail", es.config.PlatformURL, es.config.SystemKey)
	url, err := url.Parse(rawurl)
	if err != nil {
		return "<bad remote url>"
	}
	return url.String()
}
