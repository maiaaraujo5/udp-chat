// Code generated by mockery 2.9.0. DO NOT EDIT.

package mocks

import (
	in "github.com/maiaaraujo5/udp-chat/internal/app/client/domain/model/in"
	mock "github.com/stretchr/testify/mock"
)

// Receiver is an autogenerated mock type for the Receiver type
type Receiver struct {
	mock.Mock
}

// Receive provides a mock function with given fields: messages, message
func (_m *Receiver) Receive(messages []in.In, message *in.In) []in.In {
	ret := _m.Called(messages, message)

	var r0 []in.In
	if rf, ok := ret.Get(0).(func([]in.In, *in.In) []in.In); ok {
		r0 = rf(messages, message)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]in.In)
		}
	}

	return r0
}
