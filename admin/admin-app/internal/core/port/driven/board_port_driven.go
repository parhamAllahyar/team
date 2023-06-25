package driven

import (
	"admin/internal/core/domain"

	"github.com/google/uuid"
)

type BoardPortDriven interface {
	GetBoard(uuid.UUID) (*domain.Board, error)
	IndexBoard() []domain.Board
}
