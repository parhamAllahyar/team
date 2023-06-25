package postgres

import (
	"admin/internal/core/domain"
	"admin/internal/core/port/driven"
	"database/sql"
	"fmt"

)

type linkRepo struct {
	linkRepo driven.LinkPortDriven
	db       *sql.DB
}

func NewLinkRepo(db *sql.DB) *linkRepo {
	return &linkRepo{
		db: db,
	}
}

func (l *linkRepo) GetById(id string) (*domain.Link, error) {
	var link domain.Link
	query := callGorm(l.db).Table("links").Where("id = ?", id).Find(&link)
	if query.Error != nil {
		return nil, query.Error
	}
	return &link, nil
}

func (l *linkRepo) GetByLink(link string) (*domain.Link, error) {
	var inviteLink domain.Link
	query := callGorm(l.db).Table("links").Where("link = ?", link).Find(&inviteLink)
	if query.Error != nil {
		return nil, query.Error
	}
	fmt.Println("link issssss:" , inviteLink)
	return &inviteLink, nil
}

func (l *linkRepo) Update(link domain.Link) (*domain.Link, error) {
	query := callGorm(l.db).Table("links").Where("id = ?", link.ID).Updates(link).Find(&link)
	if query.Error != nil {
		return nil, query.Error
	}
	return &link, nil
}

func (l *linkRepo) Create(link domain.Link) (*domain.Link, error) {
	query := callGorm(l.db).Table("links").Create(&link).Find(&link)
	if query.Error != nil {
		return nil, query.Error
	}
	return &link, nil

}
