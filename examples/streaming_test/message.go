package streaming

import (
	"time"

	"github.com/brianvoe/gofakeit/v5"
)

// Message to be published.
type Message struct {
	Timestamp int64  `json:"timestamp"`
	Message   string `json:"message"`
}

// GenerateMessage generates a message with current timestamp and random phrase.
func GenerateMessage() *Message {
	return &Message{
		Timestamp: time.Now().Unix(),
		Message:   gofakeit.HackerPhrase(),
	}
}
