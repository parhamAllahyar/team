package driven

import (
	"admin/internal/core/domain"
)

type LinkPortDriven interface {
	GetById(string) (*domain.Link, error)
	Update(domain.Link) (*domain.Link, error)
	Create(domain.Link) (*domain.Link, error)
	GetByLink(string) (*domain.Link, error)
}
