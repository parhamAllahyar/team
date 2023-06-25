package driver

import (
	"workspace/internal/core/domain"

)

type WorkspacePortDriver interface {
	Get(id uint) (domain.Workspace, error)
	Index() []domain.Workspace
}
