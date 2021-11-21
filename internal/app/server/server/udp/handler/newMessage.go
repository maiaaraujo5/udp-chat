package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/maiaaraujo5/udp-chat/internal/app/server/domain/model"
	"github.com/maiaaraujo5/udp-chat/internal/app/server/server/udp/model/in"
	"github.com/maiaaraujo5/udp-chat/internal/app/server/server/udp/model/out"
	"net"
)

func (r *Server) handleNewMessage(parentCtx context.Context, in *in.In, remote *net.UDPAddr) error {
	fmt.Printf("received %s from %s", in, remote)

	message := &model.Message{
		UserID:  remote.String(),
		Message: in.Message,
	}

	err := r.saveMessage.Execute(parentCtx, message)

	if err != nil {
		return err
	}

	err = r.broadcastMessage(message, remote)
	if err != nil {
		return err
	}

	return nil
}

func (r *Server) broadcastMessage(message *model.Message, remote *net.UDPAddr) error {
	outMessage := out.Out{
		ID:      message.ID,
		UserID:  message.UserID,
		Message: message.Message,
		Time:    message.Time,
	}
	bytes, err := json.Marshal(outMessage)
	if err != nil {
		return err
	}

	for _, connection := range r.connections {
		if connection != remote {
			_, err = r.conn.WriteToUDP(bytes, connection)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
