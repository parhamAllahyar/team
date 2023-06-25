package usecase

import (
	"customer/pkg/auth/authenticator"
	"fmt"

)

type VerificationUsecase interface {
	VerifyPhoneRequest(token string) error
	VerifyEmailRequest(token string) error
	VerifyPhone(token string) error
	VerifyEmail(token string) error
}

func (c *customerUsecase) VerifyPhoneRequest(token string) error {

	//check 2 min or not send

	id, err := authenticator.IsCustomer()

	//TODO: change error message
	if err != nil {
		return err
	}

	//code := rand.Intn(100000)

	//TODO: store code

	//TODO: send message to notification service

	fmt.Println(id, err)
	return nil
}
func (c *customerUsecase) VerifyEmailRequest(token string) error {

	id, err := authenticator.IsCustomer()
	fmt.Println(id, err)
	return nil
}

func (c *customerUsecase) VerifyPhone(token string) error {

	id, err := authenticator.IsCustomer()
	fmt.Println(id, err)
	return nil
}
func (c *customerUsecase) VerifyEmail(token string) error {

	id, err := authenticator.IsCustomer()
	fmt.Println(id, err)
	return nil
}
