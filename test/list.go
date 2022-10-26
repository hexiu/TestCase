/*
 * @Date: 2021-09-15 14:32:26
 * @LastEditors: jaxiu
 * @LastEditTime: 2021-09-17 11:53:58
 * @FilePath: /TestCase/test/list.go
 */
package test

import "fmt"

func List() (list []*string, err error) {
	// A := "a"
	// list = append(list, &A)
	var a []*string
	fmt.Println(a)
	var uids []int
	fmt.Println(len(uids), "ok")
	return
}
func NumList() (list []int) {
	list = []int{1, 2, 3, 4, 5, 6}
	fmt.Println(list[len(list):])
	return
}
