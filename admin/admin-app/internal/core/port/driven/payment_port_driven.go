package driven

import (
	"admin/internal/core/domain"

	"github.com/google/uuid"
)

type PaymentPortDriven interface {
	Get(uuid.UUID) (*domain.Payment, error)
	Index() []domain.Payment
}
