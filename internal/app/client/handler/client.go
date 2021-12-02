package handler

import (
	"bufio"
	"github.com/maiaaraujo5/udp-chat/internal/app/client/handler/model/out"
	"net"
	"os"
	"strings"
)

const (
	NewConnection = "NEW_CONNECTION"
	NewMessage    = "NEW_MESSAGE"
	DeleteMessage = "DELETE_MESSAGE"
	Disconnection = "DISCONNECT"
)

type Client struct {
	conn     *net.UDPConn
	messages []out.Out
}

func NewClient(conn *net.UDPConn) *Client {
	return &Client{
		conn: conn,
	}
}

func (r *Client) Handle() error {
	r.handleSendMessage(NewConnection, "Joined the room!")
	go r.handleNewMessages()

	for {
		message, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			return err
		}

		message = strings.Trim(message, "\r\n")

		args := strings.Split(message, " ")

		command := strings.TrimSpace(args[0])
		msg := strings.Join(args[1:], " ")

		switch command {
		case "/msg":
			r.handleSendMessage(NewMessage, msg)
		case "/del":
			r.handleDeleteMessage(msg)
		case "/quit":
			r.handleSendMessage(Disconnection, "Bye")
		}
	}
}
