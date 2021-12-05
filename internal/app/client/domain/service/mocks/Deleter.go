// Code generated by mockery 2.9.0. DO NOT EDIT.

package mocks

import (
	in "github.com/maiaaraujo5/udp-chat/internal/app/client/domain/model/in"
	mock "github.com/stretchr/testify/mock"
)

// Deleter is an autogenerated mock type for the Deleter type
type Deleter struct {
	mock.Mock
}

// Delete provides a mock function with given fields: messages, messageID, userID
func (_m *Deleter) Delete(messages []in.In, messageID string, userID string) ([]in.In, error) {
	ret := _m.Called(messages, messageID, userID)

	var r0 []in.In
	if rf, ok := ret.Get(0).(func([]in.In, string, string) []in.In); ok {
		r0 = rf(messages, messageID, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]in.In)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]in.In, string, string) error); ok {
		r1 = rf(messages, messageID, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}