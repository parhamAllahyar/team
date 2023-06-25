package driven

import (
	"board/internal/core/domain"

	"github.com/google/uuid"
)

type BoardPortDriven interface {
	Create(domain.Board) (*domain.Board, error)
	Update(domain.Board) (*domain.Board, error)
	GetById(uuid.UUID) (*domain.Board, error)
	Index() []domain.Board
	Delete(uuid.UUID) error
}
