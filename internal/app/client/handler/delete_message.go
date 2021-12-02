package handler

import (
	"errors"
	"strings"
)

func (r *Client) handleDeleteMessage(msg string) error {

	if !r.clientIsOwnerOfTheMessage(msg) {
		return errors.New("you can't delete a message that was not sent by you")
	}

	err := r.handleSendMessage(DeleteMessage, msg)
	if err != nil {
		return err
	}
	return nil

}

func (r *Client) clientIsOwnerOfTheMessage(messageID string) bool {
	for _, message := range r.messages {
		if strings.EqualFold(message.ID, messageID) && strings.EqualFold(message.UserID, r.conn.LocalAddr().String()) {
			return true
		}
	}

	return false
}
