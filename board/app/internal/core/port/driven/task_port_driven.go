package driven

import (
	"board/internal/core/domain"

	"github.com/google/uuid"

)

type TaskPortDriven interface {
	GetById(uuid.UUID) (*domain.Task, error)
	Index() []domain.Task
	Create(domain.Task) (*domain.Task, error)
	Update(domain.Task) (*domain.Task, error)
	Delete(uuid.UUID) error
}
