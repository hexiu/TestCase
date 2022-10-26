/*
 * @Date: 2021-11-15 13:38:18
 * @LastEditors: jaxiu
 * @LastEditTime: 2021-11-15 13:48:52
 * @FilePath: /TestCase/test/uuid_test.go
 */
package test

import (
	"fmt"
	"testing"
)

func TestUUIDTest(t *testing.T) {

	for i := 0; i < 1000; i++ {
		uid := UUID("connection")
		if len(uid) > 80 {
			fmt.Println(uid)
		}
	}
}

func Benchmark_UUID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UUID("connection")
	}
}
