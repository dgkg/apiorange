package sql

import (
	"apiorange/db/sqlite"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBSQL = sqlite.SQLite

func New(dsn string) *DBSQL {
	dsn = "root@tcp(127.0.0.1:3306)/orange?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	var sqldb DBSQL
	sqldb.SetDB(db)
	sqldb.InitModel()
	return &sqldb
}
