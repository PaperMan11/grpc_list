package global

import (
	"log"
	"user/internal/dao"
)

var (
	Db *dao.DBClient
)

func Init() (err error) {
	if Db, err = dao.NewDBClient(); err != nil {
		log.Println("NewDBClient failed")
		return err
	}
	if err = Db.DBTableInit(); err != nil {
		log.Println("DBTableInit failed")
	}
	return err
}
