package handler

import (
	"github.com/maiaaraujo5/udp-chat/internal/app/client/domain/service/mocks"
	"net"
	"reflect"

	"github.com/maiaaraujo5/udp-chat/internal/app/client/domain/model/in"
	"github.com/stretchr/testify/mock"
)

func (s *ClientSuite) TestClient_receiveMessages() {
	type fields struct {
		conn     *net.UDPConn
		receiver *mocks.Receiver
		messages []in.In
	}
	type args struct {
		msg *in.In
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []in.In
		mock   func(receiver *mocks.Receiver)
	}{
		{
			name: "should successfully receive a new message",
			fields: fields{
				conn:     s.client,
				receiver: new(mocks.Receiver),
			},
			args: args{
				msg: &in.In{
					ID:      mock.Anything,
					UserID:  mock.Anything,
					Message: mock.Anything,
				},
			},
			want: []in.In{
				{
					ID:      mock.Anything,
					UserID:  mock.Anything,
					Message: mock.Anything,
				},
			},
			mock: func(receiver *mocks.Receiver) {
				receiver.On("Receive", mock.Anything, mock.Anything).Return([]in.In{
					{
						ID:      mock.Anything,
						UserID:  mock.Anything,
						Message: mock.Anything,
					},
				})
			},
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {

			tt.mock(tt.fields.receiver)

			r := &Client{
				conn:     tt.fields.conn,
				messages: tt.fields.messages,
				receiver: tt.fields.receiver,
			}

			r.receiveMessage(tt.args.msg)
			s.Assert().True(reflect.DeepEqual(r.messages, tt.want), "NewClient() = %v, want %v", r.messages, tt.want)

			tt.fields.receiver.AssertExpectations(s.T())
		})
	}
}
