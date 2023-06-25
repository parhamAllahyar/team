package grpc

import (
	"admin/internal/core/domain"
	"admin/internal/core/port/driven"
	//	"database/sql"
	"fmt"
)

type adminRepo struct {
	adminPort driven.AdminPortDriven
}

func NewAdminRepo() *adminRepo {
	return &adminRepo{}
}

func (a *adminRepo) Create(domain.Admin) error {
	return nil
}
func (a *adminRepo) Update(domain.Admin) error {
	return nil
}
func (a *adminRepo) Delete(id uint) error {
	return nil
}
func (a *adminRepo) Get(id uint) (domain.Admin, error) {
	return domain.Admin{}, nil
}
func (a *adminRepo) Index() []domain.Admin {
	return []domain.Admin{}

}
func (a *adminRepo) Test() {
	fmt.Println("hello from grpc repo")
}
