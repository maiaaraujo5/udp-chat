package handler

import (
	"errors"
	"github.com/maiaaraujo5/udp-chat/internal/app/client/domain/model/in"
	"github.com/maiaaraujo5/udp-chat/internal/app/client/domain/model/out"
	"github.com/maiaaraujo5/udp-chat/internal/app/client/domain/service/mocks"
	"github.com/maiaaraujo5/udp-chat/pkg/util"
	"github.com/stretchr/testify/mock"
	"net"
)

func (s *ClientSuite) TestClient_handleDeleteMessage() {
	type fields struct {
		conn     *net.UDPConn
		messages []in.In
		deleter  *mocks.Deleter
		creator  *mocks.Creator
	}
	type args struct {
		msg string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		mock    func(deleter *mocks.Deleter, creator *mocks.Creator)
	}{
		{
			name: "should successfully remove a message from history of messages",
			fields: fields{
				conn:    util.CreateUdpServer(),
				deleter: new(mocks.Deleter),
				creator: new(mocks.Creator),
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
					{
						ID:      mock.Anything,
						UserID:  mock.Anything,
						Message: mock.Anything,
					},
				},
			},
			args: args{
				msg: mock.Anything,
			},
			wantErr: false,
			mock: func(deleter *mocks.Deleter, creator *mocks.Creator) {
				deleter.On("Delete", mock.Anything, mock.Anything, mock.Anything).Return([]in.In{
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
				}, nil).Once()

				creator.On("Create", mock.Anything, mock.Anything).Return(&out.Out{
					ID:      mock.Anything,
					Action:  mock.Anything,
					Message: mock.Anything,
				}).Once()
			},
		},
		{
			name: "should not return error when deleter returns error and not send message to server to delete message",
			fields: fields{
				conn:    util.CreateUdpServer(),
				deleter: new(mocks.Deleter),
				creator: new(mocks.Creator),
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
					{
						ID:      mock.Anything,
						UserID:  mock.Anything,
						Message: mock.Anything,
					},
				},
			},
			args: args{
				msg: mock.Anything,
			},
			wantErr: true,
			mock: func(deleter *mocks.Deleter, creator *mocks.Creator) {
				deleter.On("Delete", mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("error")).Once()
			},
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {

			tt.mock(tt.fields.deleter, tt.fields.creator)

			r := &Client{
				conn:     s.client,
				messages: tt.fields.messages,
				deleter:  tt.fields.deleter,
				creator:  tt.fields.creator,
			}

			err := r.handleDeleteMessage(tt.args.msg)

			s.Assert().True((err != nil) == tt.wantErr, "handleDeleteMessage() error = %v, wantErr %v", err, tt.wantErr)

			tt.fields.deleter.AssertExpectations(s.T())
		})
	}
}
