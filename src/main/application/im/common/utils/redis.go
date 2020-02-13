package utils

import "github.com/garyburd/redigo/redis"

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

func ImRedisPool() *redis.Pool {
	return pool
}
