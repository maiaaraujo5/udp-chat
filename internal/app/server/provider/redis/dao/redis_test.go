package dao

import (
	"container/list"
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/maiaaraujo5/udp-chat/internal/app/server/domain/model"
	"github.com/maiaaraujo5/udp-chat/internal/app/server/domain/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"reflect"
	"strings"
	"testing"

	"github.com/alicebob/miniredis/v2"
)

type DaoSuite struct {
	suite.Suite
	client *redis.Client
}

func TestDaoSuite(t *testing.T) {
	suite.Run(t, new(DaoSuite))
}

func (s *DaoSuite) SetupSuite() {
	mr, err := miniredis.Run()
	if err != nil {
		panic(err)
	}

	s.client = redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})
}

func (s *DaoSuite) SetupTest() {
	s.client.Del(s.client.Context(), "test_message")
	s.client.RPush(s.client.Context(), "test_message", "mock-mock-mock")
}

func (s *DaoSuite) TestNewRedis() {
	type args struct {
		client *redis.Client
		config *Config
	}
	tests := []struct {
		name string
		args args
		want repository.Repository
	}{
		{
			name: "should build NewRedis successfully",
			args: args{
				client: s.client,
				config: new(Config),
			},
			want: &Redis{
				client: s.client,
				config: new(Config),
			},
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {
			got := NewRedis(tt.args.client, tt.args.config)
			s.Assert().True(reflect.DeepEqual(got, tt.want), "NewRedis() = %v, want %v", got, tt.want)
		})
	}
}

func (s *DaoSuite) TestRedis_Flush() {
	type fields struct {
		client *redis.Client
		config *Config
	}
	type args struct {
		parentCtx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		assert  func(client *redis.Client)
	}{
		{
			name: "successfully flush database",
			fields: fields{
				client: s.client,
				config: &Config{
					Key: "test_message",
				},
			},
			args: args{
				parentCtx: context.Background(),
			},
			wantErr: false,
			assert: func(client *redis.Client) {
				values, _ := client.LRange(s.client.Context(), "test_message", 0, -1).Result()
				assert.Empty(s.T(), values)
			},
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {
			r := &Redis{
				client: tt.fields.client,
				config: tt.fields.config,
			}

			err := r.Flush(tt.args.parentCtx)
			tt.assert(tt.fields.client)
			s.Assert().True((err != nil) == tt.wantErr, "Flush() error = %v, wantErr %v", err, tt.wantErr)
		})
	}
}

func (s *DaoSuite) TestRedis_List() {
	type fields struct {
		client *redis.Client
		config *Config
	}
	type args struct {
		parentCtx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *list.List
		wantErr bool
	}{
		{
			name: "should successfully return a list when key exists",
			fields: fields{
				client: s.client,
				config: &Config{
					Key: "test_message",
				},
			},
			args: args{
				parentCtx: context.Background(),
			},
			want: func() *list.List {
				l := list.New()
				l.PushBack(model.Message{ID: "mock", UserID: "mock", Message: "mock"})
				return l
			}(),
			wantErr: false,
		},
		{
			name: "should return a empty list when key does not exist",
			fields: fields{
				client: s.client,
				config: &Config{
					Key: mock.Anything,
				},
			},
			args: args{
				parentCtx: context.Background(),
			},
			want:    list.New(),
			wantErr: false,
		},
	}
	for _, tt := range tests {

		s.Run(tt.name, func() {
			r := &Redis{
				client: tt.fields.client,
				config: tt.fields.config,
			}
			got, err := r.List(tt.args.parentCtx)
			s.Assert().True((err != nil) == tt.wantErr, "List() error = %v, wantErr %v", err, tt.wantErr)
			s.Assert().True(reflect.DeepEqual(got, tt.want), "List() got = %v, want %v", got, tt.want)
		})
	}
}

func (s *DaoSuite) TestRedis_SaveAll() {
	type fields struct {
		client *redis.Client
		config *Config
	}
	type args struct {
		parentCtx context.Context
		messages  *list.List
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		assert  func(client *redis.Client)
	}{
		{
			name: "should successfully save new messages",
			fields: fields{
				client: s.client,
				config: &Config{
					Key: "test_message",
				},
			},
			args: args{
				parentCtx: context.Background(),
				messages: func() *list.List {
					l := list.New()
					l.PushBack(model.Message{ID: mock.Anything, UserID: mock.Anything, Message: mock.Anything})
					return l
				}(),
			},
			wantErr: false,
			assert: func(client *redis.Client) {
				values, _ := client.LRange(client.Context(), "test_message", 0, -1).Result()
				assert.Equal(s.T(), []string{strings.Join([]string{mock.Anything, mock.Anything, mock.Anything}, "-")}, values)
			},
		},
	}
	for _, tt := range tests {

		s.Run(tt.name, func() {
			r := &Redis{
				client: tt.fields.client,
				config: tt.fields.config,
			}

			err := r.SaveAll(tt.args.parentCtx, tt.args.messages)
			tt.assert(s.client)
			s.Assert().True((err != nil) == tt.wantErr, "SaveAll() error = %v, wantErr %v", err, tt.wantErr)
		})
	}
}
