// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	cbtest "github.com/clearblade/cbtest"
	config "github.com/clearblade/cbtest/config"

	mock "github.com/stretchr/testify/mock"
)

// Config is an autogenerated mock type for the Config type
type Config struct {
	mock.Mock
}

// Config provides a mock function with given fields: _a0
func (_m *Config) Config(_a0 cbtest.T) *config.Config {
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
