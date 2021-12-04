package handler

import (
	"encoding/json"
	"github.com/maiaaraujo5/udp-chat/internal/app/client/handler/model/out"
	"math/rand"
	"strconv"
	"time"
)

func (r *Client) handleSendMessage(action, message string) error {
	msg := &out.Out{
		ID:      r.generateId(),
		Action:  action,
		Message: message,
	}

	b, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	_, err = r.conn.Write(b)
	if err != nil {
		return err
	}

	return nil
}

func (r *Client) generateId() string {
	rand.Seed(time.Now().Unix())
	number := rand.Intn(9999)
	return strconv.Itoa(number)
}
