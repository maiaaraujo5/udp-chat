package handler

import (
	"context"
	"errors"
	"github.com/maiaaraujo5/udp-chat/internal/app/server/domain/service/mocks"
	"github.com/maiaaraujo5/udp-chat/internal/app/server/server/udp/model/in"
	"github.com/stretchr/testify/mock"
	"net"
)

func (s *ServerSuite) TestServer_handleDeleteMessage() {
	type fields struct {
		connections   map[string]net.Addr
		deleteMessage *mocks.DeleteMessage
	}
	type args struct {
		parentCtx context.Context
		in        *in.In
		remote    *net.UDPAddr
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		mock    func(service *mocks.DeleteMessage)
	}{
		{
			name: "should successfully delete one message",
			fields: fields{
				deleteMessage: new(mocks.DeleteMessage),
				connections: map[string]net.Addr{
					":3515": &net.UDPAddr{IP: net.IP{}, Port: 3515},
					":3516": &net.UDPAddr{IP: net.IP{}, Port: 3516},
				},
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
			mock: func(service *mocks.DeleteMessage) {
				service.On("Execute", mock.Anything, mock.Anything).Return(nil).Once()
			},
		},
		{
			name: "should return error when service return errors",
			fields: fields{
				deleteMessage: new(mocks.DeleteMessage),
			},
			args: args{
				parentCtx: context.Background(),
				in: &in.In{
					ID:      mock.Anything,
					Action:  mock.Anything,
					Message: mock.Anything,
				},
			},
			wantErr: true,
			mock: func(service *mocks.DeleteMessage) {
				service.On("Execute", mock.Anything, mock.Anything).Return(errors.New("error to delete message"))
			},
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {
			tt.mock(tt.fields.deleteMessage)

			r := &Server{
				connections:   tt.fields.connections,
				deleteMessage: tt.fields.deleteMessage,
				conn:          s.conn,
			}

			err := r.handleDeleteMessage(tt.args.parentCtx, tt.args.in, tt.args.remote)
			s.Assert().True((err != nil) == tt.wantErr, "Execute() error = %v, wantErr %v", err, tt.wantErr)

			tt.fields.deleteMessage.AssertExpectations(s.T())
		})
	}
}
