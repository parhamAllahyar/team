package domain

import (
	"time"

	"github.com/google/uuid"

)

type Payment struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time

}
