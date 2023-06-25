package grpc

import (
	"context"
	"fmt"
	pb "workspace/api/proto/grpcpkg"
	"workspace/internal/core/dto"

	"github.com/google/uuid"
)

func (g *grpcHandler) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {

	fmt.Println("brfhbhrfbhrhbfhbrfhbrfhbrfhbfrhbrfhbrfhbrhbfhbrfbhrhbrfhbfrhbrbfhbrhf")

	fmt.Println("title", req.GetTitle())

	fmt.Println("token", req.GetAccessToken())

	workspace, err := g.WorkspaceUsecase.Create(dto.CreateWorkspaceRequest{
		Title:       req.GetTitle(),
		AccessToken: req.GetAccessToken(),
	})

	if err != nil {
		return &pb.CreateResponse{}, err
	}

	return &pb.CreateResponse{
		Id:    workspace.ID.String(),
		Title: workspace.Title,
	}, nil
}
func (g *grpcHandler) Update(ctx context.Context, req *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	workspace, err := g.WorkspaceUsecase.Update(dto.UpdateWorkspaceRequest{
		Title:       req.GetTitle(),
		AccessToken: req.GetAccessToken(),
		ID:          uuid.MustParse(req.GetId()),
	})

	if err != nil {
		return &pb.UpdateResponse{}, err
	}

	return &pb.UpdateResponse{
		Id:    workspace.ID.String(),
		Title: workspace.Title,
	}, nil
}

func (g *grpcHandler) CustomerWorkspaces(ctx context.Context, req *pb.CustomerWorkspacesRequest) (*pb.CustomerWorkspacesResponse, error) {

	// workspaces, err := g.WorkspaceUsecase.CustomerWorkspaces(req.G)

	return &pb.CustomerWorkspacesResponse{}, nil
}

func (g *grpcHandler) Index(ctx context.Context, req *pb.IndexRequest) (*pb.IndexResponse, error) {
	return &pb.IndexResponse{}, nil
}

func (g *grpcHandler) WorkspaceMembers(ctx context.Context, req *pb.WorkspaceMembersRequest) (*pb.WorkspaceMembersResponse, error) {

	members, err := g.WorkspaceUsecase.WorkspaceMembers(dto.WorkspaceMembersRequest{
		AccessToken: req.GetAccessToken(),
		ID:          uuid.MustParse(req.GetId()),
	})

	if err != nil {
		return &pb.WorkspaceMembersResponse{}, err
	}

	fmt.Println(members)

	return &pb.WorkspaceMembersResponse{}, nil
}
