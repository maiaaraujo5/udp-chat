package handler

import (
	"encoding/json"
	in2 "github.com/maiaaraujo5/udp-chat/internal/app/client/handler/model/in"
	"log"
	"math/rand"
	"strconv"
	"time"
)

func (r *Client) handleSendMessage(action, message string) {
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
