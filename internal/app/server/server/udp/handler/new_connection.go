package handler

import (
	"context"
	"encoding/json"
	"log"
	"net"

	"github.com/maiaaraujo5/udp-chat/internal/app/server/server/udp/model/in"
	"github.com/maiaaraujo5/udp-chat/internal/app/server/server/udp/model/out"
)

func (r *Server) handleNewConnection(parentCtx context.Context, in *in.In, remote net.Addr) error {
	if r.isNewConnection(remote) {
		err := r.sendOldMessages(parentCtx, remote, r.conn)
		if err != nil {
			return err
		}

		err = r.handleNewMessage(parentCtx, in, remote)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *Server) isNewConnection(remote net.Addr) bool {
	_, ok := r.connections[remote.String()]
	if !ok {
		r.connections[remote.String()] = remote
	}
	return !ok
}

func (r *Server) sendOldMessages(parentCtx context.Context, remote net.Addr, conn net.PacketConn) error {
	log.Println("sending old messages")

	messages, err := r.messageRecover.Recover(parentCtx)
	if err != nil {
		return err
	}

	for _, message := range messages {

		outMessage, err := json.Marshal(out.FromModel(message))
		if err != nil {
			return err
		}

		_, err = conn.WriteTo(outMessage, remote)
		if err != nil {
			log.Println(err)
		}
	}

	return nil
}
