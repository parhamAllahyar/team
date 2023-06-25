package postgres

import (
	"board/internal/core/domain"
	"board/internal/core/port/driven"
	"database/sql"

	"github.com/google/uuid"

)

type taskRepo struct {
	taskPort driven.TaskPortDriven
	db       *sql.DB
}

func NewTaskRepo(db *sql.DB) *taskRepo {
	return &taskRepo{
		db: db,
	}
}

func (t *taskRepo) GetById(id uuid.UUID) (*domain.Task, error) {
	var task domain.Task
	query := callGorm(t.db).Table("tasks").Where("id = ?", id).Find(&task)
	if query.Error != nil {
		return nil, query.Error
	}
	return &task, nil
}

func (t *taskRepo) Index() []domain.Task {
	return []domain.Task{}
}

func (t *taskRepo) Create(newTask domain.Task) (*domain.Task, error) {
	query := callGorm(t.db).Table("tasks").Create(&newTask).Find(&newTask)
	if query.Error != nil {
		return nil, query.Error
	}
	return &newTask, nil
}

func (t *taskRepo) Update(task domain.Task) (*domain.Task, error) {
	query := callGorm(t.db).Table("tasks").Where("id = ?", task.ID).Updates(task).Find(&task)
	if query.Error != nil {
		return nil, query.Error
	}
	return &task, nil
}

func (t *taskRepo) Delete(id uuid.UUID) error {
	query := callGorm(t.db).Table("tasks").Where("id = ?", id).Delete(&domain.Task{})
	return query.Error
}
