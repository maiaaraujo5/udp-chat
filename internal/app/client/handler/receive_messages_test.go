package handler

import (
	"github.com/maiaaraujo5/udp-chat/internal/app/client/handler/model/in"
	"github.com/stretchr/testify/mock"
	"net"
	"reflect"
)

func (s *ClientSuite) TestClient_receiveMessages() {
	type fields struct {
		conn     *net.UDPConn
		messages []in.In
	}
	type args struct {
		msg *in.In
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []in.In
	}{
		{
			name: "should successfully receive one new message and append in the history when history is nil",
			fields: fields{
				conn: s.client,
			},
			args: args{
				msg: &in.In{
					ID:      mock.Anything,
					UserID:  mock.Anything,
					Message: mock.Anything,
				},
			},
			want: []in.In{
				{
					ID:      mock.Anything,
					UserID:  mock.Anything,
					Message: mock.Anything,
				},
			},
		},
		{
			name: "should successfully receive one new message and append in the history when history is not empty",
			fields: fields{
				conn: s.client,
				messages: []in.In{
					{
						ID:      "1",
						UserID:  mock.Anything,
						Message: mock.Anything,
					},
				},
			},
			args: args{
				msg: &in.In{
					ID:      "2",
					UserID:  mock.Anything,
					Message: mock.Anything,
				},
			},
			want: []in.In{
				{
					ID:      "1",
					UserID:  mock.Anything,
					Message: mock.Anything,
				},
				{
					ID:      "2",
					UserID:  mock.Anything,
					Message: mock.Anything,
				},
			},
		},
		{
			name: "should remove message from history when receives only the id of the message",
			fields: fields{
				conn: s.client,
				messages: []in.In{
					{
						ID:      "1",
						UserID:  mock.Anything,
						Message: mock.Anything,
					},
					{
						ID:      "2",
						UserID:  mock.Anything,
						Message: mock.Anything,
					},
					{
						ID:      "3",
						UserID:  mock.Anything,
						Message: mock.Anything,
					},
				},
			},
			args: args{
				msg: &in.In{
					ID: "2",
				},
			},
			want: []in.In{
				{
					ID:      "1",
					UserID:  mock.Anything,
					Message: mock.Anything,
				},
				{
					ID:      "3",
					UserID:  mock.Anything,
					Message: mock.Anything,
				},
			},
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {
			r := &Client{
				conn:     tt.fields.conn,
				messages: tt.fields.messages,
			}

			r.receiveMessages(tt.args.msg)
			s.Assert().True(reflect.DeepEqual(r.messages, tt.want), "NewClient() = %v, want %v", r.messages, tt.want)
		})
	}
}
