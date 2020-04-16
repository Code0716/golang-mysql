package db

import (
	"fmt"

	"../../constants"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//ConnectMySQL gormでworld databaseに接続する。
func ConnectMySQL(dbname string) *gorm.DB {
	DBMS := constants.DB
	USER := constants.DBUser
	PASS := constants.DBPassword
	PROTOCOL := constants.DBProtocol
	DBNAME := dbname
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	fmt.Println("db connected: ", &db)
	db.LogMode(true)
	// 勝手に複数形になるのを抑制
	db.SingularTable(true)

	return db
}
