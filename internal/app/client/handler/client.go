package handler

import (
	"bufio"
	"encoding/json"
	"fmt"
	in2 "github.com/maiaaraujo5/udp-chat/internal/app/client/handler/model/in"
	"github.com/maiaaraujo5/udp-chat/internal/app/client/handler/model/out"
	"log"
	"math/rand"
	"net"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type Client struct {
	conn     *net.UDPConn
	messages map[string]*out.Out
}

func NewClient(conn *net.UDPConn) *Client {
	return &Client{
		conn:     conn,
		messages: make(map[string]*out.Out),
	}
}

func (r *Client) Handle() error {
	r.sendMessage("NEW_MESSAGE", "Joined the room!")
	go r.receiveMessages()

	for {
		message, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			return err
		}

		message = strings.Trim(message, "\r\n")

		args := strings.Split(message, " ")
		command := strings.TrimSpace(args[0])

		var msg string
		for i := 1; i < len(args); i++ {
			msg += args[i] + " "
		}

		switch command {
		case "/msg":
			r.sendMessage("NEW_MESSAGE", msg)
		case "/del":
			r.sendMessage("DELETE_MESSAGE", msg)
		case "/quit":
			r.sendMessage("DISCONNECT", "Bye")
		}
	}
}

func (r *Client) receiveMessages() {
	for {
		message := make([]byte, 1024)
		rlen, _, err := r.conn.ReadFromUDP(message)
		if err != nil {
			log.Println(err)
		}

		msg, err := r.MarshalReceivedMessage(message, rlen)
		if err != nil {
			log.Println(err)
		}

		if msg.ID != "" && msg.Message == "" {
			delete(r.messages, msg.ID)
			cmd := exec.Command("clear")
			cmd.Stdout = os.Stdout
			cmd.Run()

			for _, msg := range r.messages {
				formatted := fmt.Sprintf("%s -> %s: %s", msg.ID, msg.UserID, msg.Message)
				fmt.Println(formatted)
			}

		} else {
			r.messages[msg.ID] = msg

			formatted := fmt.Sprintf("%s -> %s: %s", msg.ID, msg.UserID, msg.Message)
			fmt.Println(formatted)
		}
	}
}

func (r *Client) sendMessage(action, message string) {
	in := &in2.In{
		ID:      strconv.Itoa(rand.Intn(10000)),
		Action:  action,
		Message: message,
	}

	b, err := json.Marshal(in)
	if err != nil {
		log.Fatal(err)
	}

	_, err = r.conn.Write(b)
	if err != nil {
		log.Println(err)
	}
}

func (r *Client) MarshalReceivedMessage(message []byte, rlen int) (*out.Out, error) {
	req := &out.Out{}
	err := json.Unmarshal(message[:rlen], req)
	if err != nil {
		return nil, err
	}

	return req, nil
}
