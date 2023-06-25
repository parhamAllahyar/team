package grpc

import (
	"fmt"
	"log"
	"net"
	pb "workspace/api/proto/grpcpkg"
	"workspace/configs"
	"workspace/internal/core/usecase"

	"google.golang.org/grpc"
)

type grpcHandler struct {
	pb.UnimplementedWorkspaceServer
	WorkspaceUsecase usecase.WorkspaceUsecase
}

func NewGrpcServices(WorkspaceUsecase usecase.WorkspaceUsecase) *grpcHandler {
	return &grpcHandler{
		WorkspaceUsecase: WorkspaceUsecase,
	}
}

func InitGrpcServer(config configs.GrpcServer, handlers *grpcHandler) {

	// address := config.Host + string(config.Port)

	port := ":" + config.Port

	fmt.Println("PORT IS ", port)

	listener, err := net.Listen("tcp", ":9009")

	//TODO: change error
	if err != nil {
		log.Fatalln(err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterWorkspaceServer(grpcServer, handlers)
	err = grpcServer.Serve(listener)

	//TODO: change error
	if err != nil {
		log.Fatalln(err)
	}

}
