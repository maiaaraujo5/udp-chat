package handler

import (
	"encoding/json"
	"log"

	in "github.com/maiaaraujo5/udp-chat/internal/app/client/domain/model/in"
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
			r.receiveMessage(msg)
		}
	}
}

func (r *Client) receiveMessage(msg *in.In) {
	r.messages = r.receiver.Receive(r.messages, msg)
	r.print()
}

func (r *Client) unmarshalReceivedMessage(message []byte, rlen int) (*in.In, error) {
	req := &in.In{}
	err := json.Unmarshal(message[:rlen], req)
	if err != nil {
		return nil, err
	}

	return req, nil
}
