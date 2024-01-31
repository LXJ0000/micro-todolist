package rpc

import (
	"context"
	"errors"
	"micro-todolist/idl/pb"
	"micro-todolist/pkg/e"
)

func UserLogin(ctx context.Context, req *pb.UserRequest) (resp *pb.UserResponse, err error) {
	resp, err = UserService.UserLogin(ctx, req)
	if err != nil {
		return
	}
	if resp.Code != e.CodeSuccess {
		err = errors.New(e.GetMsg(int(resp.Code)))
	}
	return
}

func UserRegister(ctx context.Context, req *pb.UserRequest) (resp *pb.UserResponse, err error) {
	return UserService.UserRegister(ctx, req)
}
