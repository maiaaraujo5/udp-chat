package handler

import (
	"context"
	"fmt"
	"net"

	"github.com/maiaaraujo5/gostart/log/logger"
	"github.com/maiaaraujo5/udp-chat/internal/app/server/domain/model"
)

func (r *Server) handleDisconnection(parentCtx context.Context, remote net.Addr) error {
	message := &model.Message{Message: "Leave the room!", UserID: remote.String()}

	if address, ok := r.connections[remote.String()]; ok {
		delete(r.connections, address.String())

		err := r.broadcastMessage(message, remote)
		if err != nil {
			return err
		}
	}

	if !r.thereAreStillActiveConnections() {
		err := r.flusher.Flush(parentCtx)
		if err != nil {
			return err
		}

		return nil
	}

	err := r.saveMessage.Save(parentCtx, message)
	if err != nil {
		return err
	}

	return nil
}

func (r *Server) thereAreStillActiveConnections() bool {
	logger.Debug(fmt.Sprintf("Ainda há %d conexões ativas", len(r.connections)))
	return len(r.connections) > 0
}
