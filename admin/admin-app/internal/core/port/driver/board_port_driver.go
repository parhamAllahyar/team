package driver

import (
	"admin/internal/core/domain"

)

type BoardPortDriver interface {
	Get(id uint) (domain.Board, error)
	Index() []domain.Board
}
