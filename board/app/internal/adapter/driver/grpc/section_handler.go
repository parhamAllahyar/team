package grpc

import (
	pb "board/api/proto/grpcpkg"
	"board/internal/core/dto"
	"context"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"

)

func (g *grpcHandler) CreateSection(ctx context.Context, req *pb.CreateSectionRequest) (*pb.CreateSectionResponse, error) {

	section, err := g.sectionUsecase.CreateSection(dto.CreateSectionRequest{

		AccessToken: req.GetAccessToken(),
		BoardID:     uuid.MustParse(req.GetSection().BoardId),
		Title:       req.Section.GetTitle(),
		Order:       uint(req.GetSection().Order),
	})

	if err != nil {
		return nil, err
	}

	return &pb.CreateSectionResponse{
		Section: &pb.SectionItem{
			Id: section.ID.String(),
			SectionData: &pb.SectionData{
				Title:   section.Title,
				Order:   uint32(section.Order),
				BoardId: section.BoardID.String(),
			},
		}}, nil

}
func (g *grpcHandler) DeleteSection(ctx context.Context, req *pb.DeleteSectionRequest) (*emptypb.Empty, error) {

	err := g.sectionUsecase.DeleteSection(dto.DeleteSectionRequest{
		AccessToken: req.GetAccessToken(),
		SectionID:   uuid.MustParse(req.GetSectionId()),
	})

	if err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}
func (g *grpcHandler) UpdateSection(ctx context.Context, req *pb.UpdateSectionRequest) (*pb.UpdateSectionResponse, error) {

	section, err := g.sectionUsecase.UpdateSection(dto.UpdateSectionRequest{
		AccessToken: req.GetAccessToken(),
		BoardID:     uuid.MustParse(req.GetSection().BoardId),
		Title:       req.GetSection().Title,
		Order:       uint(req.GetSection().Order),
		ID:          uuid.MustParse(req.GetSectionId()),
	})

	if err != nil {
		return nil, err
	}

	return &pb.UpdateSectionResponse{
		Section: &pb.SectionItem{
			Id: section.ID.String(),
			SectionData: &pb.SectionData{
				Title:   section.Title,
				BoardId: section.BoardID.String(),
				Order:   uint32(section.Order),
			},
		},
	}, nil

}
