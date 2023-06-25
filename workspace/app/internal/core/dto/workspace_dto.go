package dto

import "github.com/google/uuid"

type CreateWorkspaceRequest struct {
	Title       string
	AccessToken string
}

type CreateWorkspaceResponse struct {
	ID    uuid.UUID
	Title string
}

type UpdateWorkspaceRequest struct {
	Title       string
	AccessToken string
	ID          uuid.UUID
}

type UpdateWorkspaceResponse struct {
	ID    uuid.UUID
	Title string
}

type WorkspaceMembersResponse struct {
}

type WorkspaceMembersRequest struct {
	AccessToken string
	ID          uuid.UUID
}
