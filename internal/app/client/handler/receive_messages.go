package handler

import (
	"encoding/json"
	"fmt"
	"github.com/maiaaraujo5/udp-chat/internal/app/client/handler/model/in"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func (r *Client) handleNewMessages() {
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

		if msg != nil {
			r.receiveMessages(msg)
		}

	}
}

func (r *Client) receiveMessages(msg *in.In) {
	if r.isDeletedMessage(msg) {
		r.clearScreen()

		var newMessages []in.In

		for _, m := range r.messages {
			if m.ID != strings.TrimSpace(msg.ID) {
				newMessages = append(newMessages, m)
				formatted := fmt.Sprintf("%s -> %s: %s", m.ID, m.UserID, m.Message)
				fmt.Println(formatted)
			}
		}

		r.messages = newMessages
		return
	}

	r.messages = append(r.messages, *msg)

	formatted := fmt.Sprintf("%s -> %s: %s", msg.ID, msg.UserID, msg.Message)
	fmt.Println(formatted)

}

func (r *Client) isDeletedMessage(msg *in.In) bool {
	return msg.ID != "" && msg.Message == ""
}

func (r *Client) unmarshalReceivedMessage(message []byte, rlen int) (*in.In, error) {
	req := &in.In{}
	err := json.Unmarshal(message[:rlen], req)
	if err != nil {
		return nil, err
	}

	return req, nil
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
