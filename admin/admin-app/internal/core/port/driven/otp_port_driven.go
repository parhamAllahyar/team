package driven

import (
	"admin/internal/core/domain"

)

type OtpPortDriven interface {
	GetById(string) (*domain.OTP, error)
	Update(domain.OTP) (*domain.OTP, error)
	Create(domain.OTP) (*domain.OTP, error)
}
