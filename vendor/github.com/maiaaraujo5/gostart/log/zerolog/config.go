package zerolog

import "github.com/maiaaraujo5/gostart/config"

const root = "gostart.zerolog"

type Config struct {
	Level string
}

func defaultConfig() {
	config.AddDefault(root+".level", "info")
}

func newConfig() (*Config, error) {
	c := &Config{}

	defaultConfig()

	err := config.ReadConfigPath(c, root)
	if err != nil {
		return nil, err
	}

	return c, nil
}
