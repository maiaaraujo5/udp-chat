// Code generated by mockery 2.9.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// Flusher is an autogenerated mock type for the Flusher type
type Flusher struct {
	mock.Mock
}

// Execute provides a mock function with given fields: parentCtx
func (_m *Flusher) Execute(parentCtx context.Context) error {
	ret := _m.Called(parentCtx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(parentCtx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
