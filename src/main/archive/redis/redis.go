package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

// 连接池配置
var pool *redis.Pool

func init() {
	pool = &redis.Pool{
		Dial: func() (conn redis.Conn, err error) {
			return redis.Dial("tcp", "amos.wang:1998", redis.DialDatabase(6))
		},
		TestOnBorrow: nil,   //
		MaxIdle:      8,     // 最大空闲数
		MaxActive:    0,     // 最大连接数, 0为无限制
		IdleTimeout:  100,   // 最大空闲时间
		Wait:         false, //
	}
}

// go get "github.com/garyburd/redigo/redis"
func main() {
	conn := pool.Get()
	defer conn.Close()
	defer pool.Close()

	/*
		1.string
		2.hash
		3.list
		4.sorted-set
	*/

	// string >>>>>>>>>>>> del set get expire
	replay, _ := conn.Do("del", "simple_name")
	fmt.Println("del simple_name", replay)

	replay, _ = conn.Do("set", "simple_name", "amos.wang")
	fmt.Println("set simple_name", replay)

	getStr, _ := redis.String(conn.Do("get", "simple_name"))
	fmt.Println("get simple_name", getStr)

	replay, _ = conn.Do("expire", "simple_name", 120)
	fmt.Println("expire simple_name 120s", replay)
	fmt.Println()

	// hash >>>>>>>>>>>> hset hget hmget
	replay, _ = conn.Do("del", "user1")
	fmt.Println("hash del user1", replay)

	replay, _ = conn.Do("hset", "user1", "name", "amos", "age", 18, "email", "admin@amos.com")
	fmt.Println("hash hset user1", replay)

	replay, _ = redis.String(conn.Do("hget", "user1", "name"))
	fmt.Println("hash hget user1 name", replay)

	replay, _ = redis.Strings(conn.Do("hmget", "user1", "name", "age", "email"))
	fmt.Println("hash hmget user1 name age email", replay)
	fmt.Println()

	// list >>>>>>>>>>>> lpush rpush lrange lpop rpop
	replay, _ = conn.Do("del", "list_number")
	fmt.Println("list del list_number", replay)

	replay, _ = conn.Do("lpush", "list_number", "1001", "1002", "1003", "1004", "1005")
	fmt.Println("list lpush list_number", replay)

	replay, _ = conn.Do("rpush", "list_number", "1006", "1007", "1008")
	fmt.Println("list rpush list_number", replay)

	replay, _ = redis.Strings(conn.Do("lrange", "list_number", 0, -1))
	fmt.Println("list lpush list_number all", replay)

	replay, _ = redis.String(conn.Do("lpop", "list_number"))
	fmt.Println("list lpop list_number", replay)

	replay, _ = redis.String(conn.Do("rpop", "list_number"))
	fmt.Println("list rpop list_number", replay)

	replay, _ = redis.Strings(conn.Do("lrange", "list_number", 0, -1))
	fmt.Println("list lrange list_number all", replay)
	fmt.Println()

	// sorted-set >>>>>>>>>>>> zadd zrange zrevrange
	replay, _ = conn.Do("del", "sort_number")
	fmt.Println("sorted-set del sort_number", replay)

	replay, _ = conn.Do("zadd", "sort_number", 5, "five", 2, "two", 4, "four", 1, "one", 3, "three")
	fmt.Println("sorted-set zadd sort_number", replay)

	replay, _ = redis.Strings(conn.Do("zrange", "sort_number", 0, -1))
	fmt.Println("sorted-set zrange sort_number all", replay)

	replay, _ = redis.Strings(conn.Do("zrange", "sort_number", 0, -1, "WITHSCORES"))
	fmt.Println("sorted-set zrange sort_number all WITHSCORES", replay)

	replay, _ = redis.Strings(conn.Do("zrevrange", "sort_number", 0, -1, "WITHSCORES"))
	fmt.Println("sorted-set zrevrange sort_number all WITHSCORES", replay)

}
