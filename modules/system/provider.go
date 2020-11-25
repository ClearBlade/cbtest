package system

import (
	cb "github.com/clearblade/Go-SDK"
	"github.com/clearblade/cbtest"
	"github.com/clearblade/cbtest/config"
	"github.com/clearblade/cbtest/modules/auth"
	"github.com/stretchr/testify/require"
)

// Config returns the config that was used for creating this system. This
// function implements the provider.Config interface.
// Panics on failure.
func (es *EphemeralSystem) Config(t cbtest.T) *config.Config {
	return es.config.Config(t)
}

// Client returns a dev client that can be used for accessing this system. This
// function implements the provider.Client interface.
// Panics on failure.
func (es *EphemeralSystem) Client(t cbtest.T) *cb.DevClient {
	client, err := es.ClientE(t)
	require.NoError(t, err)
	return client
}

// ClientE returns a dev client that can be used for accessing this system. This
// function implements the provider.Client interface.
// Returns error on failure.
func (es *EphemeralSystem) ClientE(t cbtest.T) (*cb.DevClient, error) {

	// lock client access
	es.clientLock.Lock()
	defer es.clientLock.Unlock()

	// if a client already exists, return it
	if es.client != nil {
		return es.client, nil
	}

	// login and get new client
	devClient, err := auth.LoginAsDevE(t, es.config)
	if err != nil {
		return nil, err
	}

	// cache client
	es.client = devClient
	return devClient, nil
}
