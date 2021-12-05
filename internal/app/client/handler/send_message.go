package handler

import (
	"encoding/json"
	"strings"

	"github.com/maiaaraujo5/udp-chat/internal/app/client/domain/model/in"
)

func (r *Client) handleSendMessage(action, message string) error {

	out := r.creator.Create(action, message)

	b, err := json.Marshal(out)
	if err != nil {
		return err
	}

	_, err = r.conn.Write(b)
	if err != nil {
		return err
	}

	if strings.EqualFold(action, NewMessage) {
		r.messages = append(r.messages, in.In{
			ID:      out.ID,
			UserID:  r.conn.LocalAddr().String(),
			Message: out.Message,
		})
	}

	r.print()

	return nil
}
