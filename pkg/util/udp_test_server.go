package util

import (
	"net"
)

func CreateUdpServer() *net.UDPConn {
	conn, _ := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.ParseIP("0.0.0.0"),
		Port: 3500,
	})

	return conn
}

func CreateUdpClient() *net.UDPConn {
	addr, _ := net.ResolveUDPAddr("udp", "0.0.0.0:3500")

	conn, _ := net.DialUDP("udp", nil, addr)
	return conn
}
