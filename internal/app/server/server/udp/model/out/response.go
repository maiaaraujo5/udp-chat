package out

import (
	"github.com/maiaaraujo5/udp-chat/internal/app/server/domain/model"
	"time"
)

type Out struct {
	ID      string    `json:"id,omitempty"`
	UserID  string    `json:"user_id,omitempty"`
	Message string    `json:"message,omitempty"`
	Time    time.Time `json:"time"`
}

func FromModel(message model.Message) *Out {
	return &Out{
		ID:      message.ID,
		UserID:  message.UserID,
		Message: message.Message,
		Time:    message.Time,
	}
}
