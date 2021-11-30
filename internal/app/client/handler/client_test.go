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
	conn *net.UDPConn
}

func TestClientSuite(t *testing.T) {
	suite.Run(t, new(ClientSuite))
}

func (s *ClientSuite) SetupSuite() {
	s.conn = util.CreateUdpServer()
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
				conn: s.conn,
			},
			want: &Client{
				conn: s.conn,
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
