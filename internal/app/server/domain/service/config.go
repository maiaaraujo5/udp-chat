package service

import gostartconfig "github.com/maiaaraujo5/gostart/config"

const root = "app.service"

type Config struct {
	MaxMessagesInHistory int
}

func defaultConfig() {
	gostartconfig.AddDefault(root+".maxmessagesinhistory", 20)
}

func NewConfig() (*Config, error) {
	c := &Config{}

	defaultConfig()

	err := gostartconfig.ReadConfigPath(c, root)
	if err != nil {
		return nil, err
	}

	return c, nil
}
