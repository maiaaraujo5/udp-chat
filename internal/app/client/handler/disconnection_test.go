package handler

import (
	in "github.com/maiaaraujo5/udp-chat/internal/app/client/domain/model/in"
	"github.com/maiaaraujo5/udp-chat/internal/app/client/domain/model/out"
	"github.com/maiaaraujo5/udp-chat/internal/app/client/domain/service/mocks"
	"github.com/stretchr/testify/mock"
	"net"
)

func (s *ClientSuite) TestClient_handleDisconnection() {
	type fields struct {
		conn     *net.UDPConn
		creator  *mocks.Creator
		messages []in.In
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
		mock    func(creator *mocks.Creator)
	}{
		{
			name: "should successfully handle disconnection",
			fields: fields{
				conn:    s.client,
				creator: new(mocks.Creator),
			},
			wantErr: false,
			mock: func(creator *mocks.Creator) {
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
			tt.mock(tt.fields.creator)
			r := &Client{
				conn:    tt.fields.conn,
				creator: tt.fields.creator,
			}

			err := r.handleDisconnection()
			s.Assert().True((err != nil) == tt.wantErr, "handleDeleteMessage() error = %v, wantErr %v", err, tt.wantErr)

			tt.fields.creator.AssertExpectations(s.T())
		})
	}
}
