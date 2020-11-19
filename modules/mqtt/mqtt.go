package mqtt

import (
	"crypto/rand"
	"crypto/tls"
	"encoding/hex"
	"fmt"

	cb "github.com/clearblade/Go-SDK"
	"github.com/clearblade/cbtest"
	"github.com/clearblade/cbtest/provider"
	"github.com/stretchr/testify/require"
)

const (
	mqttTimeoutSeconds = 20
)

// mqttInit interface defines a function for initialing the MQTT client. Clients
// in the ClearBlade SDK already implement this interface.
type mqttInit interface {
	InitializeMQTT(clientID string, systemKey string, timeout int, ssl *tls.Config, lastWill *cb.LastWillPacket) error
}

// InitializeMQTT initializes MQTT for the given client.
// Panics on failure.
func InitializeMQTT(t cbtest.T, provider provider.Config, client mqttInit) {
	t.Helper()
	err := InitializeMQTTE(t, provider, client)
	require.NoError(t, err)
}

// InitializeMQTTE initializes MQTT for the given client.
// Returns error on failure.
func InitializeMQTTE(t cbtest.T, provider provider.Config, client mqttInit) error {
	t.Helper()
	err := cbInitializeMQTT(t, provider, client)
	return err
}

// cbInitializeMQTT initializes the given mqttInit interface.
func cbInitializeMQTT(t cbtest.T, provider provider.Config, mqtt mqttInit) error {

	randomString, err := generateRandomString()
	if err != nil {
		return err
	}

	config := provider.Config(t)
	clientID := fmt.Sprintf("cbtest%s", randomString)
	err = mqtt.InitializeMQTT(clientID, config.SystemKey, mqttTimeoutSeconds, nil, nil)
	if err != nil {
		return err
	}

	return nil
}

// generateRandomString generates a random 16-byte value encoded as a
// hexadecimal string. See: https://stackoverflow.com/a/47677306
func generateRandomString() (string, error) {
	data := make([]byte, 16)
	_, err := rand.Read(data)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(data), nil
}
