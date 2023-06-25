package driver

import (
	"admin/internal/core/domain"
)

type UserPortDriver interface {
	Update(domain.Customer) error

	Get(id uint) (domain.Customer, error)
	Index() []domain.Customer
}
