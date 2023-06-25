package dto

// CreateTask(token, title string, id string) (dto.CreateTaskResponse, error)
// DeleteTask(token string, id uuid.UUID) error
// UpdateTask(dto.UpdateTaskRequest) (dto.UpdateTaskResponse, error)

import (
	"github.com/google/uuid"

)

type CreateTaskRequest struct {
	AccessToken string
	Title       string
	SectionID   uuid.UUID
	Order       uint
}
type CreateTaskResponse struct {
	ID        uuid.UUID
	Title     string
	SectionID uuid.UUID
	Order     uint
}

type DeleteTaskRequest struct {
	AccessToken string
	TaskID      uuid.UUID
}

type UpdateTaskRequest struct {
	AccessToken string
	ID          uuid.UUID
	Title       string
	SectionID   uuid.UUID
	Order       uint
}
type UpdateTaskResponse struct {
	ID        uuid.UUID
	Title     string
	SectionID uuid.UUID
	Order     uint
}

type AssignTaskRequest struct {
	AccessToken string
	TaskID      uuid.UUID
	CustomerID  uuid.UUID
}
