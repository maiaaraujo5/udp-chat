package client

import (
	"net"
	"strconv"
	"strings"
)

func Connect() (*net.UDPConn, error) {

	c, err := NewConfig()
	if err != nil {
		return nil, err
	}

	addr, err := net.ResolveUDPAddr("udp", strings.Join([]string{c.IP, strconv.Itoa(c.Port)}, ":"))
	if err != nil {
		return nil, err
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
