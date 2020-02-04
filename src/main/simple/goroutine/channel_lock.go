package main

import (
	"fmt"
	"time"
)

var (
	numChan = make(chan int, 200)
	sumChan = make(chan map[int]uint64, 200)
)

func batchSum(num int) {
	res := num
	for i := 1; i < num; i++ {
		res += i
	}
	time.Sleep(50 * time.Millisecond)
	sumChan <- map[int]uint64{num: uint64(res)}
	fmt.Println("写入数据:", num, uint64(res))
}

func main() {
	for i := 1; i <= 200; i++ {
		numChan <- i
	}
	close(numChan)

	for {
		num, ok := <-numChan
		if !ok {
			close(sumChan)
			break
		}
		batchSum(num)
	}

	for item := range sumChan {
		fmt.Println("读取数据:", item)
	}
}
