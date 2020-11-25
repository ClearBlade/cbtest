package system

import (
	"github.com/clearblade/cbtest"
	"github.com/clearblade/cbtest/config"
	"github.com/stretchr/testify/require"
)

// Finish closes the given system. Real behavior depends on the config and the
// flags passed to cbtest. For instance, and external system will never be destroyed;
// and an imported system will always be destroyed unless we requested its configuration
// to be saved for later use.
// Panics on error.
func Finish(t cbtest.T, s *EphemeralSystem) {
	t.Helper()
	err := FinishE(t, s)
	require.NoError(t, err)
}

// FinishE closes the given system. Real behavior depends on the config and the
// flags passed to cbtest. For instance, and external system will never be destroyed;
// and an imported system will always be destroyed unless we requested its configuration
// to be saved for later use.
// Returns error on failure.
func FinishE(t cbtest.T, s *EphemeralSystem) error {

	if s.IsExternal() {
		t.Log("Finish: closing external system")
		return nil

	} else if config.HasConfigOut() {
		t.Log("Finish: saving system")
		return config.SaveConfig(t, s.config)

	} else {
		t.Log("Finish: destroying system")
		return DestroyE(t, s)
	}
}
