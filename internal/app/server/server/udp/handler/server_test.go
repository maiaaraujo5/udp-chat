package handler

import (
	"github.com/maiaaraujo5/udp-chat/internal/app/server/domain/service"
	"github.com/maiaaraujo5/udp-chat/internal/app/server/server/udp/handler/util"
	"github.com/stretchr/testify/suite"
	"net"
	"reflect"
	"testing"
)

type ServerSuite struct {
	suite.Suite
	conn net.PacketConn
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

func TestNewServer(t *testing.T) {
	type args struct {
		messageRecover service.Recover
		saver          service.Saver
		deleteMessage  service.DeleteMessage
		flusher        service.Flusher
		conn           *net.UDPConn
	}
	tests := []struct {
		name string
		args args
		want *Server
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewServer(tt.args.messageRecover, tt.args.saver, tt.args.deleteMessage, tt.args.flusher, tt.args.conn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewServer() = %v, want %v", got, tt.want)
			}
		})
	}
}
