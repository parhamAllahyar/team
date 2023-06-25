package usecase

//for doing every thing get or use with type of port

import (
	//"customer/internal/core/domain"
	"context"
	"customer/internal/core/domain"
	"customer/internal/core/dto"
	"customer/internal/core/port/driven"
	"customer/pkg/auth/authenticator"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type CustomerUsecase interface {
	Delete(context.Context, string) error
	GetByID(context.Context, string, uuid.UUID) dto.Customer
	Update(context.Context, dto.UpdateCustomerRequest) (dto.Customer, error)
	Search(context.Context, string, string) []dto.Customer
	Index(context.Context, string) ([]dto.Customer, error)
	UpdatePhone(context.Context, string, string) error
	UpdateEmail(context.Context, string, string) error
}

// interface
type customerUsecase struct {
	customerDao driven.CustomerPortDriven
}

func NewCustomerUsecase(customerDao driven.CustomerPortDriven) CustomerUsecase {
	return &customerUsecase{
		customerDao: customerDao,
	}
}

func (c *customerUsecase) Index(ctx context.Context, token string) ([]dto.Customer, error) {

	isAdmin := authenticator.IsAdminToken(token)

	//TODO: change error message
	if !isAdmin {
		return nil, errors.New("you are not admin")
	}

	customers, err := c.customerDao.Index(ctx)

	if err != nil {
		return nil, err
	}

	fmt.Println(customers)

	//TODO:  Assign customer domain  tocustomer  dto

	//TODO: Rule
	return nil, nil

}

func (c *customerUsecase) Update(ctx context.Context, customer dto.UpdateCustomerRequest) (dto.Customer, error) {

	customerData, err := authenticator.DecodeToken(customer.AccessToken)

	//TODO: change error message
	if err != nil {
		return dto.Customer{}, err
	}

	//TODO: change id
	customerInfo, err := c.customerDao.Update(ctx, domain.Customer{
		ID:        customerData.Id,
		FirstName: customer.FirstName,
		LastName:  customer.LastName,
	})

	if err != nil {
		return dto.Customer{}, err
	}

	return dto.Customer{
		ID:        customerInfo.ID,
		FirstName: customerInfo.FirstName,
		LastName:  customerInfo.LastName,
		Email:     customerInfo.Email,
		Phone:     customerInfo.Phone,
		CreatedAt: customerInfo.CreatedAt,
		UpdatedAt: customerInfo.UpdatedAt,
	}, nil

}

func (c *customerUsecase) Delete(ctx context.Context, id string) error {
	return nil
}

func (c *customerUsecase) GetByID(ctx context.Context, token string, id uuid.UUID) dto.Customer {

	//TODO: Rule
	customer, err := c.customerDao.GetByID(ctx, id)

	if err != nil {

	}

	//Some operation
	return dto.Customer{
		//ID:        customer.ID,
		FirstName: customer.FirstName,
		LastName:  customer.LastName,
		Email:     customer.Email,
		Phone:     customer.Phone,
		Status:    customer.Status,
		CreatedAt: customer.CreatedAt,
		UpdatedAt: customer.UpdatedAt,
		DeletedAt: customer.DeletedAt,
	}
}

func (c *customerUsecase) Search(ctx context.Context, token string, data string) []dto.Customer {
	//TODO: Rule
	return nil
}

func (c *customerUsecase) UpdatePhone(ctx context.Context, token string, phone string) error {

	//TODO: Rule
	tokenData, err := authenticator.DecodeToken(token)

	fmt.Print(tokenData)

	//TODO: Change error message
	if err != nil {
		return err
	}

	//TODO: cache new email by redis (new phone and customer id)

	//TODO: call notification microservice

	return nil

}

func (c *customerUsecase) UpdateEmail(ctx context.Context, token string, email string) error {
	//TODO: Rule
	tokenData, err := authenticator.DecodeToken(token)

	fmt.Print(tokenData)

	//TODO: Change error message
	if err != nil {
		return err
	}

	//TODO: cache new email by redis (new email and user id)

	//TODO: call notification microservice

	return nil
}
