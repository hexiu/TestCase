/*
 * @Date: 2021-09-17 14:37:05
 * @LastEditors: jaxiu
 * @LastEditTime: 2021-11-15 13:45:28
 * @FilePath: /TestCase/test/page_test.go
 */
package test

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestPage(t *testing.T) {
	wg := sync.WaitGroup{}
	for j := 0; j < 10; j++ {
		wg.Add(1)
		go func(j int) {
			var link []int
			x := rand.Intn(int(time.Now().Unix() % 1000))
			for i := 0; i < x; i++ {
				link = append(link, i)
			}

			var next = true
			var list []int
			var k = 0
			for next {
				k++
				list, next = Page(link, k)
				if !next {
					fmt.Println("x:", x, k, j, next, list)
				}
			}

			wg.Done()
		}(j)
	}
	wg.Wait()
}
