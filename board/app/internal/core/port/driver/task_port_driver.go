package driver

import (
	"board/internal/core/domain"

)

type TaskPortDriver interface {
	Get(id uint) (domain.Task, error)
	Index() ([]domain.Task, error)
	Create(domain.Task) (domain.Task, error)
	Update(domain.Task) (domain.Task, error)
}
