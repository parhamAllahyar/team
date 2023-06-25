package postgres

import (
	"board/internal/core/domain"
	"board/internal/core/port/driven"
	"database/sql"

	"github.com/google/uuid"

)

type sectionRepo struct {
	sectionPort driven.SectionPortDriven
	db          *sql.DB
}

func NewSectionRepo(db *sql.DB) *sectionRepo {
	return &sectionRepo{
		db: db,
	}
}

func (s *sectionRepo) GetById(id uuid.UUID) (*domain.Section, error) {
	var section domain.Section
	query := callGorm(s.db).Table("sub_tasks").Where("id = ?", id).Find(&section)
	if query.Error != nil {
		return nil, query.Error
	}
	return &section, nil
}

func (s *sectionRepo) Index() []domain.Section {
	return []domain.Section{}
}

func (s *sectionRepo) Create(newSection domain.Section) (*domain.Section, error) {
	query := callGorm(s.db).Table("boards").Create(&newSection).Find(&newSection)
	if query.Error != nil {
		return nil, query.Error
	}
	return &newSection, nil
}

func (s *sectionRepo) Update(section domain.Section) (*domain.Section, error) {
	query := callGorm(s.db).Table("section").Where("id = ?", section.ID).Updates(section).Find(&section)
	if query.Error != nil {
		return nil, query.Error
	}
	return &section, nil
}

func (s *sectionRepo) Delete(id uuid.UUID) error {
	query := callGorm(s.db).Table("sections").Where("id = ?", id).Delete(&domain.Section{})
	return query.Error
}
