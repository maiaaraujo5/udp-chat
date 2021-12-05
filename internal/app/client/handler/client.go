package handler

import (
	"bufio"
	"context"
	"fmt"
	in2 "github.com/maiaaraujo5/udp-chat/internal/app/client/domain/model/in"
	"github.com/maiaaraujo5/udp-chat/internal/app/client/domain/service"
	"log"
	"net"
	"os"
	"os/exec"
	"runtime"
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
	receiver service.Receiver
	deleter  service.Deleter
	creator  service.Creator
	messages []in2.In
}

func NewClient(conn *net.UDPConn, receiver service.Receiver, deleter service.Deleter, creator service.Creator) *Client {
	return &Client{
		conn:     conn,
		receiver: receiver,
		creator:  creator,
		deleter:  deleter,
	}
}

func (r *Client) Handle(parentCtx context.Context) error {
	err := r.handleSendMessage(NewConnection, "Joined the room!")
	if err != nil {
		log.Printf("error to connect in room. please try again. %s", err)
	}

	go r.handleNewMessages()

	r.print()

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
			log.Printf("error to send message: %s", err)
		}
	case "/del":
		err := r.handleDeleteMessage(msg)
		if err != nil {
			log.Printf("error to delete message: %s", err)
		}
	case "/quit":
		err := r.handleDisconnection()
		if err != nil {
			log.Printf("error to quit room: %s", err)
		}
	}
}

func (r *Client) print() {
	r.clearScreen()

	fmt.Println(fmt.Sprintf("=========UDP CHAT========="))
	fmt.Println(fmt.Sprintf("my user_id: %s \n", r.conn.LocalAddr().String()))

	for _, message := range r.messages {
		fmt.Println(fmt.Sprintf("%s -> %s: %s", message.ID, message.UserID, message.Message))
	}
}

func (r *Client) clearScreen() {
	clearScreenCommands := map[string]*exec.Cmd{
		"linux":   exec.Command("clear"),
		"windows": exec.Command("cmd", "/c", "cls"),
	}

	cmd := clearScreenCommands[runtime.GOOS]
	cmd.Stdout = os.Stdout
	cmd.Run()
}
