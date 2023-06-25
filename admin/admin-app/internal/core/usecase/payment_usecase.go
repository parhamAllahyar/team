package usecase

import "admin/internal/core/port/driven"

type PaymentUsecase interface{}

type paymentUsecase struct {
	paymentDao driven.PaymentPortDriven
}

func NewPaymentUsecase(paymentDao driven.PaymentPortDriven) PaymentUsecase {
	return &paymentUsecase{
		paymentDao: paymentDao,
	}
}
