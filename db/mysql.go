package db

import (
	"database/sql"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func initMysql() *sql.DB {
	dsn := "root:root@tcp(127.0.0.1:3306)/ltest?charset=utf8mb4&parseTime=True&loc=Local"

	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Panicln(err)
		return nil
	}
	return conn
}

// 模拟数据库db测试
func InitDB(conn *sql.DB) {
	var err error
	db, err = gorm.Open(mysql.New(mysql.Config{Conn: conn}), &gorm.Config{})
	if err != nil {
		log.Panic(err)
		return
	}
}

func SetDB(conn *sql.DB) {
	var err error

	db, err = gorm.Open(mysql.New(mysql.Config{Conn: conn}), &gorm.Config{})
	if err != nil {
		log.Panic(err)
		return
	}
}

func TestName() {
	type Name struct {
		Id   int64
		Name string
	}

	nameList := make([]*Name, 0)
	err := db.Table("name").Where("id = 1").Find(&nameList).Error
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(nameList[0])
}

func Init() {
	InitDB(initMysql())
}
