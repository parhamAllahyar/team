package grpc

import (
	//"admin/internal/core/service"
	pb "board/api/proto/grpcpkg"
	"board/internal/core/dto"
	"context"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"

)

func (g *grpcHandler) CreateTask(ctx context.Context, req *pb.CreateTaskRequest) (*pb.CreateTaskResponse, error) {

	newTask, err := g.taskUsecase.CreateTask(
		dto.CreateTaskRequest{
			AccessToken: req.GetAccessToken(),
			Title:       req.GetTask().Title,
			Order:       uint(req.GetTask().Order),
			SectionID:   uuid.MustParse(req.GetTask().SectionId),
		},
	)

	if err != nil {
		return nil, err
	}

	return &pb.CreateTaskResponse{
		Task: &pb.TaskItem{
			Title:     newTask.Title,
			Order:     uint32(newTask.Order),
			SectionId: newTask.SectionID.String(),
		},
	}, nil
}
func (g *grpcHandler) DeleteTask(ctx context.Context, req *pb.DeleteTaskRequest) (*emptypb.Empty, error) {
	err := g.taskUsecase.DeleteTask(
		dto.DeleteTaskRequest{
			AccessToken: req.GetAccessToken(),
			TaskID:      uuid.MustParse(req.GetTaskId()),
		},
	)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
func (g *grpcHandler) UpdateTask(ctx context.Context, req *pb.UpdateTaskRequest) (*pb.UpdateTaskResponse, error) {

	task, err := g.taskUsecase.UpdateTask(
		dto.UpdateTaskRequest{
			AccessToken: req.GetAccessToken(),
			ID:          uuid.MustParse(req.GetTaskId()),
			Title:       req.GetTask().Title,
			SectionID:   uuid.MustParse(req.GetTask().SectionId),
			Order:       uint(req.GetTask().Order),
		},
	)

	if err != nil {
		return nil, err
	}

	return &pb.UpdateTaskResponse{
		Task: &pb.TaskItem{
			Title:     task.Title,
			Order:     uint32(task.Order),
			SectionId: task.SectionID.String(),
		},
	}, nil

}
func (g *grpcHandler) AssignTask(ctx context.Context, req *pb.AssignTaskRequest) (*emptypb.Empty, error) {
	err := g.taskUsecase.AssignTask(dto.AssignTaskRequest{
		AccessToken: req.GetAccessToken(),
		TaskID:      uuid.MustParse(req.GetTaskId()),
		CustomerID:  uuid.MustParse(req.GetUserId()),
	})
	if err != nil {
		return &emptypb.Empty{}, err
	}
	return &emptypb.Empty{}, nil
}
