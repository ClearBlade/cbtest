package mqtt

import (
	"crypto/tls"
	"testing"

	cb "github.com/clearblade/Go-SDK"
	"github.com/clearblade/cbtest/config"
	"github.com/stretchr/testify/assert"
)

type mqttClientMock struct {
	clientID  string
	systemKey string
	timeout   int
	ssl       *tls.Config
	lastWill  *cb.LastWillPacket
}

func (mc *mqttClientMock) InitializeMQTT(clientID, systemKey string, timeout int, ssl *tls.Config, lastWill *cb.LastWillPacket) error {
	mc.clientID = clientID
	mc.systemKey = systemKey
	mc.timeout = timeout
	mc.ssl = ssl
	mc.lastWill = lastWill
	return nil
}

func TestInitializeMQTTUsesProviderSucceeds(t *testing.T) {
	config := config.GetDefaultConfig()
	config.SystemKey = "system-key-override"
	clientMock := mqttClientMock{}

	InitializeMQTT(t, config, &clientMock)

	assert.NotEmpty(t, clientMock.clientID)
	assert.Equal(t, "system-key-override", clientMock.systemKey)
	assert.Greater(t, clientMock.timeout, 0)
	assert.Nil(t, clientMock.ssl)
	assert.Nil(t, clientMock.lastWill)
}
