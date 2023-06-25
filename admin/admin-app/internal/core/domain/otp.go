package domain

import (
	"time"

	"github.com/google/uuid"
)

// TODO: change type to enume
type OTP struct {
	ID        uuid.UUID
	Type      string
	NewPhone  string
	OTP       int
	Status    string
	ExpireAt  time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
