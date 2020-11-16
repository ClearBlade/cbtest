package system

import (
	"fmt"
	"net/url"

	"github.com/clearblade/cbtest/config"
)

// EphemeralSystem represents a system that exists solely for testing.
type EphemeralSystem struct {
	config     *config.Config
	localPath  string
	isExternal bool
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
