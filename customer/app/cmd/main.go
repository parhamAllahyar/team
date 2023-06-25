package main

import (
	//"context"
	"customer/configs"
	"customer/internal/adapter/driven/postgres"
	grpcHandler "customer/internal/adapter/driver/grpc"
	Repo "customer/internal/core/port/driven"
	"customer/internal/core/usecase"
	"database/sql"
	"fmt"
)

func main() {

	configs := configs.LoadConfig()

	db, err := runPostgresDB(configs.Postgre)

	if err != nil {
		panic(err)
	}

	// err = migratePostgres(db)

	// if err != nil {
	// 	panic(err)
	// }

	customerDb := postgresCustomerRepo(db)
	customerUsecase, authUsecase := initUsecases(customerDb)

	fmt.Println("use case ok")

	runGrpcServer(configs.GrpcServer, customerUsecase, authUsecase)

	fmt.Println("service ok")

}

func runPostgresDB(config configs.PostgresDBConfig) (*sql.DB, error) {
	pg, err := postgres.NewPostgres(config)
	if err != nil {
		return nil, err
	}
	return pg, nil
}

func postgresCustomerRepo(db *sql.DB) Repo.CustomerPortDriven {
	return postgres.NewCustomerRepo(db)
}

func initUsecases(customerRepo Repo.CustomerPortDriven) (usecase.CustomerUsecase, usecase.AuthUsecase) {
	customerUsecase := usecase.NewCustomerUsecase(customerRepo)
	authUsecase := usecase.NewCustomerAuthUsecase(customerRepo)
	return customerUsecase, authUsecase

}

func runGrpcServer(config configs.GrpcServer, customerUsecase usecase.CustomerUsecase, authUsecase usecase.AuthUsecase) {
	handlers := grpcHandler.NewGrpcServices(customerUsecase, authUsecase)
	grpcHandler.InitGrpcServer(config, handlers)
}
