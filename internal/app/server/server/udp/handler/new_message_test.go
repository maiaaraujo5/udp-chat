package handler

import (
	"context"
	"errors"
	"github.com/maiaaraujo5/udp-chat/internal/app/server/domain/service/mocks"
	"github.com/maiaaraujo5/udp-chat/internal/app/server/server/udp/model/in"
	"github.com/stretchr/testify/mock"
	"net"
)

func (s *ServerSuite) TestServer_handleNewMessage() {
	type fields struct {
		connections map[string]net.Addr
		saveMessage *mocks.Saver
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
		mock    func(saver *mocks.Saver)
	}{
		{
			name: "should successfully handle new message",
			fields: fields{
				connections: map[string]net.Addr{
					":3515": &net.UDPAddr{IP: net.IP{}, Port: 3515},
					":3516": &net.UDPAddr{IP: net.IP{}, Port: 3516},
				},
				saveMessage: new(mocks.Saver),
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
			mock: func(saver *mocks.Saver) {
				saver.On("Save", mock.Anything, mock.Anything).Return(nil).Once()
			},
		},
		{
			name: "should return error when saver returns error",
			fields: fields{
				connections: map[string]net.Addr{
					":3515": &net.UDPAddr{IP: net.IP{}, Port: 3515},
					":3516": &net.UDPAddr{IP: net.IP{}, Port: 3516},
				},
				saveMessage: new(mocks.Saver),
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
			mock: func(saver *mocks.Saver) {
				saver.On("Save", mock.Anything, mock.Anything).Return(errors.New("error to save message")).Once()
			},
		},
	}
	for _, tt := range tests {
		tt.mock(tt.fields.saveMessage)

		s.Run(tt.name, func() {
			r := &Server{
				connections: tt.fields.connections,
				saveMessage: tt.fields.saveMessage,
				conn:        s.conn,
			}

			err := r.handleNewMessage(tt.args.parentCtx, tt.args.in, tt.args.remote)
			s.Assert().True((err != nil) == tt.wantErr, "handleNewMessage() error = %v, wantErr %v", err, tt.wantErr)

			tt.fields.saveMessage.AssertExpectations(s.T())
		})
	}
}
