package util

import "net"

func CreateUdpServer() *net.UDPConn {
	conn, _ := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.ParseIP("0.0.0.0"),
		Port: 3500,
	})

	return conn
}
