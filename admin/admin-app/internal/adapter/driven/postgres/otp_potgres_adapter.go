package postgres

import (
	"admin/internal/core/domain"
	"admin/internal/core/port/driven"
	"database/sql"

)

type otpRepo struct {
	authRepo driven.OtpPortDriven
	db       *sql.DB
}

func NewOTPRepo(db *sql.DB) *otpRepo {
	return &otpRepo{
		db: db,
	}
}

func (o *otpRepo) GetById(string) (*domain.OTP, error) {
	return nil, nil
}
func (o *otpRepo) Update(otp domain.OTP) (*domain.OTP, error) {
	query := callGorm(o.db).Table("otps").Where("id = ?", otp.ID).Updates(otp).Find(&otp)
	if query.Error != nil {
		return nil, query.Error
	}
	return &otp, nil
}
func (o *otpRepo) Create(otp domain.OTP) (*domain.OTP, error) {
	query := callGorm(o.db).Table("otps").Create(&otp).Find(&otp)
	if query.Error != nil {
		return nil, query.Error
	}
	return &otp, nil
}
