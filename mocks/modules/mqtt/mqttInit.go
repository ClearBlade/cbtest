// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	GoSDK "github.com/clearblade/Go-SDK"
	mock "github.com/stretchr/testify/mock"

	tls "crypto/tls"
)

// mqttInit is an autogenerated mock type for the mqttInit type
type mqttInit struct {
	mock.Mock
}

// InitializeMQTT provides a mock function with given fields: clientID, systemKey, timeout, ssl, lastWill
func (_m *mqttInit) InitializeMQTT(clientID string, systemKey string, timeout int, ssl *tls.Config, lastWill *GoSDK.LastWillPacket) error {
	ret := _m.Called(clientID, systemKey, timeout, ssl, lastWill)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, int, *tls.Config, *GoSDK.LastWillPacket) error); ok {
		r0 = rf(clientID, systemKey, timeout, ssl, lastWill)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}