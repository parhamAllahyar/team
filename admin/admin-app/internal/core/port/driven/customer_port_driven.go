package driven

import (
	"admin/internal/core/domain"

	"github.com/google/uuid"
)

type CustomerPortDriven interface {
	Update(domain.Customer) error
	Get(uuid.UUID) (*domain.Customer, error)
	Index() []domain.Customer
}
