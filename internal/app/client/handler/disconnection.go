package handler

import (
	"os"
)

func (r *Client) handleDisconnection() error {
	err := r.handleSendMessage(Disconnection, "Bye")
	if err != nil {
		return err
	}

	os.Exit(0)
	return nil
}
