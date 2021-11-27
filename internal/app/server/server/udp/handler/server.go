package handler

import (
	"context"
	"encoding/json"
	"github.com/maiaaraujo5/udp-chat/internal/app/server/domain/service"
	"github.com/maiaaraujo5/udp-chat/internal/app/server/server/udp/model/in"
	"log"
	"net"
)

type Server struct {
	connections    map[string]*net.UDPAddr
	messageRecover service.Recover
	saveMessage    service.Saver
	deleteMessage  service.DeleteMessage
	flusher        service.Flusher
	conn           *net.UDPConn
}

func NewServer(messageRecover service.Recover, saver service.Saver, deleteMessage service.DeleteMessage,
	flusher service.Flusher, conn *net.UDPConn) *Server {

	return &Server{
		connections:    make(map[string]*net.UDPAddr),
		messageRecover: messageRecover,
		saveMessage:    saver,
		deleteMessage:  deleteMessage,
		flusher:        flusher,
		conn:           conn,
	}
}

func (r *Server) Handle(parentCtx context.Context) error {
	log.Println("handling...")

	for {
		message := make([]byte, 1024)
		rlen, remote, err := r.conn.ReadFromUDP(message[:])
		if err != nil {
			panic(err)
		}

		if r.isNewConnection(remote) {
			err := r.sendOldMessages(parentCtx, remote, r.conn)
			if err != nil {
				return err
			}
		}

		req, err := r.DecodeJsonIntoModel(message, rlen)
		if err != nil {
			return err
		}

		switch req.Action {
		case "NEW_MESSAGE":
			err := r.handleNewMessage(parentCtx, req, remote)
			if err != nil {
				return err
			}
		case "DELETE_MESSAGE":
			err := r.handleDeleteMessage(parentCtx, req, remote)
			if err != nil {
				return err
			}
		case "DISCONNECT":
			err := r.handleDisconnection(parentCtx, remote)
			if err != nil {
				return err
			}
		}
	}
}

func (r *Server) DecodeJsonIntoModel(message []byte, rlen int) (*in.In, error) {
	req := &in.In{}
	err := json.Unmarshal(message[:rlen], req)
	if err != nil {
		return nil, err
	}

	return req, nil
}
