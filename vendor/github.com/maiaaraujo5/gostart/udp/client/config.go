package client

import "github.com/maiaaraujo5/gostart/config"

const root = "gostart.udp.client.connect"

type Config struct {
	IP   string
	Port int
}

func defaultConfig() {
	config.AddDefault(root+".ip", "0.0.0.0")
	config.AddDefault(root+".port", "3000")
}

func NewConfig() (*Config, error) {
	c := &Config{}

	defaultConfig()

	if err := config.ReadConfigPath(c, root); err != nil {
		return nil, err
	}

	return c, nil
}