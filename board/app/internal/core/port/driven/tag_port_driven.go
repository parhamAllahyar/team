package driven

import (
	"board/internal/core/domain"

	"github.com/google/uuid"
)

type TagPortDriven interface {
	GetById(uuid.UUID) (*domain.Tag, error)
	Index() []domain.Tag
	Create(domain.Tag) (*domain.Tag, error)
	Delete(uuid.UUID) error
	Update(domain.Tag) (*domain.Tag, error)
}
