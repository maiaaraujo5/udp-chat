package service

import (
	"reflect"
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/maiaaraujo5/udp-chat/internal/app/client/domain/model/out"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type CreateMessageSuite struct {
	suite.Suite
}

func TestCreateMessageSuite(t *testing.T) {
	suite.Run(t, new(CreateMessageSuite))
}

func (s *CreateMessageSuite) SetupSuite() {
	monkey.Patch(time.Now, func() time.Time {
		return time.Date(2020, time.September, 10, 8, 1, 2, 3, time.UTC)
	})
}

func (s *CreateMessageSuite) TestNewCreator() {
	tests := []struct {
		name string
		want Creator
	}{
		{
			name: "should successfully build NewCreator",
			want: &creator{},
		},
	}
	for _, tt := range tests {

		s.Run(tt.name, func() {
			got := NewCreator()
			s.Assert().True(reflect.DeepEqual(got, tt.want), "NewCreator() = %v, want %v", got, tt.want)
		})
	}
}

func (s *CreateMessageSuite) Test_creator_Create() {
	type args struct {
		action  string
		message string
	}
	tests := []struct {
		name string
		args args
		want *out.Out
	}{
		{
			name: "should successfully create a new message",
			args: args{
				action:  mock.Anything,
				message: mock.Anything,
			},
			want: &out.Out{
				ID:      "436",
				Action:  mock.Anything,
				Message: mock.Anything,
			},
		},
	}
	for _, tt := range tests {

		s.Run(tt.name, func() {
			c := &creator{}
			got := c.Create(tt.args.action, tt.args.message)
			s.Assert().True(reflect.DeepEqual(got, tt.want), "Create() = %v, want %v", got, tt.want)
		})
	}
}
