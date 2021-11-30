package out

import (
	"time"
)

type Out struct {
	ID      string    `json:"id,omitempty"`
	UserID  string    `json:"user_id,omitempty"`
	Message string    `json:"message,omitempty"`
	Time    time.Time `json:"time"`
}
