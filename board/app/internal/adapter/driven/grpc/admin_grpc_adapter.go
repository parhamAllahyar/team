package grpc

import (
	"board/internal/core/domain"
	"board/internal/core/port/driven"
	//	"database/sql"
	"fmt"
)

type adminRepo struct {
	boardPort driven.BoardPortDriven
}

func NewAdminRepo() *adminRepo {
	return &adminRepo{}
}

func (a *adminRepo) Create(domain.Board) error {
	return nil
}
func (a *adminRepo) Update(domain.Board) error {
	return nil
}
func (a *adminRepo) Delete(id uint) error {
	return nil
}
func (a *adminRepo) Get(id uint) (domain.Board, error) {
	return domain.Board{}, nil
}
func (a *adminRepo) Index() []domain.Board {
	return []domain.Board{}

}
func (a *adminRepo) Test() {
	fmt.Println("hello from grpc repo")
}
