package main

import (
	"context"
	"fmt"
	"workspace/configs"
	"workspace/internal/adapter/driven/neo4j"
	Neo4jRepo "workspace/internal/adapter/driven/neo4j"
	grpcHandler "workspace/internal/adapter/driver/grpc"
	Repo "workspace/internal/core/port/driven"
	"workspace/internal/core/usecase"

	neo4jPkg "github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func main() {

	//config app
	config := configs.LoadConfig()

	fmt.Println(config)

	//neo4j
	neo4j := runNeo4jworkspaceDB(config.NeoConfig)

	fmt.Println("xcxccxcxcx", neo4j)

	//call repos
	workspace := neo4jRepo(neo4j, context.Background())

	//init usecase
	workspaceUsecase := initUsecase(workspace)

	//Start grpc server
	runGrpcServer(config.GrpcServer, workspaceUsecase)

}

func runNeo4jworkspaceDB(config configs.NEOConfig) neo4jPkg.DriverWithContext {
	neo4j, err := neo4j.Connection(config)
	if err != nil {
		panic(err)
	}
	return neo4j
}

func neo4jRepo(neo4j neo4jPkg.DriverWithContext, ctx context.Context) Repo.WorkspacePortDriven {
	return Neo4jRepo.NewWorkspaceRepo(neo4j, ctx)
}

func initUsecase(workspaceRepo Repo.WorkspacePortDriven) usecase.WorkspaceUsecase {
	workspaceUsecase := usecase.NewWorkspaceUsecase(workspaceRepo)
	return workspaceUsecase
}

func runGrpcServer(config configs.GrpcServer, workspace usecase.WorkspaceUsecase) {

	fmt.Println("run")
	handlers := grpcHandler.NewGrpcServices(workspace)
	grpcHandler.InitGrpcServer(config, handlers)

	fmt.Println("run")

}
