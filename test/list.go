/*
 * @Date: 2021-09-15 14:32:26
 * @LastEditors: jaxiu
 * @LastEditTime: 2021-09-15 18:17:39
 * @FilePath: /TestCase/test/test.go
 */
package test

import "fmt"

func List() (list []*string, err error) {
	// A := "a"
	// list = append(list, &A)
	var a []*string
	fmt.Println(a)
	return
}
func NumList() (list []int) {
	list = []int{1, 2, 3, 4, 5, 6}
	return
}
