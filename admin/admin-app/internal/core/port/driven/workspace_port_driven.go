package driven

import (
	"admin/internal/core/domain"

	"github.com/google/uuid"
)

type WorkspacePortDriven interface {
	Get(uuid.UUID) (*domain.Workspace, error)
	Index() []domain.Workspace
}
