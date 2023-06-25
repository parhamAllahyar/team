package domain

import (
	"time"

	"github.com/google/uuid"

)

type Admin struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
	Email     string
	Phone     string
	Password  string
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
