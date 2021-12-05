package handler

import (
	"errors"
	"net"
	"os"
	"reflect"
	"testing"

	"bou.ke/monkey"
	"github.com/maiaaraujo5/udp-chat/internal/app/client/domain/model/in"
	"github.com/maiaaraujo5/udp-chat/internal/app/client/domain/model/out"
	"github.com/maiaaraujo5/udp-chat/internal/app/client/domain/service/mocks"
	"github.com/maiaaraujo5/udp-chat/pkg/util"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type ClientSuite struct {
	suite.Suite
	server *net.UDPConn
	client *net.UDPConn
}

func TestClientSuite(t *testing.T) {
	suite.Run(t, new(ClientSuite))
}

func (s *ClientSuite) SetupSuite() {
	s.server = util.CreateUDPServer()
	s.client = util.CreateUDPClient()

	monkey.Patch(os.Exit, func(i int) {})
}

func (s *ClientSuite) TestNewClient() {
	type args struct {
		conn     *net.UDPConn
		receiver *mocks.Receiver
		deleter  *mocks.Deleter
		creator  *mocks.Creator
	}
	tests := []struct {
		name string
		args args
		want *Client
	}{
		{
			name: "should build NewClient successfully",
			args: args{
				conn:     s.client,
				receiver: new(mocks.Receiver),
				deleter:  new(mocks.Deleter),
				creator:  new(mocks.Creator),
			},
			want: &Client{
				conn:     s.client,
				receiver: new(mocks.Receiver),
				deleter:  new(mocks.Deleter),
				creator:  new(mocks.Creator),
			},
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {
			got := NewClient(tt.args.conn, tt.args.receiver, tt.args.deleter, tt.args.creator)
			s.Assert().True(reflect.DeepEqual(got, tt.want), "NewClient() = %v, want %v", got, tt.want)
		})
	}
}

func (s *ClientSuite) TestClient_handleCommands() {
	type fields struct {
		conn     *net.UDPConn
		receiver *mocks.Receiver
		deleter  *mocks.Deleter
		creator  *mocks.Creator
		messages []in.In
	}
	type args struct {
		command string
		msg     string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []in.In
		wantErr bool
		mock    func(receiver *mocks.Receiver, deleter *mocks.Deleter, creator *mocks.Creator)
	}{
		{
			name: "should successfully send a new message",
			fields: fields{
				conn:     s.client,
				receiver: new(mocks.Receiver),
				deleter:  new(mocks.Deleter),
				creator:  new(mocks.Creator),
			},
			args: args{
				command: "/msg",
				msg:     mock.Anything,
			},
			want: []in.In{
				{
					ID:      mock.Anything,
					UserID:  s.client.LocalAddr().String(),
					Message: mock.Anything,
				},
			},
			wantErr: false,
			mock: func(receiver *mocks.Receiver, deleter *mocks.Deleter, creator *mocks.Creator) {
				creator.On("Create", mock.Anything, mock.Anything).Return(&out.Out{
					ID:      mock.Anything,
					Action:  mock.Anything,
					Message: mock.Anything,
				})
			},
		},
		{
			name: "should successfully delete one message",
			fields: fields{
				conn:     s.client,
				receiver: new(mocks.Receiver),
				deleter:  new(mocks.Deleter),
				creator:  new(mocks.Creator),
				messages: []in.In{
					{
						ID:      mock.Anything,
						UserID:  s.client.LocalAddr().String(),
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
				command: "/del",
				msg:     mock.Anything,
			},
			want: []in.In{
				{
					ID:      mock.Anything,
					UserID:  mock.Anything,
					Message: mock.Anything,
				},
			},
			wantErr: false,
			mock: func(receiver *mocks.Receiver, deleter *mocks.Deleter, creator *mocks.Creator) {

				deleter.On("Delete", mock.Anything, mock.Anything, mock.Anything).Return([]in.In{
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
			name: "should not return error when handleDeleteMessage returns error and don't remove the message from history",
			fields: fields{
				conn:     s.client,
				receiver: new(mocks.Receiver),
				deleter:  new(mocks.Deleter),
				creator:  new(mocks.Creator),
				messages: []in.In{
					{
						ID:      mock.Anything,
						UserID:  s.client.LocalAddr().String(),
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
				command: "/del",
				msg:     mock.Anything,
			},
			want: []in.In{
				{
					ID:      mock.Anything,
					UserID:  s.client.LocalAddr().String(),
					Message: mock.Anything,
				},
				{
					ID:      mock.Anything,
					UserID:  mock.Anything,
					Message: mock.Anything,
				},
			},
			wantErr: false,
			mock: func(receiver *mocks.Receiver, deleter *mocks.Deleter, creator *mocks.Creator) {
				deleter.On("Delete", mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("error")).Once()
			},
		},
		{
			name: "should successfully disconnect from chat",
			fields: fields{
				conn:     s.client,
				receiver: new(mocks.Receiver),
				deleter:  new(mocks.Deleter),
				creator:  new(mocks.Creator),
			},
			args: args{
				command: "/quit",
				msg:     mock.Anything,
			},
			wantErr: false,
			mock: func(receiver *mocks.Receiver, deleter *mocks.Deleter, creator *mocks.Creator) {
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
			tt.mock(tt.fields.receiver, tt.fields.deleter, tt.fields.creator)

			r := &Client{
				conn:     tt.fields.conn,
				receiver: tt.fields.receiver,
				deleter:  tt.fields.deleter,
				creator:  tt.fields.creator,
				messages: tt.fields.messages,
			}

			r.handleCommands(tt.args.command, tt.args.msg)

			s.Assert().True(reflect.DeepEqual(r.messages, tt.want), "NewClient() = %v, want %v", r.messages, tt.want)

			tt.fields.receiver.AssertExpectations(s.T())
			tt.fields.deleter.AssertExpectations(s.T())
			tt.fields.creator.AssertExpectations(s.T())
		})
	}
}
