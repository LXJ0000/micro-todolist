package dao

import (
	"context"
	"gorm.io/gorm"
	"micro-todolist/app/user/repository/model"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &UserDao{NewDBClient(ctx)}
}

func (dao *UserDao) FindUserByUserName(userName string) (user *model.User, err error) {
	err = dao.Model(&model.User{}).
		Where("user_name = ?", userName).First(&user).Error
	return
}

func (dao *UserDao) CreateUser(user *model.User) error {
	return dao.Model(&model.User{}).Create(&user).Error
}
