package domain

import (
	"time"

	"github.com/google/uuid"
)

// TODO: change type to enume
type Link struct {
	ID        uuid.UUID
	Type      string
	Email     string
	Link      string
	Status    string
	ExpiredAt time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
