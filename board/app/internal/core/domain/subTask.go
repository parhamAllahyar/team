package domain

import (
	"time"

	"github.com/google/uuid"

)

type SubTask struct {
	ID        uuid.UUID
	TaskID    uuid.UUID
	Title     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
