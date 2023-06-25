package postgres

import (
	"context"
	"customer/internal/core/domain"
	"customer/internal/core/port/driven"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

)

type customerGORM struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey;"`
	FirstName string
	LastName  string
	Password  string
	Email     string
	Phone     string
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
	// DeletedAt time.Time
}

type customerRepo struct {
	customerPort driven.CustomerPortDriven
	db           *sql.DB
}

func NewCustomerRepo(db *sql.DB) customerRepo {
	return customerRepo{
		db: db,
	}
}

func (c customerRepo) callGorm() *gorm.DB {
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: c.db,
	}), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	return gormDB
}

func (c customerRepo) Create(ctx context.Context, customer domain.Customer) (*domain.Customer, error) {

	//TODO: return value

	newcustomer := customerGORM{
		ID:        customer.ID,
		FirstName: customer.FirstName,
		LastName:  customer.LastName,
		Password:  customer.Password,
		Email:     customer.Email,
		Phone:     customer.Phone,
		Status:    "active",
		CreatedAt: customer.CreatedAt,
	}

	query := c.callGorm().Table("customers").Create(&newcustomer)

	if query.Error != nil {
		return nil, query.Error
	}

	return &customer, nil

}
func (c customerRepo) Update(ctx context.Context, customer domain.Customer) (*domain.Customer, error) {
	coustomerData := customerGORM{
		FirstName: customer.FirstName,
		LastName:  customer.LastName,
	}
	query := c.callGorm().Table("customers").Where("id = ?", customer.ID).Updates(coustomerData).Find(&coustomerData)
	if query.Error != nil {
		return nil, query.Error
	}

	return &domain.Customer{
		ID:        coustomerData.ID,
		FirstName: coustomerData.FirstName,
		LastName:  coustomerData.LastName,
		Email:     coustomerData.Email,
		Phone:     coustomerData.Phone,
	}, nil

}
func (c customerRepo) Delete(ctx context.Context, id uuid.UUID) error {
	query := c.callGorm().Table("customers").Where("id = ?", id).Delete(&customerGORM{})

	return query.Error

}
func (c customerRepo) GetByID(ctx context.Context, id uuid.UUID) (*domain.Customer, error) {
	var customer domain.Customer
	query := c.callGorm().Table("customers").Where("id = ?", id).Find(&customer)
	if query.Error != nil {
		return nil, query.Error
	}
	fmt.Println(customer)
	return &customer, nil

}

func (c customerRepo) GetByEmail(ctx context.Context, email string) (*domain.Customer, error) {
	var customer domain.Customer
	query := c.callGorm().Table("customers").Where("email = ?", email).Find(&customer)
	if query.Error != nil {
	return	nil, query.Error
	}
	fmt.Println(customer)
	return &customer, nil
}

func (c customerRepo) Index(ctx context.Context) ([]domain.Customer, error) {

	//TODO: refactor

	var customers []domain.Customer

	query := c.callGorm().Table("customers").Find(&customers)

	if query.Error != nil {
		return nil, query.Error
	}

	return customers, nil

}
func (c customerRepo) UpdatePhone(ctx context.Context, phone string) error {

	query := c.callGorm().Table("customers").Where("phone = ?", phone).Update("phone", phone)

	if query.Error != nil {
		return query.Error
	}

	return nil

}
func (c customerRepo) UpdateEmail(ctx context.Context, email string) error {

	query := c.callGorm().Table("customers").Where("email = ?", email).Updates(&customerGORM{Email: email})

	if query.Error != nil {
		return query.Error
	}

	return nil
}
func (c customerRepo) UpdatePassword(ctx context.Context, id uuid.UUID, password string) error {

	query := c.callGorm().Table("customers").Where("id = ?", id.String()).Updates(&customerGORM{Password: password})

	if query.Error != nil {
		return query.Error
	}

	return nil

}
