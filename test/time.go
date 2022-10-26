/*
 * @Date: 2021-09-27 10:58:01
 * @LastEditors: jaxiu
 * @LastEditTime: 2021-09-27 11:18:48
 * @FilePath: /TestCase/test/time.go
 */
package test

import (
	"fmt"
	"time"
)

func Time() {
	fmt.Println(time.Now())
	t := time.Now()
	t1, _ := time.Parse("2006-01-02", t.Format("2006-01-02"))
	fmt.Println(86400-8*3600-t.Unix()+t1.Unix(), t1, t1.Unix(), t.Unix(), (t.Unix() - t1.Unix()), 3600*24)
	fmt.Println((25-t.Hour())*3600, t.Hour())
	fmt.Println(86400 - t.Hour()*3600 - t.Minute()*60 - t.Second())
}
