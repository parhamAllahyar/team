package driver

import (
	"admin/internal/core/domain"
)

type PaymentPortDriver interface {
	Get(id uint) (domain.Payment, error)
	Index() []domain.Payment
}
