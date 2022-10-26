/*
 * @Date: 2021-09-17 11:14:23
 * @LastEditors: jaxiu
 * @LastEditTime: 2022-01-11 11:46:08
 * @FilePath: /TestCase/test/map.go
 */
package test

import (
	"fmt"
	"sort"
	"sync"
	"time"
)

func Map() {
	m := make(map[string]interface{})
	m1 := make(map[string]map[string]interface{})
	mu := sync.Mutex{}
	wg := sync.WaitGroup{}
	start := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		func(i int) {
			go func(i int) {
				is := fmt.Sprintf("a-%d", i)
				m = map[string]interface{}{is: i}
				mu.Lock()
				m1[is] = m
				mu.Unlock()
				wg.Done()
				time.Sleep(1 * time.Second)
			}(i)
		}(i)
	}
	wg.Wait()
	fmt.Println(m1)
	fmt.Println(len(m1), time.Now().Sub(start))

	a := map[int64]int64{
		1: 1,
		2: 2,
		3: 4,
	}
	a1 := map[int64]int64{
		1: 1,
		6: 2,
		7: 4,
	}
	b := make(map[int64]int64)
	for k, v := range a {
		b[k] += v
	}
	for k, v := range a1 {
		b[k] += v
	}
	fmt.Println(b)
sort.Sort()
}
