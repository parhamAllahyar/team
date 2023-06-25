package driven

import (
	"admin/internal/core/domain"

	"github.com/google/uuid"
)

type AdminPortDriven interface {
	Create(domain.Admin) (*domain.Admin, error)
	Update(domain.Admin) (*domain.Admin, error)
	Delete(uuid.UUID) error
	GetById(uuid.UUID) (*domain.Admin, error)
	GetByEmail(string) (*domain.Admin, error)
	Index() []domain.Admin
}
