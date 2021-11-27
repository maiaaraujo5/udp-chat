module github.com/maiaaraujo5/udp-chat

go 1.16

require (
	github.com/go-redis/redis/v8 v8.11.4
	github.com/maiaaraujo5/gostart v0.0.0-20211122220756-17b5f5b0f7f4
	github.com/stretchr/testify v1.7.0
	go.uber.org/dig v1.12.0 // indirect
	go.uber.org/fx v1.13.1
	go.uber.org/goleak v1.1.11 // indirect
	go.uber.org/zap v1.16.0 // indirect
	golang.org/x/sys v0.0.0-20210903071746-97244b99971b // indirect
)

replace github.com/maiaaraujo5/gostart => ../gostart
