package driver

import (
	"board/internal/core/domain"

)

type SectionPortDriver interface {
	Get(id uint) (domain.Section, error)
	Index() []domain.Section
}
