package grpc

import (
	pb "board/api/proto/grpcpkg"
	"board/configs"
	"board/internal/core/usecase"

	"log"
	"net"

	"google.golang.org/grpc"
)

type grpcHandler struct {
	pb.UnimplementedBoardServer
	boardUsecase   usecase.BoardUsecase
	sectionUsecase usecase.SectionUsecase
	taskUsecase    usecase.TaskUsecase
	subTaskUsecase usecase.SubTaskUsecase
	tagUsecase     usecase.TagUsecase
}

func NewGrpcServices(boardUsecase usecase.BoardUsecase, sectionUsecase usecase.SectionUsecase, taskUsecase usecase.TaskUsecase, subTaskUsecase usecase.SubTaskUsecase, tagUsecase usecase.TagUsecase) *grpcHandler {
	return &grpcHandler{
		boardUsecase:   boardUsecase,
		sectionUsecase: sectionUsecase,
		taskUsecase:    taskUsecase,
		subTaskUsecase: subTaskUsecase,
		tagUsecase:     tagUsecase,
	}
}

func InitGrpcServer(cnfig configs.GrpcServer, handlers *grpcHandler) {

	url := ":6060"

	listener, err := net.Listen("tcp", url)

	//TODO: change error
	if err != nil {
		log.Fatalln(err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterBoardServer(grpcServer, handlers)

	err = grpcServer.Serve(listener)

	//TODO: change error
	if err != nil {
		log.Fatalln(err)
	}

}
