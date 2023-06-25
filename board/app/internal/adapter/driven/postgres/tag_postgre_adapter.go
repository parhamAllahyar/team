package postgres

import (
	"board/internal/core/domain"
	"board/internal/core/port/driven"
	"database/sql"

	"github.com/google/uuid"

)

type tagRepo struct {
	tagPort driven.TagPortDriven
	db      *sql.DB
}

func NewTagRepo(db *sql.DB) *tagRepo {
	return &tagRepo{
		db: db,
	}
}

func (t *tagRepo) GetById(id uuid.UUID) (*domain.Tag, error) {
	var tag domain.Tag
	query := callGorm(t.db).Table("tags").Where("id = ?", id).Find(&tag)
	if query.Error != nil {
		return nil, query.Error
	}
	return &tag, nil
}

func (t *tagRepo) Index() []domain.Tag {
	return []domain.Tag{}
}

func (t *tagRepo) Create(newTag domain.Tag) (*domain.Tag, error) {
	query := callGorm(t.db).Table("tags").Create(&newTag).Find(&newTag)
	if query.Error != nil {
		return nil, query.Error
	}
	return &newTag, nil
}

func (t *tagRepo) Update(tag domain.Tag) (*domain.Tag, error) {
	query := callGorm(t.db).Table("tags").Where("id = ?", tag.ID).Updates(tag).Find(&tag)
	if query.Error != nil {
		return nil, query.Error
	}
	return &tag, nil
}

func (t *tagRepo) Delete(id uuid.UUID) error {

	query := callGorm(t.db).Table("tags").Where("id = ?", id).Delete(&domain.Tag{})
	return query.Error
}
