package service

import (
	"github.com/maiaaraujo5/udp-chat/internal/app/client/domain/model/in"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"reflect"
	"testing"
)

type ReceiverSuite struct {
	suite.Suite
}

func TestReceiverSuite(t *testing.T) {
	suite.Run(t, new(ReceiverSuite))
}

func (s *ReceiverSuite) TestNewReceiver() {
	tests := []struct {
		name string
		want Receiver
	}{
		{
			name: "should successfully build NewReceiver",
			want: &receiver{},
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {
			got := NewReceiver()
			s.Assert().True(reflect.DeepEqual(got, tt.want), "NewReceiver() = %v, want %v", got, tt.want)
		})
	}
}

func (s *ReceiverSuite) Test_receiver_Receive() {
	type args struct {
		messages []in.In
		message  *in.In
	}
	tests := []struct {
		name string
		args args
		want []in.In
	}{
		{
			name: "should successfully append new message in history",
			args: args{
				message: &in.In{
					ID:      mock.Anything,
					UserID:  mock.Anything,
					Message: mock.Anything,
				},
				messages: []in.In{
					{
						ID:      mock.Anything,
						UserID:  mock.Anything,
						Message: mock.Anything,
					},
					{
						ID:      mock.Anything,
						UserID:  mock.Anything,
						Message: mock.Anything,
					},
				},
			},
			want: []in.In{
				{
					ID:      mock.Anything,
					UserID:  mock.Anything,
					Message: mock.Anything,
				},
				{
					ID:      mock.Anything,
					UserID:  mock.Anything,
					Message: mock.Anything,
				},
				{
					ID:      mock.Anything,
					UserID:  mock.Anything,
					Message: mock.Anything,
				},
			},
		},
		{
			name: "should remove a message when message was deleted",
			args: args{
				message: &in.In{
					ID:     "02",
					UserID: "user_id",
				},
				messages: []in.In{
					{
						ID:      mock.Anything,
						UserID:  mock.Anything,
						Message: mock.Anything,
					},
					{
						ID:      "02",
						UserID:  "user_id",
						Message: mock.Anything,
					},
					{
						ID:      mock.Anything,
						UserID:  mock.Anything,
						Message: mock.Anything,
					},
				},
			},
			want: []in.In{
				{
					ID:      mock.Anything,
					UserID:  mock.Anything,
					Message: mock.Anything,
				},
				{
					ID:      mock.Anything,
					UserID:  mock.Anything,
					Message: mock.Anything,
				},
			},
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {
			r := &receiver{}
			got := r.Receive(tt.args.messages, tt.args.message)
			s.Assert().True(reflect.DeepEqual(got, tt.want), "Receive() = %v, want %v", got, tt.want)
		})
	}
}
