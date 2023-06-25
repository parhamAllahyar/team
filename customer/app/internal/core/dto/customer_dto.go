package dto

import (
	"time"

	"github.com/google/uuid"

)

type CreateCustomerRequest struct {
	FirstName string
	LastName  string
	Email     string
	Phone     string
	Password  string
}

type UpdateCustomerRequest struct {
	AccessToken string
	FirstName   string
	LastName    string
}

type Customer struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
	Email     string
	Phone     string
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
