package domain

import (
	"time"

	"github.com/google/uuid"

)

type Task struct {
	ID        uuid.UUID
	Title     string
	Order     uint
	SectionID uuid.UUID
	CreatorID uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
