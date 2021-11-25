package main

import (
	"github.com/maiaaraujo5/udp-chat/internal/app/server/fx/module/runnner"
)

func main() {

	err := runnner.Run()

	if err != nil {
		panic(err)
	}
}
