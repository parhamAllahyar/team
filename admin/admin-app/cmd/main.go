package main

import (
	"admin/configs"
	"admin/internal/adapter/driven/postgres"
	httpServer "admin/internal/adapter/driver/http"
	Repo "admin/internal/core/port/driven"
	"admin/internal/core/usecase"
	"context"
	"database/sql"
	"fmt"
)

func main() {

	fmt.Println("mnnnnnnnnnnnnnnnnnn")

	//get configs
	config := configs.LoadConfig()

	//db
	db := runPostgres(config.Postgre)

	//migration
	migration(context.Background(), db)

	//repo
	//db
	adminRepo, linkRepo, otpRepo := initPostgresRepo(db)

	//grpc
	customerRepo, paymentRepo := initGRPCClient()

	//usecase
	customerUsecase, authUsecase, adminUsecase, paymentUsecase := initUsecase(customerRepo, adminRepo, paymentRepo, linkRepo, otpRepo)

	//run http server
	runHttpServer(config.HttpServer, customerUsecase, authUsecase, adminUsecase, paymentUsecase)

}

// run db
func runPostgres(config configs.PostgresDBConfig) *sql.DB {
	pg, err := postgres.NewPostgres(config)
	if err != nil {
		panic(err)
	}
	return pg
}

// migiration
func migration(ctx context.Context, db *sql.DB) {
	err := postgres.Migrate(context.Background(), db)
	if err != nil {
		panic(err)
	}
}

// //repo
// db
func initPostgresRepo(db *sql.DB) (Repo.AdminPortDriven, Repo.LinkPortDriven, Repo.OtpPortDriven) {
	adminRepo := postgres.NewAdminRepo(db)
	linkRepo := postgres.NewLinkRepo(db)
	otpRepo := postgres.NewOTPRepo(db)
	return adminRepo, linkRepo, otpRepo

}

// grpc
func initGRPCClient() (Repo.CustomerPortDriven, Repo.PaymentPortDriven) {

	return nil, nil
}

// usecase
func initUsecase(customerRepo Repo.CustomerPortDriven, adminRepo Repo.AdminPortDriven, paymentRepo Repo.PaymentPortDriven, linkRepo Repo.LinkPortDriven, otpRepo Repo.OtpPortDriven) (usecase.CustomerUsecase, usecase.AuthUsecase, usecase.AdminUsecase, usecase.PaymentUsecase) {
	customerUsecase := usecase.NewCustomerUsecase(customerRepo)
	authUsecase := usecase.NewAuthUsecase(adminRepo, linkRepo)
	admiUsecase := usecase.NewAdminUsecase(adminRepo, linkRepo, otpRepo)
	paymentUsecase := usecase.NewPaymentUsecase(paymentRepo)
	return customerUsecase, authUsecase, admiUsecase, paymentUsecase

}

// run http server
func runHttpServer(config configs.HTTPServer, customerUsecase usecase.CustomerUsecase, authUsecase usecase.AuthUsecase, admiUsecase usecase.AdminUsecase, paymentUsecase usecase.PaymentUsecase) {
	usecases := httpServer.HttpHandlers{
		AuthUsecase:     authUsecase,
		AdminUsecase:    admiUsecase,
		PaymentUsecase:  paymentUsecase,
		CustomerUsecase: customerUsecase,
	}

	fmt.Println("http server")

	err := httpServer.NewHttpServer(usecases, config)
	if err != nil {
		panic(err)
	}
}
