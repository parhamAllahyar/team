package grpc

import (
	"customer/internal/core/domain"
	"customer/internal/core/port/driven"
	//	"database/sql"
	"fmt"
)

type adminRepo struct {
	adminPort driven.CustomerPortDriven
}

func NewAdminRepo() *adminRepo {
	return &adminRepo{}
}

func (a *adminRepo) Create(domain.Customer) error {
	return nil
}
func (a *adminRepo) Update(domain.Customer) error {
	return nil
}
func (a *adminRepo) Delete(id uint) error {
	return nil
}
func (a *adminRepo) Get(id uint) (domain.Customer, error) {
	return domain.Customer{}, nil
}
func (a *adminRepo) Index() []domain.Customer {
	return []domain.Customer{}

}
func (a *adminRepo) Test() {
	fmt.Println("hello from grpc repo")
}
