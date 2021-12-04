package handler

import (
	"bufio"
	"github.com/maiaaraujo5/udp-chat/internal/app/client/handler/model/in"
	"log"
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
	messages []in.In
}

func NewClient(conn *net.UDPConn) *Client {
	return &Client{
		conn: conn,
	}
}

func (r *Client) Handle() error {
	err := r.handleSendMessage(NewConnection, "Joined the room!")
	if err != nil {
		log.Printf("error to connect in room. please try again. %s", err)
	}
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

		r.handleCommands(command, msg)

	}
}

func (r *Client) handleCommands(command, msg string) {
	switch command {
	case "/msg":
		err := r.handleSendMessage(NewMessage, msg)
		if err != nil {
			log.Printf("error to send message. please try again. %s", err)
		}
	case "/del":
		err := r.handleDeleteMessage(msg)
		if err != nil {
			log.Printf("error to delete message. please try again. %s", err)
		}
	case "/quit":
		err := r.handleSendMessage(Disconnection, "Bye")
		if err != nil {
			log.Printf("error to quit room. please try again. %s", err)
		}
	}
}
