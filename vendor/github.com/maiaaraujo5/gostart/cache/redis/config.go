package redis

import "github.com/maiaaraujo5/gostart/config"

const root = "gostart.redis"

type Config struct {
	Addr     string
	Password string
	DB       int
}

func defaultConfig() {
	config.AddDefault(root+".addr", "localhost:6379")
	config.AddDefault(root+".password", "")
	config.AddDefault(root+".db", 0)
}

func NewConfig() (*Config, error) {
	c := &Config{}

	defaultConfig()

	if err := config.ReadConfigPath(c, root); err != nil {
		return nil, err
	}

	return c, nil
}
