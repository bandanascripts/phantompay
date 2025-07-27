package structs

import "time"

type PhantomPayTransaction struct {
	ID         string
	SenderID   string
	ReceiverID string
	Amount     float64
	// Type       string
	CreatedAt  time.Time
}
