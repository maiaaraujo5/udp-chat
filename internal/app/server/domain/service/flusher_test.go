package service

import (
	"context"
	"errors"
	"github.com/maiaaraujo5/udp-chat/internal/app/server/domain/repository/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"reflect"
	"testing"
)

type FlusherSuite struct {
	suite.Suite
}

func TestFlusherSuite(t *testing.T) {
	suite.Run(t, new(FlusherSuite))
}

func (s *FlusherSuite) TestNewFlusher() {
	type args struct {
		repository *mocks.Repository
	}
	tests := []struct {
		name string
		args args
		want Flusher
	}{
		{
			name: "should build NewFlusher successfully",
			args: args{
				repository: new(mocks.Repository),
			},
			want: &FlusherImpl{
				repository: new(mocks.Repository),
			},
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {
			got := NewFlusher(tt.args.repository)
			s.Assert().True(reflect.DeepEqual(got, tt.want), "NewFlusher() = %v, want %v", got, tt.want)
		})
	}
}

func (s *FlusherSuite) TestFlusherImpl_Execute() {
	type fields struct {
		repository *mocks.Repository
	}
	type args struct {
		parentCtx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		mock    func(repository *mocks.Repository)
	}{
		{
			name: "should flush repository successfully",
			fields: fields{
				repository: new(mocks.Repository),
			},
			args: args{
				parentCtx: context.Background(),
			},
			wantErr: false,
			mock: func(repository *mocks.Repository) {
				repository.On("Flush", mock.Anything).Return(nil).Once()
			},
		},
		{
			name: "should return error when repository return error",
			fields: fields{
				repository: new(mocks.Repository),
			},
			args: args{
				parentCtx: context.Background(),
			},
			wantErr: true,
			mock: func(repository *mocks.Repository) {
				repository.On("Flush", mock.Anything).Return(errors.New("error to flush database")).Once()
			},
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {

			tt.mock(tt.fields.repository)

			r := &FlusherImpl{
				repository: tt.fields.repository,
			}

			err := r.Execute(tt.args.parentCtx)
			s.Assert().True((err != nil) == tt.wantErr, "Execute() error = %v, wantErr %v", err, tt.wantErr)

			tt.fields.repository.AssertExpectations(s.T())
		})
	}
}
