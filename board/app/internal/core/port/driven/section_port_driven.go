package driven

import (
	"board/internal/core/domain"

	"github.com/google/uuid"

)

type SectionPortDriven interface {
	GetById(uuid.UUID) (*domain.Section, error)
	Index() []domain.Section
	Delete(uuid.UUID) error
	Update(domain.Section) (*domain.Section, error)
	Create(domain.Section) (*domain.Section, error)
}
