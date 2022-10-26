package cache

import (
	"fmt"
	"log"

	redis "github.com/garyburd/redigo/redis"
)

var conn redis.Conn

func Initredis() redis.Conn {
	var err error
	conn, err = redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		log.Println(err)
		return nil
	}

	return conn
}

func ProCache(conn redis.Conn) {
	var err error
	defer conn.Close()
	_, err = conn.Do("SET", "name", "wd")
	if err != nil {
		fmt.Println("redis set error:", err)
	}
	name, err := redis.String(conn.Do("GET", "name"))
	if err != nil {
		fmt.Println("redis get error:", err)
	} else {
		fmt.Printf("Got name: %s \n", name)
	}

	if err = conn.Send("set", "a", "bb"); err != nil {
		log.Println(err)
	}
	if err = conn.Send("expire", "a", 100); err != nil {
		log.Println(err)
	}

	conn.Flush()
	log.Println( conn.Receive())
}
