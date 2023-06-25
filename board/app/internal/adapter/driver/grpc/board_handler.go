package grpc

import (
	pb "board/api/proto/grpcpkg"
	"board/internal/core/dto"
	"context"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (g *grpcHandler) CreateBoard(ctx context.Context, req *pb.CreateBoardRequest) (*pb.CreateBoardResponse, error) {

	newBoard, err := g.boardUsecase.CreateBoard(dto.CreateBoardRequest{
		AccessToken: req.GetAccessToken(),
		Title:       req.GetBoard().Title,
		WorkspaceID: uuid.MustParse(req.GetBoard().WorksoaceId),
	})

	if err != nil {
		return nil, err
	}

	return &pb.CreateBoardResponse{
		Board: &pb.BoardItem{
			Id: newBoard.ID.String(),
			BoardData: &pb.BoardData{
				Title:       newBoard.Title,
				WorksoaceId: newBoard.WorkspaceID.String(),
			},
		},
	}, nil

}
func (g *grpcHandler) DeleteBoard(ctx context.Context, req *pb.DeleteBoardRequest) (*emptypb.Empty, error) {

	err := g.boardUsecase.DeleteBoard(dto.DeleteBoardRequest{
		AccessToken: req.GetAccessToken(),
		ID:          uuid.MustParse(req.GetBoardId()),
	})

	if err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}

func (g *grpcHandler) UpdateBoard(ctx context.Context, req *pb.UpdateBoardRequest) (*pb.UpdateBoardResponse, error) {

	board, err := g.boardUsecase.UpdateBoard(dto.UpdateBoardRequest{
		AccessToken: req.GetAccessToken(),
		ID:          uuid.MustParse(req.GetBoard().Id),
		Title:       req.GetBoard().GetBoardData().Title,
		WorkspaceID: uuid.MustParse(req.GetBoard().GetBoardData().WorksoaceId),
	})

	if err != nil {
		return nil, err
	}

	return &pb.UpdateBoardResponse{
		Board: &pb.BoardItem{
			Id: board.ID.String(),
			BoardData: &pb.BoardData{
				Title:       board.Title,
				WorksoaceId: board.WorkspaceId.String(),
			},
		},
	}, nil

}
