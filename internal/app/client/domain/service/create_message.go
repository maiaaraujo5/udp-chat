package service

import (
	"github.com/maiaaraujo5/udp-chat/internal/app/client/domain/model/out"
	"math/rand"
	"strconv"
	"time"
)

type Creator interface {
	Create(action, message string) *out.Out
}

type creator struct {
}

func NewCreator() Creator {
	return &creator{}
}

func (c *creator) Create(action, message string) *out.Out {
	return &out.Out{
		ID:      c.generateId(),
		Action:  action,
		Message: message,
	}
}

func (c *creator) generateId() string {
	rand.Seed(time.Now().Unix())
	number := rand.Intn(9999)
	return strconv.Itoa(number)
}
