package service

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/maiaaraujo5/udp-chat/internal/app/client/domain/model/out"
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
		ID:      c.generateID(),
		Action:  action,
		Message: message,
	}
}

func (c *creator) generateID() string {
	rand.Seed(time.Now().Unix())
	number := rand.Intn(9999) //nolint
	return strconv.Itoa(number)
}
