package db

import (
	"fmt"
	"github.com/PlagueCat-Miao/goipfs-lab511/constdef"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func InitDataBase() error {
	db, err := gorm.Open("mysql", constdef.FileInfoMySqlUrl)
	if err != nil {
		return fmt.Errorf("[gorm-Open-err]:%v,url:%v", err, constdef.FileInfoMySqlUrl)
	}
	db.SingularTable(true) //让grom转义struct名字的时候不用加上s//https://www.jb51.net/article/151051.htm
	FileInfoDB = db
	return nil
}
