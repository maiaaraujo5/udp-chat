package handler

import (
	"github.com/maiaaraujo5/udp-chat/internal/app/client/domain/model/in"
	"github.com/maiaaraujo5/udp-chat/internal/app/client/domain/model/out"
	"github.com/maiaaraujo5/udp-chat/internal/app/client/domain/service/mocks"
	"github.com/stretchr/testify/mock"
	"net"
	"reflect"
)

func (s *ClientSuite) TestClient_handleSendMessage() {
	type fields struct {
		conn    *net.UDPConn
		creator *mocks.Creator
	}
	type args struct {
		action  string
		message string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		want    []in.In
		mock    func(creator *mocks.Creator)
	}{
		{
			name: "should send one message successfully",
			fields: fields{
				conn:    s.client,
				creator: new(mocks.Creator),
			},
			args: args{
				action:  "NEW_MESSAGE",
				message: mock.Anything,
			},
			wantErr: false,
			want: []in.In{
				{
					ID:      mock.Anything,
					UserID:  s.client.LocalAddr().String(),
					Message: mock.Anything,
				},
			},
			mock: func(creator *mocks.Creator) {
				creator.On("Create", mock.Anything, mock.Anything).Return(&out.Out{
					ID:      mock.Anything,
					Action:  mock.Anything,
					Message: mock.Anything,
				}).Once()
			},
		},
		{
			name: "should successfully send message but not append in history when action is DELETE_MESSAGE",
			fields: fields{
				conn:    s.client,
				creator: new(mocks.Creator),
			},
			args: args{
				action:  "DELETE_MESSAGE",
				message: mock.Anything,
			},
			wantErr: false,
			want:    nil,
			mock: func(creator *mocks.Creator) {
				creator.On("Create", mock.Anything, mock.Anything).Return(&out.Out{
					ID:      mock.Anything,
					Action:  mock.Anything,
					Message: mock.Anything,
				}).Once()
			},
		},
		{
			name: "should successfully send message but not append in history when action is NEW_CONNECTION",
			fields: fields{
				conn:    s.client,
				creator: new(mocks.Creator),
			},
			args: args{
				action:  "NEW_CONNECTION",
				message: mock.Anything,
			},
			wantErr: false,
			want:    nil,
			mock: func(creator *mocks.Creator) {
				creator.On("Create", mock.Anything, mock.Anything).Return(&out.Out{
					ID:      mock.Anything,
					Action:  mock.Anything,
					Message: mock.Anything,
				}).Once()
			},
		},
		{
			name: "should successfully send message but not append in history when action is DISCONNECT",
			fields: fields{
				conn:    s.client,
				creator: new(mocks.Creator),
			},
			args: args{
				action:  "DISCONNECT",
				message: mock.Anything,
			},
			wantErr: false,
			want:    nil,
			mock: func(creator *mocks.Creator) {
				creator.On("Create", mock.Anything, mock.Anything).Return(&out.Out{
					ID:      mock.Anything,
					Action:  mock.Anything,
					Message: mock.Anything,
				}).Once()
			},
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {
			tt.mock(tt.fields.creator)

			r := &Client{
				conn:    tt.fields.conn,
				creator: tt.fields.creator,
			}

			err := r.handleSendMessage(tt.args.action, tt.args.message)
			s.Assert().True((err != nil) == tt.wantErr, "Execute() error = %v, wantErr %v", err, tt.wantErr)
			s.Assert().True(reflect.DeepEqual(r.messages, tt.want), "NewClient() = %v, want %v", r.messages, tt.want)

			tt.fields.creator.AssertExpectations(s.T())
		})
	}
}
