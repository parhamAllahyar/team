package dto

import "github.com/google/uuid"

type CreateBoardRequest struct {
	AccessToken string
	Title       string
	WorkspaceID uuid.UUID
}
type CreateBoardResponse struct {
	ID          uuid.UUID
	Title       string
	WorkspaceID uuid.UUID
}
type DeleteBoardRequest struct {
	ID          uuid.UUID
	AccessToken string
}
type UpdateBoardRequest struct {
	AccessToken string
	ID          uuid.UUID
	Title       string
	WorkspaceID uuid.UUID
}
type UpdateBoardResponse struct {
	ID          uuid.UUID
	Title       string
	WorkspaceId uuid.UUID
}
