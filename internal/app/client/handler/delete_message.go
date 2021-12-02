package handler

import (
	"log"
	"strings"
)

func (r *Client) handleDeleteMessage(msg string) {
	if r.clientIsOwnerOfTheMessage(strings.TrimSpace(msg)) {
		r.handleSendMessage(DeleteMessage, msg)
		return
	}

	log.Println("you can't delete a message that was not sent by you!")
}

func (r *Client) clientIsOwnerOfTheMessage(messageID string) bool {
	for _, message := range r.messages {
		if strings.EqualFold(message.ID, messageID) && strings.EqualFold(message.UserID, r.conn.LocalAddr().String()) {
			return true
		}
	}

	return false
}
