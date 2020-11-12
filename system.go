package cbtest

import (
	"os"

	"github.com/clearblade/cbtest/types"
)

// EphemeralSystem represents a system that exists solely for testing.
type EphemeralSystem struct {
	localPath string
	System    *types.System
}

// LocalPath returns the path where the local system is located.
func (es *EphemeralSystem) LocalPath() string {
	return es.localPath
}

// RemoteURL returns the URL where the remote system is running.
func (es *EphemeralSystem) RemoteURL() string {
	return ""
}

// Destroy destroyes the remote system instance, as well as the local folder.
func (es *EphemeralSystem) Destroy() error {
	err := os.RemoveAll(es.localPath)
	return err
}
