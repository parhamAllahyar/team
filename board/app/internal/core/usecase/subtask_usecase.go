package usecase

import (
	"board/internal/core/domain"
	"board/internal/core/dto"
	"board/internal/core/port/driven"
	"board/pkg/auth"
	"fmt"

	"github.com/google/uuid"

)

type SubTaskUsecase interface {
	CreateSubTask(dto.CreateSubTaskRequest) (*dto.CreateSubTaskResponse, error)
	DeleteSubTask(dto.DeleteSubTaskRequest) error
	UpdateSubTask(dto.UpdateSubTaskRequest) (*dto.UpdateSubTaskResponse, error)
}

type subTaskUsecase struct {
	subTaskDao driven.SubTaskPortDriven
}

func NewSubTaskUsecase(subTaskDao driven.SubTaskPortDriven) SubTaskUsecase {
	return &subTaskUsecase{
		subTaskDao: subTaskDao,
	}
}

func (s *subTaskUsecase) CreateSubTask(subTask dto.CreateSubTaskRequest) (*dto.CreateSubTaskResponse, error) {
	id, err := auth.IsCustomer(subTask.AccessToken)

	if err != nil {
		return nil, err
	}

	//check permission
	fmt.Println(id)

	newSubtask, err := s.subTaskDao.Create(domain.SubTask{
		ID:     uuid.New(),
		TaskID: subTask.TsakID,
		Title:  subTask.Title,
	})

	if err != nil {
		return nil, err
	}

	return &dto.CreateSubTaskResponse{
		TaskID: newSubtask.TaskID,
		Title:  newSubtask.Title,
		ID:     newSubtask.ID,
	}, nil
}

func (s *subTaskUsecase) DeleteSubTask(subTask dto.DeleteSubTaskRequest) error {
	id, err := auth.IsCustomer(subTask.AccessToken)

	if err != nil {
		return err
	}

	//check permission
	fmt.Println(id)

	return nil
}

func (s *subTaskUsecase) UpdateSubTask(subTask dto.UpdateSubTaskRequest) (*dto.UpdateSubTaskResponse, error) {

	id, err := auth.IsCustomer(subTask.AccessToken)

	if err != nil {
		return nil, err
	}

	//check permission
	fmt.Println(id)

	subTaskData, err := s.subTaskDao.Update(domain.SubTask{
		ID:     uuid.New(),
		TaskID: subTask.TaskID,
		Title:  subTask.Title,
	})

	if err != nil {
		return nil, err
	}

	return &dto.UpdateSubTaskResponse{
		TaskID: subTaskData.TaskID,
		Title:  subTaskData.Title,
		ID:     subTaskData.ID,
	}, nil
}
