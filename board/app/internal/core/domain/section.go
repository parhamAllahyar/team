package domain

import (
	"time"

	"github.com/google/uuid"

)

type Section struct {
	ID        uuid.UUID
	BoardID   uuid.UUID
	Title     string
	Order     uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
