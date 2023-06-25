package driver

import (
	"board/internal/core/domain"
)

type BoardPortDriver interface {
	Get(id uint) (domain.Board, error)
	Index() []domain.Board
}
