package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	demo()

	wg := sync.WaitGroup{}
	wg.Add(10)
	fmt.Println("开始执行...")
	for i := 0; i < 10; i++ {
		go func(i int) {
			time.Sleep(2 * time.Second)
			fmt.Println("我来啦", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("执行结束")
}

func demo() {
	var wg sync.WaitGroup
	wg.Add(2)

	// NewTimer 创建一个 Timer，它会在最少过去时间段 d 后到期，向其自身的 C 字段发送当时的时间
	timer := time.NewTimer(5 * time.Second)

	// NewTicker 返回一个新的 Ticker，该 Ticker 包含一个通道字段，并会每隔时间段 d 就向该通道发送当时的时间。它会调
	// 整时间间隔或者丢弃 tick 信息以适应反应慢的接收者。如果d <= 0会触发panic。关闭该 Ticker 可以释放相关资源。
	ticker := time.NewTicker(5 * time.Second)

	go func(t *time.Ticker) {
		defer wg.Done()
		for {
			<-t.C
			fmt.Println("get ticker", time.Now().Format("2006-01-02 15:04:05"))
		}
	}(ticker)

	go func(t *time.Timer) {
		defer wg.Done()
		for {
			<-t.C
			fmt.Println("get timer", time.Now().Format("2006-01-02 15:04:05"))
			// Reset 使 t 重新开始计时，（本方法返回后再）等待时间段 d 过去后到期。如果调用时 t 还在等待中会返回真；如果 t已经到期或者被停止了会返回假。
			t.Reset(5 * time.Second)
		}
	}(timer)

	wg.Wait()
}
