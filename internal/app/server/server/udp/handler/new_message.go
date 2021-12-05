package handler

import (
	"context"
	"fmt"
	"net"

	"github.com/maiaaraujo5/gostart/log/logger"
	"github.com/maiaaraujo5/udp-chat/internal/app/server/domain/model"
	"github.com/maiaaraujo5/udp-chat/internal/app/server/server/udp/model/in"
)

func (r *Server) handleNewMessage(parentCtx context.Context, in *in.In, remote net.Addr) error {
	message := &model.Message{
		ID:      in.ID,
		UserID:  remote.String(),
		Message: in.Message,
	}

	logger.Info(fmt.Sprintf("receiving new message from user %s", message.UserID))

	err := r.broadcastMessage(message, remote)
	if err != nil {
		return err
	}

	err = r.saveMessage.Save(parentCtx, message)
	if err != nil {
		return err
	}
	return nil
}
