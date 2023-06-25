package main

import (
	"board/configs"
	"board/internal/adapter/driven/postgres"
	grpcHandler "board/internal/adapter/driver/grpc"
	Repo "board/internal/core/port/driven"
	"board/internal/core/usecase"
	"context"
	"database/sql"
)

func main() {

	//Call Configs
	configs := configs.LoadConfig()

	//run postgres
	db, err := runPostgreSql(configs.Postgre)

	if err != nil {
		panic("database was fail")
	}

	migration(context.Background(), db)

	bardPortDriven, sectionPortDriven, taskPortDriven, subTaskPortDriven, tagPortDriven := initRepo(db)

	boardUsecase, sectionUsecase, taskUsecase, subTaskUsecase, tagUsecase := initUsecase(bardPortDriven, sectionPortDriven, taskPortDriven, subTaskPortDriven, tagPortDriven)

	//run grpc server
	runGrpcServer(configs.GrpcServer, boardUsecase, sectionUsecase, taskUsecase, subTaskUsecase, tagUsecase)

}

func runPostgreSql(config configs.PostgresDBConfig) (*sql.DB, error) {
	pg, err := postgres.NewPostgre(config)
	if err != nil {
		return nil, err
	}
	return pg, nil
}

func migration(ctx context.Context, db *sql.DB) {
	err := postgres.Migrate(context.Background(), db)
	if err != nil {
		panic(err)
	}
}

func initRepo(db *sql.DB) (Repo.BoardPortDriven, Repo.SectionPortDriven, Repo.TaskPortDriven, Repo.SubTaskPortDriven, Repo.TagPortDriven) {
	boardRepo := postgres.NewBoardRepo(db)
	sectionRepo := postgres.NewSectionRepo(db)
	taskRepo := postgres.NewTaskRepo(db)
	subTaskRepo := postgres.NewSubTaskdRepo(db)
	tagRepo := postgres.NewTagRepo(db)
	return boardRepo, sectionRepo, taskRepo, subTaskRepo, tagRepo
}

func initUsecase(boardDao Repo.BoardPortDriven, sectionDao Repo.SectionPortDriven, taskDao Repo.TaskPortDriven, subTaskDao Repo.SubTaskPortDriven, tagDao Repo.TagPortDriven) (usecase.BoardUsecase, usecase.SectionUsecase, usecase.TaskUsecase, usecase.SubTaskUsecase, usecase.TagUsecase) {
	boardUsecase := usecase.NewBoardUsecase(boardDao)
	sectionUsecase := usecase.NewSectionUsecase(sectionDao)
	taskUsecase := usecase.NewTaskUsecase(taskDao)
	subTaskUsecase := usecase.NewSubTaskUsecase(subTaskDao)
	tagUsecase := usecase.NewTagUsecase(tagDao)
	return boardUsecase, sectionUsecase, taskUsecase, subTaskUsecase, tagUsecase
}

func runGrpcServer(config configs.GrpcServer, boardUsecase usecase.BoardUsecase, sectionUsecase usecase.SectionUsecase, taskUsecase usecase.TaskUsecase, subTaskUsecase usecase.SubTaskUsecase, tagUsecase usecase.TagUsecase) {

	handler := grpcHandler.NewGrpcServices(boardUsecase, sectionUsecase, taskUsecase, subTaskUsecase, tagUsecase)

	grpcHandler.InitGrpcServer(config, handler)

}
