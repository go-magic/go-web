package db

import (
	"dat2/conf"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func init() {
	var err error
	config := conf.GetConfig()
	db, err = gorm.Open("mysql", config.Sql)
	if err != nil {
		panic(err)
		return
	}
}

func GetDb() *gorm.DB {
	return db
}

func Close() {
	db.Close()
}
