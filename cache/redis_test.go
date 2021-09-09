/*
 * @Date: 2021-09-06 21:08:19
 * @LastEditors: jaxiu
 * @LastEditTime: 2021-09-06 21:32:15
 * @FilePath: /TestCase/cache/redis_test.go
 */
package cache

import (
	"fmt"
	"testing"

	miniredis "github.com/alicebob/miniredis/v2"
	redis "github.com/garyburd/redigo/redis"
)

func TestCache(t *testing.T) {
	ProCache(Initredis())
}

func TestCacheMock(t *testing.T) {
	s, _ := miniredis.Run()
	conn, _ := redis.Dial("tcp", s.Addr())
	fmt.Println(s.Addr())
	ProCache(conn)

}
