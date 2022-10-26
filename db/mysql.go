package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var dsn = "root:root@tcp(127.0.0.1:3306)/ltest?charset=utf8mb4&parseTime=True&loc=Local"

func initMysql() *sql.DB {

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

// sqlx
func initMysqlSqlx() {
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(db.Exec("insert into name (id, name) values (?, ?)", 2, "hjh"))

	rows, err := db.Query("select * from name")
	if err != nil {
		log.Println(err)
		return
	}
	for rows.Next() {
		var x int
		var y string
		rows.Scan(&x, &y)
		fmt.Println(x, y)
	}

	type Name struct {
		Id   int
		Name string
	}
	var nameList []Name
	fmt.Println(db.Select(&nameList, "select * from name"))
	fmt.Printf("%+v\n", nameList)

}
