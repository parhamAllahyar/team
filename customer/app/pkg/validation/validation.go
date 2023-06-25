package validation

import (
	"errors"

	"github.com/google/uuid"
	"gopkg.in/go-playground/validator.v9"

)

type SignUpValidation struct {
	ID        uuid.UUID `validate:"required"`
	FirstName string    `validate:"required"`
	LastName  string    `validate:"required"`
	Password  string    `validate:"required"`
	Email     string    `validate:"required,email"`
	Phone     string    `validate:"required"`
}

//TODO: change message

func (s SignUpValidation) SignUpValidation() error {
	var validationERR error
	v := validator.New()
	err := v.Struct(s)
	for _, e := range err.(validator.ValidationErrors) {
		validationERR = errors.Join(validationERR, errors.New(e.Field()+" IS "+e.Tag()))
	}
	
	if validationERR != nil {
		return validationERR
	}
	return nil
}

type SignInValidation struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}

func (s SignInValidation) SignInValidation() error {
	var validationERR error
	v := validator.New()
	err := v.Struct(s)
	for _, e := range err.(validator.ValidationErrors) {
		validationERR = errors.Join(validationERR, errors.New(e.Field()+" IS "+e.Tag()))
	}
	return validationERR
}

type UpdatePasswordValidation struct {
	OldPassword string `validate:"required,email"`
	NewPassword string `validate:"required"`
}

func (s SignInValidation) UpdatePasswordValidation() error {
	var validationERR error
	v := validator.New()
	err := v.Struct(s)
	for _, e := range err.(validator.ValidationErrors) {
		validationERR = errors.Join(validationERR, errors.New(e.Field()+" IS "+e.Tag()))
	}
	return validationERR
}
