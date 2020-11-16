package system

import (
	"os"

	"github.com/stretchr/testify/require"
	"golang.org/x/sync/errgroup"

	"github.com/clearblade/cbtest"
	"github.com/clearblade/cbtest/modules/auth"
)

// Destroy destroys the remote system instance, as well as the local folder. Only
// systems that are managed by this module will be destroyed, external systems
// will not.
// Panics on failure.
func Destroy(t cbtest.T, system *EphemeralSystem) {
	t.Helper()
	err := DestroyE(t, system)
	require.NoError(t, err)
}

// DestroyE destroys the remote system instance, as well as the local folder. Only
// systems that are managed by this module will be destroyed, external systems
// will not.
// Returns error on failure.
func DestroyE(t cbtest.T, system *EphemeralSystem) error {
	t.Helper()

	if system.IsExternal() {
		t.Log("External system not destroyed")
		return nil
	}

	err := doDestroy(t, system)
	if err != nil {
		t.Logf("Error destroying system: %s", err)
		return err
	}

	t.Logf("System was destroyed")
	return nil
}

func doDestroy(t cbtest.T, system *EphemeralSystem) error {

	g := errgroup.Group{}

	// remove local folder
	g.Go(func() error {
		return os.RemoveAll(system.LocalPath())
	})

	// remove remote system
	g.Go(func() error {
		devClient, err := auth.LoginAsDevE(t, system)
		if err != nil {
			return err
		}
		return devClient.DeleteSystem(system.SystemKey())
	})

	return g.Wait()
}
