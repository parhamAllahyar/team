package grpc

import (
	pb "customer/api/proto/grpcpkg"
	"customer/configs"
	"customer/internal/core/usecase"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

)

type grpcHandler struct {
	pb.UnimplementedCustomerServer
	customerUsecase usecase.CustomerUsecase
	authUsecase     usecase.AuthUsecase
	//userService service.User
}

func NewGrpcServices(customerUsecase usecase.CustomerUsecase, authUsecase usecase.AuthUsecase) *grpcHandler {
	return &grpcHandler{
		customerUsecase: customerUsecase,
		authUsecase:     authUsecase,
	}
}

func InitGrpcServer(config configs.GrpcServer , handlers *grpcHandler) {

	url := ":" + config.Port
	fmt.Println(url)
	listener, err := net.Listen("tcp", url)

	//TODO: change error
	if err != nil {
		log.Fatalln(err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterCustomerServer(grpcServer, handlers)

	err = grpcServer.Serve(listener)

	//TODO: change error
	if err != nil {
		log.Fatalln(err)
	}

}
