package models

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	Token     string
	Client    string
	IP        string
	CreatedAt time.Time
	LastUsed  time.Time
}
