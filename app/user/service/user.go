package service

import (
	"context"
	"errors"
	"micro-todolist/app/user/repository/dao"
	"micro-todolist/app/user/repository/model"
	"micro-todolist/idl/pb"
	"micro-todolist/pkg/e"
	"sync"
)

type UserSrv struct {
}

var (
	userSrvIns  *UserSrv
	userSrvOnce sync.Once
)

func NewUserSrv() *UserSrv {
	userSrvOnce.Do(func() {
		userSrvIns = &UserSrv{}
	})
	return userSrvIns
}

func (u *UserSrv) UserLogin(ctx context.Context, req *pb.UserRequest, resp *pb.UserResponse) (err error) {
	resp.Code = e.CodeSuccess
	//	用户是否存在
	user, err := dao.NewUserDao(ctx).FindUserByUserName(req.UserName)
	if err != nil {
		return
	}
	if user.ID == 0 {
		err = errors.New("用户不存在")
		resp.Code = e.CodeError
		return
	}
	if !user.CheckPassword(req.Password) {
		err = errors.New("用户密码有误")
		resp.Code = e.CodeError
		return
	}
	resp.UserInfo = BuildUser(user)
	return nil
}

func (u *UserSrv) UserRegister(ctx context.Context, req *pb.UserRequest, resp *pb.UserResponse) (err error) {
	resp.Code = e.CodeSuccess
	if req.Password != req.PasswordConfirm {
		err = errors.New("两次密码不一致")
		resp.Code = e.CodeError
		return
	}
	//	用户是否存在
	user, err := dao.NewUserDao(ctx).FindUserByUserName(req.UserName)
	if err != nil {
		return
	}
	if user.ID != 0 {
		err = errors.New("用户名已存在")
		resp.Code = e.CodeError
		return
	}
	user = &model.User{
		UserName: req.UserName,
	}
	if err = user.SetPassword(req.Password); err != nil {
		resp.Code = e.CodeError
		return
	}
	if err = dao.NewUserDao(ctx).CreateUser(user); err != nil {
		resp.Code = e.CodeError
		return
	}
	return
}

func BuildUser(item *model.User) *pb.UserModel {
	return &pb.UserModel{
		Id:        uint32(item.ID),
		UserName:  item.UserName,
		CreatedAt: item.CreatedAt.Unix(),
		UpdatedAt: item.UpdatedAt.Unix(),
	}
}
