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

type RecoverMessageSuite struct {
	suite.Suite
}

func TestRecoverMessageSuite(t *testing.T) {
	suite.Run(t, new(RecoverMessageSuite))
}

func (s *RecoverMessageSuite) TestNewRecover() {
	type args struct {
		repository *mocks.Repository
	}
	tests := []struct {
		name string
		args args
		want Recover
	}{
		{
			name: "should build NewRecover successfully",
			args: args{
				repository: new(mocks.Repository),
			},
			want: &recoverImpl{
				repository: new(mocks.Repository),
			},
		},
	}
	for _, tt := range tests {

		s.Run(tt.name, func() {
			got := NewRecover(tt.args.repository)
			s.Assert().True(reflect.DeepEqual(got, tt.want), "NewRecover() = %v, want %v", got, tt.want)
		})
	}
}

func (s *RecoverMessageSuite) TestRecoverImpl_Execute() {
	type fields struct {
		repository *mocks.Repository
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []model.Message
		wantErr bool
		mock    func(repository *mocks.Repository)
	}{
		{
			name: "should successfully return all messages that returned from repository in order",
			fields: fields{
				repository: new(mocks.Repository),
			},
			args: args{
				ctx: context.Background(),
			},
			want: []model.Message{
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
			wantErr: false,
			mock: func(repository *mocks.Repository) {
				repository.On("List", mock.Anything).Return(func() *list.List {
					l := list.New()
					l.PushBack(model.Message{ID: "1", UserID: mock.Anything, Message: mock.Anything})
					l.PushBack(model.Message{ID: "2", UserID: mock.Anything, Message: mock.Anything})
					return l
				}(), nil).Once()
			},
		},
		{
			name: "should return error when repository return error",
			fields: fields{
				repository: new(mocks.Repository),
			},
			args: args{
				ctx: context.Background(),
			},
			want:    nil,
			wantErr: true,
			mock: func(repository *mocks.Repository) {
				repository.On("List", mock.Anything).Return(nil, errors.New("error to return messages"))
			},
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {

			tt.mock(tt.fields.repository)

			r := &recoverImpl{
				repository: tt.fields.repository,
			}

			got, err := r.Recover(tt.args.ctx)
			s.Assert().True((err != nil) == tt.wantErr, "Execute() error = %v, wantErr %v", err, tt.wantErr)
			s.Assert().True(reflect.DeepEqual(got, tt.want), "Execute() = %v, want %v", got, tt.want)

			tt.fields.repository.AssertExpectations(s.T())
		})
	}
}
