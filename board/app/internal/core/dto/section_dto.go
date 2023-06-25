package dto

import "github.com/google/uuid"

type CreateSectionRequest struct {
	AccessToken string
	BoardID     uuid.UUID
	Title       string
	Order       uint
}
type CreateSectionResponse struct {
	ID          uuid.UUID
	AccessToken string
	BoardID     uuid.UUID
	Title       string
	Order       uint
}
type DeleteSectionRequest struct {
	AccessToken string
	SectionID   uuid.UUID
}
type UpdateSectionRequest struct {
	AccessToken string
	BoardID     uuid.UUID
	Title       string
	Order       uint
	ID          uuid.UUID
}
type UpdateSectionResponse struct {
	ID      uuid.UUID
	BoardID uuid.UUID
	Title   string
	Order   uint
}
