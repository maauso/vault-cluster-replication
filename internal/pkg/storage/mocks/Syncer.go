// Code generated by mockery v2.32.4. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Syncer is an autogenerated mock type for the Syncer type
type Syncer struct {
	mock.Mock
}

// PullSnapshot provides a mock function with given fields:
func (_m *Syncer) PullSnapshot() (string, error) {
	ret := _m.Called()

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func() (string, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PushSnapshot provides a mock function with given fields: fileName
func (_m *Syncer) PushSnapshot(fileName string) error {
	ret := _m.Called(fileName)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(fileName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewSyncer creates a new instance of Syncer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSyncer(t interface {
	mock.TestingT
	Cleanup(func())
}) *Syncer {
	mock := &Syncer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}