package handler

func (r *Client) handleDeleteMessage(msg string) error {

	messages, err := r.deleter.Delete(r.messages, msg, r.conn.LocalAddr().String())
	if err != nil {
		return err
	}

	r.messages = messages

	err = r.handleSendMessage(DeleteMessage, msg)
	if err != nil {
		return err
	}

	r.print()

	return nil

}
