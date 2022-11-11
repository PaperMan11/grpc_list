package dao

import "log"

const userSchema = `CREATE TABLE IF NOT EXISTS user (
    user_id BIGINT PRIMARY KEY,
    user_name VARCHAR(100) NOT NULL,
    user_password VARCHAR(100) NOT NULL 
)ENGINE=InnoDB DEFAULT CHARSET=utf8;`

func (dao *DBClient) DBTableInit() error {
	if _, err := dao.db.Exec(userSchema); err != nil {
		log.Println("create user failed")
		return err
	}
	return nil
}
