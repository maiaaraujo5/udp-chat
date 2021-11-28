package handler

import (
	"context"
	"errors"
	"github.com/maiaaraujo5/udp-chat/internal/app/server/domain/model"
	"github.com/maiaaraujo5/udp-chat/internal/app/server/domain/service/mocks"
	"github.com/maiaaraujo5/udp-chat/internal/app/server/server/udp/model/in"
	"github.com/stretchr/testify/mock"
	"net"
)

func (s *ServerSuite) TestServer_handleNewConnection() {
	type fields struct {
		connections    map[string]net.Addr
		messageRecover *mocks.Recover
		saveMessage    *mocks.Saver
	}
	type args struct {
		parentCtx context.Context
		in        *in.In
		remote    net.Addr
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		mock    func(recover *mocks.Recover, saver *mocks.Saver)
	}{
		{
			name: "should handle new connection successfully and send old messages to new client",
			fields: fields{
				connections:    map[string]net.Addr{},
				messageRecover: new(mocks.Recover),
				saveMessage:    new(mocks.Saver),
			},
			args: args{
				parentCtx: context.Background(),
				in: &in.In{
					ID:      mock.Anything,
					Action:  mock.Anything,
					Message: mock.Anything,
				},
				remote: &net.UDPAddr{
					IP:   net.IP{},
					Port: 3515,
				},
			},
			wantErr: false,
			mock: func(recover *mocks.Recover, saver *mocks.Saver) {
				recover.On("Execute", mock.Anything).Return([]model.Message{
					{
						ID:      mock.Anything,
						UserID:  mock.Anything,
						Message: mock.Anything,
					},
				}, nil).Once()

				saver.On("Execute", mock.Anything, mock.Anything).Return(nil).Once()
			},
		},
		{
			name: "should return error when recover return error",
			fields: fields{
				connections:    map[string]net.Addr{},
				messageRecover: new(mocks.Recover),
				saveMessage:    new(mocks.Saver),
			},
			args: args{
				parentCtx: context.Background(),
				in: &in.In{
					ID:      mock.Anything,
					Action:  mock.Anything,
					Message: mock.Anything,
				},
				remote: &net.UDPAddr{
					IP:   net.IP{},
					Port: 3515,
				},
			},
			wantErr: true,
			mock: func(recover *mocks.Recover, saver *mocks.Saver) {
				recover.On("Execute", mock.Anything).Return(nil, errors.New("error to recover messages")).Once()
			},
		},
		{
			name: "should return error when saver return error",
			fields: fields{
				connections:    map[string]net.Addr{},
				messageRecover: new(mocks.Recover),
				saveMessage:    new(mocks.Saver),
			},
			args: args{
				parentCtx: context.Background(),
				in: &in.In{
					ID:      mock.Anything,
					Action:  mock.Anything,
					Message: mock.Anything,
				},
				remote: &net.UDPAddr{
					IP:   net.IP{},
					Port: 3515,
				},
			},
			wantErr: true,
			mock: func(recover *mocks.Recover, saver *mocks.Saver) {
				recover.On("Execute", mock.Anything).Return([]model.Message{
					{
						ID:      mock.Anything,
						UserID:  mock.Anything,
						Message: mock.Anything,
					},
				}, nil).Once()

				saver.On("Execute", mock.Anything, mock.Anything).Return(errors.New("error to save message")).Once()
			},
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {
			tt.mock(tt.fields.messageRecover, tt.fields.saveMessage)
			r := &Server{
				connections:    tt.fields.connections,
				messageRecover: tt.fields.messageRecover,
				saveMessage:    tt.fields.saveMessage,
				conn:           s.conn,
			}

			err := r.handleNewConnection(tt.args.parentCtx, tt.args.in, tt.args.remote)
			s.Assert().True((err != nil) == tt.wantErr, "Execute() error = %v, wantErr %v", err, tt.wantErr)

			tt.fields.messageRecover.AssertExpectations(s.T())
			tt.fields.saveMessage.AssertExpectations(s.T())
		})
	}
}
