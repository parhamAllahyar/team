package postgres

import (
	"board/internal/core/domain"
	"board/internal/core/port/driven"
	"database/sql"

	"github.com/google/uuid"
)

type subTaskRepo struct {
	boardPort driven.SubTaskPortDriven
	db        *sql.DB
}

func NewSubTaskdRepo(db *sql.DB) *subTaskRepo {
	return &subTaskRepo{
		db: db,
	}
}

func (s *subTaskRepo) GetById(id uuid.UUID) (*domain.SubTask, error) {
	var subTask domain.SubTask
	query := callGorm(s.db).Table("sub_tasks").Where("id = ?", id).Find(&subTask)
	if query.Error != nil {
		return nil, query.Error
	}
	return &subTask, nil
}

func (s *subTaskRepo) Index() []domain.SubTask {
	return []domain.SubTask{}
}

func (s *subTaskRepo) Create(subTask domain.SubTask) (*domain.SubTask, error) {
	query := callGorm(s.db).Table("tasks").Create(&subTask).Find(&subTask)
	if query.Error != nil {
		return nil, query.Error
	}
	return &subTask, nil
}

func (s *subTaskRepo) Update(subTask domain.SubTask) (*domain.SubTask, error) {
	query := callGorm(s.db).Table("sub_tasks").Where("id = ?", subTask.ID).Updates(subTask).Find(&subTask)
	if query.Error != nil {
		return nil, query.Error
	}
	return &subTask, nil
}

func (s *subTaskRepo) Delete(id uuid.UUID) error {
	query := callGorm(s.db).Table("sub_tasks").Where("id = ?", id).Delete(&domain.SubTask{})
	return query.Error
}
