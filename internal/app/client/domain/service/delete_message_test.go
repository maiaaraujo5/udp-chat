package service

import (
	"github.com/maiaaraujo5/udp-chat/internal/app/client/domain/model/in"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"reflect"
	"testing"
)

type DeleteMessageSuite struct {
	suite.Suite
}

func TestDeleteMessageSuite(t *testing.T) {
	suite.Run(t, new(DeleteMessageSuite))
}

func (s *DeleteMessageSuite) TestNewDeleter() {
	tests := []struct {
		name string
		want Deleter
	}{
		{
			name: "should successfully build NewDeleter",
			want: &deleter{},
		},
	}
	for _, tt := range tests {

		s.Run(tt.name, func() {
			got := NewDeleter()
			s.Assert().True(reflect.DeepEqual(got, tt.want), "NewDeleter() = %v, want %v", got, tt.want)
		})
	}
}

func (s *DeleteMessageSuite) Test_deleter_Delete() {
	type args struct {
		messages  []in.In
		messageID string
		userID    string
	}
	tests := []struct {
		name    string
		args    args
		want    []in.In
		wantErr bool
	}{
		{
			name: "should successfully delete a message when message was sent by user who is trying to delete",
			args: args{
				messageID: "02",
				userID:    "user_id",
				messages: []in.In{
					{
						ID:      "01",
						UserID:  mock.Anything,
						Message: mock.Anything,
					},
					{
						ID:      "02",
						UserID:  "user_id",
						Message: mock.Anything,
					},
					{
						ID:      "03",
						UserID:  mock.Anything,
						Message: mock.Anything,
					},
				},
			},
			want: []in.In{
				{
					ID:      "01",
					UserID:  mock.Anything,
					Message: mock.Anything,
				},
				{
					ID:      "03",
					UserID:  mock.Anything,
					Message: mock.Anything,
				},
			},
			wantErr: false,
		},
		{
			name: "should return error when user try to delete one message that was not sent by him",
			args: args{
				messageID: "02",
				userID:    mock.Anything,
				messages: []in.In{
					{
						ID:      "01",
						UserID:  mock.Anything,
						Message: mock.Anything,
					},
					{
						ID:      "02",
						UserID:  "user_id",
						Message: mock.Anything,
					},
					{
						ID:      "03",
						UserID:  mock.Anything,
						Message: mock.Anything,
					},
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "should return error when message does not exists in array",
			args: args{
				messageID: mock.Anything,
				userID:    mock.Anything,
				messages: []in.In{
					{
						ID:      "01",
						UserID:  mock.Anything,
						Message: mock.Anything,
					},
					{
						ID:      "02",
						UserID:  "user_id",
						Message: mock.Anything,
					},
					{
						ID:      "03",
						UserID:  mock.Anything,
						Message: mock.Anything,
					},
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {

		s.Run(tt.name, func() {
			d := &deleter{}
			got, err := d.Delete(tt.args.messages, tt.args.messageID, tt.args.userID)
			s.Assert().True((err != nil) == tt.wantErr, "Delete() error = %v, wantErr %v", err, tt.wantErr)
			s.Assert().True(reflect.DeepEqual(got, tt.want), "Delete() = %v, want %v", got, tt.want)

		})
	}
}
