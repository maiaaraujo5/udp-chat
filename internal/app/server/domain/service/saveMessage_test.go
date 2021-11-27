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

type SaveMessageSuite struct {
	suite.Suite
}

func TestSaveMessageSuite(t *testing.T) {
	suite.Run(t, new(SaveMessageSuite))
}

func (s *SaveMessageSuite) TestNewSaver() {
	type args struct {
		repository *mocks.Repository
		config     *config
	}
	tests := []struct {
		name string
		args args
		want Saver
	}{
		{
			name: "should build NewSaver successfully",
			args: args{
				repository: new(mocks.Repository),
				config:     new(config),
			},
			want: &SaverImpl{
				repository: new(mocks.Repository),
				config:     new(config),
			},
		},
	}
	for _, tt := range tests {

		s.Run(tt.name, func() {
			got := NewSaver(tt.args.repository, tt.args.config)
			s.Assert().True(reflect.DeepEqual(got, tt.want), "NewSaver() = %v, want %v", got, tt.want)
		})
	}
}

func (s *SaveMessageSuite) TestSaverImpl_Execute() {
	type fields struct {
		repository *mocks.Repository
		config     *config
	}
	type args struct {
		ctx     context.Context
		message *model.Message
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		mock    func(repository *mocks.Repository)
	}{
		{
			name: "should save message successfully when repository does not have messages",
			fields: fields{
				repository: new(mocks.Repository),
				config: &config{
					maxMessagesInHistory: 3,
				},
			},
			args: args{
				ctx: context.Background(),
				message: &model.Message{
					ID:      mock.Anything,
					UserID:  mock.Anything,
					Message: mock.Anything,
				},
			},
			wantErr: false,
			mock: func(repository *mocks.Repository) {
				repository.On("List", mock.Anything).Return(list.New(), nil).Once()

				repository.On("SaveAll", mock.Anything, func() *list.List {
					l := list.New()
					l.PushBack(model.Message{ID: mock.Anything, UserID: mock.Anything, Message: mock.Anything})
					return l
				}()).Return(nil).Once()
			},
		},
		{
			name: "should successfully save messages in order when repository return one message",
			fields: fields{
				repository: new(mocks.Repository),
				config: &config{
					maxMessagesInHistory: 3,
				},
			},
			args: args{
				ctx: context.Background(),
				message: &model.Message{
					ID:      "2",
					UserID:  mock.Anything,
					Message: mock.Anything,
				},
			},
			wantErr: false,
			mock: func(repository *mocks.Repository) {
				repository.On("List", mock.Anything).Return(func() *list.List {
					l := list.New()
					l.PushBack(model.Message{ID: "1", UserID: mock.Anything, Message: mock.Anything})
					return l
				}(), nil).Once()

				repository.On("SaveAll", mock.Anything, func() *list.List {
					l := list.New()
					l.PushBack(model.Message{ID: "1", UserID: mock.Anything, Message: mock.Anything})
					l.PushBack(model.Message{ID: "2", UserID: mock.Anything, Message: mock.Anything})
					return l
				}()).Return(nil).Once()
			},
		},
		{
			name: "should remove the first message of list when length of the list is in the maximum configured to history and insert the new message in the final",
			fields: fields{
				repository: new(mocks.Repository),
				config: &config{
					maxMessagesInHistory: 3,
				},
			},
			args: args{
				ctx: context.Background(),
				message: &model.Message{
					ID:      "4",
					UserID:  mock.Anything,
					Message: mock.Anything,
				},
			},
			wantErr: false,
			mock: func(repository *mocks.Repository) {
				repository.On("List", mock.Anything).Return(func() *list.List {
					l := list.New()
					l.PushBack(model.Message{ID: "1", UserID: mock.Anything, Message: mock.Anything})
					l.PushBack(model.Message{ID: "2", UserID: mock.Anything, Message: mock.Anything})
					l.PushBack(model.Message{ID: "3", UserID: mock.Anything, Message: mock.Anything})
					return l
				}(), nil).Once()

				repository.On("SaveAll", mock.Anything, func() *list.List {
					l := list.New()
					l.PushBack(model.Message{ID: "2", UserID: mock.Anything, Message: mock.Anything})
					l.PushBack(model.Message{ID: "3", UserID: mock.Anything, Message: mock.Anything})
					l.PushBack(model.Message{ID: "4", UserID: mock.Anything, Message: mock.Anything})
					return l
				}()).Return(nil).Once()
			},
		},
		{
			name: "should return error when repository.List returns error",
			fields: fields{
				repository: new(mocks.Repository),
				config: &config{
					maxMessagesInHistory: 3,
				},
			},
			args: args{
				ctx: context.Background(),
				message: &model.Message{
					ID:      mock.Anything,
					UserID:  mock.Anything,
					Message: mock.Anything,
				},
			},
			wantErr: true,
			mock: func(repository *mocks.Repository) {
				repository.On("List", mock.Anything).Return(nil, errors.New("error to return list")).Once()
			},
		},
		{
			name: "should return error when repository.SaveAll returns error",
			fields: fields{
				repository: new(mocks.Repository),
				config: &config{
					maxMessagesInHistory: 3,
				},
			},
			args: args{
				ctx: context.Background(),
				message: &model.Message{
					ID:      "2",
					UserID:  mock.Anything,
					Message: mock.Anything,
				},
			},
			wantErr: true,
			mock: func(repository *mocks.Repository) {
				repository.On("List", mock.Anything).Return(func() *list.List {
					l := list.New()
					l.PushBack(model.Message{ID: "1", UserID: mock.Anything, Message: mock.Anything})
					return l
				}(), nil).Once()

				repository.On("SaveAll", mock.Anything, func() *list.List {
					l := list.New()
					l.PushBack(model.Message{ID: "1", UserID: mock.Anything, Message: mock.Anything})
					l.PushBack(model.Message{ID: "2", UserID: mock.Anything, Message: mock.Anything})
					return l
				}()).Return(errors.New("error to save list")).Once()
			},
		},
	}
	for _, tt := range tests {

		s.Run(tt.name, func() {
			tt.mock(tt.fields.repository)

			r := &SaverImpl{
				repository: tt.fields.repository,
				config:     tt.fields.config,
			}

			err := r.Execute(tt.args.ctx, tt.args.message)
			s.Assert().True((err != nil) == tt.wantErr, "Execute() error = %v, wantErr %v", err, tt.wantErr)

			tt.fields.repository.AssertExpectations(s.T())
		})
	}
}
