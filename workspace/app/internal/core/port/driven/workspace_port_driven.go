package driven

import (
	"workspace/internal/core/domain"

	"github.com/google/uuid"
)

type WorkspacePortDriven interface {
	Get(id uuid.UUID) (*domain.Workspace, error)
	Index() ([]domain.Workspace, error)
	Create(workspace domain.Workspace) (*domain.Workspace, error)
	Update(id uuid.UUID, title string) (*domain.Workspace, error)
	CustomerWorkspacesIndex(id uuid.UUID) ([]domain.Workspace, error)
	WorkspaceMembers(id uuid.UUID) ([]domain.Customer, error)
	IsAdmin(customerId uuid.UUID, workspacId uuid.UUID) bool
}
