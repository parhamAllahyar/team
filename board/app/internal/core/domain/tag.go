package domain

import (
	"time"

	"github.com/google/uuid"

)

type Tag struct {
	ID        uuid.UUID
	Title     string
	Pattern   string
	Order     uint
	BoardID   uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
