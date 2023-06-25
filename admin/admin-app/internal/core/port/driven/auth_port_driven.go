package driven

import (
	"admin/internal/core/domain"

	"github.com/google/uuid"
)

type AuthPortDriven interface {
	Create(domain.Admin) error
	Update(domain.Admin) error
	Delete(id uint) error
	Get(uuid.UUID) (*domain.Admin, error)
	Index() []domain.Admin
	Test() string
}
