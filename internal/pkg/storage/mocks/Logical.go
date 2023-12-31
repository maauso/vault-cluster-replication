// Code generated by mockery v2.32.4. DO NOT EDIT.

package mocks

import (
	api "github.com/hashicorp/vault/api"
	mock "github.com/stretchr/testify/mock"
)

// Logical is an autogenerated mock type for the Logical type
type Logical struct {
	mock.Mock
}

// Delete provides a mock function with given fields: path
func (_m *Logical) Delete(path string) (*api.Secret, error) {
	ret := _m.Called(path)

	var r0 *api.Secret
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*api.Secret, error)); ok {
		return rf(path)
	}
	if rf, ok := ret.Get(0).(func(string) *api.Secret); ok {
		r0 = rf(path)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*api.Secret)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: path
func (_m *Logical) List(path string) (*api.Secret, error) {
	ret := _m.Called(path)

	var r0 *api.Secret
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*api.Secret, error)); ok {
		return rf(path)
	}
	if rf, ok := ret.Get(0).(func(string) *api.Secret); ok {
		r0 = rf(path)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*api.Secret)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Read provides a mock function with given fields: path
func (_m *Logical) Read(path string) (*api.Secret, error) {
	ret := _m.Called(path)

	var r0 *api.Secret
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*api.Secret, error)); ok {
		return rf(path)
	}
	if rf, ok := ret.Get(0).(func(string) *api.Secret); ok {
		r0 = rf(path)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*api.Secret)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Write provides a mock function with given fields: path, data
func (_m *Logical) Write(path string, data map[string]interface{}) (*api.Secret, error) {
	ret := _m.Called(path, data)

	var r0 *api.Secret
	var r1 error
	if rf, ok := ret.Get(0).(func(string, map[string]interface{}) (*api.Secret, error)); ok {
		return rf(path, data)
	}
	if rf, ok := ret.Get(0).(func(string, map[string]interface{}) *api.Secret); ok {
		r0 = rf(path, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*api.Secret)
		}
	}

	if rf, ok := ret.Get(1).(func(string, map[string]interface{}) error); ok {
		r1 = rf(path, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewLogical creates a new instance of Logical. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewLogical(t interface {
	mock.TestingT
	Cleanup(func())
}) *Logical {
	mock := &Logical{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
