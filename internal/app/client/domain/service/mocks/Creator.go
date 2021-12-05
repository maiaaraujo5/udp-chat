// Code generated by mockery 2.9.0. DO NOT EDIT.

package mocks

import (
	out "github.com/maiaaraujo5/udp-chat/internal/app/client/domain/model/out"
	mock "github.com/stretchr/testify/mock"
)

// Creator is an autogenerated mock type for the Creator type
type Creator struct {
	mock.Mock
}

// Create provides a mock function with given fields: action, message
func (_m *Creator) Create(action string, message string) *out.Out {
	ret := _m.Called(action, message)

	var r0 *out.Out
	if rf, ok := ret.Get(0).(func(string, string) *out.Out); ok {
		r0 = rf(action, message)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*out.Out)
		}
	}

	return r0
}