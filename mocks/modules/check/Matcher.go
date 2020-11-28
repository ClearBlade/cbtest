// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Matcher is an autogenerated mock type for the Matcher type
type Matcher struct {
	mock.Mock
}

// FailureMessage provides a mock function with given fields: actual
func (_m *Matcher) FailureMessage(actual interface{}) string {
	ret := _m.Called(actual)

	var r0 string
	if rf, ok := ret.Get(0).(func(interface{}) string); ok {
		r0 = rf(actual)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Match provides a mock function with given fields: actual
func (_m *Matcher) Match(actual interface{}) (bool, error) {
	ret := _m.Called(actual)

	var r0 bool
	if rf, ok := ret.Get(0).(func(interface{}) bool); ok {
		r0 = rf(actual)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}) error); ok {
		r1 = rf(actual)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NegatedFailureMessage provides a mock function with given fields: actual
func (_m *Matcher) NegatedFailureMessage(actual interface{}) string {
	ret := _m.Called(actual)

	var r0 string
	if rf, ok := ret.Get(0).(func(interface{}) string); ok {
		r0 = rf(actual)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}