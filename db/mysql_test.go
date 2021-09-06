/*
 * @Date: 2021-09-06 20:05:57
 * @LastEditors: jaxiu
 * @LastEditTime: 2021-09-06 21:01:03
 * @FilePath: /TestCase/db/mysql_test.go
 */

package db

import (
	"log"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
)

func Test_Mysql_TestCase(t *testing.T) {
	Init()
	TestName()
}

func TestDB(t *testing.T) {
	conn, sm, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
		return
	}
	columns := []string{"id", "name"}
	sm.ExpectQuery("SELECT VERSION").WillReturnRows(sqlmock.NewRows([]string{"version"}).FromCSVString("10.6.4-MariaDB-1:10.6.4+maria~focal"))
	sm.ExpectQuery("SELECT (.+) FROM .*").WillReturnRows(sqlmock.NewRows(columns).FromCSVString("1,yunfan01"))
	SetDB(conn)
	TestName()
}
