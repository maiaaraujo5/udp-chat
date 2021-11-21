package handler

import (
	"context"
	"encoding/json"
	"github.com/maiaaraujo5/udp-chat/internal/app/server/server/udp/model/out"
	"log"
	"net"
)

func (r *Server) isNewConnection(remote *net.UDPAddr) bool {
	_, ok := r.connections[remote.String()]
	if !ok {
		r.connections[remote.String()] = remote
	}
	return !ok
}

func (r *Server) sendOldMessages(parentCtx context.Context, remote *net.UDPAddr, conn *net.UDPConn) error {
	log.Println("sending old messages")

	messages, err := r.messageRecover.Execute(parentCtx)
	if err != nil {
		return err
	}

	for _, message := range messages {

		outMessage, err := json.Marshal(out.FromModel(message))
		if err != nil {
			return err
		}

		_, err = conn.WriteToUDP(outMessage, remote)
		if err != nil {
			log.Println(err)
		}
	}

	return nil
}
