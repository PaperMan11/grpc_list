package service

import (
	"errors"
	"strings"
	"user/internal/global"
	"user/internal/model"
	"user/internal/pb"
	"user/pkg/encrypt"

	"user/pkg/snowflake"
)

func UserLogin(req *pb.LoginRequest) (user model.User, err error) {
	// 查找
	var u model.User
	u, err = global.Db.FindById(req.UserId)
	if err != nil {
		return u, errors.New("user not exist")
	}
	// compare password
	if ok := encrypt.ComparePassword(req.UserPassword, u.PassWord); !ok {
		return model.User{}, errors.New("password invailed")
	}

	return u, nil
}

func UserRegister(req *pb.RegRequest) (user model.User, err error) {
	var u model.User
	// 比对密码
	if ok := compare(req.UserPassword, req.UserPassword2); !ok {
		return u, errors.New("twice password not equel")
	}
	// 生成user id
	u.UserId = snowflake.GetID()
	u.UserName = req.UserName
	// 加密
	u.PassWord, err = encrypt.GenerateFromPassword(req.UserPassword)
	if err != nil {
		return model.User{}, err
	}
	// 存入数据库
	err = global.Db.CreateUser(u)
	if err != nil {
		return model.User{}, err
	}
	return u, nil
}

func compare(password, password2 string) bool {
	return strings.Compare(password, password2) == 0
}
