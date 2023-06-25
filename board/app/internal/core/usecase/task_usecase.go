package usecase

import (
	"board/internal/core/domain"
	"board/internal/core/dto"
	"board/internal/core/port/driven"
	"board/pkg/auth"
	"fmt"

	"github.com/google/uuid"

)

type TaskUsecase interface {
	CreateTask(dto.CreateTaskRequest) (*dto.CreateTaskResponse, error)
	DeleteTask(dto.DeleteTaskRequest) error
	UpdateTask(dto.UpdateTaskRequest) (*dto.UpdateTaskResponse, error)
	AssignTask(dto.AssignTaskRequest) error
}

type taskUsecase struct {
	taskDao driven.TaskPortDriven
}

func NewTaskUsecase(taskDao driven.TaskPortDriven) TaskUsecase {
	return &taskUsecase{
		taskDao: taskDao,
	}
}

func (t *taskUsecase) CreateTask(task dto.CreateTaskRequest) (*dto.CreateTaskResponse, error) {

	id, err := auth.IsCustomer(task.AccessToken)

	if err != nil {
		return nil, err
	}

	//check permission
	fmt.Println(id)

	newTask, err := t.taskDao.Create(domain.Task{
		ID:        uuid.New(),
		Title:     task.Title,
		Order:     task.Order,
		SectionID: task.SectionID,
		CreatorID: id,
	})

	return &dto.CreateTaskResponse{
		ID:        newTask.ID,
		Title:     task.Title,
		Order:     task.Order,
		SectionID: task.SectionID,
	}, nil

}

func (t *taskUsecase) DeleteTask(task dto.DeleteTaskRequest) error {

	id, err := auth.IsCustomer(task.AccessToken)

	if err != nil {
		return err
	}

	//check permission
	fmt.Println(id)

	err = t.taskDao.Delete(task.TaskID)

	if err != nil {
		return err
	}

	return nil
}

func (t *taskUsecase) UpdateTask(task dto.UpdateTaskRequest) (*dto.UpdateTaskResponse, error) {

	id, err := auth.IsCustomer(task.AccessToken)

	if err != nil {
		return nil, err
	}

	//check permission
	fmt.Println(id)

	taskData, err := t.taskDao.Update(domain.Task{
		ID:        task.ID,
		Title:     task.Title,
		Order:     task.Order,
		SectionID: task.SectionID,
	})

	return &dto.UpdateTaskResponse{
		ID:    taskData.ID,
		Title: taskData.Title,
		// SectionID: taskData.SectionId,
		// Order:     taskData.Order,
	}, nil
}

// TODO:
func (t *taskUsecase) AssignTask(task dto.AssignTaskRequest) error {

	id, err := auth.IsCustomer(task.AccessToken)

	if err != nil {
		return err
	}

	//check permission
	fmt.Println(id)

	return nil
}
