package handler

import (
	"context"
	"fmt"
	"github.com/maiaaraujo5/udp-chat/internal/app/server/domain/model"
	"github.com/maiaaraujo5/udp-chat/internal/app/server/server/udp/model/in"
	"net"
)

func (r *Server) handleDeleteMessage(parentCtx context.Context, in *in.In, remote *net.UDPAddr) error {
	fmt.Printf("received %s from %s to delete", in, remote)

	message := &model.Message{
		ID: in.Message,
	}

	err := r.broadcastMessage(message, remote)
	if err != nil {
		return err
	}

	return nil
}
