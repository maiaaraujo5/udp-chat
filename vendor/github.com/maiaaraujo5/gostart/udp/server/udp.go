package server

import "net"

func Server() (*net.UDPConn, error) {

	config, err := NewConfig()
	if err != nil {
		return nil, err
	}

	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.ParseIP(config.IP),
		Port: config.Port,
	})

	return conn, nil
}
