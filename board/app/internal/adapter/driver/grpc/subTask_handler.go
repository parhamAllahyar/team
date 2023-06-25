package grpc

import (
	//"admin/internal/core/service"
	pb "board/api/proto/grpcpkg"
	"board/internal/core/dto"
	"context"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
)

//TODO: Refactor grpc
// func (g *grpcHandler) CreateSubTask(ctx context.Context, req *pb.CreateSubTaskRequest) (*pb.CreateSubTaskResponse, error) {

// 	return &pb.CreateSubTaskResponse{
// 		Subtask: &pb.SubTaskItem{
// 			Id: 1,
// 			Subtask: &pb.SubTaskData{
// 				Title:  "aaaa",
// 				Order:  1,
// 				TaskId: 1,
// 			},
// 		},
// 	}, nil

// }
func (g *grpcHandler) DeleteSubTask(ctx context.Context, req *pb.DeleteSubTaskRequest) (*emptypb.Empty, error) {

	err := g.subTaskUsecase.DeleteSubTask(dto.DeleteSubTaskRequest{
		AccessToken: req.GetAccessToken(),
		TaskID:      uuid.MustParse(req.GetSubTaskId()),
	})

	if err != nil {
		return &emptypb.Empty{}, err
	}
	return nil, nil
}
