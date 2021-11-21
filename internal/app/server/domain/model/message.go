package model

import "time"

type Message struct {
	ID      string
	Action  string
	UserID  string
	Message string
	Time    time.Time
}
