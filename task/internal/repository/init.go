package repository

import (
	"fmt"
	"log"
	"task/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

func Init() (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True", config.Conf.MysqlConf.UserName, config.Conf.MysqlConf.PassWord,
		config.Conf.MysqlConf.Host, config.Conf.MysqlConf.Port, config.Conf.MysqlConf.DataBase)
	Db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Printf("connect db failed: %s", err)
		return err
	}
	if err = Db.Ping(); err != nil {
		log.Printf("db.Ping() failed: %s", err)
		return err
	}
	Db.SetMaxOpenConns(20)
	Db.SetMaxIdleConns(10)
	if err = migration(); err != nil {
		log.Println("migration failed")
		return err
	}

	// create table
	migration()
	return nil
}
