package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"strings"
	"sync"
)

var instance *viper.Viper
var once sync.Once

func Load() *viper.Viper {
	once.Do(func() {
		v := viper.New()
		v.AutomaticEnv()
		v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

		conf := os.Getenv("CONF")
		files := strings.Split(conf, ",")

		for _, file := range files {
			if file == "" {
				continue
			}
			v.SetConfigFile(file)
			if err := v.MergeInConfig(); err != nil {
				log.Fatalf("error to load configs: %s", err)
			}
		}

		instance = v
	})

	return instance
}

func ReadConfigPath(i interface{}, key string) error {
	instance = Load()
	err := instance.UnmarshalKey(key, i)
	if err != nil {
		return err
	}

	return nil
}

func AddDefault(path string, value interface{}) {
	instance = Load()
	instance.SetDefault(path, value)
}

func GetStringValue(key string) string {
	instance = Load()
	return instance.GetString(key)
}

func GetBoolValue(key string) bool {
	instance = Load()
	return instance.GetBool(key)
}
