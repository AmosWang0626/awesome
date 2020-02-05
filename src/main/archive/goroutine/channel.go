package main

import (
	"fmt"
)

type cat struct {
	name string
	desc string
}

func main() {

	//intChannel()
	//mapChannel()
	//catChannel()
	interfaceChannel()

	// channel 默认情况下是可写和可读的
	// 声明只写
	var c1 = make(chan<- int, 1)
	c1 <- 11
	// 声明只读
	//var c2 = make(<-chan int, 1)
	//c2 <- 11

}

func interfaceChannel() {
	var catChan chan interface{}
	catChan = make(chan interface{}, 3)
	catChan <- cat{"aa", "desc"}
	catChan <- map[string]string{"111": "amos"}
	catChan <- 211
	fmt.Printf("读取数据: %v\n", (<-catChan).(cat).name)
	fmt.Printf("读取数据: %v\n", <-catChan)
	fmt.Printf("读取数据: %v\n", <-catChan)
}

func catChannel() {
	var catChan chan cat
	catChan = make(chan cat, 3)
	catChan <- cat{"aa", "desc"}
	catChan <- cat{"bb", "desc"}
	catChan <- cat{"cc", "desc"}
	// 如果 channel 没有关闭, 遍历时会抛 deadlock 异常; 反之正常
	close(catChan)
	// channel关闭后, channel是只读不可写的, 会报异常 send on closed channel
	//catChan <- cat{"dd", "desc"}
	// channel 遍历不能使用普通的 for 循环
	for c := range catChan {
		fmt.Printf("读取数据: %v\n", c)
	}
}

func mapChannel() {
	var mapChan chan map[string]string
	mapChan = make(chan map[string]string, 3)
	map1 := map[string]string{"111": "amos"}
	map2 := map[string]string{"222": "grace"}
	map3 := map[string]string{"333": "smith"}
	mapChan <- map1
	mapChan <- map2
	mapChan <- map3

	fmt.Printf("mapChan len=%d, cap=%d\n", len(mapChan), cap(mapChan))
	fmt.Printf("读取数据: %v\n", <-mapChan)
	fmt.Printf("读取数据: %v\n", <-mapChan)
	fmt.Printf("读取数据: %v\n", <-mapChan)
}

func intChannel() {
	// 初始化 [defer栈(先进后出); chan队列(先进先出)]
	var intChan chan int
	intChan = make(chan int, 3)
	fmt.Println(intChan, &intChan)

	// 写入数据
	intChan <- 20
	num := 211
	intChan <- num

	fmt.Printf("intChan len=%d, cap=%d\n", len(intChan), cap(intChan))

	// 读取数据
	var num2 int
	num2 = <-intChan
	fmt.Println("读取数据: ", num2)
	num2 = <-intChan
	fmt.Println("读取数据: ", num2)
	num2 = <-intChan
	fmt.Println("读取数据: ", num2)
	fmt.Printf("intChan len=%d, cap=%d\n", len(intChan), cap(intChan))
}
