package service

import gostartconfig "github.com/maiaaraujo5/gostart/config"

const root = "app.service"

type config struct {
	MaxMessagesInHistory int
}

func defaultConfig() {
	gostartconfig.AddDefault(root+".maxmessagesinhistory", 20)
}

func NewConfig() (*config, error) {
	c := &config{}

	defaultConfig()

	err := gostartconfig.ReadConfigPath(c, root)
	if err != nil {
		return nil, err
	}

	return c, nil
}
