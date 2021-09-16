/*
 * @Date: 2021-09-15 14:32:04
 * @LastEditors: jaxiu
 * @LastEditTime: 2021-09-16 11:38:17
 * @FilePath: /TestCase/test/list_test.go
 */
package test

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestList(t *testing.T) {
	fmt.Println(List())
	list, _ := List()
	for i, v := range list {
		fmt.Println(i, *v)
	}
	t.Log(List())
}

func TestNumList(t *testing.T) {
	list := NumList()
	p := uintptr(unsafe.Pointer(&list[0]))
	// 模拟一个指针操作
	fmt.Printf("%d\n", *(*int)(unsafe.Pointer((uintptr(8*2) + p))))

}
