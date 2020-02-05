package main

import (
	"fmt"
	"sync"
)

var (
	mapArr = make(map[int]uint64)
	// sync.Mutex // 互斥锁
	// sync.RWMutex // 读写互斥锁
	lock sync.Mutex
)

func factorial(n int) {
	res := n
	for i := 1; i < n; i++ {
		res += i
	}
	lock.Lock()
	mapArr[n] = uint64(res)
	lock.Unlock()
}

func main() {
	for i := 1; i <= 200; i++ {
		go factorial(i)
	}

	lock.Lock()
	for index, value := range mapArr {
		fmt.Printf("map[%d:%d]\n", index, value)
	}
	lock.Unlock()
}

// 1. fatal error: concurrent map writes [ERROR :: Map并发写操作错误]
// 2. 使用race: go build -race sync_lock.go [ERROR :: race[竞争] 观察并发异常]
