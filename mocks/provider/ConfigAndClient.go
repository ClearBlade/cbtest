// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	GoSDK "github.com/clearblade/Go-SDK"
	cbtest "github.com/clearblade/cbtest"

	config "github.com/clearblade/cbtest/config"

	mock "github.com/stretchr/testify/mock"
)

// ConfigAndClient is an autogenerated mock type for the ConfigAndClient type
type ConfigAndClient struct {
	mock.Mock
}

// Client provides a mock function with given fields: _a0
func (_m *ConfigAndClient) Client(_a0 cbtest.T) *GoSDK.DevClient {
	ret := _m.Called(_a0)

	var r0 *GoSDK.DevClient
	if rf, ok := ret.Get(0).(func(cbtest.T) *GoSDK.DevClient); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*GoSDK.DevClient)
		}
	}

	return r0
}

// ClientE provides a mock function with given fields: _a0
func (_m *ConfigAndClient) ClientE(_a0 cbtest.T) (*GoSDK.DevClient, error) {
	ret := _m.Called(_a0)

	var r0 *GoSDK.DevClient
	if rf, ok := ret.Get(0).(func(cbtest.T) *GoSDK.DevClient); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*GoSDK.DevClient)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(cbtest.T) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Config provides a mock function with given fields: _a0
func (_m *ConfigAndClient) Config(_a0 cbtest.T) *config.Config {
	ret := _m.Called(_a0)

	var r0 *config.Config
	if rf, ok := ret.Get(0).(func(cbtest.T) *config.Config); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*config.Config)
		}
	}

	return r0
}
