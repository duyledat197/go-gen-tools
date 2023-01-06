package kafka_utils

import "github.com/google/uuid"

type Topic struct {
	Name string    `json:"name"`
	ID   uuid.UUID `json:"id"`
}

type Message struct {
}
