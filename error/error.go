/*
 * @Date: 2021-11-09 10:58:12
 * @LastEditors: jaxiu
 * @LastEditTime: 2021-11-09 11:00:14
 * @FilePath: /TestCase/error/error.go
 */

package main

import (
	"errors"
	"log"
)

func main() {

	test1()

}

func test1() (err error) {

	defer func() {
		if err != nil {
			log.Println("defer Error:", err)
		}
	}()

	return test2()

}

func test2() (err error) {

	err = errors.New("test2 error")
	return err
}
