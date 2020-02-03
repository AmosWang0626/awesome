package main

import (
	"fmt"
)

/*
下边代码是有问题的,如果协程没有执行完成,那么主线程也会执行完
并发: 交替执行
并行: 同时执行

协程MPG模式
1. M:操作系统主线程(物理线程); P:协程执行需要的上下文; G:协程
2. 如果主线程里有耗时操作, 线程需要阻塞等待, go 语言为了充分利用 cpu 的资源, 此时 go 语言就会先处理 P 上挂载的协程, 然后[并发-并行]继续执行
厨师做饭: 4个客人,A点了10个菜,BCD每人点2个菜; 先给A炒4个菜,先吃着,然后让他先等下; 给BCD都做一个;然后再给A做两个;再给BCD做完;再给A做完
*/

// 简单 [主程/协程] demo
func main() {
	go test()

	for i := 0; i < 10; i++ {
		fmt.Printf("main hello golang! %d\n", i)
		//time.Sleep(time.Second)
	}
}

func test() {
	for i := 0; i < 10; i++ {
		fmt.Printf("hello world! %d\n", i)
		//time.Sleep(time.Second)
	}
}
