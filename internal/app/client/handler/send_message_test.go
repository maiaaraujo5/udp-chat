package handler

import (
	"github.com/stretchr/testify/mock"
	"net"
)

func (s *ClientSuite) TestClient_handleSendMessage() {
	type fields struct {
		conn *net.UDPConn
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
	}{
		{
			name: "should send one message successfully",
			fields: fields{
				conn: s.client,
			},
			args: args{
				action:  "NEW_MESSAGE",
				message: mock.Anything,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {
			r := &Client{
				conn: tt.fields.conn,
			}

			err := r.handleSendMessage(tt.args.action, tt.args.message)
			s.Assert().True((err != nil) == tt.wantErr, "Execute() error = %v, wantErr %v", err, tt.wantErr)
		})
	}
}
