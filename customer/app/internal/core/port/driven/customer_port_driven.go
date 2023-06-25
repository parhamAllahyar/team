package driven

import (
	"context"
	"customer/internal/core/domain"

	"github.com/google/uuid"

)

type CustomerPortDriven interface {
	Create(context.Context, domain.Customer) (*domain.Customer, error)
	Update(context.Context, domain.Customer) (*domain.Customer, error)
	Delete(context.Context, uuid.UUID) error
	GetByID(context.Context, uuid.UUID) (*domain.Customer, error)
	GetByEmail(context.Context, string) (*domain.Customer, error)
	Index(context.Context) ([]domain.Customer, error)
	UpdatePhone(context.Context, string) error
	UpdateEmail(context.Context, string) error
	UpdatePassword(context.Context, uuid.UUID, string) error
}
