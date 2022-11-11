package model

type User struct {
	UserId   int64  `db:"user_id"`
	UserName string `db:"user_name"`
	PassWord string `db:"user_password"`
}
