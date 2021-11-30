package handler

import (
	"context"
	"errors"
	"github.com/maiaaraujo5/udp-chat/internal/app/server/domain/model"
	"github.com/maiaaraujo5/udp-chat/internal/app/server/domain/service/mocks"
	"github.com/maiaaraujo5/udp-chat/internal/app/server/server/udp/model/in"
	"github.com/maiaaraujo5/udp-chat/pkg/util"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"net"
	"reflect"
	"testing"
)

type ServerSuite struct {
	suite.Suite
	conn *net.UDPConn
}

func TestServerSuite(t *testing.T) {
	suite.Run(t, new(ServerSuite))
}

func (s *ServerSuite) SetupSuite() {
	s.conn = util.CreateUdpServer()
}

func (s *ServerSuite) TearDownSuite() {
	err := s.conn.Close()
	if err != nil {
		panic(err)
	}
}

func (s *ServerSuite) TestNewServer() {
	type args struct {
		messageRecover *mocks.Recover
		saver          *mocks.Saver
		deleteMessage  *mocks.DeleteMessage
		flusher        *mocks.Flusher
		conn           *net.UDPConn
	}
	tests := []struct {
		name string
		args args
		want *Server
	}{
		{
			name: "should build NewServer successfully",
			args: args{
				messageRecover: new(mocks.Recover),
				saver:          new(mocks.Saver),
				deleteMessage:  new(mocks.DeleteMessage),
				flusher:        new(mocks.Flusher),
				conn:           s.conn,
			},
			want: &Server{
				connections:    make(map[string]net.Addr),
				messageRecover: new(mocks.Recover),
				saveMessage:    new(mocks.Saver),
				deleteMessage:  new(mocks.DeleteMessage),
				flusher:        new(mocks.Flusher),
				conn:           s.conn,
			},
		},
	}
	for _, tt := range tests {

		s.Run(tt.name, func() {
			got := NewServer(tt.args.messageRecover, tt.args.saver, tt.args.deleteMessage, tt.args.flusher, tt.args.conn)
			s.Assert().True(reflect.DeepEqual(got, tt.want), "NewServer() = %v, want %v", got, tt.want)
		})
	}
}

func (s *ServerSuite) TestServer_handle() {
	type fields struct {
		connections    map[string]net.Addr
		messageRecover *mocks.Recover
		saveMessage    *mocks.Saver
		deleteMessage  *mocks.DeleteMessage
		flusher        *mocks.Flusher
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
		mock    func(recover *mocks.Recover, deleter *mocks.DeleteMessage, saver *mocks.Saver, flusher *mocks.Flusher)
	}{
		{
			name: "should successfully handle a new connection",
			fields: fields{
				connections:    make(map[string]net.Addr),
				messageRecover: new(mocks.Recover),
				saveMessage:    new(mocks.Saver),
				deleteMessage:  new(mocks.DeleteMessage),
				flusher:        new(mocks.Flusher),
			},
			args: args{
				parentCtx: context.Background(),
				in: &in.In{
					ID:      mock.Anything,
					Action:  "NEW_CONNECTION",
					Message: mock.Anything,
				},
				remote: &net.UDPAddr{
					IP:   net.IP{},
					Port: 3515,
				},
			},
			wantErr: false,
			mock: func(recover *mocks.Recover, deleter *mocks.DeleteMessage, saver *mocks.Saver, flusher *mocks.Flusher) {
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
			name: "should return error when service of handleNewConnection return errors",
			fields: fields{
				connections:    make(map[string]net.Addr),
				messageRecover: new(mocks.Recover),
				saveMessage:    new(mocks.Saver),
				deleteMessage:  new(mocks.DeleteMessage),
				flusher:        new(mocks.Flusher),
			},
			args: args{
				parentCtx: context.Background(),
				in: &in.In{
					ID:      mock.Anything,
					Action:  "NEW_CONNECTION",
					Message: mock.Anything,
				},
				remote: &net.UDPAddr{
					IP:   net.IP{},
					Port: 3515,
				},
			},
			wantErr: true,
			mock: func(recover *mocks.Recover, deleter *mocks.DeleteMessage, saver *mocks.Saver, flusher *mocks.Flusher) {
				recover.On("Execute", mock.Anything).Return(nil, errors.New("error to recover messages")).Once()
			},
		},
		{
			name: "should successfully handle a new message",
			fields: fields{
				connections:    make(map[string]net.Addr),
				messageRecover: new(mocks.Recover),
				saveMessage:    new(mocks.Saver),
				deleteMessage:  new(mocks.DeleteMessage),
				flusher:        new(mocks.Flusher),
			},
			args: args{
				parentCtx: context.Background(),
				in: &in.In{
					ID:      mock.Anything,
					Action:  "NEW_MESSAGE",
					Message: mock.Anything,
				},
				remote: &net.UDPAddr{
					IP:   net.IP{},
					Port: 3515,
				},
			},
			wantErr: false,
			mock: func(recover *mocks.Recover, deleter *mocks.DeleteMessage, saver *mocks.Saver, flusher *mocks.Flusher) {
				saver.On("Execute", mock.Anything, mock.Anything).Return(nil).Once()
			},
		},
		{
			name: "should return error when service of handleNewMessage returns error",
			fields: fields{
				connections:    make(map[string]net.Addr),
				messageRecover: new(mocks.Recover),
				saveMessage:    new(mocks.Saver),
				deleteMessage:  new(mocks.DeleteMessage),
				flusher:        new(mocks.Flusher),
			},
			args: args{
				parentCtx: context.Background(),
				in: &in.In{
					ID:      mock.Anything,
					Action:  "NEW_MESSAGE",
					Message: mock.Anything,
				},
				remote: &net.UDPAddr{
					IP:   net.IP{},
					Port: 3515,
				},
			},
			wantErr: true,
			mock: func(recover *mocks.Recover, deleter *mocks.DeleteMessage, saver *mocks.Saver, flusher *mocks.Flusher) {
				saver.On("Execute", mock.Anything, mock.Anything).Return(errors.New("error to save message")).Once()
			},
		},
		{
			name: "should successfully handle a message delete",
			fields: fields{
				connections:    make(map[string]net.Addr),
				messageRecover: new(mocks.Recover),
				saveMessage:    new(mocks.Saver),
				deleteMessage:  new(mocks.DeleteMessage),
				flusher:        new(mocks.Flusher),
			},
			args: args{
				parentCtx: context.Background(),
				in: &in.In{
					ID:      mock.Anything,
					Action:  "DELETE_MESSAGE",
					Message: mock.Anything,
				},
				remote: &net.UDPAddr{
					IP:   net.IP{},
					Port: 3515,
				},
			},
			wantErr: false,
			mock: func(recover *mocks.Recover, deleter *mocks.DeleteMessage, saver *mocks.Saver, flusher *mocks.Flusher) {
				deleter.On("Execute", mock.Anything, mock.Anything).Return(nil).Once()
			},
		},
		{
			name: "should return error when service of handleDeleteMessage returns error",
			fields: fields{
				connections:    make(map[string]net.Addr),
				messageRecover: new(mocks.Recover),
				saveMessage:    new(mocks.Saver),
				deleteMessage:  new(mocks.DeleteMessage),
				flusher:        new(mocks.Flusher),
			},
			args: args{
				parentCtx: context.Background(),
				in: &in.In{
					ID:      mock.Anything,
					Action:  "DELETE_MESSAGE",
					Message: mock.Anything,
				},
				remote: &net.UDPAddr{
					IP:   net.IP{},
					Port: 3515,
				},
			},
			wantErr: true,
			mock: func(recover *mocks.Recover, deleter *mocks.DeleteMessage, saver *mocks.Saver, flusher *mocks.Flusher) {
				deleter.On("Execute", mock.Anything, mock.Anything).Return(errors.New("error to delete message")).Once()
			},
		},
		{
			name: "should successfully handle a disconnection",
			fields: fields{
				connections: map[string]net.Addr{
					":3515": &net.UDPAddr{IP: net.IP{}, Port: 3515},
				},
				messageRecover: new(mocks.Recover),
				saveMessage:    new(mocks.Saver),
				deleteMessage:  new(mocks.DeleteMessage),
				flusher:        new(mocks.Flusher),
			},
			args: args{
				parentCtx: context.Background(),
				in: &in.In{
					ID:      mock.Anything,
					Action:  "DISCONNECT",
					Message: mock.Anything,
				},
				remote: &net.UDPAddr{
					IP:   net.IP{},
					Port: 3515,
				},
			},
			wantErr: false,
			mock: func(recover *mocks.Recover, deleter *mocks.DeleteMessage, saver *mocks.Saver, flusher *mocks.Flusher) {
				flusher.On("Execute", mock.Anything).Return(nil).Once()
			},
		},
		{
			name: "should return error when handleDisconnection returns error",
			fields: fields{
				connections: map[string]net.Addr{
					":3515": &net.UDPAddr{IP: net.IP{}, Port: 3515},
				},
				messageRecover: new(mocks.Recover),
				saveMessage:    new(mocks.Saver),
				deleteMessage:  new(mocks.DeleteMessage),
				flusher:        new(mocks.Flusher),
			},
			args: args{
				parentCtx: context.Background(),
				in: &in.In{
					ID:      mock.Anything,
					Action:  "DISCONNECT",
					Message: mock.Anything,
				},
				remote: &net.UDPAddr{
					IP:   net.IP{},
					Port: 3515,
				},
			},
			wantErr: true,
			mock: func(recover *mocks.Recover, deleter *mocks.DeleteMessage, saver *mocks.Saver, flusher *mocks.Flusher) {
				flusher.On("Execute", mock.Anything).Return(errors.New("error to flush")).Once()
			},
		},
	}
	for _, tt := range tests {
		tt.mock(tt.fields.messageRecover, tt.fields.deleteMessage, tt.fields.saveMessage, tt.fields.flusher)

		s.Run(tt.name, func() {
			r := &Server{
				connections:    tt.fields.connections,
				messageRecover: tt.fields.messageRecover,
				saveMessage:    tt.fields.saveMessage,
				deleteMessage:  tt.fields.deleteMessage,
				flusher:        tt.fields.flusher,
				conn:           s.conn,
			}
			err := r.handle(tt.args.parentCtx, tt.args.in, tt.args.remote)
			s.Assert().True((err != nil) == tt.wantErr, "handle() error = %v, wantErr %v", err, tt.wantErr)
		})

		tt.fields.deleteMessage.AssertExpectations(s.T())
		tt.fields.flusher.AssertExpectations(s.T())
		tt.fields.saveMessage.AssertExpectations(s.T())
		tt.fields.messageRecover.AssertExpectations(s.T())
	}
}
