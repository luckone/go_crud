package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"test/config"
	"log"
	"test/db/models/book"
)

var Connection *gorm.DB

func Connect() *gorm.DB {
	mysqlUser := config.Config.MySQL.User
	mysqlPassword := config.Config.MySQL.Password
	mysqlDatabase := config.Config.MySQL.Database
	mysqlHost := config.Config.MySQL.Host
	connectionParams := mysqlUser + ":" + mysqlPassword + "@tcp(" + mysqlHost + ")" + "/" + mysqlDatabase + "?charset=utf8"

	establishedConnection, err := gorm.Open("mysql", connectionParams)

	if (err != nil) {
		log.Fatal("Mysql connection error ", err)
	}

	Connection = establishedConnection
	Connection.AutoMigrate(book.BookModel{})
	return establishedConnection
}