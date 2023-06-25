package domain

import (
	"time"

	"github.com/google/uuid"

)

type Board struct {
	ID          uuid.UUID
	WorkspaceID uuid.UUID
	Title       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

func NewBoard(id uuid.UUID, WorkspaceId uuid.UUID, Title string, CreatedAt time.Time) error {
	return nil
}
