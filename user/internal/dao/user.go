package dao

import (
	"log"
	"user/internal/model"
	"user/internal/pb"
)

// 根据id查找
func (dao *DBClient) FindById(user_id int64) (model.User, error) {
	var (
		user model.User
		err  error
	)
	sqlStr := `SELECT user_id,user_name,user_password
	FROM user WHERE user_id=?`
	err = dao.db.Get(&user, sqlStr, user_id)
	if err != nil {
		log.Println("dao.db.Get failed")
	}
	return user, err
}

func (dao *DBClient) CreateUser(user model.User) error {
	sqlStr := `INSERT INTO user(user_id, user_name, user_password) VALUES (?,?,?)`
	_, err := dao.db.Exec(sqlStr, user.UserId, user.UserName, user.PassWord)
	if err != nil {
		return err
	}
	return nil
}

func BuildUserDetail(user model.User) *pb.UserModel {
	return &pb.UserModel{
		UserId:       user.UserId,
		UserName:     user.UserName,
		UserPassword: user.PassWord,
	}
}
