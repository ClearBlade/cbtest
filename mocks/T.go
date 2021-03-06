// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// T is an autogenerated mock type for the T type
type T struct {
	mock.Mock
}

// Error provides a mock function with given fields: args
func (_m *T) Error(args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Errorf provides a mock function with given fields: format, args
func (_m *T) Errorf(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Fail provides a mock function with given fields:
func (_m *T) Fail() {
	_m.Called()
}

// FailNow provides a mock function with given fields:
func (_m *T) FailNow() {
	_m.Called()
}

// Helper provides a mock function with given fields:
func (_m *T) Helper() {
	_m.Called()
}

// Log provides a mock function with given fields: args
func (_m *T) Log(args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Logf provides a mock function with given fields: format, args
func (_m *T) Logf(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Name provides a mock function with given fields:
func (_m *T) Name() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}
