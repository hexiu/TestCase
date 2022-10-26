/*
 * @Date: 2021-11-23 16:09:05
 * @LastEditors: jaxiu
 * @LastEditTime: 2021-11-23 16:11:07
 * @FilePath: /TestCase/test/struct.go
 */
package test

import (
	"encoding/json"
	"fmt"
)

type A struct {
	B string
	C *C
}

type C struct {
	D string
}

func TestStruct() {
	a := &A{B: "aaaa", C: &C{D: "bbb"}}
	b, _ := json.Marshal(a)
	fmt.Println(string(b))
}
