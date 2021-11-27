package instance

import (
	"github.com/maiaaraujo5/gostart/log"
	"github.com/maiaaraujo5/gostart/log/zerolog"
	"sync"
)

var instance log.Log
var once sync.Once

func Load() log.Log {
	once.Do(func() {
		log.DefaultConfig()
		impl := log.ImpValue()

		if impl == "zerolog" {
			instance = zerolog.NewZeroLog()
		}
	})

	return instance
}
