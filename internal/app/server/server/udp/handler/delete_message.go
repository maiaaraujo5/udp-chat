package handler

import (
	"context"
	"fmt"
	"net"

	"github.com/maiaaraujo5/udp-chat/internal/app/server/domain/model"
	"github.com/maiaaraujo5/udp-chat/internal/app/server/server/udp/model/in"
)

func (r *Server) handleDeleteMessage(parentCtx context.Context, in *in.In, remote net.Addr) error {
	fmt.Printf("received %s from %s to delete", in, remote)

	message := &model.Message{
		ID:     in.Message,
		UserID: remote.String(),
	}

	err := r.broadcastMessage(message, remote)
	if err != nil {
		return err
	}

	err = r.deleteMessage.Delete(parentCtx, message)
	if err != nil {
		return err
	}

	return nil
}
