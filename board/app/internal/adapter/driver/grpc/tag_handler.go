package grpc

import (
	pb "board/api/proto/grpcpkg"
	"board/internal/core/dto"
	"context"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"

)

func (g *grpcHandler) CreateTag(ctx context.Context, req *pb.CreateTagRequest) (*pb.CreateTagResponse, error) {

	newTag, err := g.tagUsecase.CreateTag(dto.CreateTagRequest{
		AccessToken: req.GetAccessToken(),
		Title:       req.GetTag().Title,
	})

	if err != nil {
		return nil, err
	}

	return &pb.CreateTagResponse{
		Tag: &pb.TagItem{
			Id: newTag.ID.String(),
			Tag: &pb.TagData{
				Title: newTag.Title,
				Order: uint32(newTag.Order),
			},
		},
	}, nil

}
func (g *grpcHandler) DeleteTag(ctx context.Context, req *pb.DeleteTagRequest) (*emptypb.Empty, error) {

	err := g.tagUsecase.DeleteTag(dto.DeleteTagRequest{})

	if err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}
func (g *grpcHandler) UpdateTag(ctx context.Context, req *pb.UpdateTagRequest) (*pb.UpdateTagResponse, error) {

	tag, err := g.tagUsecase.UpdateTag(
		dto.UpdateTagRequest{
			AccessToken: req.GetAccessToken(),
			ID:          uuid.MustParse(req.GetTag().Id),
			Title:       req.GetTag().GetTag().Title,
			Order:       uint(req.GetTag().GetTag().Order),
		},
	)

	if err != nil {
		return nil, err
	}

	return &pb.UpdateTagResponse{
		Tag: &pb.TagItem{
			Id: tag.ID.String(),
			Tag: &pb.TagData{
				Title: tag.Title,
				Order: uint32(tag.Order),
			},
		},
	}, nil
}
