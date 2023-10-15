// Code generated by mockery v2.32.4. DO NOT EDIT.

package mocks

import (
	io "io"

	mock "github.com/stretchr/testify/mock"
)

// Sys is an autogenerated mock type for the Sys type
type Sys struct {
	mock.Mock
}

// RaftSnapshot provides a mock function with given fields: snapWriter
func (_m *Sys) RaftSnapshot(snapWriter io.Writer) error {
	ret := _m.Called(snapWriter)

	var r0 error
	if rf, ok := ret.Get(0).(func(io.Writer) error); ok {
		r0 = rf(snapWriter)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RaftSnapshotRestore provides a mock function with given fields: snapReader, force
func (_m *Sys) RaftSnapshotRestore(snapReader io.Reader, force bool) error {
	ret := _m.Called(snapReader, force)

	var r0 error
	if rf, ok := ret.Get(0).(func(io.Reader, bool) error); ok {
		r0 = rf(snapReader, force)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewSys creates a new instance of Sys. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSys(t interface {
	mock.TestingT
	Cleanup(func())
}) *Sys {
	mock := &Sys{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
