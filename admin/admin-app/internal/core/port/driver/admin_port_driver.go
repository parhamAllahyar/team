package driver

import (
	"admin/internal/core/domain"
)

type AdminPortDriver interface {
	Create(domain.Admin) error
	Update(domain.Admin) error
	Delete(id uint) error
	Get(id uint) (domain.Admin, error)
	Index() []domain.Admin
}
