package grpc

import (
	//"admin/internal/core/service"
	"context"
	pb "customer/api/proto/grpcpkg"
	"customer/internal/core/dto"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"

)



func (g *grpcHandler) Update(ctx context.Context, req *pb.UpdateRequest) (*pb.UpdateResponse, error) {

	customer, err := g.customerUsecase.Update(ctx, dto.UpdateCustomerRequest{
		FirstName:   req.GetFirstName(),
		LastName:    req.GetLastName(),
		AccessToken: req.GetAccessToken(),
	})

	if err != nil {
		return nil, err
	}

	return &pb.UpdateResponse{
		CustomerItem: &pb.CustomerItem{
			Id: customer.ID.String(),
			CustomerData: &pb.CustomerData{
				FirstName: customer.FirstName,
				LastName:  customer.LastName,
				Email:     customer.Email,
				Phone:     customer.Phone,
			},
		},
	}, nil

}

func (g *grpcHandler) Delete(ctx context.Context, req *pb.DeleteRequest) (*emptypb.Empty, error) {

	err := g.customerUsecase.Delete(ctx, req.GetAccessToken())
	if err != nil {
		return &emptypb.Empty{}, err
	}
	return &emptypb.Empty{}, nil
}

func (g *grpcHandler) GetByID(ctx context.Context, req *pb.GetByIDRequest) (*pb.GetByIDResponse, error) {

	id, _ := uuid.Parse(req.GetId())

	customer := g.customerUsecase.GetByID(ctx, req.GetAccessToken(), id)
	return &pb.GetByIDResponse{
		CustomerItem: &pb.CustomerItem{
			Id: customer.ID.String(),
			CustomerData: &pb.CustomerData{
				FirstName: customer.FirstName,
				LastName:  customer.LastName,
				Email:     customer.Email,
				Phone:     customer.Phone,
			},
		},
	}, nil
}

func (g *grpcHandler) Index(ctx context.Context, req *pb.IndexRequest) (*pb.IndexResponse, error) {

	//TODO:

	//g.customerUsecase.Index()

	return &pb.IndexResponse{}, nil
}

func (g *grpcHandler) UpdatePhone(ctx context.Context, req *pb.UpdatePhoneRequest) (*emptypb.Empty, error) {

	err := g.customerUsecase.UpdatePhone(ctx, req.GetAccessToken(), req.GetPhone())
	if err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}

func (g *grpcHandler) UpdateEmail(ctx context.Context, req *pb.UpdateEmailRequest) (*emptypb.Empty, error) {

	err := g.customerUsecase.UpdateEmail(ctx, req.GetAccessToken(), req.GetEmail())
	if err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}
