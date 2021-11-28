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
	"time"
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
	r.sendMessage(NewConnection, "Joined the room!")
	go r.receiveMessages()

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
			r.sendMessage(NewMessage, msg)
		case "/del":
			if r.clientIsOwnerOfTheMessage(strings.TrimSpace(msg)) {
				r.sendMessage(DeleteMessage, msg)
			}
		case "/quit":
			r.sendMessage(Disconnection, "Bye")
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

		msg, err := r.unmarshalReceivedMessage(message, rlen)
		if err != nil {
			log.Println(err)
		}

		if msg.ID != "" && msg.Message == "" {
			var newMessages []out.Out
			cmd := exec.Command("clear")
			cmd.Stdout = os.Stdout
			cmd.Run()

			for _, m := range r.messages {
				if m.ID != strings.TrimSpace(msg.ID) {
					newMessages = append(newMessages, m)
					formatted := fmt.Sprintf("%s -> %s: %s", m.ID, m.UserID, m.Message)
					fmt.Println(formatted)
				}
			}

			r.messages = newMessages

		} else {
			r.messages = append(r.messages, *msg)

			formatted := fmt.Sprintf("%s -> %s: %s", msg.ID, msg.UserID, msg.Message)
			fmt.Println(formatted)
		}
	}
}

func (r *Client) sendMessage(action, message string) {
	in := &in2.In{
		ID:      r.generateId(),
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

func (r *Client) generateId() string {
	rand.Seed(time.Now().Unix())
	number := rand.Intn(9999)
	return strconv.Itoa(number)
}

func (r *Client) clientIsOwnerOfTheMessage(messageID string) bool {
	for _, message := range r.messages {
		if strings.EqualFold(message.ID, messageID) && strings.EqualFold(message.UserID, r.conn.LocalAddr().String()) {
			return true
		}
	}

	return false
}

func (r *Client) unmarshalReceivedMessage(message []byte, rlen int) (*out.Out, error) {
	req := &out.Out{}
	err := json.Unmarshal(message[:rlen], req)
	if err != nil {
		return nil, err
	}

	return req, nil
}
