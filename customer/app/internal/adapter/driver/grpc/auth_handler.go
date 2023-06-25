package grpc

import (
	// //"admin/internal/core/service"
	"context"
	pb "customer/api/proto/grpcpkg"
	"customer/internal/core/dto"

	"google.golang.org/protobuf/types/known/emptypb"

)

func (g *grpcHandler) SignIn(ctx context.Context, req *pb.SignInRequest) (*pb.SignInResponse, error) {

	response, err := g.authUsecase.SignIn(ctx, dto.SignInRequest{
		Password: req.GetPassword(),
		Email:    req.GetEmail(),
	})

	if err != nil {
		return nil, err
	}

	return &pb.SignInResponse{
		AccessToken: response.AccessToken,
	}, nil
}

func (g *grpcHandler) SignUp(ctx context.Context, req *pb.SignUpRequest) (*pb.SignUpResponse, error) {

	response, err := g.authUsecase.SignUp(ctx, dto.SignUpRequest{
		FirstName: req.GetCustomerData().FirstName,
		LastName:  req.GetCustomerData().LastName,
		Password:  req.GetPassword(),
		Email:     req.GetCustomerData().Email,
		Phone:     req.GetCustomerData().Phone,
	})

	if err != nil {
		return nil, err
	}

	return &pb.SignUpResponse{
		AccessToken: response.AccessToken,
	}, nil
}

func (g *grpcHandler) UpdatePassword(ctx context.Context, req *pb.UpdatePasswordRequest) (*emptypb.Empty, error) {

	err := g.authUsecase.UpdatePassword(ctx, dto.UpdatePasswordRequest{
		AccessToken: req.GetAccessToken(),
		OldPassword: req.GetOldPassword(),
		NewPassword: req.GetNewPassword(),
	})

	if err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil

}
