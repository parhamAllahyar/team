package postgres

import (
	"admin/internal/core/domain"
	"admin/internal/core/port/driven"
	"database/sql"

	"github.com/google/uuid"

)

type adminRepo struct {
	authRepo driven.AuthPortDriven
	db       *sql.DB
}

func NewAdminRepo(db *sql.DB) *adminRepo {
	return &adminRepo{
		db: db,
	}
}

func (a *adminRepo) Create(admin domain.Admin) (*domain.Admin, error) {
	query := callGorm(a.db).Table("admins").Create(&admin).Find(&admin)
	if query.Error != nil {
		return nil, query.Error
	}
	return &admin, nil

}
func (a *adminRepo) Update(admin domain.Admin) (*domain.Admin, error) {
	query := callGorm(a.db).Table("admins").Where("id = ?", admin.ID).Updates(admin).Find(&admin)
	if query.Error != nil {
		return nil, query.Error
	}
	return &admin, nil
}
func (a *adminRepo) Delete(id uuid.UUID) error {
	query := callGorm(a.db).Table("admins").Where("id = ?", id).Delete(&domain.Admin{})
	return query.Error
}

func (a *adminRepo) GetById(id uuid.UUID) (*domain.Admin, error) {
	var admin domain.Admin
	query := callGorm(a.db).Table("admins").Where("id = ?", id).Find(&admin)
	if query.Error != nil {
		return nil, query.Error
	}
	return &admin, nil
}

func (a *adminRepo) GetByEmail(email string) (*domain.Admin, error) {
	var admin domain.Admin
	query := callGorm(a.db).Table("admins").Where("email = ?", email).Find(&admin)
	if query.Error != nil {
		return nil, query.Error
	}
	return &admin, nil
}
func (a *adminRepo) Index() []domain.Admin {
	return nil
}
