package main

import (
	"context"
	"fmt"
	"github.com/maiaaraujo5/udp-chat/internal/app/server/domain/service"
	"github.com/maiaaraujo5/udp-chat/internal/app/server/server/udp/handler"
	"log"
	"net"
)

func main() {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.ParseIP("0.0.0.0"),
		Port: 3000,
	})

	if err != nil {
		panic(err)
	}

	defer func(conn *net.UDPConn) {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}(conn)

	fmt.Printf("server listening %s\n", conn.LocalAddr().String())

	rec := service.NewRecover()
	saver := service.NewSaver()
	flusher := service.NewFlusher()

	err = handler.NewServer(rec, saver, flusher, conn).Handle(context.Background())
	if err != nil {
		log.Println(err)
	}
}
