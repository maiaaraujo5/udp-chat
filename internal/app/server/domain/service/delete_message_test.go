package service

import (
	"container/list"
	"context"
	"errors"
	"github.com/maiaaraujo5/udp-chat/internal/app/server/domain/model"
	"github.com/maiaaraujo5/udp-chat/internal/app/server/domain/repository/mocks"
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

func (s *DeleteMessageSuite) TestNewDelete() {
	type args struct {
		repository *mocks.Repository
	}
	tests := []struct {
		name string
		args args
		want Deleter
	}{
		{
			name: "should build NewDelete successfully",
			args: args{
				repository: new(mocks.Repository),
			},
			want: &deleter{
				repository: new(mocks.Repository),
			},
		},
	}
	for _, tt := range tests {

		s.Run(tt.name, func() {
			got := NewDelete(tt.args.repository)
			s.Assert().True(reflect.DeepEqual(got, tt.want), "NewDelete() = %v, want %v", got, tt.want)
		})
	}
}

func (s *DeleteMessageSuite) Testdeleter_Execute() {
	type fields struct {
		repository *mocks.Repository
	}
	type args struct {
		parentCtx context.Context
		message   *model.Message
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		mock    func(repository *mocks.Repository)
	}{
		{
			name: "should successfully delete one message",
			fields: fields{
				repository: new(mocks.Repository),
			},
			args: args{
				parentCtx: context.Background(),
				message: &model.Message{
					ID:     "2",
					UserID: "user_id",
				},
			},
			wantErr: false,
			mock: func(repository *mocks.Repository) {
				repository.On("List", mock.Anything).Return(func() *list.List {
					l := list.New()
					l.PushBack(model.Message{ID: "1", UserID: mock.Anything, Message: mock.Anything})
					l.PushBack(model.Message{ID: "2", UserID: "user_id", Message: mock.Anything})
					l.PushBack(model.Message{ID: "3", UserID: mock.Anything, Message: mock.Anything})
					return l
				}(), nil).Once()

				repository.On("SaveAll", mock.Anything, func() *list.List {
					l := list.New()
					l.PushBack(model.Message{ID: "1", UserID: mock.Anything, Message: mock.Anything})
					l.PushBack(model.Message{ID: "3", UserID: mock.Anything, Message: mock.Anything})
					return l
				}()).Return(nil).Once()
			},
		},
		{
			name: "should not delete a message that does not belong to the user",
			fields: fields{
				repository: new(mocks.Repository),
			},
			args: args{
				parentCtx: context.Background(),
				message: &model.Message{
					ID:     "1",
					UserID: "user_id",
				},
			},
			wantErr: false,
			mock: func(repository *mocks.Repository) {
				repository.On("List", mock.Anything).Return(func() *list.List {
					l := list.New()
					l.PushBack(model.Message{ID: "1", UserID: mock.Anything, Message: mock.Anything})
					l.PushBack(model.Message{ID: "2", UserID: "user_id", Message: mock.Anything})
					l.PushBack(model.Message{ID: "3", UserID: mock.Anything, Message: mock.Anything})
					return l
				}(), nil).Once()
			},
		},
		{
			name: "should return error when repository.List returns error",
			fields: fields{
				repository: new(mocks.Repository),
			},
			args: args{
				parentCtx: context.Background(),
				message: &model.Message{
					ID:     mock.Anything,
					UserID: mock.Anything,
				},
			},
			wantErr: true,
			mock: func(repository *mocks.Repository) {
				repository.On("List", mock.Anything).Return(nil, errors.New("error to recover messages"))
			},
		},
		{
			name: "should return error when repository.SaveAll returns error",
			fields: fields{
				repository: new(mocks.Repository),
			},
			args: args{
				parentCtx: context.Background(),
				message: &model.Message{
					ID:     "2",
					UserID: "user_id",
				},
			},
			wantErr: true,
			mock: func(repository *mocks.Repository) {
				repository.On("List", mock.Anything).Return(func() *list.List {
					l := list.New()
					l.PushBack(model.Message{ID: "1", UserID: mock.Anything, Message: mock.Anything})
					l.PushBack(model.Message{ID: "2", UserID: "user_id", Message: mock.Anything})
					l.PushBack(model.Message{ID: "3", UserID: mock.Anything, Message: mock.Anything})
					return l
				}(), nil).Once()

				repository.On("SaveAll", mock.Anything, func() *list.List {
					l := list.New()
					l.PushBack(model.Message{ID: "1", UserID: mock.Anything, Message: mock.Anything})
					l.PushBack(model.Message{ID: "3", UserID: mock.Anything, Message: mock.Anything})
					return l
				}()).Return(errors.New("error to save messages")).Once()
			},
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {

			tt.mock(tt.fields.repository)

			d := &deleter{
				repository: tt.fields.repository,
			}

			err := d.Delete(tt.args.parentCtx, tt.args.message)
			s.Assert().True((err != nil) == tt.wantErr, "Execute() error = %v, wantErr %v", err, tt.wantErr)

			tt.fields.repository.AssertExpectations(s.T())
		})
	}
}
