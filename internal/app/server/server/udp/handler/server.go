package handler

import (
	"context"
	"encoding/json"
	"github.com/maiaaraujo5/gostart/log/logger"
	"github.com/maiaaraujo5/udp-chat/internal/app/server/domain/model"
	"github.com/maiaaraujo5/udp-chat/internal/app/server/domain/service"
	"github.com/maiaaraujo5/udp-chat/internal/app/server/server/udp/model/in"
	"github.com/maiaaraujo5/udp-chat/internal/app/server/server/udp/model/out"
	"net"
)

type Server struct {
	connections    map[string]net.Addr
	messageRecover service.Recover
	saveMessage    service.Saver
	deleteMessage  service.DeleteMessage
	flusher        service.Flusher
	conn           net.PacketConn
}

func NewServer(messageRecover service.Recover, saver service.Saver, deleteMessage service.DeleteMessage,
	flusher service.Flusher, conn *net.UDPConn) *Server {

	return &Server{
		connections:    make(map[string]net.Addr),
		messageRecover: messageRecover,
		saveMessage:    saver,
		deleteMessage:  deleteMessage,
		flusher:        flusher,
		conn:           conn,
	}
}

func (r *Server) Handle(parentCtx context.Context) error {
	logger.Info("Server started...")
	for {
		message := make([]byte, 1024)
		rlen, remote, err := r.conn.ReadFrom(message[:])
		if err != nil {
			return err
		}

		req, err := r.DecodeJsonIntoModel(message, rlen)
		if err != nil {
			return err
		}

		err = r.handle(parentCtx, req, remote)
		if err != nil {
			return err
		}
	}
}

func (r *Server) handle(parentCtx context.Context, message *in.In, remote net.Addr) error {
	switch message.Action {
	case "NEW_CONNECTION":
		err := r.handleNewConnection(parentCtx, message, remote)
		if err != nil {
			return err
		}
	case "NEW_MESSAGE":
		err := r.handleNewMessage(parentCtx, message, remote)
		if err != nil {
			return err
		}
	case "DELETE_MESSAGE":
		err := r.handleDeleteMessage(parentCtx, message, remote)
		if err != nil {
			return err
		}
	case "DISCONNECT":
		err := r.handleDisconnection(parentCtx, remote)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *Server) broadcastMessage(message *model.Message, remote net.Addr) error {
	outMessage := out.Out{
		ID:      message.ID,
		UserID:  message.UserID,
		Message: message.Message,
	}
	bytes, err := json.Marshal(outMessage)
	if err != nil {
		return err
	}

	for _, connection := range r.connections {
		if connection != remote {
			_, err = r.conn.WriteTo(bytes, connection)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (r *Server) DecodeJsonIntoModel(message []byte, rlen int) (*in.In, error) {
	req := &in.In{}
	err := json.Unmarshal(message[:rlen], req)
	if err != nil {
		return nil, err
	}

	return req, nil
}
