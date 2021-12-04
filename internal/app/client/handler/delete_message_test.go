package handler

import (
	"github.com/maiaaraujo5/udp-chat/internal/app/client/handler/model/in"
	"github.com/maiaaraujo5/udp-chat/pkg/util"
	"github.com/stretchr/testify/mock"
	"net"
)

func (s *ClientSuite) TestClient_handleDeleteMessage() {
	type fields struct {
		conn     *net.UDPConn
		messages []in.In
	}
	type args struct {
		msg string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "should successfully remove a message from history of messages",
			fields: fields{
				conn: util.CreateUdpServer(),
				messages: []in.In{
					{
						ID:      "1847",
						UserID:  "1",
						Message: mock.Anything,
					},
					{
						ID:      "1848",
						UserID:  s.client.LocalAddr().String(),
						Message: mock.Anything,
					},
					{
						ID:      "1849",
						UserID:  "3",
						Message: mock.Anything,
					},
				},
			},
			args: args{
				msg: "1848",
			},
			wantErr: false,
		},
		{
			name: "should not remove a message from history of messages when client is not the owner of the message",
			fields: fields{
				conn: util.CreateUdpServer(),
				messages: []in.In{
					{
						ID:      "1847",
						UserID:  "1",
						Message: mock.Anything,
					},
					{
						ID:      "1848",
						UserID:  s.client.LocalAddr().String(),
						Message: mock.Anything,
					},
					{
						ID:      "1849",
						UserID:  "3",
						Message: mock.Anything,
					},
				},
			},
			args: args{
				msg: "1847",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {
			r := &Client{
				conn:     s.client,
				messages: tt.fields.messages,
			}

			err := r.handleDeleteMessage(tt.args.msg)

			s.Assert().True((err != nil) == tt.wantErr, "Execute() error = %v, wantErr %v", err, tt.wantErr)
		})
	}
}
