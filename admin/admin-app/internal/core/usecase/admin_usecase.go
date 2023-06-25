package usecase

//for doing every thing get or use with type of port

import (
	"admin/internal/core/domain"
	"admin/internal/core/dto"
	"admin/internal/core/port/driven"
	"admin/pkg/auth/authenticator"
	"admin/pkg/auth/encryption"
	errormessage "admin/pkg/error-message"
	"admin/pkg/utils"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"

)

type AdminUsecase interface {
	Get(domain.Admin) error
	Update(dto.UpdateRequest) error
	UpdatePhone(dto.UpdatePoneRequest) error
	UpdatePhoneRequest(dto.UpdatePoneRequest) error
	GenerateInviteLink(dto.InviteLinkRequest) error
	UpdatePassword(dto.UpdatePasswordRequest) error
}

//interface

type adminUsecase struct {
	adminDao driven.AdminPortDriven
	linkDao  driven.LinkPortDriven
	otpDao   driven.OtpPortDriven
}

func NewAdminUsecase(adminDao driven.AdminPortDriven, linkDao driven.LinkPortDriven, otpDao driven.OtpPortDriven) AdminUsecase {
	return &adminUsecase{
		adminDao: adminDao,
		linkDao:  linkDao,
		otpDao:   otpDao,
	}
}

func (a *adminUsecase) Get(domain.Admin) error {
	//Some operation
	return nil
}

func (a *adminUsecase) Update(dto.UpdateRequest) error {
	//Some operation
	return nil
}

func (a *adminUsecase) UpdatePassword(request dto.UpdatePasswordRequest) error {

	//TODO: validation

	admin, err := authenticator.DecodeToken(request.AccessToken)

	if err != nil {
		return errors.New("")
	}

	fmt.Println("id is:", admin.Id)

	query, err := a.adminDao.GetById(admin.Id)

	fmt.Println("last name", query.LastName)

	if err != nil {
		return errors.New(errormessage.DatabaseError)
	}

	if query == nil && query.Status == "active" {
		return errors.New(errormessage.EmailOrPasswordIsWrong)
	}

	//compaire old password and hash vlaue
	isOldPasswordTrue := encryption.ComparePasswords(query.Password, request.OldPassword)

	if isOldPasswordTrue == false {
		return errors.New(errormessage.OldPasswordIsWrong)
	}

	hashedValue, err := encryption.HashPassword(request.NewPassword)

	if err != nil {
		return errors.New("errrrrrr")
	}

	_, err = a.adminDao.Update(domain.Admin{
		ID:       admin.Id,
		Password: hashedValue,
	})

	if err != nil {
		return errors.New("errrrrrr")
	}

	//Some operation
	return nil
}

func (a *adminUsecase) UpdatePhoneRequest(request dto.UpdatePoneRequest) error {

	//TODO: validate phone

	admin, err := authenticator.DecodeToken(request.AccessToken)

	if err != nil {
		return err
	}

	query, err := a.adminDao.GetById(admin.Id)

	otp, err := a.otpDao.Create(domain.OTP{
		ID:       uuid.New(),
		Type:     "update Phone",
		NewPhone: request.Phone,
		OTP:      rand.Intn(99999 - 10000),
		Status:   "Active",
		ExpireAt: time.Now().Add(1),
	})

	//send sms
	fmt.Println(query.Phone)
	fmt.Println(otp)

	return nil

}

func (a *adminUsecase) GenerateInviteLink(request dto.InviteLinkRequest) error {

	//TODO: validation email

	//check token
	admin, err := authenticator.DecodeToken(request.AccessToken)

	if err != nil {
		return err
	}

	//TODO: check permission
	fmt.Println(admin.Id)

	//generate link
	link, err := a.linkDao.Create(domain.Link{
		ID:        uuid.New(),
		Type:      "invite_link",
		Email:     request.Email,
		Link:      utils.RandStr(30),
		Status:    "active",
		ExpiredAt: time.Now().AddDate(0, 0, 1),
	})

	//TODO: send email
	fmt.Println(link.Link)

	return nil
}

func (a *adminUsecase) UpdatePhone(dto.UpdatePoneRequest) error {

	return nil
}
