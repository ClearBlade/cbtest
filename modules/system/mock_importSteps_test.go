// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package system

import (
	cbtest "github.com/clearblade/cbtest"
	mock "github.com/stretchr/testify/mock"

	provider "github.com/clearblade/cbtest/provider"
)

// mockImportSteps is an autogenerated mock type for the importSteps type
type mockImportSteps struct {
	mock.Mock
}

// CheckSystem provides a mock function with given fields: path
func (_m *mockImportSteps) CheckSystem(path string) error {
	ret := _m.Called(path)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(path)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Cleanup provides a mock function with given fields: tempdir
func (_m *mockImportSteps) Cleanup(tempdir string) {
	_m.Called(tempdir)
}

// DoImport provides a mock function with given fields: t, _a1, localPath
func (_m *mockImportSteps) DoImport(t cbtest.T, _a1 provider.Config, localPath string) (string, string, error) {
	ret := _m.Called(t, _a1, localPath)

	var r0 string
	if rf, ok := ret.Get(0).(func(cbtest.T, provider.Config, string) string); ok {
		r0 = rf(t, _a1, localPath)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(cbtest.T, provider.Config, string) string); ok {
		r1 = rf(t, _a1, localPath)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(cbtest.T, provider.Config, string) error); ok {
		r2 = rf(t, _a1, localPath)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Extra provides a mock function with given fields:
func (_m *mockImportSteps) Extra() []string {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// MakeTempDir provides a mock function with given fields:
func (_m *mockImportSteps) MakeTempDir() (string, func()) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 func()
	if rf, ok := ret.Get(1).(func() func()); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(func())
		}
	}

	return r0, r1
}

// MergeFolders provides a mock function with given fields: dest, srcs
func (_m *mockImportSteps) MergeFolders(dest string, srcs ...string) error {
	_va := make([]interface{}, len(srcs))
	for _i := range srcs {
		_va[_i] = srcs[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, dest)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, ...string) error); ok {
		r0 = rf(dest, srcs...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Path provides a mock function with given fields:
func (_m *mockImportSteps) Path() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// RegisterDeveloper provides a mock function with given fields: t, _a1
func (_m *mockImportSteps) RegisterDeveloper(t cbtest.T, _a1 provider.Config) error {
	ret := _m.Called(t, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(cbtest.T, provider.Config) error); ok {
		r0 = rf(t, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
