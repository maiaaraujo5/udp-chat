package handler

import (
	"context"
	"fmt"
	"log"
	"net"
)

func (r *Server) handleDisconnection(parentCtx context.Context, remote *net.UDPAddr) error {
	if address, ok := r.connections[remote.String()]; ok {
		delete(r.connections, address.String())
	}

	if !r.thereAreStillActiveConnections() {
		err := r.flusher.Execute(parentCtx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *Server) thereAreStillActiveConnections() bool {
	log.Println(fmt.Sprintf("Ainda há %d conexões ativas", len(r.connections)))
	return len(r.connections) > 0
}
