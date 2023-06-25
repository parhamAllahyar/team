package usecase

import (
	"admin/internal/core/port/driven"
)

type CustomerUsecase interface{}

type customerUsecase struct {
	customDao driven.CustomerPortDriven
}

func NewCustomerUsecase(customDao driven.CustomerPortDriven) CustomerUsecase {
	return &customerUsecase{
		customDao: customDao,
	}
}
