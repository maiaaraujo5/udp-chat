package main

import (
	"github.com/maiaaraujo5/udp-chat/internal/app/client/handler"
	"log"
	"net"
)

func main() {

	addr, err := net.ResolveUDPAddr("udp", "0.0.0.0:3000")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Established connection to %s \n", "0.0.0.0:3000")

	defer func(conn *net.UDPConn) {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}(conn)

	err = handler.NewClient(conn).Handle()
	if err != nil {
		panic(err)
	}
}
