package handler

import (
	"github.com/maiaaraujo5/udp-chat/pkg/util"
	"github.com/stretchr/testify/suite"
	"net"
	"reflect"
	"testing"
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
	s.server = util.CreateUdpServer()
	s.client = util.CreateUdpClient()
}

func (s *ClientSuite) TestNewClient() {
	type args struct {
		conn *net.UDPConn
	}
	tests := []struct {
		name string
		args args
		want *Client
	}{
		{
			name: "should build NewClient successfully",
			args: args{
				conn: s.client,
			},
			want: &Client{
				conn: s.client,
			},
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {
			got := NewClient(tt.args.conn)
			s.Assert().True(reflect.DeepEqual(got, tt.want), "NewClient() = %v, want %v", got, tt.want)
		})
	}
}
