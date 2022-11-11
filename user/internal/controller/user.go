package controller

import (
	"context"
	"user/internal/dao"
	"user/internal/model"
	pb "user/internal/pb"
	"user/internal/service"
	"user/pkg/e"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
}

var _ pb.UserServiceServer = UserService{}

func (UserService) UserLogin(ctx context.Context, req *pb.LoginRequest) (resp *pb.UserDetalResponse, err error) {
	resp = new(pb.UserDetalResponse)
	var user model.User
	user, err = service.UserLogin(req)
	if err != nil {
		resp.Code = e.ERROR
		return resp, err
	}
	resp.UserDetail = dao.BuildUserDetail(user)

	resp.Code = e.SUCCESS
	return resp, nil
}

func (UserService) UserRegister(ctx context.Context, req *pb.RegRequest) (resp *pb.UserDetalResponse, err error) {
	resp = new(pb.UserDetalResponse)
	var user model.User
	user, err = service.UserRegister(req)
	if err != nil {
		resp.Code = e.ERROR
		return resp, err
	}
	resp.Code = e.SUCCESS
	resp.UserDetail = dao.BuildUserDetail(user)
	return resp, nil
}
