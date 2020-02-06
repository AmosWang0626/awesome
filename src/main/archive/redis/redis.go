package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
	"time"
)

// go get "github.com/garyburd/redigo/redis"
func main() {
	conn, err := redis.Dial("tcp", "amos.wang:1998", redis.DialDatabase(6), redis.DialConnectTimeout(10*time.Second))
	if err != nil {
		log.Fatalln("redis dial err", err)
	}
	defer conn.Close()

	// del set get
	replay, _ := conn.Do("del", "simple_name")
	fmt.Println("del simple_name", replay)

	replay, _ = conn.Do("set", "simple_name", "amos.wang")
	fmt.Println("set simple_name", replay)

	getStr, _ := redis.String(conn.Do("get", "simple_name"))

	fmt.Println("get simple_name", getStr)
}
