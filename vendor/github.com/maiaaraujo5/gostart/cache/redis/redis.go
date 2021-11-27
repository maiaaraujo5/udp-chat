package redis

import (
	"github.com/go-redis/redis/v8"
)

func Connect() (*redis.Client, error) {
	config, err := NewConfig()
	if err != nil {
		return nil, err
	}

	return redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       config.DB,
	}), nil
}
