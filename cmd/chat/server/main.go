package main

import (
	"github.com/maiaaraujo5/udp-chat/internal/app/server/fx/module/runner"
)

func main() {

	err := runner.Run()

	if err != nil {
		panic(err)
	}
}
