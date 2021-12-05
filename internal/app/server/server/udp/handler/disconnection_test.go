package handler

import (
	"context"
	"errors"
	"github.com/maiaaraujo5/udp-chat/internal/app/server/domain/service/mocks"
	"github.com/stretchr/testify/mock"
	"net"
)

func (s *ServerSuite) TestServer_handleDisconnection() {
	type fields struct {
		flusher     *mocks.Flusher
		saver       *mocks.Saver
		connections map[string]net.Addr
	}
	type args struct {
		parentCtx context.Context
		remote    *net.UDPAddr
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		mock    func(flusher *mocks.Flusher, saver *mocks.Saver)
	}{
		{
			name: "should successfully handle disconnection when there is more than one client connected",
			fields: fields{
				flusher: new(mocks.Flusher),
				saver:   new(mocks.Saver),
				connections: map[string]net.Addr{
					":3515": &net.UDPAddr{IP: net.IP{}, Port: 3515},
					":3516": &net.UDPAddr{IP: net.IP{}, Port: 3516},
				},
			},
			args: args{
				parentCtx: context.Background(),
				remote: &net.UDPAddr{
					IP:   net.IP{},
					Port: 3515,
				},
			},
			wantErr: false,
			mock: func(flusher *mocks.Flusher, saver *mocks.Saver) {
				saver.On("Save", mock.Anything, mock.Anything).Return(nil).Once()
			},
		},
		{
			name: "should successfully handle disconnection and call flusher when just have only one active connection",
			fields: fields{
				flusher: new(mocks.Flusher),
				saver:   new(mocks.Saver),
				connections: map[string]net.Addr{
					":3515": &net.UDPAddr{IP: net.IP{}, Port: 3515},
				},
			},
			args: args{
				parentCtx: context.Background(),
				remote: &net.UDPAddr{
					IP:   net.IP{},
					Port: 3515,
				},
			},
			wantErr: false,
			mock: func(flusher *mocks.Flusher, saver *mocks.Saver) {
				flusher.On("Flush", mock.Anything).Return(nil).Once()
			},
		},
		{
			name: "should return error when save message returns error",
			fields: fields{
				flusher: new(mocks.Flusher),
				saver:   new(mocks.Saver),
				connections: map[string]net.Addr{
					":3515": &net.UDPAddr{IP: net.IP{}, Port: 3515},
					":3516": &net.UDPAddr{IP: net.IP{}, Port: 3515},
				},
			},
			args: args{
				parentCtx: context.Background(),
				remote: &net.UDPAddr{
					IP:   net.IP{},
					Port: 3515,
				},
			},
			wantErr: true,
			mock: func(flusher *mocks.Flusher, saver *mocks.Saver) {
				saver.On("Save", mock.Anything, mock.Anything).Return(errors.New("error to save message")).Once()
			},
		},
		{
			name: "should return error when flusher return error",
			fields: fields{
				flusher: new(mocks.Flusher),
				saver:   new(mocks.Saver),
				connections: map[string]net.Addr{
					":3515": &net.UDPAddr{IP: net.IP{}, Port: 3515},
				},
			},
			args: args{
				parentCtx: context.Background(),
				remote: &net.UDPAddr{
					IP:   net.IP{},
					Port: 3515,
				},
			},
			wantErr: true,
			mock: func(flusher *mocks.Flusher, saver *mocks.Saver) {
				flusher.On("Flush", mock.Anything).Return(errors.New("error to flusher")).Once()
			},
		},
	}
	for _, tt := range tests {

		s.Run(tt.name, func() {
			tt.mock(tt.fields.flusher, tt.fields.saver)

			r := &Server{
				flusher:     tt.fields.flusher,
				saveMessage: tt.fields.saver,
				connections: tt.fields.connections,
				conn:        s.conn,
			}

			err := r.handleDisconnection(tt.args.parentCtx, tt.args.remote)
			s.Assert().True((err != nil) == tt.wantErr, "handleDisconnection() error = %v, wantErr %v", err, tt.wantErr)

			tt.fields.flusher.AssertExpectations(s.T())
			tt.fields.saver.AssertExpectations(s.T())
		})
	}
}
