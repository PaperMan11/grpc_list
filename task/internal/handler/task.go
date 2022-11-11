package handler

import (
	"context"
	"task/internal/pb"
	"task/internal/service"
	"user/pkg/e"
)

type TaskServer struct {
	pb.UnimplementedTaskServiceServer
}

var _ pb.TaskServiceServer = &TaskServer{}

func (TaskServer) TaskCreate(ctx context.Context, req *pb.TaskRequest) (resp *pb.CommonResponse, err error) {
	resp = new(pb.CommonResponse)
	resp.Code = e.SUCCESS
	err = service.TaskCreate(req)
	if err != nil {
		resp.Code = e.ERROR
		resp.Msg = e.GetMsg(uint(e.ERROR))
		resp.Data = err.Error()
		return resp, err
	}
	resp.Msg = e.GetMsg(uint(e.SUCCESS))
	return resp, nil
}

func (TaskServer) TaskUpdate(ctx context.Context, req *pb.TaskRequest) (resp *pb.CommonResponse, err error) {
	resp = new(pb.CommonResponse)
	resp.Code = e.SUCCESS
	err = service.TaskUpdate(req)
	if err != nil {
		resp.Code = e.ERROR
		resp.Msg = e.GetMsg(uint(e.ERROR))
		resp.Data = err.Error()
		return resp, err
	}
	resp.Msg = e.GetMsg(uint(e.SUCCESS))
	return resp, nil
}

func (TaskServer) TaskShow(ctx context.Context, req *pb.TaskRequest) (resp *pb.TaskDetailResponse, err error) {
	resp = new(pb.TaskDetailResponse)

	return resp, nil
}

func (TaskServer) TaskDelete(ctx context.Context, req *pb.TaskRequest) (resp *pb.CommonResponse, err error) {
	resp = new(pb.CommonResponse)

	return resp, nil
}
