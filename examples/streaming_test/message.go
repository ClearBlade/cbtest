package streaming

import (
	"time"

	"github.com/brianvoe/gofakeit/v5"
)

type Message struct {
	Timestamp int64  `json:"timestamp"`
	Message   string `json:"message"`
}

func GenerateMessage() *Message {
	return &Message{
		Timestamp: time.Now().Unix(),
		Message:   gofakeit.HackerPhrase(),
	}
}
