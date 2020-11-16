package system

import (
	"github.com/clearblade/cbtest"
	"github.com/clearblade/cbtest/config"
	"github.com/stretchr/testify/require"
)

// Save saves the system config for later use with UseSystem.
// Panics on failure.
func Save(t cbtest.T, system *EphemeralSystem) {
	err := SaveE(t, system)
	require.NoError(t, err)
}

// SaveE saves the system config for later use with UseSystem.
// Returns error on failure.
func SaveE(t cbtest.T, system *EphemeralSystem) error {
	return config.SaveConfig(t, system.config)
}
