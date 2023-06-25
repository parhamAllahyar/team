package usecase

import (
	"admin/internal/core/domain"
	"admin/internal/core/dto"
	"admin/internal/core/port/driven"
	"admin/pkg/auth/authenticator"
	"admin/pkg/auth/encryption"
	errormessage "admin/pkg/error-message"
	"errors"
	"time"

	"github.com/google/uuid"

)

type AuthUsecase interface {
	SignIn(dto.SignInRequest) (*dto.AuthResponse, error)
	SignUp(dto.SignUpRequest) (*dto.AuthResponse, error)
}

type authUsecase struct {
	authDao driven.AdminPortDriven
	linkDao driven.LinkPortDriven
}

func NewAuthUsecase(authDao driven.AdminPortDriven, linkDao driven.LinkPortDriven) AuthUsecase {
	return &authUsecase{
		authDao: authDao,
		linkDao: linkDao,
	}
}

func (a *authUsecase) SignIn(request dto.SignInRequest) (*dto.AuthResponse, error) {

	//TODO: validation

	query, err := a.authDao.GetByEmail(request.Email)

	//TODO: check
	if query == nil && query.Status == "active" {
		return nil, errors.New(errormessage.EmailOrPasswordIsWrong)
	}

	if err != nil {
		return nil, errors.New(errormessage.DatabaseError)
	}

	isPasswordTrue := encryption.ComparePasswords(query.Password, request.Password)

	if isPasswordTrue == false {
		return nil, errors.New(errormessage.EmailOrPasswordIsWrong)
	}

	//Generate Token
	token, err := authenticator.ExtractClaims(query.ID)

	if err != nil {
		return nil, err
	}

	return &dto.AuthResponse{AccessToken: token}, nil
}

func (a *authUsecase) SignUp(request dto.SignUpRequest) (*dto.AuthResponse, error) {

	//TODO: check
	//TODO: validation

	link, err := a.linkDao.GetByLink(request.Link)

	if err != nil {
		return nil, err
	}

	if link.ExpiredAt.Before(time.Now()) || (link == nil) || link.Status != "Active" {
		return nil, errors.New(errormessage.LinkNotFound)
	}

	hashedValue, err := encryption.HashPassword(request.Password)

	newAdmin, err := a.authDao.Create(domain.Admin{
		ID:        uuid.New(),
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     link.Email,
		Phone:     request.Phone,
		Password:  hashedValue,
		Status:    "Active",
	})

	if err != nil {
		return nil, err
	}

	token, err := authenticator.ExtractClaims(newAdmin.ID)

	if err != nil {
		return nil, err
	}

	return &dto.AuthResponse{AccessToken: token}, nil

}

//TODO: verify phone
