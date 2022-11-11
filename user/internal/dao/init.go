package dao

import (
	"fmt"
	"log"

	"user/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type DBClient struct {
	db *sqlx.DB
}

func NewDBClient() (*DBClient, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True", config.Conf.MysqlConf.UserName, config.Conf.MysqlConf.PassWord,
		config.Conf.MysqlConf.Host, config.Conf.MysqlConf.Port, config.Conf.MysqlConf.DataBase)
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Printf("connect db failed: %s", err)
		return nil, err
	}
	if err = db.Ping(); err != nil {
		log.Printf("db.Ping() failed: %s", err)
		return nil, err
	}
	return &DBClient{db: db}, nil
}
