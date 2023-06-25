package usecase

import (
	"context"
	"customer/internal/core/domain"
	"customer/internal/core/dto"
	"customer/internal/core/port/driven"
	"customer/pkg/auth/authenticator"
	"customer/pkg/auth/encryption"
	"customer/pkg/message"
	//"customer/pkg/validation"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"

)

type AuthUsecase interface {
	SignUp(context.Context, dto.SignUpRequest) (*dto.AuthResponse, error)
	//ForgetPassword(customer dto.ForgetPassword) (dto.AuthResponse, error)
	SignIn(context.Context, dto.SignInRequest) (*dto.AuthResponse, error)
	UpdatePassword(context.Context, dto.UpdatePasswordRequest) error
}

// interface
type authUsecase struct {
	customerDao driven.CustomerPortDriven
}

func NewCustomerAuthUsecase(customerDao driven.CustomerPortDriven) AuthUsecase {
	return &authUsecase{
		customerDao: customerDao,
	}
}

func (a *authUsecase) SignIn(ctx context.Context, customer dto.SignInRequest) (*dto.AuthResponse, error) {

	//TODO: valiadtion

	// validationErr := validation.SignInValidation{
	// 	Email:    customer.Email,
	// 	Password: customer.Password,
	// }.SignInValidation()

	// if validationErr != nil {
	// 	return dto.AuthResponse{}, validationErr
	// }

	customerInfo, err := a.customerDao.GetByEmail(ctx, customer.Email)
	//TODO: change error message
	if err != nil {
		return nil, err
	}

	isPasswordValid := encryption.ComparePasswords(customerInfo.Password, customer.Password)

	//TODO: change error message
	if !isPasswordValid {
		return nil, errors.New("error")
	}

	token, err := authenticator.ExtractClaims(customerInfo.ID)

	if err != nil {
		return nil, err
	}

	//TODO: Generate Token
	return &dto.AuthResponse{
		AccessToken: token,
	}, nil

}

func (a *authUsecase) SignUp(ctx context.Context, customer dto.SignUpRequest) (*dto.AuthResponse, error) {

	isExist, err := a.customerDao.GetByEmail(ctx, customer.Email)
	if err != nil {
		return nil, err
	}

	if isExist != nil {
		return nil, errors.New(message.UserExist)
	}
	id := uuid.New()

	//TODO: validation

	// validationErrors := validation.SignUpValidation{
	// 	ID:        id,
	// 	FirstName: customer.FirstName,
	// 	LastName:  customer.LastName,
	// 	Password:  customer.Password,
	// 	Email:     customer.Email,
	// 	Phone:     customer.Phone,
	// }.SignUpValidation()

	// if validationErrors != nil {
	// 	fmt.Println(errors.Unwrap(validationErrors))
	// 	return dto.AuthResponse{}, validationErrors
	// }

	hashPassword, err := encryption.HashPassword(customer.Password)
	if err != nil {
		fmt.Print(hashPassword)
	}

	newCustomerData := domain.Customer{
		ID:        id,
		FirstName: customer.FirstName,
		LastName:  customer.LastName,
		Password:  hashPassword,
		Email:     customer.Email,
		Phone:     customer.Phone,
	}

	// err = domain.NewCustomer(newCustomerData)

	// if err != nil {
	// 	return dto.AuthResponse{}, err
	// }

	newCustomerData.CreatedAt = time.Now()
	newCustomer, err := a.customerDao.Create(ctx, newCustomerData)
	if err != nil {
		return nil, err
	}

	token, err := authenticator.ExtractClaims(newCustomer.ID)

	if err != nil {
		return nil, err
	}

	return &dto.AuthResponse{
		AccessToken: token,
	}, nil

}

func (a *authUsecase) ForgetPassword(ctx context.Context, user dto.ForgetPasswordRequest) (dto.Customer, error) {

	//TODO:

	return dto.Customer{}, nil
}

func (c *authUsecase) UpdatePassword(ctx context.Context, data dto.UpdatePasswordRequest) error {

	//TODO: check password
	tokenData, err := authenticator.DecodeToken(data.AccessToken)
	if err != nil {
		return err
	}

	// validationErr := validation.UpdatePasswordValidation()

	customer, err := c.customerDao.GetByID(ctx, tokenData.Id)
	if err != nil {
		return err
	}

	fmt.Println(customer.ID)

	isPasswordValid := encryption.ComparePasswords(customer.Password, data.OldPassword)

	if !isPasswordValid {
		return errors.New("old password is wrong")
	}

	password, err := encryption.HashPassword(data.NewPassword)

	if err != nil {
		return err
	}

	err = c.customerDao.UpdatePassword(ctx, tokenData.Id, password)

	if err != nil {
		return err
	}

	return nil

}
