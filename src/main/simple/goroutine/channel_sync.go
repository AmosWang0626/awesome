package main

import "fmt"

// 读写同时执行
func main() {

	intChan := make(chan int, 50)
	exitChan := make(chan bool, 1)

	go writeChan(intChan)

	go readChan(intChan, exitChan)

	for {
		val, ok := <-exitChan
		if !ok {
			break
		}
		fmt.Printf("退出程序:%t ", val)
	}

}

func writeChan(intChan chan int) {
	for i := 0; i < 50; i++ {
		intChan <- i
		fmt.Printf("写入数据:%v \n", i)
	}
	close(intChan)
}

func readChan(intChan chan int, exitChan chan bool) {
	for {
		val, ok := <-intChan
		if !ok {
			break
		}
		fmt.Printf("读取到数据:%v \n", val)
	}

	// 退出Flag
	exitChan <- true
	close(exitChan)
}
