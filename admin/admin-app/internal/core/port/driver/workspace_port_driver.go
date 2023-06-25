package driver

import (
	"admin/internal/core/domain"
)

type WorkspacePortDriver interface {
	Get(id uint) (domain.Workspace, error)
	Index() []domain.Workspace
}
