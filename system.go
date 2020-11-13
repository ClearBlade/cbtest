package cbtest

import (
	"fmt"
	"net/url"
	"os"
)

// EphemeralSystem represents a system that exists solely for testing.
type EphemeralSystem struct {
	config    *Config
	localPath string
}

// NewImportedSystem creates a new *EphemeralSystem from an imported system.
func NewImportedSystem(name string, config *Config, localPath string) *EphemeralSystem {
	return &EphemeralSystem{
		config:    config,
		localPath: localPath,
	}
}

// NewExternalSystem creates anew *EphemeralSystem from an external system.
func NewExternalSystem(config *Config) *EphemeralSystem {
	return &EphemeralSystem{
		config:    config,
		localPath: "<external>",
	}
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
	return es.localPath
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

// Destroy destroys the remote system instance, as well as the local folder.
func (es *EphemeralSystem) Destroy() error {
	err := os.RemoveAll(es.localPath)
	return err
}
