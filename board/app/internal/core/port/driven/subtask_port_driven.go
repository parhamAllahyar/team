package driven

import (
	"board/internal/core/domain"

	"github.com/google/uuid"
)

type SubTaskPortDriven interface {
	GetById(uuid.UUID) (*domain.SubTask, error)
	Index() []domain.SubTask
	Create(domain.SubTask) (*domain.SubTask, error)
	Update(domain.SubTask) (*domain.SubTask, error)
	Delete(uuid.UUID) error
}
