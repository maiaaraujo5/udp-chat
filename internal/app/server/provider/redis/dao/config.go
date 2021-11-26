package dao

import "github.com/maiaaraujo5/gostart/config"

const root = "app.provider.redis"

type Config struct {
	Key string
}

func defaultConfig() {
	config.AddDefault(root+".key", "messages")
}

func NewConfig() (*Config, error) {
	c := &Config{}

	defaultConfig()

	err := config.ReadConfigPath(c, root)
	if err != nil {
		return nil, err
	}

	return c, nil
}
